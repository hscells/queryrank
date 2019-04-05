package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ielab/searchrefiner"
	"net/http"
)

type QueryRankPlugin struct {
}

func (q QueryRankPlugin) Serve(s searchrefiner.Server, c *gin.Context) {
	page := c.Query("page")
	switch page {
	case "rank":
		c.Render(http.StatusOK, searchrefiner.RenderPlugin(searchrefiner.TemplatePlugin("plugin/queryrank/rank.html"), queries[0]))
		return
	default:
		// Respond to a regular request.
		c.Render(http.StatusOK, searchrefiner.RenderPlugin(searchrefiner.TemplatePlugin("plugin/queryrank/index.html"), nil))
		return
	}
}

func (q QueryRankPlugin) PermissionType() searchrefiner.PluginPermission {
	return searchrefiner.PluginPublic
}

func (q QueryRankPlugin) Details() searchrefiner.PluginDetails {
	return searchrefiner.PluginDetails{
		Title:       "Query Rank User Study",
		Description: "Rank queries in a user study.",
		Author:      "Harry Scells",
		Version:     "2.Apr.2019",
		ProjectURL:  "github.com/hscells/iqueryrank",
	}
}

var (
	Queryrank = QueryRankPlugin{}
	// Original
	// breadth_evaluation_diversifiedprecision_F0_5Measure
	// positive_precision
	// positive_recall
	queries = [][]string{
		{ // 14
			`(Bronchiectasis[Mesh Terms] OR bronchiect*[All Fields] OR bronchoect*[All Fields] OR kartagener*[All Fields] OR (ciliary[All Fields] AND dyskinesia[All Fields]) OR (bronchial*[All Fields] AND dilat*[All Fields]))`,
			`(bronchiect*[All Fields] OR bronchoect*[All Fields] OR kartagener*[All Fields] OR (bronchiectasis[Mesh Terms:noexp] AND Bronchiectasis[Mesh Terms:noexp]) OR (ciliary[All Fields] AND dyskinesia[All Fields]) OR (bronchial*[All Fields] AND dilat*[All Fields]))`,
			`(Bronchiectasis[Mesh Terms:noexp] OR bronchiect*[All Fields] OR bronchoect*[All Fields] OR kartagener*[All Fields] OR (bronchial*[All Fields] AND dilat*[All Fields]) OR (ciliary[All Fields] AND dyskinesia[All Fields]))`,
			`(bronchiect*[All Fields] OR bronchoect*[All Fields] OR kartagener*[All Fields] OR Bronchiectasis[Mesh Terms:noexp] OR (ciliary[All Fields] AND (((dyskinetic movements[All Fields] OR dyskinesia[All Fields]) OR (dyskinesia[All Fields] OR dyskinetic movements[All Fields])))) OR (bronchial*[All Fields] AND dilat*[All Fields]))`,
		},
		{ // 19
			`(bronchospas*[All Fields] OR bronchoconstrict*[All Fields] OR Asthma[Mesh Terms] OR Respiratory Sounds[Mesh Terms:noexp] OR wheez*[All Fields] OR Bronchial Spasm[Mesh Terms:noexp] OR Bronchoconstriction[Mesh Terms] OR asthma*[All Fields] OR Bronchial Hyperreactivity[Mesh Terms:noexp] OR Respiratory Hypersensitivity[Mesh Terms:noexp] OR (bronch*[All Fields] AND spasm*[All Fields]) OR (antiasthma*[All Fields] OR anti-asthma*[All Fields]) OR (bronch*[All Fields] AND constrict*[All Fields]) OR ((bronchial*[All Fields] OR respiratory[All Fields] OR airway*[All Fields] OR lung*[All Fields]) AND (hypersensitiv*[All Fields] OR hyperreactiv*[All Fields] OR allerg*[All Fields] OR insufficiency[All Fields])) OR ((dust[All Fields] OR mite*[All Fields]) AND (allerg*[All Fields] OR hypersensitiv*[All Fields])))`,
			`(Bronchial Hyperreactivity[Mesh Terms:noexp] OR asthma*[All Fields] OR Diagnostic Techniques, Respiratory System[Mesh Terms] OR Asthma[Mesh Terms:noexp] OR Respiratory Hypersensitivity[Mesh Terms:noexp] OR bronchospas*[All Fields] OR bronchoconstrict*[All Fields] OR Bronchoconstriction[Mesh Terms:noexp] OR wheez*[All Fields] OR Bronchial Spasm[Mesh Terms:noexp] OR ((dust[All Fields] OR mite*[All Fields]) AND (allerg*[All Fields] OR hypersensitiv*[All Fields])) OR (bronch*[All Fields] AND constrict*[All Fields]) OR ((bronchial*[All Fields] OR respiratory[All Fields] OR airway*[All Fields] OR lung*[All Fields]) AND (hypersensitiv*[All Fields] OR hyperreactiv*[All Fields] OR allerg*[All Fields] OR insufficiency[All Fields])) OR (bronch*[All Fields] AND spasm*[All Fields]) OR (antiasthma*[All Fields] OR anti-asthma*[All Fields]))`,
			`(asthma*[All Fields] OR Bronchoconstriction[Mesh Terms:noexp] OR Bronchial Hyperreactivity[Mesh Terms:noexp] OR Asthma[Mesh Terms:noexp] OR Respiratory Hypersensitivity[Mesh Terms:noexp] OR Bronchial Spasm[Mesh Terms:noexp] OR Respiratory Sounds[Mesh Terms:noexp] OR wheez*[All Fields] OR bronchoconstrict*[All Fields] OR bronchospas*[All Fields] OR (bronch*[All Fields] AND constrict*[All Fields]) OR (bronch*[All Fields] AND spasm*[All Fields]) OR ((bronchial*[All Fields] OR respiratory[All Fields] OR airway*[All Fields] OR lung*[All Fields]) AND (hypersensitiv*[All Fields] OR hyperreactiv*[All Fields] OR allerg*[All Fields] OR insufficiency[All Fields])) OR ((dust[All Fields] OR mite*[All Fields]) AND (allerg*[All Fields] OR hypersensitiv*[All Fields])) OR (antiasthma*[All Fields] OR anti-asthma*[All Fields]))`,
			`(wheez*[All Fields] OR Bronchial Spasm[Mesh Terms:noexp] OR Bronchoconstriction[Mesh Terms:noexp] OR Bronchial Hyperreactivity[Mesh Terms:noexp] OR bronchoconstrict*[All Fields] OR Respiratory Hypersensitivity[Mesh Terms:noexp] OR Asthma[Mesh Terms:noexp] OR bronchospas*[All Fields] OR asthma*[All Fields] OR Respiratory Sounds[Mesh Terms:noexp] OR ((dust[All Fields] OR mite*[All Fields]) AND (allerg*[All Fields] OR hypersensitiv*[All Fields])) OR (antiasthma*[All Fields] OR anti-asthma*[All Fields]) OR (bronch*[All Fields] AND spasm*[All Fields]) OR ((bronchial*[All Fields] OR airway*[All Fields] OR lung*[All Fields] OR (((((respiratory[All Fields])))))) AND (hypersensitiv*[All Fields] OR hyperreactiv*[All Fields] OR allerg*[All Fields] OR insufficiency[All Fields])) OR (bronch*[All Fields] AND constrict*[All Fields]))`,
		},
		{ // 28
			`(((randomized controlled trial[Publication Type] OR placebo[Title/Abstract] OR dt[MeSH Subheading:noexp] OR randomly[Title/Abstract] OR trial[Title/Abstract] OR groups[Title/Abstract] OR (randomized[Title/Abstract] OR randomised[Title/Abstract])) NOT (animals[Mesh Terms] AND humans[Mesh Terms])) AND ((trabeculectom*[Title/Abstract] OR sclerostom*[Title/Abstract] OR glaucoma[Mesh Terms] OR intraocular pressure[Mesh Terms] OR filtering surgery[Mesh Terms] OR (glaucoma*[Title/Abstract] AND (surg*[Title/Abstract] OR filter*[Title/Abstract] OR filtrat*[Title/Abstract])) OR (surg*[Title/Abstract] AND intra*[Title/Abstract])) AND (angiogenesis inducing agents[Mesh Terms:noexp] OR endothelial growth factors[Mesh Terms:noexp] OR vascular endothelial growth factors[Mesh Terms] OR angiogenesis inhibitors[Mesh Terms] OR (macugen*[Title/Abstract] OR pegaptanib*[Title/Abstract] OR lucentis*[Title/Abstract] OR rhufab*[Title/Abstract] OR ranibizumab*[Title/Abstract] OR bevacizumab*[Title/Abstract] OR avastin[Title/Abstract] OR aflibercept*[Title/Abstract]) OR (anti[Title/Abstract] AND VEGF*[Title/Abstract]) OR (endothelial[Title/Abstract] AND growth[Title/Abstract] AND factor*[Title/Abstract]))))`,
			`(((trial[Title/Abstract] OR groups[Title/Abstract] OR randomized controlled trial[Publication Type] OR placebo[Title/Abstract] OR dt[MeSH Subheading:noexp] OR randomly[Title/Abstract] OR (randomized[Title/Abstract] OR randomised[Title/Abstract])) NOT (animals[Mesh Terms:noexp] AND humans[Mesh Terms:noexp])) AND ((glaucoma[Mesh Terms:noexp] OR intraocular pressure[Mesh Terms:noexp] OR filtering surgery[Mesh Terms:noexp] OR trabeculectom*[Title/Abstract] OR sclerostom*[Title/Abstract] OR (glaucoma*[Title/Abstract] AND (surg*[Title/Abstract] OR filter*[Title/Abstract] OR filtrat*[Title/Abstract])) OR (surg*[Title/Abstract] AND intra*[Title/Abstract])) AND (angiogenesis inhibitors[Mesh Terms:noexp] OR endothelial growth factors[Mesh Terms:noexp] OR vascular endothelial growth factors[Mesh Terms:noexp] OR (anti[Title/Abstract] AND VEGF*[Title/Abstract]) OR (endothelial[Title/Abstract] AND growth[Title/Abstract] AND factor*[Title/Abstract]) OR (macugen*[Title/Abstract] OR pegaptanib*[Title/Abstract] OR lucentis*[Title/Abstract] OR rhufab*[Title/Abstract] OR ranibizumab*[Title/Abstract] OR bevacizumab*[Title/Abstract] OR avastin[Title/Abstract] OR aflibercept*[Title/Abstract]))))`,
			`(((placebo[Title/Abstract] OR dt[MeSH Subheading:noexp] OR randomly[Title/Abstract] OR trial[Title/Abstract] OR groups[Title/Abstract] OR randomized controlled trial[Publication Type] OR (randomized[Title/Abstract] OR randomised[Title/Abstract])) NOT (animals[Mesh Terms:noexp] AND humans[Mesh Terms:noexp])) AND ((trabeculectom*[Title/Abstract] OR sclerostom*[Title/Abstract] OR glaucoma[Mesh Terms:noexp] OR intraocular pressure[Mesh Terms:noexp] OR filtering surgery[Mesh Terms:noexp] OR (glaucoma*[Title/Abstract] AND (surg*[Title/Abstract] OR filter*[Title/Abstract] OR filtrat*[Title/Abstract])) OR (surg*[Title/Abstract] AND intra*[Title/Abstract])) AND (angiogenesis inhibitors[Mesh Terms:noexp] OR angiogenesis inducing agents[Mesh Terms:noexp] OR endothelial growth factors[Mesh Terms:noexp] OR vascular endothelial growth factors[Mesh Terms:noexp] OR (macugen*[Title/Abstract] OR pegaptanib*[Title/Abstract] OR lucentis*[Title/Abstract] OR rhufab*[Title/Abstract] OR ranibizumab*[Title/Abstract] OR bevacizumab*[Title/Abstract] OR avastin[Title/Abstract] OR aflibercept*[Title/Abstract]) OR (anti[Title/Abstract] AND VEGF*[Title/Abstract]) OR (endothelial[Title/Abstract] AND growth[Title/Abstract] AND factor*[Title/Abstract]))))`,
			`(((placebo[Title/Abstract] OR dt[MeSH Subheading:noexp] OR randomly[Title/Abstract] OR trial[Title/Abstract] OR groups[Text Word] OR randomized controlled trial[Publication Type] OR (randomized[Title/Abstract] OR randomised[Title/Abstract])) NOT (animals[Mesh Terms:noexp] AND Catarrhini[Mesh Terms:noexp])) AND ((glaucoma[Mesh Terms:noexp] OR intraocular pressure[Mesh Terms:noexp] OR filtering surgery[Mesh Terms:noexp] OR trabeculectom*[Title/Abstract] OR sclerostom*[Title/Abstract] OR (surg*[Title/Abstract] AND intra*[Title/Abstract]) OR (glaucoma*[Title/Abstract] AND (surg*[Title/Abstract] OR filter*[Title/Abstract] OR filtrat*[Title/Abstract]))) AND (Growth Inhibitors[Mesh Terms:noexp] OR angiogenesis inducing agents[Mesh Terms:noexp] OR endothelial growth factors[Mesh Terms:noexp] OR vascular endothelial growth factors[Mesh Terms:noexp] OR (endothelial[Title/Abstract] AND growth[Title/Abstract] AND factor*[Title/Abstract]) OR (macugen*[Title/Abstract] OR pegaptanib*[Title/Abstract] OR lucentis*[Title/Abstract] OR rhufab*[Title/Abstract] OR ranibizumab*[Title/Abstract] OR bevacizumab*[Title/Abstract] OR avastin[Title/Abstract] OR aflibercept*[Title/Abstract]) OR (anti[Title/Abstract] AND VEGF*[Title/Abstract]))))`,
		},
		{ // 39
			`((position*[Title/Abstract] OR half-sitting[Title/Abstract] OR supine*[Title/Abstract] OR Posture[Mesh Terms] OR posture*[Title/Abstract] OR Patient Positioning[Mesh Terms:noexp] OR Supine Position[Mesh Terms:noexp] OR (semi-recumbent*[Title/Abstract] OR semi-recumbent*[Title/Abstract]) OR (semisupin*[Title/Abstract] OR semi-supin*[Title/Abstract]) OR (semi-reclin*[Title/Abstract] OR semireclin*[Title/Abstract]) OR ((head*[Title/Abstract] OR bed[Title/Abstract] OR backrest[Title/Abstract]) AND (elevat*[Title/Abstract] OR rais*[Title/Abstract] OR inclin*[Title/Abstract] OR angle[Title/Abstract]))) AND (Pneumonia, Ventilator-Associated[Mesh Terms:noexp] OR vap[Title/Abstract] OR ((Pneumonia[Mesh Terms] OR pneumon*[Title/Abstract]) AND (Respiration, Artificial[Mesh Terms] OR Ventilators, Mechanical[Mesh Terms] OR (ventilat*[Title/Abstract] OR respirat*[Title/Abstract])))))`,
			`((Pneumonia, Ventilator-Associated[Mesh Terms:noexp] OR vap[Title/Abstract] OR ((Pneumonia[Mesh Terms:noexp] OR pneumon*[Title/Abstract]) AND (Resuscitation[Mesh Terms:noexp] OR Ventilators, Mechanical[Mesh Terms] OR (ventilat*[Title/Abstract] OR respirat*[Title/Abstract])))) AND (half-sitting[Title/Abstract] OR Posture[Mesh Terms] OR posture*[Title/Abstract] OR Patient Positioning[Mesh Terms:noexp] OR position*[Title/Abstract] OR Supine Position[Mesh Terms:noexp] OR supine*[Title/Abstract] OR ((head*[Title/Abstract] OR bed[Title/Abstract] OR backrest[Title/Abstract]) AND (elevat*[Title/Abstract] OR rais*[Title/Abstract] OR inclin*[Title/Abstract] OR angle[Title/Abstract])) OR (semi-reclin*[Title/Abstract] OR semireclin*[Title/Abstract]) OR (semisupin*[Title/Abstract] OR semi-supin*[Title/Abstract]) OR (semi-recumbent*[Title/Abstract] OR semi-recumbent*[Title/Abstract])))`,
			`((Pneumonia, Ventilator-Associated[Mesh Terms:noexp] OR vap[Title/Abstract] OR ((Pneumonia[Mesh Terms] OR pneumon*[Title/Abstract]) AND (Respiration, Artificial[Mesh Terms] OR Ventilators, Mechanical[Mesh Terms] OR (ventilat*[Title/Abstract] OR respirat*[Title/Abstract])))) AND (Patient Positioning[Mesh Terms:noexp] OR position*[Title/Abstract] OR Supine Position[Mesh Terms:noexp] OR half-sitting[Title/Abstract] OR Posture[Mesh Terms] OR posture*[Title/Abstract] OR supine*[Title/Abstract] OR (semi-recumbent*[Title/Abstract] OR semi-recumbent*[Title/Abstract]) OR ((head*[Title/Abstract] OR bed[Title/Abstract] OR backrest[Title/Abstract]) AND (elevat*[Title/Abstract] OR rais*[Title/Abstract] OR inclin*[Title/Abstract] OR angle[Title/Abstract])) OR (semi-reclin*[Title/Abstract] OR semireclin*[Title/Abstract]) OR (semisupin*[Title/Abstract] OR semi-supin*[Title/Abstract])))`,
			`((vap[Title/Abstract] AND ((Pneumonia[Mesh Terms] OR pneumon*[Title/Abstract]) AND (Respiration, Artificial[Mesh Terms] OR Ventilators, Mechanical[Mesh Terms] OR (ventilat*[Title/Abstract] OR respirat*[Title/Abstract])))) OR (Patient Positioning[Mesh Terms:noexp] OR position*[Title/Abstract] OR Supine Position[Mesh Terms:noexp] OR posture*[Title/Abstract] OR supine*[Title/Abstract] OR half-sitting[Title/Abstract] OR Posture[Mesh Terms] OR (semi-reclin*[Title/Abstract] OR semireclin*[Title/Abstract]) OR ((head*[Title/Abstract] OR bed[Title/Abstract] OR backrest[Title/Abstract]) AND (elevat*[Title/Abstract] OR rais*[Title/Abstract] OR inclin*[Title/Abstract] OR angle[Title/Abstract])) OR (semi-recumbent*[Title/Abstract] OR semi-recumbent*[Title/Abstract]) OR (semisupin*[Title/Abstract] OR semi-supin*[Title/Abstract])))`,
		},
	}
)
