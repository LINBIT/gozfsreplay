// +build ignore

package gozfsreplay

// /*
//  * CDDL HEADER START
//  *
//  * The contents of this file are subject to the terms of the
//  * Common Development and Distribution License (the "License").
//  * You may not use this file except in compliance with the License.
//  *
//  * You can obtain a copy of the license at usr/src/OPENSOLARIS.LICENSE
//  * or http://www.opensolaris.org/os/licensing.
//  * See the License for the specific language governing permissions
//  * and limitations under the License.
//  *
//  * When distributing Covered Code, include this CDDL HEADER in each
//  * file and include the License file at usr/src/OPENSOLARIS.LICENSE.
//  * If applicable, add the following below this CDDL HEADER, with the
//  * fields enclosed by brackets "[]" replaced with your own identifying
//  * information: Portions Copyright [yyyy] [name of copyright owner]
//  *
//  * CDDL HEADER END
//  */
// #include <stdint.h>
//
// #define MAXNAMELEN 256
// #define ZIO_DATA_SALT_LEN 8
// #define ZIO_DATA_IV_LEN 12
// #define ZIO_DATA_MAC_LEN 16
//
// typedef enum dmu_objset_type {
// 	DMU_OST_NONE,
// 	DMU_OST_META,
// 	DMU_OST_ZFS,
// 	DMU_OST_ZVOL,
// 	DMU_OST_OTHER,			/* For testing only! */
// 	DMU_OST_ANY,			/* Be careful! */
// 	DMU_OST_NUMTYPES
// } dmu_objset_type_t;
//
// typedef struct zio_cksum {
// 	uint64_t zc_word[4];
// } zio_cksum_t;
//
// typedef enum dmu_object_type {
//		/* TODO(rck): lots are missing */
// 	DMU_OT_ZVOL = 23,                    /* UINT8 */
// 	DMU_OT_ZVOL_PROP = 24,               /* ZAP */
// } dmu_object_type_t;
//
// typedef struct ddt_key {
// 	zio_cksum_t     ddk_cksum;      /* 256-bit block checksum */
// 	uint64_t        ddk_prop;
// } ddt_key_t;
//
// typedef enum {
// 	DRR_BEGIN, DRR_OBJECT, DRR_FREEOBJECTS,
// 	DRR_WRITE, DRR_FREE, DRR_END, DRR_WRITE_BYREF,
// 	DRR_SPILL, DRR_WRITE_EMBEDDED, DRR_OBJECT_RANGE, DRR_REDACT,
// 	DRR_NUMTYPES
// } drr_type;
//
// struct drr_begin {
// 	uint64_t drr_magic;
// 	uint64_t drr_versioninfo; /* was drr_version */
// 	uint64_t drr_creation_time;
// 	dmu_objset_type_t drr_type;
// 	uint32_t drr_flags;
// 	uint64_t drr_toguid;
// 	uint64_t drr_fromguid;
// 	char drr_toname[MAXNAMELEN];
// } drr_begin;
// struct drr_end {
// 	zio_cksum_t drr_checksum;
// 	uint64_t drr_toguid;
// } drr_end;
// struct drr_object {
// 	uint64_t drr_object;
// 	dmu_object_type_t drr_type;
// 	dmu_object_type_t drr_bonustype;
// 	uint32_t drr_blksz;
// 	uint32_t drr_bonuslen;
// 	uint8_t drr_checksumtype;
// 	uint8_t drr_compress;
// 	uint8_t drr_dn_slots;
// 	uint8_t drr_flags;
// 	uint32_t drr_raw_bonuslen;
// 	uint64_t drr_toguid;
// 	/* only (possibly) nonzero for raw streams */
// 	uint8_t drr_indblkshift;
// 	uint8_t drr_nlevels;
// 	uint8_t drr_nblkptr;
// 	uint8_t drr_pad[5];
// 	uint64_t drr_maxblkid;
// 	/* bonus content follows */
// } drr_object;
// struct drr_freeobjects {
// 	uint64_t drr_firstobj;
// 	uint64_t drr_numobjs;
// 	uint64_t drr_toguid;
// } drr_freeobjects;
// struct drr_write {
// 	uint64_t drr_object;
// 	dmu_object_type_t drr_type;
// 	uint32_t drr_pad;
// 	uint64_t drr_offset;
// 	uint64_t drr_logical_size;
// 	uint64_t drr_toguid;
// 	uint8_t drr_checksumtype;
// 	uint8_t drr_flags;
// 	uint8_t drr_compressiontype;
// 	uint8_t drr_pad2[5];
// 	/* deduplication key */
// 	ddt_key_t drr_key;
// 	/* only nonzero if drr_compressiontype is not 0 */
// 	uint64_t drr_compressed_size;
// 	/* only nonzero for raw streams */
// 	uint8_t drr_salt[ZIO_DATA_SALT_LEN];
// 	uint8_t drr_iv[ZIO_DATA_IV_LEN];
// 	uint8_t drr_mac[ZIO_DATA_MAC_LEN];
// 	/* content follows */
// } drr_write;
// struct drr_free {
// 	uint64_t drr_object;
// 	uint64_t drr_offset;
// 	uint64_t drr_length;
// 	uint64_t drr_toguid;
// } drr_free;
// struct drr_write_byref {
// 	/* where to put the data */
// 	uint64_t drr_object;
// 	uint64_t drr_offset;
// 	uint64_t drr_length;
// 	uint64_t drr_toguid;
// 	/* where to find the prior copy of the data */
// 	uint64_t drr_refguid;
// 	uint64_t drr_refobject;
// 	uint64_t drr_refoffset;
// 	/* properties of the data */
// 	uint8_t drr_checksumtype;
// 	uint8_t drr_flags;
// 	uint8_t drr_pad2[6];
// 	ddt_key_t drr_key; /* deduplication key */
// } drr_write_byref;
// struct drr_spill {
// 	uint64_t drr_object;
// 	uint64_t drr_length;
// 	uint64_t drr_toguid;
// 	uint8_t drr_flags;
// 	uint8_t drr_compressiontype;
// 	uint8_t drr_pad[6];
// 	/* only nonzero for raw streams */
// 	uint64_t drr_compressed_size;
// 	uint8_t drr_salt[ZIO_DATA_SALT_LEN];
// 	uint8_t drr_iv[ZIO_DATA_IV_LEN];
// 	uint8_t drr_mac[ZIO_DATA_MAC_LEN];
// 	dmu_object_type_t drr_type;
// 	/* spill data follows */
// } drr_spill;
// struct drr_write_embedded {
// 	uint64_t drr_object;
// 	uint64_t drr_offset;
// 	/* logical length, should equal blocksize */
// 	uint64_t drr_length;
// 	uint64_t drr_toguid;
// 	uint8_t drr_compression;
// 	uint8_t drr_etype;
// 	uint8_t drr_pad[6];
// 	uint32_t drr_lsize; /* uncompressed size of payload */
// 	uint32_t drr_psize; /* compr. (real) size of payload */
// 	/* (possibly compressed) content follows */
// } drr_write_embedded;
// struct drr_object_range {
// 	uint64_t drr_firstobj;
// 	uint64_t drr_numslots;
// 	uint64_t drr_toguid;
// 	uint8_t drr_salt[ZIO_DATA_SALT_LEN];
// 	uint8_t drr_iv[ZIO_DATA_IV_LEN];
// 	uint8_t drr_mac[ZIO_DATA_MAC_LEN];
// 	uint8_t drr_flags;
// 	uint8_t drr_pad[3];
// } drr_object_range;
// struct drr_redact {
// 	uint64_t drr_object;
// 	uint64_t drr_offset;
// 	uint64_t drr_length;
// 	uint64_t drr_toguid;
// } drr_redact;
// struct drr_checksum {
// 	uint64_t drr_pad[34];
// 	/*
// 	 * fletcher-4 checksum of everything preceding the
// 	 * checksum.
// 	 */
// 	zio_cksum_t drr_checksum;
// } drr_checksum;
//
// typedef struct dmu_replay_record {
// 	drr_type drr_type;
// 	uint32_t drr_payloadlen;
// 	union {
// 		struct drr_begin drr_begin;
// 		struct drr_end drr_end;
// 		struct drr_object drr_object;
// 		struct drr_freeobjects drr_freeobjects;
// 		struct drr_write drr_write;
// 		struct drr_free drr_free;
// 		struct drr_write_byref drr_write_byref;
// 		struct drr_spill drr_spill;
// 		struct drr_write_embedded drr_write_embedded;
// 		struct drr_object_range drr_object_range;
// 		struct drr_redact drr_redact;
//
// 		/*
// 		 * Note: drr_checksum is overlaid with all record types
// 		 * except DRR_BEGIN.  Therefore its (non-pad) members
// 		 * must not overlap with members from the other structs.
// 		 * We accomplish this by putting its members at the very
// 		 * end of the struct.
// 		 */
// 		struct drr_checksum drr_checksum;
// 	} drr_u;
// } dmu_replay_record_t;
//
// /* added by rck from dmu_replay_record */
// typedef struct dmu_replay_header {
// 	drr_type drr_type;
// 	uint32_t drr_payloadlen;
// } dmu_replay_header_t;
import "C"
import (
	"encoding/binary"
	"errors"
	"os"
)

type ZIO_cksum C.struct_zio_cksum
type DDT_key C.struct_ddt_key
type DRR_begin C.struct_drr_begin
type DRR_end C.struct_drr_end
type DRR_object C.struct_drr_object
type DRR_freeobjects C.struct_drr_freeobjects
type DRR_write C.struct_drr_write
type DRR_free C.struct_drr_free
type DRR_write_byref C.struct_drr_write_byref
type DRR_spill C.struct_drr_spill
type DRR_write_embedded C.struct_drr_write_embedded
type DRR_object_range C.struct_drr_object_range
type DRR_redact C.struct_drr_redact
type DRR_checksum C.struct_drr_checksum
type DRR_replay_record C.struct_dmu_replay_record
type DRR_replay_header C.struct_dmu_replay_header

const Sizeof_ZIO_cksum = C.sizeof_struct_zio_cksum
const Sizeof_DRR_replay_record = C.sizeof_struct_dmu_replay_record
const Sizeof_DRR_replay_header = C.sizeof_struct_dmu_replay_header
const Sizeof_DRR_freeobjects = C.sizeof_struct_drr_freeobjects
const Sizeof_DRR_object = C.sizeof_struct_drr_object
const Sizeof_DRR_free = C.sizeof_struct_drr_free
const Sizeof_DRR_write = C.sizeof_struct_drr_write

func (w DRR_write) DRR_write_compressed() bool {
	return w.Compressiontype != 0
}

func (w DRR_write) DRR_write_payload_size() uint64 {
	if w.DRR_write_compressed() {
		return w.Compressed_size
	}
	return w.Logical_size
}

type DRR_type uint32

const (
	DRR_BEGIN DRR_type = iota
	DRR_OBJECT
	DRR_FREEOBJECTS
	DRR_WRITE
	DRR_FREE
	DRR_END
	DRR_WRITE_BYREF
	DRR_SPILL
	DRR_WRITE_EMBEDDED
	DRR_OBJECT_RANGE
	DRR_REDACT
	DRR_NUMTYPES
)

func DRRRead(r *os.File, order binary.ByteOrder) (interface{}, DRR_type, error) {
	if order != binary.LittleEndian {
		return nil, 0, errors.New("zfs: currently only binary.LittleEndian is supported")
	}

	hdr := DRR_replay_header{}
	err := binary.Read(r, order, &hdr)
	if err != nil {
		return nil, 0, err
	}

	var off int64
	typ := DRR_type(hdr.Type)
	switch typ {
	case DRR_BEGIN:
		obj := DRR_begin{}
		if err := binary.Read(r, order, &obj); err != nil {
			return obj, typ, err
		}
		return obj, typ, nil
	case DRR_FREEOBJECTS:
		off = Sizeof_DRR_freeobjects
		obj := DRR_freeobjects{}
		if err := binary.Read(r, order, &obj); err != nil {
			return obj, typ, err
		}
		return obj, typ, forward(r, hdr, off)
	case DRR_OBJECT:
		off = Sizeof_DRR_object
		obj := DRR_object{}
		if err := binary.Read(r, order, &obj); err != nil {
			return obj, typ, err
		}
		return obj, typ, forward(r, hdr, off)
	case DRR_FREE:
		off = Sizeof_DRR_free
		obj := DRR_free{}
		if err := binary.Read(r, order, &obj); err != nil {
			return obj, typ, err
		}
		return obj, typ, forward(r, hdr, off)
	case DRR_WRITE:
		off = Sizeof_DRR_write
		obj := DRR_write{}
		if err := binary.Read(r, order, &obj); err != nil {
			return obj, typ, err
		}
		if _, err := r.Seek(int64(obj.DRR_write_payload_size()), 1); err != nil {
			return obj, typ, err
		}
		return obj, typ, forward(r, hdr, off)
	case DRR_END:
		obj := DRR_end{}
		if err := binary.Read(r, order, &obj); err != nil {
			return obj, typ, err
		}
		return obj, typ, nil
	default:
		return nil, 0, errors.New("zfs: unknown type")
	}

	// this can not happen
	return nil, 0, errors.New("zfs: unknown type")
}

func forward(f *os.File, hdr DRR_replay_header, off int64) error {
	typ := DRR_type(hdr.Type)
	if typ != DRR_BEGIN {
		off = Sizeof_DRR_replay_record - Sizeof_DRR_replay_header - off
		if _, err := f.Seek(off, 1); err != nil {
			return err
		}
	}
	if _, err := f.Seek(int64(hdr.Payloadlen), 1); err != nil {
		return err
	}
	return nil
}
