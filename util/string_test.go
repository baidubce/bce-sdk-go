// nolint
package util

import (
	"bytes"
	"hash/crc32"
	"io"
	"os"
	"strconv"
	"testing"
	"time"
)

func TestCalculateContentCrc32(t *testing.T) {
	// case1: normal string content
	testData1 := "hello world"
	reader1 := bytes.NewReader([]byte(testData1))
	crc32Result1, err := CalculateContentCrc32(reader1, int64(len(testData1)))
	ExpectEqual(t, nil, err)

	// manually calculate expected crc32
	expectedHash1 := crc32.NewIEEE()
	expectedHash1.Write([]byte(testData1))
	expectedCrc32_1 := strconv.FormatUint(uint64(expectedHash1.Sum32()), 10)
	ExpectEqual(t, expectedCrc32_1, crc32Result1)

	// case2: empty content
	testData2 := ""
	reader2 := bytes.NewReader([]byte(testData2))
	crc32Result2, err := CalculateContentCrc32(reader2, int64(len(testData2)))
	ExpectEqual(t, nil, err)

	expectedHash2 := crc32.NewIEEE()
	expectedHash2.Write([]byte(testData2))
	expectedCrc32_2 := strconv.FormatUint(uint64(expectedHash2.Sum32()), 10)
	ExpectEqual(t, expectedCrc32_2, crc32Result2)

	// case3: binary content
	testData3 := []byte{0x00, 0x01, 0x02, 0x03, 0xFF, 0xFE, 0xFD}
	reader3 := bytes.NewReader(testData3)
	crc32Result3, err := CalculateContentCrc32(reader3, int64(len(testData3)))
	ExpectEqual(t, nil, err)

	expectedHash3 := crc32.NewIEEE()
	expectedHash3.Write(testData3)
	expectedCrc32_3 := strconv.FormatUint(uint64(expectedHash3.Sum32()), 10)
	ExpectEqual(t, expectedCrc32_3, crc32Result3)

	// case4: size mismatch - size larger than actual data
	testData4 := "short"
	reader4 := bytes.NewReader([]byte(testData4))
	_, err = CalculateContentCrc32(reader4, 100) // size larger than actual data
	ExpectEqual(t, true, err != nil)
	ExpectEqual(t, true, err.Error() != "")

	// case5: large content (1MB)
	testData5 := make([]byte, 1024*1024)
	for i := range testData5 {
		testData5[i] = byte(i % 256)
	}
	reader5 := bytes.NewReader(testData5)
	crc32Result5, err := CalculateContentCrc32(reader5, int64(len(testData5)))
	ExpectEqual(t, nil, err)

	expectedHash5 := crc32.NewIEEE()
	expectedHash5.Write(testData5)
	expectedCrc32_5 := strconv.FormatUint(uint64(expectedHash5.Sum32()), 10)
	ExpectEqual(t, expectedCrc32_5, crc32Result5)

	// case6: size is zero
	testData6 := "some content"
	reader6 := bytes.NewReader([]byte(testData6))
	crc32Result6, err := CalculateContentCrc32(reader6, 0)
	ExpectEqual(t, nil, err)

	expectedHash6 := crc32.NewIEEE()
	// no data written for size 0
	expectedCrc32_6 := strconv.FormatUint(uint64(expectedHash6.Sum32()), 10)
	ExpectEqual(t, expectedCrc32_6, crc32Result6)

	// case7: test with specific known crc32 value
	testData7 := "test"
	reader7 := bytes.NewReader([]byte(testData7))
	crc32Result7, err := CalculateContentCrc32(reader7, int64(len(testData7)))
	ExpectEqual(t, nil, err)
	// CRC32-IEEE for "test" is 3632233996
	ExpectEqual(t, "3632233996", crc32Result7)
}

func TestCalculateContentCrc32FromFile(t *testing.T) {
	// case1: normal file read from beginning
	testFileName1 := "/tmp/test-crc32-file-" + strconv.FormatInt(time.Now().UnixMicro(), 10) + "-1"
	testData1 := []byte("hello world from file")
	err := os.WriteFile(testFileName1, testData1, 0644)
	ExpectEqual(t, nil, err)
	defer os.Remove(testFileName1)

	file1, err := os.Open(testFileName1)
	ExpectEqual(t, nil, err)
	defer file1.Close()

	crc32Result1, err := CalculateContentCrc32FromFile(file1, 0, int64(len(testData1)))
	ExpectEqual(t, nil, err)

	expectedHash1 := crc32.NewIEEE()
	expectedHash1.Write(testData1)
	expectedCrc32_1 := strconv.FormatUint(uint64(expectedHash1.Sum32()), 10)
	ExpectEqual(t, expectedCrc32_1, crc32Result1)

	// case2: read with offset
	testFileName2 := "/tmp/test-crc32-file-" + strconv.FormatInt(time.Now().UnixMicro(), 10) + "-2"
	testData2 := []byte("0123456789abcdef")
	err = os.WriteFile(testFileName2, testData2, 0644)
	ExpectEqual(t, nil, err)
	defer os.Remove(testFileName2)

	file2, err := os.Open(testFileName2)
	ExpectEqual(t, nil, err)
	defer file2.Close()

	// read "abcdef" (offset=10, size=6)
	crc32Result2, err := CalculateContentCrc32FromFile(file2, 10, 6)
	ExpectEqual(t, nil, err)

	expectedHash2 := crc32.NewIEEE()
	expectedHash2.Write([]byte("abcdef"))
	expectedCrc32_2 := strconv.FormatUint(uint64(expectedHash2.Sum32()), 10)
	ExpectEqual(t, expectedCrc32_2, crc32Result2)

	// case3: verify file position is restored after call
	testFileName3 := "/tmp/test-crc32-file-" + strconv.FormatInt(time.Now().UnixMicro(), 10) + "-3"
	testData3 := []byte("ABCDEFGHIJ")
	err = os.WriteFile(testFileName3, testData3, 0644)
	ExpectEqual(t, nil, err)
	defer os.Remove(testFileName3)

	file3, err := os.Open(testFileName3)
	ExpectEqual(t, nil, err)
	defer file3.Close()

	// set file position to 5
	file3.Seek(5, io.SeekStart)
	originalPos, _ := file3.Seek(0, io.SeekCurrent)
	ExpectEqual(t, int64(5), originalPos)

	// calculate crc32 from offset 0
	_, err = CalculateContentCrc32FromFile(file3, 0, 5)
	ExpectEqual(t, nil, err)

	// verify file position is restored
	restoredPos, _ := file3.Seek(0, io.SeekCurrent)
	ExpectEqual(t, int64(5), restoredPos)

	// case4: read entire file
	testFileName4 := "/tmp/test-crc32-file-" + strconv.FormatInt(time.Now().UnixMicro(), 10) + "-4"
	testData4 := make([]byte, 1024)
	for i := range testData4 {
		testData4[i] = byte(i % 256)
	}
	err = os.WriteFile(testFileName4, testData4, 0644)
	ExpectEqual(t, nil, err)
	defer os.Remove(testFileName4)

	file4, err := os.Open(testFileName4)
	ExpectEqual(t, nil, err)
	defer file4.Close()

	crc32Result4, err := CalculateContentCrc32FromFile(file4, 0, int64(len(testData4)))
	ExpectEqual(t, nil, err)

	expectedHash4 := crc32.NewIEEE()
	expectedHash4.Write(testData4)
	expectedCrc32_4 := strconv.FormatUint(uint64(expectedHash4.Sum32()), 10)
	ExpectEqual(t, expectedCrc32_4, crc32Result4)

	// case5: read partial file from middle
	testFileName5 := "/tmp/test-crc32-file-" + strconv.FormatInt(time.Now().UnixMicro(), 10) + "-5"
	testData5 := []byte("0123456789ABCDEFGHIJ")
	err = os.WriteFile(testFileName5, testData5, 0644)
	ExpectEqual(t, nil, err)
	defer os.Remove(testFileName5)

	file5, err := os.Open(testFileName5)
	ExpectEqual(t, nil, err)
	defer file5.Close()

	// read "56789A" (offset=5, size=6)
	crc32Result5, err := CalculateContentCrc32FromFile(file5, 5, 6)
	ExpectEqual(t, nil, err)

	expectedHash5 := crc32.NewIEEE()
	expectedHash5.Write([]byte("56789A"))
	expectedCrc32_5 := strconv.FormatUint(uint64(expectedHash5.Sum32()), 10)
	ExpectEqual(t, expectedCrc32_5, crc32Result5)

	// case6: size larger than available data from offset
	testFileName6 := "/tmp/test-crc32-file-" + strconv.FormatInt(time.Now().UnixMicro(), 10) + "-6"
	testData6 := []byte("short")
	err = os.WriteFile(testFileName6, testData6, 0644)
	ExpectEqual(t, nil, err)
	defer os.Remove(testFileName6)

	file6, err := os.Open(testFileName6)
	ExpectEqual(t, nil, err)
	defer file6.Close()

	// try to read 100 bytes but file only has 5 bytes
	_, err = CalculateContentCrc32FromFile(file6, 0, 100)
	ExpectEqual(t, true, err != nil)

	// case7: read with zero size
	testFileName7 := "/tmp/test-crc32-file-" + strconv.FormatInt(time.Now().UnixMicro(), 10) + "-7"
	testData7 := []byte("some content")
	err = os.WriteFile(testFileName7, testData7, 0644)
	ExpectEqual(t, nil, err)
	defer os.Remove(testFileName7)

	file7, err := os.Open(testFileName7)
	ExpectEqual(t, nil, err)
	defer file7.Close()

	crc32Result7, err := CalculateContentCrc32FromFile(file7, 0, 0)
	ExpectEqual(t, nil, err)

	expectedHash7 := crc32.NewIEEE()
	expectedCrc32_7 := strconv.FormatUint(uint64(expectedHash7.Sum32()), 10)
	ExpectEqual(t, expectedCrc32_7, crc32Result7)

	// case8: multiple consecutive calls on same file
	testFileName8 := "/tmp/test-crc32-file-" + strconv.FormatInt(time.Now().UnixMicro(), 10) + "-8"
	testData8 := []byte("PART1PART2PART3")
	err = os.WriteFile(testFileName8, testData8, 0644)
	ExpectEqual(t, nil, err)
	defer os.Remove(testFileName8)

	file8, err := os.Open(testFileName8)
	ExpectEqual(t, nil, err)
	defer file8.Close()

	// read PART1
	crc32Part1, err := CalculateContentCrc32FromFile(file8, 0, 5)
	ExpectEqual(t, nil, err)
	expectedHashPart1 := crc32.NewIEEE()
	expectedHashPart1.Write([]byte("PART1"))
	ExpectEqual(t, strconv.FormatUint(uint64(expectedHashPart1.Sum32()), 10), crc32Part1)

	// read PART2
	crc32Part2, err := CalculateContentCrc32FromFile(file8, 5, 5)
	ExpectEqual(t, nil, err)
	expectedHashPart2 := crc32.NewIEEE()
	expectedHashPart2.Write([]byte("PART2"))
	ExpectEqual(t, strconv.FormatUint(uint64(expectedHashPart2.Sum32()), 10), crc32Part2)

	// read PART3
	crc32Part3, err := CalculateContentCrc32FromFile(file8, 10, 5)
	ExpectEqual(t, nil, err)
	expectedHashPart3 := crc32.NewIEEE()
	expectedHashPart3.Write([]byte("PART3"))
	ExpectEqual(t, strconv.FormatUint(uint64(expectedHashPart3.Sum32()), 10), crc32Part3)
}
