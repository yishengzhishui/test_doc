## 排序算法

[十大经典排序算法（动图演示）](https://www.cnblogs.com/onepixel/p/7674659.html)

[9 种经典排序算法可视化动画](https://www.bilibili.com/video/av25136272)

### 概述

十种常见排序算法可以分为两大类：

- **比较类排序**：通过比较来决定元素间的相对次序，由于其时间复杂度不能突破O(nlogn)，因此也称为非线性时间比较类排序。
- **非比较类排序**：不通过比较来决定元素间的相对次序，它可以突破基于比较排序的时间下界，以线性时间运行，因此也称为线性时间非比较类排序。

### 初级排序-O(n^2)

#### 1、选择排序（Selection Sort ）

每次找到最小（大）值，然后放到待排序数组的起始位置

```python
def selectionSort(arr):
    n = len(arr)
    for i in range(n):
        min_index = i
        for j in range(i, n):
            if arr[min_index] > arr[j]:
                min_index = j
        arr[i], arr[min_index] = arr[min_index], arr[i]
    return arr
```

#### 2、插入排序（Insertion Sort）

从前到后逐步构建有序序列，对于未排序的数据，在已排序序列中从后向前扫描，找到相应位置并插入

```python
def insertionSort(arr):
    n = len(arr)
    for i in range(1, n):
        value = arr[i]
        j = i -1
        while j >= 0 and value < arr[j]:
            arr[j+1] = arr[j]
            j -= 1
        arr[j+1] = value
    return arr
```

#### 3、冒泡排序（Bubble Sort）

嵌套循环，每次查看相邻的元素，如果逆序则交换。

```python
def bubbleSort(arr):
    n = len(arr)
    for i in range(n):
        for j in range(n-i-1): # 最大的元素会到最后
            if arr[j] > arr[j+1]:
                arr[j], arr[j+1] = arr[j+1], arr[j]
    return arr
```

### 高级排序-O(n*logn)

### 概述：

归并与快排具有相似性，但是顺序相反

归并：先排序左右子数组，然后合并两个有序子数组

快排：先调配出左右子数组，然后对左右子数组进行排序

#### 1、快速排序（Quick Sort）

数组取标杆pivot，将小元素放到pivot左边，大元素放到右侧，然后一次对右边和左边的子数组继续快排；最后达到整体有序

##### 代码示例1

```python
def quickSort(begin,end,arr):
    if begin >= end:
        return
    pivot = partition(begin,end,arr)
    quickSort(begin, pivot-1,arr)
    quickSort(pivot+1, end, arr)
  
  
def partition(begin,end, arr):
    # pivot标杆位置；counter：小于pivot的元素个数
    pivot, counter = end, begin
    for i in range(begin,end):
        if arr[i] < arr[pivot]:
            arr[counter], arr[i] = arr[i], arr[counter]
            counter += 1
    arr[pivot], arr[counter] = arr[counter], arr[pivot]
    return counter
        
        
```

##### 代码示例2:

```python
def quick_sort(begin, end, nums):
    if begin >= end:
        return
    pivot_index = partition(begin, end, nums)
    quick_sort(begin, pivot_index-1, nums)
    quick_sort(pivot_index+1, end, nums)
  
def partition(begin, end, nums):
    pivot = nums[begin]
    mark = begin
    for i in range(begin+1, end+1):
        if nums[i] < pivot:
            mark +=1
            nums[mark], nums[i] = nums[i], nums[mark]
    nums[begin], nums[mark] = nums[mark], nums[begin]
    return mark
```

#### 2、归并排序（Merge Sort）-分治

1. 把长度为n的输入序列分成两个长度为n/2的子序列；
2. 对这两个子序列分别采用归并排序；
3. 将两个排序好的子序列合并成一个最终的排序序列

```python
def mergeSort(arr, left, right):
    if right <= left: return
    mid = (left+right)>>1
    mergeSort(arr, left, mid)
    mergeSort(arr, mid+1,right)
    merge(arr, left,mid,right)
  
def merge(arr, left,mid,right): # 两个排序好的数组合并
    temp = []
    i, j = left, mid+1
  
    while i<=mid and j <=right:
        if arr[i] <= arr[j]:
            temp.append(arr[i])
            i +=1
        else:
            temp.append(arr[j])
            j += 1
        
    while i<= mid: # 如果 前一个数组未走完
        temp.append(arr[i])
        i += 1
   
    while j <= right: # 如果 后一个数组未走完
        temp.append(arr[j])
        j += 1
    
    arr[left:right+1] = temp
```

#### 3、堆排序（Heap Sort）⚠️

[python heapq堆队列算法](https://docs.python.org/zh-cn/3/library/heapq.html)

1. 数组元素依次建立小顶堆
2. 依次取堆顶元素，并删除

##### 代码示例1

```python
def heapify(arr, n, i): 
    largest = i  
    l = 2 * i + 1     # left = 2*i + 1 
    r = 2 * i + 2     # right = 2*i + 2 
  
    if l < n and arr[i] < arr[l]: 
        largest = l 
  
    if r < n and arr[largest] < arr[r]: 
        largest = r 
  
    if largest != i: 
        arr[i],arr[largest] = arr[largest],arr[i]  # 交换
  
        heapify(arr, n, largest) 
  
def heapSort(arr): 
    n = len(arr) 
  
    # Build a maxheap. 
    for i in range(n, -1, -1): 
        heapify(arr, n, i) 
  
    # 一个个交换元素
    for i in range(n-1, 0, -1): 
        arr[i], arr[0] = arr[0], arr[i]   # 交换
        heapify(arr, i, 0) 
```

##### 代码示例2

```python
def heapify(parent_index, length, nums):
    temp = nums[parent_index]
    child_index = 2*parent_index+1
    while child_index < length:
        if child_index+1 < length and nums[child_index+1] > nums[child_index]:
            child_index = child_index+1
        if temp > nums[child_index]:
            break
        nums[parent_index] = nums[child_index]
        parent_index = child_index
        child_index = 2*parent_index + 1
    nums[parent_index] = temp


def heapsort(nums):
    for i in range((len(nums)-2)//2, -1, -1):
        heapify(i, len(nums), nums)
    for j in range(len(nums)-1, 0, -1):
        nums[j], nums[0] = nums[0], nums[j]
        heapify(0, j, nums)
```

```python
# use heapq module
import heapq


def heapSort(arr):
    n = len(arr)
    queue = []

    heapq.heapify(arr)
    for i in range(n):
        queue.append(heapq.heappop(arr))
    return queue
```

### 特殊排序-O(n)-（理解原理）

#### 计数排序(Counting Sort)

计数排序要求输入的数据必须是有确定范围的整数。将输入的数据值转化为键存 储在额外开辟的数组空间中;然后依次把计数大于 1 的填充回原数组

#### 桶排序(Bucket Sort)

桶排序 (Bucket sort)的工作的原理:假设输入数据服从均匀分布，将数据分到有 限数量的桶里，每个桶再分别排序(有可能再使用别的排序算法或是以递归方式 继续使用桶排序进行排)。

#### 基数排序(Radix Sort)

基数排序是按照低位先排序，然后收集;再按照高位排序，然后再收集;依次类 推，直到最高位。有时候有些属性是有优先级顺序的，先按低优先级排序，再按 高优先级排序。
