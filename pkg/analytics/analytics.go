package analytics

import (
	"github.com/softclub-go-0-0/database"
	"math"
)

func FilterByVisibility(visibility bool) []database.Post {
	var filteredPosts []database.Post
	for _, post := range database.Posts {
		if post.IsVisible == visibility {
			filteredPosts = append(filteredPosts, post)
		}
	}
	return filteredPosts
}

func FilterByRating(rating int) []database.Post {
	var filteredPosts []database.Post
	for _, post := range database.Posts {
		if post.Rating == rating {
			filteredPosts = append(filteredPosts, post)
		}
	}
	return filteredPosts
}

func FilterByTag(filteringTag string) []database.Post {
	var filteredPosts []database.Post
	for _, post := range database.Posts {
		add := false
		for _, tag := range post.Tags {
			if tag == filteringTag {
				add = true
				break
			}
		}
		if add {
			filteredPosts = append(filteredPosts, post)
		}
	}
	return filteredPosts
}

func FilterByTags(filteringTags []string) []database.Post {
	var filteredPosts []database.Post
	for _, post := range database.Posts {
		add := false
		for _, tag := range post.Tags {
			for _, filteringTag := range filteringTags {
				if tag == filteringTag {
					add = true
					break
				}
			}
			if add {
				break
			}
		}
		if add {
			filteredPosts = append(filteredPosts, post)
		}
	}
	return filteredPosts
}

func TotalPosts() int {
	return len(database.Posts)
}

func StatisticsByRating() map[int]int {
	statistics := make(map[int]int)
	for i := 0; i <= 5; i++ {
		statistics[i] = 0
	}
	for _, post := range database.Posts {
		statistics[post.Rating]++
	}
	return statistics
}

func StatisticsByContentLength() map[string]int {
	statistics := make(map[string]int)
	statistics["min"] = math.MaxInt
	statistics["average"] = 0
	statistics["max"] = 0

	var total uint64

	for _, post := range database.Posts {
		if statistics["min"] > len(post.Content) {
			statistics["min"] = len(post.Content)
		}
		if statistics["max"] < len(post.Content) {
			statistics["max"] = len(post.Content)
		}
		total += uint64(len(post.Content))
	}

	statistics["average"] = int(total / uint64(len(database.Posts)))

	return statistics
}
