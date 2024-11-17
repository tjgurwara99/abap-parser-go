package parser_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tjgurwara99/abap-parser-go/parser"
)

func TestReportParsing(t *testing.T) {
	testCases := []struct {
		src     string
		name    string
		program parser.Program
	}{
		{
			src:  `REPORT ZREPORT.`,
			name: "ZREPORT",
			program: parser.Program{
				IntroductoryStmt: &parser.ReportStmt{
					Name: "ZREPORT",
				},
			},
		},
		{
			src:  `REPORT ZREPORT LINE-SIZE 132.`,
			name: "ZREPORT with line size",
			program: parser.Program{
				IntroductoryStmt: &parser.ReportStmt{
					Name: "ZREPORT",
					Additions: &parser.ReportAdditions{
						ListOptions: &parser.ListOptions{
							LineSize: intPtr(132),
						},
					},
				},
			},
		},
		{
			src:  `REPORT ZREPORT LINE-SIZE 132 NO STANDARD PAGE HEADING.`,
			name: "ZREPORT with line size and no standard page heading",
			program: parser.Program{
				IntroductoryStmt: &parser.ReportStmt{
					Name: "ZREPORT",
					Additions: &parser.ReportAdditions{
						ListOptions: &parser.ListOptions{
							LineSize:              intPtr(132),
							NoStandardPageHeading: boolPtr(true),
						},
					},
				},
			},
		},
		{
			src:  `REPORT ZREPORT LINE-SIZE 132 LINE-COUNT 30(123) NO STANDARD PAGE HEADING.`,
			name: "ZREPORT with line size, page lines, footer lines and no standard page heading",
			program: parser.Program{
				IntroductoryStmt: &parser.ReportStmt{
					Name: "ZREPORT",
					Additions: &parser.ReportAdditions{
						ListOptions: &parser.ListOptions{
							LineSize:              intPtr(132),
							NoStandardPageHeading: boolPtr(true),
							PageLines:             intPtr(30),
							FooterLines:           intPtr(123),
						},
					},
				},
			},
		},
		{
			src:  `REPORT ZREPORT LINE-SIZE 132 LINE-COUNT 30 NO STANDARD PAGE HEADING.`,
			name: "ZREPORT with line size, page lines and no standard page heading",
			program: parser.Program{
				IntroductoryStmt: &parser.ReportStmt{
					Name: "ZREPORT",
					Additions: &parser.ReportAdditions{
						ListOptions: &parser.ListOptions{
							LineSize:              intPtr(132),
							NoStandardPageHeading: boolPtr(true),
							PageLines:             intPtr(30),
						},
					},
				},
			},
		},
		{
			src:  `REPORT ZREPORT LINE-SIZE 132 LINE-COUNT 30.`,
			name: "ZREPORT with line size and page lines",
			program: parser.Program{
				IntroductoryStmt: &parser.ReportStmt{
					Name: "ZREPORT",
					Additions: &parser.ReportAdditions{
						ListOptions: &parser.ListOptions{
							LineSize:  intPtr(132),
							PageLines: intPtr(30),
						},
					},
				},
			},
		},
		{
			src:  `REPORT ZREPORT LINE-SIZE 132 MESSAGE-ID z_my_messages NO STANDARD PAGE HEADING.`,
			name: "ZREPORT with line size, message id and no standard page heading",
			program: parser.Program{
				IntroductoryStmt: &parser.ReportStmt{
					Name: "ZREPORT",
					Additions: &parser.ReportAdditions{
						MsgID: strPtr("z_my_messages"),
						ListOptions: &parser.ListOptions{
							LineSize:              intPtr(132),
							NoStandardPageHeading: boolPtr(true),
						},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			parsed, err := parser.Parse("test.abap", []byte(tc.src), parser.Debug(true))
			require.NoError(t, err, "unexpected error: %v", err)
			p, ok := parsed.(*parser.Program)
			require.True(t, ok, "unexpected type: %T", parsed)
			require.Equal(t, tc.program, *p)
		})
	}

}
