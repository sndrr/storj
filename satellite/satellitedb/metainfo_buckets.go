// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package satellitedb

import (
	"context"
	"errors"

	"github.com/skyrings/skyring-common/tools/uuid"
	"github.com/zeebo/errs"

	"storj.io/storj/pkg/storj"
	"storj.io/storj/satellite/metainfo"
	dbx "storj.io/storj/satellite/satellitedb/dbx"
)

var _ metainfo.Buckets = (*buckets)(nil)

// implementation of metainfo.Buckets interface repository using spacemonkeygo/dbx orm
type buckets struct {
	db dbx.Methods
}

func (buckets *buckets) Create(ctx context.Context, bucket *metainfo.Bucket) error {
	_, err := buckets.db.Create_Bucket(ctx,
		dbx.Bucket_Id(bucket.ProjectID[:]),

		dbx.Bucket_ProjectId(bucket.ProjectID[:]),
		dbx.Bucket_Name([]byte(bucket.Name)),
		dbx.Bucket_PathCipher(int(bucket.PathCipher)),

		dbx.Bucket_CreatedAt(bucket.CreatedAt.UTC()),

		dbx.Bucket_DefaultSegmentSize(int(bucket.DefaultSegmentSize)),

		dbx.Bucket_DefaultEncryptionCipherSuite(int(bucket.DefaultEncryption.CipherSuite)),
		dbx.Bucket_DefaultEncryptionBlockSize(int(bucket.DefaultEncryption.BlockSize)),

		dbx.Bucket_DefaultRedundancyAlgorithm(int(bucket.DefaultRedundancy.Algorithm)),
		dbx.Bucket_DefaultRedundancyShareSize(int(bucket.DefaultRedundancy.ShareSize)),
		dbx.Bucket_DefaultRedundancyRequiredShares(int(bucket.DefaultRedundancy.RequiredShares)),
		dbx.Bucket_DefaultRedundancyRepairShares(int(bucket.DefaultRedundancy.RepairShares)),
		dbx.Bucket_DefaultRedundancyOptimalShares(int(bucket.DefaultRedundancy.OptimalShares)),
		dbx.Bucket_DefaultRedundancyTotalShares(int(bucket.DefaultRedundancy.TotalShares)),

		dbx.Bucket_Create_Fields{
			// AttributionID uuid.UUID, optional
		},
	)
	return err
}

func (buckets *buckets) Get(ctx context.Context, projectID uuid.UUID, name string) (*metainfo.Bucket, error) {
	dbxBucket, err := buckets.db.Get_Bucket_By_ProjectId_And_Name(ctx,
		dbx.Bucket_ProjectId(projectID[:]),
		dbx.Bucket_Name([]byte(name)),
	)
	if err != nil {
		return nil, err
	}
	return bucketFromDBX(dbxBucket)
}

func (buckets *buckets) Delete(ctx context.Context, projectID uuid.UUID, name string) error {
	_, err := buckets.db.Delete_Bucket_By_ProjectId_And_Name(ctx,
		dbx.Bucket_ProjectId(projectID[:]),
		dbx.Bucket_Name([]byte(name)),
	)
	return err
}

func (buckets *buckets) List(ctx context.Context, projectID uuid.UUID, opts metainfo.BucketListOptions) (metainfo.BucketList, error) {
	// TODO: add sanity checks

	limit := opts.Limit + 1 // add one to detect More

	var dbxBuckets []*dbx.Bucket
	var err error

	switch opts.Direction {
	case storj.Before:
		dbxBuckets, err = buckets.db.Limited_Bucket_By_ProjectId_And_Name_Less_OrderBy_Asc_Name(ctx,
			dbx.Bucket_ProjectId(projectID[:]),
			dbx.Bucket_Name([]byte(opts.Cursor)),
			0, int64(limit),
		)
	case storj.Backward:
		dbxBuckets, err = buckets.db.Limited_Bucket_By_ProjectId_And_Name_LessOrEqual_OrderBy_Asc_Name(ctx,
			dbx.Bucket_ProjectId(projectID[:]),
			dbx.Bucket_Name([]byte(opts.Cursor)),
			0, int64(limit),
		)
	case storj.After:
		dbxBuckets, err = buckets.db.Limited_Bucket_By_ProjectId_And_Name_Greater_OrderBy_Asc_Name(ctx,
			dbx.Bucket_ProjectId(projectID[:]),
			dbx.Bucket_Name([]byte(opts.Cursor)),
			0, int64(limit),
		)
	case storj.Forward:
		dbxBuckets, err = buckets.db.Limited_Bucket_By_ProjectId_And_Name_GreaterOrEqual_OrderBy_Asc_Name(ctx,
			dbx.Bucket_ProjectId(projectID[:]),
			dbx.Bucket_Name([]byte(opts.Cursor)),
			0, int64(limit),
		)
	default:
		return metainfo.BucketList{}, errors.New("unknown list direction")
	}

	if err != nil {
		return metainfo.BucketList{}, err
	}

	var result metainfo.BucketList
	result.More = len(dbxBuckets) == limit

	// cut the extra element
	if result.More {
		switch opts.Direction {
		case storj.Before, storj.Backward:
			dbxBuckets = dbxBuckets[:len(dbxBuckets)-1]
		case storj.After, storj.Forward:
			dbxBuckets = dbxBuckets[0:]
		default:
			return metainfo.BucketList{}, errors.New("unknown list direction")
		}
	}

	result.Items = make([]*metainfo.Bucket, len(dbxBuckets))

	for i, dbxBucket := range dbxBuckets {
		bucket, err := bucketFromDBX(dbxBucket)
		if err != nil {
			return metainfo.BucketList{}, err
		}
		result.Items[i] = bucket
	}

	return result, nil
}

// bucketFromDBX is used for creating Project entity from autogenerated dbx.Project struct
func bucketFromDBX(bucket *dbx.Bucket) (*metainfo.Bucket, error) {
	if bucket == nil {
		return nil, errs.New("bucket parameter is nil")
	}

	id, err := bytesToUUID(bucket.Id)
	if err != nil {
		return nil, err
	}

	projectID, err := bytesToUUID(bucket.ProjectId)
	if err != nil {
		return nil, err
	}

	var attributionID uuid.UUID
	if bucket.AttributionId == nil {
		parsedID, err := bytesToUUID(bucket.AttributionId)
		if err != nil {
			return nil, err
		}
		attributionID = parsedID
	}

	return &metainfo.Bucket{
		ID: id,

		ProjectID:  projectID,
		Name:       string(bucket.Name),
		PathCipher: storj.Cipher(bucket.PathCipher),

		AttributionID: attributionID,
		CreatedAt:     bucket.CreatedAt,

		DefaultSegmentSize: int64(bucket.DefaultSegmentSize),
		DefaultEncryption: storj.EncryptionParameters{
			CipherSuite: storj.CipherSuite(bucket.DefaultEncryptionCipherSuite),
			BlockSize:   int32(bucket.DefaultEncryptionBlockSize),
		},
		DefaultRedundancy: storj.RedundancyScheme{
			Algorithm:      storj.RedundancyAlgorithm(bucket.DefaultRedundancyAlgorithm),
			ShareSize:      int32(bucket.DefaultRedundancyShareSize),
			RequiredShares: int16(bucket.DefaultRedundancyRequiredShares),
			RepairShares:   int16(bucket.DefaultRedundancyRepairShares),
			OptimalShares:  int16(bucket.DefaultRedundancyOptimalShares),
			TotalShares:    int16(bucket.DefaultRedundancyTotalShares),
		},
	}, nil
}