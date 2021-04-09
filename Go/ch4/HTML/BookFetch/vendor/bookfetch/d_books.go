// Creating a web server to store http get result in local mongodb instance

package bookfetch

import "html/template"

var APIKey = "AIzaSyDtMtMmrqiOwKIbZvXXLXacC32S98vkQ18"
var URL = "https://www.googleapis.com/books/v1/volumes"
var Collection = "gbooks"

type Books struct {
	Count      int
	TotalItems int `json:totalItems`
	Items      []*Items
}

type Items struct {
	VolumeInfo *VolumeInfo
	AccessInfo *AccessInfo
}

type VolumeInfo struct {
	Title         string     `json:"title,omitempty"`
	Subtitle      string     `json:"subtitle,omitempty"`
	Authors       []string   `json:"authors,omitempty"`
	Publisher     string     `json:"publisher,omitempty"`
	PublishedDate string     `json:"publishedDate,omitempty"`
	Description   string     `json:"description,omitempty"`
	PageCount     int        `json:"pageCount,omitempty"`
	PreviewLink   string     `json:"previewLink,omitempty"`
	Category      []string   `json:categories,omitempty`
	ImageLinks    ImageLinks `json:imagelinks,omitempty`
}

type ImageLinks struct {
	Thumbnail string `json:"thumbnail,omitempty"`
}

type AccessInfo struct {
	WebReader string `json:"webReaderLink,omitempty"`
}

var BookMeta = template.Must(template.New("bookmeta").Parse(`
{.TotalItems}} Entries Found
{{range .Items}}----------------------------------------   
Title: {{.VolumeInfo.Title}} {{.VolumeInfo.Subtitle}}  
Publisher: {{.VolumeInfo.Publisher}}        
Authors: {{.VolumeInfo.Authors}}   
Description: {{.VolumeInfo.Description}}    
PublishedDate: {{.VolumeInfo.PublishedDate}} 
PageCount: {{.VolumeInfo.PageCount}}     
<a my:href="{{.AccessInfo.WebReader}}">Reader</a>   
{{end}}`))

var BookLinks = template.Must(template.New("BookLinks").Parse(`
{{range .Items}}
	Title: {{.VolumeInfo.Title}} {{.VolumeInfo.Subtitle}}
	Publisher: {{.VolumeInfo.Publisher}}
	Authors: {{.VolumeInfo.Authors}}
	Description: {{.VolumeInfo.Description}}    
	PublishedDate: {{.VolumeInfo.PublishedDate}} 
	PageCount: {{.VolumeInfo.PageCount}}
	<a href='{{.VolumeInfo.ImageLinks.Thumbnail}}'/a>
{{end}}
`))

var issueList = template.Must(template.New("issuelist").Parse(`
<h1>{{.TotalItems}} Entries</h1>
<table>
<tr style='text-align: left'>
  <th>List</th>
  <th>Information</th>
</tr>
{{range .Items}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))

// <a href='{{.VolumeInfo.ImageLinks.Thumbnail}}>"Image"</a><br>
// <a href='{{.AccessInfo.WebReader}}>"WebReader"</a><br>
// <a href='{{.VolumeInfo.PreviewLink}}>"Preview"</a><br>
// <b>----------------------------------------</b>
