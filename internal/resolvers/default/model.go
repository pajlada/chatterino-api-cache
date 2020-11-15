package defaultresolver

import (
	"net/http"

	"github.com/Chatterino/api/pkg/resolver"
	"github.com/Chatterino/api/pkg/utils"
	"github.com/PuerkitoBio/goquery"
)

type tooltipData struct {
	URL         string
	Title       string
	Description string
	ImageSrc    string
}

func (d *tooltipData) Truncate() {
	d.Title = utils.TruncateString(d.Title, MaxTitleLength)
	d.Description = utils.TruncateString(d.Description, MaxDescriptionLength)
}

// does this really fit in model?
func (dr *R) defaultTooltipData(doc *goquery.Document, r *http.Request, resp *http.Response) tooltipData {
	data := tooltipMetaFields(dr.baseURL, doc, r, resp, tooltipData{
		URL: resolver.CleanResponse(resp.Request.URL.String()),
	})

	if data.Title == "" {
		data.Title = doc.Find("title").First().Text()
	}

	return data
}
