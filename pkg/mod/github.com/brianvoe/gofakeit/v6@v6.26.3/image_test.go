package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleImageURL() {
	Seed(11)
	fmt.Println(ImageURL(640, 480))

	// Output: https://picsum.photos/640/480
}

func ExampleFaker_ImageURL() {
	f := New(11)
	fmt.Println(f.ImageURL(640, 480))

	// Output: https://picsum.photos/640/480
}

func BenchmarkImageURL(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ImageURL(640, 480)
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.ImageURL(640, 480)
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.ImageURL(640, 480)
		}
	})
}

func ExampleImage() {
	Seed(11)
	fmt.Println(Image(1, 1))

	// Output: &{[89 176 195 255] 4 (0,0)-(1,1)}
}

func ExampleFaker_Image() {
	f := New(11)
	fmt.Println(f.Image(1, 1))

	// Output: &{[89 176 195 255] 4 (0,0)-(1,1)}
}

func BenchmarkImage(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Image(640, 480)
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.Image(640, 480)
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.Image(640, 480)
		}
	})
}

func ExampleImageJpeg() {
	Seed(11)
	fmt.Println(ImageJpeg(1, 1))

	// Output: [255 216 255 219 0 132 0 8 6 6 7 6 5 8 7 7 7 9 9 8 10 12 20 13 12 11 11 12 25 18 19 15 20 29 26 31 30 29 26 28 28 32 36 46 39 32 34 44 35 28 28 40 55 41 44 48 49 52 52 52 31 39 57 61 56 50 60 46 51 52 50 1 9 9 9 12 11 12 24 13 13 24 50 33 28 33 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 255 192 0 17 8 0 1 0 1 3 1 34 0 2 17 1 3 17 1 255 196 1 162 0 0 1 5 1 1 1 1 1 1 0 0 0 0 0 0 0 0 1 2 3 4 5 6 7 8 9 10 11 16 0 2 1 3 3 2 4 3 5 5 4 4 0 0 1 125 1 2 3 0 4 17 5 18 33 49 65 6 19 81 97 7 34 113 20 50 129 145 161 8 35 66 177 193 21 82 209 240 36 51 98 114 130 9 10 22 23 24 25 26 37 38 39 40 41 42 52 53 54 55 56 57 58 67 68 69 70 71 72 73 74 83 84 85 86 87 88 89 90 99 100 101 102 103 104 105 106 115 116 117 118 119 120 121 122 131 132 133 134 135 136 137 138 146 147 148 149 150 151 152 153 154 162 163 164 165 166 167 168 169 170 178 179 180 181 182 183 184 185 186 194 195 196 197 198 199 200 201 202 210 211 212 213 214 215 216 217 218 225 226 227 228 229 230 231 232 233 234 241 242 243 244 245 246 247 248 249 250 1 0 3 1 1 1 1 1 1 1 1 1 0 0 0 0 0 0 1 2 3 4 5 6 7 8 9 10 11 17 0 2 1 2 4 4 3 4 7 5 4 4 0 1 2 119 0 1 2 3 17 4 5 33 49 6 18 65 81 7 97 113 19 34 50 129 8 20 66 145 161 177 193 9 35 51 82 240 21 98 114 209 10 22 36 52 225 37 241 23 24 25 26 38 39 40 41 42 53 54 55 56 57 58 67 68 69 70 71 72 73 74 83 84 85 86 87 88 89 90 99 100 101 102 103 104 105 106 115 116 117 118 119 120 121 122 130 131 132 133 134 135 136 137 138 146 147 148 149 150 151 152 153 154 162 163 164 165 166 167 168 169 170 178 179 180 181 182 183 184 185 186 194 195 196 197 198 199 200 201 202 210 211 212 213 214 215 216 217 218 226 227 228 229 230 231 232 233 234 242 243 244 245 246 247 248 249 250 255 218 0 12 3 1 0 2 17 3 17 0 63 0 216 162 138 43 213 62 92 255 217]
}

func ExampleFaker_ImageJpeg() {
	f := New(11)
	fmt.Println(f.ImageJpeg(1, 1))

	// Output: [255 216 255 219 0 132 0 8 6 6 7 6 5 8 7 7 7 9 9 8 10 12 20 13 12 11 11 12 25 18 19 15 20 29 26 31 30 29 26 28 28 32 36 46 39 32 34 44 35 28 28 40 55 41 44 48 49 52 52 52 31 39 57 61 56 50 60 46 51 52 50 1 9 9 9 12 11 12 24 13 13 24 50 33 28 33 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 50 255 192 0 17 8 0 1 0 1 3 1 34 0 2 17 1 3 17 1 255 196 1 162 0 0 1 5 1 1 1 1 1 1 0 0 0 0 0 0 0 0 1 2 3 4 5 6 7 8 9 10 11 16 0 2 1 3 3 2 4 3 5 5 4 4 0 0 1 125 1 2 3 0 4 17 5 18 33 49 65 6 19 81 97 7 34 113 20 50 129 145 161 8 35 66 177 193 21 82 209 240 36 51 98 114 130 9 10 22 23 24 25 26 37 38 39 40 41 42 52 53 54 55 56 57 58 67 68 69 70 71 72 73 74 83 84 85 86 87 88 89 90 99 100 101 102 103 104 105 106 115 116 117 118 119 120 121 122 131 132 133 134 135 136 137 138 146 147 148 149 150 151 152 153 154 162 163 164 165 166 167 168 169 170 178 179 180 181 182 183 184 185 186 194 195 196 197 198 199 200 201 202 210 211 212 213 214 215 216 217 218 225 226 227 228 229 230 231 232 233 234 241 242 243 244 245 246 247 248 249 250 1 0 3 1 1 1 1 1 1 1 1 1 0 0 0 0 0 0 1 2 3 4 5 6 7 8 9 10 11 17 0 2 1 2 4 4 3 4 7 5 4 4 0 1 2 119 0 1 2 3 17 4 5 33 49 6 18 65 81 7 97 113 19 34 50 129 8 20 66 145 161 177 193 9 35 51 82 240 21 98 114 209 10 22 36 52 225 37 241 23 24 25 26 38 39 40 41 42 53 54 55 56 57 58 67 68 69 70 71 72 73 74 83 84 85 86 87 88 89 90 99 100 101 102 103 104 105 106 115 116 117 118 119 120 121 122 130 131 132 133 134 135 136 137 138 146 147 148 149 150 151 152 153 154 162 163 164 165 166 167 168 169 170 178 179 180 181 182 183 184 185 186 194 195 196 197 198 199 200 201 202 210 211 212 213 214 215 216 217 218 226 227 228 229 230 231 232 233 234 242 243 244 245 246 247 248 249 250 255 218 0 12 3 1 0 2 17 3 17 0 63 0 216 162 138 43 213 62 92 255 217]
}

func BenchmarkImageJpeg(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ImageJpeg(640, 480)
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.ImageJpeg(640, 480)
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.ImageJpeg(640, 480)
		}
	})
}

func ExampleImagePng() {
	Seed(11)
	fmt.Println(ImagePng(1, 1))

	// Output: [137 80 78 71 13 10 26 10 0 0 0 13 73 72 68 82 0 0 0 1 0 0 0 1 8 2 0 0 0 144 119 83 222 0 0 0 16 73 68 65 84 120 156 98 138 220 112 24 16 0 0 255 255 3 58 1 207 38 214 44 234 0 0 0 0 73 69 78 68 174 66 96 130]
}

func ExampleFaker_ImagePng() {
	f := New(11)
	fmt.Println(f.ImagePng(1, 1))

	// Output: [137 80 78 71 13 10 26 10 0 0 0 13 73 72 68 82 0 0 0 1 0 0 0 1 8 2 0 0 0 144 119 83 222 0 0 0 16 73 68 65 84 120 156 98 138 220 112 24 16 0 0 255 255 3 58 1 207 38 214 44 234 0 0 0 0 73 69 78 68 174 66 96 130]
}

func BenchmarkImagePng(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ImagePng(640, 480)
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.ImagePng(640, 480)
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.ImagePng(640, 480)
		}
	})
}