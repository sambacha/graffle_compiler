// Code generated from /home/plagamedicum/Documents/yapis/graffle/GraffleParser.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // GraffleParser

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 73, 566,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9, 49, 4,
	50, 9, 50, 4, 51, 9, 51, 3, 2, 5, 2, 104, 10, 2, 3, 2, 7, 2, 107, 10, 2,
	12, 2, 14, 2, 110, 11, 2, 3, 2, 3, 2, 5, 2, 114, 10, 2, 3, 3, 5, 3, 117,
	10, 3, 3, 3, 3, 3, 6, 3, 121, 10, 3, 13, 3, 14, 3, 122, 3, 3, 5, 3, 126,
	10, 3, 6, 3, 128, 10, 3, 13, 3, 14, 3, 129, 3, 4, 3, 4, 3, 4, 3, 4, 5,
	4, 136, 10, 4, 3, 5, 3, 5, 3, 5, 7, 5, 141, 10, 5, 12, 5, 14, 5, 144, 11,
	5, 3, 5, 5, 5, 147, 10, 5, 3, 6, 3, 6, 5, 6, 151, 10, 6, 3, 7, 3, 7, 3,
	7, 5, 7, 156, 10, 7, 3, 8, 3, 8, 3, 8, 5, 8, 161, 10, 8, 3, 8, 3, 8, 3,
	8, 3, 8, 5, 8, 167, 10, 8, 5, 8, 169, 10, 8, 3, 8, 6, 8, 172, 10, 8, 13,
	8, 14, 8, 173, 3, 8, 5, 8, 177, 10, 8, 3, 8, 3, 8, 5, 8, 181, 10, 8, 3,
	8, 6, 8, 184, 10, 8, 13, 8, 14, 8, 185, 5, 8, 188, 10, 8, 3, 8, 3, 8, 3,
	9, 3, 9, 3, 9, 7, 9, 195, 10, 9, 12, 9, 14, 9, 198, 11, 9, 3, 9, 5, 9,
	201, 10, 9, 3, 9, 3, 9, 5, 9, 205, 10, 9, 3, 9, 6, 9, 208, 10, 9, 13, 9,
	14, 9, 209, 5, 9, 212, 10, 9, 3, 9, 3, 9, 3, 10, 5, 10, 217, 10, 10, 3,
	10, 3, 10, 3, 10, 5, 10, 222, 10, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10,
	228, 10, 10, 5, 10, 230, 10, 10, 3, 10, 6, 10, 233, 10, 10, 13, 10, 14,
	10, 234, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13,
	3, 13, 3, 13, 5, 13, 248, 10, 13, 3, 14, 3, 14, 3, 14, 5, 14, 253, 10,
	14, 3, 15, 3, 15, 3, 15, 5, 15, 258, 10, 15, 3, 16, 3, 16, 3, 16, 5, 16,
	263, 10, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 272,
	10, 16, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 278, 10, 16, 3, 16, 5, 16, 281,
	10, 16, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 287, 10, 16, 5, 16, 289, 10,
	16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 5, 17, 296, 10, 17, 3, 18, 5, 18,
	299, 10, 18, 3, 18, 3, 18, 5, 18, 303, 10, 18, 3, 18, 3, 18, 5, 18, 307,
	10, 18, 3, 18, 3, 18, 5, 18, 311, 10, 18, 3, 18, 5, 18, 314, 10, 18, 3,
	19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21,
	3, 21, 7, 21, 328, 10, 21, 12, 21, 14, 21, 331, 11, 21, 5, 21, 333, 10,
	21, 3, 21, 3, 21, 3, 21, 3, 21, 5, 21, 339, 10, 21, 3, 21, 5, 21, 342,
	10, 21, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 24,
	3, 24, 3, 24, 3, 24, 3, 24, 7, 24, 357, 10, 24, 12, 24, 14, 24, 360, 11,
	24, 5, 24, 362, 10, 24, 3, 24, 3, 24, 5, 24, 366, 10, 24, 3, 25, 3, 25,
	3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 5, 25, 378, 10,
	25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26,
	3, 26, 3, 26, 3, 26, 5, 26, 393, 10, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3,
	27, 3, 27, 5, 27, 401, 10, 27, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28,
	3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 30, 3,
	30, 3, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31,
	3, 31, 3, 31, 3, 31, 3, 31, 5, 31, 434, 10, 31, 3, 32, 3, 32, 3, 32, 3,
	32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 7, 32, 448,
	10, 32, 12, 32, 14, 32, 451, 11, 32, 3, 32, 3, 32, 3, 32, 3, 32, 5, 32,
	457, 10, 32, 3, 33, 3, 33, 5, 33, 461, 10, 33, 3, 33, 3, 33, 5, 33, 465,
	10, 33, 3, 33, 3, 33, 5, 33, 469, 10, 33, 5, 33, 471, 10, 33, 3, 34, 3,
	34, 5, 34, 475, 10, 34, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35,
	3, 35, 3, 35, 3, 35, 5, 35, 487, 10, 35, 3, 36, 3, 36, 3, 36, 3, 36, 3,
	36, 3, 36, 3, 36, 3, 36, 3, 36, 5, 36, 498, 10, 36, 3, 37, 3, 37, 5, 37,
	502, 10, 37, 3, 38, 3, 38, 3, 39, 3, 39, 3, 40, 3, 40, 3, 40, 3, 40, 3,
	41, 3, 41, 5, 41, 514, 10, 41, 3, 42, 3, 42, 3, 43, 3, 43, 3, 44, 3, 44,
	5, 44, 522, 10, 44, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 5, 45, 530,
	10, 45, 3, 46, 3, 46, 3, 46, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 7, 47,
	540, 10, 47, 12, 47, 14, 47, 543, 11, 47, 5, 47, 545, 10, 47, 3, 47, 3,
	47, 3, 48, 3, 48, 3, 49, 3, 49, 3, 49, 3, 49, 5, 49, 555, 10, 49, 3, 50,
	3, 50, 3, 51, 5, 51, 560, 10, 51, 3, 51, 3, 51, 5, 51, 564, 10, 51, 3,
	51, 2, 2, 52, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32,
	34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68,
	70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 2, 8,
	4, 2, 58, 58, 60, 60, 3, 2, 31, 41, 3, 2, 27, 30, 3, 2, 23, 26, 3, 2, 7,
	8, 5, 2, 3, 3, 6, 6, 9, 9, 2, 617, 2, 103, 3, 2, 2, 2, 4, 116, 3, 2, 2,
	2, 6, 135, 3, 2, 2, 2, 8, 137, 3, 2, 2, 2, 10, 150, 3, 2, 2, 2, 12, 155,
	3, 2, 2, 2, 14, 157, 3, 2, 2, 2, 16, 191, 3, 2, 2, 2, 18, 216, 3, 2, 2,
	2, 20, 236, 3, 2, 2, 2, 22, 239, 3, 2, 2, 2, 24, 247, 3, 2, 2, 2, 26, 249,
	3, 2, 2, 2, 28, 254, 3, 2, 2, 2, 30, 288, 3, 2, 2, 2, 32, 290, 3, 2, 2,
	2, 34, 313, 3, 2, 2, 2, 36, 315, 3, 2, 2, 2, 38, 318, 3, 2, 2, 2, 40, 322,
	3, 2, 2, 2, 42, 343, 3, 2, 2, 2, 44, 346, 3, 2, 2, 2, 46, 351, 3, 2, 2,
	2, 48, 377, 3, 2, 2, 2, 50, 392, 3, 2, 2, 2, 52, 400, 3, 2, 2, 2, 54, 402,
	3, 2, 2, 2, 56, 408, 3, 2, 2, 2, 58, 414, 3, 2, 2, 2, 60, 433, 3, 2, 2,
	2, 62, 456, 3, 2, 2, 2, 64, 470, 3, 2, 2, 2, 66, 474, 3, 2, 2, 2, 68, 486,
	3, 2, 2, 2, 70, 497, 3, 2, 2, 2, 72, 501, 3, 2, 2, 2, 74, 503, 3, 2, 2,
	2, 76, 505, 3, 2, 2, 2, 78, 507, 3, 2, 2, 2, 80, 513, 3, 2, 2, 2, 82, 515,
	3, 2, 2, 2, 84, 517, 3, 2, 2, 2, 86, 521, 3, 2, 2, 2, 88, 529, 3, 2, 2,
	2, 90, 531, 3, 2, 2, 2, 92, 534, 3, 2, 2, 2, 94, 548, 3, 2, 2, 2, 96, 554,
	3, 2, 2, 2, 98, 556, 3, 2, 2, 2, 100, 559, 3, 2, 2, 2, 102, 104, 7, 60,
	2, 2, 103, 102, 3, 2, 2, 2, 103, 104, 3, 2, 2, 2, 104, 108, 3, 2, 2, 2,
	105, 107, 5, 34, 18, 2, 106, 105, 3, 2, 2, 2, 107, 110, 3, 2, 2, 2, 108,
	106, 3, 2, 2, 2, 108, 109, 3, 2, 2, 2, 109, 111, 3, 2, 2, 2, 110, 108,
	3, 2, 2, 2, 111, 113, 5, 4, 3, 2, 112, 114, 7, 2, 2, 3, 113, 112, 3, 2,
	2, 2, 113, 114, 3, 2, 2, 2, 114, 3, 3, 2, 2, 2, 115, 117, 7, 60, 2, 2,
	116, 115, 3, 2, 2, 2, 116, 117, 3, 2, 2, 2, 117, 127, 3, 2, 2, 2, 118,
	125, 5, 6, 4, 2, 119, 121, 9, 2, 2, 2, 120, 119, 3, 2, 2, 2, 121, 122,
	3, 2, 2, 2, 122, 120, 3, 2, 2, 2, 122, 123, 3, 2, 2, 2, 123, 126, 3, 2,
	2, 2, 124, 126, 7, 2, 2, 3, 125, 120, 3, 2, 2, 2, 125, 124, 3, 2, 2, 2,
	126, 128, 3, 2, 2, 2, 127, 118, 3, 2, 2, 2, 128, 129, 3, 2, 2, 2, 129,
	127, 3, 2, 2, 2, 129, 130, 3, 2, 2, 2, 130, 5, 3, 2, 2, 2, 131, 136, 5,
	14, 8, 2, 132, 136, 5, 16, 9, 2, 133, 136, 5, 22, 12, 2, 134, 136, 5, 8,
	5, 2, 135, 131, 3, 2, 2, 2, 135, 132, 3, 2, 2, 2, 135, 133, 3, 2, 2, 2,
	135, 134, 3, 2, 2, 2, 136, 7, 3, 2, 2, 2, 137, 142, 5, 10, 6, 2, 138, 139,
	7, 58, 2, 2, 139, 141, 5, 10, 6, 2, 140, 138, 3, 2, 2, 2, 141, 144, 3,
	2, 2, 2, 142, 140, 3, 2, 2, 2, 142, 143, 3, 2, 2, 2, 143, 146, 3, 2, 2,
	2, 144, 142, 3, 2, 2, 2, 145, 147, 7, 58, 2, 2, 146, 145, 3, 2, 2, 2, 146,
	147, 3, 2, 2, 2, 147, 9, 3, 2, 2, 2, 148, 151, 5, 12, 7, 2, 149, 151, 5,
	20, 11, 2, 150, 148, 3, 2, 2, 2, 150, 149, 3, 2, 2, 2, 151, 11, 3, 2, 2,
	2, 152, 156, 5, 48, 25, 2, 153, 156, 5, 92, 47, 2, 154, 156, 5, 86, 44,
	2, 155, 152, 3, 2, 2, 2, 155, 153, 3, 2, 2, 2, 155, 154, 3, 2, 2, 2, 156,
	13, 3, 2, 2, 2, 157, 158, 7, 43, 2, 2, 158, 168, 5, 70, 36, 2, 159, 161,
	7, 59, 2, 2, 160, 159, 3, 2, 2, 2, 160, 161, 3, 2, 2, 2, 161, 166, 3, 2,
	2, 2, 162, 167, 7, 44, 2, 2, 163, 167, 7, 55, 2, 2, 164, 165, 7, 44, 2,
	2, 165, 167, 7, 55, 2, 2, 166, 162, 3, 2, 2, 2, 166, 163, 3, 2, 2, 2, 166,
	164, 3, 2, 2, 2, 167, 169, 3, 2, 2, 2, 168, 160, 3, 2, 2, 2, 168, 169,
	3, 2, 2, 2, 169, 171, 3, 2, 2, 2, 170, 172, 5, 4, 3, 2, 171, 170, 3, 2,
	2, 2, 172, 173, 3, 2, 2, 2, 173, 171, 3, 2, 2, 2, 173, 174, 3, 2, 2, 2,
	174, 176, 3, 2, 2, 2, 175, 177, 7, 60, 2, 2, 176, 175, 3, 2, 2, 2, 176,
	177, 3, 2, 2, 2, 177, 187, 3, 2, 2, 2, 178, 180, 7, 45, 2, 2, 179, 181,
	7, 55, 2, 2, 180, 179, 3, 2, 2, 2, 180, 181, 3, 2, 2, 2, 181, 183, 3, 2,
	2, 2, 182, 184, 5, 4, 3, 2, 183, 182, 3, 2, 2, 2, 184, 185, 3, 2, 2, 2,
	185, 183, 3, 2, 2, 2, 185, 186, 3, 2, 2, 2, 186, 188, 3, 2, 2, 2, 187,
	178, 3, 2, 2, 2, 187, 188, 3, 2, 2, 2, 188, 189, 3, 2, 2, 2, 189, 190,
	5, 100, 51, 2, 190, 15, 3, 2, 2, 2, 191, 192, 7, 43, 2, 2, 192, 196, 5,
	96, 49, 2, 193, 195, 5, 18, 10, 2, 194, 193, 3, 2, 2, 2, 195, 198, 3, 2,
	2, 2, 196, 194, 3, 2, 2, 2, 196, 197, 3, 2, 2, 2, 197, 200, 3, 2, 2, 2,
	198, 196, 3, 2, 2, 2, 199, 201, 7, 60, 2, 2, 200, 199, 3, 2, 2, 2, 200,
	201, 3, 2, 2, 2, 201, 211, 3, 2, 2, 2, 202, 204, 7, 47, 2, 2, 203, 205,
	7, 55, 2, 2, 204, 203, 3, 2, 2, 2, 204, 205, 3, 2, 2, 2, 205, 207, 3, 2,
	2, 2, 206, 208, 5, 4, 3, 2, 207, 206, 3, 2, 2, 2, 208, 209, 3, 2, 2, 2,
	209, 207, 3, 2, 2, 2, 209, 210, 3, 2, 2, 2, 210, 212, 3, 2, 2, 2, 211,
	202, 3, 2, 2, 2, 211, 212, 3, 2, 2, 2, 212, 213, 3, 2, 2, 2, 213, 214,
	5, 100, 51, 2, 214, 17, 3, 2, 2, 2, 215, 217, 7, 60, 2, 2, 216, 215, 3,
	2, 2, 2, 216, 217, 3, 2, 2, 2, 217, 218, 3, 2, 2, 2, 218, 219, 7, 46, 2,
	2, 219, 229, 5, 96, 49, 2, 220, 222, 7, 59, 2, 2, 221, 220, 3, 2, 2, 2,
	221, 222, 3, 2, 2, 2, 222, 227, 3, 2, 2, 2, 223, 228, 7, 44, 2, 2, 224,
	228, 7, 55, 2, 2, 225, 226, 7, 44, 2, 2, 226, 228, 7, 55, 2, 2, 227, 223,
	3, 2, 2, 2, 227, 224, 3, 2, 2, 2, 227, 225, 3, 2, 2, 2, 228, 230, 3, 2,
	2, 2, 229, 221, 3, 2, 2, 2, 229, 230, 3, 2, 2, 2, 230, 232, 3, 2, 2, 2,
	231, 233, 5, 4, 3, 2, 232, 231, 3, 2, 2, 2, 233, 234, 3, 2, 2, 2, 234,
	232, 3, 2, 2, 2, 234, 235, 3, 2, 2, 2, 235, 19, 3, 2, 2, 2, 236, 237, 5,
	24, 13, 2, 237, 238, 5, 8, 5, 2, 238, 21, 3, 2, 2, 2, 239, 240, 5, 24,
	13, 2, 240, 241, 5, 4, 3, 2, 241, 242, 5, 100, 51, 2, 242, 23, 3, 2, 2,
	2, 243, 248, 5, 26, 14, 2, 244, 248, 5, 28, 15, 2, 245, 248, 5, 30, 16,
	2, 246, 248, 5, 32, 17, 2, 247, 243, 3, 2, 2, 2, 247, 244, 3, 2, 2, 2,
	247, 245, 3, 2, 2, 2, 247, 246, 3, 2, 2, 2, 248, 25, 3, 2, 2, 2, 249, 250,
	7, 48, 2, 2, 250, 252, 5, 70, 36, 2, 251, 253, 7, 55, 2, 2, 252, 251, 3,
	2, 2, 2, 252, 253, 3, 2, 2, 2, 253, 27, 3, 2, 2, 2, 254, 255, 7, 49, 2,
	2, 255, 257, 5, 70, 36, 2, 256, 258, 7, 55, 2, 2, 257, 256, 3, 2, 2, 2,
	257, 258, 3, 2, 2, 2, 258, 29, 3, 2, 2, 2, 259, 260, 7, 50, 2, 2, 260,
	262, 5, 70, 36, 2, 261, 263, 7, 55, 2, 2, 262, 261, 3, 2, 2, 2, 262, 263,
	3, 2, 2, 2, 263, 289, 3, 2, 2, 2, 264, 265, 7, 50, 2, 2, 265, 266, 5, 12,
	7, 2, 266, 267, 7, 59, 2, 2, 267, 268, 5, 70, 36, 2, 268, 269, 7, 59, 2,
	2, 269, 271, 5, 12, 7, 2, 270, 272, 7, 55, 2, 2, 271, 270, 3, 2, 2, 2,
	271, 272, 3, 2, 2, 2, 272, 289, 3, 2, 2, 2, 273, 274, 7, 50, 2, 2, 274,
	275, 7, 69, 2, 2, 275, 277, 7, 51, 2, 2, 276, 278, 7, 52, 2, 2, 277, 276,
	3, 2, 2, 2, 277, 278, 3, 2, 2, 2, 278, 280, 3, 2, 2, 2, 279, 281, 7, 53,
	2, 2, 280, 279, 3, 2, 2, 2, 280, 281, 3, 2, 2, 2, 281, 282, 3, 2, 2, 2,
	282, 283, 5, 66, 34, 2, 283, 284, 7, 54, 2, 2, 284, 286, 5, 66, 34, 2,
	285, 287, 7, 55, 2, 2, 286, 285, 3, 2, 2, 2, 286, 287, 3, 2, 2, 2, 287,
	289, 3, 2, 2, 2, 288, 259, 3, 2, 2, 2, 288, 264, 3, 2, 2, 2, 288, 273,
	3, 2, 2, 2, 289, 31, 3, 2, 2, 2, 290, 291, 7, 53, 2, 2, 291, 292, 5, 66,
	34, 2, 292, 293, 7, 54, 2, 2, 293, 295, 5, 66, 34, 2, 294, 296, 7, 55,
	2, 2, 295, 294, 3, 2, 2, 2, 295, 296, 3, 2, 2, 2, 296, 33, 3, 2, 2, 2,
	297, 299, 7, 60, 2, 2, 298, 297, 3, 2, 2, 2, 298, 299, 3, 2, 2, 2, 299,
	300, 3, 2, 2, 2, 300, 314, 5, 36, 19, 2, 301, 303, 7, 60, 2, 2, 302, 301,
	3, 2, 2, 2, 302, 303, 3, 2, 2, 2, 303, 304, 3, 2, 2, 2, 304, 314, 5, 38,
	20, 2, 305, 307, 7, 60, 2, 2, 306, 305, 3, 2, 2, 2, 306, 307, 3, 2, 2,
	2, 307, 308, 3, 2, 2, 2, 308, 314, 5, 42, 22, 2, 309, 311, 7, 60, 2, 2,
	310, 309, 3, 2, 2, 2, 310, 311, 3, 2, 2, 2, 311, 312, 3, 2, 2, 2, 312,
	314, 5, 44, 23, 2, 313, 298, 3, 2, 2, 2, 313, 302, 3, 2, 2, 2, 313, 306,
	3, 2, 2, 2, 313, 310, 3, 2, 2, 2, 314, 35, 3, 2, 2, 2, 315, 316, 5, 40,
	21, 2, 316, 317, 5, 8, 5, 2, 317, 37, 3, 2, 2, 2, 318, 319, 5, 40, 21,
	2, 319, 320, 5, 4, 3, 2, 320, 321, 5, 100, 51, 2, 321, 39, 3, 2, 2, 2,
	322, 323, 7, 70, 2, 2, 323, 332, 7, 10, 2, 2, 324, 329, 7, 69, 2, 2, 325,
	326, 7, 59, 2, 2, 326, 328, 7, 69, 2, 2, 327, 325, 3, 2, 2, 2, 328, 331,
	3, 2, 2, 2, 329, 327, 3, 2, 2, 2, 329, 330, 3, 2, 2, 2, 330, 333, 3, 2,
	2, 2, 331, 329, 3, 2, 2, 2, 332, 324, 3, 2, 2, 2, 332, 333, 3, 2, 2, 2,
	333, 334, 3, 2, 2, 2, 334, 335, 7, 11, 2, 2, 335, 336, 7, 22, 2, 2, 336,
	341, 5, 96, 49, 2, 337, 339, 7, 59, 2, 2, 338, 337, 3, 2, 2, 2, 338, 339,
	3, 2, 2, 2, 339, 340, 3, 2, 2, 2, 340, 342, 7, 68, 2, 2, 341, 338, 3, 2,
	2, 2, 341, 342, 3, 2, 2, 2, 342, 41, 3, 2, 2, 2, 343, 344, 5, 46, 24, 2,
	344, 345, 5, 8, 5, 2, 345, 43, 3, 2, 2, 2, 346, 347, 5, 46, 24, 2, 347,
	348, 7, 60, 2, 2, 348, 349, 5, 4, 3, 2, 349, 350, 5, 100, 51, 2, 350, 45,
	3, 2, 2, 2, 351, 352, 7, 70, 2, 2, 352, 361, 7, 10, 2, 2, 353, 358, 7,
	69, 2, 2, 354, 355, 7, 59, 2, 2, 355, 357, 7, 69, 2, 2, 356, 354, 3, 2,
	2, 2, 357, 360, 3, 2, 2, 2, 358, 356, 3, 2, 2, 2, 358, 359, 3, 2, 2, 2,
	359, 362, 3, 2, 2, 2, 360, 358, 3, 2, 2, 2, 361, 353, 3, 2, 2, 2, 361,
	362, 3, 2, 2, 2, 362, 363, 3, 2, 2, 2, 363, 365, 7, 11, 2, 2, 364, 366,
	7, 22, 2, 2, 365, 364, 3, 2, 2, 2, 365, 366, 3, 2, 2, 2, 366, 47, 3, 2,
	2, 2, 367, 368, 7, 70, 2, 2, 368, 369, 7, 22, 2, 2, 369, 378, 7, 69, 2,
	2, 370, 371, 7, 70, 2, 2, 371, 372, 7, 22, 2, 2, 372, 378, 5, 66, 34, 2,
	373, 378, 5, 50, 26, 2, 374, 378, 5, 60, 31, 2, 375, 378, 5, 62, 32, 2,
	376, 378, 5, 64, 33, 2, 377, 367, 3, 2, 2, 2, 377, 370, 3, 2, 2, 2, 377,
	373, 3, 2, 2, 2, 377, 374, 3, 2, 2, 2, 377, 375, 3, 2, 2, 2, 377, 376,
	3, 2, 2, 2, 378, 49, 3, 2, 2, 2, 379, 380, 7, 70, 2, 2, 380, 381, 7, 22,
	2, 2, 381, 382, 7, 67, 2, 2, 382, 383, 7, 10, 2, 2, 383, 384, 5, 96, 49,
	2, 384, 385, 7, 11, 2, 2, 385, 393, 3, 2, 2, 2, 386, 387, 7, 70, 2, 2,
	387, 388, 7, 22, 2, 2, 388, 389, 5, 96, 49, 2, 389, 390, 5, 52, 27, 2,
	390, 391, 5, 96, 49, 2, 391, 393, 3, 2, 2, 2, 392, 379, 3, 2, 2, 2, 392,
	386, 3, 2, 2, 2, 393, 51, 3, 2, 2, 2, 394, 401, 7, 16, 2, 2, 395, 401,
	5, 54, 28, 2, 396, 401, 7, 17, 2, 2, 397, 401, 5, 56, 29, 2, 398, 401,
	7, 20, 2, 2, 399, 401, 5, 58, 30, 2, 400, 394, 3, 2, 2, 2, 400, 395, 3,
	2, 2, 2, 400, 396, 3, 2, 2, 2, 400, 397, 3, 2, 2, 2, 400, 398, 3, 2, 2,
	2, 400, 399, 3, 2, 2, 2, 401, 53, 3, 2, 2, 2, 402, 403, 7, 28, 2, 2, 403,
	404, 7, 14, 2, 2, 404, 405, 7, 3, 2, 2, 405, 406, 7, 15, 2, 2, 406, 407,
	7, 16, 2, 2, 407, 55, 3, 2, 2, 2, 408, 409, 7, 17, 2, 2, 409, 410, 7, 14,
	2, 2, 410, 411, 7, 3, 2, 2, 411, 412, 7, 15, 2, 2, 412, 413, 7, 28, 2,
	2, 413, 57, 3, 2, 2, 2, 414, 415, 7, 28, 2, 2, 415, 416, 7, 14, 2, 2, 416,
	417, 7, 3, 2, 2, 417, 418, 7, 15, 2, 2, 418, 419, 7, 28, 2, 2, 419, 59,
	3, 2, 2, 2, 420, 421, 7, 70, 2, 2, 421, 422, 7, 22, 2, 2, 422, 423, 7,
	10, 2, 2, 423, 424, 7, 66, 2, 2, 424, 425, 7, 11, 2, 2, 425, 434, 5, 96,
	49, 2, 426, 427, 7, 70, 2, 2, 427, 428, 7, 22, 2, 2, 428, 434, 5, 96, 49,
	2, 429, 430, 7, 70, 2, 2, 430, 431, 5, 84, 43, 2, 431, 432, 5, 96, 49,
	2, 432, 434, 3, 2, 2, 2, 433, 420, 3, 2, 2, 2, 433, 426, 3, 2, 2, 2, 433,
	429, 3, 2, 2, 2, 434, 61, 3, 2, 2, 2, 435, 436, 7, 70, 2, 2, 436, 437,
	7, 22, 2, 2, 437, 438, 7, 65, 2, 2, 438, 439, 7, 10, 2, 2, 439, 440, 5,
	96, 49, 2, 440, 441, 7, 11, 2, 2, 441, 457, 3, 2, 2, 2, 442, 443, 7, 70,
	2, 2, 443, 444, 7, 22, 2, 2, 444, 449, 5, 96, 49, 2, 445, 446, 7, 59, 2,
	2, 446, 448, 5, 96, 49, 2, 447, 445, 3, 2, 2, 2, 448, 451, 3, 2, 2, 2,
	449, 447, 3, 2, 2, 2, 449, 450, 3, 2, 2, 2, 450, 457, 3, 2, 2, 2, 451,
	449, 3, 2, 2, 2, 452, 453, 7, 70, 2, 2, 453, 454, 5, 84, 43, 2, 454, 455,
	5, 96, 49, 2, 455, 457, 3, 2, 2, 2, 456, 435, 3, 2, 2, 2, 456, 442, 3,
	2, 2, 2, 456, 452, 3, 2, 2, 2, 457, 63, 3, 2, 2, 2, 458, 460, 5, 60, 31,
	2, 459, 461, 5, 94, 48, 2, 460, 459, 3, 2, 2, 2, 460, 461, 3, 2, 2, 2,
	461, 471, 3, 2, 2, 2, 462, 464, 5, 50, 26, 2, 463, 465, 5, 94, 48, 2, 464,
	463, 3, 2, 2, 2, 464, 465, 3, 2, 2, 2, 465, 471, 3, 2, 2, 2, 466, 468,
	5, 62, 32, 2, 467, 469, 5, 94, 48, 2, 468, 467, 3, 2, 2, 2, 468, 469, 3,
	2, 2, 2, 469, 471, 3, 2, 2, 2, 470, 458, 3, 2, 2, 2, 470, 462, 3, 2, 2,
	2, 470, 466, 3, 2, 2, 2, 471, 65, 3, 2, 2, 2, 472, 475, 7, 69, 2, 2, 473,
	475, 5, 68, 35, 2, 474, 472, 3, 2, 2, 2, 474, 473, 3, 2, 2, 2, 475, 67,
	3, 2, 2, 2, 476, 487, 5, 92, 47, 2, 477, 487, 5, 98, 50, 2, 478, 479, 7,
	10, 2, 2, 479, 480, 5, 70, 36, 2, 480, 481, 7, 11, 2, 2, 481, 487, 3, 2,
	2, 2, 482, 483, 7, 10, 2, 2, 483, 484, 5, 78, 40, 2, 484, 485, 7, 11, 2,
	2, 485, 487, 3, 2, 2, 2, 486, 476, 3, 2, 2, 2, 486, 477, 3, 2, 2, 2, 486,
	478, 3, 2, 2, 2, 486, 482, 3, 2, 2, 2, 487, 69, 3, 2, 2, 2, 488, 489, 5,
	66, 34, 2, 489, 490, 5, 74, 38, 2, 490, 491, 5, 72, 37, 2, 491, 498, 3,
	2, 2, 2, 492, 493, 5, 76, 39, 2, 493, 494, 5, 72, 37, 2, 494, 498, 3, 2,
	2, 2, 495, 498, 5, 78, 40, 2, 496, 498, 5, 66, 34, 2, 497, 488, 3, 2, 2,
	2, 497, 492, 3, 2, 2, 2, 497, 495, 3, 2, 2, 2, 497, 496, 3, 2, 2, 2, 498,
	71, 3, 2, 2, 2, 499, 502, 5, 66, 34, 2, 500, 502, 5, 70, 36, 2, 501, 499,
	3, 2, 2, 2, 501, 500, 3, 2, 2, 2, 502, 73, 3, 2, 2, 2, 503, 504, 9, 3,
	2, 2, 504, 75, 3, 2, 2, 2, 505, 506, 7, 42, 2, 2, 506, 77, 3, 2, 2, 2,
	507, 508, 5, 66, 34, 2, 508, 509, 5, 82, 42, 2, 509, 510, 5, 80, 41, 2,
	510, 79, 3, 2, 2, 2, 511, 514, 5, 66, 34, 2, 512, 514, 5, 78, 40, 2, 513,
	511, 3, 2, 2, 2, 513, 512, 3, 2, 2, 2, 514, 81, 3, 2, 2, 2, 515, 516, 9,
	4, 2, 2, 516, 83, 3, 2, 2, 2, 517, 518, 9, 5, 2, 2, 518, 85, 3, 2, 2, 2,
	519, 522, 5, 88, 45, 2, 520, 522, 5, 90, 46, 2, 521, 519, 3, 2, 2, 2, 521,
	520, 3, 2, 2, 2, 522, 87, 3, 2, 2, 2, 523, 524, 7, 63, 2, 2, 524, 530,
	7, 69, 2, 2, 525, 526, 7, 63, 2, 2, 526, 530, 5, 96, 49, 2, 527, 528, 7,
	63, 2, 2, 528, 530, 5, 92, 47, 2, 529, 523, 3, 2, 2, 2, 529, 525, 3, 2,
	2, 2, 529, 527, 3, 2, 2, 2, 530, 89, 3, 2, 2, 2, 531, 532, 7, 64, 2, 2,
	532, 533, 7, 70, 2, 2, 533, 91, 3, 2, 2, 2, 534, 535, 7, 70, 2, 2, 535,
	544, 7, 10, 2, 2, 536, 541, 5, 96, 49, 2, 537, 538, 7, 59, 2, 2, 538, 540,
	5, 96, 49, 2, 539, 537, 3, 2, 2, 2, 540, 543, 3, 2, 2, 2, 541, 539, 3,
	2, 2, 2, 541, 542, 3, 2, 2, 2, 542, 545, 3, 2, 2, 2, 543, 541, 3, 2, 2,
	2, 544, 536, 3, 2, 2, 2, 544, 545, 3, 2, 2, 2, 545, 546, 3, 2, 2, 2, 546,
	547, 7, 11, 2, 2, 547, 93, 3, 2, 2, 2, 548, 549, 9, 6, 2, 2, 549, 95, 3,
	2, 2, 2, 550, 555, 5, 70, 36, 2, 551, 555, 5, 78, 40, 2, 552, 555, 5, 66,
	34, 2, 553, 555, 5, 98, 50, 2, 554, 550, 3, 2, 2, 2, 554, 551, 3, 2, 2,
	2, 554, 552, 3, 2, 2, 2, 554, 553, 3, 2, 2, 2, 555, 97, 3, 2, 2, 2, 556,
	557, 9, 7, 2, 2, 557, 99, 3, 2, 2, 2, 558, 560, 7, 60, 2, 2, 559, 558,
	3, 2, 2, 2, 559, 560, 3, 2, 2, 2, 560, 561, 3, 2, 2, 2, 561, 563, 7, 62,
	2, 2, 562, 564, 7, 60, 2, 2, 563, 562, 3, 2, 2, 2, 563, 564, 3, 2, 2, 2,
	564, 101, 3, 2, 2, 2, 76, 103, 108, 113, 116, 122, 125, 129, 135, 142,
	146, 150, 155, 160, 166, 168, 173, 176, 180, 185, 187, 196, 200, 204, 209,
	211, 216, 221, 227, 229, 234, 247, 252, 257, 262, 271, 277, 280, 286, 288,
	295, 298, 302, 306, 310, 313, 329, 332, 338, 341, 358, 361, 365, 377, 392,
	400, 433, 449, 456, 460, 464, 468, 470, 474, 486, 497, 501, 513, 521, 529,
	541, 544, 554, 559, 563,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "", "", "", "", "", "", "'('", "')'", "'{'", "'}'", "'['", "']'",
	"'->'", "'<-'", "", "", "'--'", "", "'='", "'+='", "'-='", "'*='", "'/='",
	"'+'", "'-'", "'*'", "'/'", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "'.'",
	"','",
}
var symbolicNames = []string{
	"", "NUMBER", "FLOAT", "INT", "STRING", "LABEL", "ML_LABEL", "BOOL", "L_PAREN",
	"R_PAREN", "L_CURLY", "R_CURLY", "L_BRACKET", "R_BRACKET", "OR_ARC_LR",
	"OR_ARC_RL", "OR_W_ARC_LR", "OR_W_ARC_RL", "UNOR_ARC", "UNOR_W_ARC", "ASSIGN",
	"ADD_ASSIGN", "SUB_ASSIGN", "MULT_ASSIGN", "DIV_ASSIGN", "ADD", "SUB",
	"MULT", "DIV", "NEQ", "EQUALS", "LESS_THAN", "GR_THAN", "LESS_THAN_E",
	"GR_THAN_E", "AND", "OR", "XOR", "NOR", "NAND", "NOT", "IF", "THEN", "ELSE",
	"IS", "DEFAULT", "WHILE", "UNTIL", "FOR", "IN", "RANGE", "FROM", "TO",
	"DO", "SKIP_ITERATION", "BREAK", "ACT_DELIM", "ARG_DELIM", "NEWLINE", "BLOCK_BEGIN",
	"BLOCK_END", "PRINTER", "KEY_INPUT", "G_N", "V_N", "E_N", "WHERE", "VAR",
	"ID", "WS", "LINE_COMMENT", "M_LINE_COMMENT",
}

var ruleNames = []string{
	"file", "sequence", "sequence_element", "sequence_line", "one_line_sequence_element",
	"atom_action", "if_stmnt", "if_is_stmnt", "case_stmnt", "one_line_stmnt",
	"mult_line_stmnt", "stmnt", "while_stmnt", "until_stmnt", "for_stmnt",
	"from_to_stmnt", "function_declaration", "one_line_function_declaration",
	"mult_line_function_declaration", "function_declaration_head", "one_line_procedure_declaration",
	"mult_line_procedure_declaration", "procedure_declaration_head", "var_declaration",
	"arc_declaration", "arc", "or_w_arc_lr", "or_w_arc_rl", "unor_w_arc", "vertice_declaration",
	"graph_declaration", "labeled_declaration", "expr", "integral_expr", "logical_expr",
	"log_expr_operand", "bin_log_operator", "unar_log_operator", "arithm_expr",
	"arithm_expr_operand", "bin_arithm_operator", "arithm_assign_operator",
	"builtin_function_call", "built_func_print", "built_func_input", "function_call",
	"label", "value", "builtin", "block_end",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type GraffleParser struct {
	*antlr.BaseParser
}

func NewGraffleParser(input antlr.TokenStream) *GraffleParser {
	this := new(GraffleParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "GraffleParser.g4"

	return this
}

// GraffleParser tokens.
const (
	GraffleParserEOF            = antlr.TokenEOF
	GraffleParserNUMBER         = 1
	GraffleParserFLOAT          = 2
	GraffleParserINT            = 3
	GraffleParserSTRING         = 4
	GraffleParserLABEL          = 5
	GraffleParserML_LABEL       = 6
	GraffleParserBOOL           = 7
	GraffleParserL_PAREN        = 8
	GraffleParserR_PAREN        = 9
	GraffleParserL_CURLY        = 10
	GraffleParserR_CURLY        = 11
	GraffleParserL_BRACKET      = 12
	GraffleParserR_BRACKET      = 13
	GraffleParserOR_ARC_LR      = 14
	GraffleParserOR_ARC_RL      = 15
	GraffleParserOR_W_ARC_LR    = 16
	GraffleParserOR_W_ARC_RL    = 17
	GraffleParserUNOR_ARC       = 18
	GraffleParserUNOR_W_ARC     = 19
	GraffleParserASSIGN         = 20
	GraffleParserADD_ASSIGN     = 21
	GraffleParserSUB_ASSIGN     = 22
	GraffleParserMULT_ASSIGN    = 23
	GraffleParserDIV_ASSIGN     = 24
	GraffleParserADD            = 25
	GraffleParserSUB            = 26
	GraffleParserMULT           = 27
	GraffleParserDIV            = 28
	GraffleParserNEQ            = 29
	GraffleParserEQUALS         = 30
	GraffleParserLESS_THAN      = 31
	GraffleParserGR_THAN        = 32
	GraffleParserLESS_THAN_E    = 33
	GraffleParserGR_THAN_E      = 34
	GraffleParserAND            = 35
	GraffleParserOR             = 36
	GraffleParserXOR            = 37
	GraffleParserNOR            = 38
	GraffleParserNAND           = 39
	GraffleParserNOT            = 40
	GraffleParserIF             = 41
	GraffleParserTHEN           = 42
	GraffleParserELSE           = 43
	GraffleParserIS             = 44
	GraffleParserDEFAULT        = 45
	GraffleParserWHILE          = 46
	GraffleParserUNTIL          = 47
	GraffleParserFOR            = 48
	GraffleParserIN             = 49
	GraffleParserRANGE          = 50
	GraffleParserFROM           = 51
	GraffleParserTO             = 52
	GraffleParserDO             = 53
	GraffleParserSKIP_ITERATION = 54
	GraffleParserBREAK          = 55
	GraffleParserACT_DELIM      = 56
	GraffleParserARG_DELIM      = 57
	GraffleParserNEWLINE        = 58
	GraffleParserBLOCK_BEGIN    = 59
	GraffleParserBLOCK_END      = 60
	GraffleParserPRINTER        = 61
	GraffleParserKEY_INPUT      = 62
	GraffleParserG_N            = 63
	GraffleParserV_N            = 64
	GraffleParserE_N            = 65
	GraffleParserWHERE          = 66
	GraffleParserVAR            = 67
	GraffleParserID             = 68
	GraffleParserWS             = 69
	GraffleParserLINE_COMMENT   = 70
	GraffleParserM_LINE_COMMENT = 71
)

// GraffleParser rules.
const (
	GraffleParserRULE_file                            = 0
	GraffleParserRULE_sequence                        = 1
	GraffleParserRULE_sequence_element                = 2
	GraffleParserRULE_sequence_line                   = 3
	GraffleParserRULE_one_line_sequence_element       = 4
	GraffleParserRULE_atom_action                     = 5
	GraffleParserRULE_if_stmnt                        = 6
	GraffleParserRULE_if_is_stmnt                     = 7
	GraffleParserRULE_case_stmnt                      = 8
	GraffleParserRULE_one_line_stmnt                  = 9
	GraffleParserRULE_mult_line_stmnt                 = 10
	GraffleParserRULE_stmnt                           = 11
	GraffleParserRULE_while_stmnt                     = 12
	GraffleParserRULE_until_stmnt                     = 13
	GraffleParserRULE_for_stmnt                       = 14
	GraffleParserRULE_from_to_stmnt                   = 15
	GraffleParserRULE_function_declaration            = 16
	GraffleParserRULE_one_line_function_declaration   = 17
	GraffleParserRULE_mult_line_function_declaration  = 18
	GraffleParserRULE_function_declaration_head       = 19
	GraffleParserRULE_one_line_procedure_declaration  = 20
	GraffleParserRULE_mult_line_procedure_declaration = 21
	GraffleParserRULE_procedure_declaration_head      = 22
	GraffleParserRULE_var_declaration                 = 23
	GraffleParserRULE_arc_declaration                 = 24
	GraffleParserRULE_arc                             = 25
	GraffleParserRULE_or_w_arc_lr                     = 26
	GraffleParserRULE_or_w_arc_rl                     = 27
	GraffleParserRULE_unor_w_arc                      = 28
	GraffleParserRULE_vertice_declaration             = 29
	GraffleParserRULE_graph_declaration               = 30
	GraffleParserRULE_labeled_declaration             = 31
	GraffleParserRULE_expr                            = 32
	GraffleParserRULE_integral_expr                   = 33
	GraffleParserRULE_logical_expr                    = 34
	GraffleParserRULE_log_expr_operand                = 35
	GraffleParserRULE_bin_log_operator                = 36
	GraffleParserRULE_unar_log_operator               = 37
	GraffleParserRULE_arithm_expr                     = 38
	GraffleParserRULE_arithm_expr_operand             = 39
	GraffleParserRULE_bin_arithm_operator             = 40
	GraffleParserRULE_arithm_assign_operator          = 41
	GraffleParserRULE_builtin_function_call           = 42
	GraffleParserRULE_built_func_print                = 43
	GraffleParserRULE_built_func_input                = 44
	GraffleParserRULE_function_call                   = 45
	GraffleParserRULE_label                           = 46
	GraffleParserRULE_value                           = 47
	GraffleParserRULE_builtin                         = 48
	GraffleParserRULE_block_end                       = 49
)

// IFileContext is an interface to support dynamic dispatch.
type IFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetMain_body returns the main_body rule contexts.
	GetMain_body() ISequenceContext

	// SetMain_body sets the main_body rule contexts.
	SetMain_body(ISequenceContext)

	// IsFileContext differentiates from other interfaces.
	IsFileContext()
}

type FileContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	main_body ISequenceContext
}

func NewEmptyFileContext() *FileContext {
	var p = new(FileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_file
	return p
}

func (*FileContext) IsFileContext() {}

func NewFileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FileContext {
	var p = new(FileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_file

	return p
}

func (s *FileContext) GetParser() antlr.Parser { return s.parser }

func (s *FileContext) GetMain_body() ISequenceContext { return s.main_body }

func (s *FileContext) SetMain_body(v ISequenceContext) { s.main_body = v }

func (s *FileContext) Sequence() ISequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISequenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISequenceContext)
}

func (s *FileContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(GraffleParserNEWLINE, 0)
}

func (s *FileContext) AllFunction_declaration() []IFunction_declarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFunction_declarationContext)(nil)).Elem())
	var tst = make([]IFunction_declarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFunction_declarationContext)
		}
	}

	return tst
}

func (s *FileContext) Function_declaration(i int) IFunction_declarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunction_declarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFunction_declarationContext)
}

func (s *FileContext) EOF() antlr.TerminalNode {
	return s.GetToken(GraffleParserEOF, 0)
}

func (s *FileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterFile(s)
	}
}

func (s *FileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitFile(s)
	}
}

func (s *FileContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitFile(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) File() (localctx IFileContext) {
	localctx = NewFileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GraffleParserRULE_file)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(101)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(100)
			p.Match(GraffleParserNEWLINE)
		}

	}
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(103)
				p.Function_declaration()
			}

		}
		p.SetState(108)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
	}
	{
		p.SetState(109)

		var _x = p.Sequence()

		localctx.(*FileContext).main_body = _x
	}
	p.SetState(111)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(110)
			p.Match(GraffleParserEOF)
		}

	}

	return localctx
}

// ISequenceContext is an interface to support dynamic dispatch.
type ISequenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSequenceContext differentiates from other interfaces.
	IsSequenceContext()
}

type SequenceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySequenceContext() *SequenceContext {
	var p = new(SequenceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_sequence
	return p
}

func (*SequenceContext) IsSequenceContext() {}

func NewSequenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SequenceContext {
	var p = new(SequenceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_sequence

	return p
}

func (s *SequenceContext) GetParser() antlr.Parser { return s.parser }

func (s *SequenceContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(GraffleParserNEWLINE)
}

func (s *SequenceContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(GraffleParserNEWLINE, i)
}

func (s *SequenceContext) AllSequence_element() []ISequence_elementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISequence_elementContext)(nil)).Elem())
	var tst = make([]ISequence_elementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISequence_elementContext)
		}
	}

	return tst
}

func (s *SequenceContext) Sequence_element(i int) ISequence_elementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISequence_elementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISequence_elementContext)
}

func (s *SequenceContext) AllEOF() []antlr.TerminalNode {
	return s.GetTokens(GraffleParserEOF)
}

func (s *SequenceContext) EOF(i int) antlr.TerminalNode {
	return s.GetToken(GraffleParserEOF, i)
}

func (s *SequenceContext) AllACT_DELIM() []antlr.TerminalNode {
	return s.GetTokens(GraffleParserACT_DELIM)
}

func (s *SequenceContext) ACT_DELIM(i int) antlr.TerminalNode {
	return s.GetToken(GraffleParserACT_DELIM, i)
}

func (s *SequenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SequenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SequenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterSequence(s)
	}
}

func (s *SequenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitSequence(s)
	}
}

func (s *SequenceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitSequence(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Sequence() (localctx ISequenceContext) {
	localctx = NewSequenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GraffleParserRULE_sequence)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraffleParserNEWLINE {
		{
			p.SetState(113)
			p.Match(GraffleParserNEWLINE)
		}

	}
	p.SetState(125)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(116)
				p.Sequence_element()
			}
			p.SetState(123)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case GraffleParserACT_DELIM, GraffleParserNEWLINE:
				p.SetState(118)
				p.GetErrorHandler().Sync(p)
				_alt = 1
				for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
					switch _alt {
					case 1:
						{
							p.SetState(117)
							_la = p.GetTokenStream().LA(1)

							if !(_la == GraffleParserACT_DELIM || _la == GraffleParserNEWLINE) {
								p.GetErrorHandler().RecoverInline(p)
							} else {
								p.GetErrorHandler().ReportMatch(p)
								p.Consume()
							}
						}

					default:
						panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
					}

					p.SetState(120)
					p.GetErrorHandler().Sync(p)
					_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
				}

			case GraffleParserEOF:
				{
					p.SetState(122)
					p.Match(GraffleParserEOF)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(127)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}

	return localctx
}

// ISequence_elementContext is an interface to support dynamic dispatch.
type ISequence_elementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSequence_elementContext differentiates from other interfaces.
	IsSequence_elementContext()
}

type Sequence_elementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySequence_elementContext() *Sequence_elementContext {
	var p = new(Sequence_elementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_sequence_element
	return p
}

func (*Sequence_elementContext) IsSequence_elementContext() {}

func NewSequence_elementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Sequence_elementContext {
	var p = new(Sequence_elementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_sequence_element

	return p
}

func (s *Sequence_elementContext) GetParser() antlr.Parser { return s.parser }

func (s *Sequence_elementContext) If_stmnt() IIf_stmntContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIf_stmntContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIf_stmntContext)
}

func (s *Sequence_elementContext) If_is_stmnt() IIf_is_stmntContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIf_is_stmntContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIf_is_stmntContext)
}

func (s *Sequence_elementContext) Mult_line_stmnt() IMult_line_stmntContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMult_line_stmntContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMult_line_stmntContext)
}

func (s *Sequence_elementContext) Sequence_line() ISequence_lineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISequence_lineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISequence_lineContext)
}

func (s *Sequence_elementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Sequence_elementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Sequence_elementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterSequence_element(s)
	}
}

func (s *Sequence_elementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitSequence_element(s)
	}
}

func (s *Sequence_elementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitSequence_element(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Sequence_element() (localctx ISequence_elementContext) {
	localctx = NewSequence_elementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, GraffleParserRULE_sequence_element)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(129)
			p.If_stmnt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(130)
			p.If_is_stmnt()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(131)
			p.Mult_line_stmnt()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(132)
			p.Sequence_line()
		}

	}

	return localctx
}

// ISequence_lineContext is an interface to support dynamic dispatch.
type ISequence_lineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSequence_lineContext differentiates from other interfaces.
	IsSequence_lineContext()
}

type Sequence_lineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySequence_lineContext() *Sequence_lineContext {
	var p = new(Sequence_lineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_sequence_line
	return p
}

func (*Sequence_lineContext) IsSequence_lineContext() {}

func NewSequence_lineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Sequence_lineContext {
	var p = new(Sequence_lineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_sequence_line

	return p
}

func (s *Sequence_lineContext) GetParser() antlr.Parser { return s.parser }

func (s *Sequence_lineContext) AllOne_line_sequence_element() []IOne_line_sequence_elementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOne_line_sequence_elementContext)(nil)).Elem())
	var tst = make([]IOne_line_sequence_elementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOne_line_sequence_elementContext)
		}
	}

	return tst
}

func (s *Sequence_lineContext) One_line_sequence_element(i int) IOne_line_sequence_elementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOne_line_sequence_elementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOne_line_sequence_elementContext)
}

func (s *Sequence_lineContext) AllACT_DELIM() []antlr.TerminalNode {
	return s.GetTokens(GraffleParserACT_DELIM)
}

func (s *Sequence_lineContext) ACT_DELIM(i int) antlr.TerminalNode {
	return s.GetToken(GraffleParserACT_DELIM, i)
}

func (s *Sequence_lineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Sequence_lineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Sequence_lineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterSequence_line(s)
	}
}

func (s *Sequence_lineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitSequence_line(s)
	}
}

func (s *Sequence_lineContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitSequence_line(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Sequence_line() (localctx ISequence_lineContext) {
	localctx = NewSequence_lineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, GraffleParserRULE_sequence_line)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(135)
		p.One_line_sequence_element()
	}
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(136)
				p.Match(GraffleParserACT_DELIM)
			}
			{
				p.SetState(137)
				p.One_line_sequence_element()
			}

		}
		p.SetState(142)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
	}
	p.SetState(144)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(143)
			p.Match(GraffleParserACT_DELIM)
		}

	}

	return localctx
}

// IOne_line_sequence_elementContext is an interface to support dynamic dispatch.
type IOne_line_sequence_elementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOne_line_sequence_elementContext differentiates from other interfaces.
	IsOne_line_sequence_elementContext()
}

type One_line_sequence_elementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOne_line_sequence_elementContext() *One_line_sequence_elementContext {
	var p = new(One_line_sequence_elementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_one_line_sequence_element
	return p
}

func (*One_line_sequence_elementContext) IsOne_line_sequence_elementContext() {}

func NewOne_line_sequence_elementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *One_line_sequence_elementContext {
	var p = new(One_line_sequence_elementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_one_line_sequence_element

	return p
}

func (s *One_line_sequence_elementContext) GetParser() antlr.Parser { return s.parser }

func (s *One_line_sequence_elementContext) Atom_action() IAtom_actionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtom_actionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtom_actionContext)
}

func (s *One_line_sequence_elementContext) One_line_stmnt() IOne_line_stmntContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOne_line_stmntContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOne_line_stmntContext)
}

func (s *One_line_sequence_elementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *One_line_sequence_elementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *One_line_sequence_elementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterOne_line_sequence_element(s)
	}
}

func (s *One_line_sequence_elementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitOne_line_sequence_element(s)
	}
}

func (s *One_line_sequence_elementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitOne_line_sequence_element(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) One_line_sequence_element() (localctx IOne_line_sequence_elementContext) {
	localctx = NewOne_line_sequence_elementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, GraffleParserRULE_one_line_sequence_element)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(148)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GraffleParserPRINTER, GraffleParserKEY_INPUT, GraffleParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(146)
			p.Atom_action()
		}

	case GraffleParserWHILE, GraffleParserUNTIL, GraffleParserFOR, GraffleParserFROM:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(147)
			p.One_line_stmnt()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAtom_actionContext is an interface to support dynamic dispatch.
type IAtom_actionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtom_actionContext differentiates from other interfaces.
	IsAtom_actionContext()
}

type Atom_actionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtom_actionContext() *Atom_actionContext {
	var p = new(Atom_actionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_atom_action
	return p
}

func (*Atom_actionContext) IsAtom_actionContext() {}

func NewAtom_actionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Atom_actionContext {
	var p = new(Atom_actionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_atom_action

	return p
}

func (s *Atom_actionContext) GetParser() antlr.Parser { return s.parser }

func (s *Atom_actionContext) Var_declaration() IVar_declarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVar_declarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVar_declarationContext)
}

func (s *Atom_actionContext) Function_call() IFunction_callContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunction_callContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunction_callContext)
}

func (s *Atom_actionContext) Builtin_function_call() IBuiltin_function_callContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBuiltin_function_callContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBuiltin_function_callContext)
}

func (s *Atom_actionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Atom_actionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Atom_actionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterAtom_action(s)
	}
}

func (s *Atom_actionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitAtom_action(s)
	}
}

func (s *Atom_actionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitAtom_action(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Atom_action() (localctx IAtom_actionContext) {
	localctx = NewAtom_actionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, GraffleParserRULE_atom_action)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(153)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(150)
			p.Var_declaration()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(151)
			p.Function_call()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(152)
			p.Builtin_function_call()
		}

	}

	return localctx
}

// IIf_stmntContext is an interface to support dynamic dispatch.
type IIf_stmntContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetCond returns the cond rule contexts.
	GetCond() ILogical_exprContext

	// SetCond sets the cond rule contexts.
	SetCond(ILogical_exprContext)

	// IsIf_stmntContext differentiates from other interfaces.
	IsIf_stmntContext()
}

type If_stmntContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	cond   ILogical_exprContext
}

func NewEmptyIf_stmntContext() *If_stmntContext {
	var p = new(If_stmntContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_if_stmnt
	return p
}

func (*If_stmntContext) IsIf_stmntContext() {}

func NewIf_stmntContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_stmntContext {
	var p = new(If_stmntContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_if_stmnt

	return p
}

func (s *If_stmntContext) GetParser() antlr.Parser { return s.parser }

func (s *If_stmntContext) GetCond() ILogical_exprContext { return s.cond }

func (s *If_stmntContext) SetCond(v ILogical_exprContext) { s.cond = v }

func (s *If_stmntContext) IF() antlr.TerminalNode {
	return s.GetToken(GraffleParserIF, 0)
}

func (s *If_stmntContext) Block_end() IBlock_endContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlock_endContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlock_endContext)
}

func (s *If_stmntContext) Logical_expr() ILogical_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILogical_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILogical_exprContext)
}

func (s *If_stmntContext) AllSequence() []ISequenceContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISequenceContext)(nil)).Elem())
	var tst = make([]ISequenceContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISequenceContext)
		}
	}

	return tst
}

func (s *If_stmntContext) Sequence(i int) ISequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISequenceContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISequenceContext)
}

func (s *If_stmntContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(GraffleParserNEWLINE, 0)
}

func (s *If_stmntContext) ELSE() antlr.TerminalNode {
	return s.GetToken(GraffleParserELSE, 0)
}

func (s *If_stmntContext) THEN() antlr.TerminalNode {
	return s.GetToken(GraffleParserTHEN, 0)
}

func (s *If_stmntContext) AllDO() []antlr.TerminalNode {
	return s.GetTokens(GraffleParserDO)
}

func (s *If_stmntContext) DO(i int) antlr.TerminalNode {
	return s.GetToken(GraffleParserDO, i)
}

func (s *If_stmntContext) ARG_DELIM() antlr.TerminalNode {
	return s.GetToken(GraffleParserARG_DELIM, 0)
}

func (s *If_stmntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_stmntContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *If_stmntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterIf_stmnt(s)
	}
}

func (s *If_stmntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitIf_stmnt(s)
	}
}

func (s *If_stmntContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitIf_stmnt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) If_stmnt() (localctx IIf_stmntContext) {
	localctx = NewIf_stmntContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, GraffleParserRULE_if_stmnt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(155)
		p.Match(GraffleParserIF)
	}
	{
		p.SetState(156)

		var _x = p.Logical_expr()

		localctx.(*If_stmntContext).cond = _x
	}
	p.SetState(166)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-42)&-(0x1f+1)) == 0 && ((1<<uint((_la-42)))&((1<<(GraffleParserTHEN-42))|(1<<(GraffleParserDO-42))|(1<<(GraffleParserARG_DELIM-42)))) != 0 {
		p.SetState(158)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraffleParserARG_DELIM {
			{
				p.SetState(157)
				p.Match(GraffleParserARG_DELIM)
			}

		}
		p.SetState(164)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(160)
				p.Match(GraffleParserTHEN)
			}

		case 2:
			{
				p.SetState(161)
				p.Match(GraffleParserDO)
			}

		case 3:
			{
				p.SetState(162)
				p.Match(GraffleParserTHEN)
			}
			{
				p.SetState(163)
				p.Match(GraffleParserDO)
			}

		}

	}
	p.SetState(169)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(168)
				p.Sequence()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(171)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())
	}
	p.SetState(174)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(173)
			p.Match(GraffleParserNEWLINE)
		}

	}
	p.SetState(185)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraffleParserELSE {
		{
			p.SetState(176)
			p.Match(GraffleParserELSE)
		}
		p.SetState(178)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraffleParserDO {
			{
				p.SetState(177)
				p.Match(GraffleParserDO)
			}

		}
		p.SetState(181)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(180)
					p.Sequence()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(183)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())
		}

	}
	{
		p.SetState(187)
		p.Block_end()
	}

	return localctx
}

// IIf_is_stmntContext is an interface to support dynamic dispatch.
type IIf_is_stmntContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIf_is_stmntContext differentiates from other interfaces.
	IsIf_is_stmntContext()
}

type If_is_stmntContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_is_stmntContext() *If_is_stmntContext {
	var p = new(If_is_stmntContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_if_is_stmnt
	return p
}

func (*If_is_stmntContext) IsIf_is_stmntContext() {}

func NewIf_is_stmntContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_is_stmntContext {
	var p = new(If_is_stmntContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_if_is_stmnt

	return p
}

func (s *If_is_stmntContext) GetParser() antlr.Parser { return s.parser }

func (s *If_is_stmntContext) IF() antlr.TerminalNode {
	return s.GetToken(GraffleParserIF, 0)
}

func (s *If_is_stmntContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *If_is_stmntContext) Block_end() IBlock_endContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlock_endContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlock_endContext)
}

func (s *If_is_stmntContext) AllCase_stmnt() []ICase_stmntContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICase_stmntContext)(nil)).Elem())
	var tst = make([]ICase_stmntContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICase_stmntContext)
		}
	}

	return tst
}

func (s *If_is_stmntContext) Case_stmnt(i int) ICase_stmntContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICase_stmntContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICase_stmntContext)
}

func (s *If_is_stmntContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(GraffleParserNEWLINE, 0)
}

func (s *If_is_stmntContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(GraffleParserDEFAULT, 0)
}

func (s *If_is_stmntContext) DO() antlr.TerminalNode {
	return s.GetToken(GraffleParserDO, 0)
}

func (s *If_is_stmntContext) AllSequence() []ISequenceContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISequenceContext)(nil)).Elem())
	var tst = make([]ISequenceContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISequenceContext)
		}
	}

	return tst
}

func (s *If_is_stmntContext) Sequence(i int) ISequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISequenceContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISequenceContext)
}

func (s *If_is_stmntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_is_stmntContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *If_is_stmntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterIf_is_stmnt(s)
	}
}

func (s *If_is_stmntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitIf_is_stmnt(s)
	}
}

func (s *If_is_stmntContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitIf_is_stmnt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) If_is_stmnt() (localctx IIf_is_stmntContext) {
	localctx = NewIf_is_stmntContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, GraffleParserRULE_if_is_stmnt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(189)
		p.Match(GraffleParserIF)
	}
	{
		p.SetState(190)
		p.Value()
	}
	p.SetState(194)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(191)
				p.Case_stmnt()
			}

		}
		p.SetState(196)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())
	}
	p.SetState(198)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(197)
			p.Match(GraffleParserNEWLINE)
		}

	}
	p.SetState(209)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraffleParserDEFAULT {
		{
			p.SetState(200)
			p.Match(GraffleParserDEFAULT)
		}
		p.SetState(202)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraffleParserDO {
			{
				p.SetState(201)
				p.Match(GraffleParserDO)
			}

		}
		p.SetState(205)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(204)
					p.Sequence()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(207)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())
		}

	}
	{
		p.SetState(211)
		p.Block_end()
	}

	return localctx
}

// ICase_stmntContext is an interface to support dynamic dispatch.
type ICase_stmntContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCase_stmntContext differentiates from other interfaces.
	IsCase_stmntContext()
}

type Case_stmntContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCase_stmntContext() *Case_stmntContext {
	var p = new(Case_stmntContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_case_stmnt
	return p
}

func (*Case_stmntContext) IsCase_stmntContext() {}

func NewCase_stmntContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Case_stmntContext {
	var p = new(Case_stmntContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_case_stmnt

	return p
}

func (s *Case_stmntContext) GetParser() antlr.Parser { return s.parser }

func (s *Case_stmntContext) IS() antlr.TerminalNode {
	return s.GetToken(GraffleParserIS, 0)
}

func (s *Case_stmntContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *Case_stmntContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(GraffleParserNEWLINE, 0)
}

func (s *Case_stmntContext) AllSequence() []ISequenceContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISequenceContext)(nil)).Elem())
	var tst = make([]ISequenceContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISequenceContext)
		}
	}

	return tst
}

func (s *Case_stmntContext) Sequence(i int) ISequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISequenceContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISequenceContext)
}

func (s *Case_stmntContext) THEN() antlr.TerminalNode {
	return s.GetToken(GraffleParserTHEN, 0)
}

func (s *Case_stmntContext) DO() antlr.TerminalNode {
	return s.GetToken(GraffleParserDO, 0)
}

func (s *Case_stmntContext) ARG_DELIM() antlr.TerminalNode {
	return s.GetToken(GraffleParserARG_DELIM, 0)
}

func (s *Case_stmntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Case_stmntContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Case_stmntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterCase_stmnt(s)
	}
}

func (s *Case_stmntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitCase_stmnt(s)
	}
}

func (s *Case_stmntContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitCase_stmnt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Case_stmnt() (localctx ICase_stmntContext) {
	localctx = NewCase_stmntContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, GraffleParserRULE_case_stmnt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(214)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraffleParserNEWLINE {
		{
			p.SetState(213)
			p.Match(GraffleParserNEWLINE)
		}

	}
	{
		p.SetState(216)
		p.Match(GraffleParserIS)
	}
	{
		p.SetState(217)
		p.Value()
	}
	p.SetState(227)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-42)&-(0x1f+1)) == 0 && ((1<<uint((_la-42)))&((1<<(GraffleParserTHEN-42))|(1<<(GraffleParserDO-42))|(1<<(GraffleParserARG_DELIM-42)))) != 0 {
		p.SetState(219)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraffleParserARG_DELIM {
			{
				p.SetState(218)
				p.Match(GraffleParserARG_DELIM)
			}

		}
		p.SetState(225)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(221)
				p.Match(GraffleParserTHEN)
			}

		case 2:
			{
				p.SetState(222)
				p.Match(GraffleParserDO)
			}

		case 3:
			{
				p.SetState(223)
				p.Match(GraffleParserTHEN)
			}
			{
				p.SetState(224)
				p.Match(GraffleParserDO)
			}

		}

	}
	p.SetState(230)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(229)
				p.Sequence()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(232)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext())
	}

	return localctx
}

// IOne_line_stmntContext is an interface to support dynamic dispatch.
type IOne_line_stmntContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOne_line_stmntContext differentiates from other interfaces.
	IsOne_line_stmntContext()
}

type One_line_stmntContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOne_line_stmntContext() *One_line_stmntContext {
	var p = new(One_line_stmntContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_one_line_stmnt
	return p
}

func (*One_line_stmntContext) IsOne_line_stmntContext() {}

func NewOne_line_stmntContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *One_line_stmntContext {
	var p = new(One_line_stmntContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_one_line_stmnt

	return p
}

func (s *One_line_stmntContext) GetParser() antlr.Parser { return s.parser }

func (s *One_line_stmntContext) Stmnt() IStmntContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStmntContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStmntContext)
}

func (s *One_line_stmntContext) Sequence_line() ISequence_lineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISequence_lineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISequence_lineContext)
}

func (s *One_line_stmntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *One_line_stmntContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *One_line_stmntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterOne_line_stmnt(s)
	}
}

func (s *One_line_stmntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitOne_line_stmnt(s)
	}
}

func (s *One_line_stmntContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitOne_line_stmnt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) One_line_stmnt() (localctx IOne_line_stmntContext) {
	localctx = NewOne_line_stmntContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, GraffleParserRULE_one_line_stmnt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(234)
		p.Stmnt()
	}
	{
		p.SetState(235)
		p.Sequence_line()
	}

	return localctx
}

// IMult_line_stmntContext is an interface to support dynamic dispatch.
type IMult_line_stmntContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMult_line_stmntContext differentiates from other interfaces.
	IsMult_line_stmntContext()
}

type Mult_line_stmntContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMult_line_stmntContext() *Mult_line_stmntContext {
	var p = new(Mult_line_stmntContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_mult_line_stmnt
	return p
}

func (*Mult_line_stmntContext) IsMult_line_stmntContext() {}

func NewMult_line_stmntContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Mult_line_stmntContext {
	var p = new(Mult_line_stmntContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_mult_line_stmnt

	return p
}

func (s *Mult_line_stmntContext) GetParser() antlr.Parser { return s.parser }

func (s *Mult_line_stmntContext) Stmnt() IStmntContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStmntContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStmntContext)
}

func (s *Mult_line_stmntContext) Sequence() ISequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISequenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISequenceContext)
}

func (s *Mult_line_stmntContext) Block_end() IBlock_endContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlock_endContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlock_endContext)
}

func (s *Mult_line_stmntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Mult_line_stmntContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Mult_line_stmntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterMult_line_stmnt(s)
	}
}

func (s *Mult_line_stmntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitMult_line_stmnt(s)
	}
}

func (s *Mult_line_stmntContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitMult_line_stmnt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Mult_line_stmnt() (localctx IMult_line_stmntContext) {
	localctx = NewMult_line_stmntContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, GraffleParserRULE_mult_line_stmnt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(237)
		p.Stmnt()
	}
	{
		p.SetState(238)
		p.Sequence()
	}
	{
		p.SetState(239)
		p.Block_end()
	}

	return localctx
}

// IStmntContext is an interface to support dynamic dispatch.
type IStmntContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStmntContext differentiates from other interfaces.
	IsStmntContext()
}

type StmntContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmntContext() *StmntContext {
	var p = new(StmntContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_stmnt
	return p
}

func (*StmntContext) IsStmntContext() {}

func NewStmntContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmntContext {
	var p = new(StmntContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_stmnt

	return p
}

func (s *StmntContext) GetParser() antlr.Parser { return s.parser }

func (s *StmntContext) While_stmnt() IWhile_stmntContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhile_stmntContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhile_stmntContext)
}

func (s *StmntContext) Until_stmnt() IUntil_stmntContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUntil_stmntContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUntil_stmntContext)
}

func (s *StmntContext) For_stmnt() IFor_stmntContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFor_stmntContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFor_stmntContext)
}

func (s *StmntContext) From_to_stmnt() IFrom_to_stmntContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFrom_to_stmntContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFrom_to_stmntContext)
}

func (s *StmntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmntContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterStmnt(s)
	}
}

func (s *StmntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitStmnt(s)
	}
}

func (s *StmntContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitStmnt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Stmnt() (localctx IStmntContext) {
	localctx = NewStmntContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, GraffleParserRULE_stmnt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(245)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GraffleParserWHILE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(241)
			p.While_stmnt()
		}

	case GraffleParserUNTIL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(242)
			p.Until_stmnt()
		}

	case GraffleParserFOR:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(243)
			p.For_stmnt()
		}

	case GraffleParserFROM:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(244)
			p.From_to_stmnt()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IWhile_stmntContext is an interface to support dynamic dispatch.
type IWhile_stmntContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetCond returns the cond rule contexts.
	GetCond() ILogical_exprContext

	// SetCond sets the cond rule contexts.
	SetCond(ILogical_exprContext)

	// IsWhile_stmntContext differentiates from other interfaces.
	IsWhile_stmntContext()
}

type While_stmntContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	cond   ILogical_exprContext
}

func NewEmptyWhile_stmntContext() *While_stmntContext {
	var p = new(While_stmntContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_while_stmnt
	return p
}

func (*While_stmntContext) IsWhile_stmntContext() {}

func NewWhile_stmntContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *While_stmntContext {
	var p = new(While_stmntContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_while_stmnt

	return p
}

func (s *While_stmntContext) GetParser() antlr.Parser { return s.parser }

func (s *While_stmntContext) GetCond() ILogical_exprContext { return s.cond }

func (s *While_stmntContext) SetCond(v ILogical_exprContext) { s.cond = v }

func (s *While_stmntContext) WHILE() antlr.TerminalNode {
	return s.GetToken(GraffleParserWHILE, 0)
}

func (s *While_stmntContext) Logical_expr() ILogical_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILogical_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILogical_exprContext)
}

func (s *While_stmntContext) DO() antlr.TerminalNode {
	return s.GetToken(GraffleParserDO, 0)
}

func (s *While_stmntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *While_stmntContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *While_stmntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterWhile_stmnt(s)
	}
}

func (s *While_stmntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitWhile_stmnt(s)
	}
}

func (s *While_stmntContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitWhile_stmnt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) While_stmnt() (localctx IWhile_stmntContext) {
	localctx = NewWhile_stmntContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, GraffleParserRULE_while_stmnt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(247)
		p.Match(GraffleParserWHILE)
	}
	{
		p.SetState(248)

		var _x = p.Logical_expr()

		localctx.(*While_stmntContext).cond = _x
	}
	p.SetState(250)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraffleParserDO {
		{
			p.SetState(249)
			p.Match(GraffleParserDO)
		}

	}

	return localctx
}

// IUntil_stmntContext is an interface to support dynamic dispatch.
type IUntil_stmntContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetCond returns the cond rule contexts.
	GetCond() ILogical_exprContext

	// SetCond sets the cond rule contexts.
	SetCond(ILogical_exprContext)

	// IsUntil_stmntContext differentiates from other interfaces.
	IsUntil_stmntContext()
}

type Until_stmntContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	cond   ILogical_exprContext
}

func NewEmptyUntil_stmntContext() *Until_stmntContext {
	var p = new(Until_stmntContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_until_stmnt
	return p
}

func (*Until_stmntContext) IsUntil_stmntContext() {}

func NewUntil_stmntContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Until_stmntContext {
	var p = new(Until_stmntContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_until_stmnt

	return p
}

func (s *Until_stmntContext) GetParser() antlr.Parser { return s.parser }

func (s *Until_stmntContext) GetCond() ILogical_exprContext { return s.cond }

func (s *Until_stmntContext) SetCond(v ILogical_exprContext) { s.cond = v }

func (s *Until_stmntContext) UNTIL() antlr.TerminalNode {
	return s.GetToken(GraffleParserUNTIL, 0)
}

func (s *Until_stmntContext) Logical_expr() ILogical_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILogical_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILogical_exprContext)
}

func (s *Until_stmntContext) DO() antlr.TerminalNode {
	return s.GetToken(GraffleParserDO, 0)
}

func (s *Until_stmntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Until_stmntContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Until_stmntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterUntil_stmnt(s)
	}
}

func (s *Until_stmntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitUntil_stmnt(s)
	}
}

func (s *Until_stmntContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitUntil_stmnt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Until_stmnt() (localctx IUntil_stmntContext) {
	localctx = NewUntil_stmntContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, GraffleParserRULE_until_stmnt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(252)
		p.Match(GraffleParserUNTIL)
	}
	{
		p.SetState(253)

		var _x = p.Logical_expr()

		localctx.(*Until_stmntContext).cond = _x
	}
	p.SetState(255)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraffleParserDO {
		{
			p.SetState(254)
			p.Match(GraffleParserDO)
		}

	}

	return localctx
}

// IFor_stmntContext is an interface to support dynamic dispatch.
type IFor_stmntContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFor_stmntContext differentiates from other interfaces.
	IsFor_stmntContext()
}

type For_stmntContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFor_stmntContext() *For_stmntContext {
	var p = new(For_stmntContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_for_stmnt
	return p
}

func (*For_stmntContext) IsFor_stmntContext() {}

func NewFor_stmntContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *For_stmntContext {
	var p = new(For_stmntContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_for_stmnt

	return p
}

func (s *For_stmntContext) GetParser() antlr.Parser { return s.parser }

func (s *For_stmntContext) CopyFrom(ctx *For_stmntContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *For_stmntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *For_stmntContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ForLogicalContext struct {
	*For_stmntContext
	cond ILogical_exprContext
}

func NewForLogicalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForLogicalContext {
	var p = new(ForLogicalContext)

	p.For_stmntContext = NewEmptyFor_stmntContext()
	p.parser = parser
	p.CopyFrom(ctx.(*For_stmntContext))

	return p
}

func (s *ForLogicalContext) GetCond() ILogical_exprContext { return s.cond }

func (s *ForLogicalContext) SetCond(v ILogical_exprContext) { s.cond = v }

func (s *ForLogicalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForLogicalContext) FOR() antlr.TerminalNode {
	return s.GetToken(GraffleParserFOR, 0)
}

func (s *ForLogicalContext) Logical_expr() ILogical_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILogical_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILogical_exprContext)
}

func (s *ForLogicalContext) DO() antlr.TerminalNode {
	return s.GetToken(GraffleParserDO, 0)
}

func (s *ForLogicalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterForLogical(s)
	}
}

func (s *ForLogicalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitForLogical(s)
	}
}

func (s *ForLogicalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitForLogical(s)

	default:
		return t.VisitChildren(s)
	}
}

type ForVarContext struct {
	*For_stmntContext
	pre_act  IAtom_actionContext
	cond     ILogical_exprContext
	post_act IAtom_actionContext
}

func NewForVarContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForVarContext {
	var p = new(ForVarContext)

	p.For_stmntContext = NewEmptyFor_stmntContext()
	p.parser = parser
	p.CopyFrom(ctx.(*For_stmntContext))

	return p
}

func (s *ForVarContext) GetPre_act() IAtom_actionContext { return s.pre_act }

func (s *ForVarContext) GetCond() ILogical_exprContext { return s.cond }

func (s *ForVarContext) GetPost_act() IAtom_actionContext { return s.post_act }

func (s *ForVarContext) SetPre_act(v IAtom_actionContext) { s.pre_act = v }

func (s *ForVarContext) SetCond(v ILogical_exprContext) { s.cond = v }

func (s *ForVarContext) SetPost_act(v IAtom_actionContext) { s.post_act = v }

func (s *ForVarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForVarContext) FOR() antlr.TerminalNode {
	return s.GetToken(GraffleParserFOR, 0)
}

func (s *ForVarContext) AllARG_DELIM() []antlr.TerminalNode {
	return s.GetTokens(GraffleParserARG_DELIM)
}

func (s *ForVarContext) ARG_DELIM(i int) antlr.TerminalNode {
	return s.GetToken(GraffleParserARG_DELIM, i)
}

func (s *ForVarContext) AllAtom_action() []IAtom_actionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAtom_actionContext)(nil)).Elem())
	var tst = make([]IAtom_actionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAtom_actionContext)
		}
	}

	return tst
}

func (s *ForVarContext) Atom_action(i int) IAtom_actionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtom_actionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAtom_actionContext)
}

func (s *ForVarContext) Logical_expr() ILogical_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILogical_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILogical_exprContext)
}

func (s *ForVarContext) DO() antlr.TerminalNode {
	return s.GetToken(GraffleParserDO, 0)
}

func (s *ForVarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterForVar(s)
	}
}

func (s *ForVarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitForVar(s)
	}
}

func (s *ForVarContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitForVar(s)

	default:
		return t.VisitChildren(s)
	}
}

type ForRangeContext struct {
	*For_stmntContext
	from IExprContext
	to   IExprContext
}

func NewForRangeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForRangeContext {
	var p = new(ForRangeContext)

	p.For_stmntContext = NewEmptyFor_stmntContext()
	p.parser = parser
	p.CopyFrom(ctx.(*For_stmntContext))

	return p
}

func (s *ForRangeContext) GetFrom() IExprContext { return s.from }

func (s *ForRangeContext) GetTo() IExprContext { return s.to }

func (s *ForRangeContext) SetFrom(v IExprContext) { s.from = v }

func (s *ForRangeContext) SetTo(v IExprContext) { s.to = v }

func (s *ForRangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForRangeContext) FOR() antlr.TerminalNode {
	return s.GetToken(GraffleParserFOR, 0)
}

func (s *ForRangeContext) VAR() antlr.TerminalNode {
	return s.GetToken(GraffleParserVAR, 0)
}

func (s *ForRangeContext) IN() antlr.TerminalNode {
	return s.GetToken(GraffleParserIN, 0)
}

func (s *ForRangeContext) TO() antlr.TerminalNode {
	return s.GetToken(GraffleParserTO, 0)
}

func (s *ForRangeContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ForRangeContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ForRangeContext) RANGE() antlr.TerminalNode {
	return s.GetToken(GraffleParserRANGE, 0)
}

func (s *ForRangeContext) FROM() antlr.TerminalNode {
	return s.GetToken(GraffleParserFROM, 0)
}

func (s *ForRangeContext) DO() antlr.TerminalNode {
	return s.GetToken(GraffleParserDO, 0)
}

func (s *ForRangeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterForRange(s)
	}
}

func (s *ForRangeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitForRange(s)
	}
}

func (s *ForRangeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitForRange(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) For_stmnt() (localctx IFor_stmntContext) {
	localctx = NewFor_stmntContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, GraffleParserRULE_for_stmnt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(286)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext()) {
	case 1:
		localctx = NewForLogicalContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(257)
			p.Match(GraffleParserFOR)
		}
		{
			p.SetState(258)

			var _x = p.Logical_expr()

			localctx.(*ForLogicalContext).cond = _x
		}
		p.SetState(260)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraffleParserDO {
			{
				p.SetState(259)
				p.Match(GraffleParserDO)
			}

		}

	case 2:
		localctx = NewForVarContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(262)
			p.Match(GraffleParserFOR)
		}
		{
			p.SetState(263)

			var _x = p.Atom_action()

			localctx.(*ForVarContext).pre_act = _x
		}
		{
			p.SetState(264)
			p.Match(GraffleParserARG_DELIM)
		}
		{
			p.SetState(265)

			var _x = p.Logical_expr()

			localctx.(*ForVarContext).cond = _x
		}
		{
			p.SetState(266)
			p.Match(GraffleParserARG_DELIM)
		}
		{
			p.SetState(267)

			var _x = p.Atom_action()

			localctx.(*ForVarContext).post_act = _x
		}
		p.SetState(269)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraffleParserDO {
			{
				p.SetState(268)
				p.Match(GraffleParserDO)
			}

		}

	case 3:
		localctx = NewForRangeContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(271)
			p.Match(GraffleParserFOR)
		}
		{
			p.SetState(272)
			p.Match(GraffleParserVAR)
		}
		{
			p.SetState(273)
			p.Match(GraffleParserIN)
		}
		p.SetState(275)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraffleParserRANGE {
			{
				p.SetState(274)
				p.Match(GraffleParserRANGE)
			}

		}
		p.SetState(278)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraffleParserFROM {
			{
				p.SetState(277)
				p.Match(GraffleParserFROM)
			}

		}
		{
			p.SetState(280)

			var _x = p.Expr()

			localctx.(*ForRangeContext).from = _x
		}
		{
			p.SetState(281)
			p.Match(GraffleParserTO)
		}
		{
			p.SetState(282)

			var _x = p.Expr()

			localctx.(*ForRangeContext).to = _x
		}
		p.SetState(284)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraffleParserDO {
			{
				p.SetState(283)
				p.Match(GraffleParserDO)
			}

		}

	}

	return localctx
}

// IFrom_to_stmntContext is an interface to support dynamic dispatch.
type IFrom_to_stmntContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetFrom returns the from rule contexts.
	GetFrom() IExprContext

	// GetTo returns the to rule contexts.
	GetTo() IExprContext

	// SetFrom sets the from rule contexts.
	SetFrom(IExprContext)

	// SetTo sets the to rule contexts.
	SetTo(IExprContext)

	// IsFrom_to_stmntContext differentiates from other interfaces.
	IsFrom_to_stmntContext()
}

type From_to_stmntContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	from   IExprContext
	to     IExprContext
}

func NewEmptyFrom_to_stmntContext() *From_to_stmntContext {
	var p = new(From_to_stmntContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_from_to_stmnt
	return p
}

func (*From_to_stmntContext) IsFrom_to_stmntContext() {}

func NewFrom_to_stmntContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *From_to_stmntContext {
	var p = new(From_to_stmntContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_from_to_stmnt

	return p
}

func (s *From_to_stmntContext) GetParser() antlr.Parser { return s.parser }

func (s *From_to_stmntContext) GetFrom() IExprContext { return s.from }

func (s *From_to_stmntContext) GetTo() IExprContext { return s.to }

func (s *From_to_stmntContext) SetFrom(v IExprContext) { s.from = v }

func (s *From_to_stmntContext) SetTo(v IExprContext) { s.to = v }

func (s *From_to_stmntContext) FROM() antlr.TerminalNode {
	return s.GetToken(GraffleParserFROM, 0)
}

func (s *From_to_stmntContext) TO() antlr.TerminalNode {
	return s.GetToken(GraffleParserTO, 0)
}

func (s *From_to_stmntContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *From_to_stmntContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *From_to_stmntContext) DO() antlr.TerminalNode {
	return s.GetToken(GraffleParserDO, 0)
}

func (s *From_to_stmntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *From_to_stmntContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *From_to_stmntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterFrom_to_stmnt(s)
	}
}

func (s *From_to_stmntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitFrom_to_stmnt(s)
	}
}

func (s *From_to_stmntContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitFrom_to_stmnt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) From_to_stmnt() (localctx IFrom_to_stmntContext) {
	localctx = NewFrom_to_stmntContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, GraffleParserRULE_from_to_stmnt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(288)
		p.Match(GraffleParserFROM)
	}
	{
		p.SetState(289)

		var _x = p.Expr()

		localctx.(*From_to_stmntContext).from = _x
	}
	{
		p.SetState(290)
		p.Match(GraffleParserTO)
	}
	{
		p.SetState(291)

		var _x = p.Expr()

		localctx.(*From_to_stmntContext).to = _x
	}
	p.SetState(293)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraffleParserDO {
		{
			p.SetState(292)
			p.Match(GraffleParserDO)
		}

	}

	return localctx
}

// IFunction_declarationContext is an interface to support dynamic dispatch.
type IFunction_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunction_declarationContext differentiates from other interfaces.
	IsFunction_declarationContext()
}

type Function_declarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_declarationContext() *Function_declarationContext {
	var p = new(Function_declarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_function_declaration
	return p
}

func (*Function_declarationContext) IsFunction_declarationContext() {}

func NewFunction_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_declarationContext {
	var p = new(Function_declarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_function_declaration

	return p
}

func (s *Function_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_declarationContext) One_line_function_declaration() IOne_line_function_declarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOne_line_function_declarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOne_line_function_declarationContext)
}

func (s *Function_declarationContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(GraffleParserNEWLINE, 0)
}

func (s *Function_declarationContext) Mult_line_function_declaration() IMult_line_function_declarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMult_line_function_declarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMult_line_function_declarationContext)
}

func (s *Function_declarationContext) One_line_procedure_declaration() IOne_line_procedure_declarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOne_line_procedure_declarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOne_line_procedure_declarationContext)
}

func (s *Function_declarationContext) Mult_line_procedure_declaration() IMult_line_procedure_declarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMult_line_procedure_declarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMult_line_procedure_declarationContext)
}

func (s *Function_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterFunction_declaration(s)
	}
}

func (s *Function_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitFunction_declaration(s)
	}
}

func (s *Function_declarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitFunction_declaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Function_declaration() (localctx IFunction_declarationContext) {
	localctx = NewFunction_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, GraffleParserRULE_function_declaration)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(311)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 44, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(296)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraffleParserNEWLINE {
			{
				p.SetState(295)
				p.Match(GraffleParserNEWLINE)
			}

		}
		{
			p.SetState(298)
			p.One_line_function_declaration()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(300)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraffleParserNEWLINE {
			{
				p.SetState(299)
				p.Match(GraffleParserNEWLINE)
			}

		}
		{
			p.SetState(302)
			p.Mult_line_function_declaration()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(304)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraffleParserNEWLINE {
			{
				p.SetState(303)
				p.Match(GraffleParserNEWLINE)
			}

		}
		{
			p.SetState(306)
			p.One_line_procedure_declaration()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		p.SetState(308)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraffleParserNEWLINE {
			{
				p.SetState(307)
				p.Match(GraffleParserNEWLINE)
			}

		}
		{
			p.SetState(310)
			p.Mult_line_procedure_declaration()
		}

	}

	return localctx
}

// IOne_line_function_declarationContext is an interface to support dynamic dispatch.
type IOne_line_function_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOne_line_function_declarationContext differentiates from other interfaces.
	IsOne_line_function_declarationContext()
}

type One_line_function_declarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOne_line_function_declarationContext() *One_line_function_declarationContext {
	var p = new(One_line_function_declarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_one_line_function_declaration
	return p
}

func (*One_line_function_declarationContext) IsOne_line_function_declarationContext() {}

func NewOne_line_function_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *One_line_function_declarationContext {
	var p = new(One_line_function_declarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_one_line_function_declaration

	return p
}

func (s *One_line_function_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *One_line_function_declarationContext) Function_declaration_head() IFunction_declaration_headContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunction_declaration_headContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunction_declaration_headContext)
}

func (s *One_line_function_declarationContext) Sequence_line() ISequence_lineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISequence_lineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISequence_lineContext)
}

func (s *One_line_function_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *One_line_function_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *One_line_function_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterOne_line_function_declaration(s)
	}
}

func (s *One_line_function_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitOne_line_function_declaration(s)
	}
}

func (s *One_line_function_declarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitOne_line_function_declaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) One_line_function_declaration() (localctx IOne_line_function_declarationContext) {
	localctx = NewOne_line_function_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, GraffleParserRULE_one_line_function_declaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(313)
		p.Function_declaration_head()
	}
	{
		p.SetState(314)
		p.Sequence_line()
	}

	return localctx
}

// IMult_line_function_declarationContext is an interface to support dynamic dispatch.
type IMult_line_function_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMult_line_function_declarationContext differentiates from other interfaces.
	IsMult_line_function_declarationContext()
}

type Mult_line_function_declarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMult_line_function_declarationContext() *Mult_line_function_declarationContext {
	var p = new(Mult_line_function_declarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_mult_line_function_declaration
	return p
}

func (*Mult_line_function_declarationContext) IsMult_line_function_declarationContext() {}

func NewMult_line_function_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Mult_line_function_declarationContext {
	var p = new(Mult_line_function_declarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_mult_line_function_declaration

	return p
}

func (s *Mult_line_function_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Mult_line_function_declarationContext) Function_declaration_head() IFunction_declaration_headContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunction_declaration_headContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunction_declaration_headContext)
}

func (s *Mult_line_function_declarationContext) Sequence() ISequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISequenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISequenceContext)
}

func (s *Mult_line_function_declarationContext) Block_end() IBlock_endContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlock_endContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlock_endContext)
}

func (s *Mult_line_function_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Mult_line_function_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Mult_line_function_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterMult_line_function_declaration(s)
	}
}

func (s *Mult_line_function_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitMult_line_function_declaration(s)
	}
}

func (s *Mult_line_function_declarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitMult_line_function_declaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Mult_line_function_declaration() (localctx IMult_line_function_declarationContext) {
	localctx = NewMult_line_function_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, GraffleParserRULE_mult_line_function_declaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(316)
		p.Function_declaration_head()
	}
	{
		p.SetState(317)
		p.Sequence()
	}
	{
		p.SetState(318)
		p.Block_end()
	}

	return localctx
}

// IFunction_declaration_headContext is an interface to support dynamic dispatch.
type IFunction_declaration_headContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOpd1 returns the opd1 token.
	GetOpd1() antlr.Token

	// GetOpd2 returns the opd2 token.
	GetOpd2() antlr.Token

	// SetOpd1 sets the opd1 token.
	SetOpd1(antlr.Token)

	// SetOpd2 sets the opd2 token.
	SetOpd2(antlr.Token)

	// GetReturn_val returns the return_val rule contexts.
	GetReturn_val() IValueContext

	// SetReturn_val sets the return_val rule contexts.
	SetReturn_val(IValueContext)

	// IsFunction_declaration_headContext differentiates from other interfaces.
	IsFunction_declaration_headContext()
}

type Function_declaration_headContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	opd1       antlr.Token
	opd2       antlr.Token
	return_val IValueContext
}

func NewEmptyFunction_declaration_headContext() *Function_declaration_headContext {
	var p = new(Function_declaration_headContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_function_declaration_head
	return p
}

func (*Function_declaration_headContext) IsFunction_declaration_headContext() {}

func NewFunction_declaration_headContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_declaration_headContext {
	var p = new(Function_declaration_headContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_function_declaration_head

	return p
}

func (s *Function_declaration_headContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_declaration_headContext) GetOpd1() antlr.Token { return s.opd1 }

func (s *Function_declaration_headContext) GetOpd2() antlr.Token { return s.opd2 }

func (s *Function_declaration_headContext) SetOpd1(v antlr.Token) { s.opd1 = v }

func (s *Function_declaration_headContext) SetOpd2(v antlr.Token) { s.opd2 = v }

func (s *Function_declaration_headContext) GetReturn_val() IValueContext { return s.return_val }

func (s *Function_declaration_headContext) SetReturn_val(v IValueContext) { s.return_val = v }

func (s *Function_declaration_headContext) ID() antlr.TerminalNode {
	return s.GetToken(GraffleParserID, 0)
}

func (s *Function_declaration_headContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(GraffleParserL_PAREN, 0)
}

func (s *Function_declaration_headContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(GraffleParserR_PAREN, 0)
}

func (s *Function_declaration_headContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(GraffleParserASSIGN, 0)
}

func (s *Function_declaration_headContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *Function_declaration_headContext) WHERE() antlr.TerminalNode {
	return s.GetToken(GraffleParserWHERE, 0)
}

func (s *Function_declaration_headContext) AllVAR() []antlr.TerminalNode {
	return s.GetTokens(GraffleParserVAR)
}

func (s *Function_declaration_headContext) VAR(i int) antlr.TerminalNode {
	return s.GetToken(GraffleParserVAR, i)
}

func (s *Function_declaration_headContext) AllARG_DELIM() []antlr.TerminalNode {
	return s.GetTokens(GraffleParserARG_DELIM)
}

func (s *Function_declaration_headContext) ARG_DELIM(i int) antlr.TerminalNode {
	return s.GetToken(GraffleParserARG_DELIM, i)
}

func (s *Function_declaration_headContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_declaration_headContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_declaration_headContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterFunction_declaration_head(s)
	}
}

func (s *Function_declaration_headContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitFunction_declaration_head(s)
	}
}

func (s *Function_declaration_headContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitFunction_declaration_head(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Function_declaration_head() (localctx IFunction_declaration_headContext) {
	localctx = NewFunction_declaration_headContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, GraffleParserRULE_function_declaration_head)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(320)
		p.Match(GraffleParserID)
	}
	{
		p.SetState(321)
		p.Match(GraffleParserL_PAREN)
	}
	p.SetState(330)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraffleParserVAR {
		{
			p.SetState(322)

			var _m = p.Match(GraffleParserVAR)

			localctx.(*Function_declaration_headContext).opd1 = _m
		}
		p.SetState(327)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == GraffleParserARG_DELIM {
			{
				p.SetState(323)
				p.Match(GraffleParserARG_DELIM)
			}
			{
				p.SetState(324)

				var _m = p.Match(GraffleParserVAR)

				localctx.(*Function_declaration_headContext).opd2 = _m
			}

			p.SetState(329)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(332)
		p.Match(GraffleParserR_PAREN)
	}
	{
		p.SetState(333)
		p.Match(GraffleParserASSIGN)
	}
	{
		p.SetState(334)

		var _x = p.Value()

		localctx.(*Function_declaration_headContext).return_val = _x
	}
	p.SetState(339)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraffleParserARG_DELIM || _la == GraffleParserWHERE {
		p.SetState(336)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraffleParserARG_DELIM {
			{
				p.SetState(335)
				p.Match(GraffleParserARG_DELIM)
			}

		}
		{
			p.SetState(338)
			p.Match(GraffleParserWHERE)
		}

	}

	return localctx
}

// IOne_line_procedure_declarationContext is an interface to support dynamic dispatch.
type IOne_line_procedure_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOne_line_procedure_declarationContext differentiates from other interfaces.
	IsOne_line_procedure_declarationContext()
}

type One_line_procedure_declarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOne_line_procedure_declarationContext() *One_line_procedure_declarationContext {
	var p = new(One_line_procedure_declarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_one_line_procedure_declaration
	return p
}

func (*One_line_procedure_declarationContext) IsOne_line_procedure_declarationContext() {}

func NewOne_line_procedure_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *One_line_procedure_declarationContext {
	var p = new(One_line_procedure_declarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_one_line_procedure_declaration

	return p
}

func (s *One_line_procedure_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *One_line_procedure_declarationContext) Procedure_declaration_head() IProcedure_declaration_headContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProcedure_declaration_headContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProcedure_declaration_headContext)
}

func (s *One_line_procedure_declarationContext) Sequence_line() ISequence_lineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISequence_lineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISequence_lineContext)
}

func (s *One_line_procedure_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *One_line_procedure_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *One_line_procedure_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterOne_line_procedure_declaration(s)
	}
}

func (s *One_line_procedure_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitOne_line_procedure_declaration(s)
	}
}

func (s *One_line_procedure_declarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitOne_line_procedure_declaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) One_line_procedure_declaration() (localctx IOne_line_procedure_declarationContext) {
	localctx = NewOne_line_procedure_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, GraffleParserRULE_one_line_procedure_declaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(341)
		p.Procedure_declaration_head()
	}
	{
		p.SetState(342)
		p.Sequence_line()
	}

	return localctx
}

// IMult_line_procedure_declarationContext is an interface to support dynamic dispatch.
type IMult_line_procedure_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMult_line_procedure_declarationContext differentiates from other interfaces.
	IsMult_line_procedure_declarationContext()
}

type Mult_line_procedure_declarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMult_line_procedure_declarationContext() *Mult_line_procedure_declarationContext {
	var p = new(Mult_line_procedure_declarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_mult_line_procedure_declaration
	return p
}

func (*Mult_line_procedure_declarationContext) IsMult_line_procedure_declarationContext() {}

func NewMult_line_procedure_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Mult_line_procedure_declarationContext {
	var p = new(Mult_line_procedure_declarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_mult_line_procedure_declaration

	return p
}

func (s *Mult_line_procedure_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Mult_line_procedure_declarationContext) Procedure_declaration_head() IProcedure_declaration_headContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProcedure_declaration_headContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProcedure_declaration_headContext)
}

func (s *Mult_line_procedure_declarationContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(GraffleParserNEWLINE, 0)
}

func (s *Mult_line_procedure_declarationContext) Sequence() ISequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISequenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISequenceContext)
}

func (s *Mult_line_procedure_declarationContext) Block_end() IBlock_endContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlock_endContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlock_endContext)
}

func (s *Mult_line_procedure_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Mult_line_procedure_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Mult_line_procedure_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterMult_line_procedure_declaration(s)
	}
}

func (s *Mult_line_procedure_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitMult_line_procedure_declaration(s)
	}
}

func (s *Mult_line_procedure_declarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitMult_line_procedure_declaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Mult_line_procedure_declaration() (localctx IMult_line_procedure_declarationContext) {
	localctx = NewMult_line_procedure_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, GraffleParserRULE_mult_line_procedure_declaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(344)
		p.Procedure_declaration_head()
	}
	{
		p.SetState(345)
		p.Match(GraffleParserNEWLINE)
	}
	{
		p.SetState(346)
		p.Sequence()
	}
	{
		p.SetState(347)
		p.Block_end()
	}

	return localctx
}

// IProcedure_declaration_headContext is an interface to support dynamic dispatch.
type IProcedure_declaration_headContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOpd1 returns the opd1 token.
	GetOpd1() antlr.Token

	// GetOpd2 returns the opd2 token.
	GetOpd2() antlr.Token

	// SetOpd1 sets the opd1 token.
	SetOpd1(antlr.Token)

	// SetOpd2 sets the opd2 token.
	SetOpd2(antlr.Token)

	// IsProcedure_declaration_headContext differentiates from other interfaces.
	IsProcedure_declaration_headContext()
}

type Procedure_declaration_headContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	opd1   antlr.Token
	opd2   antlr.Token
}

func NewEmptyProcedure_declaration_headContext() *Procedure_declaration_headContext {
	var p = new(Procedure_declaration_headContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_procedure_declaration_head
	return p
}

func (*Procedure_declaration_headContext) IsProcedure_declaration_headContext() {}

func NewProcedure_declaration_headContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Procedure_declaration_headContext {
	var p = new(Procedure_declaration_headContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_procedure_declaration_head

	return p
}

func (s *Procedure_declaration_headContext) GetParser() antlr.Parser { return s.parser }

func (s *Procedure_declaration_headContext) GetOpd1() antlr.Token { return s.opd1 }

func (s *Procedure_declaration_headContext) GetOpd2() antlr.Token { return s.opd2 }

func (s *Procedure_declaration_headContext) SetOpd1(v antlr.Token) { s.opd1 = v }

func (s *Procedure_declaration_headContext) SetOpd2(v antlr.Token) { s.opd2 = v }

func (s *Procedure_declaration_headContext) ID() antlr.TerminalNode {
	return s.GetToken(GraffleParserID, 0)
}

func (s *Procedure_declaration_headContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(GraffleParserL_PAREN, 0)
}

func (s *Procedure_declaration_headContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(GraffleParserR_PAREN, 0)
}

func (s *Procedure_declaration_headContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(GraffleParserASSIGN, 0)
}

func (s *Procedure_declaration_headContext) AllVAR() []antlr.TerminalNode {
	return s.GetTokens(GraffleParserVAR)
}

func (s *Procedure_declaration_headContext) VAR(i int) antlr.TerminalNode {
	return s.GetToken(GraffleParserVAR, i)
}

func (s *Procedure_declaration_headContext) AllARG_DELIM() []antlr.TerminalNode {
	return s.GetTokens(GraffleParserARG_DELIM)
}

func (s *Procedure_declaration_headContext) ARG_DELIM(i int) antlr.TerminalNode {
	return s.GetToken(GraffleParserARG_DELIM, i)
}

func (s *Procedure_declaration_headContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Procedure_declaration_headContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Procedure_declaration_headContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterProcedure_declaration_head(s)
	}
}

func (s *Procedure_declaration_headContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitProcedure_declaration_head(s)
	}
}

func (s *Procedure_declaration_headContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitProcedure_declaration_head(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Procedure_declaration_head() (localctx IProcedure_declaration_headContext) {
	localctx = NewProcedure_declaration_headContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, GraffleParserRULE_procedure_declaration_head)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(349)
		p.Match(GraffleParserID)
	}
	{
		p.SetState(350)
		p.Match(GraffleParserL_PAREN)
	}
	p.SetState(359)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraffleParserVAR {
		{
			p.SetState(351)

			var _m = p.Match(GraffleParserVAR)

			localctx.(*Procedure_declaration_headContext).opd1 = _m
		}
		p.SetState(356)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == GraffleParserARG_DELIM {
			{
				p.SetState(352)
				p.Match(GraffleParserARG_DELIM)
			}
			{
				p.SetState(353)

				var _m = p.Match(GraffleParserVAR)

				localctx.(*Procedure_declaration_headContext).opd2 = _m
			}

			p.SetState(358)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(361)
		p.Match(GraffleParserR_PAREN)
	}
	p.SetState(363)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraffleParserASSIGN {
		{
			p.SetState(362)
			p.Match(GraffleParserASSIGN)
		}

	}

	return localctx
}

// IVar_declarationContext is an interface to support dynamic dispatch.
type IVar_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetVariable returns the variable token.
	GetVariable() antlr.Token

	// GetVal returns the val token.
	GetVal() antlr.Token

	// SetVariable sets the variable token.
	SetVariable(antlr.Token)

	// SetVal sets the val token.
	SetVal(antlr.Token)

	// IsVar_declarationContext differentiates from other interfaces.
	IsVar_declarationContext()
}

type Var_declarationContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	variable antlr.Token
	val      antlr.Token
}

func NewEmptyVar_declarationContext() *Var_declarationContext {
	var p = new(Var_declarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_var_declaration
	return p
}

func (*Var_declarationContext) IsVar_declarationContext() {}

func NewVar_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_declarationContext {
	var p = new(Var_declarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_var_declaration

	return p
}

func (s *Var_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_declarationContext) GetVariable() antlr.Token { return s.variable }

func (s *Var_declarationContext) GetVal() antlr.Token { return s.val }

func (s *Var_declarationContext) SetVariable(v antlr.Token) { s.variable = v }

func (s *Var_declarationContext) SetVal(v antlr.Token) { s.val = v }

func (s *Var_declarationContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(GraffleParserASSIGN, 0)
}

func (s *Var_declarationContext) ID() antlr.TerminalNode {
	return s.GetToken(GraffleParserID, 0)
}

func (s *Var_declarationContext) VAR() antlr.TerminalNode {
	return s.GetToken(GraffleParserVAR, 0)
}

func (s *Var_declarationContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Var_declarationContext) Arc_declaration() IArc_declarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArc_declarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArc_declarationContext)
}

func (s *Var_declarationContext) Vertice_declaration() IVertice_declarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVertice_declarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVertice_declarationContext)
}

func (s *Var_declarationContext) Graph_declaration() IGraph_declarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGraph_declarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGraph_declarationContext)
}

func (s *Var_declarationContext) Labeled_declaration() ILabeled_declarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILabeled_declarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILabeled_declarationContext)
}

func (s *Var_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterVar_declaration(s)
	}
}

func (s *Var_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitVar_declaration(s)
	}
}

func (s *Var_declarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitVar_declaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Var_declaration() (localctx IVar_declarationContext) {
	localctx = NewVar_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, GraffleParserRULE_var_declaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(375)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 52, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(365)

			var _m = p.Match(GraffleParserID)

			localctx.(*Var_declarationContext).variable = _m
		}
		{
			p.SetState(366)
			p.Match(GraffleParserASSIGN)
		}
		{
			p.SetState(367)

			var _m = p.Match(GraffleParserVAR)

			localctx.(*Var_declarationContext).val = _m
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(368)

			var _m = p.Match(GraffleParserID)

			localctx.(*Var_declarationContext).variable = _m
		}
		{
			p.SetState(369)
			p.Match(GraffleParserASSIGN)
		}
		{
			p.SetState(370)
			p.Expr()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(371)
			p.Arc_declaration()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(372)
			p.Vertice_declaration()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(373)
			p.Graph_declaration()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(374)
			p.Labeled_declaration()
		}

	}

	return localctx
}

// IArc_declarationContext is an interface to support dynamic dispatch.
type IArc_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetVariable returns the variable token.
	GetVariable() antlr.Token

	// SetVariable sets the variable token.
	SetVariable(antlr.Token)

	// IsArc_declarationContext differentiates from other interfaces.
	IsArc_declarationContext()
}

type Arc_declarationContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	variable antlr.Token
}

func NewEmptyArc_declarationContext() *Arc_declarationContext {
	var p = new(Arc_declarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_arc_declaration
	return p
}

func (*Arc_declarationContext) IsArc_declarationContext() {}

func NewArc_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Arc_declarationContext {
	var p = new(Arc_declarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_arc_declaration

	return p
}

func (s *Arc_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Arc_declarationContext) GetVariable() antlr.Token { return s.variable }

func (s *Arc_declarationContext) SetVariable(v antlr.Token) { s.variable = v }

func (s *Arc_declarationContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(GraffleParserASSIGN, 0)
}

func (s *Arc_declarationContext) E_N() antlr.TerminalNode {
	return s.GetToken(GraffleParserE_N, 0)
}

func (s *Arc_declarationContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(GraffleParserL_PAREN, 0)
}

func (s *Arc_declarationContext) AllValue() []IValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueContext)(nil)).Elem())
	var tst = make([]IValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueContext)
		}
	}

	return tst
}

func (s *Arc_declarationContext) Value(i int) IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *Arc_declarationContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(GraffleParserR_PAREN, 0)
}

func (s *Arc_declarationContext) ID() antlr.TerminalNode {
	return s.GetToken(GraffleParserID, 0)
}

func (s *Arc_declarationContext) Arc() IArcContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArcContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArcContext)
}

func (s *Arc_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Arc_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Arc_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterArc_declaration(s)
	}
}

func (s *Arc_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitArc_declaration(s)
	}
}

func (s *Arc_declarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitArc_declaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Arc_declaration() (localctx IArc_declarationContext) {
	localctx = NewArc_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, GraffleParserRULE_arc_declaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(390)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 53, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(377)

			var _m = p.Match(GraffleParserID)

			localctx.(*Arc_declarationContext).variable = _m
		}
		{
			p.SetState(378)
			p.Match(GraffleParserASSIGN)
		}
		{
			p.SetState(379)
			p.Match(GraffleParserE_N)
		}
		{
			p.SetState(380)
			p.Match(GraffleParserL_PAREN)
		}
		{
			p.SetState(381)
			p.Value()
		}
		{
			p.SetState(382)
			p.Match(GraffleParserR_PAREN)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(384)

			var _m = p.Match(GraffleParserID)

			localctx.(*Arc_declarationContext).variable = _m
		}
		{
			p.SetState(385)
			p.Match(GraffleParserASSIGN)
		}
		{
			p.SetState(386)
			p.Value()
		}
		{
			p.SetState(387)
			p.Arc()
		}
		{
			p.SetState(388)
			p.Value()
		}

	}

	return localctx
}

// IArcContext is an interface to support dynamic dispatch.
type IArcContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArcContext differentiates from other interfaces.
	IsArcContext()
}

type ArcContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArcContext() *ArcContext {
	var p = new(ArcContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_arc
	return p
}

func (*ArcContext) IsArcContext() {}

func NewArcContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArcContext {
	var p = new(ArcContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_arc

	return p
}

func (s *ArcContext) GetParser() antlr.Parser { return s.parser }

func (s *ArcContext) OR_ARC_LR() antlr.TerminalNode {
	return s.GetToken(GraffleParserOR_ARC_LR, 0)
}

func (s *ArcContext) Or_w_arc_lr() IOr_w_arc_lrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOr_w_arc_lrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOr_w_arc_lrContext)
}

func (s *ArcContext) OR_ARC_RL() antlr.TerminalNode {
	return s.GetToken(GraffleParserOR_ARC_RL, 0)
}

func (s *ArcContext) Or_w_arc_rl() IOr_w_arc_rlContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOr_w_arc_rlContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOr_w_arc_rlContext)
}

func (s *ArcContext) UNOR_ARC() antlr.TerminalNode {
	return s.GetToken(GraffleParserUNOR_ARC, 0)
}

func (s *ArcContext) Unor_w_arc() IUnor_w_arcContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnor_w_arcContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnor_w_arcContext)
}

func (s *ArcContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArcContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArcContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterArc(s)
	}
}

func (s *ArcContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitArc(s)
	}
}

func (s *ArcContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitArc(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Arc() (localctx IArcContext) {
	localctx = NewArcContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, GraffleParserRULE_arc)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(398)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 54, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(392)
			p.Match(GraffleParserOR_ARC_LR)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(393)
			p.Or_w_arc_lr()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(394)
			p.Match(GraffleParserOR_ARC_RL)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(395)
			p.Or_w_arc_rl()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(396)
			p.Match(GraffleParserUNOR_ARC)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(397)
			p.Unor_w_arc()
		}

	}

	return localctx
}

// IOr_w_arc_lrContext is an interface to support dynamic dispatch.
type IOr_w_arc_lrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetWeight returns the weight token.
	GetWeight() antlr.Token

	// SetWeight sets the weight token.
	SetWeight(antlr.Token)

	// IsOr_w_arc_lrContext differentiates from other interfaces.
	IsOr_w_arc_lrContext()
}

type Or_w_arc_lrContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	weight antlr.Token
}

func NewEmptyOr_w_arc_lrContext() *Or_w_arc_lrContext {
	var p = new(Or_w_arc_lrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_or_w_arc_lr
	return p
}

func (*Or_w_arc_lrContext) IsOr_w_arc_lrContext() {}

func NewOr_w_arc_lrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Or_w_arc_lrContext {
	var p = new(Or_w_arc_lrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_or_w_arc_lr

	return p
}

func (s *Or_w_arc_lrContext) GetParser() antlr.Parser { return s.parser }

func (s *Or_w_arc_lrContext) GetWeight() antlr.Token { return s.weight }

func (s *Or_w_arc_lrContext) SetWeight(v antlr.Token) { s.weight = v }

func (s *Or_w_arc_lrContext) SUB() antlr.TerminalNode {
	return s.GetToken(GraffleParserSUB, 0)
}

func (s *Or_w_arc_lrContext) L_BRACKET() antlr.TerminalNode {
	return s.GetToken(GraffleParserL_BRACKET, 0)
}

func (s *Or_w_arc_lrContext) R_BRACKET() antlr.TerminalNode {
	return s.GetToken(GraffleParserR_BRACKET, 0)
}

func (s *Or_w_arc_lrContext) OR_ARC_LR() antlr.TerminalNode {
	return s.GetToken(GraffleParserOR_ARC_LR, 0)
}

func (s *Or_w_arc_lrContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(GraffleParserNUMBER, 0)
}

func (s *Or_w_arc_lrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Or_w_arc_lrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Or_w_arc_lrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterOr_w_arc_lr(s)
	}
}

func (s *Or_w_arc_lrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitOr_w_arc_lr(s)
	}
}

func (s *Or_w_arc_lrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitOr_w_arc_lr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Or_w_arc_lr() (localctx IOr_w_arc_lrContext) {
	localctx = NewOr_w_arc_lrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, GraffleParserRULE_or_w_arc_lr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(400)
		p.Match(GraffleParserSUB)
	}
	{
		p.SetState(401)
		p.Match(GraffleParserL_BRACKET)
	}
	{
		p.SetState(402)

		var _m = p.Match(GraffleParserNUMBER)

		localctx.(*Or_w_arc_lrContext).weight = _m
	}
	{
		p.SetState(403)
		p.Match(GraffleParserR_BRACKET)
	}
	{
		p.SetState(404)
		p.Match(GraffleParserOR_ARC_LR)
	}

	return localctx
}

// IOr_w_arc_rlContext is an interface to support dynamic dispatch.
type IOr_w_arc_rlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetWeight returns the weight token.
	GetWeight() antlr.Token

	// SetWeight sets the weight token.
	SetWeight(antlr.Token)

	// IsOr_w_arc_rlContext differentiates from other interfaces.
	IsOr_w_arc_rlContext()
}

type Or_w_arc_rlContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	weight antlr.Token
}

func NewEmptyOr_w_arc_rlContext() *Or_w_arc_rlContext {
	var p = new(Or_w_arc_rlContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_or_w_arc_rl
	return p
}

func (*Or_w_arc_rlContext) IsOr_w_arc_rlContext() {}

func NewOr_w_arc_rlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Or_w_arc_rlContext {
	var p = new(Or_w_arc_rlContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_or_w_arc_rl

	return p
}

func (s *Or_w_arc_rlContext) GetParser() antlr.Parser { return s.parser }

func (s *Or_w_arc_rlContext) GetWeight() antlr.Token { return s.weight }

func (s *Or_w_arc_rlContext) SetWeight(v antlr.Token) { s.weight = v }

func (s *Or_w_arc_rlContext) OR_ARC_RL() antlr.TerminalNode {
	return s.GetToken(GraffleParserOR_ARC_RL, 0)
}

func (s *Or_w_arc_rlContext) L_BRACKET() antlr.TerminalNode {
	return s.GetToken(GraffleParserL_BRACKET, 0)
}

func (s *Or_w_arc_rlContext) R_BRACKET() antlr.TerminalNode {
	return s.GetToken(GraffleParserR_BRACKET, 0)
}

func (s *Or_w_arc_rlContext) SUB() antlr.TerminalNode {
	return s.GetToken(GraffleParserSUB, 0)
}

func (s *Or_w_arc_rlContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(GraffleParserNUMBER, 0)
}

func (s *Or_w_arc_rlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Or_w_arc_rlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Or_w_arc_rlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterOr_w_arc_rl(s)
	}
}

func (s *Or_w_arc_rlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitOr_w_arc_rl(s)
	}
}

func (s *Or_w_arc_rlContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitOr_w_arc_rl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Or_w_arc_rl() (localctx IOr_w_arc_rlContext) {
	localctx = NewOr_w_arc_rlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, GraffleParserRULE_or_w_arc_rl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(406)
		p.Match(GraffleParserOR_ARC_RL)
	}
	{
		p.SetState(407)
		p.Match(GraffleParserL_BRACKET)
	}
	{
		p.SetState(408)

		var _m = p.Match(GraffleParserNUMBER)

		localctx.(*Or_w_arc_rlContext).weight = _m
	}
	{
		p.SetState(409)
		p.Match(GraffleParserR_BRACKET)
	}
	{
		p.SetState(410)
		p.Match(GraffleParserSUB)
	}

	return localctx
}

// IUnor_w_arcContext is an interface to support dynamic dispatch.
type IUnor_w_arcContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetWeight returns the weight token.
	GetWeight() antlr.Token

	// SetWeight sets the weight token.
	SetWeight(antlr.Token)

	// IsUnor_w_arcContext differentiates from other interfaces.
	IsUnor_w_arcContext()
}

type Unor_w_arcContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	weight antlr.Token
}

func NewEmptyUnor_w_arcContext() *Unor_w_arcContext {
	var p = new(Unor_w_arcContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_unor_w_arc
	return p
}

func (*Unor_w_arcContext) IsUnor_w_arcContext() {}

func NewUnor_w_arcContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Unor_w_arcContext {
	var p = new(Unor_w_arcContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_unor_w_arc

	return p
}

func (s *Unor_w_arcContext) GetParser() antlr.Parser { return s.parser }

func (s *Unor_w_arcContext) GetWeight() antlr.Token { return s.weight }

func (s *Unor_w_arcContext) SetWeight(v antlr.Token) { s.weight = v }

func (s *Unor_w_arcContext) AllSUB() []antlr.TerminalNode {
	return s.GetTokens(GraffleParserSUB)
}

func (s *Unor_w_arcContext) SUB(i int) antlr.TerminalNode {
	return s.GetToken(GraffleParserSUB, i)
}

func (s *Unor_w_arcContext) L_BRACKET() antlr.TerminalNode {
	return s.GetToken(GraffleParserL_BRACKET, 0)
}

func (s *Unor_w_arcContext) R_BRACKET() antlr.TerminalNode {
	return s.GetToken(GraffleParserR_BRACKET, 0)
}

func (s *Unor_w_arcContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(GraffleParserNUMBER, 0)
}

func (s *Unor_w_arcContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Unor_w_arcContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Unor_w_arcContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterUnor_w_arc(s)
	}
}

func (s *Unor_w_arcContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitUnor_w_arc(s)
	}
}

func (s *Unor_w_arcContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitUnor_w_arc(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Unor_w_arc() (localctx IUnor_w_arcContext) {
	localctx = NewUnor_w_arcContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, GraffleParserRULE_unor_w_arc)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(412)
		p.Match(GraffleParserSUB)
	}
	{
		p.SetState(413)
		p.Match(GraffleParserL_BRACKET)
	}
	{
		p.SetState(414)

		var _m = p.Match(GraffleParserNUMBER)

		localctx.(*Unor_w_arcContext).weight = _m
	}
	{
		p.SetState(415)
		p.Match(GraffleParserR_BRACKET)
	}
	{
		p.SetState(416)
		p.Match(GraffleParserSUB)
	}

	return localctx
}

// IVertice_declarationContext is an interface to support dynamic dispatch.
type IVertice_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVertice_declarationContext differentiates from other interfaces.
	IsVertice_declarationContext()
}

type Vertice_declarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVertice_declarationContext() *Vertice_declarationContext {
	var p = new(Vertice_declarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_vertice_declaration
	return p
}

func (*Vertice_declarationContext) IsVertice_declarationContext() {}

func NewVertice_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Vertice_declarationContext {
	var p = new(Vertice_declarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_vertice_declaration

	return p
}

func (s *Vertice_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Vertice_declarationContext) ID() antlr.TerminalNode {
	return s.GetToken(GraffleParserID, 0)
}

func (s *Vertice_declarationContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(GraffleParserASSIGN, 0)
}

func (s *Vertice_declarationContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(GraffleParserL_PAREN, 0)
}

func (s *Vertice_declarationContext) V_N() antlr.TerminalNode {
	return s.GetToken(GraffleParserV_N, 0)
}

func (s *Vertice_declarationContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(GraffleParserR_PAREN, 0)
}

func (s *Vertice_declarationContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *Vertice_declarationContext) Arithm_assign_operator() IArithm_assign_operatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArithm_assign_operatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArithm_assign_operatorContext)
}

func (s *Vertice_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Vertice_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Vertice_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterVertice_declaration(s)
	}
}

func (s *Vertice_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitVertice_declaration(s)
	}
}

func (s *Vertice_declarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitVertice_declaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Vertice_declaration() (localctx IVertice_declarationContext) {
	localctx = NewVertice_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, GraffleParserRULE_vertice_declaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(431)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 55, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(418)
			p.Match(GraffleParserID)
		}
		{
			p.SetState(419)
			p.Match(GraffleParserASSIGN)
		}
		{
			p.SetState(420)
			p.Match(GraffleParserL_PAREN)
		}
		{
			p.SetState(421)
			p.Match(GraffleParserV_N)
		}
		{
			p.SetState(422)
			p.Match(GraffleParserR_PAREN)
		}
		{
			p.SetState(423)
			p.Value()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(424)
			p.Match(GraffleParserID)
		}
		{
			p.SetState(425)
			p.Match(GraffleParserASSIGN)
		}
		{
			p.SetState(426)
			p.Value()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(427)
			p.Match(GraffleParserID)
		}
		{
			p.SetState(428)
			p.Arithm_assign_operator()
		}
		{
			p.SetState(429)
			p.Value()
		}

	}

	return localctx
}

// IGraph_declarationContext is an interface to support dynamic dispatch.
type IGraph_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGraph_declarationContext differentiates from other interfaces.
	IsGraph_declarationContext()
}

type Graph_declarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGraph_declarationContext() *Graph_declarationContext {
	var p = new(Graph_declarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_graph_declaration
	return p
}

func (*Graph_declarationContext) IsGraph_declarationContext() {}

func NewGraph_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Graph_declarationContext {
	var p = new(Graph_declarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_graph_declaration

	return p
}

func (s *Graph_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Graph_declarationContext) ID() antlr.TerminalNode {
	return s.GetToken(GraffleParserID, 0)
}

func (s *Graph_declarationContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(GraffleParserASSIGN, 0)
}

func (s *Graph_declarationContext) G_N() antlr.TerminalNode {
	return s.GetToken(GraffleParserG_N, 0)
}

func (s *Graph_declarationContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(GraffleParserL_PAREN, 0)
}

func (s *Graph_declarationContext) AllValue() []IValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueContext)(nil)).Elem())
	var tst = make([]IValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueContext)
		}
	}

	return tst
}

func (s *Graph_declarationContext) Value(i int) IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *Graph_declarationContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(GraffleParserR_PAREN, 0)
}

func (s *Graph_declarationContext) AllARG_DELIM() []antlr.TerminalNode {
	return s.GetTokens(GraffleParserARG_DELIM)
}

func (s *Graph_declarationContext) ARG_DELIM(i int) antlr.TerminalNode {
	return s.GetToken(GraffleParserARG_DELIM, i)
}

func (s *Graph_declarationContext) Arithm_assign_operator() IArithm_assign_operatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArithm_assign_operatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArithm_assign_operatorContext)
}

func (s *Graph_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Graph_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Graph_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterGraph_declaration(s)
	}
}

func (s *Graph_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitGraph_declaration(s)
	}
}

func (s *Graph_declarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitGraph_declaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Graph_declaration() (localctx IGraph_declarationContext) {
	localctx = NewGraph_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, GraffleParserRULE_graph_declaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.SetState(454)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 57, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(433)
			p.Match(GraffleParserID)
		}
		{
			p.SetState(434)
			p.Match(GraffleParserASSIGN)
		}
		{
			p.SetState(435)
			p.Match(GraffleParserG_N)
		}
		{
			p.SetState(436)
			p.Match(GraffleParserL_PAREN)
		}
		{
			p.SetState(437)
			p.Value()
		}
		{
			p.SetState(438)
			p.Match(GraffleParserR_PAREN)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(440)
			p.Match(GraffleParserID)
		}
		{
			p.SetState(441)
			p.Match(GraffleParserASSIGN)
		}
		{
			p.SetState(442)
			p.Value()
		}
		p.SetState(447)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 56, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(443)
					p.Match(GraffleParserARG_DELIM)
				}
				{
					p.SetState(444)
					p.Value()
				}

			}
			p.SetState(449)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 56, p.GetParserRuleContext())
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(450)
			p.Match(GraffleParserID)
		}
		{
			p.SetState(451)
			p.Arithm_assign_operator()
		}
		{
			p.SetState(452)
			p.Value()
		}

	}

	return localctx
}

// ILabeled_declarationContext is an interface to support dynamic dispatch.
type ILabeled_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLabeled_declarationContext differentiates from other interfaces.
	IsLabeled_declarationContext()
}

type Labeled_declarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLabeled_declarationContext() *Labeled_declarationContext {
	var p = new(Labeled_declarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_labeled_declaration
	return p
}

func (*Labeled_declarationContext) IsLabeled_declarationContext() {}

func NewLabeled_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Labeled_declarationContext {
	var p = new(Labeled_declarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_labeled_declaration

	return p
}

func (s *Labeled_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Labeled_declarationContext) Vertice_declaration() IVertice_declarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVertice_declarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVertice_declarationContext)
}

func (s *Labeled_declarationContext) Label() ILabelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILabelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILabelContext)
}

func (s *Labeled_declarationContext) Arc_declaration() IArc_declarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArc_declarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArc_declarationContext)
}

func (s *Labeled_declarationContext) Graph_declaration() IGraph_declarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGraph_declarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGraph_declarationContext)
}

func (s *Labeled_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Labeled_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Labeled_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterLabeled_declaration(s)
	}
}

func (s *Labeled_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitLabeled_declaration(s)
	}
}

func (s *Labeled_declarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitLabeled_declaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Labeled_declaration() (localctx ILabeled_declarationContext) {
	localctx = NewLabeled_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, GraffleParserRULE_labeled_declaration)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(468)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 61, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(456)
			p.Vertice_declaration()
		}
		p.SetState(458)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraffleParserLABEL || _la == GraffleParserML_LABEL {
			{
				p.SetState(457)
				p.Label()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(460)
			p.Arc_declaration()
		}
		p.SetState(462)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraffleParserLABEL || _la == GraffleParserML_LABEL {
			{
				p.SetState(461)
				p.Label()
			}

		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(464)
			p.Graph_declaration()
		}
		p.SetState(466)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraffleParserLABEL || _la == GraffleParserML_LABEL {
			{
				p.SetState(465)
				p.Label()
			}

		}

	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) VAR() antlr.TerminalNode {
	return s.GetToken(GraffleParserVAR, 0)
}

func (s *ExprContext) Integral_expr() IIntegral_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegral_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntegral_exprContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (s *ExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Expr() (localctx IExprContext) {
	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, GraffleParserRULE_expr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(472)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GraffleParserVAR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(470)
			p.Match(GraffleParserVAR)
		}

	case GraffleParserNUMBER, GraffleParserSTRING, GraffleParserBOOL, GraffleParserL_PAREN, GraffleParserID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(471)
			p.Integral_expr()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIntegral_exprContext is an interface to support dynamic dispatch.
type IIntegral_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntegral_exprContext differentiates from other interfaces.
	IsIntegral_exprContext()
}

type Integral_exprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntegral_exprContext() *Integral_exprContext {
	var p = new(Integral_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_integral_expr
	return p
}

func (*Integral_exprContext) IsIntegral_exprContext() {}

func NewIntegral_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Integral_exprContext {
	var p = new(Integral_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_integral_expr

	return p
}

func (s *Integral_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Integral_exprContext) Function_call() IFunction_callContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunction_callContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunction_callContext)
}

func (s *Integral_exprContext) Builtin() IBuiltinContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBuiltinContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBuiltinContext)
}

func (s *Integral_exprContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(GraffleParserL_PAREN, 0)
}

func (s *Integral_exprContext) Logical_expr() ILogical_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILogical_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILogical_exprContext)
}

func (s *Integral_exprContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(GraffleParserR_PAREN, 0)
}

func (s *Integral_exprContext) Arithm_expr() IArithm_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArithm_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArithm_exprContext)
}

func (s *Integral_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Integral_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Integral_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterIntegral_expr(s)
	}
}

func (s *Integral_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitIntegral_expr(s)
	}
}

func (s *Integral_exprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitIntegral_expr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Integral_expr() (localctx IIntegral_exprContext) {
	localctx = NewIntegral_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, GraffleParserRULE_integral_expr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(484)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 63, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(474)
			p.Function_call()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(475)
			p.Builtin()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(476)
			p.Match(GraffleParserL_PAREN)
		}
		{
			p.SetState(477)
			p.Logical_expr()
		}
		{
			p.SetState(478)
			p.Match(GraffleParserR_PAREN)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(480)
			p.Match(GraffleParserL_PAREN)
		}
		{
			p.SetState(481)
			p.Arithm_expr()
		}
		{
			p.SetState(482)
			p.Match(GraffleParserR_PAREN)
		}

	}

	return localctx
}

// ILogical_exprContext is an interface to support dynamic dispatch.
type ILogical_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLeft returns the left rule contexts.
	GetLeft() IExprContext

	// GetBin_op returns the bin_op rule contexts.
	GetBin_op() IBin_log_operatorContext

	// GetRight returns the right rule contexts.
	GetRight() ILog_expr_operandContext

	// GetUn_op returns the un_op rule contexts.
	GetUn_op() IUnar_log_operatorContext

	// SetLeft sets the left rule contexts.
	SetLeft(IExprContext)

	// SetBin_op sets the bin_op rule contexts.
	SetBin_op(IBin_log_operatorContext)

	// SetRight sets the right rule contexts.
	SetRight(ILog_expr_operandContext)

	// SetUn_op sets the un_op rule contexts.
	SetUn_op(IUnar_log_operatorContext)

	// IsLogical_exprContext differentiates from other interfaces.
	IsLogical_exprContext()
}

type Logical_exprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	left   IExprContext
	bin_op IBin_log_operatorContext
	right  ILog_expr_operandContext
	un_op  IUnar_log_operatorContext
}

func NewEmptyLogical_exprContext() *Logical_exprContext {
	var p = new(Logical_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_logical_expr
	return p
}

func (*Logical_exprContext) IsLogical_exprContext() {}

func NewLogical_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Logical_exprContext {
	var p = new(Logical_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_logical_expr

	return p
}

func (s *Logical_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Logical_exprContext) GetLeft() IExprContext { return s.left }

func (s *Logical_exprContext) GetBin_op() IBin_log_operatorContext { return s.bin_op }

func (s *Logical_exprContext) GetRight() ILog_expr_operandContext { return s.right }

func (s *Logical_exprContext) GetUn_op() IUnar_log_operatorContext { return s.un_op }

func (s *Logical_exprContext) SetLeft(v IExprContext) { s.left = v }

func (s *Logical_exprContext) SetBin_op(v IBin_log_operatorContext) { s.bin_op = v }

func (s *Logical_exprContext) SetRight(v ILog_expr_operandContext) { s.right = v }

func (s *Logical_exprContext) SetUn_op(v IUnar_log_operatorContext) { s.un_op = v }

func (s *Logical_exprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Logical_exprContext) Bin_log_operator() IBin_log_operatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBin_log_operatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBin_log_operatorContext)
}

func (s *Logical_exprContext) Log_expr_operand() ILog_expr_operandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILog_expr_operandContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILog_expr_operandContext)
}

func (s *Logical_exprContext) Unar_log_operator() IUnar_log_operatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnar_log_operatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnar_log_operatorContext)
}

func (s *Logical_exprContext) Arithm_expr() IArithm_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArithm_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArithm_exprContext)
}

func (s *Logical_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Logical_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Logical_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterLogical_expr(s)
	}
}

func (s *Logical_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitLogical_expr(s)
	}
}

func (s *Logical_exprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitLogical_expr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Logical_expr() (localctx ILogical_exprContext) {
	localctx = NewLogical_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, GraffleParserRULE_logical_expr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(495)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 64, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(486)

			var _x = p.Expr()

			localctx.(*Logical_exprContext).left = _x
		}
		{
			p.SetState(487)

			var _x = p.Bin_log_operator()

			localctx.(*Logical_exprContext).bin_op = _x
		}
		{
			p.SetState(488)

			var _x = p.Log_expr_operand()

			localctx.(*Logical_exprContext).right = _x
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(490)

			var _x = p.Unar_log_operator()

			localctx.(*Logical_exprContext).un_op = _x
		}
		{
			p.SetState(491)
			p.Log_expr_operand()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(493)
			p.Arithm_expr()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(494)
			p.Expr()
		}

	}

	return localctx
}

// ILog_expr_operandContext is an interface to support dynamic dispatch.
type ILog_expr_operandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLog_expr_operandContext differentiates from other interfaces.
	IsLog_expr_operandContext()
}

type Log_expr_operandContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLog_expr_operandContext() *Log_expr_operandContext {
	var p = new(Log_expr_operandContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_log_expr_operand
	return p
}

func (*Log_expr_operandContext) IsLog_expr_operandContext() {}

func NewLog_expr_operandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Log_expr_operandContext {
	var p = new(Log_expr_operandContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_log_expr_operand

	return p
}

func (s *Log_expr_operandContext) GetParser() antlr.Parser { return s.parser }

func (s *Log_expr_operandContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Log_expr_operandContext) Logical_expr() ILogical_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILogical_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILogical_exprContext)
}

func (s *Log_expr_operandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Log_expr_operandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Log_expr_operandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterLog_expr_operand(s)
	}
}

func (s *Log_expr_operandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitLog_expr_operand(s)
	}
}

func (s *Log_expr_operandContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitLog_expr_operand(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Log_expr_operand() (localctx ILog_expr_operandContext) {
	localctx = NewLog_expr_operandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, GraffleParserRULE_log_expr_operand)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(499)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 65, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(497)
			p.Expr()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(498)
			p.Logical_expr()
		}

	}

	return localctx
}

// IBin_log_operatorContext is an interface to support dynamic dispatch.
type IBin_log_operatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBin_log_operatorContext differentiates from other interfaces.
	IsBin_log_operatorContext()
}

type Bin_log_operatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBin_log_operatorContext() *Bin_log_operatorContext {
	var p = new(Bin_log_operatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_bin_log_operator
	return p
}

func (*Bin_log_operatorContext) IsBin_log_operatorContext() {}

func NewBin_log_operatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Bin_log_operatorContext {
	var p = new(Bin_log_operatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_bin_log_operator

	return p
}

func (s *Bin_log_operatorContext) GetParser() antlr.Parser { return s.parser }

func (s *Bin_log_operatorContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(GraffleParserEQUALS, 0)
}

func (s *Bin_log_operatorContext) NEQ() antlr.TerminalNode {
	return s.GetToken(GraffleParserNEQ, 0)
}

func (s *Bin_log_operatorContext) LESS_THAN() antlr.TerminalNode {
	return s.GetToken(GraffleParserLESS_THAN, 0)
}

func (s *Bin_log_operatorContext) GR_THAN() antlr.TerminalNode {
	return s.GetToken(GraffleParserGR_THAN, 0)
}

func (s *Bin_log_operatorContext) LESS_THAN_E() antlr.TerminalNode {
	return s.GetToken(GraffleParserLESS_THAN_E, 0)
}

func (s *Bin_log_operatorContext) GR_THAN_E() antlr.TerminalNode {
	return s.GetToken(GraffleParserGR_THAN_E, 0)
}

func (s *Bin_log_operatorContext) AND() antlr.TerminalNode {
	return s.GetToken(GraffleParserAND, 0)
}

func (s *Bin_log_operatorContext) NAND() antlr.TerminalNode {
	return s.GetToken(GraffleParserNAND, 0)
}

func (s *Bin_log_operatorContext) OR() antlr.TerminalNode {
	return s.GetToken(GraffleParserOR, 0)
}

func (s *Bin_log_operatorContext) NOR() antlr.TerminalNode {
	return s.GetToken(GraffleParserNOR, 0)
}

func (s *Bin_log_operatorContext) XOR() antlr.TerminalNode {
	return s.GetToken(GraffleParserXOR, 0)
}

func (s *Bin_log_operatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Bin_log_operatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Bin_log_operatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterBin_log_operator(s)
	}
}

func (s *Bin_log_operatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitBin_log_operator(s)
	}
}

func (s *Bin_log_operatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitBin_log_operator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Bin_log_operator() (localctx IBin_log_operatorContext) {
	localctx = NewBin_log_operatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, GraffleParserRULE_bin_log_operator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(501)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-29)&-(0x1f+1)) == 0 && ((1<<uint((_la-29)))&((1<<(GraffleParserNEQ-29))|(1<<(GraffleParserEQUALS-29))|(1<<(GraffleParserLESS_THAN-29))|(1<<(GraffleParserGR_THAN-29))|(1<<(GraffleParserLESS_THAN_E-29))|(1<<(GraffleParserGR_THAN_E-29))|(1<<(GraffleParserAND-29))|(1<<(GraffleParserOR-29))|(1<<(GraffleParserXOR-29))|(1<<(GraffleParserNOR-29))|(1<<(GraffleParserNAND-29)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IUnar_log_operatorContext is an interface to support dynamic dispatch.
type IUnar_log_operatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnar_log_operatorContext differentiates from other interfaces.
	IsUnar_log_operatorContext()
}

type Unar_log_operatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnar_log_operatorContext() *Unar_log_operatorContext {
	var p = new(Unar_log_operatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_unar_log_operator
	return p
}

func (*Unar_log_operatorContext) IsUnar_log_operatorContext() {}

func NewUnar_log_operatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Unar_log_operatorContext {
	var p = new(Unar_log_operatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_unar_log_operator

	return p
}

func (s *Unar_log_operatorContext) GetParser() antlr.Parser { return s.parser }

func (s *Unar_log_operatorContext) NOT() antlr.TerminalNode {
	return s.GetToken(GraffleParserNOT, 0)
}

func (s *Unar_log_operatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Unar_log_operatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Unar_log_operatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterUnar_log_operator(s)
	}
}

func (s *Unar_log_operatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitUnar_log_operator(s)
	}
}

func (s *Unar_log_operatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitUnar_log_operator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Unar_log_operator() (localctx IUnar_log_operatorContext) {
	localctx = NewUnar_log_operatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, GraffleParserRULE_unar_log_operator)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(503)
		p.Match(GraffleParserNOT)
	}

	return localctx
}

// IArithm_exprContext is an interface to support dynamic dispatch.
type IArithm_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLeft returns the left rule contexts.
	GetLeft() IExprContext

	// GetOp returns the op rule contexts.
	GetOp() IBin_arithm_operatorContext

	// GetRight returns the right rule contexts.
	GetRight() IArithm_expr_operandContext

	// SetLeft sets the left rule contexts.
	SetLeft(IExprContext)

	// SetOp sets the op rule contexts.
	SetOp(IBin_arithm_operatorContext)

	// SetRight sets the right rule contexts.
	SetRight(IArithm_expr_operandContext)

	// IsArithm_exprContext differentiates from other interfaces.
	IsArithm_exprContext()
}

type Arithm_exprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	left   IExprContext
	op     IBin_arithm_operatorContext
	right  IArithm_expr_operandContext
}

func NewEmptyArithm_exprContext() *Arithm_exprContext {
	var p = new(Arithm_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_arithm_expr
	return p
}

func (*Arithm_exprContext) IsArithm_exprContext() {}

func NewArithm_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Arithm_exprContext {
	var p = new(Arithm_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_arithm_expr

	return p
}

func (s *Arithm_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Arithm_exprContext) GetLeft() IExprContext { return s.left }

func (s *Arithm_exprContext) GetOp() IBin_arithm_operatorContext { return s.op }

func (s *Arithm_exprContext) GetRight() IArithm_expr_operandContext { return s.right }

func (s *Arithm_exprContext) SetLeft(v IExprContext) { s.left = v }

func (s *Arithm_exprContext) SetOp(v IBin_arithm_operatorContext) { s.op = v }

func (s *Arithm_exprContext) SetRight(v IArithm_expr_operandContext) { s.right = v }

func (s *Arithm_exprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Arithm_exprContext) Bin_arithm_operator() IBin_arithm_operatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBin_arithm_operatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBin_arithm_operatorContext)
}

func (s *Arithm_exprContext) Arithm_expr_operand() IArithm_expr_operandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArithm_expr_operandContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArithm_expr_operandContext)
}

func (s *Arithm_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Arithm_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Arithm_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterArithm_expr(s)
	}
}

func (s *Arithm_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitArithm_expr(s)
	}
}

func (s *Arithm_exprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitArithm_expr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Arithm_expr() (localctx IArithm_exprContext) {
	localctx = NewArithm_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, GraffleParserRULE_arithm_expr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(505)

		var _x = p.Expr()

		localctx.(*Arithm_exprContext).left = _x
	}
	{
		p.SetState(506)

		var _x = p.Bin_arithm_operator()

		localctx.(*Arithm_exprContext).op = _x
	}
	{
		p.SetState(507)

		var _x = p.Arithm_expr_operand()

		localctx.(*Arithm_exprContext).right = _x
	}

	return localctx
}

// IArithm_expr_operandContext is an interface to support dynamic dispatch.
type IArithm_expr_operandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArithm_expr_operandContext differentiates from other interfaces.
	IsArithm_expr_operandContext()
}

type Arithm_expr_operandContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArithm_expr_operandContext() *Arithm_expr_operandContext {
	var p = new(Arithm_expr_operandContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_arithm_expr_operand
	return p
}

func (*Arithm_expr_operandContext) IsArithm_expr_operandContext() {}

func NewArithm_expr_operandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Arithm_expr_operandContext {
	var p = new(Arithm_expr_operandContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_arithm_expr_operand

	return p
}

func (s *Arithm_expr_operandContext) GetParser() antlr.Parser { return s.parser }

func (s *Arithm_expr_operandContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Arithm_expr_operandContext) Arithm_expr() IArithm_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArithm_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArithm_exprContext)
}

func (s *Arithm_expr_operandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Arithm_expr_operandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Arithm_expr_operandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterArithm_expr_operand(s)
	}
}

func (s *Arithm_expr_operandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitArithm_expr_operand(s)
	}
}

func (s *Arithm_expr_operandContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitArithm_expr_operand(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Arithm_expr_operand() (localctx IArithm_expr_operandContext) {
	localctx = NewArithm_expr_operandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, GraffleParserRULE_arithm_expr_operand)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(511)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 66, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(509)
			p.Expr()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(510)
			p.Arithm_expr()
		}

	}

	return localctx
}

// IBin_arithm_operatorContext is an interface to support dynamic dispatch.
type IBin_arithm_operatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBin_arithm_operatorContext differentiates from other interfaces.
	IsBin_arithm_operatorContext()
}

type Bin_arithm_operatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBin_arithm_operatorContext() *Bin_arithm_operatorContext {
	var p = new(Bin_arithm_operatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_bin_arithm_operator
	return p
}

func (*Bin_arithm_operatorContext) IsBin_arithm_operatorContext() {}

func NewBin_arithm_operatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Bin_arithm_operatorContext {
	var p = new(Bin_arithm_operatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_bin_arithm_operator

	return p
}

func (s *Bin_arithm_operatorContext) GetParser() antlr.Parser { return s.parser }

func (s *Bin_arithm_operatorContext) MULT() antlr.TerminalNode {
	return s.GetToken(GraffleParserMULT, 0)
}

func (s *Bin_arithm_operatorContext) DIV() antlr.TerminalNode {
	return s.GetToken(GraffleParserDIV, 0)
}

func (s *Bin_arithm_operatorContext) ADD() antlr.TerminalNode {
	return s.GetToken(GraffleParserADD, 0)
}

func (s *Bin_arithm_operatorContext) SUB() antlr.TerminalNode {
	return s.GetToken(GraffleParserSUB, 0)
}

func (s *Bin_arithm_operatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Bin_arithm_operatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Bin_arithm_operatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterBin_arithm_operator(s)
	}
}

func (s *Bin_arithm_operatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitBin_arithm_operator(s)
	}
}

func (s *Bin_arithm_operatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitBin_arithm_operator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Bin_arithm_operator() (localctx IBin_arithm_operatorContext) {
	localctx = NewBin_arithm_operatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, GraffleParserRULE_bin_arithm_operator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(513)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GraffleParserADD)|(1<<GraffleParserSUB)|(1<<GraffleParserMULT)|(1<<GraffleParserDIV))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IArithm_assign_operatorContext is an interface to support dynamic dispatch.
type IArithm_assign_operatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArithm_assign_operatorContext differentiates from other interfaces.
	IsArithm_assign_operatorContext()
}

type Arithm_assign_operatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArithm_assign_operatorContext() *Arithm_assign_operatorContext {
	var p = new(Arithm_assign_operatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_arithm_assign_operator
	return p
}

func (*Arithm_assign_operatorContext) IsArithm_assign_operatorContext() {}

func NewArithm_assign_operatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Arithm_assign_operatorContext {
	var p = new(Arithm_assign_operatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_arithm_assign_operator

	return p
}

func (s *Arithm_assign_operatorContext) GetParser() antlr.Parser { return s.parser }

func (s *Arithm_assign_operatorContext) ADD_ASSIGN() antlr.TerminalNode {
	return s.GetToken(GraffleParserADD_ASSIGN, 0)
}

func (s *Arithm_assign_operatorContext) SUB_ASSIGN() antlr.TerminalNode {
	return s.GetToken(GraffleParserSUB_ASSIGN, 0)
}

func (s *Arithm_assign_operatorContext) MULT_ASSIGN() antlr.TerminalNode {
	return s.GetToken(GraffleParserMULT_ASSIGN, 0)
}

func (s *Arithm_assign_operatorContext) DIV_ASSIGN() antlr.TerminalNode {
	return s.GetToken(GraffleParserDIV_ASSIGN, 0)
}

func (s *Arithm_assign_operatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Arithm_assign_operatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Arithm_assign_operatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterArithm_assign_operator(s)
	}
}

func (s *Arithm_assign_operatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitArithm_assign_operator(s)
	}
}

func (s *Arithm_assign_operatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitArithm_assign_operator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Arithm_assign_operator() (localctx IArithm_assign_operatorContext) {
	localctx = NewArithm_assign_operatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, GraffleParserRULE_arithm_assign_operator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(515)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GraffleParserADD_ASSIGN)|(1<<GraffleParserSUB_ASSIGN)|(1<<GraffleParserMULT_ASSIGN)|(1<<GraffleParserDIV_ASSIGN))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IBuiltin_function_callContext is an interface to support dynamic dispatch.
type IBuiltin_function_callContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBuiltin_function_callContext differentiates from other interfaces.
	IsBuiltin_function_callContext()
}

type Builtin_function_callContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBuiltin_function_callContext() *Builtin_function_callContext {
	var p = new(Builtin_function_callContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_builtin_function_call
	return p
}

func (*Builtin_function_callContext) IsBuiltin_function_callContext() {}

func NewBuiltin_function_callContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Builtin_function_callContext {
	var p = new(Builtin_function_callContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_builtin_function_call

	return p
}

func (s *Builtin_function_callContext) GetParser() antlr.Parser { return s.parser }

func (s *Builtin_function_callContext) Built_func_print() IBuilt_func_printContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBuilt_func_printContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBuilt_func_printContext)
}

func (s *Builtin_function_callContext) Built_func_input() IBuilt_func_inputContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBuilt_func_inputContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBuilt_func_inputContext)
}

func (s *Builtin_function_callContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Builtin_function_callContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Builtin_function_callContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterBuiltin_function_call(s)
	}
}

func (s *Builtin_function_callContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitBuiltin_function_call(s)
	}
}

func (s *Builtin_function_callContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitBuiltin_function_call(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Builtin_function_call() (localctx IBuiltin_function_callContext) {
	localctx = NewBuiltin_function_callContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, GraffleParserRULE_builtin_function_call)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(519)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GraffleParserPRINTER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(517)
			p.Built_func_print()
		}

	case GraffleParserKEY_INPUT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(518)
			p.Built_func_input()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBuilt_func_printContext is an interface to support dynamic dispatch.
type IBuilt_func_printContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBuilt_func_printContext differentiates from other interfaces.
	IsBuilt_func_printContext()
}

type Built_func_printContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBuilt_func_printContext() *Built_func_printContext {
	var p = new(Built_func_printContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_built_func_print
	return p
}

func (*Built_func_printContext) IsBuilt_func_printContext() {}

func NewBuilt_func_printContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Built_func_printContext {
	var p = new(Built_func_printContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_built_func_print

	return p
}

func (s *Built_func_printContext) GetParser() antlr.Parser { return s.parser }

func (s *Built_func_printContext) PRINTER() antlr.TerminalNode {
	return s.GetToken(GraffleParserPRINTER, 0)
}

func (s *Built_func_printContext) VAR() antlr.TerminalNode {
	return s.GetToken(GraffleParserVAR, 0)
}

func (s *Built_func_printContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *Built_func_printContext) Function_call() IFunction_callContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunction_callContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunction_callContext)
}

func (s *Built_func_printContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Built_func_printContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Built_func_printContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterBuilt_func_print(s)
	}
}

func (s *Built_func_printContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitBuilt_func_print(s)
	}
}

func (s *Built_func_printContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitBuilt_func_print(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Built_func_print() (localctx IBuilt_func_printContext) {
	localctx = NewBuilt_func_printContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, GraffleParserRULE_built_func_print)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(527)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 68, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(521)
			p.Match(GraffleParserPRINTER)
		}
		{
			p.SetState(522)
			p.Match(GraffleParserVAR)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(523)
			p.Match(GraffleParserPRINTER)
		}
		{
			p.SetState(524)
			p.Value()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(525)
			p.Match(GraffleParserPRINTER)
		}
		{
			p.SetState(526)
			p.Function_call()
		}

	}

	return localctx
}

// IBuilt_func_inputContext is an interface to support dynamic dispatch.
type IBuilt_func_inputContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBuilt_func_inputContext differentiates from other interfaces.
	IsBuilt_func_inputContext()
}

type Built_func_inputContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBuilt_func_inputContext() *Built_func_inputContext {
	var p = new(Built_func_inputContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_built_func_input
	return p
}

func (*Built_func_inputContext) IsBuilt_func_inputContext() {}

func NewBuilt_func_inputContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Built_func_inputContext {
	var p = new(Built_func_inputContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_built_func_input

	return p
}

func (s *Built_func_inputContext) GetParser() antlr.Parser { return s.parser }

func (s *Built_func_inputContext) KEY_INPUT() antlr.TerminalNode {
	return s.GetToken(GraffleParserKEY_INPUT, 0)
}

func (s *Built_func_inputContext) ID() antlr.TerminalNode {
	return s.GetToken(GraffleParserID, 0)
}

func (s *Built_func_inputContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Built_func_inputContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Built_func_inputContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterBuilt_func_input(s)
	}
}

func (s *Built_func_inputContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitBuilt_func_input(s)
	}
}

func (s *Built_func_inputContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitBuilt_func_input(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Built_func_input() (localctx IBuilt_func_inputContext) {
	localctx = NewBuilt_func_inputContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, GraffleParserRULE_built_func_input)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(529)
		p.Match(GraffleParserKEY_INPUT)
	}
	{
		p.SetState(530)
		p.Match(GraffleParserID)
	}

	return localctx
}

// IFunction_callContext is an interface to support dynamic dispatch.
type IFunction_callContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunction_callContext differentiates from other interfaces.
	IsFunction_callContext()
}

type Function_callContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_callContext() *Function_callContext {
	var p = new(Function_callContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_function_call
	return p
}

func (*Function_callContext) IsFunction_callContext() {}

func NewFunction_callContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_callContext {
	var p = new(Function_callContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_function_call

	return p
}

func (s *Function_callContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_callContext) ID() antlr.TerminalNode {
	return s.GetToken(GraffleParserID, 0)
}

func (s *Function_callContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(GraffleParserL_PAREN, 0)
}

func (s *Function_callContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(GraffleParserR_PAREN, 0)
}

func (s *Function_callContext) AllValue() []IValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueContext)(nil)).Elem())
	var tst = make([]IValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueContext)
		}
	}

	return tst
}

func (s *Function_callContext) Value(i int) IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *Function_callContext) AllARG_DELIM() []antlr.TerminalNode {
	return s.GetTokens(GraffleParserARG_DELIM)
}

func (s *Function_callContext) ARG_DELIM(i int) antlr.TerminalNode {
	return s.GetToken(GraffleParserARG_DELIM, i)
}

func (s *Function_callContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_callContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_callContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterFunction_call(s)
	}
}

func (s *Function_callContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitFunction_call(s)
	}
}

func (s *Function_callContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitFunction_call(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Function_call() (localctx IFunction_callContext) {
	localctx = NewFunction_callContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, GraffleParserRULE_function_call)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(532)
		p.Match(GraffleParserID)
	}
	{
		p.SetState(533)
		p.Match(GraffleParserL_PAREN)
	}
	p.SetState(542)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GraffleParserNUMBER)|(1<<GraffleParserSTRING)|(1<<GraffleParserBOOL)|(1<<GraffleParserL_PAREN))) != 0) || (((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(GraffleParserNOT-40))|(1<<(GraffleParserVAR-40))|(1<<(GraffleParserID-40)))) != 0) {
		{
			p.SetState(534)
			p.Value()
		}
		p.SetState(539)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == GraffleParserARG_DELIM {
			{
				p.SetState(535)
				p.Match(GraffleParserARG_DELIM)
			}
			{
				p.SetState(536)
				p.Value()
			}

			p.SetState(541)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(544)
		p.Match(GraffleParserR_PAREN)
	}

	return localctx
}

// ILabelContext is an interface to support dynamic dispatch.
type ILabelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLabelContext differentiates from other interfaces.
	IsLabelContext()
}

type LabelContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLabelContext() *LabelContext {
	var p = new(LabelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_label
	return p
}

func (*LabelContext) IsLabelContext() {}

func NewLabelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LabelContext {
	var p = new(LabelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_label

	return p
}

func (s *LabelContext) GetParser() antlr.Parser { return s.parser }

func (s *LabelContext) LABEL() antlr.TerminalNode {
	return s.GetToken(GraffleParserLABEL, 0)
}

func (s *LabelContext) ML_LABEL() antlr.TerminalNode {
	return s.GetToken(GraffleParserML_LABEL, 0)
}

func (s *LabelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LabelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterLabel(s)
	}
}

func (s *LabelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitLabel(s)
	}
}

func (s *LabelContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitLabel(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Label() (localctx ILabelContext) {
	localctx = NewLabelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, GraffleParserRULE_label)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(546)
		_la = p.GetTokenStream().LA(1)

		if !(_la == GraffleParserLABEL || _la == GraffleParserML_LABEL) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) Logical_expr() ILogical_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILogical_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILogical_exprContext)
}

func (s *ValueContext) Arithm_expr() IArithm_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArithm_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArithm_exprContext)
}

func (s *ValueContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ValueContext) Builtin() IBuiltinContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBuiltinContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBuiltinContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitValue(s)
	}
}

func (s *ValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, GraffleParserRULE_value)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(552)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 71, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(548)
			p.Logical_expr()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(549)
			p.Arithm_expr()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(550)
			p.Expr()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(551)
			p.Builtin()
		}

	}

	return localctx
}

// IBuiltinContext is an interface to support dynamic dispatch.
type IBuiltinContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBuiltinContext differentiates from other interfaces.
	IsBuiltinContext()
}

type BuiltinContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBuiltinContext() *BuiltinContext {
	var p = new(BuiltinContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_builtin
	return p
}

func (*BuiltinContext) IsBuiltinContext() {}

func NewBuiltinContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BuiltinContext {
	var p = new(BuiltinContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_builtin

	return p
}

func (s *BuiltinContext) GetParser() antlr.Parser { return s.parser }

func (s *BuiltinContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(GraffleParserNUMBER, 0)
}

func (s *BuiltinContext) STRING() antlr.TerminalNode {
	return s.GetToken(GraffleParserSTRING, 0)
}

func (s *BuiltinContext) BOOL() antlr.TerminalNode {
	return s.GetToken(GraffleParserBOOL, 0)
}

func (s *BuiltinContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BuiltinContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BuiltinContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterBuiltin(s)
	}
}

func (s *BuiltinContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitBuiltin(s)
	}
}

func (s *BuiltinContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitBuiltin(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Builtin() (localctx IBuiltinContext) {
	localctx = NewBuiltinContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, GraffleParserRULE_builtin)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(554)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GraffleParserNUMBER)|(1<<GraffleParserSTRING)|(1<<GraffleParserBOOL))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IBlock_endContext is an interface to support dynamic dispatch.
type IBlock_endContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlock_endContext differentiates from other interfaces.
	IsBlock_endContext()
}

type Block_endContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlock_endContext() *Block_endContext {
	var p = new(Block_endContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraffleParserRULE_block_end
	return p
}

func (*Block_endContext) IsBlock_endContext() {}

func NewBlock_endContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Block_endContext {
	var p = new(Block_endContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraffleParserRULE_block_end

	return p
}

func (s *Block_endContext) GetParser() antlr.Parser { return s.parser }

func (s *Block_endContext) BLOCK_END() antlr.TerminalNode {
	return s.GetToken(GraffleParserBLOCK_END, 0)
}

func (s *Block_endContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(GraffleParserNEWLINE)
}

func (s *Block_endContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(GraffleParserNEWLINE, i)
}

func (s *Block_endContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Block_endContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Block_endContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.EnterBlock_end(s)
	}
}

func (s *Block_endContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraffleParserListener); ok {
		listenerT.ExitBlock_end(s)
	}
}

func (s *Block_endContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraffleParserVisitor:
		return t.VisitBlock_end(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraffleParser) Block_end() (localctx IBlock_endContext) {
	localctx = NewBlock_endContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, GraffleParserRULE_block_end)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(557)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraffleParserNEWLINE {
		{
			p.SetState(556)
			p.Match(GraffleParserNEWLINE)
		}

	}
	{
		p.SetState(559)
		p.Match(GraffleParserBLOCK_END)
	}
	p.SetState(561)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 73, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(560)
			p.Match(GraffleParserNEWLINE)
		}

	}

	return localctx
}
