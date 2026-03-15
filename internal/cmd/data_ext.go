package cmd

import (
	"github.com/alpacahq/cli/internal/api"
	"github.com/alpacahq/cli/internal/cmdutil"
	"github.com/spf13/cobra"
)

var newsCmd = fetchCmd("news", api.NewsOp, func(cmd *cobra.Command, args []string) (any, error) {
	params := newsParamsFromFlags(cmd)
	if cmdutil.Bool(cmd, "all") {
		max := cmdutil.Int(cmd, "max")
		var allNews []api.News
		for page := 0; page < maxPages; page++ {
			resp, err := dataClient.News(params)
			if err != nil {
				return nil, err
			}
			allNews = append(allNews, resp.News...)
			if max > 0 && len(allNews) >= max {
				allNews = allNews[:max]
				break
			}
			if resp.NextPageToken == "" {
				break
			}
			params.PageToken = resp.NextPageToken
		}
		return allNews, nil
	}

	return dataClient.News(params)
}, func(c *cobra.Command) {
	addPaginationFlags(c)
	c.Example = `  alpaca data news
  alpaca data news --symbols AAPL,MSFT --limit 10
  alpaca data news --symbols AAPL --all --max 100`
})

func init() {
	dataCmd.AddCommand(newsCmd)
}
