package main

import (
	"fmt"
	"github.com/softclub-go-0-0/auth"
	"github.com/softclub-go-0-0/siyovush-analytics/pkg/analytics"
)

func main() {
	var user, password string
	fmt.Println("Enter your username and password:")
	fmt.Scan(&user, &password)
	if !auth.Validate(user, password) {
		fmt.Println("\nAccess Denied!")
		return
	}
	fmt.Println("\nAccess granted!")
	fmt.Println("\nHere are the statistics of our platform:")

	fmt.Printf("\tTotal posts:\t%d\n", analytics.TotalPosts())
	fmt.Println("\tAll the posts:", analytics.FilterByVisibility(true))

	fmt.Printf("\tTotal posts filtered by %t visibility:\t%d\n", true, len(analytics.FilterByVisibility(true)))
	fmt.Println("\tPosts with IsVisible == true:", analytics.FilterByVisibility(true))

	fmt.Printf("\tTotal posts filtered by %t visibility:\t%d\n", false, len(analytics.FilterByVisibility(false)))
	fmt.Println("\tPosts with IsVisible == false:", analytics.FilterByVisibility(false))

	fmt.Printf("\tTotal posts filtered by %d rating:\t%d\n", 4, len(analytics.FilterByRating(4)))
	fmt.Println("\tPosts with Rating == 4:", analytics.FilterByRating(4))

	fmt.Printf("\tTotal posts filtered by %q tag:\t%d\n", "mollit", len(analytics.FilterByTag("mollit")))
	fmt.Println("\tPosts with at least one Tag == mollit:", analytics.FilterByTag("mollit"))

	fmt.Printf("\tTotal posts filtered by %q, %q and %q tags:\t%d\n", "nostrud", "eiusmod", "enim", len(analytics.FilterByTags([]string{"nostrud", "eiusmod", "enim"})))
	fmt.Println("\tPosts with at least one Tag == nostrud or eiusmod or enim:", analytics.FilterByTags([]string{"nostrud", "eiusmod", "enim"}))

	fmt.Printf("\tTotal posts filtered by %q, %q, %q and %q tags:\t%d\n", "nostrud", "eiusmod", "enim", "mollit", len(analytics.FilterByTags([]string{"nostrud", "eiusmod", "enim", "mollit"})))
	fmt.Println("\tPosts with at least one Tag == nostrud or eiusmod or enim or mollit:", analytics.FilterByTags([]string{"nostrud", "eiusmod", "enim", "mollit"}))

	fmt.Printf("\tStatistics of posts by rating:\t%v\n", analytics.StatisticsByRating())

	fmt.Printf("\tStatistics of posts by content length:\t%v\n", analytics.StatisticsByContentLength())
}
