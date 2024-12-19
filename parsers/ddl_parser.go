// Code generated from DDLParser.g4 by ANTLR 4.7.1. DO NOT EDIT.

package ddlparser // DDLParser
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 100, 450,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 3, 2, 3, 2, 3, 2, 7, 2, 56, 10,
	2, 12, 2, 14, 2, 59, 11, 2, 3, 2, 5, 2, 62, 10, 2, 3, 2, 3, 2, 3, 3, 3,
	3, 5, 3, 68, 10, 3, 3, 3, 3, 3, 5, 3, 72, 10, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	5, 3, 78, 10, 3, 3, 3, 7, 3, 81, 10, 3, 12, 3, 14, 3, 84, 11, 3, 5, 3,
	86, 10, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6,
	7, 6, 98, 10, 6, 12, 6, 14, 6, 101, 11, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7,
	5, 7, 108, 10, 7, 3, 8, 3, 8, 7, 8, 112, 10, 8, 12, 8, 14, 8, 115, 11,
	8, 3, 9, 3, 9, 5, 9, 119, 10, 9, 3, 9, 5, 9, 122, 10, 9, 3, 9, 3, 9, 3,
	9, 5, 9, 127, 10, 9, 3, 9, 5, 9, 130, 10, 9, 3, 9, 3, 9, 5, 9, 134, 10,
	9, 3, 9, 3, 9, 3, 9, 5, 9, 139, 10, 9, 3, 9, 5, 9, 142, 10, 9, 3, 9, 3,
	9, 3, 9, 5, 9, 147, 10, 9, 3, 9, 5, 9, 150, 10, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 5, 9, 156, 10, 9, 3, 9, 5, 9, 159, 10, 9, 3, 9, 3, 9, 5, 9, 163, 10,
	9, 3, 9, 5, 9, 166, 10, 9, 3, 9, 5, 9, 169, 10, 9, 3, 9, 3, 9, 5, 9, 173,
	10, 9, 3, 9, 5, 9, 176, 10, 9, 3, 9, 5, 9, 179, 10, 9, 3, 9, 3, 9, 5, 9,
	183, 10, 9, 3, 9, 5, 9, 186, 10, 9, 3, 9, 5, 9, 189, 10, 9, 3, 9, 5, 9,
	192, 10, 9, 3, 9, 3, 9, 5, 9, 196, 10, 9, 3, 9, 5, 9, 199, 10, 9, 3, 9,
	5, 9, 202, 10, 9, 3, 9, 3, 9, 3, 9, 5, 9, 207, 10, 9, 3, 9, 3, 9, 3, 9,
	5, 9, 212, 10, 9, 3, 9, 3, 9, 3, 9, 5, 9, 217, 10, 9, 3, 9, 5, 9, 220,
	10, 9, 5, 9, 222, 10, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 238, 10, 12, 3,
	12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 7, 13, 246, 10, 13, 12, 13, 14,
	13, 249, 11, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 5, 14, 260, 10, 14, 3, 14, 3, 14, 3, 14, 3, 14, 5, 14, 266, 10,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 5, 14, 274, 10, 14, 3, 15,
	5, 15, 277, 10, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 5,
	17, 286, 10, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 5, 17, 293, 10, 17,
	5, 17, 295, 10, 17, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3,
	20, 3, 20, 3, 20, 3, 20, 3, 20, 5, 20, 309, 10, 20, 3, 20, 5, 20, 312,
	10, 20, 3, 21, 3, 21, 3, 21, 5, 21, 317, 10, 21, 3, 21, 5, 21, 320, 10,
	21, 3, 21, 3, 21, 3, 21, 5, 21, 325, 10, 21, 3, 21, 5, 21, 328, 10, 21,
	3, 22, 3, 22, 5, 22, 332, 10, 22, 5, 22, 334, 10, 22, 3, 22, 3, 22, 5,
	22, 338, 10, 22, 3, 22, 5, 22, 341, 10, 22, 3, 22, 3, 22, 7, 22, 345, 10,
	22, 12, 22, 14, 22, 348, 11, 22, 3, 22, 3, 22, 5, 22, 352, 10, 22, 5, 22,
	354, 10, 22, 3, 22, 5, 22, 357, 10, 22, 3, 22, 5, 22, 360, 10, 22, 3, 22,
	5, 22, 363, 10, 22, 3, 22, 5, 22, 366, 10, 22, 3, 22, 3, 22, 7, 22, 370,
	10, 22, 12, 22, 14, 22, 373, 11, 22, 5, 22, 375, 10, 22, 3, 23, 3, 23,
	5, 23, 379, 10, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 5, 23, 387,
	10, 23, 3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 5, 25,
	397, 10, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 5, 25, 404, 10, 25, 7,
	25, 406, 10, 25, 12, 25, 14, 25, 409, 11, 25, 3, 25, 3, 25, 3, 26, 3, 26,
	5, 26, 415, 10, 26, 3, 26, 3, 26, 3, 26, 5, 26, 420, 10, 26, 3, 26, 3,
	26, 5, 26, 424, 10, 26, 3, 26, 3, 26, 3, 26, 5, 26, 429, 10, 26, 3, 26,
	5, 26, 432, 10, 26, 3, 26, 3, 26, 5, 26, 436, 10, 26, 3, 26, 3, 26, 5,
	26, 440, 10, 26, 3, 26, 3, 26, 3, 26, 5, 26, 445, 10, 26, 3, 26, 5, 26,
	448, 10, 26, 3, 26, 2, 2, 27, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
	26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 2, 19, 3, 2, 96, 97,
	10, 2, 13, 13, 42, 42, 45, 45, 47, 47, 52, 52, 64, 64, 69, 69, 77, 77,
	4, 2, 5, 5, 96, 97, 4, 2, 11, 11, 77, 77, 4, 2, 11, 11, 13, 13, 7, 2, 4,
	4, 34, 35, 44, 44, 61, 61, 68, 68, 4, 2, 60, 60, 72, 72, 5, 2, 21, 22,
	29, 30, 51, 51, 10, 2, 7, 7, 9, 10, 20, 20, 36, 36, 41, 41, 43, 43, 58,
	58, 67, 67, 7, 2, 5, 6, 19, 19, 65, 66, 76, 76, 80, 80, 4, 2, 26, 26, 59,
	59, 4, 2, 48, 48, 87, 90, 4, 2, 28, 28, 70, 70, 4, 2, 18, 18, 39, 40, 4,
	2, 94, 94, 97, 97, 4, 2, 32, 32, 38, 38, 4, 2, 8, 8, 31, 31, 2, 530, 2,
	52, 3, 2, 2, 2, 4, 65, 3, 2, 2, 2, 6, 87, 3, 2, 2, 2, 8, 89, 3, 2, 2, 2,
	10, 93, 3, 2, 2, 2, 12, 107, 3, 2, 2, 2, 14, 109, 3, 2, 2, 2, 16, 221,
	3, 2, 2, 2, 18, 223, 3, 2, 2, 2, 20, 227, 3, 2, 2, 2, 22, 233, 3, 2, 2,
	2, 24, 241, 3, 2, 2, 2, 26, 273, 3, 2, 2, 2, 28, 276, 3, 2, 2, 2, 30, 280,
	3, 2, 2, 2, 32, 294, 3, 2, 2, 2, 34, 296, 3, 2, 2, 2, 36, 299, 3, 2, 2,
	2, 38, 311, 3, 2, 2, 2, 40, 327, 3, 2, 2, 2, 42, 374, 3, 2, 2, 2, 44, 386,
	3, 2, 2, 2, 46, 388, 3, 2, 2, 2, 48, 391, 3, 2, 2, 2, 50, 447, 3, 2, 2,
	2, 52, 57, 5, 4, 3, 2, 53, 54, 7, 83, 2, 2, 54, 56, 5, 4, 3, 2, 55, 53,
	3, 2, 2, 2, 56, 59, 3, 2, 2, 2, 57, 55, 3, 2, 2, 2, 57, 58, 3, 2, 2, 2,
	58, 61, 3, 2, 2, 2, 59, 57, 3, 2, 2, 2, 60, 62, 7, 83, 2, 2, 61, 60, 3,
	2, 2, 2, 61, 62, 3, 2, 2, 2, 62, 63, 3, 2, 2, 2, 63, 64, 7, 2, 2, 3, 64,
	3, 3, 2, 2, 2, 65, 67, 7, 17, 2, 2, 66, 68, 7, 63, 2, 2, 67, 66, 3, 2,
	2, 2, 67, 68, 3, 2, 2, 2, 68, 69, 3, 2, 2, 2, 69, 71, 7, 62, 2, 2, 70,
	72, 5, 8, 5, 2, 71, 70, 3, 2, 2, 2, 71, 72, 3, 2, 2, 2, 72, 73, 3, 2, 2,
	2, 73, 74, 5, 6, 4, 2, 74, 85, 5, 10, 6, 2, 75, 82, 5, 50, 26, 2, 76, 78,
	7, 84, 2, 2, 77, 76, 3, 2, 2, 2, 77, 78, 3, 2, 2, 2, 78, 79, 3, 2, 2, 2,
	79, 81, 5, 50, 26, 2, 80, 77, 3, 2, 2, 2, 81, 84, 3, 2, 2, 2, 82, 80, 3,
	2, 2, 2, 82, 83, 3, 2, 2, 2, 83, 86, 3, 2, 2, 2, 84, 82, 3, 2, 2, 2, 85,
	75, 3, 2, 2, 2, 85, 86, 3, 2, 2, 2, 86, 5, 3, 2, 2, 2, 87, 88, 9, 2, 2,
	2, 88, 7, 3, 2, 2, 2, 89, 90, 7, 33, 2, 2, 90, 91, 7, 48, 2, 2, 91, 92,
	7, 27, 2, 2, 92, 9, 3, 2, 2, 2, 93, 94, 7, 85, 2, 2, 94, 99, 5, 12, 7,
	2, 95, 96, 7, 84, 2, 2, 96, 98, 5, 12, 7, 2, 97, 95, 3, 2, 2, 2, 98, 101,
	3, 2, 2, 2, 99, 97, 3, 2, 2, 2, 99, 100, 3, 2, 2, 2, 100, 102, 3, 2, 2,
	2, 101, 99, 3, 2, 2, 2, 102, 103, 7, 86, 2, 2, 103, 11, 3, 2, 2, 2, 104,
	105, 9, 2, 2, 2, 105, 108, 5, 14, 8, 2, 106, 108, 5, 42, 22, 2, 107, 104,
	3, 2, 2, 2, 107, 106, 3, 2, 2, 2, 108, 13, 3, 2, 2, 2, 109, 113, 5, 16,
	9, 2, 110, 112, 5, 26, 14, 2, 111, 110, 3, 2, 2, 2, 112, 115, 3, 2, 2,
	2, 113, 111, 3, 2, 2, 2, 113, 114, 3, 2, 2, 2, 114, 15, 3, 2, 2, 2, 115,
	113, 3, 2, 2, 2, 116, 118, 9, 3, 2, 2, 117, 119, 5, 18, 10, 2, 118, 117,
	3, 2, 2, 2, 118, 119, 3, 2, 2, 2, 119, 121, 3, 2, 2, 2, 120, 122, 7, 5,
	2, 2, 121, 120, 3, 2, 2, 2, 121, 122, 3, 2, 2, 2, 122, 129, 3, 2, 2, 2,
	123, 124, 7, 11, 2, 2, 124, 127, 7, 59, 2, 2, 125, 127, 7, 12, 2, 2, 126,
	123, 3, 2, 2, 2, 126, 125, 3, 2, 2, 2, 127, 128, 3, 2, 2, 2, 128, 130,
	9, 4, 2, 2, 129, 126, 3, 2, 2, 2, 129, 130, 3, 2, 2, 2, 130, 133, 3, 2,
	2, 2, 131, 132, 7, 14, 2, 2, 132, 134, 9, 2, 2, 2, 133, 131, 3, 2, 2, 2,
	133, 134, 3, 2, 2, 2, 134, 222, 3, 2, 2, 2, 135, 136, 7, 46, 2, 2, 136,
	138, 9, 5, 2, 2, 137, 139, 5, 18, 10, 2, 138, 137, 3, 2, 2, 2, 138, 139,
	3, 2, 2, 2, 139, 141, 3, 2, 2, 2, 140, 142, 7, 5, 2, 2, 141, 140, 3, 2,
	2, 2, 141, 142, 3, 2, 2, 2, 142, 222, 3, 2, 2, 2, 143, 144, 7, 47, 2, 2,
	144, 146, 7, 77, 2, 2, 145, 147, 5, 18, 10, 2, 146, 145, 3, 2, 2, 2, 146,
	147, 3, 2, 2, 2, 147, 149, 3, 2, 2, 2, 148, 150, 7, 5, 2, 2, 149, 148,
	3, 2, 2, 2, 149, 150, 3, 2, 2, 2, 150, 222, 3, 2, 2, 2, 151, 152, 7, 46,
	2, 2, 152, 153, 9, 6, 2, 2, 153, 155, 7, 78, 2, 2, 154, 156, 5, 18, 10,
	2, 155, 154, 3, 2, 2, 2, 155, 156, 3, 2, 2, 2, 156, 158, 3, 2, 2, 2, 157,
	159, 7, 5, 2, 2, 158, 157, 3, 2, 2, 2, 158, 159, 3, 2, 2, 2, 159, 222,
	3, 2, 2, 2, 160, 162, 9, 7, 2, 2, 161, 163, 5, 18, 10, 2, 162, 161, 3,
	2, 2, 2, 162, 163, 3, 2, 2, 2, 163, 165, 3, 2, 2, 2, 164, 166, 9, 8, 2,
	2, 165, 164, 3, 2, 2, 2, 165, 166, 3, 2, 2, 2, 166, 168, 3, 2, 2, 2, 167,
	169, 7, 81, 2, 2, 168, 167, 3, 2, 2, 2, 168, 169, 3, 2, 2, 2, 169, 222,
	3, 2, 2, 2, 170, 172, 7, 57, 2, 2, 171, 173, 5, 20, 11, 2, 172, 171, 3,
	2, 2, 2, 172, 173, 3, 2, 2, 2, 173, 175, 3, 2, 2, 2, 174, 176, 9, 8, 2,
	2, 175, 174, 3, 2, 2, 2, 175, 176, 3, 2, 2, 2, 176, 178, 3, 2, 2, 2, 177,
	179, 7, 81, 2, 2, 178, 177, 3, 2, 2, 2, 178, 179, 3, 2, 2, 2, 179, 222,
	3, 2, 2, 2, 180, 182, 7, 24, 2, 2, 181, 183, 7, 55, 2, 2, 182, 181, 3,
	2, 2, 2, 182, 183, 3, 2, 2, 2, 183, 185, 3, 2, 2, 2, 184, 186, 5, 20, 11,
	2, 185, 184, 3, 2, 2, 2, 185, 186, 3, 2, 2, 2, 186, 188, 3, 2, 2, 2, 187,
	189, 9, 8, 2, 2, 188, 187, 3, 2, 2, 2, 188, 189, 3, 2, 2, 2, 189, 191,
	3, 2, 2, 2, 190, 192, 7, 81, 2, 2, 191, 190, 3, 2, 2, 2, 191, 192, 3, 2,
	2, 2, 192, 222, 3, 2, 2, 2, 193, 195, 9, 9, 2, 2, 194, 196, 5, 22, 12,
	2, 195, 194, 3, 2, 2, 2, 195, 196, 3, 2, 2, 2, 196, 198, 3, 2, 2, 2, 197,
	199, 9, 8, 2, 2, 198, 197, 3, 2, 2, 2, 198, 199, 3, 2, 2, 2, 199, 201,
	3, 2, 2, 2, 200, 202, 7, 81, 2, 2, 201, 200, 3, 2, 2, 2, 201, 202, 3, 2,
	2, 2, 202, 222, 3, 2, 2, 2, 203, 222, 9, 10, 2, 2, 204, 206, 9, 11, 2,
	2, 205, 207, 5, 18, 10, 2, 206, 205, 3, 2, 2, 2, 206, 207, 3, 2, 2, 2,
	207, 222, 3, 2, 2, 2, 208, 209, 9, 12, 2, 2, 209, 211, 5, 24, 13, 2, 210,
	212, 7, 5, 2, 2, 211, 210, 3, 2, 2, 2, 211, 212, 3, 2, 2, 2, 212, 219,
	3, 2, 2, 2, 213, 214, 7, 11, 2, 2, 214, 217, 7, 59, 2, 2, 215, 217, 7,
	12, 2, 2, 216, 213, 3, 2, 2, 2, 216, 215, 3, 2, 2, 2, 217, 218, 3, 2, 2,
	2, 218, 220, 9, 4, 2, 2, 219, 216, 3, 2, 2, 2, 219, 220, 3, 2, 2, 2, 220,
	222, 3, 2, 2, 2, 221, 116, 3, 2, 2, 2, 221, 135, 3, 2, 2, 2, 221, 143,
	3, 2, 2, 2, 221, 151, 3, 2, 2, 2, 221, 160, 3, 2, 2, 2, 221, 170, 3, 2,
	2, 2, 221, 180, 3, 2, 2, 2, 221, 193, 3, 2, 2, 2, 221, 203, 3, 2, 2, 2,
	221, 204, 3, 2, 2, 2, 221, 208, 3, 2, 2, 2, 222, 17, 3, 2, 2, 2, 223, 224,
	7, 85, 2, 2, 224, 225, 7, 91, 2, 2, 225, 226, 7, 86, 2, 2, 226, 19, 3,
	2, 2, 2, 227, 228, 7, 85, 2, 2, 228, 229, 7, 91, 2, 2, 229, 230, 7, 84,
	2, 2, 230, 231, 7, 91, 2, 2, 231, 232, 7, 86, 2, 2, 232, 21, 3, 2, 2, 2,
	233, 234, 7, 85, 2, 2, 234, 237, 7, 91, 2, 2, 235, 236, 7, 84, 2, 2, 236,
	238, 7, 91, 2, 2, 237, 235, 3, 2, 2, 2, 237, 238, 3, 2, 2, 2, 238, 239,
	3, 2, 2, 2, 239, 240, 7, 86, 2, 2, 240, 23, 3, 2, 2, 2, 241, 242, 7, 85,
	2, 2, 242, 247, 7, 94, 2, 2, 243, 244, 7, 84, 2, 2, 244, 246, 7, 94, 2,
	2, 245, 243, 3, 2, 2, 2, 246, 249, 3, 2, 2, 2, 247, 245, 3, 2, 2, 2, 247,
	248, 3, 2, 2, 2, 248, 250, 3, 2, 2, 2, 249, 247, 3, 2, 2, 2, 250, 251,
	7, 86, 2, 2, 251, 25, 3, 2, 2, 2, 252, 274, 5, 28, 15, 2, 253, 254, 7,
	23, 2, 2, 254, 274, 5, 32, 17, 2, 255, 260, 7, 3, 2, 2, 256, 257, 7, 53,
	2, 2, 257, 258, 7, 73, 2, 2, 258, 260, 5, 40, 21, 2, 259, 255, 3, 2, 2,
	2, 259, 256, 3, 2, 2, 2, 260, 274, 3, 2, 2, 2, 261, 274, 5, 34, 18, 2,
	262, 274, 7, 38, 2, 2, 263, 265, 7, 71, 2, 2, 264, 266, 7, 38, 2, 2, 265,
	264, 3, 2, 2, 2, 265, 266, 3, 2, 2, 2, 266, 274, 3, 2, 2, 2, 267, 274,
	5, 30, 16, 2, 268, 269, 7, 14, 2, 2, 269, 274, 7, 96, 2, 2, 270, 271, 7,
	58, 2, 2, 271, 272, 7, 23, 2, 2, 272, 274, 7, 75, 2, 2, 273, 252, 3, 2,
	2, 2, 273, 253, 3, 2, 2, 2, 273, 259, 3, 2, 2, 2, 273, 261, 3, 2, 2, 2,
	273, 262, 3, 2, 2, 2, 273, 263, 3, 2, 2, 2, 273, 267, 3, 2, 2, 2, 273,
	268, 3, 2, 2, 2, 273, 270, 3, 2, 2, 2, 274, 27, 3, 2, 2, 2, 275, 277, 7,
	48, 2, 2, 276, 275, 3, 2, 2, 2, 276, 277, 3, 2, 2, 2, 277, 278, 3, 2, 2,
	2, 278, 279, 7, 50, 2, 2, 279, 29, 3, 2, 2, 2, 280, 281, 7, 15, 2, 2, 281,
	282, 7, 94, 2, 2, 282, 31, 3, 2, 2, 2, 283, 295, 7, 50, 2, 2, 284, 286,
	5, 36, 19, 2, 285, 284, 3, 2, 2, 2, 285, 286, 3, 2, 2, 2, 286, 287, 3,
	2, 2, 2, 287, 295, 5, 38, 20, 2, 288, 292, 5, 40, 21, 2, 289, 290, 7, 53,
	2, 2, 290, 291, 7, 73, 2, 2, 291, 293, 5, 40, 21, 2, 292, 289, 3, 2, 2,
	2, 292, 293, 3, 2, 2, 2, 293, 295, 3, 2, 2, 2, 294, 283, 3, 2, 2, 2, 294,
	285, 3, 2, 2, 2, 294, 288, 3, 2, 2, 2, 295, 33, 3, 2, 2, 2, 296, 297, 7,
	56, 2, 2, 297, 298, 7, 38, 2, 2, 298, 35, 3, 2, 2, 2, 299, 300, 9, 13,
	2, 2, 300, 37, 3, 2, 2, 2, 301, 312, 7, 94, 2, 2, 302, 312, 7, 91, 2, 2,
	303, 304, 7, 90, 2, 2, 304, 312, 7, 91, 2, 2, 305, 312, 9, 14, 2, 2, 306,
	312, 7, 92, 2, 2, 307, 309, 7, 48, 2, 2, 308, 307, 3, 2, 2, 2, 308, 309,
	3, 2, 2, 2, 309, 310, 3, 2, 2, 2, 310, 312, 7, 50, 2, 2, 311, 301, 3, 2,
	2, 2, 311, 302, 3, 2, 2, 2, 311, 303, 3, 2, 2, 2, 311, 305, 3, 2, 2, 2,
	311, 306, 3, 2, 2, 2, 311, 308, 3, 2, 2, 2, 312, 39, 3, 2, 2, 2, 313, 319,
	9, 15, 2, 2, 314, 316, 7, 85, 2, 2, 315, 317, 7, 91, 2, 2, 316, 315, 3,
	2, 2, 2, 316, 317, 3, 2, 2, 2, 317, 318, 3, 2, 2, 2, 318, 320, 7, 86, 2,
	2, 319, 314, 3, 2, 2, 2, 319, 320, 3, 2, 2, 2, 320, 328, 3, 2, 2, 2, 321,
	322, 7, 49, 2, 2, 322, 324, 7, 85, 2, 2, 323, 325, 7, 91, 2, 2, 324, 323,
	3, 2, 2, 2, 324, 325, 3, 2, 2, 2, 325, 326, 3, 2, 2, 2, 326, 328, 7, 86,
	2, 2, 327, 313, 3, 2, 2, 2, 327, 321, 3, 2, 2, 2, 328, 41, 3, 2, 2, 2,
	329, 331, 7, 16, 2, 2, 330, 332, 9, 16, 2, 2, 331, 330, 3, 2, 2, 2, 331,
	332, 3, 2, 2, 2, 332, 334, 3, 2, 2, 2, 333, 329, 3, 2, 2, 2, 333, 334,
	3, 2, 2, 2, 334, 335, 3, 2, 2, 2, 335, 337, 5, 34, 18, 2, 336, 338, 9,
	16, 2, 2, 337, 336, 3, 2, 2, 2, 337, 338, 3, 2, 2, 2, 338, 340, 3, 2, 2,
	2, 339, 341, 9, 2, 2, 2, 340, 339, 3, 2, 2, 2, 340, 341, 3, 2, 2, 2, 341,
	342, 3, 2, 2, 2, 342, 346, 5, 48, 25, 2, 343, 345, 5, 44, 23, 2, 344, 343,
	3, 2, 2, 2, 345, 348, 3, 2, 2, 2, 346, 344, 3, 2, 2, 2, 346, 347, 3, 2,
	2, 2, 347, 375, 3, 2, 2, 2, 348, 346, 3, 2, 2, 2, 349, 351, 7, 16, 2, 2,
	350, 352, 9, 16, 2, 2, 351, 350, 3, 2, 2, 2, 351, 352, 3, 2, 2, 2, 352,
	354, 3, 2, 2, 2, 353, 349, 3, 2, 2, 2, 353, 354, 3, 2, 2, 2, 354, 356,
	3, 2, 2, 2, 355, 357, 7, 71, 2, 2, 356, 355, 3, 2, 2, 2, 356, 357, 3, 2,
	2, 2, 357, 359, 3, 2, 2, 2, 358, 360, 9, 17, 2, 2, 359, 358, 3, 2, 2, 2,
	359, 360, 3, 2, 2, 2, 360, 362, 3, 2, 2, 2, 361, 363, 9, 2, 2, 2, 362,
	361, 3, 2, 2, 2, 362, 363, 3, 2, 2, 2, 363, 365, 3, 2, 2, 2, 364, 366,
	5, 46, 24, 2, 365, 364, 3, 2, 2, 2, 365, 366, 3, 2, 2, 2, 366, 367, 3,
	2, 2, 2, 367, 371, 5, 48, 25, 2, 368, 370, 5, 44, 23, 2, 369, 368, 3, 2,
	2, 2, 370, 373, 3, 2, 2, 2, 371, 369, 3, 2, 2, 2, 371, 372, 3, 2, 2, 2,
	372, 375, 3, 2, 2, 2, 373, 371, 3, 2, 2, 2, 374, 333, 3, 2, 2, 2, 374,
	353, 3, 2, 2, 2, 375, 43, 3, 2, 2, 2, 376, 378, 7, 37, 2, 2, 377, 379,
	7, 82, 2, 2, 378, 377, 3, 2, 2, 2, 378, 379, 3, 2, 2, 2, 379, 380, 3, 2,
	2, 2, 380, 387, 7, 93, 2, 2, 381, 382, 7, 79, 2, 2, 382, 383, 7, 54, 2,
	2, 383, 387, 9, 16, 2, 2, 384, 385, 7, 15, 2, 2, 385, 387, 7, 94, 2, 2,
	386, 376, 3, 2, 2, 2, 386, 381, 3, 2, 2, 2, 386, 384, 3, 2, 2, 2, 387,
	45, 3, 2, 2, 2, 388, 389, 7, 74, 2, 2, 389, 390, 9, 18, 2, 2, 390, 47,
	3, 2, 2, 2, 391, 392, 7, 85, 2, 2, 392, 396, 9, 2, 2, 2, 393, 394, 7, 85,
	2, 2, 394, 395, 7, 91, 2, 2, 395, 397, 7, 86, 2, 2, 396, 393, 3, 2, 2,
	2, 396, 397, 3, 2, 2, 2, 397, 407, 3, 2, 2, 2, 398, 399, 7, 84, 2, 2, 399,
	403, 9, 2, 2, 2, 400, 401, 7, 85, 2, 2, 401, 402, 7, 91, 2, 2, 402, 404,
	7, 86, 2, 2, 403, 400, 3, 2, 2, 2, 403, 404, 3, 2, 2, 2, 404, 406, 3, 2,
	2, 2, 405, 398, 3, 2, 2, 2, 406, 409, 3, 2, 2, 2, 407, 405, 3, 2, 2, 2,
	407, 408, 3, 2, 2, 2, 408, 410, 3, 2, 2, 2, 409, 407, 3, 2, 2, 2, 410,
	411, 7, 86, 2, 2, 411, 49, 3, 2, 2, 2, 412, 414, 7, 25, 2, 2, 413, 415,
	7, 82, 2, 2, 414, 413, 3, 2, 2, 2, 414, 415, 3, 2, 2, 2, 415, 416, 3, 2,
	2, 2, 416, 448, 9, 2, 2, 2, 417, 419, 7, 3, 2, 2, 418, 420, 7, 82, 2, 2,
	419, 418, 3, 2, 2, 2, 419, 420, 3, 2, 2, 2, 420, 421, 3, 2, 2, 2, 421,
	448, 7, 91, 2, 2, 422, 424, 7, 23, 2, 2, 423, 422, 3, 2, 2, 2, 423, 424,
	3, 2, 2, 2, 424, 428, 3, 2, 2, 2, 425, 426, 7, 11, 2, 2, 426, 429, 7, 59,
	2, 2, 427, 429, 7, 12, 2, 2, 428, 425, 3, 2, 2, 2, 428, 427, 3, 2, 2, 2,
	429, 431, 3, 2, 2, 2, 430, 432, 7, 82, 2, 2, 431, 430, 3, 2, 2, 2, 431,
	432, 3, 2, 2, 2, 432, 433, 3, 2, 2, 2, 433, 448, 9, 2, 2, 2, 434, 436,
	7, 23, 2, 2, 435, 434, 3, 2, 2, 2, 435, 436, 3, 2, 2, 2, 436, 437, 3, 2,
	2, 2, 437, 439, 7, 14, 2, 2, 438, 440, 7, 82, 2, 2, 439, 438, 3, 2, 2,
	2, 439, 440, 3, 2, 2, 2, 440, 441, 3, 2, 2, 2, 441, 448, 9, 2, 2, 2, 442,
	444, 7, 15, 2, 2, 443, 445, 7, 82, 2, 2, 444, 443, 3, 2, 2, 2, 444, 445,
	3, 2, 2, 2, 445, 446, 3, 2, 2, 2, 446, 448, 7, 94, 2, 2, 447, 412, 3, 2,
	2, 2, 447, 417, 3, 2, 2, 2, 447, 423, 3, 2, 2, 2, 447, 435, 3, 2, 2, 2,
	447, 442, 3, 2, 2, 2, 448, 51, 3, 2, 2, 2, 83, 57, 61, 67, 71, 77, 82,
	85, 99, 107, 113, 118, 121, 126, 129, 133, 138, 141, 146, 149, 155, 158,
	162, 165, 168, 172, 175, 178, 182, 185, 188, 191, 195, 198, 201, 206, 211,
	216, 219, 221, 237, 247, 259, 265, 273, 276, 285, 292, 294, 308, 311, 316,
	319, 324, 327, 331, 333, 337, 340, 346, 351, 353, 356, 359, 362, 365, 371,
	374, 378, 386, 396, 403, 407, 414, 419, 423, 428, 431, 435, 439, 444, 447,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "'='", "';'", "','", "'('", "')'", "'~'",
	"'!'", "'+'", "'-'",
}
var symbolicNames = []string{
	"", "AUTO_INCREMENT", "BIGINT", "BINARY", "BIT", "BLOB", "BTREE", "BOOLEAN",
	"BOOL", "CHARACTER", "CHARSET", "CHAR", "COLLATE", "COMMENT", "CONSTRAINT",
	"CREATE", "CURRENT_TIMESTAMP", "DATETIME", "DATE", "DECIMAL", "DEC", "DEFAULT",
	"DOUBLE", "ENGINE", "ENUM", "EXISTS", "FALSE", "FIXED", "FLOAT", "HASH",
	"INDEX", "IF", "INTEGER", "INT", "JSON", "KEY_BLOCK_SIZE", "KEY", "LOCALTIMESTAMP",
	"LOCALTIME", "LONGBLOB", "LONGTEXT", "MEDIUMBLOB", "MEDIUMINT", "MEDIUMTEXT",
	"NATIONAL", "NCHAR", "NOT", "NOW", "NULL", "NUMERIC", "NVARCHAR", "ON",
	"PARSER", "PRECISION", "PRIMARY", "REAL", "SERIAL", "SET", "SIGNED", "SMALLINT",
	"TABLE", "TEMPORARY", "TEXT", "TIMESTAMP", "TIME", "TINYBLOB", "TINYINT",
	"TINYTEXT", "TRUE", "UNIQUE", "UNSIGNED", "UPDATE", "USING", "VALUE", "VARBINARY",
	"VARCHAR", "VARYING", "WITH", "YEAR", "ZEROFILL", "Equal", "Semicolon",
	"Comma", "LeftParenthesis", "RightParenthesis", "Tilde", "Exclamation",
	"Plus", "Minus", "Integer", "Number", "FilesizeLiteral", "Literal", "PlaceholderString",
	"BackQuotedString", "RawString", "SPACE", "BLOCK_COMMENT", "LINE_COMMENT",
}

var ruleNames = []string{
	"root", "createTableDDL", "tableName", "ifNotExists", "createDefinitions",
	"createDefinition", "columnDefinition", "dataType", "lengthOneDimension",
	"lengthTwoDimension", "lengthTwoOptionalDimension", "collectionOptions",
	"columnConstraint", "nullNotNull", "comment", "defaultValue", "primaryKey",
	"unaryOperator", "constant", "currentTimestamp", "tableConstraint", "indexOption",
	"indexType", "indexColumnNames", "tableOption",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type DDLParser struct {
	*antlr.BaseParser
}

func NewDDLParser(input antlr.TokenStream) *DDLParser {
	this := new(DDLParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "DDLParser.g4"

	return this
}

// DDLParser tokens.
const (
	DDLParserEOF               = antlr.TokenEOF
	DDLParserAUTO_INCREMENT    = 1
	DDLParserBIGINT            = 2
	DDLParserBINARY            = 3
	DDLParserBIT               = 4
	DDLParserBLOB              = 5
	DDLParserBTREE             = 6
	DDLParserBOOLEAN           = 7
	DDLParserBOOL              = 8
	DDLParserCHARACTER         = 9
	DDLParserCHARSET           = 10
	DDLParserCHAR              = 11
	DDLParserCOLLATE           = 12
	DDLParserCOMMENT           = 13
	DDLParserCONSTRAINT        = 14
	DDLParserCREATE            = 15
	DDLParserCURRENT_TIMESTAMP = 16
	DDLParserDATETIME          = 17
	DDLParserDATE              = 18
	DDLParserDECIMAL           = 19
	DDLParserDEC               = 20
	DDLParserDEFAULT           = 21
	DDLParserDOUBLE            = 22
	DDLParserENGINE            = 23
	DDLParserENUM              = 24
	DDLParserEXISTS            = 25
	DDLParserFALSE             = 26
	DDLParserFIXED             = 27
	DDLParserFLOAT             = 28
	DDLParserHASH              = 29
	DDLParserINDEX             = 30
	DDLParserIF                = 31
	DDLParserINTEGER           = 32
	DDLParserINT               = 33
	DDLParserJSON              = 34
	DDLParserKEY_BLOCK_SIZE    = 35
	DDLParserKEY               = 36
	DDLParserLOCALTIMESTAMP    = 37
	DDLParserLOCALTIME         = 38
	DDLParserLONGBLOB          = 39
	DDLParserLONGTEXT          = 40
	DDLParserMEDIUMBLOB        = 41
	DDLParserMEDIUMINT         = 42
	DDLParserMEDIUMTEXT        = 43
	DDLParserNATIONAL          = 44
	DDLParserNCHAR             = 45
	DDLParserNOT               = 46
	DDLParserNOW               = 47
	DDLParserNULL              = 48
	DDLParserNUMERIC           = 49
	DDLParserNVARCHAR          = 50
	DDLParserON                = 51
	DDLParserPARSER            = 52
	DDLParserPRECISION         = 53
	DDLParserPRIMARY           = 54
	DDLParserREAL              = 55
	DDLParserSERIAL            = 56
	DDLParserSET               = 57
	DDLParserSIGNED            = 58
	DDLParserSMALLINT          = 59
	DDLParserTABLE             = 60
	DDLParserTEMPORARY         = 61
	DDLParserTEXT              = 62
	DDLParserTIMESTAMP         = 63
	DDLParserTIME              = 64
	DDLParserTINYBLOB          = 65
	DDLParserTINYINT           = 66
	DDLParserTINYTEXT          = 67
	DDLParserTRUE              = 68
	DDLParserUNIQUE            = 69
	DDLParserUNSIGNED          = 70
	DDLParserUPDATE            = 71
	DDLParserUSING             = 72
	DDLParserVALUE             = 73
	DDLParserVARBINARY         = 74
	DDLParserVARCHAR           = 75
	DDLParserVARYING           = 76
	DDLParserWITH              = 77
	DDLParserYEAR              = 78
	DDLParserZEROFILL          = 79
	DDLParserEqual             = 80
	DDLParserSemicolon         = 81
	DDLParserComma             = 82
	DDLParserLeftParenthesis   = 83
	DDLParserRightParenthesis  = 84
	DDLParserTilde             = 85
	DDLParserExclamation       = 86
	DDLParserPlus              = 87
	DDLParserMinus             = 88
	DDLParserInteger           = 89
	DDLParserNumber            = 90
	DDLParserFilesizeLiteral   = 91
	DDLParserLiteral           = 92
	DDLParserPlaceholderString = 93
	DDLParserBackQuotedString  = 94
	DDLParserRawString         = 95
	DDLParserSPACE             = 96
	DDLParserBLOCK_COMMENT     = 97
	DDLParserLINE_COMMENT      = 98
)

// DDLParser rules.
const (
	DDLParserRULE_root                       = 0
	DDLParserRULE_createTableDDL             = 1
	DDLParserRULE_tableName                  = 2
	DDLParserRULE_ifNotExists                = 3
	DDLParserRULE_createDefinitions          = 4
	DDLParserRULE_createDefinition           = 5
	DDLParserRULE_columnDefinition           = 6
	DDLParserRULE_dataType                   = 7
	DDLParserRULE_lengthOneDimension         = 8
	DDLParserRULE_lengthTwoDimension         = 9
	DDLParserRULE_lengthTwoOptionalDimension = 10
	DDLParserRULE_collectionOptions          = 11
	DDLParserRULE_columnConstraint           = 12
	DDLParserRULE_nullNotNull                = 13
	DDLParserRULE_comment                    = 14
	DDLParserRULE_defaultValue               = 15
	DDLParserRULE_primaryKey                 = 16
	DDLParserRULE_unaryOperator              = 17
	DDLParserRULE_constant                   = 18
	DDLParserRULE_currentTimestamp           = 19
	DDLParserRULE_tableConstraint            = 20
	DDLParserRULE_indexOption                = 21
	DDLParserRULE_indexType                  = 22
	DDLParserRULE_indexColumnNames           = 23
	DDLParserRULE_tableOption                = 24
)

// IRootContext is an interface to support dynamic dispatch.
type IRootContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRootContext differentiates from other interfaces.
	IsRootContext()
}

type RootContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRootContext() *RootContext {
	var p = new(RootContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DDLParserRULE_root
	return p
}

func (*RootContext) IsRootContext() {}

func NewRootContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RootContext {
	var p = new(RootContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DDLParserRULE_root

	return p
}

func (s *RootContext) GetParser() antlr.Parser { return s.parser }

func (s *RootContext) AllCreateTableDDL() []ICreateTableDDLContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICreateTableDDLContext)(nil)).Elem())
	var tst = make([]ICreateTableDDLContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICreateTableDDLContext)
		}
	}

	return tst
}

func (s *RootContext) CreateTableDDL(i int) ICreateTableDDLContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICreateTableDDLContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICreateTableDDLContext)
}

func (s *RootContext) EOF() antlr.TerminalNode {
	return s.GetToken(DDLParserEOF, 0)
}

func (s *RootContext) AllSemicolon() []antlr.TerminalNode {
	return s.GetTokens(DDLParserSemicolon)
}

func (s *RootContext) Semicolon(i int) antlr.TerminalNode {
	return s.GetToken(DDLParserSemicolon, i)
}

func (s *RootContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RootContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RootContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DDLParserListener); ok {
		listenerT.EnterRoot(s)
	}
}

func (s *RootContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DDLParserListener); ok {
		listenerT.ExitRoot(s)
	}
}

func (p *DDLParser) Root() (localctx IRootContext) {
	localctx = NewRootContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DDLParserRULE_root)
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
		p.SetState(50)
		p.CreateTableDDL()
	}
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(51)
				p.Match(DDLParserSemicolon)
			}
			{
				p.SetState(52)
				p.CreateTableDDL()
			}

		}
		p.SetState(57)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
	}
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DDLParserSemicolon {
		{
			p.SetState(58)
			p.Match(DDLParserSemicolon)
		}

	}
	{
		p.SetState(61)
		p.Match(DDLParserEOF)
	}

	return localctx
}

// ICreateTableDDLContext is an interface to support dynamic dispatch.
type ICreateTableDDLContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCreateTableDDLContext differentiates from other interfaces.
	IsCreateTableDDLContext()
}

type CreateTableDDLContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCreateTableDDLContext() *CreateTableDDLContext {
	var p = new(CreateTableDDLContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DDLParserRULE_createTableDDL
	return p
}

func (*CreateTableDDLContext) IsCreateTableDDLContext() {}

func NewCreateTableDDLContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CreateTableDDLContext {
	var p = new(CreateTableDDLContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DDLParserRULE_createTableDDL

	return p
}

func (s *CreateTableDDLContext) GetParser() antlr.Parser { return s.parser }

func (s *CreateTableDDLContext) CREATE() antlr.TerminalNode {
	return s.GetToken(DDLParserCREATE, 0)
}

func (s *CreateTableDDLContext) TABLE() antlr.TerminalNode {
	return s.GetToken(DDLParserTABLE, 0)
}

func (s *CreateTableDDLContext) TableName() ITableNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITableNameContext)
}

func (s *CreateTableDDLContext) CreateDefinitions() ICreateDefinitionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICreateDefinitionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICreateDefinitionsContext)
}

func (s *CreateTableDDLContext) TEMPORARY() antlr.TerminalNode {
	return s.GetToken(DDLParserTEMPORARY, 0)
}

func (s *CreateTableDDLContext) IfNotExists() IIfNotExistsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfNotExistsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfNotExistsContext)
}

func (s *CreateTableDDLContext) AllTableOption() []ITableOptionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITableOptionContext)(nil)).Elem())
	var tst = make([]ITableOptionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITableOptionContext)
		}
	}

	return tst
}

func (s *CreateTableDDLContext) TableOption(i int) ITableOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableOptionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITableOptionContext)
}

func (s *CreateTableDDLContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(DDLParserComma)
}

func (s *CreateTableDDLContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(DDLParserComma, i)
}

func (s *CreateTableDDLContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CreateTableDDLContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CreateTableDDLContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DDLParserListener); ok {
		listenerT.EnterCreateTableDDL(s)
	}
}

func (s *CreateTableDDLContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DDLParserListener); ok {
		listenerT.ExitCreateTableDDL(s)
	}
}

func (p *DDLParser) CreateTableDDL() (localctx ICreateTableDDLContext) {
	localctx = NewCreateTableDDLContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DDLParserRULE_createTableDDL)
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
		p.SetState(63)
		p.Match(DDLParserCREATE)
	}
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DDLParserTEMPORARY {
		{
			p.SetState(64)
			p.Match(DDLParserTEMPORARY)
		}

	}
	{
		p.SetState(67)
		p.Match(DDLParserTABLE)
	}
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DDLParserIF {
		{
			p.SetState(68)
			p.IfNotExists()
		}

	}
	{
		p.SetState(71)
		p.TableName()
	}
	{
		p.SetState(72)
		p.CreateDefinitions()
	}
	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DDLParserAUTO_INCREMENT)|(1<<DDLParserCHARACTER)|(1<<DDLParserCHARSET)|(1<<DDLParserCOLLATE)|(1<<DDLParserCOMMENT)|(1<<DDLParserDEFAULT)|(1<<DDLParserENGINE))) != 0 {
		{
			p.SetState(73)
			p.TableOption()
		}
		p.SetState(80)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DDLParserAUTO_INCREMENT)|(1<<DDLParserCHARACTER)|(1<<DDLParserCHARSET)|(1<<DDLParserCOLLATE)|(1<<DDLParserCOMMENT)|(1<<DDLParserDEFAULT)|(1<<DDLParserENGINE))) != 0) || _la == DDLParserComma {
			p.SetState(75)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == DDLParserComma {
				{
					p.SetState(74)
					p.Match(DDLParserComma)
				}

			}
			{
				p.SetState(77)
				p.TableOption()
			}

			p.SetState(82)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}

	return localctx
}

// ITableNameContext is an interface to support dynamic dispatch.
type ITableNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTableNameContext differentiates from other interfaces.
	IsTableNameContext()
}

type TableNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableNameContext() *TableNameContext {
	var p = new(TableNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DDLParserRULE_tableName
	return p
}

func (*TableNameContext) IsTableNameContext() {}

func NewTableNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableNameContext {
	var p = new(TableNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DDLParserRULE_tableName

	return p
}

func (s *TableNameContext) GetParser() antlr.Parser { return s.parser }

func (s *TableNameContext) BackQuotedString() antlr.TerminalNode {
	return s.GetToken(DDLParserBackQuotedString, 0)
}

func (s *TableNameContext) RawString() antlr.TerminalNode {
	return s.GetToken(DDLParserRawString, 0)
}

func (s *TableNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TableNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DDLParserListener); ok {
		listenerT.EnterTableName(s)
	}
}

func (s *TableNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DDLParserListener); ok {
		listenerT.ExitTableName(s)
	}
}

func (p *DDLParser) TableName() (localctx ITableNameContext) {
	localctx = NewTableNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, DDLParserRULE_tableName)
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
		p.SetState(85)
		_la = p.GetTokenStream().LA(1)

		if !(_la == DDLParserBackQuotedString || _la == DDLParserRawString) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IIfNotExistsContext is an interface to support dynamic dispatch.
type IIfNotExistsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfNotExistsContext differentiates from other interfaces.
	IsIfNotExistsContext()
}

type IfNotExistsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfNotExistsContext() *IfNotExistsContext {
	var p = new(IfNotExistsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DDLParserRULE_ifNotExists
	return p
}

func (*IfNotExistsContext) IsIfNotExistsContext() {}

func NewIfNotExistsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfNotExistsContext {
	var p = new(IfNotExistsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DDLParserRULE_ifNotExists

	return p
}

func (s *IfNotExistsContext) GetParser() antlr.Parser { return s.parser }

func (s *IfNotExistsContext) IF() antlr.TerminalNode {
	return s.GetToken(DDLParserIF, 0)
}

func (s *IfNotExistsContext) NOT() antlr.TerminalNode {
	return s.GetToken(DDLParserNOT, 0)
}

func (s *IfNotExistsContext) EXISTS() antlr.TerminalNode {
	return s.GetToken(DDLParserEXISTS, 0)
}

func (s *IfNotExistsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfNotExistsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfNotExistsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DDLParserListener); ok {
		listenerT.EnterIfNotExists(s)
	}
}

func (s *IfNotExistsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DDLParserListener); ok {
		listenerT.ExitIfNotExists(s)
	}
}

func (p *DDLParser) IfNotExists() (localctx IIfNotExistsContext) {
	localctx = NewIfNotExistsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, DDLParserRULE_ifNotExists)

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
		p.SetState(87)
		p.Match(DDLParserIF)
	}
	{
		p.SetState(88)
		p.Match(DDLParserNOT)
	}
	{
		p.SetState(89)
		p.Match(DDLParserEXISTS)
	}

	return localctx
}

// ICreateDefinitionsContext is an interface to support dynamic dispatch.
type ICreateDefinitionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCreateDefinitionsContext differentiates from other interfaces.
	IsCreateDefinitionsContext()
}

type CreateDefinitionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCreateDefinitionsContext() *CreateDefinitionsContext {
	var p = new(CreateDefinitionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DDLParserRULE_createDefinitions
	return p
}

func (*CreateDefinitionsContext) IsCreateDefinitionsContext() {}

func NewCreateDefinitionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CreateDefinitionsContext {
	var p = new(CreateDefinitionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DDLParserRULE_createDefinitions

	return p
}

func (s *CreateDefinitionsContext) GetParser() antlr.Parser { return s.parser }

func (s *CreateDefinitionsContext) LeftParenthesis() antlr.TerminalNode {
	return s.GetToken(DDLParserLeftParenthesis, 0)
}

func (s *CreateDefinitionsContext) AllCreateDefinition() []ICreateDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICreateDefinitionContext)(nil)).Elem())
	var tst = make([]ICreateDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICreateDefinitionContext)
		}
	}

	return tst
}

func (s *CreateDefinitionsContext) CreateDefinition(i int) ICreateDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICreateDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICreateDefinitionContext)
}

func (s *CreateDefinitionsContext) RightParenthesis() antlr.TerminalNode {
	return s.GetToken(DDLParserRightParenthesis, 0)
}

func (s *CreateDefinitionsContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(DDLParserComma)
}

func (s *CreateDefinitionsContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(DDLParserComma, i)
}

func (s *CreateDefinitionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CreateDefinitionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CreateDefinitionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DDLParserListener); ok {
		listenerT.EnterCreateDefinitions(s)
	}
}

func (s *CreateDefinitionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DDLParserListener); ok {
		listenerT.ExitCreateDefinitions(s)
	}
}

func (p *DDLParser) CreateDefinitions() (localctx ICreateDefinitionsContext) {
	localctx = NewCreateDefinitionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, DDLParserRULE_createDefinitions)
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
		p.SetState(91)
		p.Match(DDLParserLeftParenthesis)
	}
	{
		p.SetState(92)
		p.CreateDefinition()
	}
	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DDLParserComma {
		{
			p.SetState(93)
			p.Match(DDLParserComma)
		}
		{
			p.SetState(94)
			p.CreateDefinition()
		}

		p.SetState(99)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(100)
		p.Match(DDLParserRightParenthesis)
	}

	return localctx
}

// ICreateDefinitionContext is an interface to support dynamic dispatch.
type ICreateDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetField returns the field token.
	GetField() antlr.Token

	// SetField sets the field token.
	SetField(antlr.Token)

	// IsCreateDefinitionContext differentiates from other interfaces.
	IsCreateDefinitionContext()
}

type CreateDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	field  antlr.Token
}

func NewEmptyCreateDefinitionContext() *CreateDefinitionContext {
	var p = new(CreateDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DDLParserRULE_createDefinition
	return p
}

func (*CreateDefinitionContext) IsCreateDefinitionContext() {}

func NewCreateDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CreateDefinitionContext {
	var p = new(CreateDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DDLParserRULE_createDefinition

	return p
}

func (s *CreateDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *CreateDefinitionContext) GetField() antlr.Token { return s.field }

func (s *CreateDefinitionContext) SetField(v antlr.Token) { s.field = v }

func (s *CreateDefinitionContext) ColumnDefinition() IColumnDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumnDefinitionContext)
}

func (s *CreateDefinitionContext) BackQuotedString() antlr.TerminalNode {
	return s.GetToken(DDLParserBackQuotedString, 0)
}

func (s *CreateDefinitionContext) RawString() antlr.TerminalNode {
	return s.GetToken(DDLParserRawString, 0)
}

func (s *CreateDefinitionContext) TableConstraint() ITableConstraintContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableConstraintContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITableConstraintContext)
}

func (s *CreateDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CreateDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CreateDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DDLParserListener); ok {
		listenerT.EnterCreateDefinition(s)
	}
}

func (s *CreateDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DDLParserListener); ok {
		listenerT.ExitCreateDefinition(s)
	}
}

func (p *DDLParser) CreateDefinition() (localctx ICreateDefinitionContext) {
	localctx = NewCreateDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, DDLParserRULE_createDefinition)
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

	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(102)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*CreateDefinitionContext).field = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == DDLParserBackQuotedString || _la == DDLParserRawString) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*CreateDefinitionContext).field = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(103)
			p.ColumnDefinition()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(104)
			p.TableConstraint()
		}

	}

	return localctx
}

// IColumnDefinitionContext is an interface to support dynamic dispatch.
type IColumnDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsColumnDefinitionContext differentiates from other interfaces.
	IsColumnDefinitionContext()
}

type ColumnDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumnDefinitionContext() *ColumnDefinitionContext {
	var p = new(ColumnDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DDLParserRULE_columnDefinition
	return p
}

func (*ColumnDefinitionContext) IsColumnDefinitionContext() {}

func NewColumnDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnDefinitionContext {
	var p = new(ColumnDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DDLParserRULE_columnDefinition

	return p
}

func (s *ColumnDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumnDefinitionContext) DataType() IDataTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDataTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDataTypeContext)
}

func (s *ColumnDefinitionContext) AllColumnConstraint() []IColumnConstraintContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IColumnConstraintContext)(nil)).Elem())
	var tst = make([]IColumnConstraintContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IColumnConstraintContext)
		}
	}

	return tst
}

func (s *ColumnDefinitionContext) ColumnConstraint(i int) IColumnConstraintContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnConstraintContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IColumnConstraintContext)
}

func (s *ColumnDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColumnDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DDLParserListener); ok {
		listenerT.EnterColumnDefinition(s)
	}
}

func (s *ColumnDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DDLParserListener); ok {
		listenerT.ExitColumnDefinition(s)
	}
}

func (p *DDLParser) ColumnDefinition() (localctx IColumnDefinitionContext) {
	localctx = NewColumnDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, DDLParserRULE_columnDefinition)
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
		p.SetState(107)
		p.DataType()
	}
	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DDLParserAUTO_INCREMENT)|(1<<DDLParserCOLLATE)|(1<<DDLParserCOMMENT)|(1<<DDLParserDEFAULT))) != 0) || (((_la-36)&-(0x1f+1)) == 0 && ((1<<uint((_la-36)))&((1<<(DDLParserKEY-36))|(1<<(DDLParserNOT-36))|(1<<(DDLParserNULL-36))|(1<<(DDLParserON-36))|(1<<(DDLParserPRIMARY-36))|(1<<(DDLParserSERIAL-36)))) != 0) || _la == DDLParserUNIQUE {
		{
			p.SetState(108)
			p.ColumnConstraint()
		}

		p.SetState(113)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDataTypeContext is an interface to support dynamic dispatch.
type IDataTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTypeName returns the typeName token.
	GetTypeName() antlr.Token

	// GetSign returns the sign token.
	GetSign() antlr.Token

	// SetTypeName sets the typeName token.
	SetTypeName(antlr.Token)

	// SetSign sets the sign token.
	SetSign(antlr.Token)

	// IsDataTypeContext differentiates from other interfaces.
	IsDataTypeContext()
}

type DataTypeContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	typeName antlr.Token
	sign     antlr.Token
}

func NewEmptyDataTypeContext() *DataTypeContext {
	var p = new(DataTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DDLParserRULE_dataType
	return p
}

func (*DataTypeContext) IsDataTypeContext() {}

func NewDataTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DataTypeContext {
	var p = new(DataTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DDLParserRULE_dataType

	return p
}

func (s *DataTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *DataTypeContext) GetTypeName() antlr.Token { return s.typeName }

func (s *DataTypeContext) GetSign() antlr.Token { return s.sign }

func (s *DataTypeContext) SetTypeName(v antlr.Token) { s.typeName = v }

func (s *DataTypeContext) SetSign(v antlr.Token) { s.sign = v }

func (s *DataTypeContext) CHAR() antlr.TerminalNode {
	return s.GetToken(DDLParserCHAR, 0)
}

func (s *DataTypeContext) VARCHAR() antlr.TerminalNode {
	return s.GetToken(DDLParserVARCHAR, 0)
}

func (s *DataTypeContext) TINYTEXT() antlr.TerminalNode {
	return s.GetToken(DDLParserTINYTEXT, 0)
}

func (s *DataTypeContext) TEXT() antlr.TerminalNode {
	return s.GetToken(DDLParserTEXT, 0)
}

func (s *DataTypeContext) MEDIUMTEXT() antlr.TerminalNode {
	return s.GetToken(DDLParserMEDIUMTEXT, 0)
}

func (s *DataTypeContext) LONGTEXT() antlr.TerminalNode {
	return s.GetToken(DDLParserLONGTEXT, 0)
}

func (s *DataTypeContext) NCHAR() antlr.TerminalNode {
	return s.GetToken(DDLParserNCHAR, 0)
}

func (s *DataTypeContext) NVARCHAR() antlr.TerminalNode {
	return s.GetToken(DDLParserNVARCHAR, 0)
}

func (s *DataTypeContext) LengthOneDimension() ILengthOneDimensionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILengthOneDimensionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILengthOneDimensionContext)
}

func (s *DataTypeContext) AllBINARY() []antlr.TerminalNode {
	return s.GetTokens(DDLParserBINARY)
}

func (s *DataTypeContext) BINARY(i int) antlr.TerminalNode {
	return s.GetToken(DDLParserBINARY, i)
}

func (s *DataTypeContext) COLLATE() antlr.TerminalNode {
	return s.GetToken(DDLParserCOLLATE, 0)
}

func (s *DataTypeContext) AllRawString() []antlr.TerminalNode {
	return s.GetTokens(DDLParserRawString)
}

func (s *DataTypeContext) RawString(i int) antlr.TerminalNode {
	return s.GetToken(DDLParserRawString, i)
}

func (s *DataTypeContext) AllBackQuotedString() []antlr.TerminalNode {
	return s.GetTokens(DDLParserBackQuotedString)
}

func (s *DataTypeContext) BackQuotedString(i int) antlr.TerminalNode {
	return s.GetToken(DDLParserBackQuotedString, i)
}

func (s *DataTypeContext) CHARACTER() antlr.TerminalNode {
	return s.GetToken(DDLParserCHARACTER, 0)
}

func (s *DataTypeContext) AllSET() []antlr.TerminalNode {
	return s.GetTokens(DDLParserSET)
}

func (s *DataTypeContext) SET(i int) antlr.TerminalNode {
	return s.GetToken(DDLParserSET, i)
}

func (s *DataTypeContext) CHARSET() antlr.TerminalNode {
	return s.GetToken(DDLParserCHARSET, 0)
}

func (s *DataTypeContext) NATIONAL() antlr.TerminalNode {
	return s.GetToken(DDLParserNATIONAL, 0)
}

func (s *DataTypeContext) VARYING() antlr.TerminalNode {
	return s.GetToken(DDLParserVARYING, 0)
}

func (s *DataTypeContext) TINYINT() antlr.TerminalNode {
	return s.GetToken(DDLParserTINYINT, 0)
}

func (s *DataTypeContext) SMALLINT() antlr.TerminalNode {
	return s.GetToken(DDLParserSMALLINT, 0)
}

func (s *DataTypeContext) MEDIUMINT() antlr.TerminalNode {
	return s.GetToken(DDLParserMEDIUMINT, 0)
}

func (s *DataTypeContext) INT() antlr.TerminalNode {
	return s.GetToken(DDLParserINT, 0)
}

func (s *DataTypeContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(DDLParserINTEGER, 0)
}

func (s *DataTypeContext) BIGINT() antlr.TerminalNode {
	return s.GetToken(DDLParserBIGINT, 0)
}

func (s *DataTypeContext) ZEROFILL() antlr.TerminalNode {
	return s.GetToken(DDLParserZEROFILL, 0)
}

func (s *DataTypeContext) SIGNED() antlr.TerminalNode {
	return s.GetToken(DDLParserSIGNED, 0)
}

func (s *DataTypeContext) UNSIGNED() antlr.TerminalNode {
	return s.GetToken(DDLParserUNSIGNED, 0)
}

func (s *DataTypeContext) REAL() antlr.TerminalNode {
	return s.GetToken(DDLParserREAL, 0)
}

func (s *DataTypeContext) LengthTwoDimension() ILengthTwoDimensionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILengthTwoDimensionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILengthTwoDimensionContext)
}

func (s *DataTypeContext) DOUBLE() antlr.TerminalNode {
	return s.GetToken(DDLParserDOUBLE, 0)
}

func (s *DataTypeContext) PRECISION() antlr.TerminalNode {
	return s.GetToken(DDLParserPRECISION, 0)
}

func (s *DataTypeContext) DECIMAL() antlr.TerminalNode {
	return s.GetToken(DDLParserDECIMAL, 0)
}

func (s *DataTypeContext) DEC() antlr.TerminalNode {
	return s.GetToken(DDLParserDEC, 0)
}

func (s *DataTypeContext) FIXED() antlr.TerminalNode {
	return s.GetToken(DDLParserFIXED, 0)
}

func (s *DataTypeContext) NUMERIC() antlr.TerminalNode {
	return s.GetToken(DDLParserNUMERIC, 0)
}

func (s *DataTypeContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(DDLParserFLOAT, 0)
}

func (s *DataTypeContext) LengthTwoOptionalDimension() ILengthTwoOptionalDimensionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILengthTwoOptionalDimensionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILengthTwoOptionalDimensionContext)
}

func (s *DataTypeContext) DATE() antlr.TerminalNode {
	return s.GetToken(DDLParserDATE, 0)
}

func (s *DataTypeContext) TINYBLOB() antlr.TerminalNode {
	return s.GetToken(DDLParserTINYBLOB, 0)
}

func (s *DataTypeContext) BLOB() antlr.TerminalNode {
	return s.GetToken(DDLParserBLOB, 0)
}

func (s *DataTypeContext) MEDIUMBLOB() antlr.TerminalNode {
	return s.GetToken(DDLParserMEDIUMBLOB, 0)
}

func (s *DataTypeContext) LONGBLOB() antlr.TerminalNode {
	return s.GetToken(DDLParserLONGBLOB, 0)
}

func (s *DataTypeContext) BOOL() antlr.TerminalNode {
	return s.GetToken(DDLParserBOOL, 0)
}

func (s *DataTypeContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(DDLParserBOOLEAN, 0)
}

func (s *DataTypeContext) SERIAL() antlr.TerminalNode {
	return s.GetToken(DDLParserSERIAL, 0)
}

func (s *DataTypeContext) JSON() antlr.TerminalNode {
	return s.GetToken(DDLParserJSON, 0)
}

func (s *DataTypeContext) BIT() antlr.TerminalNode {
	return s.GetToken(DDLParserBIT, 0)
}

func (s *DataTypeContext) TIME() antlr.TerminalNode {
	return s.GetToken(DDLParserTIME, 0)
}

func (s *DataTypeContext) TIMESTAMP() antlr.TerminalNode {
	return s.GetToken(DDLParserTIMESTAMP, 0)
}

func (s *DataTypeContext) DATETIME() antlr.TerminalNode {
	return s.GetToken(DDLParserDATETIME, 0)
}

func (s *DataTypeContext) VARBINARY() antlr.TerminalNode {
	return s.GetToken(DDLParserVARBINARY, 0)
}

func (s *DataTypeContext) YEAR() antlr.TerminalNode {
	return s.GetToken(DDLParserYEAR, 0)
}

func (s *DataTypeContext) CollectionOptions() ICollectionOptionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICollectionOptionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICollectionOptionsContext)
}

func (s *DataTypeContext) ENUM() antlr.TerminalNode {
	return s.GetToken(DDLParserENUM, 0)
}

func (s *DataTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DataTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DDLParserListener); ok {
		listenerT.EnterDataType(s)
	}
}

func (s *DataTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DDLParserListener); ok {
		listenerT.ExitDataType(s)
	}
}

func (p *DDLParser) DataType() (localctx IDataTypeContext) {
	localctx = NewDataTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, DDLParserRULE_dataType)
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

	p.SetState(219)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(114)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*DataTypeContext).typeName = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == DDLParserCHAR || (((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(DDLParserLONGTEXT-40))|(1<<(DDLParserMEDIUMTEXT-40))|(1<<(DDLParserNCHAR-40))|(1<<(DDLParserNVARCHAR-40))|(1<<(DDLParserTEXT-40))|(1<<(DDLParserTINYTEXT-40)))) != 0) || _la == DDLParserVARCHAR) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*DataTypeContext).typeName = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		p.SetState(116)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DDLParserLeftParenthesis {
			{
				p.SetState(115)
				p.LengthOneDimension()
			}

		}
		p.SetState(119)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DDLParserBINARY {
			{
				p.SetState(118)
				p.Match(DDLParserBINARY)
			}

		}
		p.SetState(127)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DDLParserCHARACTER || _la == DDLParserCHARSET {
			p.SetState(124)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case DDLParserCHARACTER:
				{
					p.SetState(121)
					p.Match(DDLParserCHARACTER)
				}
				{
					p.SetState(122)
					p.Match(DDLParserSET)
				}

			case DDLParserCHARSET:
				{
					p.SetState(123)
					p.Match(DDLParserCHARSET)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}
			{
				p.SetState(126)
				_la = p.GetTokenStream().LA(1)

				if !(_la == DDLParserBINARY || _la == DDLParserBackQuotedString || _la == DDLParserRawString) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		p.SetState(131)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(129)
				p.Match(DDLParserCOLLATE)
			}
			{
				p.SetState(130)
				_la = p.GetTokenStream().LA(1)

				if !(_la == DDLParserBackQuotedString || _la == DDLParserRawString) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(133)
			p.Match(DDLParserNATIONAL)
		}
		{
			p.SetState(134)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*DataTypeContext).typeName = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == DDLParserCHARACTER || _la == DDLParserVARCHAR) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*DataTypeContext).typeName = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		p.SetState(136)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DDLParserLeftParenthesis {
			{
				p.SetState(135)
				p.LengthOneDimension()
			}

		}
		p.SetState(139)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DDLParserBINARY {
			{
				p.SetState(138)
				p.Match(DDLParserBINARY)
			}

		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(141)
			p.Match(DDLParserNCHAR)
		}
		{
			p.SetState(142)

			var _m = p.Match(DDLParserVARCHAR)

			localctx.(*DataTypeContext).typeName = _m
		}
		p.SetState(144)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DDLParserLeftParenthesis {
			{
				p.SetState(143)
				p.LengthOneDimension()
			}

		}
		p.SetState(147)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DDLParserBINARY {
			{
				p.SetState(146)
				p.Match(DDLParserBINARY)
			}

		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(149)
			p.Match(DDLParserNATIONAL)
		}
		{
			p.SetState(150)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*DataTypeContext).typeName = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == DDLParserCHARACTER || _la == DDLParserCHAR) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*DataTypeContext).typeName = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(151)
			p.Match(DDLParserVARYING)
		}
		p.SetState(153)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DDLParserLeftParenthesis {
			{
				p.SetState(152)
				p.LengthOneDimension()
			}

		}
		p.SetState(156)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DDLParserBINARY {
			{
				p.SetState(155)
				p.Match(DDLParserBINARY)
			}

		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(158)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*DataTypeContext).typeName = _lt

			_la = p.GetTokenStream().LA(1)

			if !((((_la-2)&-(0x1f+1)) == 0 && ((1<<uint((_la-2)))&((1<<(DDLParserBIGINT-2))|(1<<(DDLParserINTEGER-2))|(1<<(DDLParserINT-2)))) != 0) || (((_la-42)&-(0x1f+1)) == 0 && ((1<<uint((_la-42)))&((1<<(DDLParserMEDIUMINT-42))|(1<<(DDLParserSMALLINT-42))|(1<<(DDLParserTINYINT-42)))) != 0)) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*DataTypeContext).typeName = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		p.SetState(160)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DDLParserLeftParenthesis {
			{
				p.SetState(159)
				p.LengthOneDimension()
			}

		}
		p.SetState(163)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DDLParserSIGNED || _la == DDLParserUNSIGNED {
			{
				p.SetState(162)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*DataTypeContext).sign = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == DDLParserSIGNED || _la == DDLParserUNSIGNED) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*DataTypeContext).sign = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		p.SetState(166)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DDLParserZEROFILL {
			{
				p.SetState(165)
				p.Match(DDLParserZEROFILL)
			}

		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(168)

			var _m = p.Match(DDLParserREAL)

			localctx.(*DataTypeContext).typeName = _m
		}
		p.SetState(170)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DDLParserLeftParenthesis {
			{
				p.SetState(169)
				p.LengthTwoDimension()
			}

		}
		p.SetState(173)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DDLParserSIGNED || _la == DDLParserUNSIGNED {
			{
				p.SetState(172)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*DataTypeContext).sign = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == DDLParserSIGNED || _la == DDLParserUNSIGNED) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*DataTypeContext).sign = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		p.SetState(176)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DDLParserZEROFILL {
			{
				p.SetState(175)
				p.Match(DDLParserZEROFILL)
			}

		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(178)

			var _m = p.Match(DDLParserDOUBLE)

			localctx.(*DataTypeContext).typeName = _m
		}
		p.SetState(180)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DDLParserPRECISION {
			{
				p.SetState(179)
				p.Match(DDLParserPRECISION)
			}

		}
		p.SetState(183)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DDLParserLeftParenthesis {
			{
				p.SetState(182)
				p.LengthTwoDimension()
			}

		}
		p.SetState(186)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DDLParserSIGNED || _la == DDLParserUNSIGNED {
			{
				p.SetState(185)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*DataTypeContext).sign = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == DDLParserSIGNED || _la == DDLParserUNSIGNED) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*DataTypeContext).sign = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		p.SetState(189)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DDLParserZEROFILL {
			{
				p.SetState(188)
				p.Match(DDLParserZEROFILL)
			}

		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(191)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*DataTypeContext).typeName = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la-19)&-(0x1f+1)) == 0 && ((1<<uint((_la-19)))&((1<<(DDLParserDECIMAL-19))|(1<<(DDLParserDEC-19))|(1<<(DDLParserFIXED-19))|(1<<(DDLParserFLOAT-19))|(1<<(DDLParserNUMERIC-19)))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*DataTypeContext).typeName = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		p.SetState(193)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DDLParserLeftParenthesis {
			{
				p.SetState(192)
				p.LengthTwoOptionalDimension()
			}

		}
		p.SetState(196)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DDLParserSIGNED || _la == DDLParserUNSIGNED {
			{
				p.SetState(195)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*DataTypeContext).sign = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == DDLParserSIGNED || _la == DDLParserUNSIGNED) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*DataTypeContext).sign = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		p.SetState(199)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DDLParserZEROFILL {
			{
				p.SetState(198)
				p.Match(DDLParserZEROFILL)
			}

		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(201)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*DataTypeContext).typeName = _lt

			_la = p.GetTokenStream().LA(1)

			if !((((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DDLParserBLOB)|(1<<DDLParserBOOLEAN)|(1<<DDLParserBOOL)|(1<<DDLParserDATE))) != 0) || (((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(DDLParserJSON-34))|(1<<(DDLParserLONGBLOB-34))|(1<<(DDLParserMEDIUMBLOB-34))|(1<<(DDLParserSERIAL-34))|(1<<(DDLParserTINYBLOB-34)))) != 0)) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*DataTypeContext).typeName = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(202)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*DataTypeContext).typeName = _lt

			_la = p.GetTokenStream().LA(1)

			if !((((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DDLParserBINARY)|(1<<DDLParserBIT)|(1<<DDLParserDATETIME))) != 0) || (((_la-63)&-(0x1f+1)) == 0 && ((1<<uint((_la-63)))&((1<<(DDLParserTIMESTAMP-63))|(1<<(DDLParserTIME-63))|(1<<(DDLParserVARBINARY-63))|(1<<(DDLParserYEAR-63)))) != 0)) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*DataTypeContext).typeName = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		p.SetState(204)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DDLParserLeftParenthesis {
			{
				p.SetState(203)
				p.LengthOneDimension()
			}

		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(206)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*DataTypeContext).typeName = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == DDLParserENUM || _la == DDLParserSET) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*DataTypeContext).typeName = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(207)
			p.CollectionOptions()
		}
		p.SetState(209)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DDLParserBINARY {
			{
				p.SetState(208)
				p.Match(DDLParserBINARY)
			}

		}
		p.SetState(217)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DDLParserCHARACTER || _la == DDLParserCHARSET {
			p.SetState(214)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case DDLParserCHARACTER:
				{
					p.SetState(211)
					p.Match(DDLParserCHARACTER)
				}
				{
					p.SetState(212)
					p.Match(DDLParserSET)
				}

			case DDLParserCHARSET:
				{
					p.SetState(213)
					p.Match(DDLParserCHARSET)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}
			{
				p.SetState(216)
				_la = p.GetTokenStream().LA(1)

				if !(_la == DDLParserBINARY || _la == DDLParserBackQuotedString || _la == DDLParserRawString) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}

	}

	return localctx
}

// ILengthOneDimensionContext is an interface to support dynamic dispatch.
type ILengthOneDimensionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLengthOneDimensionContext differentiates from other interfaces.
	IsLengthOneDimensionContext()
}

type LengthOneDimensionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLengthOneDimensionContext() *LengthOneDimensionContext {
	var p = new(LengthOneDimensionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DDLParserRULE_lengthOneDimension
	return p
}

func (*LengthOneDimensionContext) IsLengthOneDimensionContext() {}

func NewLengthOneDimensionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LengthOneDimensionContext {
	var p = new(LengthOneDimensionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DDLParserRULE_lengthOneDimension

	return p
}

func (s *LengthOneDimensionContext) GetParser() antlr.Parser { return s.parser }

func (s *LengthOneDimensionContext) LeftParenthesis() antlr.TerminalNode {
	return s.GetToken(DDLParserLeftParenthesis, 0)
}

func (s *LengthOneDimensionContext) Integer() antlr.TerminalNode {
	return s.GetToken(DDLParserInteger, 0)
}

func (s *LengthOneDimensionContext) RightParenthesis() antlr.TerminalNode {
	return s.GetToken(DDLParserRightParenthesis, 0)
}

func (s *LengthOneDimensionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LengthOneDimensionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LengthOneDimensionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DDLParserListener); ok {
		listenerT.EnterLengthOneDimension(s)
	}
}

func (s *LengthOneDimensionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DDLParserListener); ok {
		listenerT.ExitLengthOneDimension(s)
	}
}

func (p *DDLParser) LengthOneDimension() (localctx ILengthOneDimensionContext) {
	localctx = NewLengthOneDimensionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, DDLParserRULE_lengthOneDimension)

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
		p.SetState(221)
		p.Match(DDLParserLeftParenthesis)
	}
	{
		p.SetState(222)
		p.Match(DDLParserInteger)
	}
	{
		p.SetState(223)
		p.Match(DDLParserRightParenthesis)
	}

	return localctx
}

// ILengthTwoDimensionContext is an interface to support dynamic dispatch.
type ILengthTwoDimensionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLengthTwoDimensionContext differentiates from other interfaces.
	IsLengthTwoDimensionContext()
}

type LengthTwoDimensionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLengthTwoDimensionContext() *LengthTwoDimensionContext {
	var p = new(LengthTwoDimensionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DDLParserRULE_lengthTwoDimension
	return p
}

func (*LengthTwoDimensionContext) IsLengthTwoDimensionContext() {}

func NewLengthTwoDimensionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LengthTwoDimensionContext {
	var p = new(LengthTwoDimensionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DDLParserRULE_lengthTwoDimension

	return p
}

func (s *LengthTwoDimensionContext) GetParser() antlr.Parser { return s.parser }

func (s *LengthTwoDimensionContext) LeftParenthesis() antlr.TerminalNode {
	return s.GetToken(DDLParserLeftParenthesis, 0)
}

func (s *LengthTwoDimensionContext) AllInteger() []antlr.TerminalNode {
	return s.GetTokens(DDLParserInteger)
}

func (s *LengthTwoDimensionContext) Integer(i int) antlr.TerminalNode {
	return s.GetToken(DDLParserInteger, i)
}

func (s *LengthTwoDimensionContext) Comma() antlr.TerminalNode {
	return s.GetToken(DDLParserComma, 0)
}

func (s *LengthTwoDimensionContext) RightParenthesis() antlr.TerminalNode {
	return s.GetToken(DDLParserRightParenthesis, 0)
}

func (s *LengthTwoDimensionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LengthTwoDimensionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LengthTwoDimensionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DDLParserListener); ok {
		listenerT.EnterLengthTwoDimension(s)
	}
}

func (s *LengthTwoDimensionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DDLParserListener); ok {
		listenerT.ExitLengthTwoDimension(s)
	}
}

func (p *DDLParser) LengthTwoDimension() (localctx ILengthTwoDimensionContext) {
	localctx = NewLengthTwoDimensionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, DDLParserRULE_lengthTwoDimension)

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
		p.SetState(225)
		p.Match(DDLParserLeftParenthesis)
	}
	{
		p.SetState(226)
		p.Match(DDLParserInteger)
	}
	{
		p.SetState(227)
		p.Match(DDLParserComma)
	}
	{
		p.SetState(228)
		p.Match(DDLParserInteger)
	}
	{
		p.SetState(229)
		p.Match(DDLParserRightParenthesis)
	}

	return localctx
}

// ILengthTwoOptionalDimensionContext is an interface to support dynamic dispatch.
type ILengthTwoOptionalDimensionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLengthTwoOptionalDimensionContext differentiates from other interfaces.
	IsLengthTwoOptionalDimensionContext()
}

type LengthTwoOptionalDimensionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLengthTwoOptionalDimensionContext() *LengthTwoOptionalDimensionContext {
	var p = new(LengthTwoOptionalDimensionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DDLParserRULE_lengthTwoOptionalDimension
	return p
}

func (*LengthTwoOptionalDimensionContext) IsLengthTwoOptionalDimensionContext() {}

func NewLengthTwoOptionalDimensionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LengthTwoOptionalDimensionContext {
	var p = new(LengthTwoOptionalDimensionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DDLParserRULE_lengthTwoOptionalDimension

	return p
}

func (s *LengthTwoOptionalDimensionContext) GetParser() antlr.Parser { return s.parser }

func (s *LengthTwoOptionalDimensionContext) LeftParenthesis() antlr.TerminalNode {
	return s.GetToken(DDLParserLeftParenthesis, 0)
}

func (s *LengthTwoOptionalDimensionContext) AllInteger() []antlr.TerminalNode {
	return s.GetTokens(DDLParserInteger)
}

func (s *LengthTwoOptionalDimensionContext) Integer(i int) antlr.TerminalNode {
	return s.GetToken(DDLParserInteger, i)
}

func (s *LengthTwoOptionalDimensionContext) RightParenthesis() antlr.TerminalNode {
	return s.GetToken(DDLParserRightParenthesis, 0)
}

func (s *LengthTwoOptionalDimensionContext) Comma() antlr.TerminalNode {
	return s.GetToken(DDLParserComma, 0)
}

func (s *LengthTwoOptionalDimensionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LengthTwoOptionalDimensionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LengthTwoOptionalDimensionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DDLParserListener); ok {
		listenerT.EnterLengthTwoOptionalDimension(s)
	}
}

func (s *LengthTwoOptionalDimensionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DDLParserListener); ok {
		listenerT.ExitLengthTwoOptionalDimension(s)
	}
}

func (p *DDLParser) LengthTwoOptionalDimension() (localctx ILengthTwoOptionalDimensionContext) {
	localctx = NewLengthTwoOptionalDimensionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, DDLParserRULE_lengthTwoOptionalDimension)
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
		p.SetState(231)
		p.Match(DDLParserLeftParenthesis)
	}
	{
		p.SetState(232)
		p.Match(DDLParserInteger)
	}
	p.SetState(235)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DDLParserComma {
		{
			p.SetState(233)
			p.Match(DDLParserComma)
		}
		{
			p.SetState(234)
			p.Match(DDLParserInteger)
		}

	}
	{
		p.SetState(237)
		p.Match(DDLParserRightParenthesis)
	}

	return localctx
}

// ICollectionOptionsContext is an interface to support dynamic dispatch.
type ICollectionOptionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCollectionOptionsContext differentiates from other interfaces.
	IsCollectionOptionsContext()
}

type CollectionOptionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCollectionOptionsContext() *CollectionOptionsContext {
	var p = new(CollectionOptionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DDLParserRULE_collectionOptions
	return p
}

func (*CollectionOptionsContext) IsCollectionOptionsContext() {}

func NewCollectionOptionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CollectionOptionsContext {
	var p = new(CollectionOptionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DDLParserRULE_collectionOptions

	return p
}

func (s *CollectionOptionsContext) GetParser() antlr.Parser { return s.parser }

func (s *CollectionOptionsContext) LeftParenthesis() antlr.TerminalNode {
	return s.GetToken(DDLParserLeftParenthesis, 0)
}

func (s *CollectionOptionsContext) AllLiteral() []antlr.TerminalNode {
	return s.GetTokens(DDLParserLiteral)
}

func (s *CollectionOptionsContext) Literal(i int) antlr.TerminalNode {
	return s.GetToken(DDLParserLiteral, i)
}

func (s *CollectionOptionsContext) RightParenthesis() antlr.TerminalNode {
	return s.GetToken(DDLParserRightParenthesis, 0)
}

func (s *CollectionOptionsContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(DDLParserComma)
}

func (s *CollectionOptionsContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(DDLParserComma, i)
}

func (s *CollectionOptionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CollectionOptionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CollectionOptionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DDLParserListener); ok {
		listenerT.EnterCollectionOptions(s)
	}
}

func (s *CollectionOptionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DDLParserListener); ok {
		listenerT.ExitCollectionOptions(s)
	}
}

func (p *DDLParser) CollectionOptions() (localctx ICollectionOptionsContext) {
	localctx = NewCollectionOptionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, DDLParserRULE_collectionOptions)
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
		p.SetState(239)
		p.Match(DDLParserLeftParenthesis)
	}
	{
		p.SetState(240)
		p.Match(DDLParserLiteral)
	}
	p.SetState(245)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DDLParserComma {
		{
			p.SetState(241)
			p.Match(DDLParserComma)
		}
		{
			p.SetState(242)
			p.Match(DDLParserLiteral)
		}

		p.SetState(247)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(248)
		p.Match(DDLParserRightParenthesis)
	}

	return localctx
}

// IColumnConstraintContext is an interface to support dynamic dispatch.
type IColumnConstraintContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsColumnConstraintContext differentiates from other interfaces.
	IsColumnConstraintContext()
}

type ColumnConstraintContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumnConstraintContext() *ColumnConstraintContext {
	var p = new(ColumnConstraintContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DDLParserRULE_columnConstraint
	return p
}

func (*ColumnConstraintContext) IsColumnConstraintContext() {}

func NewColumnConstraintContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnConstraintContext {
	var p = new(ColumnConstraintContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DDLParserRULE_columnConstraint

	return p
}

func (s *ColumnConstraintContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumnConstraintContext) NullNotNull() INullNotNullContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INullNotNullContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INullNotNullContext)
}

func (s *ColumnConstraintContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(DDLParserDEFAULT, 0)
}

func (s *ColumnConstraintContext) DefaultValue() IDefaultValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefaultValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefaultValueContext)
}

func (s *ColumnConstraintContext) AUTO_INCREMENT() antlr.TerminalNode {
	return s.GetToken(DDLParserAUTO_INCREMENT, 0)
}

func (s *ColumnConstraintContext) ON() antlr.TerminalNode {
	return s.GetToken(DDLParserON, 0)
}

func (s *ColumnConstraintContext) UPDATE() antlr.TerminalNode {
	return s.GetToken(DDLParserUPDATE, 0)
}

func (s *ColumnConstraintContext) CurrentTimestamp() ICurrentTimestampContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICurrentTimestampContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICurrentTimestampContext)
}

func (s *ColumnConstraintContext) PrimaryKey() IPrimaryKeyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryKeyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimaryKeyContext)
}

func (s *ColumnConstraintContext) KEY() antlr.TerminalNode {
	return s.GetToken(DDLParserKEY, 0)
}

func (s *ColumnConstraintContext) UNIQUE() antlr.TerminalNode {
	return s.GetToken(DDLParserUNIQUE, 0)
}

func (s *ColumnConstraintContext) Comment() ICommentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommentContext)
}

func (s *ColumnConstraintContext) COLLATE() antlr.TerminalNode {
	return s.GetToken(DDLParserCOLLATE, 0)
}

func (s *ColumnConstraintContext) BackQuotedString() antlr.TerminalNode {
	return s.GetToken(DDLParserBackQuotedString, 0)
}

func (s *ColumnConstraintContext) SERIAL() antlr.TerminalNode {
	return s.GetToken(DDLParserSERIAL, 0)
}

func (s *ColumnConstraintContext) VALUE() antlr.TerminalNode {
	return s.GetToken(DDLParserVALUE, 0)
}

func (s *ColumnConstraintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnConstraintContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColumnConstraintContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DDLParserListener); ok {
		listenerT.EnterColumnConstraint(s)
	}
}

func (s *ColumnConstraintContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DDLParserListener); ok {
		listenerT.ExitColumnConstraint(s)
	}
}

func (p *DDLParser) ColumnConstraint() (localctx IColumnConstraintContext) {
	localctx = NewColumnConstraintContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, DDLParserRULE_columnConstraint)

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

	p.SetState(271)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DDLParserNOT, DDLParserNULL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(250)
			p.NullNotNull()
		}

	case DDLParserDEFAULT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(251)
			p.Match(DDLParserDEFAULT)
		}
		{
			p.SetState(252)
			p.DefaultValue()
		}

	case DDLParserAUTO_INCREMENT, DDLParserON:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(257)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case DDLParserAUTO_INCREMENT:
			{
				p.SetState(253)
				p.Match(DDLParserAUTO_INCREMENT)
			}

		case DDLParserON:
			{
				p.SetState(254)
				p.Match(DDLParserON)
			}
			{
				p.SetState(255)
				p.Match(DDLParserUPDATE)
			}
			{
				p.SetState(256)
				p.CurrentTimestamp()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

	case DDLParserPRIMARY:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(259)
			p.PrimaryKey()
		}

	case DDLParserKEY:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(260)
			p.Match(DDLParserKEY)
		}

	case DDLParserUNIQUE:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(261)
			p.Match(DDLParserUNIQUE)
		}
		p.SetState(263)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 42, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(262)
				p.Match(DDLParserKEY)
			}

		}

	case DDLParserCOMMENT:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(265)
			p.Comment()
		}

	case DDLParserCOLLATE:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(266)
			p.Match(DDLParserCOLLATE)
		}
		{
			p.SetState(267)
			p.Match(DDLParserBackQuotedString)
		}

	case DDLParserSERIAL:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(268)
			p.Match(DDLParserSERIAL)
		}
		{
			p.SetState(269)
			p.Match(DDLParserDEFAULT)
		}
		{
			p.SetState(270)
			p.Match(DDLParserVALUE)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// INullNotNullContext is an interface to support dynamic dispatch.
type INullNotNullContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNullNotNullContext differentiates from other interfaces.
	IsNullNotNullContext()
}

type NullNotNullContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNullNotNullContext() *NullNotNullContext {
	var p = new(NullNotNullContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DDLParserRULE_nullNotNull
	return p
}

func (*NullNotNullContext) IsNullNotNullContext() {}

func NewNullNotNullContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NullNotNullContext {
	var p = new(NullNotNullContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DDLParserRULE_nullNotNull

	return p
}

func (s *NullNotNullContext) GetParser() antlr.Parser { return s.parser }

func (s *NullNotNullContext) NULL() antlr.TerminalNode {
	return s.GetToken(DDLParserNULL, 0)
}

func (s *NullNotNullContext) NOT() antlr.TerminalNode {
	return s.GetToken(DDLParserNOT, 0)
}

func (s *NullNotNullContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullNotNullContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NullNotNullContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DDLParserListener); ok {
		listenerT.EnterNullNotNull(s)
	}
}

func (s *NullNotNullContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DDLParserListener); ok {
		listenerT.ExitNullNotNull(s)
	}
}

func (p *DDLParser) NullNotNull() (localctx INullNotNullContext) {
	localctx = NewNullNotNullContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, DDLParserRULE_nullNotNull)
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
	p.SetState(274)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DDLParserNOT {
		{
			p.SetState(273)
			p.Match(DDLParserNOT)
		}

	}
	{
		p.SetState(276)
		p.Match(DDLParserNULL)
	}

	return localctx
}

// ICommentContext is an interface to support dynamic dispatch.
type ICommentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetContent returns the content token.
	GetContent() antlr.Token

	// SetContent sets the content token.
	SetContent(antlr.Token)

	// IsCommentContext differentiates from other interfaces.
	IsCommentContext()
}

type CommentContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	content antlr.Token
}

func NewEmptyCommentContext() *CommentContext {
	var p = new(CommentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DDLParserRULE_comment
	return p
}

func (*CommentContext) IsCommentContext() {}

func NewCommentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommentContext {
	var p = new(CommentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DDLParserRULE_comment

	return p
}

func (s *CommentContext) GetParser() antlr.Parser { return s.parser }

func (s *CommentContext) GetContent() antlr.Token { return s.content }

func (s *CommentContext) SetContent(v antlr.Token) { s.content = v }

func (s *CommentContext) COMMENT() antlr.TerminalNode {
	return s.GetToken(DDLParserCOMMENT, 0)
}

func (s *CommentContext) Literal() antlr.TerminalNode {
	return s.GetToken(DDLParserLiteral, 0)
}

func (s *CommentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DDLParserListener); ok {
		listenerT.EnterComment(s)
	}
}

func (s *CommentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DDLParserListener); ok {
		listenerT.ExitComment(s)
	}
}

func (p *DDLParser) Comment() (localctx ICommentContext) {
	localctx = NewCommentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, DDLParserRULE_comment)

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
		p.SetState(278)
		p.Match(DDLParserCOMMENT)
	}
	{
		p.SetState(279)

		var _m = p.Match(DDLParserLiteral)

		localctx.(*CommentContext).content = _m
	}

	return localctx
}

// IDefaultValueContext is an interface to support dynamic dispatch.
type IDefaultValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefaultValueContext differentiates from other interfaces.
	IsDefaultValueContext()
}

type DefaultValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefaultValueContext() *DefaultValueContext {
	var p = new(DefaultValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DDLParserRULE_defaultValue
	return p
}

func (*DefaultValueContext) IsDefaultValueContext() {}

func NewDefaultValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefaultValueContext {
	var p = new(DefaultValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DDLParserRULE_defaultValue

	return p
}

func (s *DefaultValueContext) GetParser() antlr.Parser { return s.parser }

func (s *DefaultValueContext) NULL() antlr.TerminalNode {
	return s.GetToken(DDLParserNULL, 0)
}

func (s *DefaultValueContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *DefaultValueContext) UnaryOperator() IUnaryOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnaryOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnaryOperatorContext)
}

func (s *DefaultValueContext) AllCurrentTimestamp() []ICurrentTimestampContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICurrentTimestampContext)(nil)).Elem())
	var tst = make([]ICurrentTimestampContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICurrentTimestampContext)
		}
	}

	return tst
}

func (s *DefaultValueContext) CurrentTimestamp(i int) ICurrentTimestampContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICurrentTimestampContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICurrentTimestampContext)
}

func (s *DefaultValueContext) ON() antlr.TerminalNode {
	return s.GetToken(DDLParserON, 0)
}

func (s *DefaultValueContext) UPDATE() antlr.TerminalNode {
	return s.GetToken(DDLParserUPDATE, 0)
}

func (s *DefaultValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefaultValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DDLParserListener); ok {
		listenerT.EnterDefaultValue(s)
	}
}

func (s *DefaultValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DDLParserListener); ok {
		listenerT.ExitDefaultValue(s)
	}
}

func (p *DDLParser) DefaultValue() (localctx IDefaultValueContext) {
	localctx = NewDefaultValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, DDLParserRULE_defaultValue)

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

	p.SetState(292)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 47, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(281)
			p.Match(DDLParserNULL)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(283)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 45, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(282)
				p.UnaryOperator()
			}

		}
		{
			p.SetState(285)
			p.Constant()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(286)
			p.CurrentTimestamp()
		}
		p.SetState(290)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 46, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(287)
				p.Match(DDLParserON)
			}
			{
				p.SetState(288)
				p.Match(DDLParserUPDATE)
			}
			{
				p.SetState(289)
				p.CurrentTimestamp()
			}

		}

	}

	return localctx
}

// IPrimaryKeyContext is an interface to support dynamic dispatch.
type IPrimaryKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimaryKeyContext differentiates from other interfaces.
	IsPrimaryKeyContext()
}

type PrimaryKeyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryKeyContext() *PrimaryKeyContext {
	var p = new(PrimaryKeyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DDLParserRULE_primaryKey
	return p
}

func (*PrimaryKeyContext) IsPrimaryKeyContext() {}

func NewPrimaryKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryKeyContext {
	var p = new(PrimaryKeyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DDLParserRULE_primaryKey

	return p
}

func (s *PrimaryKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryKeyContext) PRIMARY() antlr.TerminalNode {
	return s.GetToken(DDLParserPRIMARY, 0)
}

func (s *PrimaryKeyContext) KEY() antlr.TerminalNode {
	return s.GetToken(DDLParserKEY, 0)
}

func (s *PrimaryKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryKeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DDLParserListener); ok {
		listenerT.EnterPrimaryKey(s)
	}
}

func (s *PrimaryKeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DDLParserListener); ok {
		listenerT.ExitPrimaryKey(s)
	}
}

func (p *DDLParser) PrimaryKey() (localctx IPrimaryKeyContext) {
	localctx = NewPrimaryKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, DDLParserRULE_primaryKey)

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
		p.SetState(294)
		p.Match(DDLParserPRIMARY)
	}
	{
		p.SetState(295)
		p.Match(DDLParserKEY)
	}

	return localctx
}

// IUnaryOperatorContext is an interface to support dynamic dispatch.
type IUnaryOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnaryOperatorContext differentiates from other interfaces.
	IsUnaryOperatorContext()
}

type UnaryOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryOperatorContext() *UnaryOperatorContext {
	var p = new(UnaryOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DDLParserRULE_unaryOperator
	return p
}

func (*UnaryOperatorContext) IsUnaryOperatorContext() {}

func NewUnaryOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryOperatorContext {
	var p = new(UnaryOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DDLParserRULE_unaryOperator

	return p
}

func (s *UnaryOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryOperatorContext) Exclamation() antlr.TerminalNode {
	return s.GetToken(DDLParserExclamation, 0)
}

func (s *UnaryOperatorContext) Tilde() antlr.TerminalNode {
	return s.GetToken(DDLParserTilde, 0)
}

func (s *UnaryOperatorContext) Plus() antlr.TerminalNode {
	return s.GetToken(DDLParserPlus, 0)
}

func (s *UnaryOperatorContext) Minus() antlr.TerminalNode {
	return s.GetToken(DDLParserMinus, 0)
}

func (s *UnaryOperatorContext) NOT() antlr.TerminalNode {
	return s.GetToken(DDLParserNOT, 0)
}

func (s *UnaryOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DDLParserListener); ok {
		listenerT.EnterUnaryOperator(s)
	}
}

func (s *UnaryOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DDLParserListener); ok {
		listenerT.ExitUnaryOperator(s)
	}
}

func (p *DDLParser) UnaryOperator() (localctx IUnaryOperatorContext) {
	localctx = NewUnaryOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, DDLParserRULE_unaryOperator)
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
		p.SetState(297)
		_la = p.GetTokenStream().LA(1)

		if !(_la == DDLParserNOT || (((_la-85)&-(0x1f+1)) == 0 && ((1<<uint((_la-85)))&((1<<(DDLParserTilde-85))|(1<<(DDLParserExclamation-85))|(1<<(DDLParserPlus-85))|(1<<(DDLParserMinus-85)))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IConstantContext is an interface to support dynamic dispatch.
type IConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstantContext differentiates from other interfaces.
	IsConstantContext()
}

type ConstantContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantContext() *ConstantContext {
	var p = new(ConstantContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DDLParserRULE_constant
	return p
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DDLParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) Literal() antlr.TerminalNode {
	return s.GetToken(DDLParserLiteral, 0)
}

func (s *ConstantContext) Integer() antlr.TerminalNode {
	return s.GetToken(DDLParserInteger, 0)
}

func (s *ConstantContext) Minus() antlr.TerminalNode {
	return s.GetToken(DDLParserMinus, 0)
}

func (s *ConstantContext) TRUE() antlr.TerminalNode {
	return s.GetToken(DDLParserTRUE, 0)
}

func (s *ConstantContext) FALSE() antlr.TerminalNode {
	return s.GetToken(DDLParserFALSE, 0)
}

func (s *ConstantContext) Number() antlr.TerminalNode {
	return s.GetToken(DDLParserNumber, 0)
}

func (s *ConstantContext) NULL() antlr.TerminalNode {
	return s.GetToken(DDLParserNULL, 0)
}

func (s *ConstantContext) NOT() antlr.TerminalNode {
	return s.GetToken(DDLParserNOT, 0)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DDLParserListener); ok {
		listenerT.EnterConstant(s)
	}
}

func (s *ConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DDLParserListener); ok {
		listenerT.ExitConstant(s)
	}
}

func (p *DDLParser) Constant() (localctx IConstantContext) {
	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, DDLParserRULE_constant)
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

	p.SetState(309)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DDLParserLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(299)
			p.Match(DDLParserLiteral)
		}

	case DDLParserInteger:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(300)
			p.Match(DDLParserInteger)
		}

	case DDLParserMinus:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(301)
			p.Match(DDLParserMinus)
		}
		{
			p.SetState(302)
			p.Match(DDLParserInteger)
		}

	case DDLParserFALSE, DDLParserTRUE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(303)
			_la = p.GetTokenStream().LA(1)

			if !(_la == DDLParserFALSE || _la == DDLParserTRUE) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case DDLParserNumber:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(304)
			p.Match(DDLParserNumber)
		}

	case DDLParserNOT, DDLParserNULL:
		p.EnterOuterAlt(localctx, 6)
		p.SetState(306)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DDLParserNOT {
			{
				p.SetState(305)
				p.Match(DDLParserNOT)
			}

		}
		{
			p.SetState(308)
			p.Match(DDLParserNULL)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ICurrentTimestampContext is an interface to support dynamic dispatch.
type ICurrentTimestampContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCurrentTimestampContext differentiates from other interfaces.
	IsCurrentTimestampContext()
}

type CurrentTimestampContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCurrentTimestampContext() *CurrentTimestampContext {
	var p = new(CurrentTimestampContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DDLParserRULE_currentTimestamp
	return p
}

func (*CurrentTimestampContext) IsCurrentTimestampContext() {}

func NewCurrentTimestampContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CurrentTimestampContext {
	var p = new(CurrentTimestampContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DDLParserRULE_currentTimestamp

	return p
}

func (s *CurrentTimestampContext) GetParser() antlr.Parser { return s.parser }

func (s *CurrentTimestampContext) NOW() antlr.TerminalNode {
	return s.GetToken(DDLParserNOW, 0)
}

func (s *CurrentTimestampContext) LeftParenthesis() antlr.TerminalNode {
	return s.GetToken(DDLParserLeftParenthesis, 0)
}

func (s *CurrentTimestampContext) RightParenthesis() antlr.TerminalNode {
	return s.GetToken(DDLParserRightParenthesis, 0)
}

func (s *CurrentTimestampContext) CURRENT_TIMESTAMP() antlr.TerminalNode {
	return s.GetToken(DDLParserCURRENT_TIMESTAMP, 0)
}

func (s *CurrentTimestampContext) LOCALTIME() antlr.TerminalNode {
	return s.GetToken(DDLParserLOCALTIME, 0)
}

func (s *CurrentTimestampContext) LOCALTIMESTAMP() antlr.TerminalNode {
	return s.GetToken(DDLParserLOCALTIMESTAMP, 0)
}

func (s *CurrentTimestampContext) Integer() antlr.TerminalNode {
	return s.GetToken(DDLParserInteger, 0)
}

func (s *CurrentTimestampContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CurrentTimestampContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CurrentTimestampContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DDLParserListener); ok {
		listenerT.EnterCurrentTimestamp(s)
	}
}

func (s *CurrentTimestampContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DDLParserListener); ok {
		listenerT.ExitCurrentTimestamp(s)
	}
}

func (p *DDLParser) CurrentTimestamp() (localctx ICurrentTimestampContext) {
	localctx = NewCurrentTimestampContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, DDLParserRULE_currentTimestamp)
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
	p.SetState(325)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DDLParserCURRENT_TIMESTAMP, DDLParserLOCALTIMESTAMP, DDLParserLOCALTIME:
		{
			p.SetState(311)
			_la = p.GetTokenStream().LA(1)

			if !(((_la-16)&-(0x1f+1)) == 0 && ((1<<uint((_la-16)))&((1<<(DDLParserCURRENT_TIMESTAMP-16))|(1<<(DDLParserLOCALTIMESTAMP-16))|(1<<(DDLParserLOCALTIME-16)))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		p.SetState(317)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DDLParserLeftParenthesis {
			{
				p.SetState(312)
				p.Match(DDLParserLeftParenthesis)
			}
			p.SetState(314)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == DDLParserInteger {
				{
					p.SetState(313)
					p.Match(DDLParserInteger)
				}

			}
			{
				p.SetState(316)
				p.Match(DDLParserRightParenthesis)
			}

		}

	case DDLParserNOW:
		{
			p.SetState(319)
			p.Match(DDLParserNOW)
		}
		{
			p.SetState(320)
			p.Match(DDLParserLeftParenthesis)
		}
		p.SetState(322)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DDLParserInteger {
			{
				p.SetState(321)
				p.Match(DDLParserInteger)
			}

		}
		{
			p.SetState(324)
			p.Match(DDLParserRightParenthesis)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITableConstraintContext is an interface to support dynamic dispatch.
type ITableConstraintContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// GetIndex returns the index token.
	GetIndex() antlr.Token

	// GetIndexFormat returns the indexFormat token.
	GetIndexFormat() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// SetIndex sets the index token.
	SetIndex(antlr.Token)

	// SetIndexFormat sets the indexFormat token.
	SetIndexFormat(antlr.Token)

	// GetPk returns the pk rule contexts.
	GetPk() IPrimaryKeyContext

	// SetPk sets the pk rule contexts.
	SetPk(IPrimaryKeyContext)

	// IsTableConstraintContext differentiates from other interfaces.
	IsTableConstraintContext()
}

type TableConstraintContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	name        antlr.Token
	pk          IPrimaryKeyContext
	index       antlr.Token
	indexFormat antlr.Token
}

func NewEmptyTableConstraintContext() *TableConstraintContext {
	var p = new(TableConstraintContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DDLParserRULE_tableConstraint
	return p
}

func (*TableConstraintContext) IsTableConstraintContext() {}

func NewTableConstraintContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableConstraintContext {
	var p = new(TableConstraintContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DDLParserRULE_tableConstraint

	return p
}

func (s *TableConstraintContext) GetParser() antlr.Parser { return s.parser }

func (s *TableConstraintContext) GetName() antlr.Token { return s.name }

func (s *TableConstraintContext) GetIndex() antlr.Token { return s.index }

func (s *TableConstraintContext) GetIndexFormat() antlr.Token { return s.indexFormat }

func (s *TableConstraintContext) SetName(v antlr.Token) { s.name = v }

func (s *TableConstraintContext) SetIndex(v antlr.Token) { s.index = v }

func (s *TableConstraintContext) SetIndexFormat(v antlr.Token) { s.indexFormat = v }

func (s *TableConstraintContext) GetPk() IPrimaryKeyContext { return s.pk }

func (s *TableConstraintContext) SetPk(v IPrimaryKeyContext) { s.pk = v }

func (s *TableConstraintContext) IndexColumnNames() IIndexColumnNamesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndexColumnNamesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndexColumnNamesContext)
}

func (s *TableConstraintContext) PrimaryKey() IPrimaryKeyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryKeyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimaryKeyContext)
}

func (s *TableConstraintContext) CONSTRAINT() antlr.TerminalNode {
	return s.GetToken(DDLParserCONSTRAINT, 0)
}

func (s *TableConstraintContext) AllIndexOption() []IIndexOptionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIndexOptionContext)(nil)).Elem())
	var tst = make([]IIndexOptionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIndexOptionContext)
		}
	}

	return tst
}

func (s *TableConstraintContext) IndexOption(i int) IIndexOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndexOptionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIndexOptionContext)
}

func (s *TableConstraintContext) BackQuotedString() antlr.TerminalNode {
	return s.GetToken(DDLParserBackQuotedString, 0)
}

func (s *TableConstraintContext) AllRawString() []antlr.TerminalNode {
	return s.GetTokens(DDLParserRawString)
}

func (s *TableConstraintContext) RawString(i int) antlr.TerminalNode {
	return s.GetToken(DDLParserRawString, i)
}

func (s *TableConstraintContext) AllLiteral() []antlr.TerminalNode {
	return s.GetTokens(DDLParserLiteral)
}

func (s *TableConstraintContext) Literal(i int) antlr.TerminalNode {
	return s.GetToken(DDLParserLiteral, i)
}

func (s *TableConstraintContext) UNIQUE() antlr.TerminalNode {
	return s.GetToken(DDLParserUNIQUE, 0)
}

func (s *TableConstraintContext) IndexType() IIndexTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndexTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndexTypeContext)
}

func (s *TableConstraintContext) INDEX() antlr.TerminalNode {
	return s.GetToken(DDLParserINDEX, 0)
}

func (s *TableConstraintContext) KEY() antlr.TerminalNode {
	return s.GetToken(DDLParserKEY, 0)
}

func (s *TableConstraintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableConstraintContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TableConstraintContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DDLParserListener); ok {
		listenerT.EnterTableConstraint(s)
	}
}

func (s *TableConstraintContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DDLParserListener); ok {
		listenerT.ExitTableConstraint(s)
	}
}

func (p *DDLParser) TableConstraint() (localctx ITableConstraintContext) {
	localctx = NewTableConstraintContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, DDLParserRULE_tableConstraint)
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

	p.SetState(372)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 66, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(331)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DDLParserCONSTRAINT {
			{
				p.SetState(327)
				p.Match(DDLParserCONSTRAINT)
			}
			p.SetState(329)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == DDLParserLiteral || _la == DDLParserRawString {
				{
					p.SetState(328)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*TableConstraintContext).name = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == DDLParserLiteral || _la == DDLParserRawString) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*TableConstraintContext).name = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}

			}

		}
		{
			p.SetState(333)

			var _x = p.PrimaryKey()

			localctx.(*TableConstraintContext).pk = _x
		}
		p.SetState(335)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 56, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(334)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*TableConstraintContext).index = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == DDLParserLiteral || _la == DDLParserRawString) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*TableConstraintContext).index = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		p.SetState(338)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DDLParserBackQuotedString || _la == DDLParserRawString {
			{
				p.SetState(337)
				_la = p.GetTokenStream().LA(1)

				if !(_la == DDLParserBackQuotedString || _la == DDLParserRawString) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		{
			p.SetState(340)
			p.IndexColumnNames()
		}
		p.SetState(344)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == DDLParserCOMMENT || _la == DDLParserKEY_BLOCK_SIZE || _la == DDLParserWITH {
			{
				p.SetState(341)
				p.IndexOption()
			}

			p.SetState(346)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(351)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DDLParserCONSTRAINT {
			{
				p.SetState(347)
				p.Match(DDLParserCONSTRAINT)
			}
			p.SetState(349)
			p.GetErrorHandler().Sync(p)

			if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 59, p.GetParserRuleContext()) == 1 {
				{
					p.SetState(348)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*TableConstraintContext).name = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == DDLParserLiteral || _la == DDLParserRawString) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*TableConstraintContext).name = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}

			}

		}
		p.SetState(354)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DDLParserUNIQUE {
			{
				p.SetState(353)
				p.Match(DDLParserUNIQUE)
			}

		}
		p.SetState(357)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DDLParserINDEX || _la == DDLParserKEY {
			{
				p.SetState(356)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*TableConstraintContext).indexFormat = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == DDLParserINDEX || _la == DDLParserKEY) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*TableConstraintContext).indexFormat = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		p.SetState(360)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DDLParserBackQuotedString || _la == DDLParserRawString {
			{
				p.SetState(359)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*TableConstraintContext).index = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == DDLParserBackQuotedString || _la == DDLParserRawString) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*TableConstraintContext).index = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		p.SetState(363)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DDLParserUSING {
			{
				p.SetState(362)
				p.IndexType()
			}

		}
		{
			p.SetState(365)
			p.IndexColumnNames()
		}
		p.SetState(369)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == DDLParserCOMMENT || _la == DDLParserKEY_BLOCK_SIZE || _la == DDLParserWITH {
			{
				p.SetState(366)
				p.IndexOption()
			}

			p.SetState(371)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}

	return localctx
}

// IIndexOptionContext is an interface to support dynamic dispatch.
type IIndexOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIndexOptionContext differentiates from other interfaces.
	IsIndexOptionContext()
}

type IndexOptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexOptionContext() *IndexOptionContext {
	var p = new(IndexOptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DDLParserRULE_indexOption
	return p
}

func (*IndexOptionContext) IsIndexOptionContext() {}

func NewIndexOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexOptionContext {
	var p = new(IndexOptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DDLParserRULE_indexOption

	return p
}

func (s *IndexOptionContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexOptionContext) KEY_BLOCK_SIZE() antlr.TerminalNode {
	return s.GetToken(DDLParserKEY_BLOCK_SIZE, 0)
}

func (s *IndexOptionContext) FilesizeLiteral() antlr.TerminalNode {
	return s.GetToken(DDLParserFilesizeLiteral, 0)
}

func (s *IndexOptionContext) Equal() antlr.TerminalNode {
	return s.GetToken(DDLParserEqual, 0)
}

func (s *IndexOptionContext) WITH() antlr.TerminalNode {
	return s.GetToken(DDLParserWITH, 0)
}

func (s *IndexOptionContext) PARSER() antlr.TerminalNode {
	return s.GetToken(DDLParserPARSER, 0)
}

func (s *IndexOptionContext) Literal() antlr.TerminalNode {
	return s.GetToken(DDLParserLiteral, 0)
}

func (s *IndexOptionContext) RawString() antlr.TerminalNode {
	return s.GetToken(DDLParserRawString, 0)
}

func (s *IndexOptionContext) COMMENT() antlr.TerminalNode {
	return s.GetToken(DDLParserCOMMENT, 0)
}

func (s *IndexOptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexOptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IndexOptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DDLParserListener); ok {
		listenerT.EnterIndexOption(s)
	}
}

func (s *IndexOptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DDLParserListener); ok {
		listenerT.ExitIndexOption(s)
	}
}

func (p *DDLParser) IndexOption() (localctx IIndexOptionContext) {
	localctx = NewIndexOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, DDLParserRULE_indexOption)
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

	p.SetState(384)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DDLParserKEY_BLOCK_SIZE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(374)
			p.Match(DDLParserKEY_BLOCK_SIZE)
		}
		p.SetState(376)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DDLParserEqual {
			{
				p.SetState(375)
				p.Match(DDLParserEqual)
			}

		}
		{
			p.SetState(378)
			p.Match(DDLParserFilesizeLiteral)
		}

	case DDLParserWITH:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(379)
			p.Match(DDLParserWITH)
		}
		{
			p.SetState(380)
			p.Match(DDLParserPARSER)
		}
		{
			p.SetState(381)
			_la = p.GetTokenStream().LA(1)

			if !(_la == DDLParserLiteral || _la == DDLParserRawString) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case DDLParserCOMMENT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(382)
			p.Match(DDLParserCOMMENT)
		}
		{
			p.SetState(383)
			p.Match(DDLParserLiteral)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIndexTypeContext is an interface to support dynamic dispatch.
type IIndexTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIndexTypeContext differentiates from other interfaces.
	IsIndexTypeContext()
}

type IndexTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexTypeContext() *IndexTypeContext {
	var p = new(IndexTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DDLParserRULE_indexType
	return p
}

func (*IndexTypeContext) IsIndexTypeContext() {}

func NewIndexTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexTypeContext {
	var p = new(IndexTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DDLParserRULE_indexType

	return p
}

func (s *IndexTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexTypeContext) USING() antlr.TerminalNode {
	return s.GetToken(DDLParserUSING, 0)
}

func (s *IndexTypeContext) BTREE() antlr.TerminalNode {
	return s.GetToken(DDLParserBTREE, 0)
}

func (s *IndexTypeContext) HASH() antlr.TerminalNode {
	return s.GetToken(DDLParserHASH, 0)
}

func (s *IndexTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IndexTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DDLParserListener); ok {
		listenerT.EnterIndexType(s)
	}
}

func (s *IndexTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DDLParserListener); ok {
		listenerT.ExitIndexType(s)
	}
}

func (p *DDLParser) IndexType() (localctx IIndexTypeContext) {
	localctx = NewIndexTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, DDLParserRULE_indexType)
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
		p.SetState(386)
		p.Match(DDLParserUSING)
	}
	{
		p.SetState(387)
		_la = p.GetTokenStream().LA(1)

		if !(_la == DDLParserBTREE || _la == DDLParserHASH) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IIndexColumnNamesContext is an interface to support dynamic dispatch.
type IIndexColumnNamesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIndexColumnNamesContext differentiates from other interfaces.
	IsIndexColumnNamesContext()
}

type IndexColumnNamesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexColumnNamesContext() *IndexColumnNamesContext {
	var p = new(IndexColumnNamesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DDLParserRULE_indexColumnNames
	return p
}

func (*IndexColumnNamesContext) IsIndexColumnNamesContext() {}

func NewIndexColumnNamesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexColumnNamesContext {
	var p = new(IndexColumnNamesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DDLParserRULE_indexColumnNames

	return p
}

func (s *IndexColumnNamesContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexColumnNamesContext) AllLeftParenthesis() []antlr.TerminalNode {
	return s.GetTokens(DDLParserLeftParenthesis)
}

func (s *IndexColumnNamesContext) LeftParenthesis(i int) antlr.TerminalNode {
	return s.GetToken(DDLParserLeftParenthesis, i)
}

func (s *IndexColumnNamesContext) AllRightParenthesis() []antlr.TerminalNode {
	return s.GetTokens(DDLParserRightParenthesis)
}

func (s *IndexColumnNamesContext) RightParenthesis(i int) antlr.TerminalNode {
	return s.GetToken(DDLParserRightParenthesis, i)
}

func (s *IndexColumnNamesContext) AllBackQuotedString() []antlr.TerminalNode {
	return s.GetTokens(DDLParserBackQuotedString)
}

func (s *IndexColumnNamesContext) BackQuotedString(i int) antlr.TerminalNode {
	return s.GetToken(DDLParserBackQuotedString, i)
}

func (s *IndexColumnNamesContext) AllRawString() []antlr.TerminalNode {
	return s.GetTokens(DDLParserRawString)
}

func (s *IndexColumnNamesContext) RawString(i int) antlr.TerminalNode {
	return s.GetToken(DDLParserRawString, i)
}

func (s *IndexColumnNamesContext) AllInteger() []antlr.TerminalNode {
	return s.GetTokens(DDLParserInteger)
}

func (s *IndexColumnNamesContext) Integer(i int) antlr.TerminalNode {
	return s.GetToken(DDLParserInteger, i)
}

func (s *IndexColumnNamesContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(DDLParserComma)
}

func (s *IndexColumnNamesContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(DDLParserComma, i)
}

func (s *IndexColumnNamesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexColumnNamesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IndexColumnNamesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DDLParserListener); ok {
		listenerT.EnterIndexColumnNames(s)
	}
}

func (s *IndexColumnNamesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DDLParserListener); ok {
		listenerT.ExitIndexColumnNames(s)
	}
}

func (p *DDLParser) IndexColumnNames() (localctx IIndexColumnNamesContext) {
	localctx = NewIndexColumnNamesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, DDLParserRULE_indexColumnNames)
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
		p.SetState(389)
		p.Match(DDLParserLeftParenthesis)
	}
	{
		p.SetState(390)
		_la = p.GetTokenStream().LA(1)

		if !(_la == DDLParserBackQuotedString || _la == DDLParserRawString) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(394)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DDLParserLeftParenthesis {
		{
			p.SetState(391)
			p.Match(DDLParserLeftParenthesis)
		}
		{
			p.SetState(392)
			p.Match(DDLParserInteger)
		}
		{
			p.SetState(393)
			p.Match(DDLParserRightParenthesis)
		}

	}
	p.SetState(405)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DDLParserComma {
		{
			p.SetState(396)
			p.Match(DDLParserComma)
		}
		{
			p.SetState(397)
			_la = p.GetTokenStream().LA(1)

			if !(_la == DDLParserBackQuotedString || _la == DDLParserRawString) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		p.SetState(401)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DDLParserLeftParenthesis {
			{
				p.SetState(398)
				p.Match(DDLParserLeftParenthesis)
			}
			{
				p.SetState(399)
				p.Match(DDLParserInteger)
			}
			{
				p.SetState(400)
				p.Match(DDLParserRightParenthesis)
			}

		}

		p.SetState(407)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(408)
		p.Match(DDLParserRightParenthesis)
	}

	return localctx
}

// ITableOptionContext is an interface to support dynamic dispatch.
type ITableOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTableComment returns the tableComment token.
	GetTableComment() antlr.Token

	// SetTableComment sets the tableComment token.
	SetTableComment(antlr.Token)

	// IsTableOptionContext differentiates from other interfaces.
	IsTableOptionContext()
}

type TableOptionContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	tableComment antlr.Token
}

func NewEmptyTableOptionContext() *TableOptionContext {
	var p = new(TableOptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DDLParserRULE_tableOption
	return p
}

func (*TableOptionContext) IsTableOptionContext() {}

func NewTableOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableOptionContext {
	var p = new(TableOptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DDLParserRULE_tableOption

	return p
}

func (s *TableOptionContext) GetParser() antlr.Parser { return s.parser }

func (s *TableOptionContext) GetTableComment() antlr.Token { return s.tableComment }

func (s *TableOptionContext) SetTableComment(v antlr.Token) { s.tableComment = v }

func (s *TableOptionContext) ENGINE() antlr.TerminalNode {
	return s.GetToken(DDLParserENGINE, 0)
}

func (s *TableOptionContext) BackQuotedString() antlr.TerminalNode {
	return s.GetToken(DDLParserBackQuotedString, 0)
}

func (s *TableOptionContext) RawString() antlr.TerminalNode {
	return s.GetToken(DDLParserRawString, 0)
}

func (s *TableOptionContext) Equal() antlr.TerminalNode {
	return s.GetToken(DDLParserEqual, 0)
}

func (s *TableOptionContext) AUTO_INCREMENT() antlr.TerminalNode {
	return s.GetToken(DDLParserAUTO_INCREMENT, 0)
}

func (s *TableOptionContext) Integer() antlr.TerminalNode {
	return s.GetToken(DDLParserInteger, 0)
}

func (s *TableOptionContext) CHARACTER() antlr.TerminalNode {
	return s.GetToken(DDLParserCHARACTER, 0)
}

func (s *TableOptionContext) SET() antlr.TerminalNode {
	return s.GetToken(DDLParserSET, 0)
}

func (s *TableOptionContext) CHARSET() antlr.TerminalNode {
	return s.GetToken(DDLParserCHARSET, 0)
}

func (s *TableOptionContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(DDLParserDEFAULT, 0)
}

func (s *TableOptionContext) COLLATE() antlr.TerminalNode {
	return s.GetToken(DDLParserCOLLATE, 0)
}

func (s *TableOptionContext) COMMENT() antlr.TerminalNode {
	return s.GetToken(DDLParserCOMMENT, 0)
}

func (s *TableOptionContext) Literal() antlr.TerminalNode {
	return s.GetToken(DDLParserLiteral, 0)
}

func (s *TableOptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableOptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TableOptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DDLParserListener); ok {
		listenerT.EnterTableOption(s)
	}
}

func (s *TableOptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DDLParserListener); ok {
		listenerT.ExitTableOption(s)
	}
}

func (p *DDLParser) TableOption() (localctx ITableOptionContext) {
	localctx = NewTableOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, DDLParserRULE_tableOption)
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

	p.SetState(445)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 80, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(410)
			p.Match(DDLParserENGINE)
		}
		p.SetState(412)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DDLParserEqual {
			{
				p.SetState(411)
				p.Match(DDLParserEqual)
			}

		}
		{
			p.SetState(414)
			_la = p.GetTokenStream().LA(1)

			if !(_la == DDLParserBackQuotedString || _la == DDLParserRawString) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(415)
			p.Match(DDLParserAUTO_INCREMENT)
		}
		p.SetState(417)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DDLParserEqual {
			{
				p.SetState(416)
				p.Match(DDLParserEqual)
			}

		}
		{
			p.SetState(419)
			p.Match(DDLParserInteger)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(421)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DDLParserDEFAULT {
			{
				p.SetState(420)
				p.Match(DDLParserDEFAULT)
			}

		}
		p.SetState(426)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case DDLParserCHARACTER:
			{
				p.SetState(423)
				p.Match(DDLParserCHARACTER)
			}
			{
				p.SetState(424)
				p.Match(DDLParserSET)
			}

		case DDLParserCHARSET:
			{
				p.SetState(425)
				p.Match(DDLParserCHARSET)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		p.SetState(429)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DDLParserEqual {
			{
				p.SetState(428)
				p.Match(DDLParserEqual)
			}

		}
		{
			p.SetState(431)
			_la = p.GetTokenStream().LA(1)

			if !(_la == DDLParserBackQuotedString || _la == DDLParserRawString) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		p.SetState(433)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DDLParserDEFAULT {
			{
				p.SetState(432)
				p.Match(DDLParserDEFAULT)
			}

		}
		{
			p.SetState(435)
			p.Match(DDLParserCOLLATE)
		}
		p.SetState(437)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DDLParserEqual {
			{
				p.SetState(436)
				p.Match(DDLParserEqual)
			}

		}
		{
			p.SetState(439)
			_la = p.GetTokenStream().LA(1)

			if !(_la == DDLParserBackQuotedString || _la == DDLParserRawString) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(440)
			p.Match(DDLParserCOMMENT)
		}
		p.SetState(442)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DDLParserEqual {
			{
				p.SetState(441)
				p.Match(DDLParserEqual)
			}

		}
		{
			p.SetState(444)

			var _m = p.Match(DDLParserLiteral)

			localctx.(*TableOptionContext).tableComment = _m
		}

	}

	return localctx
}
