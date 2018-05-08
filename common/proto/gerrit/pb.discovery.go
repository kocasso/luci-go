// AUTOGENERATED. DO NOT EDIT.

package gerrit

import discovery "go.chromium.org/luci/grpc/discovery"

import "github.com/golang/protobuf/protoc-gen-go/descriptor"

func init() {
	discovery.RegisterDescriptorSetCompressed(
		[]string{
			"gerrit.Gerrit",
		},
		[]byte{31, 139,
			8, 0, 0, 0, 0, 0, 0, 255, 212, 122, 239, 111, 35, 199,
			121, 191, 102, 102, 73, 73, 163, 19, 69, 14, 41, 157, 110, 19,
			223, 119, 190, 74, 108, 223, 217, 18, 101, 159, 155, 216, 61, 167,
			137, 121, 20, 117, 199, 156, 78, 188, 144, 146, 175, 70, 3, 8,
			203, 229, 144, 220, 120, 185, 203, 219, 217, 149, 78, 233, 33, 22,
			250, 34, 5, 106, 56, 64, 2, 183, 136, 81, 35, 1, 106, 160,
			117, 81, 196, 239, 140, 182, 112, 10, 24, 70, 157, 247, 69, 223,
			22, 8, 208, 151, 69, 209, 63, 160, 64, 139, 226, 153, 217, 89,
			146, 247, 35, 118, 218, 162, 69, 239, 205, 241, 153, 157, 125, 158,
			207, 243, 204, 243, 107, 158, 21, 253, 247, 61, 250, 149, 65, 88,
			117, 135, 81, 56, 242, 146, 81, 53, 140, 6, 219, 126, 226, 122,
			219, 110, 56, 26, 133, 193, 246, 56, 10, 227, 112, 123, 32, 162,
			200, 139, 211, 255, 170, 106, 141, 229, 53, 181, 241, 26, 45, 94,
			23, 113, 125, 232, 4, 3, 209, 22, 119, 19, 33, 99, 182, 70,
			243, 65, 50, 234, 138, 104, 29, 113, 116, 137, 180, 83, 138, 109,
			209, 249, 112, 28, 123, 97, 32, 215, 49, 39, 151, 10, 87, 202,
			213, 148, 231, 183, 18, 17, 157, 182, 212, 179, 182, 217, 179, 241,
			93, 202, 234, 67, 225, 190, 94, 115, 93, 33, 165, 97, 190, 78,
			231, 199, 81, 248, 29, 225, 198, 138, 251, 98, 219, 144, 172, 72,
			73, 36, 250, 235, 88, 173, 194, 79, 118, 145, 210, 177, 136, 70,
			158, 148, 94, 24, 172, 19, 245, 96, 106, 5, 120, 57, 174, 27,
			38, 65, 188, 110, 105, 94, 41, 185, 241, 167, 136, 150, 103, 132,
			203, 113, 24, 72, 193, 174, 210, 188, 140, 157, 56, 145, 74, 120,
			225, 202, 134, 209, 224, 17, 155, 171, 29, 181, 179, 157, 190, 1,
			102, 137, 132, 35, 195, 32, 133, 152, 82, 27, 47, 211, 188, 222,
			201, 150, 232, 124, 109, 111, 175, 117, 167, 177, 83, 156, 99, 203,
			116, 113, 183, 213, 190, 214, 220, 217, 105, 236, 23, 17, 91, 165,
			165, 219, 237, 214, 55, 27, 245, 131, 163, 253, 214, 193, 209, 110,
			235, 112, 127, 167, 136, 55, 190, 71, 151, 106, 26, 115, 51, 232,
			135, 140, 81, 43, 112, 70, 34, 53, 141, 250, 205, 42, 52, 39,
			70, 142, 231, 167, 98, 53, 193, 46, 211, 162, 20, 110, 24, 244,
			156, 232, 244, 72, 45, 201, 117, 194, 201, 165, 197, 246, 74, 182,
			222, 80, 203, 204, 166, 11, 137, 20, 145, 98, 172, 237, 148, 209,
			27, 45, 74, 245, 225, 43, 241, 143, 59, 249, 203, 52, 23, 158,
			4, 34, 82, 16, 150, 38, 231, 62, 5, 189, 173, 119, 60, 243,
			247, 132, 46, 77, 185, 3, 91, 163, 172, 117, 251, 160, 217, 218,
			63, 58, 220, 239, 220, 110, 212, 155, 187, 77, 101, 30, 74, 243,
			123, 181, 107, 141, 189, 78, 17, 177, 50, 93, 217, 105, 28, 212,
			154, 123, 141, 157, 163, 116, 17, 179, 10, 45, 214, 15, 219, 237,
			198, 254, 193, 81, 187, 241, 106, 179, 211, 108, 237, 23, 45, 86,
			162, 203, 181, 189, 189, 108, 165, 83, 92, 96, 140, 22, 204, 198,
			122, 235, 214, 173, 230, 65, 177, 200, 86, 232, 18, 108, 211, 116,
			167, 200, 225, 61, 179, 105, 183, 185, 215, 232, 20, 95, 97, 5,
			186, 8, 123, 52, 121, 134, 216, 26, 45, 101, 40, 106, 245, 122,
			235, 112, 255, 160, 83, 60, 195, 108, 149, 22, 65, 92, 227, 78,
			163, 125, 116, 120, 123, 167, 118, 0, 219, 45, 182, 76, 23, 110,
			53, 58, 157, 218, 117, 32, 23, 88, 133, 174, 24, 1, 181, 250,
			129, 130, 118, 86, 100, 101, 90, 168, 223, 168, 237, 95, 111, 76,
			22, 57, 188, 153, 50, 220, 41, 158, 189, 194, 42, 180, 208, 185,
			217, 188, 125, 116, 171, 209, 190, 222, 168, 93, 219, 107, 20, 207,
			206, 16, 43, 209, 165, 206, 225, 181, 91, 205, 131, 131, 116, 9,
			179, 21, 186, 120, 167, 113, 237, 104, 175, 185, 127, 179, 83, 60,
			59, 179, 216, 18, 205, 213, 111, 52, 234, 55, 139, 103, 10, 64,
			65, 171, 123, 180, 219, 106, 29, 52, 218, 176, 165, 200, 206, 211,
			210, 237, 195, 206, 141, 163, 122, 163, 125, 208, 220, 109, 214, 53,
			250, 51, 206, 24, 61, 119, 208, 174, 213, 111, 54, 247, 175, 31,
			53, 119, 96, 237, 21, 182, 78, 75, 59, 173, 59, 251, 123, 173,
			218, 142, 50, 93, 109, 95, 61, 56, 67, 87, 222, 66, 52, 127,
			93, 157, 57, 251, 77, 186, 152, 101, 12, 182, 110, 60, 225, 193,
			36, 98, 179, 73, 100, 101, 238, 117, 131, 46, 77, 197, 25, 179,
			31, 25, 124, 250, 245, 47, 252, 138, 192, 220, 152, 251, 230, 191,
			93, 163, 243, 44, 87, 152, 251, 5, 66, 244, 111, 45, 138, 206,
			49, 82, 152, 99, 87, 254, 218, 226, 245, 112, 124, 26, 121, 131,
			97, 204, 175, 60, 247, 252, 75, 252, 96, 40, 248, 222, 97, 189,
			201, 107, 73, 60, 12, 35, 89, 165, 148, 239, 121, 174, 8, 164,
			232, 241, 36, 232, 137, 136, 199, 67, 193, 107, 99, 199, 133, 157,
			250, 201, 38, 127, 85, 68, 144, 97, 248, 149, 234, 115, 252, 18,
			108, 216, 72, 31, 109, 92, 126, 153, 242, 211, 48, 225, 35, 231,
			148, 7, 97, 204, 19, 41, 120, 60, 244, 36, 239, 123, 190, 224,
			226, 158, 43, 198, 49, 247, 2, 238, 134, 163, 177, 239, 57, 129,
			43, 248, 137, 23, 15, 149, 148, 148, 71, 149, 242, 215, 82, 14,
			97, 55, 118, 188, 128, 59, 220, 13, 199, 167, 60, 236, 79, 111,
			227, 78, 76, 41, 87, 255, 134, 113, 60, 190, 186, 189, 125, 114,
			114, 82, 117, 20, 82, 157, 233, 245, 62, 185, 189, 215, 172, 55,
			246, 59, 141, 173, 43, 213, 231, 40, 229, 135, 129, 47, 164, 228,
			145, 184, 155, 120, 145, 232, 241, 238, 41, 119, 198, 99, 223, 115,
			157, 174, 47, 184, 239, 156, 240, 48, 226, 206, 32, 18, 162, 199,
			227, 16, 176, 158, 68, 94, 236, 5, 131, 77, 46, 195, 126, 124,
			226, 68, 130, 242, 158, 39, 227, 200, 235, 38, 241, 140, 153, 12,
			50, 79, 206, 108, 8, 3, 238, 4, 124, 163, 214, 225, 205, 206,
			6, 191, 86, 235, 52, 59, 155, 148, 223, 105, 30, 220, 104, 29,
			30, 240, 59, 181, 118, 187, 182, 127, 208, 108, 116, 120, 171, 205,
			235, 173, 253, 157, 166, 10, 5, 222, 218, 229, 181, 253, 215, 248,
			205, 230, 254, 206, 38, 23, 94, 60, 20, 17, 23, 247, 198, 17,
			160, 15, 35, 238, 129, 1, 69, 175, 74, 121, 71, 136, 25, 241,
			253, 80, 195, 145, 99, 225, 122, 125, 207, 229, 190, 19, 12, 18,
			103, 32, 248, 32, 60, 22, 81, 224, 5, 3, 62, 41, 19, 146,
			59, 65, 143, 114, 223, 27, 121, 177, 163, 74, 211, 195, 26, 85,
			41, 93, 160, 8, 51, 82, 92, 40, 80, 74, 113, 126, 142, 89,
			108, 238, 255, 33, 74, 41, 201, 207, 33, 70, 216, 66, 129, 126,
			153, 90, 249, 57, 60, 199, 72, 5, 191, 104, 159, 231, 123, 161,
			211, 147, 112, 116, 202, 193, 193, 204, 94, 175, 74, 233, 57, 154,
			131, 93, 136, 145, 74, 126, 197, 80, 152, 145, 74, 145, 27, 138,
			48, 82, 121, 246, 43, 244, 159, 176, 98, 136, 24, 185, 136, 119,
			237, 127, 192, 92, 249, 186, 228, 94, 159, 59, 224, 87, 17, 31,
			58, 32, 96, 162, 10, 88, 122, 224, 29, 139, 128, 71, 98, 28,
			130, 94, 60, 18, 125, 240, 233, 67, 9, 74, 43, 71, 172, 221,
			110, 154, 195, 151, 252, 85, 79, 156, 232, 232, 225, 174, 51, 118,
			186, 158, 239, 197, 167, 192, 70, 71, 51, 151, 34, 58, 22, 81,
			149, 242, 102, 95, 177, 11, 3, 255, 20, 0, 128, 109, 92, 199,
			247, 69, 196, 123, 161, 144, 193, 211, 49, 31, 58, 199, 130, 203,
			196, 29, 78, 113, 218, 84, 27, 35, 33, 19, 63, 6, 175, 16,
			81, 20, 70, 20, 14, 168, 235, 245, 122, 34, 0, 104, 251, 97,
			44, 54, 249, 56, 10, 143, 189, 30, 128, 12, 194, 96, 75, 220,
			243, 100, 44, 130, 152, 167, 197, 159, 67, 49, 2, 6, 16, 82,
			78, 160, 249, 108, 242, 110, 18, 167, 204, 189, 96, 64, 249, 35,
			114, 1, 63, 241, 124, 159, 67, 133, 6, 23, 113, 38, 12, 79,
			28, 205, 172, 31, 38, 193, 228, 88, 16, 24, 59, 95, 50, 20,
			102, 228, 34, 251, 178, 161, 8, 35, 23, 183, 175, 209, 191, 66,
			20, 231, 230, 152, 245, 228, 220, 8, 217, 127, 129, 120, 71, 123,
			153, 144, 252, 4, 4, 136, 123, 113, 228, 112, 47, 232, 135, 209,
			72, 185, 147, 14, 34, 215, 79, 122, 2, 130, 41, 53, 136, 110,
			33, 40, 229, 157, 48, 137, 92, 161, 98, 59, 74, 226, 225, 85,
			170, 162, 89, 94, 221, 78, 27, 178, 173, 72, 28, 123, 226, 164,
			58, 8, 195, 129, 47, 164, 218, 93, 117, 195, 209, 246, 78, 232,
			38, 35, 17, 104, 159, 221, 142, 132, 140, 183, 156, 177, 183, 165,
			221, 77, 86, 135, 241, 200, 255, 210, 93, 168, 180, 91, 105, 199,
			69, 193, 91, 115, 224, 121, 79, 230, 138, 116, 137, 90, 57, 229,
			173, 79, 225, 11, 160, 97, 78, 59, 229, 83, 184, 98, 40, 204,
			200, 83, 231, 215, 233, 239, 170, 141, 136, 145, 203, 120, 217, 14,
			120, 141, 203, 100, 52, 114, 34, 149, 142, 132, 227, 14, 185, 239,
			116, 133, 63, 201, 39, 16, 125, 50, 233, 142, 188, 120, 83, 185,
			140, 51, 134, 179, 21, 145, 212, 7, 0, 110, 66, 249, 32, 114,
			2, 72, 13, 151, 194, 136, 71, 2, 78, 68, 244, 46, 155, 76,
			232, 196, 154, 167, 62, 150, 156, 62, 150, 203, 120, 193, 80, 152,
			145, 203, 75, 231, 232, 191, 34, 133, 12, 51, 82, 197, 107, 246,
			63, 35, 190, 35, 98, 199, 243, 69, 47, 69, 52, 117, 6, 155,
			233, 9, 40, 255, 74, 70, 34, 242, 92, 126, 236, 248, 137, 144,
			160, 133, 227, 251, 92, 121, 156, 242, 34, 141, 215, 241, 229, 38,
			143, 132, 27, 14, 2, 239, 187, 25, 75, 253, 206, 166, 121, 87,
			69, 94, 28, 235, 68, 217, 21, 92, 138, 24, 226, 60, 30, 10,
			202, 221, 36, 138, 192, 131, 33, 80, 55, 149, 8, 125, 144, 96,
			136, 238, 169, 118, 201, 205, 52, 68, 205, 186, 82, 29, 242, 126,
			23, 124, 100, 20, 30, 171, 244, 76, 117, 180, 77, 241, 203, 12,
			3, 199, 82, 197, 37, 67, 129, 41, 42, 171, 233, 137, 17, 70,
			158, 199, 231, 237, 128, 239, 8, 233, 70, 94, 87, 204, 176, 1,
			161, 42, 101, 92, 26, 59, 177, 59, 4, 232, 151, 77, 129, 209,
			46, 52, 109, 51, 173, 81, 56, 26, 121, 49, 239, 220, 168, 109,
			61, 175, 128, 31, 182, 247, 36, 104, 222, 23, 192, 161, 31, 133,
			163, 12, 24, 65, 32, 157, 25, 10, 51, 242, 252, 234, 26, 125,
			73, 1, 179, 24, 121, 1, 87, 236, 103, 39, 192, 140, 117, 84,
			58, 222, 84, 113, 249, 157, 68, 198, 6, 108, 198, 213, 66, 240,
			234, 138, 161, 48, 35, 47, 176, 50, 253, 115, 237, 7, 57, 70,
			94, 196, 107, 246, 31, 35, 126, 219, 137, 160, 60, 66, 186, 74,
			226, 113, 18, 43, 254, 67, 225, 64, 82, 239, 123, 194, 239, 73,
			5, 86, 235, 170, 149, 10, 187, 224, 131, 83, 42, 83, 62, 18,
			82, 58, 3, 81, 229, 45, 200, 121, 199, 142, 239, 245, 248, 201,
			80, 232, 8, 126, 176, 245, 132, 130, 52, 211, 120, 114, 29, 116,
			220, 147, 148, 75, 225, 43, 7, 207, 244, 200, 33, 192, 90, 52,
			20, 102, 228, 197, 242, 42, 253, 80, 235, 145, 103, 228, 101, 92,
			182, 255, 236, 215, 211, 35, 221, 145, 89, 177, 10, 25, 27, 178,
			53, 125, 24, 43, 228, 190, 72, 55, 82, 224, 186, 160, 147, 202,
			235, 143, 242, 143, 127, 60, 123, 71, 26, 35, 245, 156, 216, 161,
			58, 163, 118, 141, 196, 76, 167, 60, 2, 220, 203, 134, 194, 140,
			188, 92, 100, 244, 83, 173, 211, 60, 35, 223, 192, 171, 246, 95,
			34, 190, 231, 201, 88, 245, 66, 146, 143, 194, 30, 228, 205, 94,
			26, 50, 70, 10, 232, 59, 114, 6, 158, 171, 183, 77, 187, 97,
			215, 145, 80, 205, 189, 64, 121, 99, 18, 196, 146, 123, 129, 20,
			81, 44, 122, 219, 61, 225, 11, 80, 103, 172, 140, 227, 127, 254,
			115, 163, 143, 59, 184, 135, 207, 109, 30, 129, 30, 198, 255, 230,
			49, 35, 223, 96, 21, 250, 75, 173, 227, 2, 35, 117, 204, 236,
			191, 251, 31, 213, 49, 61, 228, 71, 235, 118, 226, 72, 250, 152,
			131, 118, 226, 20, 138, 58, 223, 12, 165, 6, 253, 152, 19, 94,
			64, 160, 225, 146, 161, 48, 35, 245, 66, 137, 254, 182, 82, 126,
			145, 145, 6, 182, 237, 155, 188, 153, 86, 186, 163, 244, 170, 125,
			228, 245, 54, 185, 186, 127, 42, 165, 205, 197, 210, 248, 175, 58,
			152, 72, 244, 69, 36, 2, 23, 212, 79, 95, 147, 153, 212, 69,
			4, 172, 203, 134, 194, 140, 52, 214, 46, 208, 150, 146, 74, 25,
			185, 142, 47, 216, 215, 50, 169, 201, 184, 231, 196, 66, 37, 165,
			73, 86, 133, 172, 236, 72, 222, 78, 23, 14, 213, 30, 184, 112,
			112, 17, 196, 94, 236, 137, 137, 48, 138, 128, 163, 73, 91, 20,
			51, 114, 125, 117, 157, 190, 172, 132, 45, 49, 210, 196, 204, 174,
			102, 194, 210, 4, 33, 185, 35, 101, 232, 122, 14, 216, 56, 235,
			229, 117, 14, 205, 24, 47, 33, 120, 59, 163, 48, 35, 205, 229,
			18, 253, 129, 246, 156, 115, 140, 236, 225, 11, 246, 25, 202, 88,
			79, 183, 15, 208, 61, 31, 59, 158, 175, 218, 115, 199, 213, 29,
			170, 105, 113, 211, 230, 18, 76, 235, 197, 210, 68, 47, 205, 194,
			183, 202, 155, 131, 32, 132, 178, 60, 219, 180, 153, 86, 42, 1,
			167, 136, 61, 215, 153, 118, 243, 115, 8, 0, 153, 170, 114, 14,
			51, 178, 87, 49, 125, 192, 50, 35, 45, 188, 110, 7, 159, 13,
			213, 64, 123, 36, 226, 105, 92, 244, 115, 3, 91, 70, 32, 221,
			228, 205, 101, 204, 72, 171, 124, 158, 254, 12, 43, 100, 5, 70,
			94, 197, 204, 254, 41, 206, 160, 233, 86, 75, 157, 122, 79, 59,
			156, 234, 157, 125, 223, 20, 186, 126, 232, 251, 225, 137, 114, 187,
			72, 64, 255, 37, 174, 82, 190, 53, 109, 88, 79, 242, 112, 44,
			130, 108, 53, 3, 57, 3, 240, 129, 199, 208, 150, 67, 116, 137,
			32, 189, 253, 76, 49, 28, 133, 17, 128, 114, 69, 16, 235, 64,
			212, 79, 125, 71, 198, 169, 247, 194, 157, 110, 82, 158, 244, 91,
			106, 118, 178, 201, 189, 170, 168, 234, 38, 62, 93, 63, 9, 19,
			191, 199, 229, 48, 60, 225, 201, 120, 170, 187, 76, 252, 24, 58,
			27, 96, 101, 12, 208, 61, 189, 42, 133, 223, 207, 108, 89, 64,
			96, 175, 140, 194, 140, 188, 186, 92, 162, 127, 162, 61, 114, 133,
			145, 215, 240, 5, 251, 109, 196, 59, 175, 123, 99, 197, 117, 36,
			162, 129, 80, 7, 155, 90, 50, 224, 147, 171, 123, 149, 239, 134,
			17, 239, 131, 18, 163, 240, 88, 221, 175, 116, 159, 45, 55, 21,
			94, 154, 190, 52, 130, 170, 174, 154, 27, 184, 255, 234, 203, 97,
			63, 22, 193, 38, 63, 25, 122, 238, 80, 37, 93, 63, 60, 81,
			206, 98, 56, 232, 152, 234, 122, 3, 30, 71, 98, 42, 86, 87,
			16, 128, 52, 206, 176, 130, 25, 121, 173, 188, 78, 191, 175, 21,
			40, 50, 242, 109, 188, 102, 159, 206, 248, 130, 238, 74, 227, 199,
			41, 97, 64, 184, 78, 0, 24, 19, 169, 251, 186, 88, 248, 254,
			196, 75, 51, 183, 200, 28, 75, 183, 112, 78, 239, 116, 170, 243,
			205, 64, 22, 17, 0, 49, 85, 177, 136, 25, 249, 118, 113, 149,
			254, 158, 6, 89, 98, 228, 8, 87, 236, 100, 6, 228, 137, 232,
			30, 249, 94, 240, 186, 156, 130, 168, 50, 181, 134, 8, 183, 96,
			209, 7, 39, 10, 39, 229, 205, 11, 84, 119, 214, 245, 2, 29,
			135, 202, 98, 179, 67, 47, 211, 160, 164, 35, 175, 12, 96, 9,
			1, 8, 147, 212, 75, 152, 145, 163, 66, 153, 94, 85, 248, 24,
			35, 14, 46, 217, 91, 25, 190, 113, 8, 151, 50, 207, 241, 225,
			112, 186, 190, 24, 201, 199, 38, 60, 134, 224, 229, 121, 67, 97,
			70, 28, 90, 164, 67, 197, 183, 204, 72, 15, 219, 246, 239, 204,
			232, 221, 79, 124, 223, 148, 199, 52, 179, 106, 230, 250, 38, 186,
			149, 93, 230, 211, 61, 253, 48, 140, 33, 187, 107, 167, 167, 42,
			189, 67, 198, 83, 238, 104, 80, 148, 17, 136, 50, 62, 82, 198,
			140, 244, 202, 23, 232, 247, 20, 138, 10, 35, 3, 252, 132, 125,
			119, 162, 93, 34, 135, 220, 21, 81, 12, 98, 224, 194, 56, 157,
			218, 210, 216, 154, 145, 50, 157, 91, 165, 55, 8, 32, 17, 40,
			38, 105, 18, 19, 1, 184, 90, 150, 1, 204, 93, 58, 69, 83,
			65, 0, 192, 20, 182, 10, 102, 100, 176, 246, 69, 250, 45, 133,
			109, 149, 17, 15, 95, 176, 119, 50, 108, 166, 68, 234, 218, 38,
			238, 197, 80, 70, 125, 30, 71, 142, 251, 58, 132, 155, 60, 149,
			49, 28, 135, 35, 249, 65, 186, 214, 236, 205, 152, 98, 21, 1,
			207, 130, 161, 48, 35, 94, 105, 157, 254, 161, 246, 196, 53, 70,
			124, 124, 209, 254, 125, 52, 115, 36, 96, 105, 39, 232, 77, 121,
			34, 172, 238, 66, 183, 175, 202, 103, 95, 93, 222, 178, 134, 243,
			129, 126, 139, 254, 186, 141, 242, 195, 253, 214, 26, 2, 92, 198,
			68, 107, 152, 17, 127, 237, 9, 74, 41, 182, 230, 152, 21, 206,
			37, 106, 254, 98, 193, 181, 53, 92, 88, 167, 79, 80, 203, 82,
			55, 218, 187, 184, 108, 23, 211, 192, 230, 122, 68, 93, 165, 116,
			153, 230, 224, 177, 197, 172, 187, 56, 180, 129, 39, 144, 57, 216,
			62, 111, 40, 196, 200, 221, 133, 130, 161, 8, 35, 119, 75, 140,
			110, 43, 190, 136, 145, 24, 127, 201, 222, 224, 119, 160, 121, 250,
			149, 119, 122, 253, 58, 178, 224, 141, 140, 202, 51, 18, 47, 173,
			25, 10, 184, 157, 191, 104, 40, 194, 72, 252, 255, 55, 232, 87,
			41, 220, 112, 114, 247, 230, 126, 128, 144, 125, 153, 167, 179, 206,
			44, 26, 192, 220, 15, 79, 61, 171, 234, 90, 111, 1, 199, 123,
			11, 54, 248, 182, 133, 192, 8, 247, 241, 170, 125, 215, 140, 113,
			210, 92, 186, 201, 69, 117, 80, 229, 27, 230, 131, 208, 182, 140,
			220, 13, 62, 118, 162, 88, 165, 16, 51, 120, 200, 190, 23, 61,
			52, 115, 152, 126, 113, 251, 217, 237, 145, 35, 99, 17, 81, 133,
			20, 238, 254, 169, 149, 145, 178, 242, 125, 124, 239, 139, 74, 65,
			164, 172, 124, 95, 223, 223, 129, 66, 140, 220, 95, 44, 26, 138,
			48, 114, 191, 92, 1, 199, 183, 16, 88, 249, 13, 92, 178, 119,
			212, 96, 54, 18, 125, 165, 180, 78, 203, 113, 200, 93, 80, 27,
			90, 68, 165, 119, 38, 86, 109, 26, 133, 50, 158, 158, 234, 101,
			88, 144, 197, 172, 55, 240, 253, 213, 84, 30, 202, 129, 8, 131,
			5, 204, 246, 198, 226, 57, 67, 17, 70, 222, 88, 41, 210, 247,
			176, 2, 131, 153, 245, 38, 194, 235, 246, 219, 56, 179, 99, 198,
			191, 74, 41, 223, 85, 189, 181, 42, 228, 10, 153, 238, 204, 39,
			99, 197, 8, 218, 9, 254, 140, 170, 14, 234, 7, 36, 7, 245,
			67, 77, 19, 182, 190, 182, 7, 255, 125, 125, 235, 107, 251, 206,
			72, 124, 157, 95, 130, 195, 217, 76, 159, 213, 195, 158, 216, 210,
			45, 235, 101, 245, 138, 46, 44, 84, 21, 218, 16, 170, 128, 220,
			228, 110, 24, 168, 201, 90, 138, 174, 23, 186, 242, 63, 59, 63,
			210, 70, 221, 114, 195, 32, 142, 66, 95, 79, 143, 244, 218, 17,
			228, 194, 65, 24, 121, 66, 210, 153, 163, 46, 40, 155, 97, 139,
			229, 222, 68, 248, 141, 82, 106, 110, 156, 83, 86, 91, 48, 36,
			2, 114, 177, 108, 72, 2, 228, 218, 121, 186, 171, 44, 76, 152,
			245, 22, 194, 171, 246, 75, 70, 133, 180, 253, 231, 94, 15, 10,
			77, 223, 19, 145, 30, 41, 58, 50, 189, 68, 132, 17, 159, 220,
			44, 170, 148, 174, 40, 182, 196, 98, 249, 183, 16, 126, 19, 173,
			167, 114, 72, 78, 113, 54, 40, 8, 2, 114, 177, 104, 72, 37,
			183, 92, 161, 47, 81, 108, 97, 150, 255, 33, 154, 251, 41, 66,
			246, 51, 60, 27, 27, 126, 102, 204, 45, 81, 98, 129, 110, 63,
			68, 11, 95, 160, 5, 106, 89, 144, 147, 242, 63, 66, 248, 29,
			68, 148, 20, 12, 121, 201, 250, 17, 154, 95, 166, 187, 52, 15,
			36, 158, 99, 214, 219, 200, 42, 218, 95, 229, 233, 103, 64, 62,
			18, 78, 32, 39, 179, 92, 200, 35, 83, 211, 220, 174, 128, 212,
			158, 186, 150, 210, 117, 94, 243, 65, 138, 209, 210, 100, 1, 195,
			66, 97, 133, 182, 83, 73, 136, 89, 127, 132, 44, 102, 95, 227,
			217, 55, 198, 105, 89, 189, 80, 232, 10, 165, 102, 182, 159, 83,
			42, 210, 76, 151, 39, 11, 24, 22, 138, 37, 250, 74, 42, 21,
			51, 235, 199, 200, 178, 237, 231, 248, 67, 159, 50, 83, 233, 102,
			250, 106, 102, 198, 106, 234, 54, 37, 3, 128, 255, 24, 89, 171,
			147, 5, 197, 115, 253, 2, 253, 13, 48, 49, 88, 240, 93, 132,
			43, 246, 83, 92, 127, 82, 133, 170, 161, 74, 171, 166, 210, 182,
			222, 11, 212, 220, 51, 245, 14, 12, 217, 40, 255, 238, 212, 201,
			224, 185, 188, 226, 179, 96, 72, 4, 228, 226, 138, 33, 9, 144,
			172, 76, 111, 42, 153, 136, 89, 63, 1, 153, 191, 197, 219, 234,
			131, 46, 200, 28, 38, 35, 39, 80, 113, 173, 175, 58, 190, 19,
			121, 253, 83, 176, 155, 113, 28, 175, 175, 239, 49, 112, 193, 16,
			189, 12, 10, 178, 88, 254, 39, 8, 191, 139, 42, 169, 48, 148,
			83, 236, 13, 20, 164, 164, 101, 80, 16, 1, 146, 149, 233, 59,
			136, 98, 139, 176, 252, 123, 104, 238, 231, 8, 217, 127, 0, 133,
			122, 210, 156, 56, 221, 48, 81, 115, 241, 52, 52, 170, 15, 143,
			149, 255, 139, 83, 229, 236, 90, 110, 18, 3, 80, 91, 208, 32,
			233, 80, 128, 0, 123, 15, 45, 148, 233, 223, 32, 106, 89, 4,
			14, 234, 125, 132, 153, 253, 1, 82, 153, 92, 53, 119, 234, 222,
			159, 30, 145, 158, 98, 234, 150, 1, 238, 231, 94, 159, 247, 204,
			244, 54, 203, 1, 211, 237, 215, 212, 196, 42, 253, 226, 147, 246,
			13, 15, 125, 147, 85, 49, 155, 246, 233, 224, 7, 144, 183, 244,
			16, 109, 250, 5, 189, 205, 136, 74, 247, 165, 199, 68, 148, 199,
			188, 143, 240, 123, 104, 85, 29, 4, 129, 10, 6, 10, 45, 24,
			18, 1, 185, 184, 108, 72, 2, 100, 177, 68, 255, 69, 107, 143,
			152, 245, 1, 194, 101, 251, 151, 90, 251, 116, 252, 209, 235, 169,
			79, 87, 70, 125, 62, 86, 77, 157, 76, 103, 199, 144, 121, 29,
			87, 207, 105, 162, 48, 25, 12, 255, 119, 204, 35, 51, 251, 192,
			26, 116, 106, 141, 91, 159, 105, 46, 240, 234, 15, 16, 126, 31,
			177, 212, 32, 224, 213, 31, 76, 204, 133, 148, 65, 22, 11, 134,
			36, 64, 150, 24, 253, 0, 43, 115, 97, 102, 125, 136, 240, 211,
			112, 117, 175, 113, 223, 147, 177, 241, 146, 236, 47, 21, 102, 77,
			40, 228, 227, 220, 232, 17, 32, 39, 179, 191, 41, 101, 82, 227,
			152, 239, 132, 201, 96, 32, 100, 76, 161, 19, 114, 70, 34, 214,
			55, 124, 41, 226, 42, 239, 204, 34, 80, 69, 61, 253, 8, 166,
			219, 191, 153, 193, 10, 36, 0, 128, 68, 179, 108, 126, 43, 236,
			121, 253, 83, 158, 254, 33, 132, 158, 247, 15, 161, 143, 87, 51,
			4, 157, 28, 192, 1, 164, 120, 72, 91, 125, 129, 87, 165, 94,
			49, 53, 29, 13, 129, 146, 11, 6, 203, 200, 28, 144, 75, 37,
			67, 34, 32, 217, 134, 33, 9, 144, 79, 62, 69, 127, 174, 125,
			147, 48, 235, 35, 132, 215, 236, 159, 105, 223, 204, 6, 114, 255,
			55, 2, 19, 10, 253, 71, 8, 127, 136, 158, 78, 213, 131, 66,
			255, 209, 196, 211, 32, 15, 125, 132, 22, 141, 45, 136, 210, 182,
			178, 10, 247, 28, 108, 89, 44, 255, 49, 154, 251, 5, 66, 246,
			247, 31, 153, 63, 179, 59, 236, 127, 119, 246, 156, 249, 38, 167,
			137, 169, 220, 105, 33, 102, 125, 140, 22, 24, 221, 160, 150, 101,
			65, 234, 252, 4, 146, 71, 69, 157, 143, 251, 192, 45, 6, 204,
			96, 169, 252, 244, 9, 194, 31, 167, 101, 196, 82, 249, 233, 19,
			132, 231, 13, 137, 128, 92, 40, 24, 146, 0, 89, 98, 244, 146,
			146, 128, 152, 245, 41, 244, 181, 182, 146, 160, 102, 75, 179, 159,
			126, 50, 57, 16, 216, 159, 34, 252, 9, 42, 167, 156, 80, 94,
			189, 187, 108, 72, 197, 170, 144, 61, 37, 64, 174, 157, 239, 230,
			213, 95, 145, 189, 240, 31, 1, 0, 0, 255, 255, 116, 246, 171,
			233, 129, 38, 0, 0},
	)
}

// FileDescriptorSet returns a descriptor set for this proto package, which
// includes all defined services, and all transitive dependencies.
//
// Will not return nil.
//
// Do NOT modify the returned descriptor.
func FileDescriptorSet() *descriptor.FileDescriptorSet {
	// We just need ONE of the service names to look up the FileDescriptorSet.
	ret, err := discovery.GetDescriptorSet("gerrit.Gerrit")
	if err != nil {
		panic(err)
	}
	return ret
}
