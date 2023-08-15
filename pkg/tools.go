package tools

import "im/internal/service/models"

//对消息进行快排
func QuickSort(l int, r int, a *[]models.Message) {
	if l >= r {
		return
	}
	x := (*a)[l].ID
	i := l - 1
	j := r + 1
	for i < j {
		for i++; (*a)[i].ID < x; {
			i++
		}
		for j--; (*a)[j].ID > x; {
			j--
		}
		if i < j {
			(*a)[i], (*a)[j] = (*a)[j], (*a)[i]
		}
	}
	QuickSort(l, j, a)
	QuickSort(j+1, r, a)
}
