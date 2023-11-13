package models

type Vulnerability struct {
	Id            string `json:"id"`
	Title         string `json:"title"`
	PublishedDate string `json:"published_date"`
	Description   string `json:"description"`
	VulnType      string `json:"vuln_type"`
	References    struct {
		Url []string `json:"url"`
		Cve []string `json:"cve"`
	} `json:"references"`
	Cvss struct {
		Score    string `json:"score"`
		Vector   string `json:"vector"`
		Severity string `json:"severity"`
	} `json:"cvss"`
	Verified     bool   `json:"verified"`
	FixedIn      string `json:"fixed_in"`
	IntroducedIn string `json:"introduced_in"`
}

type VulnInfo struct {
	ReleaseDate string          `json:"release_date"`
	Status      string          `json:"status"`
	Vulns       []Vulnerability `json:"vulnerabilities"`
}

type WPScan map[string]VulnInfo

/*
{
    "4.9.4": {
      "release_date": "2018-02-06",
      "changelog_url": "https://codex.wordpress.org/Version_4.9.4",
      "status": "insecure",
      "vulnerabilities": [
        {
          "id": "5e0c1ddd-fdd0-421b-bdbe-3eee6b75c919",
          "old_id": 9021,
          "title": "WordPress <= 4.9.4 - Application Denial of Service (DoS) (unpatched)",
          "created_at": "2018-02-05T16:50:40.000Z",
          "updated_at": "2020-09-22T07:24:12.000Z",
          "published_date": "2018-02-05T00:00:00.000Z",
          "description": "An application Denial of Service (DoS) was found to affect WordPress versions 4.9.4 and below. We are not aware of a patch for this issue.",
          "poc": "string",
          "vuln_type": "DOS",
          "references": {
            "url": "https://baraktawily.blogspot.fr/2018/02/how-to-dos-29-of-world-wide-websites.html",
            "cve": "2018-6389"
          },
          "cvss": {
            "score": "7.5",
            "vector": "CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:N/I:N/A:H",
            "severity": "high"
          },
          "verified": false,
          "fixed_in": "4.9.5",
          "introduced_in": "1.0"
        }
      ]
    }
  }

*/
