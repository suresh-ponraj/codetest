package lottery

type Price struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type GameIncrement struct {
	Number4 int `json:"4"`
}

type Set struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type NumberSet struct {
	First int   `json:"first"`
	Last  int   `json:"last"`
	Sets  []Set `json:"sets"`
}

type Combination struct {
	Number10 []int `json:"10"`
}

type GameOffer struct {
	Key             string        `json:"key"`
	Name            string        `json:"name"`
	Description     string        `json:"description"`
	Price           Price         `json:"price"`
	MinGames        int           `json:"min_games"`
	MaxGames        int           `json:"max_games"`
	Multiple        int           `json:"multiple"`
	Ordered         bool          `json:"ordered"`
	GameIncrement   GameIncrement `json:"game_increment"`
	EquivalentGames int           `json:"equivalent_games"`
	NumberSets      []NumberSet   `json:"number_sets"`
	Combinations    Combination   `json:"combinations"`
	DisplayRange    string        `json:"display_range"`
}

type GameType struct {
	Key         string      `json:"key"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	GameOffers  []GameOffer `json:"game_offers"`
}

type PrizePool struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type JackpotImage struct {
	ImageName          string `json:"image_name"`
	ImageURL           string `json:"image_url"`
	SvgURL             string `json:"svg_url"`
	ImageWidth         int    `json:"image_width"`
	ImageHeight        int    `json:"image_height"`
	ContentDescription string `json:"content_description"`
}

type Content struct {
	SalesPitchHeading1    string   `json:"sales_pitch_heading_1"`
	SalesPitchSubHeading1 string   `json:"sales_pitch_sub_heading_1"`
	Paragraph1            string   `json:"paragraph_1"`
	Paragraph2            string   `json:"paragraph_2"`
	Paragraph3            string   `json:"paragraph_3"`
	Image                 string   `json:"image"`
	SalesPitchHeading2    string   `json:"sales_pitch_heading_2"`
	SalesPitchSubHeading2 string   `json:"sales_pitch_sub_heading_2"`
	Features              []string `json:""`
}

type Value struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type Prize struct {
	Type             string   `json:"type"`
	CardTitle        string   `json:"card_title"`
	Name             string   `json:"name"`
	Description      string   `json:"description"`
	Content          Content  `json:"content"`
	Value            Value    `json:"value"`
	ValueIsExact     bool     `json:"value_is_exact"`
	HeroImage        string   `json:"hero_image"`
	CarouselImages   []string `json:"carousel_images"`
	FeatureDrawImage string   `json:"feature_draw_image"`
	EdmImage         string   `json:"edm_image"`
}

type Offer struct {
	Name           string `json:"name"`
	Key            string `json:"key"`
	NumTickets     int    `json:"num_tickets"`
	Price          Price  `json:"price"`
	PricePerTicket Price  `json:"price_per_ticket"`
	Ribbon         string `json:"ribbon"`
	BonusPrize     string `json:"bonus_prize"`
}

type Draws struct {
	Name         string       `json:"name"`
	Date         string       `json:"date"`
	Stop         string       `json:"stop"`
	DrawNo       int          `json:"draw_no"`
	PrizePool    PrizePool    `json:"prize_pool"`
	JackpotImage JackpotImage `json:"jackpot_image"`
}

type Draw struct {
	Description           string  `json:"description"`
	DrawNumber            int     `json:"draw_number"`
	DrawStop              string  `json:"draw_stop"`
	DrawDate              string  `json:"draw_date"`
	Prize                 Prize   `json:"prize"`
	Offers                []Offer `json:"offers"`
	TermsAndConditionsURL string  `json:"terms_and_conditions_url"`
}

type Day struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type Lottery struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Desc         string `json:"desc"`
	Multidraw    bool   `json:"multidraw"`
	Type         string `json:"type"`
	IconURL      string `json:"icon_url"`
	IconWhiteURL string `json:"icon_white_url"`
	PlayURL      string `json:"play_url"`
	LotteryID    int    `json:"lottery_id"`
}

type Result struct {
	Type           string     `json:"type"`
	Key            string     `json:"key"`
	Name           string     `json:"name"`
	Autoplayable   string     `json:"autoplayable"`
	Lottery        Lottery    `json:"lottery"`
	GameType       []GameType `json:"game_types"`
	Draws          []Draws    `json:"draws"`
	Draw           Draw       `json:"draw"`
	Days           []Day      `json:"days"`
	Addons         []string   `json:"addons"`
	QuickpickSizes []int      `json:"quickpick_sizes"`
}

type Ticket struct {
	Results  []Result `json:"result"`
	Messages []string `json:"messages"`
}
