# count-negative-numbers-in-a-sorted-matrix

## 題目解讀：

### 題目來源:
[count-negative-numbers-in-a-sorted-matrix](https://leetcode.com/problems/count-negative-numbers-in-a-sorted-matrix/)

### 原文:
Given a m * n matrix grid which is sorted in non-increasing order both row-wise and column-wise. 

Return the number of negative numbers in grid.

### 解讀：

給定一個m by n 的矩陣 grid 每個元素的在 列跟行值都是以遞減的方式排列

找出所有負數的個數

## 初步解法:
### 初步觀察:

因為 矩陣 grid 每個元素的在 列跟行值都是以遞減的方式排列

所需要從右下往左上找是最佳搜尋順

由下而上 由右而左

依序找尋 小於0的元素

### 初步設計:
given a m x n matrix grid,( [][]int)

Step 0: let idxY = len(grid) - 1, count = 0

Step 1: if idxY < 0 go to step 8

Step 2: let idxX = len(grid[idxY]) -1

Step 3: if idxX < 0 go to step 7

Step 4: if grid[idxY][idxX] < 0 set connt += 1

Step 5: if grid[idxY][idxX] set idxY = idxY - 1 go to Step 1

Step 6: idxX = idxX - 1 go to step 3

Step 7: idxY = idxY - 1 go to step 1

Step 8: return count

## 遇到的困難
### 題目上理解的問題
因為英文不是筆者母語

所以在題意解讀上 容易被英文用詞解讀給搞模糊

### pseudo code撰寫

一開始不習慣把pseudo code寫下來

因此 不太容易把自己的code做解析

### golang table driven test不熟
對於table driven test還不太熟析

所以對於寫test還是耗費不少時間
## 我的github source code
```golang
package count_neg

func countNegatives(grid [][]int) int {
	count := 0
	lenY := len(grid)
	lenX := len(grid[0])
	for idxY := lenY - 1; idxY >= 0; idxY-- {
		for idxX := lenX - 1; idxX >= 0; idxX-- {
			if grid[idxY][idxX] < 0 {
				count += 1
			} else {
				continue
			}
		}
	}
	return count
}

```
## 測資的撰寫
```golang
package count_neg

import "testing"

func Test_countNegatives(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				grid: [][]int{{4, 3, 2, -1}, {3, 2, 1, -1}, {1, 1, -1, -2}, {-1, -1, -2, -3}},
			},
			want: 8,
		},
		{
			name: "Example2",
			args: args{
				grid: [][]int{{3, 2}, {1, 0}},
			},
			want: 0,
		},
		{
			name: "Example3",
			args: args{
				grid: [][]int{{1, -1}, {-1, -1}},
			},
			want: 3,
		},
		{
			name: "Example4",
			args: args{
				grid: [][]int{{-1}},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countNegatives(tt.args.grid); got != tt.want {
				t.Errorf("countNegatives() = %v, want %v", got, tt.want)
			}
		})
	}
}

```
## my self record
[golang leetcode 30day 28th day](https://hackmd.io/fMOzB-TNSouKsSjtg_qGAg?view)
## 參考文章

[golang test](https://ithelp.ithome.com.tw/articles/10204692)