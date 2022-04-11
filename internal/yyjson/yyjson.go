// Copyright 2022 The yyjson-go Authors
// SPDX-License-Identifier: BSD-3-Clause

package yyjson

type READ_FLAG = Uint32_t /* yyjson.h:479:18 */

// Yyjson_read_flags re-defines using constants because ccgo/v3 unexported and use variable.
const (
	READ_NOFLAG                READ_FLAG = READ_FLAG(0)
	READ_INSITU                READ_FLAG = READ_FLAG(int32(1) << 0) /* yyjson.h:497:31 */
	READ_STOP_WHEN_DONE        READ_FLAG = READ_FLAG(int32(1) << 1) /* yyjson.h:502:31 */
	READ_ALLOW_TRAILING_COMMAS READ_FLAG = READ_FLAG(int32(1) << 2) /* yyjson.h:506:31 */
	READ_ALLOW_COMMENTS        READ_FLAG = READ_FLAG(int32(1) << 3) /* yyjson.h:509:31 */
	READ_ALLOW_INF_AND_NAN     READ_FLAG = READ_FLAG(int32(1) << 4) /* yyjson.h:513:31 */
)

type READ_CODE = Uint32_t /* yyjson.h:518:18 */

// Yyjson_read_code re-defines using constants because ccgo/v3 unexported and use variable.
const (
	READ_ERROR_INVALID_PARAMETER    READ_CODE = READ_CODE(1)  /* yyjson.h:524:31 */
	READ_ERROR_MEMORY_ALLOCATION    READ_CODE = READ_CODE(2)  /* yyjson.h:527:31 */
	READ_ERROR_EMPTY_CONTENT        READ_CODE = READ_CODE(3)  /* yyjson.h:530:31 */
	READ_ERROR_UNEXPECTED_CONTENT   READ_CODE = READ_CODE(4)  /* yyjson.h:533:31 */
	READ_ERROR_UNEXPECTED_END       READ_CODE = READ_CODE(5)  /* yyjson.h:536:31 */
	READ_ERROR_UNEXPECTED_CHARACTER READ_CODE = READ_CODE(6)  /* yyjson.h:539:31 */
	READ_ERROR_JSON_STRUCTURE       READ_CODE = READ_CODE(7)  /* yyjson.h:542:31 */
	READ_ERROR_INVALID_COMMENT      READ_CODE = READ_CODE(8)  /* yyjson.h:545:31 */
	READ_ERROR_INVALID_NUMBER       READ_CODE = READ_CODE(9)  /* yyjson.h:548:31 */
	READ_ERROR_INVALID_STRING       READ_CODE = READ_CODE(10) /* yyjson.h:551:31 */
	READ_ERROR_LITERAL              READ_CODE = READ_CODE(11) /* yyjson.h:554:31 */
	READ_ERROR_FILE_OPEN            READ_CODE = READ_CODE(12) /* yyjson.h:557:31 */
	READ_ERROR_FILE_READ            READ_CODE = READ_CODE(13) /* yyjson.h:560:31 */
)
