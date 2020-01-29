package emoji

// Emojis - Map of Emoji Runes as Hex keys to their description
var Emojis = map[string]Emoji{
	"1F004": {
		Key:        "1F004",
		Value:      "🀄",
		Descriptor: "Mahjong Red Dragon",
	},
	"1F0CF": {
		Key:        "1F0CF",
		Value:      "🃏",
		Descriptor: "Joker",
	},
	"1F170-FE0F": {
		Key:        "1F170-FE0F",
		Value:      "🅰️",
		Descriptor: "A Button (Blood Type)",
	},
	"1F171-FE0F": {
		Key:        "1F171-FE0F",
		Value:      "🅱️",
		Descriptor: "B Button (Blood Type)",
	},
	"1F17E-FE0F": {
		Key:        "1F17E-FE0F",
		Value:      "🅾️",
		Descriptor: "O Button (Blood Type)",
	},
	"1F17F-FE0F": {
		Key:        "1F17F-FE0F",
		Value:      "🅿️",
		Descriptor: "P Button",
	},
	"1F18E": {
		Key:        "1F18E",
		Value:      "🆎",
		Descriptor: "AB Button (Blood Type)",
	},
	"1F191": {
		Key:        "1F191",
		Value:      "🆑",
		Descriptor: "CL Button",
	},
	"1F192": {
		Key:        "1F192",
		Value:      "🆒",
		Descriptor: "Cool Button",
	},
	"1F193": {
		Key:        "1F193",
		Value:      "🆓",
		Descriptor: "Free Button",
	},
	"1F194": {
		Key:        "1F194",
		Value:      "🆔",
		Descriptor: "ID Button",
	},
	"1F195": {
		Key:        "1F195",
		Value:      "🆕",
		Descriptor: "New Button",
	},
	"1F196": {
		Key:        "1F196",
		Value:      "🆖",
		Descriptor: "NG Button",
	},
	"1F197": {
		Key:        "1F197",
		Value:      "🆗",
		Descriptor: "OK Button",
	},
	"1F198": {
		Key:        "1F198",
		Value:      "🆘",
		Descriptor: "SOS Button",
	},
	"1F199": {
		Key:        "1F199",
		Value:      "🆙",
		Descriptor: "Up! Button",
	},
	"1F19A": {
		Key:        "1F19A",
		Value:      "🆚",
		Descriptor: "Vs Button",
	},
	"1F1E6-1F1E8": {
		Key:        "1F1E6-1F1E8",
		Value:      "🇦🇨",
		Descriptor: "Flag: Ascension Island",
	},
	"1F1E6-1F1E9": {
		Key:        "1F1E6-1F1E9",
		Value:      "🇦🇩",
		Descriptor: "Flag: Andorra",
	},
	"1F1E6-1F1EA": {
		Key:        "1F1E6-1F1EA",
		Value:      "🇦🇪",
		Descriptor: "Flag: United Arab Emirates",
	},
	"1F1E6-1F1EB": {
		Key:        "1F1E6-1F1EB",
		Value:      "🇦🇫",
		Descriptor: "Flag: Afghanistan",
	},
	"1F1E6-1F1EC": {
		Key:        "1F1E6-1F1EC",
		Value:      "🇦🇬",
		Descriptor: "Flag: Antigua \u0026 Barbuda",
	},
	"1F1E6-1F1EE": {
		Key:        "1F1E6-1F1EE",
		Value:      "🇦🇮",
		Descriptor: "Flag: Anguilla",
	},
	"1F1E6-1F1F1": {
		Key:        "1F1E6-1F1F1",
		Value:      "🇦🇱",
		Descriptor: "Flag: Albania",
	},
	"1F1E6-1F1F2": {
		Key:        "1F1E6-1F1F2",
		Value:      "🇦🇲",
		Descriptor: "Flag: Armenia",
	},
	"1F1E6-1F1F4": {
		Key:        "1F1E6-1F1F4",
		Value:      "🇦🇴",
		Descriptor: "Flag: Angola",
	},
	"1F1E6-1F1F6": {
		Key:        "1F1E6-1F1F6",
		Value:      "🇦🇶",
		Descriptor: "Flag: Antarctica",
	},
	"1F1E6-1F1F7": {
		Key:        "1F1E6-1F1F7",
		Value:      "🇦🇷",
		Descriptor: "Flag: Argentina",
	},
	"1F1E6-1F1F8": {
		Key:        "1F1E6-1F1F8",
		Value:      "🇦🇸",
		Descriptor: "Flag: American Samoa",
	},
	"1F1E6-1F1F9": {
		Key:        "1F1E6-1F1F9",
		Value:      "🇦🇹",
		Descriptor: "Flag: Austria",
	},
	"1F1E6-1F1FA": {
		Key:        "1F1E6-1F1FA",
		Value:      "🇦🇺",
		Descriptor: "Flag: Australia",
	},
	"1F1E6-1F1FC": {
		Key:        "1F1E6-1F1FC",
		Value:      "🇦🇼",
		Descriptor: "Flag: Aruba",
	},
	"1F1E6-1F1FD": {
		Key:        "1F1E6-1F1FD",
		Value:      "🇦🇽",
		Descriptor: "Flag: Åland Islands",
	},
	"1F1E6-1F1FF": {
		Key:        "1F1E6-1F1FF",
		Value:      "🇦🇿",
		Descriptor: "Flag: Azerbaijan",
	},
	"1F1E7-1F1E6": {
		Key:        "1F1E7-1F1E6",
		Value:      "🇧🇦",
		Descriptor: "Flag: Bosnia \u0026 Herzegovina",
	},
	"1F1E7-1F1E7": {
		Key:        "1F1E7-1F1E7",
		Value:      "🇧🇧",
		Descriptor: "Flag: Barbados",
	},
	"1F1E7-1F1E9": {
		Key:        "1F1E7-1F1E9",
		Value:      "🇧🇩",
		Descriptor: "Flag: Bangladesh",
	},
	"1F1E7-1F1EA": {
		Key:        "1F1E7-1F1EA",
		Value:      "🇧🇪",
		Descriptor: "Flag: Belgium",
	},
	"1F1E7-1F1EB": {
		Key:        "1F1E7-1F1EB",
		Value:      "🇧🇫",
		Descriptor: "Flag: Burkina Faso",
	},
	"1F1E7-1F1EC": {
		Key:        "1F1E7-1F1EC",
		Value:      "🇧🇬",
		Descriptor: "Flag: Bulgaria",
	},
	"1F1E7-1F1ED": {
		Key:        "1F1E7-1F1ED",
		Value:      "🇧🇭",
		Descriptor: "Flag: Bahrain",
	},
	"1F1E7-1F1EE": {
		Key:        "1F1E7-1F1EE",
		Value:      "🇧🇮",
		Descriptor: "Flag: Burundi",
	},
	"1F1E7-1F1EF": {
		Key:        "1F1E7-1F1EF",
		Value:      "🇧🇯",
		Descriptor: "Flag: Benin",
	},
	"1F1E7-1F1F1": {
		Key:        "1F1E7-1F1F1",
		Value:      "🇧🇱",
		Descriptor: "Flag: St. Barthélemy",
	},
	"1F1E7-1F1F2": {
		Key:        "1F1E7-1F1F2",
		Value:      "🇧🇲",
		Descriptor: "Flag: Bermuda",
	},
	"1F1E7-1F1F3": {
		Key:        "1F1E7-1F1F3",
		Value:      "🇧🇳",
		Descriptor: "Flag: Brunei",
	},
	"1F1E7-1F1F4": {
		Key:        "1F1E7-1F1F4",
		Value:      "🇧🇴",
		Descriptor: "Flag: Bolivia",
	},
	"1F1E7-1F1F6": {
		Key:        "1F1E7-1F1F6",
		Value:      "🇧🇶",
		Descriptor: "Flag: Caribbean Netherlands",
	},
	"1F1E7-1F1F7": {
		Key:        "1F1E7-1F1F7",
		Value:      "🇧🇷",
		Descriptor: "Flag: Brazil",
	},
	"1F1E7-1F1F8": {
		Key:        "1F1E7-1F1F8",
		Value:      "🇧🇸",
		Descriptor: "Flag: Bahamas",
	},
	"1F1E7-1F1F9": {
		Key:        "1F1E7-1F1F9",
		Value:      "🇧🇹",
		Descriptor: "Flag: Bhutan",
	},
	"1F1E7-1F1FB": {
		Key:        "1F1E7-1F1FB",
		Value:      "🇧🇻",
		Descriptor: "Flag: Bouvet Island",
	},
	"1F1E7-1F1FC": {
		Key:        "1F1E7-1F1FC",
		Value:      "🇧🇼",
		Descriptor: "Flag: Botswana",
	},
	"1F1E7-1F1FE": {
		Key:        "1F1E7-1F1FE",
		Value:      "🇧🇾",
		Descriptor: "Flag: Belarus",
	},
	"1F1E7-1F1FF": {
		Key:        "1F1E7-1F1FF",
		Value:      "🇧🇿",
		Descriptor: "Flag: Belize",
	},
	"1F1E8-1F1E6": {
		Key:        "1F1E8-1F1E6",
		Value:      "🇨🇦",
		Descriptor: "Flag: Canada",
	},
	"1F1E8-1F1E8": {
		Key:        "1F1E8-1F1E8",
		Value:      "🇨🇨",
		Descriptor: "Flag: Cocos (Keeling) Islands",
	},
	"1F1E8-1F1E9": {
		Key:        "1F1E8-1F1E9",
		Value:      "🇨🇩",
		Descriptor: "Flag: Congo - Kinshasa",
	},
	"1F1E8-1F1EB": {
		Key:        "1F1E8-1F1EB",
		Value:      "🇨🇫",
		Descriptor: "Flag: Central African Republic",
	},
	"1F1E8-1F1EC": {
		Key:        "1F1E8-1F1EC",
		Value:      "🇨🇬",
		Descriptor: "Flag: Congo - Brazzaville",
	},
	"1F1E8-1F1ED": {
		Key:        "1F1E8-1F1ED",
		Value:      "🇨🇭",
		Descriptor: "Flag: Switzerland",
	},
	"1F1E8-1F1EE": {
		Key:        "1F1E8-1F1EE",
		Value:      "🇨🇮",
		Descriptor: "Flag: Côte d’Ivoire",
	},
	"1F1E8-1F1F0": {
		Key:        "1F1E8-1F1F0",
		Value:      "🇨🇰",
		Descriptor: "Flag: Cook Islands",
	},
	"1F1E8-1F1F1": {
		Key:        "1F1E8-1F1F1",
		Value:      "🇨🇱",
		Descriptor: "Flag: Chile",
	},
	"1F1E8-1F1F2": {
		Key:        "1F1E8-1F1F2",
		Value:      "🇨🇲",
		Descriptor: "Flag: Cameroon",
	},
	"1F1E8-1F1F3": {
		Key:        "1F1E8-1F1F3",
		Value:      "🇨🇳",
		Descriptor: "Flag: China",
	},
	"1F1E8-1F1F4": {
		Key:        "1F1E8-1F1F4",
		Value:      "🇨🇴",
		Descriptor: "Flag: Colombia",
	},
	"1F1E8-1F1F5": {
		Key:        "1F1E8-1F1F5",
		Value:      "🇨🇵",
		Descriptor: "Flag: Clipperton Island",
	},
	"1F1E8-1F1F7": {
		Key:        "1F1E8-1F1F7",
		Value:      "🇨🇷",
		Descriptor: "Flag: Costa Rica",
	},
	"1F1E8-1F1FA": {
		Key:        "1F1E8-1F1FA",
		Value:      "🇨🇺",
		Descriptor: "Flag: Cuba",
	},
	"1F1E8-1F1FB": {
		Key:        "1F1E8-1F1FB",
		Value:      "🇨🇻",
		Descriptor: "Flag: Cape Verde",
	},
	"1F1E8-1F1FC": {
		Key:        "1F1E8-1F1FC",
		Value:      "🇨🇼",
		Descriptor: "Flag: Curaçao",
	},
	"1F1E8-1F1FD": {
		Key:        "1F1E8-1F1FD",
		Value:      "🇨🇽",
		Descriptor: "Flag: Christmas Island",
	},
	"1F1E8-1F1FE": {
		Key:        "1F1E8-1F1FE",
		Value:      "🇨🇾",
		Descriptor: "Flag: Cyprus",
	},
	"1F1E8-1F1FF": {
		Key:        "1F1E8-1F1FF",
		Value:      "🇨🇿",
		Descriptor: "Flag: Czechia",
	},
	"1F1E9-1F1EA": {
		Key:        "1F1E9-1F1EA",
		Value:      "🇩🇪",
		Descriptor: "Flag: Germany",
	},
	"1F1E9-1F1EC": {
		Key:        "1F1E9-1F1EC",
		Value:      "🇩🇬",
		Descriptor: "Flag: Diego Garcia",
	},
	"1F1E9-1F1EF": {
		Key:        "1F1E9-1F1EF",
		Value:      "🇩🇯",
		Descriptor: "Flag: Djibouti",
	},
	"1F1E9-1F1F0": {
		Key:        "1F1E9-1F1F0",
		Value:      "🇩🇰",
		Descriptor: "Flag: Denmark",
	},
	"1F1E9-1F1F2": {
		Key:        "1F1E9-1F1F2",
		Value:      "🇩🇲",
		Descriptor: "Flag: Dominica",
	},
	"1F1E9-1F1F4": {
		Key:        "1F1E9-1F1F4",
		Value:      "🇩🇴",
		Descriptor: "Flag: Dominican Republic",
	},
	"1F1E9-1F1FF": {
		Key:        "1F1E9-1F1FF",
		Value:      "🇩🇿",
		Descriptor: "Flag: Algeria",
	},
	"1F1EA-1F1E6": {
		Key:        "1F1EA-1F1E6",
		Value:      "🇪🇦",
		Descriptor: "Flag: Ceuta \u0026 Melilla",
	},
	"1F1EA-1F1E8": {
		Key:        "1F1EA-1F1E8",
		Value:      "🇪🇨",
		Descriptor: "Flag: Ecuador",
	},
	"1F1EA-1F1EA": {
		Key:        "1F1EA-1F1EA",
		Value:      "🇪🇪",
		Descriptor: "Flag: Estonia",
	},
	"1F1EA-1F1EC": {
		Key:        "1F1EA-1F1EC",
		Value:      "🇪🇬",
		Descriptor: "Flag: Egypt",
	},
	"1F1EA-1F1ED": {
		Key:        "1F1EA-1F1ED",
		Value:      "🇪🇭",
		Descriptor: "Flag: Western Sahara",
	},
	"1F1EA-1F1F7": {
		Key:        "1F1EA-1F1F7",
		Value:      "🇪🇷",
		Descriptor: "Flag: Eritrea",
	},
	"1F1EA-1F1F8": {
		Key:        "1F1EA-1F1F8",
		Value:      "🇪🇸",
		Descriptor: "Flag: Spain",
	},
	"1F1EA-1F1F9": {
		Key:        "1F1EA-1F1F9",
		Value:      "🇪🇹",
		Descriptor: "Flag: Ethiopia",
	},
	"1F1EA-1F1FA": {
		Key:        "1F1EA-1F1FA",
		Value:      "🇪🇺",
		Descriptor: "Flag: European Union",
	},
	"1F1EB-1F1EE": {
		Key:        "1F1EB-1F1EE",
		Value:      "🇫🇮",
		Descriptor: "Flag: Finland",
	},
	"1F1EB-1F1EF": {
		Key:        "1F1EB-1F1EF",
		Value:      "🇫🇯",
		Descriptor: "Flag: Fiji",
	},
	"1F1EB-1F1F0": {
		Key:        "1F1EB-1F1F0",
		Value:      "🇫🇰",
		Descriptor: "Flag: Falkland Islands",
	},
	"1F1EB-1F1F2": {
		Key:        "1F1EB-1F1F2",
		Value:      "🇫🇲",
		Descriptor: "Flag: Micronesia",
	},
	"1F1EB-1F1F4": {
		Key:        "1F1EB-1F1F4",
		Value:      "🇫🇴",
		Descriptor: "Flag: Faroe Islands",
	},
	"1F1EB-1F1F7": {
		Key:        "1F1EB-1F1F7",
		Value:      "🇫🇷",
		Descriptor: "Flag: France",
	},
	"1F1EC-1F1E6": {
		Key:        "1F1EC-1F1E6",
		Value:      "🇬🇦",
		Descriptor: "Flag: Gabon",
	},
	"1F1EC-1F1E7": {
		Key:        "1F1EC-1F1E7",
		Value:      "🇬🇧",
		Descriptor: "Flag: United Kingdom",
	},
	"1F1EC-1F1E9": {
		Key:        "1F1EC-1F1E9",
		Value:      "🇬🇩",
		Descriptor: "Flag: Grenada",
	},
	"1F1EC-1F1EA": {
		Key:        "1F1EC-1F1EA",
		Value:      "🇬🇪",
		Descriptor: "Flag: Georgia",
	},
	"1F1EC-1F1EB": {
		Key:        "1F1EC-1F1EB",
		Value:      "🇬🇫",
		Descriptor: "Flag: French Guiana",
	},
	"1F1EC-1F1EC": {
		Key:        "1F1EC-1F1EC",
		Value:      "🇬🇬",
		Descriptor: "Flag: Guernsey",
	},
	"1F1EC-1F1ED": {
		Key:        "1F1EC-1F1ED",
		Value:      "🇬🇭",
		Descriptor: "Flag: Ghana",
	},
	"1F1EC-1F1EE": {
		Key:        "1F1EC-1F1EE",
		Value:      "🇬🇮",
		Descriptor: "Flag: Gibraltar",
	},
	"1F1EC-1F1F1": {
		Key:        "1F1EC-1F1F1",
		Value:      "🇬🇱",
		Descriptor: "Flag: Greenland",
	},
	"1F1EC-1F1F2": {
		Key:        "1F1EC-1F1F2",
		Value:      "🇬🇲",
		Descriptor: "Flag: Gambia",
	},
	"1F1EC-1F1F3": {
		Key:        "1F1EC-1F1F3",
		Value:      "🇬🇳",
		Descriptor: "Flag: Guinea",
	},
	"1F1EC-1F1F5": {
		Key:        "1F1EC-1F1F5",
		Value:      "🇬🇵",
		Descriptor: "Flag: Guadeloupe",
	},
	"1F1EC-1F1F6": {
		Key:        "1F1EC-1F1F6",
		Value:      "🇬🇶",
		Descriptor: "Flag: Equatorial Guinea",
	},
	"1F1EC-1F1F7": {
		Key:        "1F1EC-1F1F7",
		Value:      "🇬🇷",
		Descriptor: "Flag: Greece",
	},
	"1F1EC-1F1F8": {
		Key:        "1F1EC-1F1F8",
		Value:      "🇬🇸",
		Descriptor: "Flag: South Georgia \u0026 South Sandwich Islands",
	},
	"1F1EC-1F1F9": {
		Key:        "1F1EC-1F1F9",
		Value:      "🇬🇹",
		Descriptor: "Flag: Guatemala",
	},
	"1F1EC-1F1FA": {
		Key:        "1F1EC-1F1FA",
		Value:      "🇬🇺",
		Descriptor: "Flag: Guam",
	},
	"1F1EC-1F1FC": {
		Key:        "1F1EC-1F1FC",
		Value:      "🇬🇼",
		Descriptor: "Flag: Guinea-Bissau",
	},
	"1F1EC-1F1FE": {
		Key:        "1F1EC-1F1FE",
		Value:      "🇬🇾",
		Descriptor: "Flag: Guyana",
	},
	"1F1ED-1F1F0": {
		Key:        "1F1ED-1F1F0",
		Value:      "🇭🇰",
		Descriptor: "Flag: Hong Kong SAR China",
	},
	"1F1ED-1F1F2": {
		Key:        "1F1ED-1F1F2",
		Value:      "🇭🇲",
		Descriptor: "Flag: Heard \u0026 McDonald Islands",
	},
	"1F1ED-1F1F3": {
		Key:        "1F1ED-1F1F3",
		Value:      "🇭🇳",
		Descriptor: "Flag: Honduras",
	},
	"1F1ED-1F1F7": {
		Key:        "1F1ED-1F1F7",
		Value:      "🇭🇷",
		Descriptor: "Flag: Croatia",
	},
	"1F1ED-1F1F9": {
		Key:        "1F1ED-1F1F9",
		Value:      "🇭🇹",
		Descriptor: "Flag: Haiti",
	},
	"1F1ED-1F1FA": {
		Key:        "1F1ED-1F1FA",
		Value:      "🇭🇺",
		Descriptor: "Flag: Hungary",
	},
	"1F1EE-1F1E8": {
		Key:        "1F1EE-1F1E8",
		Value:      "🇮🇨",
		Descriptor: "Flag: Canary Islands",
	},
	"1F1EE-1F1E9": {
		Key:        "1F1EE-1F1E9",
		Value:      "🇮🇩",
		Descriptor: "Flag: Indonesia",
	},
	"1F1EE-1F1EA": {
		Key:        "1F1EE-1F1EA",
		Value:      "🇮🇪",
		Descriptor: "Flag: Ireland",
	},
	"1F1EE-1F1F1": {
		Key:        "1F1EE-1F1F1",
		Value:      "🇮🇱",
		Descriptor: "Flag: Israel",
	},
	"1F1EE-1F1F2": {
		Key:        "1F1EE-1F1F2",
		Value:      "🇮🇲",
		Descriptor: "Flag: Isle of Man",
	},
	"1F1EE-1F1F3": {
		Key:        "1F1EE-1F1F3",
		Value:      "🇮🇳",
		Descriptor: "Flag: India",
	},
	"1F1EE-1F1F4": {
		Key:        "1F1EE-1F1F4",
		Value:      "🇮🇴",
		Descriptor: "Flag: British Indian Ocean Territory",
	},
	"1F1EE-1F1F6": {
		Key:        "1F1EE-1F1F6",
		Value:      "🇮🇶",
		Descriptor: "Flag: Iraq",
	},
	"1F1EE-1F1F7": {
		Key:        "1F1EE-1F1F7",
		Value:      "🇮🇷",
		Descriptor: "Flag: Iran",
	},
	"1F1EE-1F1F8": {
		Key:        "1F1EE-1F1F8",
		Value:      "🇮🇸",
		Descriptor: "Flag: Iceland",
	},
	"1F1EE-1F1F9": {
		Key:        "1F1EE-1F1F9",
		Value:      "🇮🇹",
		Descriptor: "Flag: Italy",
	},
	"1F1EF-1F1EA": {
		Key:        "1F1EF-1F1EA",
		Value:      "🇯🇪",
		Descriptor: "Flag: Jersey",
	},
	"1F1EF-1F1F2": {
		Key:        "1F1EF-1F1F2",
		Value:      "🇯🇲",
		Descriptor: "Flag: Jamaica",
	},
	"1F1EF-1F1F4": {
		Key:        "1F1EF-1F1F4",
		Value:      "🇯🇴",
		Descriptor: "Flag: Jordan",
	},
	"1F1EF-1F1F5": {
		Key:        "1F1EF-1F1F5",
		Value:      "🇯🇵",
		Descriptor: "Flag: Japan",
	},
	"1F1F0-1F1EA": {
		Key:        "1F1F0-1F1EA",
		Value:      "🇰🇪",
		Descriptor: "Flag: Kenya",
	},
	"1F1F0-1F1EC": {
		Key:        "1F1F0-1F1EC",
		Value:      "🇰🇬",
		Descriptor: "Flag: Kyrgyzstan",
	},
	"1F1F0-1F1ED": {
		Key:        "1F1F0-1F1ED",
		Value:      "🇰🇭",
		Descriptor: "Flag: Cambodia",
	},
	"1F1F0-1F1EE": {
		Key:        "1F1F0-1F1EE",
		Value:      "🇰🇮",
		Descriptor: "Flag: Kiribati",
	},
	"1F1F0-1F1F2": {
		Key:        "1F1F0-1F1F2",
		Value:      "🇰🇲",
		Descriptor: "Flag: Comoros",
	},
	"1F1F0-1F1F3": {
		Key:        "1F1F0-1F1F3",
		Value:      "🇰🇳",
		Descriptor: "Flag: St. Kitts \u0026 Nevis",
	},
	"1F1F0-1F1F5": {
		Key:        "1F1F0-1F1F5",
		Value:      "🇰🇵",
		Descriptor: "Flag: North Korea",
	},
	"1F1F0-1F1F7": {
		Key:        "1F1F0-1F1F7",
		Value:      "🇰🇷",
		Descriptor: "Flag: South Korea",
	},
	"1F1F0-1F1FC": {
		Key:        "1F1F0-1F1FC",
		Value:      "🇰🇼",
		Descriptor: "Flag: Kuwait",
	},
	"1F1F0-1F1FE": {
		Key:        "1F1F0-1F1FE",
		Value:      "🇰🇾",
		Descriptor: "Flag: Cayman Islands",
	},
	"1F1F0-1F1FF": {
		Key:        "1F1F0-1F1FF",
		Value:      "🇰🇿",
		Descriptor: "Flag: Kazakhstan",
	},
	"1F1F1-1F1E6": {
		Key:        "1F1F1-1F1E6",
		Value:      "🇱🇦",
		Descriptor: "Flag: Laos",
	},
	"1F1F1-1F1E7": {
		Key:        "1F1F1-1F1E7",
		Value:      "🇱🇧",
		Descriptor: "Flag: Lebanon",
	},
	"1F1F1-1F1E8": {
		Key:        "1F1F1-1F1E8",
		Value:      "🇱🇨",
		Descriptor: "Flag: St. Lucia",
	},
	"1F1F1-1F1EE": {
		Key:        "1F1F1-1F1EE",
		Value:      "🇱🇮",
		Descriptor: "Flag: Liechtenstein",
	},
	"1F1F1-1F1F0": {
		Key:        "1F1F1-1F1F0",
		Value:      "🇱🇰",
		Descriptor: "Flag: Sri Lanka",
	},
	"1F1F1-1F1F7": {
		Key:        "1F1F1-1F1F7",
		Value:      "🇱🇷",
		Descriptor: "Flag: Liberia",
	},
	"1F1F1-1F1F8": {
		Key:        "1F1F1-1F1F8",
		Value:      "🇱🇸",
		Descriptor: "Flag: Lesotho",
	},
	"1F1F1-1F1F9": {
		Key:        "1F1F1-1F1F9",
		Value:      "🇱🇹",
		Descriptor: "Flag: Lithuania",
	},
	"1F1F1-1F1FA": {
		Key:        "1F1F1-1F1FA",
		Value:      "🇱🇺",
		Descriptor: "Flag: Luxembourg",
	},
	"1F1F1-1F1FB": {
		Key:        "1F1F1-1F1FB",
		Value:      "🇱🇻",
		Descriptor: "Flag: Latvia",
	},
	"1F1F1-1F1FE": {
		Key:        "1F1F1-1F1FE",
		Value:      "🇱🇾",
		Descriptor: "Flag: Libya",
	},
	"1F1F2-1F1E6": {
		Key:        "1F1F2-1F1E6",
		Value:      "🇲🇦",
		Descriptor: "Flag: Morocco",
	},
	"1F1F2-1F1E8": {
		Key:        "1F1F2-1F1E8",
		Value:      "🇲🇨",
		Descriptor: "Flag: Monaco",
	},
	"1F1F2-1F1E9": {
		Key:        "1F1F2-1F1E9",
		Value:      "🇲🇩",
		Descriptor: "Flag: Moldova",
	},
	"1F1F2-1F1EA": {
		Key:        "1F1F2-1F1EA",
		Value:      "🇲🇪",
		Descriptor: "Flag: Montenegro",
	},
	"1F1F2-1F1EB": {
		Key:        "1F1F2-1F1EB",
		Value:      "🇲🇫",
		Descriptor: "Flag: St. Martin",
	},
	"1F1F2-1F1EC": {
		Key:        "1F1F2-1F1EC",
		Value:      "🇲🇬",
		Descriptor: "Flag: Madagascar",
	},
	"1F1F2-1F1ED": {
		Key:        "1F1F2-1F1ED",
		Value:      "🇲🇭",
		Descriptor: "Flag: Marshall Islands",
	},
	"1F1F2-1F1F0": {
		Key:        "1F1F2-1F1F0",
		Value:      "🇲🇰",
		Descriptor: "Flag: North Macedonia",
	},
	"1F1F2-1F1F1": {
		Key:        "1F1F2-1F1F1",
		Value:      "🇲🇱",
		Descriptor: "Flag: Mali",
	},
	"1F1F2-1F1F2": {
		Key:        "1F1F2-1F1F2",
		Value:      "🇲🇲",
		Descriptor: "Flag: Myanmar (Burma)",
	},
	"1F1F2-1F1F3": {
		Key:        "1F1F2-1F1F3",
		Value:      "🇲🇳",
		Descriptor: "Flag: Mongolia",
	},
	"1F1F2-1F1F4": {
		Key:        "1F1F2-1F1F4",
		Value:      "🇲🇴",
		Descriptor: "Flag: Macao Sar China",
	},
	"1F1F2-1F1F5": {
		Key:        "1F1F2-1F1F5",
		Value:      "🇲🇵",
		Descriptor: "Flag: Northern Mariana Islands",
	},
	"1F1F2-1F1F6": {
		Key:        "1F1F2-1F1F6",
		Value:      "🇲🇶",
		Descriptor: "Flag: Martinique",
	},
	"1F1F2-1F1F7": {
		Key:        "1F1F2-1F1F7",
		Value:      "🇲🇷",
		Descriptor: "Flag: Mauritania",
	},
	"1F1F2-1F1F8": {
		Key:        "1F1F2-1F1F8",
		Value:      "🇲🇸",
		Descriptor: "Flag: Montserrat",
	},
	"1F1F2-1F1F9": {
		Key:        "1F1F2-1F1F9",
		Value:      "🇲🇹",
		Descriptor: "Flag: Malta",
	},
	"1F1F2-1F1FA": {
		Key:        "1F1F2-1F1FA",
		Value:      "🇲🇺",
		Descriptor: "Flag: Mauritius",
	},
	"1F1F2-1F1FB": {
		Key:        "1F1F2-1F1FB",
		Value:      "🇲🇻",
		Descriptor: "Flag: Maldives",
	},
	"1F1F2-1F1FC": {
		Key:        "1F1F2-1F1FC",
		Value:      "🇲🇼",
		Descriptor: "Flag: Malawi",
	},
	"1F1F2-1F1FD": {
		Key:        "1F1F2-1F1FD",
		Value:      "🇲🇽",
		Descriptor: "Flag: Mexico",
	},
	"1F1F2-1F1FE": {
		Key:        "1F1F2-1F1FE",
		Value:      "🇲🇾",
		Descriptor: "Flag: Malaysia",
	},
	"1F1F2-1F1FF": {
		Key:        "1F1F2-1F1FF",
		Value:      "🇲🇿",
		Descriptor: "Flag: Mozambique",
	},
	"1F1F3-1F1E6": {
		Key:        "1F1F3-1F1E6",
		Value:      "🇳🇦",
		Descriptor: "Flag: Namibia",
	},
	"1F1F3-1F1E8": {
		Key:        "1F1F3-1F1E8",
		Value:      "🇳🇨",
		Descriptor: "Flag: New Caledonia",
	},
	"1F1F3-1F1EA": {
		Key:        "1F1F3-1F1EA",
		Value:      "🇳🇪",
		Descriptor: "Flag: Niger",
	},
	"1F1F3-1F1EB": {
		Key:        "1F1F3-1F1EB",
		Value:      "🇳🇫",
		Descriptor: "Flag: Norfolk Island",
	},
	"1F1F3-1F1EC": {
		Key:        "1F1F3-1F1EC",
		Value:      "🇳🇬",
		Descriptor: "Flag: Nigeria",
	},
	"1F1F3-1F1EE": {
		Key:        "1F1F3-1F1EE",
		Value:      "🇳🇮",
		Descriptor: "Flag: Nicaragua",
	},
	"1F1F3-1F1F1": {
		Key:        "1F1F3-1F1F1",
		Value:      "🇳🇱",
		Descriptor: "Flag: Netherlands",
	},
	"1F1F3-1F1F4": {
		Key:        "1F1F3-1F1F4",
		Value:      "🇳🇴",
		Descriptor: "Flag: Norway",
	},
	"1F1F3-1F1F5": {
		Key:        "1F1F3-1F1F5",
		Value:      "🇳🇵",
		Descriptor: "Flag: Nepal",
	},
	"1F1F3-1F1F7": {
		Key:        "1F1F3-1F1F7",
		Value:      "🇳🇷",
		Descriptor: "Flag: Nauru",
	},
	"1F1F3-1F1FA": {
		Key:        "1F1F3-1F1FA",
		Value:      "🇳🇺",
		Descriptor: "Flag: Niue",
	},
	"1F1F3-1F1FF": {
		Key:        "1F1F3-1F1FF",
		Value:      "🇳🇿",
		Descriptor: "Flag: New Zealand",
	},
	"1F1F4-1F1F2": {
		Key:        "1F1F4-1F1F2",
		Value:      "🇴🇲",
		Descriptor: "Flag: Oman",
	},
	"1F1F5-1F1E6": {
		Key:        "1F1F5-1F1E6",
		Value:      "🇵🇦",
		Descriptor: "Flag: Panama",
	},
	"1F1F5-1F1EA": {
		Key:        "1F1F5-1F1EA",
		Value:      "🇵🇪",
		Descriptor: "Flag: Peru",
	},
	"1F1F5-1F1EB": {
		Key:        "1F1F5-1F1EB",
		Value:      "🇵🇫",
		Descriptor: "Flag: French Polynesia",
	},
	"1F1F5-1F1EC": {
		Key:        "1F1F5-1F1EC",
		Value:      "🇵🇬",
		Descriptor: "Flag: Papua New Guinea",
	},
	"1F1F5-1F1ED": {
		Key:        "1F1F5-1F1ED",
		Value:      "🇵🇭",
		Descriptor: "Flag: Philippines",
	},
	"1F1F5-1F1F0": {
		Key:        "1F1F5-1F1F0",
		Value:      "🇵🇰",
		Descriptor: "Flag: Pakistan",
	},
	"1F1F5-1F1F1": {
		Key:        "1F1F5-1F1F1",
		Value:      "🇵🇱",
		Descriptor: "Flag: Poland",
	},
	"1F1F5-1F1F2": {
		Key:        "1F1F5-1F1F2",
		Value:      "🇵🇲",
		Descriptor: "Flag: St. Pierre \u0026 Miquelon",
	},
	"1F1F5-1F1F3": {
		Key:        "1F1F5-1F1F3",
		Value:      "🇵🇳",
		Descriptor: "Flag: Pitcairn Islands",
	},
	"1F1F5-1F1F7": {
		Key:        "1F1F5-1F1F7",
		Value:      "🇵🇷",
		Descriptor: "Flag: Puerto Rico",
	},
	"1F1F5-1F1F8": {
		Key:        "1F1F5-1F1F8",
		Value:      "🇵🇸",
		Descriptor: "Flag: Palestinian Territories",
	},
	"1F1F5-1F1F9": {
		Key:        "1F1F5-1F1F9",
		Value:      "🇵🇹",
		Descriptor: "Flag: Portugal",
	},
	"1F1F5-1F1FC": {
		Key:        "1F1F5-1F1FC",
		Value:      "🇵🇼",
		Descriptor: "Flag: Palau",
	},
	"1F1F5-1F1FE": {
		Key:        "1F1F5-1F1FE",
		Value:      "🇵🇾",
		Descriptor: "Flag: Paraguay",
	},
	"1F1F6-1F1E6": {
		Key:        "1F1F6-1F1E6",
		Value:      "🇶🇦",
		Descriptor: "Flag: Qatar",
	},
	"1F1F7-1F1EA": {
		Key:        "1F1F7-1F1EA",
		Value:      "🇷🇪",
		Descriptor: "Flag: Réunion",
	},
	"1F1F7-1F1F4": {
		Key:        "1F1F7-1F1F4",
		Value:      "🇷🇴",
		Descriptor: "Flag: Romania",
	},
	"1F1F7-1F1F8": {
		Key:        "1F1F7-1F1F8",
		Value:      "🇷🇸",
		Descriptor: "Flag: Serbia",
	},
	"1F1F7-1F1FA": {
		Key:        "1F1F7-1F1FA",
		Value:      "🇷🇺",
		Descriptor: "Flag: Russia",
	},
	"1F1F7-1F1FC": {
		Key:        "1F1F7-1F1FC",
		Value:      "🇷🇼",
		Descriptor: "Flag: Rwanda",
	},
	"1F1F8-1F1E6": {
		Key:        "1F1F8-1F1E6",
		Value:      "🇸🇦",
		Descriptor: "Flag: Saudi Arabia",
	},
	"1F1F8-1F1E7": {
		Key:        "1F1F8-1F1E7",
		Value:      "🇸🇧",
		Descriptor: "Flag: Solomon Islands",
	},
	"1F1F8-1F1E8": {
		Key:        "1F1F8-1F1E8",
		Value:      "🇸🇨",
		Descriptor: "Flag: Seychelles",
	},
	"1F1F8-1F1E9": {
		Key:        "1F1F8-1F1E9",
		Value:      "🇸🇩",
		Descriptor: "Flag: Sudan",
	},
	"1F1F8-1F1EA": {
		Key:        "1F1F8-1F1EA",
		Value:      "🇸🇪",
		Descriptor: "Flag: Sweden",
	},
	"1F1F8-1F1EC": {
		Key:        "1F1F8-1F1EC",
		Value:      "🇸🇬",
		Descriptor: "Flag: Singapore",
	},
	"1F1F8-1F1ED": {
		Key:        "1F1F8-1F1ED",
		Value:      "🇸🇭",
		Descriptor: "Flag: St. Helena",
	},
	"1F1F8-1F1EE": {
		Key:        "1F1F8-1F1EE",
		Value:      "🇸🇮",
		Descriptor: "Flag: Slovenia",
	},
	"1F1F8-1F1EF": {
		Key:        "1F1F8-1F1EF",
		Value:      "🇸🇯",
		Descriptor: "Flag: Svalbard \u0026 Jan Mayen",
	},
	"1F1F8-1F1F0": {
		Key:        "1F1F8-1F1F0",
		Value:      "🇸🇰",
		Descriptor: "Flag: Slovakia",
	},
	"1F1F8-1F1F1": {
		Key:        "1F1F8-1F1F1",
		Value:      "🇸🇱",
		Descriptor: "Flag: Sierra Leone",
	},
	"1F1F8-1F1F2": {
		Key:        "1F1F8-1F1F2",
		Value:      "🇸🇲",
		Descriptor: "Flag: San Marino",
	},
	"1F1F8-1F1F3": {
		Key:        "1F1F8-1F1F3",
		Value:      "🇸🇳",
		Descriptor: "Flag: Senegal",
	},
	"1F1F8-1F1F4": {
		Key:        "1F1F8-1F1F4",
		Value:      "🇸🇴",
		Descriptor: "Flag: Somalia",
	},
	"1F1F8-1F1F7": {
		Key:        "1F1F8-1F1F7",
		Value:      "🇸🇷",
		Descriptor: "Flag: Suriname",
	},
	"1F1F8-1F1F8": {
		Key:        "1F1F8-1F1F8",
		Value:      "🇸🇸",
		Descriptor: "Flag: South Sudan",
	},
	"1F1F8-1F1F9": {
		Key:        "1F1F8-1F1F9",
		Value:      "🇸🇹",
		Descriptor: "Flag: São Tomé \u0026 Príncipe",
	},
	"1F1F8-1F1FB": {
		Key:        "1F1F8-1F1FB",
		Value:      "🇸🇻",
		Descriptor: "Flag: El Salvador",
	},
	"1F1F8-1F1FD": {
		Key:        "1F1F8-1F1FD",
		Value:      "🇸🇽",
		Descriptor: "Flag: Sint Maarten",
	},
	"1F1F8-1F1FE": {
		Key:        "1F1F8-1F1FE",
		Value:      "🇸🇾",
		Descriptor: "Flag: Syria",
	},
	"1F1F8-1F1FF": {
		Key:        "1F1F8-1F1FF",
		Value:      "🇸🇿",
		Descriptor: "Flag: Eswatini",
	},
	"1F1F9-1F1E6": {
		Key:        "1F1F9-1F1E6",
		Value:      "🇹🇦",
		Descriptor: "Flag: Tristan Da Cunha",
	},
	"1F1F9-1F1E8": {
		Key:        "1F1F9-1F1E8",
		Value:      "🇹🇨",
		Descriptor: "Flag: Turks \u0026 Caicos Islands",
	},
	"1F1F9-1F1E9": {
		Key:        "1F1F9-1F1E9",
		Value:      "🇹🇩",
		Descriptor: "Flag: Chad",
	},
	"1F1F9-1F1EB": {
		Key:        "1F1F9-1F1EB",
		Value:      "🇹🇫",
		Descriptor: "Flag: French Southern Territories",
	},
	"1F1F9-1F1EC": {
		Key:        "1F1F9-1F1EC",
		Value:      "🇹🇬",
		Descriptor: "Flag: Togo",
	},
	"1F1F9-1F1ED": {
		Key:        "1F1F9-1F1ED",
		Value:      "🇹🇭",
		Descriptor: "Flag: Thailand",
	},
	"1F1F9-1F1EF": {
		Key:        "1F1F9-1F1EF",
		Value:      "🇹🇯",
		Descriptor: "Flag: Tajikistan",
	},
	"1F1F9-1F1F0": {
		Key:        "1F1F9-1F1F0",
		Value:      "🇹🇰",
		Descriptor: "Flag: Tokelau",
	},
	"1F1F9-1F1F1": {
		Key:        "1F1F9-1F1F1",
		Value:      "🇹🇱",
		Descriptor: "Flag: Timor-Leste",
	},
	"1F1F9-1F1F2": {
		Key:        "1F1F9-1F1F2",
		Value:      "🇹🇲",
		Descriptor: "Flag: Turkmenistan",
	},
	"1F1F9-1F1F3": {
		Key:        "1F1F9-1F1F3",
		Value:      "🇹🇳",
		Descriptor: "Flag: Tunisia",
	},
	"1F1F9-1F1F4": {
		Key:        "1F1F9-1F1F4",
		Value:      "🇹🇴",
		Descriptor: "Flag: Tonga",
	},
	"1F1F9-1F1F7": {
		Key:        "1F1F9-1F1F7",
		Value:      "🇹🇷",
		Descriptor: "Flag: Turkey",
	},
	"1F1F9-1F1F9": {
		Key:        "1F1F9-1F1F9",
		Value:      "🇹🇹",
		Descriptor: "Flag: Trinidad \u0026 Tobago",
	},
	"1F1F9-1F1FB": {
		Key:        "1F1F9-1F1FB",
		Value:      "🇹🇻",
		Descriptor: "Flag: Tuvalu",
	},
	"1F1F9-1F1FC": {
		Key:        "1F1F9-1F1FC",
		Value:      "🇹🇼",
		Descriptor: "Flag: Taiwan",
	},
	"1F1F9-1F1FF": {
		Key:        "1F1F9-1F1FF",
		Value:      "🇹🇿",
		Descriptor: "Flag: Tanzania",
	},
	"1F1FA-1F1E6": {
		Key:        "1F1FA-1F1E6",
		Value:      "🇺🇦",
		Descriptor: "Flag: Ukraine",
	},
	"1F1FA-1F1EC": {
		Key:        "1F1FA-1F1EC",
		Value:      "🇺🇬",
		Descriptor: "Flag: Uganda",
	},
	"1F1FA-1F1F2": {
		Key:        "1F1FA-1F1F2",
		Value:      "🇺🇲",
		Descriptor: "Flag: U.S. Outlying Islands",
	},
	"1F1FA-1F1F3": {
		Key:        "1F1FA-1F1F3",
		Value:      "🇺🇳",
		Descriptor: "Flag: United Nations",
	},
	"1F1FA-1F1F8": {
		Key:        "1F1FA-1F1F8",
		Value:      "🇺🇸",
		Descriptor: "Flag: United States",
	},
	"1F1FA-1F1FE": {
		Key:        "1F1FA-1F1FE",
		Value:      "🇺🇾",
		Descriptor: "Flag: Uruguay",
	},
	"1F1FA-1F1FF": {
		Key:        "1F1FA-1F1FF",
		Value:      "🇺🇿",
		Descriptor: "Flag: Uzbekistan",
	},
	"1F1FB-1F1E6": {
		Key:        "1F1FB-1F1E6",
		Value:      "🇻🇦",
		Descriptor: "Flag: Vatican City",
	},
	"1F1FB-1F1E8": {
		Key:        "1F1FB-1F1E8",
		Value:      "🇻🇨",
		Descriptor: "Flag: St. Vincent \u0026 Grenadines",
	},
	"1F1FB-1F1EA": {
		Key:        "1F1FB-1F1EA",
		Value:      "🇻🇪",
		Descriptor: "Flag: Venezuela",
	},
	"1F1FB-1F1EC": {
		Key:        "1F1FB-1F1EC",
		Value:      "🇻🇬",
		Descriptor: "Flag: British Virgin Islands",
	},
	"1F1FB-1F1EE": {
		Key:        "1F1FB-1F1EE",
		Value:      "🇻🇮",
		Descriptor: "Flag: U.S. Virgin Islands",
	},
	"1F1FB-1F1F3": {
		Key:        "1F1FB-1F1F3",
		Value:      "🇻🇳",
		Descriptor: "Flag: Vietnam",
	},
	"1F1FB-1F1FA": {
		Key:        "1F1FB-1F1FA",
		Value:      "🇻🇺",
		Descriptor: "Flag: Vanuatu",
	},
	"1F1FC-1F1EB": {
		Key:        "1F1FC-1F1EB",
		Value:      "🇼🇫",
		Descriptor: "Flag: Wallis \u0026 Futuna",
	},
	"1F1FC-1F1F8": {
		Key:        "1F1FC-1F1F8",
		Value:      "🇼🇸",
		Descriptor: "Flag: Samoa",
	},
	"1F1FD-1F1F0": {
		Key:        "1F1FD-1F1F0",
		Value:      "🇽🇰",
		Descriptor: "Flag: Kosovo",
	},
	"1F1FE-1F1EA": {
		Key:        "1F1FE-1F1EA",
		Value:      "🇾🇪",
		Descriptor: "Flag: Yemen",
	},
	"1F1FE-1F1F9": {
		Key:        "1F1FE-1F1F9",
		Value:      "🇾🇹",
		Descriptor: "Flag: Mayotte",
	},
	"1F1FF-1F1E6": {
		Key:        "1F1FF-1F1E6",
		Value:      "🇿🇦",
		Descriptor: "Flag: South Africa",
	},
	"1F1FF-1F1F2": {
		Key:        "1F1FF-1F1F2",
		Value:      "🇿🇲",
		Descriptor: "Flag: Zambia",
	},
	"1F1FF-1F1FC": {
		Key:        "1F1FF-1F1FC",
		Value:      "🇿🇼",
		Descriptor: "Flag: Zimbabwe",
	},
	"1F201": {
		Key:        "1F201",
		Value:      "🈁",
		Descriptor: "Japanese “Here” Button",
	},
	"1F202-FE0F": {
		Key:        "1F202-FE0F",
		Value:      "🈂️",
		Descriptor: "Japanese “Service Charge” Button",
	},
	"1F21A": {
		Key:        "1F21A",
		Value:      "🈚",
		Descriptor: "Japanese “Free of Charge” Button",
	},
	"1F22F": {
		Key:        "1F22F",
		Value:      "🈯",
		Descriptor: "Japanese “Reserved” Button",
	},
	"1F232": {
		Key:        "1F232",
		Value:      "🈲",
		Descriptor: "Japanese “Prohibited” Button",
	},
	"1F233": {
		Key:        "1F233",
		Value:      "🈳",
		Descriptor: "Japanese “Vacancy” Button",
	},
	"1F234": {
		Key:        "1F234",
		Value:      "🈴",
		Descriptor: "Japanese “Passing Grade” Button",
	},
	"1F235": {
		Key:        "1F235",
		Value:      "🈵",
		Descriptor: "Japanese “No Vacancy” Button",
	},
	"1F236": {
		Key:        "1F236",
		Value:      "🈶",
		Descriptor: "Japanese “Not Free of Charge” Button",
	},
	"1F237-FE0F": {
		Key:        "1F237-FE0F",
		Value:      "🈷️",
		Descriptor: "Japanese “Monthly Amount” Button",
	},
	"1F238": {
		Key:        "1F238",
		Value:      "🈸",
		Descriptor: "Japanese “Application” Button",
	},
	"1F239": {
		Key:        "1F239",
		Value:      "🈹",
		Descriptor: "Japanese “Discount” Button",
	},
	"1F23A": {
		Key:        "1F23A",
		Value:      "🈺",
		Descriptor: "Japanese “Open for Business” Button",
	},
	"1F250": {
		Key:        "1F250",
		Value:      "🉐",
		Descriptor: "Japanese “Bargain” Button",
	},
	"1F251": {
		Key:        "1F251",
		Value:      "🉑",
		Descriptor: "Japanese “Acceptable” Button",
	},
	"1F300": {
		Key:        "1F300",
		Value:      "🌀",
		Descriptor: "Cyclone",
	},
	"1F301": {
		Key:        "1F301",
		Value:      "🌁",
		Descriptor: "Foggy",
	},
	"1F302": {
		Key:        "1F302",
		Value:      "🌂",
		Descriptor: "Closed Umbrella",
	},
	"1F303": {
		Key:        "1F303",
		Value:      "🌃",
		Descriptor: "Night With Stars",
	},
	"1F304": {
		Key:        "1F304",
		Value:      "🌄",
		Descriptor: "Sunrise Over Mountains",
	},
	"1F305": {
		Key:        "1F305",
		Value:      "🌅",
		Descriptor: "Sunrise",
	},
	"1F306": {
		Key:        "1F306",
		Value:      "🌆",
		Descriptor: "Cityscape at Dusk",
	},
	"1F307": {
		Key:        "1F307",
		Value:      "🌇",
		Descriptor: "Sunset",
	},
	"1F308": {
		Key:        "1F308",
		Value:      "🌈",
		Descriptor: "Rainbow",
	},
	"1F309": {
		Key:        "1F309",
		Value:      "🌉",
		Descriptor: "Bridge at Night",
	},
	"1F30A": {
		Key:        "1F30A",
		Value:      "🌊",
		Descriptor: "Water Wave",
	},
	"1F30B": {
		Key:        "1F30B",
		Value:      "🌋",
		Descriptor: "Volcano",
	},
	"1F30C": {
		Key:        "1F30C",
		Value:      "🌌",
		Descriptor: "Milky Way",
	},
	"1F30D": {
		Key:        "1F30D",
		Value:      "🌍",
		Descriptor: "Globe Showing Europe-Africa",
	},
	"1F30E": {
		Key:        "1F30E",
		Value:      "🌎",
		Descriptor: "Globe Showing Americas",
	},
	"1F30F": {
		Key:        "1F30F",
		Value:      "🌏",
		Descriptor: "Globe Showing Asia-Australia",
	},
	"1F310": {
		Key:        "1F310",
		Value:      "🌐",
		Descriptor: "Globe With Meridians",
	},
	"1F311": {
		Key:        "1F311",
		Value:      "🌑",
		Descriptor: "New Moon",
	},
	"1F312": {
		Key:        "1F312",
		Value:      "🌒",
		Descriptor: "Waxing Crescent Moon",
	},
	"1F313": {
		Key:        "1F313",
		Value:      "🌓",
		Descriptor: "First Quarter Moon",
	},
	"1F314": {
		Key:        "1F314",
		Value:      "🌔",
		Descriptor: "Waxing Gibbous Moon",
	},
	"1F315": {
		Key:        "1F315",
		Value:      "🌕",
		Descriptor: "Full Moon",
	},
	"1F316": {
		Key:        "1F316",
		Value:      "🌖",
		Descriptor: "Waning Gibbous Moon",
	},
	"1F317": {
		Key:        "1F317",
		Value:      "🌗",
		Descriptor: "Last Quarter Moon",
	},
	"1F318": {
		Key:        "1F318",
		Value:      "🌘",
		Descriptor: "Waning Crescent Moon",
	},
	"1F319": {
		Key:        "1F319",
		Value:      "🌙",
		Descriptor: "Crescent Moon",
	},
	"1F31A": {
		Key:        "1F31A",
		Value:      "🌚",
		Descriptor: "New Moon Face",
	},
	"1F31B": {
		Key:        "1F31B",
		Value:      "🌛",
		Descriptor: "First Quarter Moon Face",
	},
	"1F31C": {
		Key:        "1F31C",
		Value:      "🌜",
		Descriptor: "Last Quarter Moon Face",
	},
	"1F31D": {
		Key:        "1F31D",
		Value:      "🌝",
		Descriptor: "Full Moon Face",
	},
	"1F31E": {
		Key:        "1F31E",
		Value:      "🌞",
		Descriptor: "Sun With Face",
	},
	"1F31F": {
		Key:        "1F31F",
		Value:      "🌟",
		Descriptor: "Glowing Star",
	},
	"1F320": {
		Key:        "1F320",
		Value:      "🌠",
		Descriptor: "Shooting Star",
	},
	"1F321-FE0F": {
		Key:        "1F321-FE0F",
		Value:      "🌡️",
		Descriptor: "Thermometer",
	},
	"1F324-FE0F": {
		Key:        "1F324-FE0F",
		Value:      "🌤️",
		Descriptor: "Sun Behind Small Cloud",
	},
	"1F325-FE0F": {
		Key:        "1F325-FE0F",
		Value:      "🌥️",
		Descriptor: "Sun Behind Large Cloud",
	},
	"1F326-FE0F": {
		Key:        "1F326-FE0F",
		Value:      "🌦️",
		Descriptor: "Sun Behind Rain Cloud",
	},
	"1F327-FE0F": {
		Key:        "1F327-FE0F",
		Value:      "🌧️",
		Descriptor: "Cloud With Rain",
	},
	"1F328-FE0F": {
		Key:        "1F328-FE0F",
		Value:      "🌨️",
		Descriptor: "Cloud With Snow",
	},
	"1F329-FE0F": {
		Key:        "1F329-FE0F",
		Value:      "🌩️",
		Descriptor: "Cloud With Lightning",
	},
	"1F32A-FE0F": {
		Key:        "1F32A-FE0F",
		Value:      "🌪️",
		Descriptor: "Tornado",
	},
	"1F32B-FE0F": {
		Key:        "1F32B-FE0F",
		Value:      "🌫️",
		Descriptor: "Fog",
	},
	"1F32C-FE0F": {
		Key:        "1F32C-FE0F",
		Value:      "🌬️",
		Descriptor: "Wind Face",
	},
	"1F32D": {
		Key:        "1F32D",
		Value:      "🌭",
		Descriptor: "Hot Dog",
	},
	"1F32E": {
		Key:        "1F32E",
		Value:      "🌮",
		Descriptor: "Taco",
	},
	"1F32F": {
		Key:        "1F32F",
		Value:      "🌯",
		Descriptor: "Burrito",
	},
	"1F330": {
		Key:        "1F330",
		Value:      "🌰",
		Descriptor: "Chestnut",
	},
	"1F331": {
		Key:        "1F331",
		Value:      "🌱",
		Descriptor: "Seedling",
	},
	"1F332": {
		Key:        "1F332",
		Value:      "🌲",
		Descriptor: "Evergreen Tree",
	},
	"1F333": {
		Key:        "1F333",
		Value:      "🌳",
		Descriptor: "Deciduous Tree",
	},
	"1F334": {
		Key:        "1F334",
		Value:      "🌴",
		Descriptor: "Palm Tree",
	},
	"1F335": {
		Key:        "1F335",
		Value:      "🌵",
		Descriptor: "Cactus",
	},
	"1F336-FE0F": {
		Key:        "1F336-FE0F",
		Value:      "🌶️",
		Descriptor: "Hot Pepper",
	},
	"1F337": {
		Key:        "1F337",
		Value:      "🌷",
		Descriptor: "Tulip",
	},
	"1F338": {
		Key:        "1F338",
		Value:      "🌸",
		Descriptor: "Cherry Blossom",
	},
	"1F339": {
		Key:        "1F339",
		Value:      "🌹",
		Descriptor: "Rose",
	},
	"1F33A": {
		Key:        "1F33A",
		Value:      "🌺",
		Descriptor: "Hibiscus",
	},
	"1F33B": {
		Key:        "1F33B",
		Value:      "🌻",
		Descriptor: "Sunflower",
	},
	"1F33C": {
		Key:        "1F33C",
		Value:      "🌼",
		Descriptor: "Blossom",
	},
	"1F33D": {
		Key:        "1F33D",
		Value:      "🌽",
		Descriptor: "Ear of Corn",
	},
	"1F33E": {
		Key:        "1F33E",
		Value:      "🌾",
		Descriptor: "Sheaf of Rice",
	},
	"1F33F": {
		Key:        "1F33F",
		Value:      "🌿",
		Descriptor: "Herb",
	},
	"1F340": {
		Key:        "1F340",
		Value:      "🍀",
		Descriptor: "Four Leaf Clover",
	},
	"1F341": {
		Key:        "1F341",
		Value:      "🍁",
		Descriptor: "Maple Leaf",
	},
	"1F342": {
		Key:        "1F342",
		Value:      "🍂",
		Descriptor: "Fallen Leaf",
	},
	"1F343": {
		Key:        "1F343",
		Value:      "🍃",
		Descriptor: "Leaf Fluttering in Wind",
	},
	"1F344": {
		Key:        "1F344",
		Value:      "🍄",
		Descriptor: "Mushroom",
	},
	"1F345": {
		Key:        "1F345",
		Value:      "🍅",
		Descriptor: "Tomato",
	},
	"1F346": {
		Key:        "1F346",
		Value:      "🍆",
		Descriptor: "Eggplant",
	},
	"1F347": {
		Key:        "1F347",
		Value:      "🍇",
		Descriptor: "Grapes",
	},
	"1F348": {
		Key:        "1F348",
		Value:      "🍈",
		Descriptor: "Melon",
	},
	"1F349": {
		Key:        "1F349",
		Value:      "🍉",
		Descriptor: "Watermelon",
	},
	"1F34A": {
		Key:        "1F34A",
		Value:      "🍊",
		Descriptor: "Tangerine",
	},
	"1F34B": {
		Key:        "1F34B",
		Value:      "🍋",
		Descriptor: "Lemon",
	},
	"1F34C": {
		Key:        "1F34C",
		Value:      "🍌",
		Descriptor: "Banana",
	},
	"1F34D": {
		Key:        "1F34D",
		Value:      "🍍",
		Descriptor: "Pineapple",
	},
	"1F34E": {
		Key:        "1F34E",
		Value:      "🍎",
		Descriptor: "Red Apple",
	},
	"1F34F": {
		Key:        "1F34F",
		Value:      "🍏",
		Descriptor: "Green Apple",
	},
	"1F350": {
		Key:        "1F350",
		Value:      "🍐",
		Descriptor: "Pear",
	},
	"1F351": {
		Key:        "1F351",
		Value:      "🍑",
		Descriptor: "Peach",
	},
	"1F352": {
		Key:        "1F352",
		Value:      "🍒",
		Descriptor: "Cherries",
	},
	"1F353": {
		Key:        "1F353",
		Value:      "🍓",
		Descriptor: "Strawberry",
	},
	"1F354": {
		Key:        "1F354",
		Value:      "🍔",
		Descriptor: "Hamburger",
	},
	"1F355": {
		Key:        "1F355",
		Value:      "🍕",
		Descriptor: "Pizza",
	},
	"1F356": {
		Key:        "1F356",
		Value:      "🍖",
		Descriptor: "Meat on Bone",
	},
	"1F357": {
		Key:        "1F357",
		Value:      "🍗",
		Descriptor: "Poultry Leg",
	},
	"1F358": {
		Key:        "1F358",
		Value:      "🍘",
		Descriptor: "Rice Cracker",
	},
	"1F359": {
		Key:        "1F359",
		Value:      "🍙",
		Descriptor: "Rice Ball",
	},
	"1F35A": {
		Key:        "1F35A",
		Value:      "🍚",
		Descriptor: "Cooked Rice",
	},
	"1F35B": {
		Key:        "1F35B",
		Value:      "🍛",
		Descriptor: "Curry Rice",
	},
	"1F35C": {
		Key:        "1F35C",
		Value:      "🍜",
		Descriptor: "Steaming Bowl",
	},
	"1F35D": {
		Key:        "1F35D",
		Value:      "🍝",
		Descriptor: "Spaghetti",
	},
	"1F35E": {
		Key:        "1F35E",
		Value:      "🍞",
		Descriptor: "Bread",
	},
	"1F35F": {
		Key:        "1F35F",
		Value:      "🍟",
		Descriptor: "French Fries",
	},
	"1F360": {
		Key:        "1F360",
		Value:      "🍠",
		Descriptor: "Roasted Sweet Potato",
	},
	"1F361": {
		Key:        "1F361",
		Value:      "🍡",
		Descriptor: "Dango",
	},
	"1F362": {
		Key:        "1F362",
		Value:      "🍢",
		Descriptor: "Oden",
	},
	"1F363": {
		Key:        "1F363",
		Value:      "🍣",
		Descriptor: "Sushi",
	},
	"1F364": {
		Key:        "1F364",
		Value:      "🍤",
		Descriptor: "Fried Shrimp",
	},
	"1F365": {
		Key:        "1F365",
		Value:      "🍥",
		Descriptor: "Fish Cake With Swirl",
	},
	"1F366": {
		Key:        "1F366",
		Value:      "🍦",
		Descriptor: "Soft Ice Cream",
	},
	"1F367": {
		Key:        "1F367",
		Value:      "🍧",
		Descriptor: "Shaved Ice",
	},
	"1F368": {
		Key:        "1F368",
		Value:      "🍨",
		Descriptor: "Ice Cream",
	},
	"1F369": {
		Key:        "1F369",
		Value:      "🍩",
		Descriptor: "Doughnut",
	},
	"1F36A": {
		Key:        "1F36A",
		Value:      "🍪",
		Descriptor: "Cookie",
	},
	"1F36B": {
		Key:        "1F36B",
		Value:      "🍫",
		Descriptor: "Chocolate Bar",
	},
	"1F36C": {
		Key:        "1F36C",
		Value:      "🍬",
		Descriptor: "Candy",
	},
	"1F36D": {
		Key:        "1F36D",
		Value:      "🍭",
		Descriptor: "Lollipop",
	},
	"1F36E": {
		Key:        "1F36E",
		Value:      "🍮",
		Descriptor: "Custard",
	},
	"1F36F": {
		Key:        "1F36F",
		Value:      "🍯",
		Descriptor: "Honey Pot",
	},
	"1F370": {
		Key:        "1F370",
		Value:      "🍰",
		Descriptor: "Shortcake",
	},
	"1F371": {
		Key:        "1F371",
		Value:      "🍱",
		Descriptor: "Bento Box",
	},
	"1F372": {
		Key:        "1F372",
		Value:      "🍲",
		Descriptor: "Pot of Food",
	},
	"1F373": {
		Key:        "1F373",
		Value:      "🍳",
		Descriptor: "Cooking",
	},
	"1F374": {
		Key:        "1F374",
		Value:      "🍴",
		Descriptor: "Fork and Knife",
	},
	"1F375": {
		Key:        "1F375",
		Value:      "🍵",
		Descriptor: "Teacup Without Handle",
	},
	"1F376": {
		Key:        "1F376",
		Value:      "🍶",
		Descriptor: "Sake",
	},
	"1F377": {
		Key:        "1F377",
		Value:      "🍷",
		Descriptor: "Wine Glass",
	},
	"1F378": {
		Key:        "1F378",
		Value:      "🍸",
		Descriptor: "Cocktail Glass",
	},
	"1F379": {
		Key:        "1F379",
		Value:      "🍹",
		Descriptor: "Tropical Drink",
	},
	"1F37A": {
		Key:        "1F37A",
		Value:      "🍺",
		Descriptor: "Beer Mug",
	},
	"1F37B": {
		Key:        "1F37B",
		Value:      "🍻",
		Descriptor: "Clinking Beer Mugs",
	},
	"1F37C": {
		Key:        "1F37C",
		Value:      "🍼",
		Descriptor: "Baby Bottle",
	},
	"1F37D-FE0F": {
		Key:        "1F37D-FE0F",
		Value:      "🍽️",
		Descriptor: "Fork and Knife With Plate",
	},
	"1F37E": {
		Key:        "1F37E",
		Value:      "🍾",
		Descriptor: "Bottle With Popping Cork",
	},
	"1F37F": {
		Key:        "1F37F",
		Value:      "🍿",
		Descriptor: "Popcorn",
	},
	"1F380": {
		Key:        "1F380",
		Value:      "🎀",
		Descriptor: "Ribbon",
	},
	"1F381": {
		Key:        "1F381",
		Value:      "🎁",
		Descriptor: "Wrapped Gift",
	},
	"1F382": {
		Key:        "1F382",
		Value:      "🎂",
		Descriptor: "Birthday Cake",
	},
	"1F383": {
		Key:        "1F383",
		Value:      "🎃",
		Descriptor: "Jack-O-Lantern",
	},
	"1F384": {
		Key:        "1F384",
		Value:      "🎄",
		Descriptor: "Christmas Tree",
	},
	"1F385": {
		Key:        "1F385",
		Value:      "🎅",
		Descriptor: "Santa Claus",
	},
	"1F385-1F3FB": {
		Key:        "1F385-1F3FB",
		Value:      "🎅🏻",
		Descriptor: "Santa Claus: Light Skin Tone",
	},
	"1F385-1F3FC": {
		Key:        "1F385-1F3FC",
		Value:      "🎅🏼",
		Descriptor: "Santa Claus: Medium-Light Skin Tone",
	},
	"1F385-1F3FD": {
		Key:        "1F385-1F3FD",
		Value:      "🎅🏽",
		Descriptor: "Santa Claus: Medium Skin Tone",
	},
	"1F385-1F3FE": {
		Key:        "1F385-1F3FE",
		Value:      "🎅🏾",
		Descriptor: "Santa Claus: Medium-Dark Skin Tone",
	},
	"1F385-1F3FF": {
		Key:        "1F385-1F3FF",
		Value:      "🎅🏿",
		Descriptor: "Santa Claus: Dark Skin Tone",
	},
	"1F386": {
		Key:        "1F386",
		Value:      "🎆",
		Descriptor: "Fireworks",
	},
	"1F387": {
		Key:        "1F387",
		Value:      "🎇",
		Descriptor: "Sparkler",
	},
	"1F388": {
		Key:        "1F388",
		Value:      "🎈",
		Descriptor: "Balloon",
	},
	"1F389": {
		Key:        "1F389",
		Value:      "🎉",
		Descriptor: "Party Popper",
	},
	"1F38A": {
		Key:        "1F38A",
		Value:      "🎊",
		Descriptor: "Confetti Ball",
	},
	"1F38B": {
		Key:        "1F38B",
		Value:      "🎋",
		Descriptor: "Tanabata Tree",
	},
	"1F38C": {
		Key:        "1F38C",
		Value:      "🎌",
		Descriptor: "Crossed Flags",
	},
	"1F38D": {
		Key:        "1F38D",
		Value:      "🎍",
		Descriptor: "Pine Decoration",
	},
	"1F38E": {
		Key:        "1F38E",
		Value:      "🎎",
		Descriptor: "Japanese Dolls",
	},
	"1F38F": {
		Key:        "1F38F",
		Value:      "🎏",
		Descriptor: "Carp Streamer",
	},
	"1F390": {
		Key:        "1F390",
		Value:      "🎐",
		Descriptor: "Wind Chime",
	},
	"1F391": {
		Key:        "1F391",
		Value:      "🎑",
		Descriptor: "Moon Viewing Ceremony",
	},
	"1F392": {
		Key:        "1F392",
		Value:      "🎒",
		Descriptor: "Backpack",
	},
	"1F393": {
		Key:        "1F393",
		Value:      "🎓",
		Descriptor: "Graduation Cap",
	},
	"1F396-FE0F": {
		Key:        "1F396-FE0F",
		Value:      "🎖️",
		Descriptor: "Military Medal",
	},
	"1F397-FE0F": {
		Key:        "1F397-FE0F",
		Value:      "🎗️",
		Descriptor: "Reminder Ribbon",
	},
	"1F399-FE0F": {
		Key:        "1F399-FE0F",
		Value:      "🎙️",
		Descriptor: "Studio Microphone",
	},
	"1F39A-FE0F": {
		Key:        "1F39A-FE0F",
		Value:      "🎚️",
		Descriptor: "Level Slider",
	},
	"1F39B-FE0F": {
		Key:        "1F39B-FE0F",
		Value:      "🎛️",
		Descriptor: "Control Knobs",
	},
	"1F39E-FE0F": {
		Key:        "1F39E-FE0F",
		Value:      "🎞️",
		Descriptor: "Film Frames",
	},
	"1F39F-FE0F": {
		Key:        "1F39F-FE0F",
		Value:      "🎟️",
		Descriptor: "Admission Tickets",
	},
	"1F3A0": {
		Key:        "1F3A0",
		Value:      "🎠",
		Descriptor: "Carousel Horse",
	},
	"1F3A1": {
		Key:        "1F3A1",
		Value:      "🎡",
		Descriptor: "Ferris Wheel",
	},
	"1F3A2": {
		Key:        "1F3A2",
		Value:      "🎢",
		Descriptor: "Roller Coaster",
	},
	"1F3A3": {
		Key:        "1F3A3",
		Value:      "🎣",
		Descriptor: "Fishing Pole",
	},
	"1F3A4": {
		Key:        "1F3A4",
		Value:      "🎤",
		Descriptor: "Microphone",
	},
	"1F3A5": {
		Key:        "1F3A5",
		Value:      "🎥",
		Descriptor: "Movie Camera",
	},
	"1F3A6": {
		Key:        "1F3A6",
		Value:      "🎦",
		Descriptor: "Cinema",
	},
	"1F3A7": {
		Key:        "1F3A7",
		Value:      "🎧",
		Descriptor: "Headphone",
	},
	"1F3A8": {
		Key:        "1F3A8",
		Value:      "🎨",
		Descriptor: "Artist Palette",
	},
	"1F3A9": {
		Key:        "1F3A9",
		Value:      "🎩",
		Descriptor: "Top Hat",
	},
	"1F3AA": {
		Key:        "1F3AA",
		Value:      "🎪",
		Descriptor: "Circus Tent",
	},
	"1F3AB": {
		Key:        "1F3AB",
		Value:      "🎫",
		Descriptor: "Ticket",
	},
	"1F3AC": {
		Key:        "1F3AC",
		Value:      "🎬",
		Descriptor: "Clapper Board",
	},
	"1F3AD": {
		Key:        "1F3AD",
		Value:      "🎭",
		Descriptor: "Performing Arts",
	},
	"1F3AE": {
		Key:        "1F3AE",
		Value:      "🎮",
		Descriptor: "Video Game",
	},
	"1F3AF": {
		Key:        "1F3AF",
		Value:      "🎯",
		Descriptor: "Direct Hit",
	},
	"1F3B0": {
		Key:        "1F3B0",
		Value:      "🎰",
		Descriptor: "Slot Machine",
	},
	"1F3B1": {
		Key:        "1F3B1",
		Value:      "🎱",
		Descriptor: "Pool 8 Ball",
	},
	"1F3B2": {
		Key:        "1F3B2",
		Value:      "🎲",
		Descriptor: "Game Die",
	},
	"1F3B3": {
		Key:        "1F3B3",
		Value:      "🎳",
		Descriptor: "Bowling",
	},
	"1F3B4": {
		Key:        "1F3B4",
		Value:      "🎴",
		Descriptor: "Flower Playing Cards",
	},
	"1F3B5": {
		Key:        "1F3B5",
		Value:      "🎵",
		Descriptor: "Musical Note",
	},
	"1F3B6": {
		Key:        "1F3B6",
		Value:      "🎶",
		Descriptor: "Musical Notes",
	},
	"1F3B7": {
		Key:        "1F3B7",
		Value:      "🎷",
		Descriptor: "Saxophone",
	},
	"1F3B8": {
		Key:        "1F3B8",
		Value:      "🎸",
		Descriptor: "Guitar",
	},
	"1F3B9": {
		Key:        "1F3B9",
		Value:      "🎹",
		Descriptor: "Musical Keyboard",
	},
	"1F3BA": {
		Key:        "1F3BA",
		Value:      "🎺",
		Descriptor: "Trumpet",
	},
	"1F3BB": {
		Key:        "1F3BB",
		Value:      "🎻",
		Descriptor: "Violin",
	},
	"1F3BC": {
		Key:        "1F3BC",
		Value:      "🎼",
		Descriptor: "Musical Score",
	},
	"1F3BD": {
		Key:        "1F3BD",
		Value:      "🎽",
		Descriptor: "Running Shirt",
	},
	"1F3BE": {
		Key:        "1F3BE",
		Value:      "🎾",
		Descriptor: "Tennis",
	},
	"1F3BF": {
		Key:        "1F3BF",
		Value:      "🎿",
		Descriptor: "Skis",
	},
	"1F3C0": {
		Key:        "1F3C0",
		Value:      "🏀",
		Descriptor: "Basketball",
	},
	"1F3C1": {
		Key:        "1F3C1",
		Value:      "🏁",
		Descriptor: "Chequered Flag",
	},
	"1F3C2": {
		Key:        "1F3C2",
		Value:      "🏂",
		Descriptor: "Snowboarder",
	},
	"1F3C3": {
		Key:        "1F3C3",
		Value:      "🏃",
		Descriptor: "Person Running",
	},
	"1F3C3-1F3FB": {
		Key:        "1F3C3-1F3FB",
		Value:      "🏃🏻",
		Descriptor: "Person Running: Light Skin Tone",
	},
	"1F3C3-1F3FB-200D-2640-FE0F": {
		Key:        "1F3C3-1F3FB-200D-2640-FE0F",
		Value:      "🏃🏻‍♀️",
		Descriptor: "Woman Running: Light Skin Tone",
	},
	"1F3C3-1F3FB-200D-2642-FE0F": {
		Key:        "1F3C3-1F3FB-200D-2642-FE0F",
		Value:      "🏃🏻‍♂️",
		Descriptor: "Man Running: Light Skin Tone",
	},
	"1F3C3-1F3FC": {
		Key:        "1F3C3-1F3FC",
		Value:      "🏃🏼",
		Descriptor: "Person Running: Medium-Light Skin Tone",
	},
	"1F3C3-1F3FC-200D-2640-FE0F": {
		Key:        "1F3C3-1F3FC-200D-2640-FE0F",
		Value:      "🏃🏼‍♀️",
		Descriptor: "Woman Running: Medium-Light Skin Tone",
	},
	"1F3C3-1F3FC-200D-2642-FE0F": {
		Key:        "1F3C3-1F3FC-200D-2642-FE0F",
		Value:      "🏃🏼‍♂️",
		Descriptor: "Man Running: Medium-Light Skin Tone",
	},
	"1F3C3-1F3FD": {
		Key:        "1F3C3-1F3FD",
		Value:      "🏃🏽",
		Descriptor: "Person Running: Medium Skin Tone",
	},
	"1F3C3-1F3FD-200D-2640-FE0F": {
		Key:        "1F3C3-1F3FD-200D-2640-FE0F",
		Value:      "🏃🏽‍♀️",
		Descriptor: "Woman Running: Medium Skin Tone",
	},
	"1F3C3-1F3FD-200D-2642-FE0F": {
		Key:        "1F3C3-1F3FD-200D-2642-FE0F",
		Value:      "🏃🏽‍♂️",
		Descriptor: "Man Running: Medium Skin Tone",
	},
	"1F3C3-1F3FE": {
		Key:        "1F3C3-1F3FE",
		Value:      "🏃🏾",
		Descriptor: "Person Running: Medium-Dark Skin Tone",
	},
	"1F3C3-1F3FE-200D-2640-FE0F": {
		Key:        "1F3C3-1F3FE-200D-2640-FE0F",
		Value:      "🏃🏾‍♀️",
		Descriptor: "Woman Running: Medium-Dark Skin Tone",
	},
	"1F3C3-1F3FE-200D-2642-FE0F": {
		Key:        "1F3C3-1F3FE-200D-2642-FE0F",
		Value:      "🏃🏾‍♂️",
		Descriptor: "Man Running: Medium-Dark Skin Tone",
	},
	"1F3C3-1F3FF": {
		Key:        "1F3C3-1F3FF",
		Value:      "🏃🏿",
		Descriptor: "Person Running: Dark Skin Tone",
	},
	"1F3C3-1F3FF-200D-2640-FE0F": {
		Key:        "1F3C3-1F3FF-200D-2640-FE0F",
		Value:      "🏃🏿‍♀️",
		Descriptor: "Woman Running: Dark Skin Tone",
	},
	"1F3C3-1F3FF-200D-2642-FE0F": {
		Key:        "1F3C3-1F3FF-200D-2642-FE0F",
		Value:      "🏃🏿‍♂️",
		Descriptor: "Man Running: Dark Skin Tone",
	},
	"1F3C3-200D-2640-FE0F": {
		Key:        "1F3C3-200D-2640-FE0F",
		Value:      "🏃‍♀️",
		Descriptor: "Woman Running",
	},
	"1F3C3-200D-2642-FE0F": {
		Key:        "1F3C3-200D-2642-FE0F",
		Value:      "🏃‍♂️",
		Descriptor: "Man Running",
	},
	"1F3C4": {
		Key:        "1F3C4",
		Value:      "🏄",
		Descriptor: "Person Surfing",
	},
	"1F3C4-1F3FB": {
		Key:        "1F3C4-1F3FB",
		Value:      "🏄🏻",
		Descriptor: "Person Surfing: Light Skin Tone",
	},
	"1F3C4-1F3FB-200D-2640-FE0F": {
		Key:        "1F3C4-1F3FB-200D-2640-FE0F",
		Value:      "🏄🏻‍♀️",
		Descriptor: "Woman Surfing: Light Skin Tone",
	},
	"1F3C4-1F3FB-200D-2642-FE0F": {
		Key:        "1F3C4-1F3FB-200D-2642-FE0F",
		Value:      "🏄🏻‍♂️",
		Descriptor: "Man Surfing: Light Skin Tone",
	},
	"1F3C4-1F3FC": {
		Key:        "1F3C4-1F3FC",
		Value:      "🏄🏼",
		Descriptor: "Person Surfing: Medium-Light Skin Tone",
	},
	"1F3C4-1F3FC-200D-2640-FE0F": {
		Key:        "1F3C4-1F3FC-200D-2640-FE0F",
		Value:      "🏄🏼‍♀️",
		Descriptor: "Woman Surfing: Medium-Light Skin Tone",
	},
	"1F3C4-1F3FC-200D-2642-FE0F": {
		Key:        "1F3C4-1F3FC-200D-2642-FE0F",
		Value:      "🏄🏼‍♂️",
		Descriptor: "Man Surfing: Medium-Light Skin Tone",
	},
	"1F3C4-1F3FD": {
		Key:        "1F3C4-1F3FD",
		Value:      "🏄🏽",
		Descriptor: "Person Surfing: Medium Skin Tone",
	},
	"1F3C4-1F3FD-200D-2640-FE0F": {
		Key:        "1F3C4-1F3FD-200D-2640-FE0F",
		Value:      "🏄🏽‍♀️",
		Descriptor: "Woman Surfing: Medium Skin Tone",
	},
	"1F3C4-1F3FD-200D-2642-FE0F": {
		Key:        "1F3C4-1F3FD-200D-2642-FE0F",
		Value:      "🏄🏽‍♂️",
		Descriptor: "Man Surfing: Medium Skin Tone",
	},
	"1F3C4-1F3FE": {
		Key:        "1F3C4-1F3FE",
		Value:      "🏄🏾",
		Descriptor: "Person Surfing: Medium-Dark Skin Tone",
	},
	"1F3C4-1F3FE-200D-2640-FE0F": {
		Key:        "1F3C4-1F3FE-200D-2640-FE0F",
		Value:      "🏄🏾‍♀️",
		Descriptor: "Woman Surfing: Medium-Dark Skin Tone",
	},
	"1F3C4-1F3FE-200D-2642-FE0F": {
		Key:        "1F3C4-1F3FE-200D-2642-FE0F",
		Value:      "🏄🏾‍♂️",
		Descriptor: "Man Surfing: Medium-Dark Skin Tone",
	},
	"1F3C4-1F3FF": {
		Key:        "1F3C4-1F3FF",
		Value:      "🏄🏿",
		Descriptor: "Person Surfing: Dark Skin Tone",
	},
	"1F3C4-1F3FF-200D-2640-FE0F": {
		Key:        "1F3C4-1F3FF-200D-2640-FE0F",
		Value:      "🏄🏿‍♀️",
		Descriptor: "Woman Surfing: Dark Skin Tone",
	},
	"1F3C4-1F3FF-200D-2642-FE0F": {
		Key:        "1F3C4-1F3FF-200D-2642-FE0F",
		Value:      "🏄🏿‍♂️",
		Descriptor: "Man Surfing: Dark Skin Tone",
	},
	"1F3C4-200D-2640-FE0F": {
		Key:        "1F3C4-200D-2640-FE0F",
		Value:      "🏄‍♀️",
		Descriptor: "Woman Surfing",
	},
	"1F3C4-200D-2642-FE0F": {
		Key:        "1F3C4-200D-2642-FE0F",
		Value:      "🏄‍♂️",
		Descriptor: "Man Surfing",
	},
	"1F3C5": {
		Key:        "1F3C5",
		Value:      "🏅",
		Descriptor: "Sports Medal",
	},
	"1F3C6": {
		Key:        "1F3C6",
		Value:      "🏆",
		Descriptor: "Trophy",
	},
	"1F3C7": {
		Key:        "1F3C7",
		Value:      "🏇",
		Descriptor: "Horse Racing",
	},
	"1F3C7-1F3FB": {
		Key:        "1F3C7-1F3FB",
		Value:      "🏇🏻",
		Descriptor: "Horse Racing: Light Skin Tone",
	},
	"1F3C7-1F3FC": {
		Key:        "1F3C7-1F3FC",
		Value:      "🏇🏼",
		Descriptor: "Horse Racing: Medium-Light Skin Tone",
	},
	"1F3C7-1F3FD": {
		Key:        "1F3C7-1F3FD",
		Value:      "🏇🏽",
		Descriptor: "Horse Racing: Medium Skin Tone",
	},
	"1F3C7-1F3FE": {
		Key:        "1F3C7-1F3FE",
		Value:      "🏇🏾",
		Descriptor: "Horse Racing: Medium-Dark Skin Tone",
	},
	"1F3C7-1F3FF": {
		Key:        "1F3C7-1F3FF",
		Value:      "🏇🏿",
		Descriptor: "Horse Racing: Dark Skin Tone",
	},
	"1F3C8": {
		Key:        "1F3C8",
		Value:      "🏈",
		Descriptor: "American Football",
	},
	"1F3C9": {
		Key:        "1F3C9",
		Value:      "🏉",
		Descriptor: "Rugby Football",
	},
	"1F3CA": {
		Key:        "1F3CA",
		Value:      "🏊",
		Descriptor: "Person Swimming",
	},
	"1F3CA-1F3FB": {
		Key:        "1F3CA-1F3FB",
		Value:      "🏊🏻",
		Descriptor: "Person Swimming: Light Skin Tone",
	},
	"1F3CA-1F3FB-200D-2640-FE0F": {
		Key:        "1F3CA-1F3FB-200D-2640-FE0F",
		Value:      "🏊🏻‍♀️",
		Descriptor: "Woman Swimming: Light Skin Tone",
	},
	"1F3CA-1F3FB-200D-2642-FE0F": {
		Key:        "1F3CA-1F3FB-200D-2642-FE0F",
		Value:      "🏊🏻‍♂️",
		Descriptor: "Man Swimming: Light Skin Tone",
	},
	"1F3CA-1F3FC": {
		Key:        "1F3CA-1F3FC",
		Value:      "🏊🏼",
		Descriptor: "Person Swimming: Medium-Light Skin Tone",
	},
	"1F3CA-1F3FC-200D-2640-FE0F": {
		Key:        "1F3CA-1F3FC-200D-2640-FE0F",
		Value:      "🏊🏼‍♀️",
		Descriptor: "Woman Swimming: Medium-Light Skin Tone",
	},
	"1F3CA-1F3FC-200D-2642-FE0F": {
		Key:        "1F3CA-1F3FC-200D-2642-FE0F",
		Value:      "🏊🏼‍♂️",
		Descriptor: "Man Swimming: Medium-Light Skin Tone",
	},
	"1F3CA-1F3FD": {
		Key:        "1F3CA-1F3FD",
		Value:      "🏊🏽",
		Descriptor: "Person Swimming: Medium Skin Tone",
	},
	"1F3CA-1F3FD-200D-2640-FE0F": {
		Key:        "1F3CA-1F3FD-200D-2640-FE0F",
		Value:      "🏊🏽‍♀️",
		Descriptor: "Woman Swimming: Medium Skin Tone",
	},
	"1F3CA-1F3FD-200D-2642-FE0F": {
		Key:        "1F3CA-1F3FD-200D-2642-FE0F",
		Value:      "🏊🏽‍♂️",
		Descriptor: "Man Swimming: Medium Skin Tone",
	},
	"1F3CA-1F3FE": {
		Key:        "1F3CA-1F3FE",
		Value:      "🏊🏾",
		Descriptor: "Person Swimming: Medium-Dark Skin Tone",
	},
	"1F3CA-1F3FE-200D-2640-FE0F": {
		Key:        "1F3CA-1F3FE-200D-2640-FE0F",
		Value:      "🏊🏾‍♀️",
		Descriptor: "Woman Swimming: Medium-Dark Skin Tone",
	},
	"1F3CA-1F3FE-200D-2642-FE0F": {
		Key:        "1F3CA-1F3FE-200D-2642-FE0F",
		Value:      "🏊🏾‍♂️",
		Descriptor: "Man Swimming: Medium-Dark Skin Tone",
	},
	"1F3CA-1F3FF": {
		Key:        "1F3CA-1F3FF",
		Value:      "🏊🏿",
		Descriptor: "Person Swimming: Dark Skin Tone",
	},
	"1F3CA-1F3FF-200D-2640-FE0F": {
		Key:        "1F3CA-1F3FF-200D-2640-FE0F",
		Value:      "🏊🏿‍♀️",
		Descriptor: "Woman Swimming: Dark Skin Tone",
	},
	"1F3CA-1F3FF-200D-2642-FE0F": {
		Key:        "1F3CA-1F3FF-200D-2642-FE0F",
		Value:      "🏊🏿‍♂️",
		Descriptor: "Man Swimming: Dark Skin Tone",
	},
	"1F3CA-200D-2640-FE0F": {
		Key:        "1F3CA-200D-2640-FE0F",
		Value:      "🏊‍♀️",
		Descriptor: "Woman Swimming",
	},
	"1F3CA-200D-2642-FE0F": {
		Key:        "1F3CA-200D-2642-FE0F",
		Value:      "🏊‍♂️",
		Descriptor: "Man Swimming",
	},
	"1F3CB-1F3FB": {
		Key:        "1F3CB-1F3FB",
		Value:      "🏋🏻",
		Descriptor: "Person Lifting Weights: Light Skin Tone",
	},
	"1F3CB-1F3FB-200D-2640-FE0F": {
		Key:        "1F3CB-1F3FB-200D-2640-FE0F",
		Value:      "🏋🏻‍♀️",
		Descriptor: "Woman Lifting Weights: Light Skin Tone",
	},
	"1F3CB-1F3FB-200D-2642-FE0F": {
		Key:        "1F3CB-1F3FB-200D-2642-FE0F",
		Value:      "🏋🏻‍♂️",
		Descriptor: "Man Lifting Weights: Light Skin Tone",
	},
	"1F3CB-1F3FC": {
		Key:        "1F3CB-1F3FC",
		Value:      "🏋🏼",
		Descriptor: "Person Lifting Weights: Medium-Light Skin Tone",
	},
	"1F3CB-1F3FC-200D-2640-FE0F": {
		Key:        "1F3CB-1F3FC-200D-2640-FE0F",
		Value:      "🏋🏼‍♀️",
		Descriptor: "Woman Lifting Weights: Medium-Light Skin Tone",
	},
	"1F3CB-1F3FC-200D-2642-FE0F": {
		Key:        "1F3CB-1F3FC-200D-2642-FE0F",
		Value:      "🏋🏼‍♂️",
		Descriptor: "Man Lifting Weights: Medium-Light Skin Tone",
	},
	"1F3CB-1F3FD": {
		Key:        "1F3CB-1F3FD",
		Value:      "🏋🏽",
		Descriptor: "Person Lifting Weights: Medium Skin Tone",
	},
	"1F3CB-1F3FD-200D-2640-FE0F": {
		Key:        "1F3CB-1F3FD-200D-2640-FE0F",
		Value:      "🏋🏽‍♀️",
		Descriptor: "Woman Lifting Weights: Medium Skin Tone",
	},
	"1F3CB-1F3FD-200D-2642-FE0F": {
		Key:        "1F3CB-1F3FD-200D-2642-FE0F",
		Value:      "🏋🏽‍♂️",
		Descriptor: "Man Lifting Weights: Medium Skin Tone",
	},
	"1F3CB-1F3FE": {
		Key:        "1F3CB-1F3FE",
		Value:      "🏋🏾",
		Descriptor: "Person Lifting Weights: Medium-Dark Skin Tone",
	},
	"1F3CB-1F3FE-200D-2640-FE0F": {
		Key:        "1F3CB-1F3FE-200D-2640-FE0F",
		Value:      "🏋🏾‍♀️",
		Descriptor: "Woman Lifting Weights: Medium-Dark Skin Tone",
	},
	"1F3CB-1F3FE-200D-2642-FE0F": {
		Key:        "1F3CB-1F3FE-200D-2642-FE0F",
		Value:      "🏋🏾‍♂️",
		Descriptor: "Man Lifting Weights: Medium-Dark Skin Tone",
	},
	"1F3CB-1F3FF": {
		Key:        "1F3CB-1F3FF",
		Value:      "🏋🏿",
		Descriptor: "Person Lifting Weights: Dark Skin Tone",
	},
	"1F3CB-1F3FF-200D-2640-FE0F": {
		Key:        "1F3CB-1F3FF-200D-2640-FE0F",
		Value:      "🏋🏿‍♀️",
		Descriptor: "Woman Lifting Weights: Dark Skin Tone",
	},
	"1F3CB-1F3FF-200D-2642-FE0F": {
		Key:        "1F3CB-1F3FF-200D-2642-FE0F",
		Value:      "🏋🏿‍♂️",
		Descriptor: "Man Lifting Weights: Dark Skin Tone",
	},
	"1F3CB-FE0F": {
		Key:        "1F3CB-FE0F",
		Value:      "🏋️",
		Descriptor: "Person Lifting Weights",
	},
	"1F3CB-FE0F-200D-2640-FE0F": {
		Key:        "1F3CB-FE0F-200D-2640-FE0F",
		Value:      "🏋️‍♀️",
		Descriptor: "Woman Lifting Weights",
	},
	"1F3CB-FE0F-200D-2642-FE0F": {
		Key:        "1F3CB-FE0F-200D-2642-FE0F",
		Value:      "🏋️‍♂️",
		Descriptor: "Man Lifting Weights",
	},
	"1F3CC-1F3FB": {
		Key:        "1F3CC-1F3FB",
		Value:      "🏌🏻",
		Descriptor: "Person Golfing: Light Skin Tone",
	},
	"1F3CC-1F3FB-200D-2640-FE0F": {
		Key:        "1F3CC-1F3FB-200D-2640-FE0F",
		Value:      "🏌🏻‍♀️",
		Descriptor: "Woman Golfing: Light Skin Tone",
	},
	"1F3CC-1F3FB-200D-2642-FE0F": {
		Key:        "1F3CC-1F3FB-200D-2642-FE0F",
		Value:      "🏌🏻‍♂️",
		Descriptor: "Man Golfing: Light Skin Tone",
	},
	"1F3CC-1F3FC": {
		Key:        "1F3CC-1F3FC",
		Value:      "🏌🏼",
		Descriptor: "Person Golfing: Medium-Light Skin Tone",
	},
	"1F3CC-1F3FC-200D-2640-FE0F": {
		Key:        "1F3CC-1F3FC-200D-2640-FE0F",
		Value:      "🏌🏼‍♀️",
		Descriptor: "Woman Golfing: Medium-Light Skin Tone",
	},
	"1F3CC-1F3FC-200D-2642-FE0F": {
		Key:        "1F3CC-1F3FC-200D-2642-FE0F",
		Value:      "🏌🏼‍♂️",
		Descriptor: "Man Golfing: Medium-Light Skin Tone",
	},
	"1F3CC-1F3FD": {
		Key:        "1F3CC-1F3FD",
		Value:      "🏌🏽",
		Descriptor: "Person Golfing: Medium Skin Tone",
	},
	"1F3CC-1F3FD-200D-2640-FE0F": {
		Key:        "1F3CC-1F3FD-200D-2640-FE0F",
		Value:      "🏌🏽‍♀️",
		Descriptor: "Woman Golfing: Medium Skin Tone",
	},
	"1F3CC-1F3FD-200D-2642-FE0F": {
		Key:        "1F3CC-1F3FD-200D-2642-FE0F",
		Value:      "🏌🏽‍♂️",
		Descriptor: "Man Golfing: Medium Skin Tone",
	},
	"1F3CC-1F3FE": {
		Key:        "1F3CC-1F3FE",
		Value:      "🏌🏾",
		Descriptor: "Person Golfing: Medium-Dark Skin Tone",
	},
	"1F3CC-1F3FE-200D-2640-FE0F": {
		Key:        "1F3CC-1F3FE-200D-2640-FE0F",
		Value:      "🏌🏾‍♀️",
		Descriptor: "Woman Golfing: Medium-Dark Skin Tone",
	},
	"1F3CC-1F3FE-200D-2642-FE0F": {
		Key:        "1F3CC-1F3FE-200D-2642-FE0F",
		Value:      "🏌🏾‍♂️",
		Descriptor: "Man Golfing: Medium-Dark Skin Tone",
	},
	"1F3CC-1F3FF": {
		Key:        "1F3CC-1F3FF",
		Value:      "🏌🏿",
		Descriptor: "Person Golfing: Dark Skin Tone",
	},
	"1F3CC-1F3FF-200D-2640-FE0F": {
		Key:        "1F3CC-1F3FF-200D-2640-FE0F",
		Value:      "🏌🏿‍♀️",
		Descriptor: "Woman Golfing: Dark Skin Tone",
	},
	"1F3CC-1F3FF-200D-2642-FE0F": {
		Key:        "1F3CC-1F3FF-200D-2642-FE0F",
		Value:      "🏌🏿‍♂️",
		Descriptor: "Man Golfing: Dark Skin Tone",
	},
	"1F3CC-FE0F": {
		Key:        "1F3CC-FE0F",
		Value:      "🏌️",
		Descriptor: "Person Golfing",
	},
	"1F3CC-FE0F-200D-2640-FE0F": {
		Key:        "1F3CC-FE0F-200D-2640-FE0F",
		Value:      "🏌️‍♀️",
		Descriptor: "Woman Golfing",
	},
	"1F3CC-FE0F-200D-2642-FE0F": {
		Key:        "1F3CC-FE0F-200D-2642-FE0F",
		Value:      "🏌️‍♂️",
		Descriptor: "Man Golfing",
	},
	"1F3CD-FE0F": {
		Key:        "1F3CD-FE0F",
		Value:      "🏍️",
		Descriptor: "Motorcycle",
	},
	"1F3CE-FE0F": {
		Key:        "1F3CE-FE0F",
		Value:      "🏎️",
		Descriptor: "Racing Car",
	},
	"1F3CF": {
		Key:        "1F3CF",
		Value:      "🏏",
		Descriptor: "Cricket Game",
	},
	"1F3D0": {
		Key:        "1F3D0",
		Value:      "🏐",
		Descriptor: "Volleyball",
	},
	"1F3D1": {
		Key:        "1F3D1",
		Value:      "🏑",
		Descriptor: "Field Hockey",
	},
	"1F3D2": {
		Key:        "1F3D2",
		Value:      "🏒",
		Descriptor: "Ice Hockey",
	},
	"1F3D3": {
		Key:        "1F3D3",
		Value:      "🏓",
		Descriptor: "Ping Pong",
	},
	"1F3D4-FE0F": {
		Key:        "1F3D4-FE0F",
		Value:      "🏔️",
		Descriptor: "Snow-Capped Mountain",
	},
	"1F3D5-FE0F": {
		Key:        "1F3D5-FE0F",
		Value:      "🏕️",
		Descriptor: "Camping",
	},
	"1F3D6-FE0F": {
		Key:        "1F3D6-FE0F",
		Value:      "🏖️",
		Descriptor: "Beach With Umbrella",
	},
	"1F3D7-FE0F": {
		Key:        "1F3D7-FE0F",
		Value:      "🏗️",
		Descriptor: "Building Construction",
	},
	"1F3D8-FE0F": {
		Key:        "1F3D8-FE0F",
		Value:      "🏘️",
		Descriptor: "Houses",
	},
	"1F3D9-FE0F": {
		Key:        "1F3D9-FE0F",
		Value:      "🏙️",
		Descriptor: "Cityscape",
	},
	"1F3DA-FE0F": {
		Key:        "1F3DA-FE0F",
		Value:      "🏚️",
		Descriptor: "Derelict House",
	},
	"1F3DB-FE0F": {
		Key:        "1F3DB-FE0F",
		Value:      "🏛️",
		Descriptor: "Classical Building",
	},
	"1F3DC-FE0F": {
		Key:        "1F3DC-FE0F",
		Value:      "🏜️",
		Descriptor: "Desert",
	},
	"1F3DD-FE0F": {
		Key:        "1F3DD-FE0F",
		Value:      "🏝️",
		Descriptor: "Desert Island",
	},
	"1F3DE-FE0F": {
		Key:        "1F3DE-FE0F",
		Value:      "🏞️",
		Descriptor: "National Park",
	},
	"1F3DF-FE0F": {
		Key:        "1F3DF-FE0F",
		Value:      "🏟️",
		Descriptor: "Stadium",
	},
	"1F3E0": {
		Key:        "1F3E0",
		Value:      "🏠",
		Descriptor: "House",
	},
	"1F3E1": {
		Key:        "1F3E1",
		Value:      "🏡",
		Descriptor: "House With Garden",
	},
	"1F3E2": {
		Key:        "1F3E2",
		Value:      "🏢",
		Descriptor: "Office Building",
	},
	"1F3E3": {
		Key:        "1F3E3",
		Value:      "🏣",
		Descriptor: "Japanese Post Office",
	},
	"1F3E4": {
		Key:        "1F3E4",
		Value:      "🏤",
		Descriptor: "Post Office",
	},
	"1F3E5": {
		Key:        "1F3E5",
		Value:      "🏥",
		Descriptor: "Hospital",
	},
	"1F3E6": {
		Key:        "1F3E6",
		Value:      "🏦",
		Descriptor: "Bank",
	},
	"1F3E7": {
		Key:        "1F3E7",
		Value:      "🏧",
		Descriptor: "ATM Sign",
	},
	"1F3E8": {
		Key:        "1F3E8",
		Value:      "🏨",
		Descriptor: "Hotel",
	},
	"1F3E9": {
		Key:        "1F3E9",
		Value:      "🏩",
		Descriptor: "Love Hotel",
	},
	"1F3EA": {
		Key:        "1F3EA",
		Value:      "🏪",
		Descriptor: "Convenience Store",
	},
	"1F3EB": {
		Key:        "1F3EB",
		Value:      "🏫",
		Descriptor: "School",
	},
	"1F3EC": {
		Key:        "1F3EC",
		Value:      "🏬",
		Descriptor: "Department Store",
	},
	"1F3ED": {
		Key:        "1F3ED",
		Value:      "🏭",
		Descriptor: "Factory",
	},
	"1F3EE": {
		Key:        "1F3EE",
		Value:      "🏮",
		Descriptor: "Red Paper Lantern",
	},
	"1F3EF": {
		Key:        "1F3EF",
		Value:      "🏯",
		Descriptor: "Japanese Castle",
	},
	"1F3F0": {
		Key:        "1F3F0",
		Value:      "🏰",
		Descriptor: "Castle",
	},
	"1F3F3-FE0F": {
		Key:        "1F3F3-FE0F",
		Value:      "🏳️",
		Descriptor: "White Flag",
	},
	"1F3F3-FE0F-200D-1F308": {
		Key:        "1F3F3-FE0F-200D-1F308",
		Value:      "🏳️‍🌈",
		Descriptor: "Rainbow Flag",
	},
	"1F3F4": {
		Key:        "1F3F4",
		Value:      "🏴",
		Descriptor: "Black Flag",
	},
	"1F3F4-200D-2620-FE0F": {
		Key:        "1F3F4-200D-2620-FE0F",
		Value:      "🏴‍☠️",
		Descriptor: "Pirate Flag",
	},
	"1F3F4-E0067-E0062-E0065-E006E-E0067-E007F": {
		Key:        "1F3F4-E0067-E0062-E0065-E006E-E0067-E007F",
		Value:      "🏴󠁧󠁢󠁥󠁮󠁧󠁿",
		Descriptor: "Flag: England",
	},
	"1F3F4-E0067-E0062-E0073-E0063-E0074-E007F": {
		Key:        "1F3F4-E0067-E0062-E0073-E0063-E0074-E007F",
		Value:      "🏴󠁧󠁢󠁳󠁣󠁴󠁿",
		Descriptor: "Flag: Scotland",
	},
	"1F3F4-E0067-E0062-E0077-E006C-E0073-E007F": {
		Key:        "1F3F4-E0067-E0062-E0077-E006C-E0073-E007F",
		Value:      "🏴󠁧󠁢󠁷󠁬󠁳󠁿",
		Descriptor: "Flag: Wales",
	},
	"1F3F5-FE0F": {
		Key:        "1F3F5-FE0F",
		Value:      "🏵️",
		Descriptor: "Rosette",
	},
	"1F3F7-FE0F": {
		Key:        "1F3F7-FE0F",
		Value:      "🏷️",
		Descriptor: "Label",
	},
	"1F3F8": {
		Key:        "1F3F8",
		Value:      "🏸",
		Descriptor: "Badminton",
	},
	"1F3F9": {
		Key:        "1F3F9",
		Value:      "🏹",
		Descriptor: "Bow and Arrow",
	},
	"1F3FA": {
		Key:        "1F3FA",
		Value:      "🏺",
		Descriptor: "Amphora",
	},
	"1F3FB": {
		Key:        "1F3FB",
		Value:      "🏻",
		Descriptor: "Light Skin Tone",
	},
	"1F3FC": {
		Key:        "1F3FC",
		Value:      "🏼",
		Descriptor: "Medium-Light Skin Tone",
	},
	"1F3FD": {
		Key:        "1F3FD",
		Value:      "🏽",
		Descriptor: "Medium Skin Tone",
	},
	"1F3FE": {
		Key:        "1F3FE",
		Value:      "🏾",
		Descriptor: "Medium-Dark Skin Tone",
	},
	"1F3FF": {
		Key:        "1F3FF",
		Value:      "🏿",
		Descriptor: "Dark Skin Tone",
	},
	"1F400": {
		Key:        "1F400",
		Value:      "🐀",
		Descriptor: "Rat",
	},
	"1F401": {
		Key:        "1F401",
		Value:      "🐁",
		Descriptor: "Mouse",
	},
	"1F402": {
		Key:        "1F402",
		Value:      "🐂",
		Descriptor: "Ox",
	},
	"1F403": {
		Key:        "1F403",
		Value:      "🐃",
		Descriptor: "Water Buffalo",
	},
	"1F404": {
		Key:        "1F404",
		Value:      "🐄",
		Descriptor: "Cow",
	},
	"1F405": {
		Key:        "1F405",
		Value:      "🐅",
		Descriptor: "Tiger",
	},
	"1F406": {
		Key:        "1F406",
		Value:      "🐆",
		Descriptor: "Leopard",
	},
	"1F407": {
		Key:        "1F407",
		Value:      "🐇",
		Descriptor: "Rabbit",
	},
	"1F408": {
		Key:        "1F408",
		Value:      "🐈",
		Descriptor: "Cat",
	},
	"1F409": {
		Key:        "1F409",
		Value:      "🐉",
		Descriptor: "Dragon",
	},
	"1F40A": {
		Key:        "1F40A",
		Value:      "🐊",
		Descriptor: "Crocodile",
	},
	"1F40B": {
		Key:        "1F40B",
		Value:      "🐋",
		Descriptor: "Whale",
	},
	"1F40C": {
		Key:        "1F40C",
		Value:      "🐌",
		Descriptor: "Snail",
	},
	"1F40D": {
		Key:        "1F40D",
		Value:      "🐍",
		Descriptor: "Snake",
	},
	"1F40E": {
		Key:        "1F40E",
		Value:      "🐎",
		Descriptor: "Horse",
	},
	"1F40F": {
		Key:        "1F40F",
		Value:      "🐏",
		Descriptor: "Ram",
	},
	"1F410": {
		Key:        "1F410",
		Value:      "🐐",
		Descriptor: "Goat",
	},
	"1F411": {
		Key:        "1F411",
		Value:      "🐑",
		Descriptor: "Ewe",
	},
	"1F412": {
		Key:        "1F412",
		Value:      "🐒",
		Descriptor: "Monkey",
	},
	"1F413": {
		Key:        "1F413",
		Value:      "🐓",
		Descriptor: "Rooster",
	},
	"1F414": {
		Key:        "1F414",
		Value:      "🐔",
		Descriptor: "Chicken",
	},
	"1F415": {
		Key:        "1F415",
		Value:      "🐕",
		Descriptor: "Dog",
	},
	"1F415-200D-1F9BA": {
		Key:        "1F415-200D-1F9BA",
		Value:      "🐕‍🦺",
		Descriptor: "Service Dog",
	},
	"1F416": {
		Key:        "1F416",
		Value:      "🐖",
		Descriptor: "Pig",
	},
	"1F417": {
		Key:        "1F417",
		Value:      "🐗",
		Descriptor: "Boar",
	},
	"1F418": {
		Key:        "1F418",
		Value:      "🐘",
		Descriptor: "Elephant",
	},
	"1F419": {
		Key:        "1F419",
		Value:      "🐙",
		Descriptor: "Octopus",
	},
	"1F41A": {
		Key:        "1F41A",
		Value:      "🐚",
		Descriptor: "Spiral Shell",
	},
	"1F41B": {
		Key:        "1F41B",
		Value:      "🐛",
		Descriptor: "Bug",
	},
	"1F41C": {
		Key:        "1F41C",
		Value:      "🐜",
		Descriptor: "Ant",
	},
	"1F41D": {
		Key:        "1F41D",
		Value:      "🐝",
		Descriptor: "Honeybee",
	},
	"1F41E": {
		Key:        "1F41E",
		Value:      "🐞",
		Descriptor: "Lady Beetle",
	},
	"1F41F": {
		Key:        "1F41F",
		Value:      "🐟",
		Descriptor: "Fish",
	},
	"1F420": {
		Key:        "1F420",
		Value:      "🐠",
		Descriptor: "Tropical Fish",
	},
	"1F421": {
		Key:        "1F421",
		Value:      "🐡",
		Descriptor: "Blowfish",
	},
	"1F422": {
		Key:        "1F422",
		Value:      "🐢",
		Descriptor: "Turtle",
	},
	"1F423": {
		Key:        "1F423",
		Value:      "🐣",
		Descriptor: "Hatching Chick",
	},
	"1F424": {
		Key:        "1F424",
		Value:      "🐤",
		Descriptor: "Baby Chick",
	},
	"1F425": {
		Key:        "1F425",
		Value:      "🐥",
		Descriptor: "Front-Facing Baby Chick",
	},
	"1F426": {
		Key:        "1F426",
		Value:      "🐦",
		Descriptor: "Bird",
	},
	"1F427": {
		Key:        "1F427",
		Value:      "🐧",
		Descriptor: "Penguin",
	},
	"1F428": {
		Key:        "1F428",
		Value:      "🐨",
		Descriptor: "Koala",
	},
	"1F429": {
		Key:        "1F429",
		Value:      "🐩",
		Descriptor: "Poodle",
	},
	"1F42A": {
		Key:        "1F42A",
		Value:      "🐪",
		Descriptor: "Camel",
	},
	"1F42B": {
		Key:        "1F42B",
		Value:      "🐫",
		Descriptor: "Two-Hump Camel",
	},
	"1F42C": {
		Key:        "1F42C",
		Value:      "🐬",
		Descriptor: "Dolphin",
	},
	"1F42D": {
		Key:        "1F42D",
		Value:      "🐭",
		Descriptor: "Mouse Face",
	},
	"1F42E": {
		Key:        "1F42E",
		Value:      "🐮",
		Descriptor: "Cow Face",
	},
	"1F42F": {
		Key:        "1F42F",
		Value:      "🐯",
		Descriptor: "Tiger Face",
	},
	"1F430": {
		Key:        "1F430",
		Value:      "🐰",
		Descriptor: "Rabbit Face",
	},
	"1F431": {
		Key:        "1F431",
		Value:      "🐱",
		Descriptor: "Cat Face",
	},
	"1F432": {
		Key:        "1F432",
		Value:      "🐲",
		Descriptor: "Dragon Face",
	},
	"1F433": {
		Key:        "1F433",
		Value:      "🐳",
		Descriptor: "Spouting Whale",
	},
	"1F434": {
		Key:        "1F434",
		Value:      "🐴",
		Descriptor: "Horse Face",
	},
	"1F435": {
		Key:        "1F435",
		Value:      "🐵",
		Descriptor: "Monkey Face",
	},
	"1F436": {
		Key:        "1F436",
		Value:      "🐶",
		Descriptor: "Dog Face",
	},
	"1F437": {
		Key:        "1F437",
		Value:      "🐷",
		Descriptor: "Pig Face",
	},
	"1F438": {
		Key:        "1F438",
		Value:      "🐸",
		Descriptor: "Frog",
	},
	"1F439": {
		Key:        "1F439",
		Value:      "🐹",
		Descriptor: "Hamster",
	},
	"1F43A": {
		Key:        "1F43A",
		Value:      "🐺",
		Descriptor: "Wolf",
	},
	"1F43B": {
		Key:        "1F43B",
		Value:      "🐻",
		Descriptor: "Bear",
	},
	"1F43C": {
		Key:        "1F43C",
		Value:      "🐼",
		Descriptor: "Panda",
	},
	"1F43D": {
		Key:        "1F43D",
		Value:      "🐽",
		Descriptor: "Pig Nose",
	},
	"1F43E": {
		Key:        "1F43E",
		Value:      "🐾",
		Descriptor: "Paw Prints",
	},
	"1F43F-FE0F": {
		Key:        "1F43F-FE0F",
		Value:      "🐿️",
		Descriptor: "Chipmunk",
	},
	"1F440": {
		Key:        "1F440",
		Value:      "👀",
		Descriptor: "Eyes",
	},
	"1F441-FE0F": {
		Key:        "1F441-FE0F",
		Value:      "👁️",
		Descriptor: "Eye",
	},
	"1F441-FE0F-200D-1F5E8-FE0F": {
		Key:        "1F441-FE0F-200D-1F5E8-FE0F",
		Value:      "👁️‍🗨️",
		Descriptor: "Eye in Speech Bubble",
	},
	"1F442": {
		Key:        "1F442",
		Value:      "👂",
		Descriptor: "Ear",
	},
	"1F442-1F3FB": {
		Key:        "1F442-1F3FB",
		Value:      "👂🏻",
		Descriptor: "Ear: Light Skin Tone",
	},
	"1F442-1F3FC": {
		Key:        "1F442-1F3FC",
		Value:      "👂🏼",
		Descriptor: "Ear: Medium-Light Skin Tone",
	},
	"1F442-1F3FD": {
		Key:        "1F442-1F3FD",
		Value:      "👂🏽",
		Descriptor: "Ear: Medium Skin Tone",
	},
	"1F442-1F3FE": {
		Key:        "1F442-1F3FE",
		Value:      "👂🏾",
		Descriptor: "Ear: Medium-Dark Skin Tone",
	},
	"1F442-1F3FF": {
		Key:        "1F442-1F3FF",
		Value:      "👂🏿",
		Descriptor: "Ear: Dark Skin Tone",
	},
	"1F443": {
		Key:        "1F443",
		Value:      "👃",
		Descriptor: "Nose",
	},
	"1F443-1F3FB": {
		Key:        "1F443-1F3FB",
		Value:      "👃🏻",
		Descriptor: "Nose: Light Skin Tone",
	},
	"1F443-1F3FC": {
		Key:        "1F443-1F3FC",
		Value:      "👃🏼",
		Descriptor: "Nose: Medium-Light Skin Tone",
	},
	"1F443-1F3FD": {
		Key:        "1F443-1F3FD",
		Value:      "👃🏽",
		Descriptor: "Nose: Medium Skin Tone",
	},
	"1F443-1F3FE": {
		Key:        "1F443-1F3FE",
		Value:      "👃🏾",
		Descriptor: "Nose: Medium-Dark Skin Tone",
	},
	"1F443-1F3FF": {
		Key:        "1F443-1F3FF",
		Value:      "👃🏿",
		Descriptor: "Nose: Dark Skin Tone",
	},
	"1F444": {
		Key:        "1F444",
		Value:      "👄",
		Descriptor: "Mouth",
	},
	"1F445": {
		Key:        "1F445",
		Value:      "👅",
		Descriptor: "Tongue",
	},
	"1F446": {
		Key:        "1F446",
		Value:      "👆",
		Descriptor: "Backhand Index Pointing Up",
	},
	"1F446-1F3FB": {
		Key:        "1F446-1F3FB",
		Value:      "👆🏻",
		Descriptor: "Backhand Index Pointing Up: Light Skin Tone",
	},
	"1F446-1F3FC": {
		Key:        "1F446-1F3FC",
		Value:      "👆🏼",
		Descriptor: "Backhand Index Pointing Up: Medium-Light Skin Tone",
	},
	"1F446-1F3FD": {
		Key:        "1F446-1F3FD",
		Value:      "👆🏽",
		Descriptor: "Backhand Index Pointing Up: Medium Skin Tone",
	},
	"1F446-1F3FE": {
		Key:        "1F446-1F3FE",
		Value:      "👆🏾",
		Descriptor: "Backhand Index Pointing Up: Medium-Dark Skin Tone",
	},
	"1F446-1F3FF": {
		Key:        "1F446-1F3FF",
		Value:      "👆🏿",
		Descriptor: "Backhand Index Pointing Up: Dark Skin Tone",
	},
	"1F447": {
		Key:        "1F447",
		Value:      "👇",
		Descriptor: "Backhand Index Pointing Down",
	},
	"1F447-1F3FB": {
		Key:        "1F447-1F3FB",
		Value:      "👇🏻",
		Descriptor: "Backhand Index Pointing Down: Light Skin Tone",
	},
	"1F447-1F3FC": {
		Key:        "1F447-1F3FC",
		Value:      "👇🏼",
		Descriptor: "Backhand Index Pointing Down: Medium-Light Skin Tone",
	},
	"1F447-1F3FD": {
		Key:        "1F447-1F3FD",
		Value:      "👇🏽",
		Descriptor: "Backhand Index Pointing Down: Medium Skin Tone",
	},
	"1F447-1F3FE": {
		Key:        "1F447-1F3FE",
		Value:      "👇🏾",
		Descriptor: "Backhand Index Pointing Down: Medium-Dark Skin Tone",
	},
	"1F447-1F3FF": {
		Key:        "1F447-1F3FF",
		Value:      "👇🏿",
		Descriptor: "Backhand Index Pointing Down: Dark Skin Tone",
	},
	"1F448": {
		Key:        "1F448",
		Value:      "👈",
		Descriptor: "Backhand Index Pointing Left",
	},
	"1F448-1F3FB": {
		Key:        "1F448-1F3FB",
		Value:      "👈🏻",
		Descriptor: "Backhand Index Pointing Left: Light Skin Tone",
	},
	"1F448-1F3FC": {
		Key:        "1F448-1F3FC",
		Value:      "👈🏼",
		Descriptor: "Backhand Index Pointing Left: Medium-Light Skin Tone",
	},
	"1F448-1F3FD": {
		Key:        "1F448-1F3FD",
		Value:      "👈🏽",
		Descriptor: "Backhand Index Pointing Left: Medium Skin Tone",
	},
	"1F448-1F3FE": {
		Key:        "1F448-1F3FE",
		Value:      "👈🏾",
		Descriptor: "Backhand Index Pointing Left: Medium-Dark Skin Tone",
	},
	"1F448-1F3FF": {
		Key:        "1F448-1F3FF",
		Value:      "👈🏿",
		Descriptor: "Backhand Index Pointing Left: Dark Skin Tone",
	},
	"1F449": {
		Key:        "1F449",
		Value:      "👉",
		Descriptor: "Backhand Index Pointing Right",
	},
	"1F449-1F3FB": {
		Key:        "1F449-1F3FB",
		Value:      "👉🏻",
		Descriptor: "Backhand Index Pointing Right: Light Skin Tone",
	},
	"1F449-1F3FC": {
		Key:        "1F449-1F3FC",
		Value:      "👉🏼",
		Descriptor: "Backhand Index Pointing Right: Medium-Light Skin Tone",
	},
	"1F449-1F3FD": {
		Key:        "1F449-1F3FD",
		Value:      "👉🏽",
		Descriptor: "Backhand Index Pointing Right: Medium Skin Tone",
	},
	"1F449-1F3FE": {
		Key:        "1F449-1F3FE",
		Value:      "👉🏾",
		Descriptor: "Backhand Index Pointing Right: Medium-Dark Skin Tone",
	},
	"1F449-1F3FF": {
		Key:        "1F449-1F3FF",
		Value:      "👉🏿",
		Descriptor: "Backhand Index Pointing Right: Dark Skin Tone",
	},
	"1F44A": {
		Key:        "1F44A",
		Value:      "👊",
		Descriptor: "Oncoming Fist",
	},
	"1F44A-1F3FB": {
		Key:        "1F44A-1F3FB",
		Value:      "👊🏻",
		Descriptor: "Oncoming Fist: Light Skin Tone",
	},
	"1F44A-1F3FC": {
		Key:        "1F44A-1F3FC",
		Value:      "👊🏼",
		Descriptor: "Oncoming Fist: Medium-Light Skin Tone",
	},
	"1F44A-1F3FD": {
		Key:        "1F44A-1F3FD",
		Value:      "👊🏽",
		Descriptor: "Oncoming Fist: Medium Skin Tone",
	},
	"1F44A-1F3FE": {
		Key:        "1F44A-1F3FE",
		Value:      "👊🏾",
		Descriptor: "Oncoming Fist: Medium-Dark Skin Tone",
	},
	"1F44A-1F3FF": {
		Key:        "1F44A-1F3FF",
		Value:      "👊🏿",
		Descriptor: "Oncoming Fist: Dark Skin Tone",
	},
	"1F44B": {
		Key:        "1F44B",
		Value:      "👋",
		Descriptor: "Waving Hand",
	},
	"1F44B-1F3FB": {
		Key:        "1F44B-1F3FB",
		Value:      "👋🏻",
		Descriptor: "Waving Hand: Light Skin Tone",
	},
	"1F44B-1F3FC": {
		Key:        "1F44B-1F3FC",
		Value:      "👋🏼",
		Descriptor: "Waving Hand: Medium-Light Skin Tone",
	},
	"1F44B-1F3FD": {
		Key:        "1F44B-1F3FD",
		Value:      "👋🏽",
		Descriptor: "Waving Hand: Medium Skin Tone",
	},
	"1F44B-1F3FE": {
		Key:        "1F44B-1F3FE",
		Value:      "👋🏾",
		Descriptor: "Waving Hand: Medium-Dark Skin Tone",
	},
	"1F44B-1F3FF": {
		Key:        "1F44B-1F3FF",
		Value:      "👋🏿",
		Descriptor: "Waving Hand: Dark Skin Tone",
	},
	"1F44C": {
		Key:        "1F44C",
		Value:      "👌",
		Descriptor: "OK Hand",
	},
	"1F44C-1F3FB": {
		Key:        "1F44C-1F3FB",
		Value:      "👌🏻",
		Descriptor: "OK Hand: Light Skin Tone",
	},
	"1F44C-1F3FC": {
		Key:        "1F44C-1F3FC",
		Value:      "👌🏼",
		Descriptor: "OK Hand: Medium-Light Skin Tone",
	},
	"1F44C-1F3FD": {
		Key:        "1F44C-1F3FD",
		Value:      "👌🏽",
		Descriptor: "OK Hand: Medium Skin Tone",
	},
	"1F44C-1F3FE": {
		Key:        "1F44C-1F3FE",
		Value:      "👌🏾",
		Descriptor: "OK Hand: Medium-Dark Skin Tone",
	},
	"1F44C-1F3FF": {
		Key:        "1F44C-1F3FF",
		Value:      "👌🏿",
		Descriptor: "OK Hand: Dark Skin Tone",
	},
	"1F44D": {
		Key:        "1F44D",
		Value:      "👍",
		Descriptor: "Thumbs Up",
	},
	"1F44D-1F3FB": {
		Key:        "1F44D-1F3FB",
		Value:      "👍🏻",
		Descriptor: "Thumbs Up: Light Skin Tone",
	},
	"1F44D-1F3FC": {
		Key:        "1F44D-1F3FC",
		Value:      "👍🏼",
		Descriptor: "Thumbs Up: Medium-Light Skin Tone",
	},
	"1F44D-1F3FD": {
		Key:        "1F44D-1F3FD",
		Value:      "👍🏽",
		Descriptor: "Thumbs Up: Medium Skin Tone",
	},
	"1F44D-1F3FE": {
		Key:        "1F44D-1F3FE",
		Value:      "👍🏾",
		Descriptor: "Thumbs Up: Medium-Dark Skin Tone",
	},
	"1F44D-1F3FF": {
		Key:        "1F44D-1F3FF",
		Value:      "👍🏿",
		Descriptor: "Thumbs Up: Dark Skin Tone",
	},
	"1F44E": {
		Key:        "1F44E",
		Value:      "👎",
		Descriptor: "Thumbs Down",
	},
	"1F44E-1F3FB": {
		Key:        "1F44E-1F3FB",
		Value:      "👎🏻",
		Descriptor: "Thumbs Down: Light Skin Tone",
	},
	"1F44E-1F3FC": {
		Key:        "1F44E-1F3FC",
		Value:      "👎🏼",
		Descriptor: "Thumbs Down: Medium-Light Skin Tone",
	},
	"1F44E-1F3FD": {
		Key:        "1F44E-1F3FD",
		Value:      "👎🏽",
		Descriptor: "Thumbs Down: Medium Skin Tone",
	},
	"1F44E-1F3FE": {
		Key:        "1F44E-1F3FE",
		Value:      "👎🏾",
		Descriptor: "Thumbs Down: Medium-Dark Skin Tone",
	},
	"1F44E-1F3FF": {
		Key:        "1F44E-1F3FF",
		Value:      "👎🏿",
		Descriptor: "Thumbs Down: Dark Skin Tone",
	},
	"1F44F": {
		Key:        "1F44F",
		Value:      "👏",
		Descriptor: "Clapping Hands",
	},
	"1F44F-1F3FB": {
		Key:        "1F44F-1F3FB",
		Value:      "👏🏻",
		Descriptor: "Clapping Hands: Light Skin Tone",
	},
	"1F44F-1F3FC": {
		Key:        "1F44F-1F3FC",
		Value:      "👏🏼",
		Descriptor: "Clapping Hands: Medium-Light Skin Tone",
	},
	"1F44F-1F3FD": {
		Key:        "1F44F-1F3FD",
		Value:      "👏🏽",
		Descriptor: "Clapping Hands: Medium Skin Tone",
	},
	"1F44F-1F3FE": {
		Key:        "1F44F-1F3FE",
		Value:      "👏🏾",
		Descriptor: "Clapping Hands: Medium-Dark Skin Tone",
	},
	"1F44F-1F3FF": {
		Key:        "1F44F-1F3FF",
		Value:      "👏🏿",
		Descriptor: "Clapping Hands: Dark Skin Tone",
	},
	"1F450": {
		Key:        "1F450",
		Value:      "👐",
		Descriptor: "Open Hands",
	},
	"1F450-1F3FB": {
		Key:        "1F450-1F3FB",
		Value:      "👐🏻",
		Descriptor: "Open Hands: Light Skin Tone",
	},
	"1F450-1F3FC": {
		Key:        "1F450-1F3FC",
		Value:      "👐🏼",
		Descriptor: "Open Hands: Medium-Light Skin Tone",
	},
	"1F450-1F3FD": {
		Key:        "1F450-1F3FD",
		Value:      "👐🏽",
		Descriptor: "Open Hands: Medium Skin Tone",
	},
	"1F450-1F3FE": {
		Key:        "1F450-1F3FE",
		Value:      "👐🏾",
		Descriptor: "Open Hands: Medium-Dark Skin Tone",
	},
	"1F450-1F3FF": {
		Key:        "1F450-1F3FF",
		Value:      "👐🏿",
		Descriptor: "Open Hands: Dark Skin Tone",
	},
	"1F451": {
		Key:        "1F451",
		Value:      "👑",
		Descriptor: "Crown",
	},
	"1F452": {
		Key:        "1F452",
		Value:      "👒",
		Descriptor: "Woman’s Hat",
	},
	"1F453": {
		Key:        "1F453",
		Value:      "👓",
		Descriptor: "Glasses",
	},
	"1F454": {
		Key:        "1F454",
		Value:      "👔",
		Descriptor: "Necktie",
	},
	"1F455": {
		Key:        "1F455",
		Value:      "👕",
		Descriptor: "T-Shirt",
	},
	"1F456": {
		Key:        "1F456",
		Value:      "👖",
		Descriptor: "Jeans",
	},
	"1F457": {
		Key:        "1F457",
		Value:      "👗",
		Descriptor: "Dress",
	},
	"1F458": {
		Key:        "1F458",
		Value:      "👘",
		Descriptor: "Kimono",
	},
	"1F459": {
		Key:        "1F459",
		Value:      "👙",
		Descriptor: "Bikini",
	},
	"1F45A": {
		Key:        "1F45A",
		Value:      "👚",
		Descriptor: "Woman’s Clothes",
	},
	"1F45B": {
		Key:        "1F45B",
		Value:      "👛",
		Descriptor: "Purse",
	},
	"1F45C": {
		Key:        "1F45C",
		Value:      "👜",
		Descriptor: "Handbag",
	},
	"1F45D": {
		Key:        "1F45D",
		Value:      "👝",
		Descriptor: "Clutch Bag",
	},
	"1F45E": {
		Key:        "1F45E",
		Value:      "👞",
		Descriptor: "Man’s Shoe",
	},
	"1F45F": {
		Key:        "1F45F",
		Value:      "👟",
		Descriptor: "Running Shoe",
	},
	"1F460": {
		Key:        "1F460",
		Value:      "👠",
		Descriptor: "High-Heeled Shoe",
	},
	"1F461": {
		Key:        "1F461",
		Value:      "👡",
		Descriptor: "Woman’s Sandal",
	},
	"1F462": {
		Key:        "1F462",
		Value:      "👢",
		Descriptor: "Woman’s Boot",
	},
	"1F463": {
		Key:        "1F463",
		Value:      "👣",
		Descriptor: "Footprints",
	},
	"1F464": {
		Key:        "1F464",
		Value:      "👤",
		Descriptor: "Bust in Silhouette",
	},
	"1F465": {
		Key:        "1F465",
		Value:      "👥",
		Descriptor: "Busts in Silhouette",
	},
	"1F466": {
		Key:        "1F466",
		Value:      "👦",
		Descriptor: "Boy",
	},
	"1F466-1F3FB": {
		Key:        "1F466-1F3FB",
		Value:      "👦🏻",
		Descriptor: "Boy: Light Skin Tone",
	},
	"1F466-1F3FC": {
		Key:        "1F466-1F3FC",
		Value:      "👦🏼",
		Descriptor: "Boy: Medium-Light Skin Tone",
	},
	"1F466-1F3FD": {
		Key:        "1F466-1F3FD",
		Value:      "👦🏽",
		Descriptor: "Boy: Medium Skin Tone",
	},
	"1F466-1F3FE": {
		Key:        "1F466-1F3FE",
		Value:      "👦🏾",
		Descriptor: "Boy: Medium-Dark Skin Tone",
	},
	"1F466-1F3FF": {
		Key:        "1F466-1F3FF",
		Value:      "👦🏿",
		Descriptor: "Boy: Dark Skin Tone",
	},
	"1F467": {
		Key:        "1F467",
		Value:      "👧",
		Descriptor: "Girl",
	},
	"1F467-1F3FB": {
		Key:        "1F467-1F3FB",
		Value:      "👧🏻",
		Descriptor: "Girl: Light Skin Tone",
	},
	"1F467-1F3FC": {
		Key:        "1F467-1F3FC",
		Value:      "👧🏼",
		Descriptor: "Girl: Medium-Light Skin Tone",
	},
	"1F467-1F3FD": {
		Key:        "1F467-1F3FD",
		Value:      "👧🏽",
		Descriptor: "Girl: Medium Skin Tone",
	},
	"1F467-1F3FE": {
		Key:        "1F467-1F3FE",
		Value:      "👧🏾",
		Descriptor: "Girl: Medium-Dark Skin Tone",
	},
	"1F467-1F3FF": {
		Key:        "1F467-1F3FF",
		Value:      "👧🏿",
		Descriptor: "Girl: Dark Skin Tone",
	},
	"1F468": {
		Key:        "1F468",
		Value:      "👨",
		Descriptor: "Man",
	},
	"1F468-1F3FB": {
		Key:        "1F468-1F3FB",
		Value:      "👨🏻",
		Descriptor: "Man: Light Skin Tone",
	},
	"1F468-1F3FB-200D-1F33E": {
		Key:        "1F468-1F3FB-200D-1F33E",
		Value:      "👨🏻‍🌾",
		Descriptor: "Man Farmer: Light Skin Tone",
	},
	"1F468-1F3FB-200D-1F373": {
		Key:        "1F468-1F3FB-200D-1F373",
		Value:      "👨🏻‍🍳",
		Descriptor: "Man Cook: Light Skin Tone",
	},
	"1F468-1F3FB-200D-1F393": {
		Key:        "1F468-1F3FB-200D-1F393",
		Value:      "👨🏻‍🎓",
		Descriptor: "Man Student: Light Skin Tone",
	},
	"1F468-1F3FB-200D-1F3A4": {
		Key:        "1F468-1F3FB-200D-1F3A4",
		Value:      "👨🏻‍🎤",
		Descriptor: "Man Singer: Light Skin Tone",
	},
	"1F468-1F3FB-200D-1F3A8": {
		Key:        "1F468-1F3FB-200D-1F3A8",
		Value:      "👨🏻‍🎨",
		Descriptor: "Man Artist: Light Skin Tone",
	},
	"1F468-1F3FB-200D-1F3EB": {
		Key:        "1F468-1F3FB-200D-1F3EB",
		Value:      "👨🏻‍🏫",
		Descriptor: "Man Teacher: Light Skin Tone",
	},
	"1F468-1F3FB-200D-1F3ED": {
		Key:        "1F468-1F3FB-200D-1F3ED",
		Value:      "👨🏻‍🏭",
		Descriptor: "Man Factory Worker: Light Skin Tone",
	},
	"1F468-1F3FB-200D-1F4BB": {
		Key:        "1F468-1F3FB-200D-1F4BB",
		Value:      "👨🏻‍💻",
		Descriptor: "Man Technologist: Light Skin Tone",
	},
	"1F468-1F3FB-200D-1F4BC": {
		Key:        "1F468-1F3FB-200D-1F4BC",
		Value:      "👨🏻‍💼",
		Descriptor: "Man Office Worker: Light Skin Tone",
	},
	"1F468-1F3FB-200D-1F527": {
		Key:        "1F468-1F3FB-200D-1F527",
		Value:      "👨🏻‍🔧",
		Descriptor: "Man Mechanic: Light Skin Tone",
	},
	"1F468-1F3FB-200D-1F52C": {
		Key:        "1F468-1F3FB-200D-1F52C",
		Value:      "👨🏻‍🔬",
		Descriptor: "Man Scientist: Light Skin Tone",
	},
	"1F468-1F3FB-200D-1F680": {
		Key:        "1F468-1F3FB-200D-1F680",
		Value:      "👨🏻‍🚀",
		Descriptor: "Man Astronaut: Light Skin Tone",
	},
	"1F468-1F3FB-200D-1F692": {
		Key:        "1F468-1F3FB-200D-1F692",
		Value:      "👨🏻‍🚒",
		Descriptor: "Man Firefighter: Light Skin Tone",
	},
	"1F468-1F3FB-200D-1F91D-200D-1F468-1F3FC": {
		Key:        "1F468-1F3FB-200D-1F91D-200D-1F468-1F3FC",
		Value:      "👨🏻‍🤝‍👨🏼",
		Descriptor: "Men Holding Hands: Light Skin Tone, Medium-Light Skin Tone",
	},
	"1F468-1F3FB-200D-1F91D-200D-1F468-1F3FD": {
		Key:        "1F468-1F3FB-200D-1F91D-200D-1F468-1F3FD",
		Value:      "👨🏻‍🤝‍👨🏽",
		Descriptor: "Men Holding Hands: Light Skin Tone, Medium Skin Tone",
	},
	"1F468-1F3FB-200D-1F91D-200D-1F468-1F3FE": {
		Key:        "1F468-1F3FB-200D-1F91D-200D-1F468-1F3FE",
		Value:      "👨🏻‍🤝‍👨🏾",
		Descriptor: "Men Holding Hands: Light Skin Tone, Medium-Dark Skin Tone",
	},
	"1F468-1F3FB-200D-1F91D-200D-1F468-1F3FF": {
		Key:        "1F468-1F3FB-200D-1F91D-200D-1F468-1F3FF",
		Value:      "👨🏻‍🤝‍👨🏿",
		Descriptor: "Men Holding Hands: Light Skin Tone, Dark Skin Tone",
	},
	"1F468-1F3FB-200D-1F9AF": {
		Key:        "1F468-1F3FB-200D-1F9AF",
		Value:      "👨🏻‍🦯",
		Descriptor: "Man With Probing Cane: Light Skin Tone",
	},
	"1F468-1F3FB-200D-1F9B0": {
		Key:        "1F468-1F3FB-200D-1F9B0",
		Value:      "👨🏻‍🦰",
		Descriptor: "Man: Light Skin Tone, Red Hair",
	},
	"1F468-1F3FB-200D-1F9B1": {
		Key:        "1F468-1F3FB-200D-1F9B1",
		Value:      "👨🏻‍🦱",
		Descriptor: "Man: Light Skin Tone, Curly Hair",
	},
	"1F468-1F3FB-200D-1F9B2": {
		Key:        "1F468-1F3FB-200D-1F9B2",
		Value:      "👨🏻‍🦲",
		Descriptor: "Man: Light Skin Tone, Bald",
	},
	"1F468-1F3FB-200D-1F9B3": {
		Key:        "1F468-1F3FB-200D-1F9B3",
		Value:      "👨🏻‍🦳",
		Descriptor: "Man: Light Skin Tone, White Hair",
	},
	"1F468-1F3FB-200D-1F9BC": {
		Key:        "1F468-1F3FB-200D-1F9BC",
		Value:      "👨🏻‍🦼",
		Descriptor: "Man in Motorized Wheelchair: Light Skin Tone",
	},
	"1F468-1F3FB-200D-1F9BD": {
		Key:        "1F468-1F3FB-200D-1F9BD",
		Value:      "👨🏻‍🦽",
		Descriptor: "Man in Manual Wheelchair: Light Skin Tone",
	},
	"1F468-1F3FB-200D-2695-FE0F": {
		Key:        "1F468-1F3FB-200D-2695-FE0F",
		Value:      "👨🏻‍⚕️",
		Descriptor: "Man Health Worker: Light Skin Tone",
	},
	"1F468-1F3FB-200D-2696-FE0F": {
		Key:        "1F468-1F3FB-200D-2696-FE0F",
		Value:      "👨🏻‍⚖️",
		Descriptor: "Man Judge: Light Skin Tone",
	},
	"1F468-1F3FB-200D-2708-FE0F": {
		Key:        "1F468-1F3FB-200D-2708-FE0F",
		Value:      "👨🏻‍✈️",
		Descriptor: "Man Pilot: Light Skin Tone",
	},
	"1F468-1F3FC": {
		Key:        "1F468-1F3FC",
		Value:      "👨🏼",
		Descriptor: "Man: Medium-Light Skin Tone",
	},
	"1F468-1F3FC-200D-1F33E": {
		Key:        "1F468-1F3FC-200D-1F33E",
		Value:      "👨🏼‍🌾",
		Descriptor: "Man Farmer: Medium-Light Skin Tone",
	},
	"1F468-1F3FC-200D-1F373": {
		Key:        "1F468-1F3FC-200D-1F373",
		Value:      "👨🏼‍🍳",
		Descriptor: "Man Cook: Medium-Light Skin Tone",
	},
	"1F468-1F3FC-200D-1F393": {
		Key:        "1F468-1F3FC-200D-1F393",
		Value:      "👨🏼‍🎓",
		Descriptor: "Man Student: Medium-Light Skin Tone",
	},
	"1F468-1F3FC-200D-1F3A4": {
		Key:        "1F468-1F3FC-200D-1F3A4",
		Value:      "👨🏼‍🎤",
		Descriptor: "Man Singer: Medium-Light Skin Tone",
	},
	"1F468-1F3FC-200D-1F3A8": {
		Key:        "1F468-1F3FC-200D-1F3A8",
		Value:      "👨🏼‍🎨",
		Descriptor: "Man Artist: Medium-Light Skin Tone",
	},
	"1F468-1F3FC-200D-1F3EB": {
		Key:        "1F468-1F3FC-200D-1F3EB",
		Value:      "👨🏼‍🏫",
		Descriptor: "Man Teacher: Medium-Light Skin Tone",
	},
	"1F468-1F3FC-200D-1F3ED": {
		Key:        "1F468-1F3FC-200D-1F3ED",
		Value:      "👨🏼‍🏭",
		Descriptor: "Man Factory Worker: Medium-Light Skin Tone",
	},
	"1F468-1F3FC-200D-1F4BB": {
		Key:        "1F468-1F3FC-200D-1F4BB",
		Value:      "👨🏼‍💻",
		Descriptor: "Man Technologist: Medium-Light Skin Tone",
	},
	"1F468-1F3FC-200D-1F4BC": {
		Key:        "1F468-1F3FC-200D-1F4BC",
		Value:      "👨🏼‍💼",
		Descriptor: "Man Office Worker: Medium-Light Skin Tone",
	},
	"1F468-1F3FC-200D-1F527": {
		Key:        "1F468-1F3FC-200D-1F527",
		Value:      "👨🏼‍🔧",
		Descriptor: "Man Mechanic: Medium-Light Skin Tone",
	},
	"1F468-1F3FC-200D-1F52C": {
		Key:        "1F468-1F3FC-200D-1F52C",
		Value:      "👨🏼‍🔬",
		Descriptor: "Man Scientist: Medium-Light Skin Tone",
	},
	"1F468-1F3FC-200D-1F680": {
		Key:        "1F468-1F3FC-200D-1F680",
		Value:      "👨🏼‍🚀",
		Descriptor: "Man Astronaut: Medium-Light Skin Tone",
	},
	"1F468-1F3FC-200D-1F692": {
		Key:        "1F468-1F3FC-200D-1F692",
		Value:      "👨🏼‍🚒",
		Descriptor: "Man Firefighter: Medium-Light Skin Tone",
	},
	"1F468-1F3FC-200D-1F91D-200D-1F468-1F3FB": {
		Key:        "1F468-1F3FC-200D-1F91D-200D-1F468-1F3FB",
		Value:      "👨🏼‍🤝‍👨🏻",
		Descriptor: "Men Holding Hands: Medium-Light Skin Tone, Light Skin Tone",
	},
	"1F468-1F3FC-200D-1F91D-200D-1F468-1F3FD": {
		Key:        "1F468-1F3FC-200D-1F91D-200D-1F468-1F3FD",
		Value:      "👨🏼‍🤝‍👨🏽",
		Descriptor: "Men Holding Hands: Medium-Light Skin Tone, Medium Skin Tone",
	},
	"1F468-1F3FC-200D-1F91D-200D-1F468-1F3FE": {
		Key:        "1F468-1F3FC-200D-1F91D-200D-1F468-1F3FE",
		Value:      "👨🏼‍🤝‍👨🏾",
		Descriptor: "Men Holding Hands: Medium-Light Skin Tone, Medium-Dark Skin Tone",
	},
	"1F468-1F3FC-200D-1F91D-200D-1F468-1F3FF": {
		Key:        "1F468-1F3FC-200D-1F91D-200D-1F468-1F3FF",
		Value:      "👨🏼‍🤝‍👨🏿",
		Descriptor: "Men Holding Hands: Medium-Light Skin Tone, Dark Skin Tone",
	},
	"1F468-1F3FC-200D-1F9AF": {
		Key:        "1F468-1F3FC-200D-1F9AF",
		Value:      "👨🏼‍🦯",
		Descriptor: "Man With Probing Cane: Medium-Light Skin Tone",
	},
	"1F468-1F3FC-200D-1F9B0": {
		Key:        "1F468-1F3FC-200D-1F9B0",
		Value:      "👨🏼‍🦰",
		Descriptor: "Man: Medium-Light Skin Tone, Red Hair",
	},
	"1F468-1F3FC-200D-1F9B1": {
		Key:        "1F468-1F3FC-200D-1F9B1",
		Value:      "👨🏼‍🦱",
		Descriptor: "Man: Medium-Light Skin Tone, Curly Hair",
	},
	"1F468-1F3FC-200D-1F9B2": {
		Key:        "1F468-1F3FC-200D-1F9B2",
		Value:      "👨🏼‍🦲",
		Descriptor: "Man: Medium-Light Skin Tone, Bald",
	},
	"1F468-1F3FC-200D-1F9B3": {
		Key:        "1F468-1F3FC-200D-1F9B3",
		Value:      "👨🏼‍🦳",
		Descriptor: "Man: Medium-Light Skin Tone, White Hair",
	},
	"1F468-1F3FC-200D-1F9BC": {
		Key:        "1F468-1F3FC-200D-1F9BC",
		Value:      "👨🏼‍🦼",
		Descriptor: "Man in Motorized Wheelchair: Medium-Light Skin Tone",
	},
	"1F468-1F3FC-200D-1F9BD": {
		Key:        "1F468-1F3FC-200D-1F9BD",
		Value:      "👨🏼‍🦽",
		Descriptor: "Man in Manual Wheelchair: Medium-Light Skin Tone",
	},
	"1F468-1F3FC-200D-2695-FE0F": {
		Key:        "1F468-1F3FC-200D-2695-FE0F",
		Value:      "👨🏼‍⚕️",
		Descriptor: "Man Health Worker: Medium-Light Skin Tone",
	},
	"1F468-1F3FC-200D-2696-FE0F": {
		Key:        "1F468-1F3FC-200D-2696-FE0F",
		Value:      "👨🏼‍⚖️",
		Descriptor: "Man Judge: Medium-Light Skin Tone",
	},
	"1F468-1F3FC-200D-2708-FE0F": {
		Key:        "1F468-1F3FC-200D-2708-FE0F",
		Value:      "👨🏼‍✈️",
		Descriptor: "Man Pilot: Medium-Light Skin Tone",
	},
	"1F468-1F3FD": {
		Key:        "1F468-1F3FD",
		Value:      "👨🏽",
		Descriptor: "Man: Medium Skin Tone",
	},
	"1F468-1F3FD-200D-1F33E": {
		Key:        "1F468-1F3FD-200D-1F33E",
		Value:      "👨🏽‍🌾",
		Descriptor: "Man Farmer: Medium Skin Tone",
	},
	"1F468-1F3FD-200D-1F373": {
		Key:        "1F468-1F3FD-200D-1F373",
		Value:      "👨🏽‍🍳",
		Descriptor: "Man Cook: Medium Skin Tone",
	},
	"1F468-1F3FD-200D-1F393": {
		Key:        "1F468-1F3FD-200D-1F393",
		Value:      "👨🏽‍🎓",
		Descriptor: "Man Student: Medium Skin Tone",
	},
	"1F468-1F3FD-200D-1F3A4": {
		Key:        "1F468-1F3FD-200D-1F3A4",
		Value:      "👨🏽‍🎤",
		Descriptor: "Man Singer: Medium Skin Tone",
	},
	"1F468-1F3FD-200D-1F3A8": {
		Key:        "1F468-1F3FD-200D-1F3A8",
		Value:      "👨🏽‍🎨",
		Descriptor: "Man Artist: Medium Skin Tone",
	},
	"1F468-1F3FD-200D-1F3EB": {
		Key:        "1F468-1F3FD-200D-1F3EB",
		Value:      "👨🏽‍🏫",
		Descriptor: "Man Teacher: Medium Skin Tone",
	},
	"1F468-1F3FD-200D-1F3ED": {
		Key:        "1F468-1F3FD-200D-1F3ED",
		Value:      "👨🏽‍🏭",
		Descriptor: "Man Factory Worker: Medium Skin Tone",
	},
	"1F468-1F3FD-200D-1F4BB": {
		Key:        "1F468-1F3FD-200D-1F4BB",
		Value:      "👨🏽‍💻",
		Descriptor: "Man Technologist: Medium Skin Tone",
	},
	"1F468-1F3FD-200D-1F4BC": {
		Key:        "1F468-1F3FD-200D-1F4BC",
		Value:      "👨🏽‍💼",
		Descriptor: "Man Office Worker: Medium Skin Tone",
	},
	"1F468-1F3FD-200D-1F527": {
		Key:        "1F468-1F3FD-200D-1F527",
		Value:      "👨🏽‍🔧",
		Descriptor: "Man Mechanic: Medium Skin Tone",
	},
	"1F468-1F3FD-200D-1F52C": {
		Key:        "1F468-1F3FD-200D-1F52C",
		Value:      "👨🏽‍🔬",
		Descriptor: "Man Scientist: Medium Skin Tone",
	},
	"1F468-1F3FD-200D-1F680": {
		Key:        "1F468-1F3FD-200D-1F680",
		Value:      "👨🏽‍🚀",
		Descriptor: "Man Astronaut: Medium Skin Tone",
	},
	"1F468-1F3FD-200D-1F692": {
		Key:        "1F468-1F3FD-200D-1F692",
		Value:      "👨🏽‍🚒",
		Descriptor: "Man Firefighter: Medium Skin Tone",
	},
	"1F468-1F3FD-200D-1F91D-200D-1F468-1F3FB": {
		Key:        "1F468-1F3FD-200D-1F91D-200D-1F468-1F3FB",
		Value:      "👨🏽‍🤝‍👨🏻",
		Descriptor: "Men Holding Hands: Medium Skin Tone, Light Skin Tone",
	},
	"1F468-1F3FD-200D-1F91D-200D-1F468-1F3FC": {
		Key:        "1F468-1F3FD-200D-1F91D-200D-1F468-1F3FC",
		Value:      "👨🏽‍🤝‍👨🏼",
		Descriptor: "Men Holding Hands: Medium Skin Tone, Medium-Light Skin Tone",
	},
	"1F468-1F3FD-200D-1F91D-200D-1F468-1F3FE": {
		Key:        "1F468-1F3FD-200D-1F91D-200D-1F468-1F3FE",
		Value:      "👨🏽‍🤝‍👨🏾",
		Descriptor: "Men Holding Hands: Medium Skin Tone, Medium-Dark Skin Tone",
	},
	"1F468-1F3FD-200D-1F91D-200D-1F468-1F3FF": {
		Key:        "1F468-1F3FD-200D-1F91D-200D-1F468-1F3FF",
		Value:      "👨🏽‍🤝‍👨🏿",
		Descriptor: "Men Holding Hands: Medium Skin Tone, Dark Skin Tone",
	},
	"1F468-1F3FD-200D-1F9AF": {
		Key:        "1F468-1F3FD-200D-1F9AF",
		Value:      "👨🏽‍🦯",
		Descriptor: "Man With Probing Cane: Medium Skin Tone",
	},
	"1F468-1F3FD-200D-1F9B0": {
		Key:        "1F468-1F3FD-200D-1F9B0",
		Value:      "👨🏽‍🦰",
		Descriptor: "Man: Medium Skin Tone, Red Hair",
	},
	"1F468-1F3FD-200D-1F9B1": {
		Key:        "1F468-1F3FD-200D-1F9B1",
		Value:      "👨🏽‍🦱",
		Descriptor: "Man: Medium Skin Tone, Curly Hair",
	},
	"1F468-1F3FD-200D-1F9B2": {
		Key:        "1F468-1F3FD-200D-1F9B2",
		Value:      "👨🏽‍🦲",
		Descriptor: "Man: Medium Skin Tone, Bald",
	},
	"1F468-1F3FD-200D-1F9B3": {
		Key:        "1F468-1F3FD-200D-1F9B3",
		Value:      "👨🏽‍🦳",
		Descriptor: "Man: Medium Skin Tone, White Hair",
	},
	"1F468-1F3FD-200D-1F9BC": {
		Key:        "1F468-1F3FD-200D-1F9BC",
		Value:      "👨🏽‍🦼",
		Descriptor: "Man in Motorized Wheelchair: Medium Skin Tone",
	},
	"1F468-1F3FD-200D-1F9BD": {
		Key:        "1F468-1F3FD-200D-1F9BD",
		Value:      "👨🏽‍🦽",
		Descriptor: "Man in Manual Wheelchair: Medium Skin Tone",
	},
	"1F468-1F3FD-200D-2695-FE0F": {
		Key:        "1F468-1F3FD-200D-2695-FE0F",
		Value:      "👨🏽‍⚕️",
		Descriptor: "Man Health Worker: Medium Skin Tone",
	},
	"1F468-1F3FD-200D-2696-FE0F": {
		Key:        "1F468-1F3FD-200D-2696-FE0F",
		Value:      "👨🏽‍⚖️",
		Descriptor: "Man Judge: Medium Skin Tone",
	},
	"1F468-1F3FD-200D-2708-FE0F": {
		Key:        "1F468-1F3FD-200D-2708-FE0F",
		Value:      "👨🏽‍✈️",
		Descriptor: "Man Pilot: Medium Skin Tone",
	},
	"1F468-1F3FE": {
		Key:        "1F468-1F3FE",
		Value:      "👨🏾",
		Descriptor: "Man: Medium-Dark Skin Tone",
	},
	"1F468-1F3FE-200D-1F33E": {
		Key:        "1F468-1F3FE-200D-1F33E",
		Value:      "👨🏾‍🌾",
		Descriptor: "Man Farmer: Medium-Dark Skin Tone",
	},
	"1F468-1F3FE-200D-1F373": {
		Key:        "1F468-1F3FE-200D-1F373",
		Value:      "👨🏾‍🍳",
		Descriptor: "Man Cook: Medium-Dark Skin Tone",
	},
	"1F468-1F3FE-200D-1F393": {
		Key:        "1F468-1F3FE-200D-1F393",
		Value:      "👨🏾‍🎓",
		Descriptor: "Man Student: Medium-Dark Skin Tone",
	},
	"1F468-1F3FE-200D-1F3A4": {
		Key:        "1F468-1F3FE-200D-1F3A4",
		Value:      "👨🏾‍🎤",
		Descriptor: "Man Singer: Medium-Dark Skin Tone",
	},
	"1F468-1F3FE-200D-1F3A8": {
		Key:        "1F468-1F3FE-200D-1F3A8",
		Value:      "👨🏾‍🎨",
		Descriptor: "Man Artist: Medium-Dark Skin Tone",
	},
	"1F468-1F3FE-200D-1F3EB": {
		Key:        "1F468-1F3FE-200D-1F3EB",
		Value:      "👨🏾‍🏫",
		Descriptor: "Man Teacher: Medium-Dark Skin Tone",
	},
	"1F468-1F3FE-200D-1F3ED": {
		Key:        "1F468-1F3FE-200D-1F3ED",
		Value:      "👨🏾‍🏭",
		Descriptor: "Man Factory Worker: Medium-Dark Skin Tone",
	},
	"1F468-1F3FE-200D-1F4BB": {
		Key:        "1F468-1F3FE-200D-1F4BB",
		Value:      "👨🏾‍💻",
		Descriptor: "Man Technologist: Medium-Dark Skin Tone",
	},
	"1F468-1F3FE-200D-1F4BC": {
		Key:        "1F468-1F3FE-200D-1F4BC",
		Value:      "👨🏾‍💼",
		Descriptor: "Man Office Worker: Medium-Dark Skin Tone",
	},
	"1F468-1F3FE-200D-1F527": {
		Key:        "1F468-1F3FE-200D-1F527",
		Value:      "👨🏾‍🔧",
		Descriptor: "Man Mechanic: Medium-Dark Skin Tone",
	},
	"1F468-1F3FE-200D-1F52C": {
		Key:        "1F468-1F3FE-200D-1F52C",
		Value:      "👨🏾‍🔬",
		Descriptor: "Man Scientist: Medium-Dark Skin Tone",
	},
	"1F468-1F3FE-200D-1F680": {
		Key:        "1F468-1F3FE-200D-1F680",
		Value:      "👨🏾‍🚀",
		Descriptor: "Man Astronaut: Medium-Dark Skin Tone",
	},
	"1F468-1F3FE-200D-1F692": {
		Key:        "1F468-1F3FE-200D-1F692",
		Value:      "👨🏾‍🚒",
		Descriptor: "Man Firefighter: Medium-Dark Skin Tone",
	},
	"1F468-1F3FE-200D-1F91D-200D-1F468-1F3FB": {
		Key:        "1F468-1F3FE-200D-1F91D-200D-1F468-1F3FB",
		Value:      "👨🏾‍🤝‍👨🏻",
		Descriptor: "Men Holding Hands: Medium-Dark Skin Tone, Light Skin Tone",
	},
	"1F468-1F3FE-200D-1F91D-200D-1F468-1F3FC": {
		Key:        "1F468-1F3FE-200D-1F91D-200D-1F468-1F3FC",
		Value:      "👨🏾‍🤝‍👨🏼",
		Descriptor: "Men Holding Hands: Medium-Dark Skin Tone, Medium-Light Skin Tone",
	},
	"1F468-1F3FE-200D-1F91D-200D-1F468-1F3FD": {
		Key:        "1F468-1F3FE-200D-1F91D-200D-1F468-1F3FD",
		Value:      "👨🏾‍🤝‍👨🏽",
		Descriptor: "Men Holding Hands: Medium-Dark Skin Tone, Medium Skin Tone",
	},
	"1F468-1F3FE-200D-1F91D-200D-1F468-1F3FF": {
		Key:        "1F468-1F3FE-200D-1F91D-200D-1F468-1F3FF",
		Value:      "👨🏾‍🤝‍👨🏿",
		Descriptor: "Men Holding Hands: Medium-Dark Skin Tone, Dark Skin Tone",
	},
	"1F468-1F3FE-200D-1F9AF": {
		Key:        "1F468-1F3FE-200D-1F9AF",
		Value:      "👨🏾‍🦯",
		Descriptor: "Man With Probing Cane: Medium-Dark Skin Tone",
	},
	"1F468-1F3FE-200D-1F9B0": {
		Key:        "1F468-1F3FE-200D-1F9B0",
		Value:      "👨🏾‍🦰",
		Descriptor: "Man: Medium-Dark Skin Tone, Red Hair",
	},
	"1F468-1F3FE-200D-1F9B1": {
		Key:        "1F468-1F3FE-200D-1F9B1",
		Value:      "👨🏾‍🦱",
		Descriptor: "Man: Medium-Dark Skin Tone, Curly Hair",
	},
	"1F468-1F3FE-200D-1F9B2": {
		Key:        "1F468-1F3FE-200D-1F9B2",
		Value:      "👨🏾‍🦲",
		Descriptor: "Man: Medium-Dark Skin Tone, Bald",
	},
	"1F468-1F3FE-200D-1F9B3": {
		Key:        "1F468-1F3FE-200D-1F9B3",
		Value:      "👨🏾‍🦳",
		Descriptor: "Man: Medium-Dark Skin Tone, White Hair",
	},
	"1F468-1F3FE-200D-1F9BC": {
		Key:        "1F468-1F3FE-200D-1F9BC",
		Value:      "👨🏾‍🦼",
		Descriptor: "Man in Motorized Wheelchair: Medium-Dark Skin Tone",
	},
	"1F468-1F3FE-200D-1F9BD": {
		Key:        "1F468-1F3FE-200D-1F9BD",
		Value:      "👨🏾‍🦽",
		Descriptor: "Man in Manual Wheelchair: Medium-Dark Skin Tone",
	},
	"1F468-1F3FE-200D-2695-FE0F": {
		Key:        "1F468-1F3FE-200D-2695-FE0F",
		Value:      "👨🏾‍⚕️",
		Descriptor: "Man Health Worker: Medium-Dark Skin Tone",
	},
	"1F468-1F3FE-200D-2696-FE0F": {
		Key:        "1F468-1F3FE-200D-2696-FE0F",
		Value:      "👨🏾‍⚖️",
		Descriptor: "Man Judge: Medium-Dark Skin Tone",
	},
	"1F468-1F3FE-200D-2708-FE0F": {
		Key:        "1F468-1F3FE-200D-2708-FE0F",
		Value:      "👨🏾‍✈️",
		Descriptor: "Man Pilot: Medium-Dark Skin Tone",
	},
	"1F468-1F3FF": {
		Key:        "1F468-1F3FF",
		Value:      "👨🏿",
		Descriptor: "Man: Dark Skin Tone",
	},
	"1F468-1F3FF-200D-1F33E": {
		Key:        "1F468-1F3FF-200D-1F33E",
		Value:      "👨🏿‍🌾",
		Descriptor: "Man Farmer: Dark Skin Tone",
	},
	"1F468-1F3FF-200D-1F373": {
		Key:        "1F468-1F3FF-200D-1F373",
		Value:      "👨🏿‍🍳",
		Descriptor: "Man Cook: Dark Skin Tone",
	},
	"1F468-1F3FF-200D-1F393": {
		Key:        "1F468-1F3FF-200D-1F393",
		Value:      "👨🏿‍🎓",
		Descriptor: "Man Student: Dark Skin Tone",
	},
	"1F468-1F3FF-200D-1F3A4": {
		Key:        "1F468-1F3FF-200D-1F3A4",
		Value:      "👨🏿‍🎤",
		Descriptor: "Man Singer: Dark Skin Tone",
	},
	"1F468-1F3FF-200D-1F3A8": {
		Key:        "1F468-1F3FF-200D-1F3A8",
		Value:      "👨🏿‍🎨",
		Descriptor: "Man Artist: Dark Skin Tone",
	},
	"1F468-1F3FF-200D-1F3EB": {
		Key:        "1F468-1F3FF-200D-1F3EB",
		Value:      "👨🏿‍🏫",
		Descriptor: "Man Teacher: Dark Skin Tone",
	},
	"1F468-1F3FF-200D-1F3ED": {
		Key:        "1F468-1F3FF-200D-1F3ED",
		Value:      "👨🏿‍🏭",
		Descriptor: "Man Factory Worker: Dark Skin Tone",
	},
	"1F468-1F3FF-200D-1F4BB": {
		Key:        "1F468-1F3FF-200D-1F4BB",
		Value:      "👨🏿‍💻",
		Descriptor: "Man Technologist: Dark Skin Tone",
	},
	"1F468-1F3FF-200D-1F4BC": {
		Key:        "1F468-1F3FF-200D-1F4BC",
		Value:      "👨🏿‍💼",
		Descriptor: "Man Office Worker: Dark Skin Tone",
	},
	"1F468-1F3FF-200D-1F527": {
		Key:        "1F468-1F3FF-200D-1F527",
		Value:      "👨🏿‍🔧",
		Descriptor: "Man Mechanic: Dark Skin Tone",
	},
	"1F468-1F3FF-200D-1F52C": {
		Key:        "1F468-1F3FF-200D-1F52C",
		Value:      "👨🏿‍🔬",
		Descriptor: "Man Scientist: Dark Skin Tone",
	},
	"1F468-1F3FF-200D-1F680": {
		Key:        "1F468-1F3FF-200D-1F680",
		Value:      "👨🏿‍🚀",
		Descriptor: "Man Astronaut: Dark Skin Tone",
	},
	"1F468-1F3FF-200D-1F692": {
		Key:        "1F468-1F3FF-200D-1F692",
		Value:      "👨🏿‍🚒",
		Descriptor: "Man Firefighter: Dark Skin Tone",
	},
	"1F468-1F3FF-200D-1F91D-200D-1F468-1F3FB": {
		Key:        "1F468-1F3FF-200D-1F91D-200D-1F468-1F3FB",
		Value:      "👨🏿‍🤝‍👨🏻",
		Descriptor: "Men Holding Hands: Dark Skin Tone, Light Skin Tone",
	},
	"1F468-1F3FF-200D-1F91D-200D-1F468-1F3FC": {
		Key:        "1F468-1F3FF-200D-1F91D-200D-1F468-1F3FC",
		Value:      "👨🏿‍🤝‍👨🏼",
		Descriptor: "Men Holding Hands: Dark Skin Tone, Medium-Light Skin Tone",
	},
	"1F468-1F3FF-200D-1F91D-200D-1F468-1F3FD": {
		Key:        "1F468-1F3FF-200D-1F91D-200D-1F468-1F3FD",
		Value:      "👨🏿‍🤝‍👨🏽",
		Descriptor: "Men Holding Hands: Dark Skin Tone, Medium Skin Tone",
	},
	"1F468-1F3FF-200D-1F91D-200D-1F468-1F3FE": {
		Key:        "1F468-1F3FF-200D-1F91D-200D-1F468-1F3FE",
		Value:      "👨🏿‍🤝‍👨🏾",
		Descriptor: "Men Holding Hands: Dark Skin Tone, Medium-Dark Skin Tone",
	},
	"1F468-1F3FF-200D-1F9AF": {
		Key:        "1F468-1F3FF-200D-1F9AF",
		Value:      "👨🏿‍🦯",
		Descriptor: "Man With Probing Cane: Dark Skin Tone",
	},
	"1F468-1F3FF-200D-1F9B0": {
		Key:        "1F468-1F3FF-200D-1F9B0",
		Value:      "👨🏿‍🦰",
		Descriptor: "Man: Dark Skin Tone, Red Hair",
	},
	"1F468-1F3FF-200D-1F9B1": {
		Key:        "1F468-1F3FF-200D-1F9B1",
		Value:      "👨🏿‍🦱",
		Descriptor: "Man: Dark Skin Tone, Curly Hair",
	},
	"1F468-1F3FF-200D-1F9B2": {
		Key:        "1F468-1F3FF-200D-1F9B2",
		Value:      "👨🏿‍🦲",
		Descriptor: "Man: Dark Skin Tone, Bald",
	},
	"1F468-1F3FF-200D-1F9B3": {
		Key:        "1F468-1F3FF-200D-1F9B3",
		Value:      "👨🏿‍🦳",
		Descriptor: "Man: Dark Skin Tone, White Hair",
	},
	"1F468-1F3FF-200D-1F9BC": {
		Key:        "1F468-1F3FF-200D-1F9BC",
		Value:      "👨🏿‍🦼",
		Descriptor: "Man in Motorized Wheelchair: Dark Skin Tone",
	},
	"1F468-1F3FF-200D-1F9BD": {
		Key:        "1F468-1F3FF-200D-1F9BD",
		Value:      "👨🏿‍🦽",
		Descriptor: "Man in Manual Wheelchair: Dark Skin Tone",
	},
	"1F468-1F3FF-200D-2695-FE0F": {
		Key:        "1F468-1F3FF-200D-2695-FE0F",
		Value:      "👨🏿‍⚕️",
		Descriptor: "Man Health Worker: Dark Skin Tone",
	},
	"1F468-1F3FF-200D-2696-FE0F": {
		Key:        "1F468-1F3FF-200D-2696-FE0F",
		Value:      "👨🏿‍⚖️",
		Descriptor: "Man Judge: Dark Skin Tone",
	},
	"1F468-1F3FF-200D-2708-FE0F": {
		Key:        "1F468-1F3FF-200D-2708-FE0F",
		Value:      "👨🏿‍✈️",
		Descriptor: "Man Pilot: Dark Skin Tone",
	},
	"1F468-200D-1F33E": {
		Key:        "1F468-200D-1F33E",
		Value:      "👨‍🌾",
		Descriptor: "Man Farmer",
	},
	"1F468-200D-1F373": {
		Key:        "1F468-200D-1F373",
		Value:      "👨‍🍳",
		Descriptor: "Man Cook",
	},
	"1F468-200D-1F393": {
		Key:        "1F468-200D-1F393",
		Value:      "👨‍🎓",
		Descriptor: "Man Student",
	},
	"1F468-200D-1F3A4": {
		Key:        "1F468-200D-1F3A4",
		Value:      "👨‍🎤",
		Descriptor: "Man Singer",
	},
	"1F468-200D-1F3A8": {
		Key:        "1F468-200D-1F3A8",
		Value:      "👨‍🎨",
		Descriptor: "Man Artist",
	},
	"1F468-200D-1F3EB": {
		Key:        "1F468-200D-1F3EB",
		Value:      "👨‍🏫",
		Descriptor: "Man Teacher",
	},
	"1F468-200D-1F3ED": {
		Key:        "1F468-200D-1F3ED",
		Value:      "👨‍🏭",
		Descriptor: "Man Factory Worker",
	},
	"1F468-200D-1F466": {
		Key:        "1F468-200D-1F466",
		Value:      "👨‍👦",
		Descriptor: "Family: Man, Boy",
	},
	"1F468-200D-1F466-200D-1F466": {
		Key:        "1F468-200D-1F466-200D-1F466",
		Value:      "👨‍👦‍👦",
		Descriptor: "Family: Man, Boy, Boy",
	},
	"1F468-200D-1F467": {
		Key:        "1F468-200D-1F467",
		Value:      "👨‍👧",
		Descriptor: "Family: Man, Girl",
	},
	"1F468-200D-1F467-200D-1F466": {
		Key:        "1F468-200D-1F467-200D-1F466",
		Value:      "👨‍👧‍👦",
		Descriptor: "Family: Man, Girl, Boy",
	},
	"1F468-200D-1F467-200D-1F467": {
		Key:        "1F468-200D-1F467-200D-1F467",
		Value:      "👨‍👧‍👧",
		Descriptor: "Family: Man, Girl, Girl",
	},
	"1F468-200D-1F468-200D-1F466": {
		Key:        "1F468-200D-1F468-200D-1F466",
		Value:      "👨‍👨‍👦",
		Descriptor: "Family: Man, Man, Boy",
	},
	"1F468-200D-1F468-200D-1F466-200D-1F466": {
		Key:        "1F468-200D-1F468-200D-1F466-200D-1F466",
		Value:      "👨‍👨‍👦‍👦",
		Descriptor: "Family: Man, Man, Boy, Boy",
	},
	"1F468-200D-1F468-200D-1F467": {
		Key:        "1F468-200D-1F468-200D-1F467",
		Value:      "👨‍👨‍👧",
		Descriptor: "Family: Man, Man, Girl",
	},
	"1F468-200D-1F468-200D-1F467-200D-1F466": {
		Key:        "1F468-200D-1F468-200D-1F467-200D-1F466",
		Value:      "👨‍👨‍👧‍👦",
		Descriptor: "Family: Man, Man, Girl, Boy",
	},
	"1F468-200D-1F468-200D-1F467-200D-1F467": {
		Key:        "1F468-200D-1F468-200D-1F467-200D-1F467",
		Value:      "👨‍👨‍👧‍👧",
		Descriptor: "Family: Man, Man, Girl, Girl",
	},
	"1F468-200D-1F469-200D-1F466": {
		Key:        "1F468-200D-1F469-200D-1F466",
		Value:      "👨‍👩‍👦",
		Descriptor: "Family: Man, Woman, Boy",
	},
	"1F468-200D-1F469-200D-1F466-200D-1F466": {
		Key:        "1F468-200D-1F469-200D-1F466-200D-1F466",
		Value:      "👨‍👩‍👦‍👦",
		Descriptor: "Family: Man, Woman, Boy, Boy",
	},
	"1F468-200D-1F469-200D-1F467": {
		Key:        "1F468-200D-1F469-200D-1F467",
		Value:      "👨‍👩‍👧",
		Descriptor: "Family: Man, Woman, Girl",
	},
	"1F468-200D-1F469-200D-1F467-200D-1F466": {
		Key:        "1F468-200D-1F469-200D-1F467-200D-1F466",
		Value:      "👨‍👩‍👧‍👦",
		Descriptor: "Family: Man, Woman, Girl, Boy",
	},
	"1F468-200D-1F469-200D-1F467-200D-1F467": {
		Key:        "1F468-200D-1F469-200D-1F467-200D-1F467",
		Value:      "👨‍👩‍👧‍👧",
		Descriptor: "Family: Man, Woman, Girl, Girl",
	},
	"1F468-200D-1F4BB": {
		Key:        "1F468-200D-1F4BB",
		Value:      "👨‍💻",
		Descriptor: "Man Technologist",
	},
	"1F468-200D-1F4BC": {
		Key:        "1F468-200D-1F4BC",
		Value:      "👨‍💼",
		Descriptor: "Man Office Worker",
	},
	"1F468-200D-1F527": {
		Key:        "1F468-200D-1F527",
		Value:      "👨‍🔧",
		Descriptor: "Man Mechanic",
	},
	"1F468-200D-1F52C": {
		Key:        "1F468-200D-1F52C",
		Value:      "👨‍🔬",
		Descriptor: "Man Scientist",
	},
	"1F468-200D-1F680": {
		Key:        "1F468-200D-1F680",
		Value:      "👨‍🚀",
		Descriptor: "Man Astronaut",
	},
	"1F468-200D-1F692": {
		Key:        "1F468-200D-1F692",
		Value:      "👨‍🚒",
		Descriptor: "Man Firefighter",
	},
	"1F468-200D-1F9AF": {
		Key:        "1F468-200D-1F9AF",
		Value:      "👨‍🦯",
		Descriptor: "Man With Probing Cane",
	},
	"1F468-200D-1F9B0": {
		Key:        "1F468-200D-1F9B0",
		Value:      "👨‍🦰",
		Descriptor: "Man: Red Hair",
	},
	"1F468-200D-1F9B1": {
		Key:        "1F468-200D-1F9B1",
		Value:      "👨‍🦱",
		Descriptor: "Man: Curly Hair",
	},
	"1F468-200D-1F9B2": {
		Key:        "1F468-200D-1F9B2",
		Value:      "👨‍🦲",
		Descriptor: "Man: Bald",
	},
	"1F468-200D-1F9B3": {
		Key:        "1F468-200D-1F9B3",
		Value:      "👨‍🦳",
		Descriptor: "Man: White Hair",
	},
	"1F468-200D-1F9BC": {
		Key:        "1F468-200D-1F9BC",
		Value:      "👨‍🦼",
		Descriptor: "Man in Motorized Wheelchair",
	},
	"1F468-200D-1F9BD": {
		Key:        "1F468-200D-1F9BD",
		Value:      "👨‍🦽",
		Descriptor: "Man in Manual Wheelchair",
	},
	"1F468-200D-2695-FE0F": {
		Key:        "1F468-200D-2695-FE0F",
		Value:      "👨‍⚕️",
		Descriptor: "Man Health Worker",
	},
	"1F468-200D-2696-FE0F": {
		Key:        "1F468-200D-2696-FE0F",
		Value:      "👨‍⚖️",
		Descriptor: "Man Judge",
	},
	"1F468-200D-2708-FE0F": {
		Key:        "1F468-200D-2708-FE0F",
		Value:      "👨‍✈️",
		Descriptor: "Man Pilot",
	},
	"1F468-200D-2764-FE0F-200D-1F468": {
		Key:        "1F468-200D-2764-FE0F-200D-1F468",
		Value:      "👨‍❤️‍👨",
		Descriptor: "Couple With Heart: Man, Man",
	},
	"1F468-200D-2764-FE0F-200D-1F48B-200D-1F468": {
		Key:        "1F468-200D-2764-FE0F-200D-1F48B-200D-1F468",
		Value:      "👨‍❤️‍💋‍👨",
		Descriptor: "Kiss: Man, Man",
	},
	"1F469": {
		Key:        "1F469",
		Value:      "👩",
		Descriptor: "Woman",
	},
	"1F469-1F3FB": {
		Key:        "1F469-1F3FB",
		Value:      "👩🏻",
		Descriptor: "Woman: Light Skin Tone",
	},
	"1F469-1F3FB-200D-1F33E": {
		Key:        "1F469-1F3FB-200D-1F33E",
		Value:      "👩🏻‍🌾",
		Descriptor: "Woman Farmer: Light Skin Tone",
	},
	"1F469-1F3FB-200D-1F373": {
		Key:        "1F469-1F3FB-200D-1F373",
		Value:      "👩🏻‍🍳",
		Descriptor: "Woman Cook: Light Skin Tone",
	},
	"1F469-1F3FB-200D-1F393": {
		Key:        "1F469-1F3FB-200D-1F393",
		Value:      "👩🏻‍🎓",
		Descriptor: "Woman Student: Light Skin Tone",
	},
	"1F469-1F3FB-200D-1F3A4": {
		Key:        "1F469-1F3FB-200D-1F3A4",
		Value:      "👩🏻‍🎤",
		Descriptor: "Woman Singer: Light Skin Tone",
	},
	"1F469-1F3FB-200D-1F3A8": {
		Key:        "1F469-1F3FB-200D-1F3A8",
		Value:      "👩🏻‍🎨",
		Descriptor: "Woman Artist: Light Skin Tone",
	},
	"1F469-1F3FB-200D-1F3EB": {
		Key:        "1F469-1F3FB-200D-1F3EB",
		Value:      "👩🏻‍🏫",
		Descriptor: "Woman Teacher: Light Skin Tone",
	},
	"1F469-1F3FB-200D-1F3ED": {
		Key:        "1F469-1F3FB-200D-1F3ED",
		Value:      "👩🏻‍🏭",
		Descriptor: "Woman Factory Worker: Light Skin Tone",
	},
	"1F469-1F3FB-200D-1F4BB": {
		Key:        "1F469-1F3FB-200D-1F4BB",
		Value:      "👩🏻‍💻",
		Descriptor: "Woman Technologist: Light Skin Tone",
	},
	"1F469-1F3FB-200D-1F4BC": {
		Key:        "1F469-1F3FB-200D-1F4BC",
		Value:      "👩🏻‍💼",
		Descriptor: "Woman Office Worker: Light Skin Tone",
	},
	"1F469-1F3FB-200D-1F527": {
		Key:        "1F469-1F3FB-200D-1F527",
		Value:      "👩🏻‍🔧",
		Descriptor: "Woman Mechanic: Light Skin Tone",
	},
	"1F469-1F3FB-200D-1F52C": {
		Key:        "1F469-1F3FB-200D-1F52C",
		Value:      "👩🏻‍🔬",
		Descriptor: "Woman Scientist: Light Skin Tone",
	},
	"1F469-1F3FB-200D-1F680": {
		Key:        "1F469-1F3FB-200D-1F680",
		Value:      "👩🏻‍🚀",
		Descriptor: "Woman Astronaut: Light Skin Tone",
	},
	"1F469-1F3FB-200D-1F692": {
		Key:        "1F469-1F3FB-200D-1F692",
		Value:      "👩🏻‍🚒",
		Descriptor: "Woman Firefighter: Light Skin Tone",
	},
	"1F469-1F3FB-200D-1F91D-200D-1F468-1F3FC": {
		Key:        "1F469-1F3FB-200D-1F91D-200D-1F468-1F3FC",
		Value:      "👩🏻‍🤝‍👨🏼",
		Descriptor: "Woman and Man Holding Hands: Light Skin Tone, Medium-Light Skin Tone",
	},
	"1F469-1F3FB-200D-1F91D-200D-1F468-1F3FD": {
		Key:        "1F469-1F3FB-200D-1F91D-200D-1F468-1F3FD",
		Value:      "👩🏻‍🤝‍👨🏽",
		Descriptor: "Woman and Man Holding Hands: Light Skin Tone, Medium Skin Tone",
	},
	"1F469-1F3FB-200D-1F91D-200D-1F468-1F3FE": {
		Key:        "1F469-1F3FB-200D-1F91D-200D-1F468-1F3FE",
		Value:      "👩🏻‍🤝‍👨🏾",
		Descriptor: "Woman and Man Holding Hands: Light Skin Tone, Medium-Dark Skin Tone",
	},
	"1F469-1F3FB-200D-1F91D-200D-1F468-1F3FF": {
		Key:        "1F469-1F3FB-200D-1F91D-200D-1F468-1F3FF",
		Value:      "👩🏻‍🤝‍👨🏿",
		Descriptor: "Woman and Man Holding Hands: Light Skin Tone, Dark Skin Tone",
	},
	"1F469-1F3FB-200D-1F91D-200D-1F469-1F3FC": {
		Key:        "1F469-1F3FB-200D-1F91D-200D-1F469-1F3FC",
		Value:      "👩🏻‍🤝‍👩🏼",
		Descriptor: "Women Holding Hands: Light Skin Tone, Medium-Light Skin Tone",
	},
	"1F469-1F3FB-200D-1F91D-200D-1F469-1F3FD": {
		Key:        "1F469-1F3FB-200D-1F91D-200D-1F469-1F3FD",
		Value:      "👩🏻‍🤝‍👩🏽",
		Descriptor: "Women Holding Hands: Light Skin Tone, Medium Skin Tone",
	},
	"1F469-1F3FB-200D-1F91D-200D-1F469-1F3FE": {
		Key:        "1F469-1F3FB-200D-1F91D-200D-1F469-1F3FE",
		Value:      "👩🏻‍🤝‍👩🏾",
		Descriptor: "Women Holding Hands: Light Skin Tone, Medium-Dark Skin Tone",
	},
	"1F469-1F3FB-200D-1F91D-200D-1F469-1F3FF": {
		Key:        "1F469-1F3FB-200D-1F91D-200D-1F469-1F3FF",
		Value:      "👩🏻‍🤝‍👩🏿",
		Descriptor: "Women Holding Hands: Light Skin Tone, Dark Skin Tone",
	},
	"1F469-1F3FB-200D-1F9AF": {
		Key:        "1F469-1F3FB-200D-1F9AF",
		Value:      "👩🏻‍🦯",
		Descriptor: "Woman With Probing Cane: Light Skin Tone",
	},
	"1F469-1F3FB-200D-1F9B0": {
		Key:        "1F469-1F3FB-200D-1F9B0",
		Value:      "👩🏻‍🦰",
		Descriptor: "Woman: Light Skin Tone, Red Hair",
	},
	"1F469-1F3FB-200D-1F9B1": {
		Key:        "1F469-1F3FB-200D-1F9B1",
		Value:      "👩🏻‍🦱",
		Descriptor: "Woman: Light Skin Tone, Curly Hair",
	},
	"1F469-1F3FB-200D-1F9B2": {
		Key:        "1F469-1F3FB-200D-1F9B2",
		Value:      "👩🏻‍🦲",
		Descriptor: "Woman: Light Skin Tone, Bald",
	},
	"1F469-1F3FB-200D-1F9B3": {
		Key:        "1F469-1F3FB-200D-1F9B3",
		Value:      "👩🏻‍🦳",
		Descriptor: "Woman: Light Skin Tone, White Hair",
	},
	"1F469-1F3FB-200D-1F9BC": {
		Key:        "1F469-1F3FB-200D-1F9BC",
		Value:      "👩🏻‍🦼",
		Descriptor: "Woman in Motorized Wheelchair: Light Skin Tone",
	},
	"1F469-1F3FB-200D-1F9BD": {
		Key:        "1F469-1F3FB-200D-1F9BD",
		Value:      "👩🏻‍🦽",
		Descriptor: "Woman in Manual Wheelchair: Light Skin Tone",
	},
	"1F469-1F3FB-200D-2695-FE0F": {
		Key:        "1F469-1F3FB-200D-2695-FE0F",
		Value:      "👩🏻‍⚕️",
		Descriptor: "Woman Health Worker: Light Skin Tone",
	},
	"1F469-1F3FB-200D-2696-FE0F": {
		Key:        "1F469-1F3FB-200D-2696-FE0F",
		Value:      "👩🏻‍⚖️",
		Descriptor: "Woman Judge: Light Skin Tone",
	},
	"1F469-1F3FB-200D-2708-FE0F": {
		Key:        "1F469-1F3FB-200D-2708-FE0F",
		Value:      "👩🏻‍✈️",
		Descriptor: "Woman Pilot: Light Skin Tone",
	},
	"1F469-1F3FC": {
		Key:        "1F469-1F3FC",
		Value:      "👩🏼",
		Descriptor: "Woman: Medium-Light Skin Tone",
	},
	"1F469-1F3FC-200D-1F33E": {
		Key:        "1F469-1F3FC-200D-1F33E",
		Value:      "👩🏼‍🌾",
		Descriptor: "Woman Farmer: Medium-Light Skin Tone",
	},
	"1F469-1F3FC-200D-1F373": {
		Key:        "1F469-1F3FC-200D-1F373",
		Value:      "👩🏼‍🍳",
		Descriptor: "Woman Cook: Medium-Light Skin Tone",
	},
	"1F469-1F3FC-200D-1F393": {
		Key:        "1F469-1F3FC-200D-1F393",
		Value:      "👩🏼‍🎓",
		Descriptor: "Woman Student: Medium-Light Skin Tone",
	},
	"1F469-1F3FC-200D-1F3A4": {
		Key:        "1F469-1F3FC-200D-1F3A4",
		Value:      "👩🏼‍🎤",
		Descriptor: "Woman Singer: Medium-Light Skin Tone",
	},
	"1F469-1F3FC-200D-1F3A8": {
		Key:        "1F469-1F3FC-200D-1F3A8",
		Value:      "👩🏼‍🎨",
		Descriptor: "Woman Artist: Medium-Light Skin Tone",
	},
	"1F469-1F3FC-200D-1F3EB": {
		Key:        "1F469-1F3FC-200D-1F3EB",
		Value:      "👩🏼‍🏫",
		Descriptor: "Woman Teacher: Medium-Light Skin Tone",
	},
	"1F469-1F3FC-200D-1F3ED": {
		Key:        "1F469-1F3FC-200D-1F3ED",
		Value:      "👩🏼‍🏭",
		Descriptor: "Woman Factory Worker: Medium-Light Skin Tone",
	},
	"1F469-1F3FC-200D-1F4BB": {
		Key:        "1F469-1F3FC-200D-1F4BB",
		Value:      "👩🏼‍💻",
		Descriptor: "Woman Technologist: Medium-Light Skin Tone",
	},
	"1F469-1F3FC-200D-1F4BC": {
		Key:        "1F469-1F3FC-200D-1F4BC",
		Value:      "👩🏼‍💼",
		Descriptor: "Woman Office Worker: Medium-Light Skin Tone",
	},
	"1F469-1F3FC-200D-1F527": {
		Key:        "1F469-1F3FC-200D-1F527",
		Value:      "👩🏼‍🔧",
		Descriptor: "Woman Mechanic: Medium-Light Skin Tone",
	},
	"1F469-1F3FC-200D-1F52C": {
		Key:        "1F469-1F3FC-200D-1F52C",
		Value:      "👩🏼‍🔬",
		Descriptor: "Woman Scientist: Medium-Light Skin Tone",
	},
	"1F469-1F3FC-200D-1F680": {
		Key:        "1F469-1F3FC-200D-1F680",
		Value:      "👩🏼‍🚀",
		Descriptor: "Woman Astronaut: Medium-Light Skin Tone",
	},
	"1F469-1F3FC-200D-1F692": {
		Key:        "1F469-1F3FC-200D-1F692",
		Value:      "👩🏼‍🚒",
		Descriptor: "Woman Firefighter: Medium-Light Skin Tone",
	},
	"1F469-1F3FC-200D-1F91D-200D-1F468-1F3FB": {
		Key:        "1F469-1F3FC-200D-1F91D-200D-1F468-1F3FB",
		Value:      "👩🏼‍🤝‍👨🏻",
		Descriptor: "Woman and Man Holding Hands: Medium-Light Skin Tone, Light Skin Tone",
	},
	"1F469-1F3FC-200D-1F91D-200D-1F468-1F3FD": {
		Key:        "1F469-1F3FC-200D-1F91D-200D-1F468-1F3FD",
		Value:      "👩🏼‍🤝‍👨🏽",
		Descriptor: "Woman and Man Holding Hands: Medium-Light Skin Tone, Medium Skin Tone",
	},
	"1F469-1F3FC-200D-1F91D-200D-1F468-1F3FE": {
		Key:        "1F469-1F3FC-200D-1F91D-200D-1F468-1F3FE",
		Value:      "👩🏼‍🤝‍👨🏾",
		Descriptor: "Woman and Man Holding Hands: Medium-Light Skin Tone, Medium-Dark Skin Tone",
	},
	"1F469-1F3FC-200D-1F91D-200D-1F468-1F3FF": {
		Key:        "1F469-1F3FC-200D-1F91D-200D-1F468-1F3FF",
		Value:      "👩🏼‍🤝‍👨🏿",
		Descriptor: "Woman and Man Holding Hands: Medium-Light Skin Tone, Dark Skin Tone",
	},
	"1F469-1F3FC-200D-1F91D-200D-1F469-1F3FB": {
		Key:        "1F469-1F3FC-200D-1F91D-200D-1F469-1F3FB",
		Value:      "👩🏼‍🤝‍👩🏻",
		Descriptor: "Women Holding Hands: Medium-Light Skin Tone, Light Skin Tone",
	},
	"1F469-1F3FC-200D-1F91D-200D-1F469-1F3FD": {
		Key:        "1F469-1F3FC-200D-1F91D-200D-1F469-1F3FD",
		Value:      "👩🏼‍🤝‍👩🏽",
		Descriptor: "Women Holding Hands: Medium-Light Skin Tone, Medium Skin Tone",
	},
	"1F469-1F3FC-200D-1F91D-200D-1F469-1F3FE": {
		Key:        "1F469-1F3FC-200D-1F91D-200D-1F469-1F3FE",
		Value:      "👩🏼‍🤝‍👩🏾",
		Descriptor: "Women Holding Hands: Medium-Light Skin Tone, Medium-Dark Skin Tone",
	},
	"1F469-1F3FC-200D-1F91D-200D-1F469-1F3FF": {
		Key:        "1F469-1F3FC-200D-1F91D-200D-1F469-1F3FF",
		Value:      "👩🏼‍🤝‍👩🏿",
		Descriptor: "Women Holding Hands: Medium-Light Skin Tone, Dark Skin Tone",
	},
	"1F469-1F3FC-200D-1F9AF": {
		Key:        "1F469-1F3FC-200D-1F9AF",
		Value:      "👩🏼‍🦯",
		Descriptor: "Woman With Probing Cane: Medium-Light Skin Tone",
	},
	"1F469-1F3FC-200D-1F9B0": {
		Key:        "1F469-1F3FC-200D-1F9B0",
		Value:      "👩🏼‍🦰",
		Descriptor: "Woman: Medium-Light Skin Tone, Red Hair",
	},
	"1F469-1F3FC-200D-1F9B1": {
		Key:        "1F469-1F3FC-200D-1F9B1",
		Value:      "👩🏼‍🦱",
		Descriptor: "Woman: Medium-Light Skin Tone, Curly Hair",
	},
	"1F469-1F3FC-200D-1F9B2": {
		Key:        "1F469-1F3FC-200D-1F9B2",
		Value:      "👩🏼‍🦲",
		Descriptor: "Woman: Medium-Light Skin Tone, Bald",
	},
	"1F469-1F3FC-200D-1F9B3": {
		Key:        "1F469-1F3FC-200D-1F9B3",
		Value:      "👩🏼‍🦳",
		Descriptor: "Woman: Medium-Light Skin Tone, White Hair",
	},
	"1F469-1F3FC-200D-1F9BC": {
		Key:        "1F469-1F3FC-200D-1F9BC",
		Value:      "👩🏼‍🦼",
		Descriptor: "Woman in Motorized Wheelchair: Medium-Light Skin Tone",
	},
	"1F469-1F3FC-200D-1F9BD": {
		Key:        "1F469-1F3FC-200D-1F9BD",
		Value:      "👩🏼‍🦽",
		Descriptor: "Woman in Manual Wheelchair: Medium-Light Skin Tone",
	},
	"1F469-1F3FC-200D-2695-FE0F": {
		Key:        "1F469-1F3FC-200D-2695-FE0F",
		Value:      "👩🏼‍⚕️",
		Descriptor: "Woman Health Worker: Medium-Light Skin Tone",
	},
	"1F469-1F3FC-200D-2696-FE0F": {
		Key:        "1F469-1F3FC-200D-2696-FE0F",
		Value:      "👩🏼‍⚖️",
		Descriptor: "Woman Judge: Medium-Light Skin Tone",
	},
	"1F469-1F3FC-200D-2708-FE0F": {
		Key:        "1F469-1F3FC-200D-2708-FE0F",
		Value:      "👩🏼‍✈️",
		Descriptor: "Woman Pilot: Medium-Light Skin Tone",
	},
	"1F469-1F3FD": {
		Key:        "1F469-1F3FD",
		Value:      "👩🏽",
		Descriptor: "Woman: Medium Skin Tone",
	},
	"1F469-1F3FD-200D-1F33E": {
		Key:        "1F469-1F3FD-200D-1F33E",
		Value:      "👩🏽‍🌾",
		Descriptor: "Woman Farmer: Medium Skin Tone",
	},
	"1F469-1F3FD-200D-1F373": {
		Key:        "1F469-1F3FD-200D-1F373",
		Value:      "👩🏽‍🍳",
		Descriptor: "Woman Cook: Medium Skin Tone",
	},
	"1F469-1F3FD-200D-1F393": {
		Key:        "1F469-1F3FD-200D-1F393",
		Value:      "👩🏽‍🎓",
		Descriptor: "Woman Student: Medium Skin Tone",
	},
	"1F469-1F3FD-200D-1F3A4": {
		Key:        "1F469-1F3FD-200D-1F3A4",
		Value:      "👩🏽‍🎤",
		Descriptor: "Woman Singer: Medium Skin Tone",
	},
	"1F469-1F3FD-200D-1F3A8": {
		Key:        "1F469-1F3FD-200D-1F3A8",
		Value:      "👩🏽‍🎨",
		Descriptor: "Woman Artist: Medium Skin Tone",
	},
	"1F469-1F3FD-200D-1F3EB": {
		Key:        "1F469-1F3FD-200D-1F3EB",
		Value:      "👩🏽‍🏫",
		Descriptor: "Woman Teacher: Medium Skin Tone",
	},
	"1F469-1F3FD-200D-1F3ED": {
		Key:        "1F469-1F3FD-200D-1F3ED",
		Value:      "👩🏽‍🏭",
		Descriptor: "Woman Factory Worker: Medium Skin Tone",
	},
	"1F469-1F3FD-200D-1F4BB": {
		Key:        "1F469-1F3FD-200D-1F4BB",
		Value:      "👩🏽‍💻",
		Descriptor: "Woman Technologist: Medium Skin Tone",
	},
	"1F469-1F3FD-200D-1F4BC": {
		Key:        "1F469-1F3FD-200D-1F4BC",
		Value:      "👩🏽‍💼",
		Descriptor: "Woman Office Worker: Medium Skin Tone",
	},
	"1F469-1F3FD-200D-1F527": {
		Key:        "1F469-1F3FD-200D-1F527",
		Value:      "👩🏽‍🔧",
		Descriptor: "Woman Mechanic: Medium Skin Tone",
	},
	"1F469-1F3FD-200D-1F52C": {
		Key:        "1F469-1F3FD-200D-1F52C",
		Value:      "👩🏽‍🔬",
		Descriptor: "Woman Scientist: Medium Skin Tone",
	},
	"1F469-1F3FD-200D-1F680": {
		Key:        "1F469-1F3FD-200D-1F680",
		Value:      "👩🏽‍🚀",
		Descriptor: "Woman Astronaut: Medium Skin Tone",
	},
	"1F469-1F3FD-200D-1F692": {
		Key:        "1F469-1F3FD-200D-1F692",
		Value:      "👩🏽‍🚒",
		Descriptor: "Woman Firefighter: Medium Skin Tone",
	},
	"1F469-1F3FD-200D-1F91D-200D-1F468-1F3FB": {
		Key:        "1F469-1F3FD-200D-1F91D-200D-1F468-1F3FB",
		Value:      "👩🏽‍🤝‍👨🏻",
		Descriptor: "Woman and Man Holding Hands: Medium Skin Tone, Light Skin Tone",
	},
	"1F469-1F3FD-200D-1F91D-200D-1F468-1F3FC": {
		Key:        "1F469-1F3FD-200D-1F91D-200D-1F468-1F3FC",
		Value:      "👩🏽‍🤝‍👨🏼",
		Descriptor: "Woman and Man Holding Hands: Medium Skin Tone, Medium-Light Skin Tone",
	},
	"1F469-1F3FD-200D-1F91D-200D-1F468-1F3FE": {
		Key:        "1F469-1F3FD-200D-1F91D-200D-1F468-1F3FE",
		Value:      "👩🏽‍🤝‍👨🏾",
		Descriptor: "Woman and Man Holding Hands: Medium Skin Tone, Medium-Dark Skin Tone",
	},
	"1F469-1F3FD-200D-1F91D-200D-1F468-1F3FF": {
		Key:        "1F469-1F3FD-200D-1F91D-200D-1F468-1F3FF",
		Value:      "👩🏽‍🤝‍👨🏿",
		Descriptor: "Woman and Man Holding Hands: Medium Skin Tone, Dark Skin Tone",
	},
	"1F469-1F3FD-200D-1F91D-200D-1F469-1F3FB": {
		Key:        "1F469-1F3FD-200D-1F91D-200D-1F469-1F3FB",
		Value:      "👩🏽‍🤝‍👩🏻",
		Descriptor: "Women Holding Hands: Medium Skin Tone, Light Skin Tone",
	},
	"1F469-1F3FD-200D-1F91D-200D-1F469-1F3FC": {
		Key:        "1F469-1F3FD-200D-1F91D-200D-1F469-1F3FC",
		Value:      "👩🏽‍🤝‍👩🏼",
		Descriptor: "Women Holding Hands: Medium Skin Tone, Medium-Light Skin Tone",
	},
	"1F469-1F3FD-200D-1F91D-200D-1F469-1F3FE": {
		Key:        "1F469-1F3FD-200D-1F91D-200D-1F469-1F3FE",
		Value:      "👩🏽‍🤝‍👩🏾",
		Descriptor: "Women Holding Hands: Medium Skin Tone, Medium-Dark Skin Tone",
	},
	"1F469-1F3FD-200D-1F91D-200D-1F469-1F3FF": {
		Key:        "1F469-1F3FD-200D-1F91D-200D-1F469-1F3FF",
		Value:      "👩🏽‍🤝‍👩🏿",
		Descriptor: "Women Holding Hands: Medium Skin Tone, Dark Skin Tone",
	},
	"1F469-1F3FD-200D-1F9AF": {
		Key:        "1F469-1F3FD-200D-1F9AF",
		Value:      "👩🏽‍🦯",
		Descriptor: "Woman With Probing Cane: Medium Skin Tone",
	},
	"1F469-1F3FD-200D-1F9B0": {
		Key:        "1F469-1F3FD-200D-1F9B0",
		Value:      "👩🏽‍🦰",
		Descriptor: "Woman: Medium Skin Tone, Red Hair",
	},
	"1F469-1F3FD-200D-1F9B1": {
		Key:        "1F469-1F3FD-200D-1F9B1",
		Value:      "👩🏽‍🦱",
		Descriptor: "Woman: Medium Skin Tone, Curly Hair",
	},
	"1F469-1F3FD-200D-1F9B2": {
		Key:        "1F469-1F3FD-200D-1F9B2",
		Value:      "👩🏽‍🦲",
		Descriptor: "Woman: Medium Skin Tone, Bald",
	},
	"1F469-1F3FD-200D-1F9B3": {
		Key:        "1F469-1F3FD-200D-1F9B3",
		Value:      "👩🏽‍🦳",
		Descriptor: "Woman: Medium Skin Tone, White Hair",
	},
	"1F469-1F3FD-200D-1F9BC": {
		Key:        "1F469-1F3FD-200D-1F9BC",
		Value:      "👩🏽‍🦼",
		Descriptor: "Woman in Motorized Wheelchair: Medium Skin Tone",
	},
	"1F469-1F3FD-200D-1F9BD": {
		Key:        "1F469-1F3FD-200D-1F9BD",
		Value:      "👩🏽‍🦽",
		Descriptor: "Woman in Manual Wheelchair: Medium Skin Tone",
	},
	"1F469-1F3FD-200D-2695-FE0F": {
		Key:        "1F469-1F3FD-200D-2695-FE0F",
		Value:      "👩🏽‍⚕️",
		Descriptor: "Woman Health Worker: Medium Skin Tone",
	},
	"1F469-1F3FD-200D-2696-FE0F": {
		Key:        "1F469-1F3FD-200D-2696-FE0F",
		Value:      "👩🏽‍⚖️",
		Descriptor: "Woman Judge: Medium Skin Tone",
	},
	"1F469-1F3FD-200D-2708-FE0F": {
		Key:        "1F469-1F3FD-200D-2708-FE0F",
		Value:      "👩🏽‍✈️",
		Descriptor: "Woman Pilot: Medium Skin Tone",
	},
	"1F469-1F3FE": {
		Key:        "1F469-1F3FE",
		Value:      "👩🏾",
		Descriptor: "Woman: Medium-Dark Skin Tone",
	},
	"1F469-1F3FE-200D-1F33E": {
		Key:        "1F469-1F3FE-200D-1F33E",
		Value:      "👩🏾‍🌾",
		Descriptor: "Woman Farmer: Medium-Dark Skin Tone",
	},
	"1F469-1F3FE-200D-1F373": {
		Key:        "1F469-1F3FE-200D-1F373",
		Value:      "👩🏾‍🍳",
		Descriptor: "Woman Cook: Medium-Dark Skin Tone",
	},
	"1F469-1F3FE-200D-1F393": {
		Key:        "1F469-1F3FE-200D-1F393",
		Value:      "👩🏾‍🎓",
		Descriptor: "Woman Student: Medium-Dark Skin Tone",
	},
	"1F469-1F3FE-200D-1F3A4": {
		Key:        "1F469-1F3FE-200D-1F3A4",
		Value:      "👩🏾‍🎤",
		Descriptor: "Woman Singer: Medium-Dark Skin Tone",
	},
	"1F469-1F3FE-200D-1F3A8": {
		Key:        "1F469-1F3FE-200D-1F3A8",
		Value:      "👩🏾‍🎨",
		Descriptor: "Woman Artist: Medium-Dark Skin Tone",
	},
	"1F469-1F3FE-200D-1F3EB": {
		Key:        "1F469-1F3FE-200D-1F3EB",
		Value:      "👩🏾‍🏫",
		Descriptor: "Woman Teacher: Medium-Dark Skin Tone",
	},
	"1F469-1F3FE-200D-1F3ED": {
		Key:        "1F469-1F3FE-200D-1F3ED",
		Value:      "👩🏾‍🏭",
		Descriptor: "Woman Factory Worker: Medium-Dark Skin Tone",
	},
	"1F469-1F3FE-200D-1F4BB": {
		Key:        "1F469-1F3FE-200D-1F4BB",
		Value:      "👩🏾‍💻",
		Descriptor: "Woman Technologist: Medium-Dark Skin Tone",
	},
	"1F469-1F3FE-200D-1F4BC": {
		Key:        "1F469-1F3FE-200D-1F4BC",
		Value:      "👩🏾‍💼",
		Descriptor: "Woman Office Worker: Medium-Dark Skin Tone",
	},
	"1F469-1F3FE-200D-1F527": {
		Key:        "1F469-1F3FE-200D-1F527",
		Value:      "👩🏾‍🔧",
		Descriptor: "Woman Mechanic: Medium-Dark Skin Tone",
	},
	"1F469-1F3FE-200D-1F52C": {
		Key:        "1F469-1F3FE-200D-1F52C",
		Value:      "👩🏾‍🔬",
		Descriptor: "Woman Scientist: Medium-Dark Skin Tone",
	},
	"1F469-1F3FE-200D-1F680": {
		Key:        "1F469-1F3FE-200D-1F680",
		Value:      "👩🏾‍🚀",
		Descriptor: "Woman Astronaut: Medium-Dark Skin Tone",
	},
	"1F469-1F3FE-200D-1F692": {
		Key:        "1F469-1F3FE-200D-1F692",
		Value:      "👩🏾‍🚒",
		Descriptor: "Woman Firefighter: Medium-Dark Skin Tone",
	},
	"1F469-1F3FE-200D-1F91D-200D-1F468-1F3FB": {
		Key:        "1F469-1F3FE-200D-1F91D-200D-1F468-1F3FB",
		Value:      "👩🏾‍🤝‍👨🏻",
		Descriptor: "Woman and Man Holding Hands: Medium-Dark Skin Tone, Light Skin Tone",
	},
	"1F469-1F3FE-200D-1F91D-200D-1F468-1F3FC": {
		Key:        "1F469-1F3FE-200D-1F91D-200D-1F468-1F3FC",
		Value:      "👩🏾‍🤝‍👨🏼",
		Descriptor: "Woman and Man Holding Hands: Medium-Dark Skin Tone, Medium-Light Skin Tone",
	},
	"1F469-1F3FE-200D-1F91D-200D-1F468-1F3FD": {
		Key:        "1F469-1F3FE-200D-1F91D-200D-1F468-1F3FD",
		Value:      "👩🏾‍🤝‍👨🏽",
		Descriptor: "Woman and Man Holding Hands: Medium-Dark Skin Tone, Medium Skin Tone",
	},
	"1F469-1F3FE-200D-1F91D-200D-1F468-1F3FF": {
		Key:        "1F469-1F3FE-200D-1F91D-200D-1F468-1F3FF",
		Value:      "👩🏾‍🤝‍👨🏿",
		Descriptor: "Woman and Man Holding Hands: Medium-Dark Skin Tone, Dark Skin Tone",
	},
	"1F469-1F3FE-200D-1F91D-200D-1F469-1F3FB": {
		Key:        "1F469-1F3FE-200D-1F91D-200D-1F469-1F3FB",
		Value:      "👩🏾‍🤝‍👩🏻",
		Descriptor: "Women Holding Hands: Medium-Dark Skin Tone, Light Skin Tone",
	},
	"1F469-1F3FE-200D-1F91D-200D-1F469-1F3FC": {
		Key:        "1F469-1F3FE-200D-1F91D-200D-1F469-1F3FC",
		Value:      "👩🏾‍🤝‍👩🏼",
		Descriptor: "Women Holding Hands: Medium-Dark Skin Tone, Medium-Light Skin Tone",
	},
	"1F469-1F3FE-200D-1F91D-200D-1F469-1F3FD": {
		Key:        "1F469-1F3FE-200D-1F91D-200D-1F469-1F3FD",
		Value:      "👩🏾‍🤝‍👩🏽",
		Descriptor: "Women Holding Hands: Medium-Dark Skin Tone, Medium Skin Tone",
	},
	"1F469-1F3FE-200D-1F91D-200D-1F469-1F3FF": {
		Key:        "1F469-1F3FE-200D-1F91D-200D-1F469-1F3FF",
		Value:      "👩🏾‍🤝‍👩🏿",
		Descriptor: "Women Holding Hands: Medium-Dark Skin Tone, Dark Skin Tone",
	},
	"1F469-1F3FE-200D-1F9AF": {
		Key:        "1F469-1F3FE-200D-1F9AF",
		Value:      "👩🏾‍🦯",
		Descriptor: "Woman With Probing Cane: Medium-Dark Skin Tone",
	},
	"1F469-1F3FE-200D-1F9B0": {
		Key:        "1F469-1F3FE-200D-1F9B0",
		Value:      "👩🏾‍🦰",
		Descriptor: "Woman: Medium-Dark Skin Tone, Red Hair",
	},
	"1F469-1F3FE-200D-1F9B1": {
		Key:        "1F469-1F3FE-200D-1F9B1",
		Value:      "👩🏾‍🦱",
		Descriptor: "Woman: Medium-Dark Skin Tone, Curly Hair",
	},
	"1F469-1F3FE-200D-1F9B2": {
		Key:        "1F469-1F3FE-200D-1F9B2",
		Value:      "👩🏾‍🦲",
		Descriptor: "Woman: Medium-Dark Skin Tone, Bald",
	},
	"1F469-1F3FE-200D-1F9B3": {
		Key:        "1F469-1F3FE-200D-1F9B3",
		Value:      "👩🏾‍🦳",
		Descriptor: "Woman: Medium-Dark Skin Tone, White Hair",
	},
	"1F469-1F3FE-200D-1F9BC": {
		Key:        "1F469-1F3FE-200D-1F9BC",
		Value:      "👩🏾‍🦼",
		Descriptor: "Woman in Motorized Wheelchair: Medium-Dark Skin Tone",
	},
	"1F469-1F3FE-200D-1F9BD": {
		Key:        "1F469-1F3FE-200D-1F9BD",
		Value:      "👩🏾‍🦽",
		Descriptor: "Woman in Manual Wheelchair: Medium-Dark Skin Tone",
	},
	"1F469-1F3FE-200D-2695-FE0F": {
		Key:        "1F469-1F3FE-200D-2695-FE0F",
		Value:      "👩🏾‍⚕️",
		Descriptor: "Woman Health Worker: Medium-Dark Skin Tone",
	},
	"1F469-1F3FE-200D-2696-FE0F": {
		Key:        "1F469-1F3FE-200D-2696-FE0F",
		Value:      "👩🏾‍⚖️",
		Descriptor: "Woman Judge: Medium-Dark Skin Tone",
	},
	"1F469-1F3FE-200D-2708-FE0F": {
		Key:        "1F469-1F3FE-200D-2708-FE0F",
		Value:      "👩🏾‍✈️",
		Descriptor: "Woman Pilot: Medium-Dark Skin Tone",
	},
	"1F469-1F3FF": {
		Key:        "1F469-1F3FF",
		Value:      "👩🏿",
		Descriptor: "Woman: Dark Skin Tone",
	},
	"1F469-1F3FF-200D-1F33E": {
		Key:        "1F469-1F3FF-200D-1F33E",
		Value:      "👩🏿‍🌾",
		Descriptor: "Woman Farmer: Dark Skin Tone",
	},
	"1F469-1F3FF-200D-1F373": {
		Key:        "1F469-1F3FF-200D-1F373",
		Value:      "👩🏿‍🍳",
		Descriptor: "Woman Cook: Dark Skin Tone",
	},
	"1F469-1F3FF-200D-1F393": {
		Key:        "1F469-1F3FF-200D-1F393",
		Value:      "👩🏿‍🎓",
		Descriptor: "Woman Student: Dark Skin Tone",
	},
	"1F469-1F3FF-200D-1F3A4": {
		Key:        "1F469-1F3FF-200D-1F3A4",
		Value:      "👩🏿‍🎤",
		Descriptor: "Woman Singer: Dark Skin Tone",
	},
	"1F469-1F3FF-200D-1F3A8": {
		Key:        "1F469-1F3FF-200D-1F3A8",
		Value:      "👩🏿‍🎨",
		Descriptor: "Woman Artist: Dark Skin Tone",
	},
	"1F469-1F3FF-200D-1F3EB": {
		Key:        "1F469-1F3FF-200D-1F3EB",
		Value:      "👩🏿‍🏫",
		Descriptor: "Woman Teacher: Dark Skin Tone",
	},
	"1F469-1F3FF-200D-1F3ED": {
		Key:        "1F469-1F3FF-200D-1F3ED",
		Value:      "👩🏿‍🏭",
		Descriptor: "Woman Factory Worker: Dark Skin Tone",
	},
	"1F469-1F3FF-200D-1F4BB": {
		Key:        "1F469-1F3FF-200D-1F4BB",
		Value:      "👩🏿‍💻",
		Descriptor: "Woman Technologist: Dark Skin Tone",
	},
	"1F469-1F3FF-200D-1F4BC": {
		Key:        "1F469-1F3FF-200D-1F4BC",
		Value:      "👩🏿‍💼",
		Descriptor: "Woman Office Worker: Dark Skin Tone",
	},
	"1F469-1F3FF-200D-1F527": {
		Key:        "1F469-1F3FF-200D-1F527",
		Value:      "👩🏿‍🔧",
		Descriptor: "Woman Mechanic: Dark Skin Tone",
	},
	"1F469-1F3FF-200D-1F52C": {
		Key:        "1F469-1F3FF-200D-1F52C",
		Value:      "👩🏿‍🔬",
		Descriptor: "Woman Scientist: Dark Skin Tone",
	},
	"1F469-1F3FF-200D-1F680": {
		Key:        "1F469-1F3FF-200D-1F680",
		Value:      "👩🏿‍🚀",
		Descriptor: "Woman Astronaut: Dark Skin Tone",
	},
	"1F469-1F3FF-200D-1F692": {
		Key:        "1F469-1F3FF-200D-1F692",
		Value:      "👩🏿‍🚒",
		Descriptor: "Woman Firefighter: Dark Skin Tone",
	},
	"1F469-1F3FF-200D-1F91D-200D-1F468-1F3FB": {
		Key:        "1F469-1F3FF-200D-1F91D-200D-1F468-1F3FB",
		Value:      "👩🏿‍🤝‍👨🏻",
		Descriptor: "Woman and Man Holding Hands: Dark Skin Tone, Light Skin Tone",
	},
	"1F469-1F3FF-200D-1F91D-200D-1F468-1F3FC": {
		Key:        "1F469-1F3FF-200D-1F91D-200D-1F468-1F3FC",
		Value:      "👩🏿‍🤝‍👨🏼",
		Descriptor: "Woman and Man Holding Hands: Dark Skin Tone, Medium-Light Skin Tone",
	},
	"1F469-1F3FF-200D-1F91D-200D-1F468-1F3FD": {
		Key:        "1F469-1F3FF-200D-1F91D-200D-1F468-1F3FD",
		Value:      "👩🏿‍🤝‍👨🏽",
		Descriptor: "Woman and Man Holding Hands: Dark Skin Tone, Medium Skin Tone",
	},
	"1F469-1F3FF-200D-1F91D-200D-1F468-1F3FE": {
		Key:        "1F469-1F3FF-200D-1F91D-200D-1F468-1F3FE",
		Value:      "👩🏿‍🤝‍👨🏾",
		Descriptor: "Woman and Man Holding Hands: Dark Skin Tone, Medium-Dark Skin Tone",
	},
	"1F469-1F3FF-200D-1F91D-200D-1F469-1F3FB": {
		Key:        "1F469-1F3FF-200D-1F91D-200D-1F469-1F3FB",
		Value:      "👩🏿‍🤝‍👩🏻",
		Descriptor: "Women Holding Hands: Dark Skin Tone, Light Skin Tone",
	},
	"1F469-1F3FF-200D-1F91D-200D-1F469-1F3FC": {
		Key:        "1F469-1F3FF-200D-1F91D-200D-1F469-1F3FC",
		Value:      "👩🏿‍🤝‍👩🏼",
		Descriptor: "Women Holding Hands: Dark Skin Tone, Medium-Light Skin Tone",
	},
	"1F469-1F3FF-200D-1F91D-200D-1F469-1F3FD": {
		Key:        "1F469-1F3FF-200D-1F91D-200D-1F469-1F3FD",
		Value:      "👩🏿‍🤝‍👩🏽",
		Descriptor: "Women Holding Hands: Dark Skin Tone, Medium Skin Tone",
	},
	"1F469-1F3FF-200D-1F91D-200D-1F469-1F3FE": {
		Key:        "1F469-1F3FF-200D-1F91D-200D-1F469-1F3FE",
		Value:      "👩🏿‍🤝‍👩🏾",
		Descriptor: "Women Holding Hands: Dark Skin Tone, Medium-Dark Skin Tone",
	},
	"1F469-1F3FF-200D-1F9AF": {
		Key:        "1F469-1F3FF-200D-1F9AF",
		Value:      "👩🏿‍🦯",
		Descriptor: "Woman With Probing Cane: Dark Skin Tone",
	},
	"1F469-1F3FF-200D-1F9B0": {
		Key:        "1F469-1F3FF-200D-1F9B0",
		Value:      "👩🏿‍🦰",
		Descriptor: "Woman: Dark Skin Tone, Red Hair",
	},
	"1F469-1F3FF-200D-1F9B1": {
		Key:        "1F469-1F3FF-200D-1F9B1",
		Value:      "👩🏿‍🦱",
		Descriptor: "Woman: Dark Skin Tone, Curly Hair",
	},
	"1F469-1F3FF-200D-1F9B2": {
		Key:        "1F469-1F3FF-200D-1F9B2",
		Value:      "👩🏿‍🦲",
		Descriptor: "Woman: Dark Skin Tone, Bald",
	},
	"1F469-1F3FF-200D-1F9B3": {
		Key:        "1F469-1F3FF-200D-1F9B3",
		Value:      "👩🏿‍🦳",
		Descriptor: "Woman: Dark Skin Tone, White Hair",
	},
	"1F469-1F3FF-200D-1F9BC": {
		Key:        "1F469-1F3FF-200D-1F9BC",
		Value:      "👩🏿‍🦼",
		Descriptor: "Woman in Motorized Wheelchair: Dark Skin Tone",
	},
	"1F469-1F3FF-200D-1F9BD": {
		Key:        "1F469-1F3FF-200D-1F9BD",
		Value:      "👩🏿‍🦽",
		Descriptor: "Woman in Manual Wheelchair: Dark Skin Tone",
	},
	"1F469-1F3FF-200D-2695-FE0F": {
		Key:        "1F469-1F3FF-200D-2695-FE0F",
		Value:      "👩🏿‍⚕️",
		Descriptor: "Woman Health Worker: Dark Skin Tone",
	},
	"1F469-1F3FF-200D-2696-FE0F": {
		Key:        "1F469-1F3FF-200D-2696-FE0F",
		Value:      "👩🏿‍⚖️",
		Descriptor: "Woman Judge: Dark Skin Tone",
	},
	"1F469-1F3FF-200D-2708-FE0F": {
		Key:        "1F469-1F3FF-200D-2708-FE0F",
		Value:      "👩🏿‍✈️",
		Descriptor: "Woman Pilot: Dark Skin Tone",
	},
	"1F469-200D-1F33E": {
		Key:        "1F469-200D-1F33E",
		Value:      "👩‍🌾",
		Descriptor: "Woman Farmer",
	},
	"1F469-200D-1F373": {
		Key:        "1F469-200D-1F373",
		Value:      "👩‍🍳",
		Descriptor: "Woman Cook",
	},
	"1F469-200D-1F393": {
		Key:        "1F469-200D-1F393",
		Value:      "👩‍🎓",
		Descriptor: "Woman Student",
	},
	"1F469-200D-1F3A4": {
		Key:        "1F469-200D-1F3A4",
		Value:      "👩‍🎤",
		Descriptor: "Woman Singer",
	},
	"1F469-200D-1F3A8": {
		Key:        "1F469-200D-1F3A8",
		Value:      "👩‍🎨",
		Descriptor: "Woman Artist",
	},
	"1F469-200D-1F3EB": {
		Key:        "1F469-200D-1F3EB",
		Value:      "👩‍🏫",
		Descriptor: "Woman Teacher",
	},
	"1F469-200D-1F3ED": {
		Key:        "1F469-200D-1F3ED",
		Value:      "👩‍🏭",
		Descriptor: "Woman Factory Worker",
	},
	"1F469-200D-1F466": {
		Key:        "1F469-200D-1F466",
		Value:      "👩‍👦",
		Descriptor: "Family: Woman, Boy",
	},
	"1F469-200D-1F466-200D-1F466": {
		Key:        "1F469-200D-1F466-200D-1F466",
		Value:      "👩‍👦‍👦",
		Descriptor: "Family: Woman, Boy, Boy",
	},
	"1F469-200D-1F467": {
		Key:        "1F469-200D-1F467",
		Value:      "👩‍👧",
		Descriptor: "Family: Woman, Girl",
	},
	"1F469-200D-1F467-200D-1F466": {
		Key:        "1F469-200D-1F467-200D-1F466",
		Value:      "👩‍👧‍👦",
		Descriptor: "Family: Woman, Girl, Boy",
	},
	"1F469-200D-1F467-200D-1F467": {
		Key:        "1F469-200D-1F467-200D-1F467",
		Value:      "👩‍👧‍👧",
		Descriptor: "Family: Woman, Girl, Girl",
	},
	"1F469-200D-1F469-200D-1F466": {
		Key:        "1F469-200D-1F469-200D-1F466",
		Value:      "👩‍👩‍👦",
		Descriptor: "Family: Woman, Woman, Boy",
	},
	"1F469-200D-1F469-200D-1F466-200D-1F466": {
		Key:        "1F469-200D-1F469-200D-1F466-200D-1F466",
		Value:      "👩‍👩‍👦‍👦",
		Descriptor: "Family: Woman, Woman, Boy, Boy",
	},
	"1F469-200D-1F469-200D-1F467": {
		Key:        "1F469-200D-1F469-200D-1F467",
		Value:      "👩‍👩‍👧",
		Descriptor: "Family: Woman, Woman, Girl",
	},
	"1F469-200D-1F469-200D-1F467-200D-1F466": {
		Key:        "1F469-200D-1F469-200D-1F467-200D-1F466",
		Value:      "👩‍👩‍👧‍👦",
		Descriptor: "Family: Woman, Woman, Girl, Boy",
	},
	"1F469-200D-1F469-200D-1F467-200D-1F467": {
		Key:        "1F469-200D-1F469-200D-1F467-200D-1F467",
		Value:      "👩‍👩‍👧‍👧",
		Descriptor: "Family: Woman, Woman, Girl, Girl",
	},
	"1F469-200D-1F4BB": {
		Key:        "1F469-200D-1F4BB",
		Value:      "👩‍💻",
		Descriptor: "Woman Technologist",
	},
	"1F469-200D-1F4BC": {
		Key:        "1F469-200D-1F4BC",
		Value:      "👩‍💼",
		Descriptor: "Woman Office Worker",
	},
	"1F469-200D-1F527": {
		Key:        "1F469-200D-1F527",
		Value:      "👩‍🔧",
		Descriptor: "Woman Mechanic",
	},
	"1F469-200D-1F52C": {
		Key:        "1F469-200D-1F52C",
		Value:      "👩‍🔬",
		Descriptor: "Woman Scientist",
	},
	"1F469-200D-1F680": {
		Key:        "1F469-200D-1F680",
		Value:      "👩‍🚀",
		Descriptor: "Woman Astronaut",
	},
	"1F469-200D-1F692": {
		Key:        "1F469-200D-1F692",
		Value:      "👩‍🚒",
		Descriptor: "Woman Firefighter",
	},
	"1F469-200D-1F9AF": {
		Key:        "1F469-200D-1F9AF",
		Value:      "👩‍🦯",
		Descriptor: "Woman With Probing Cane",
	},
	"1F469-200D-1F9B0": {
		Key:        "1F469-200D-1F9B0",
		Value:      "👩‍🦰",
		Descriptor: "Woman: Red Hair",
	},
	"1F469-200D-1F9B1": {
		Key:        "1F469-200D-1F9B1",
		Value:      "👩‍🦱",
		Descriptor: "Woman: Curly Hair",
	},
	"1F469-200D-1F9B2": {
		Key:        "1F469-200D-1F9B2",
		Value:      "👩‍🦲",
		Descriptor: "Woman: Bald",
	},
	"1F469-200D-1F9B3": {
		Key:        "1F469-200D-1F9B3",
		Value:      "👩‍🦳",
		Descriptor: "Woman: White Hair",
	},
	"1F469-200D-1F9BC": {
		Key:        "1F469-200D-1F9BC",
		Value:      "👩‍🦼",
		Descriptor: "Woman in Motorized Wheelchair",
	},
	"1F469-200D-1F9BD": {
		Key:        "1F469-200D-1F9BD",
		Value:      "👩‍🦽",
		Descriptor: "Woman in Manual Wheelchair",
	},
	"1F469-200D-2695-FE0F": {
		Key:        "1F469-200D-2695-FE0F",
		Value:      "👩‍⚕️",
		Descriptor: "Woman Health Worker",
	},
	"1F469-200D-2696-FE0F": {
		Key:        "1F469-200D-2696-FE0F",
		Value:      "👩‍⚖️",
		Descriptor: "Woman Judge",
	},
	"1F469-200D-2708-FE0F": {
		Key:        "1F469-200D-2708-FE0F",
		Value:      "👩‍✈️",
		Descriptor: "Woman Pilot",
	},
	"1F469-200D-2764-FE0F-200D-1F468": {
		Key:        "1F469-200D-2764-FE0F-200D-1F468",
		Value:      "👩‍❤️‍👨",
		Descriptor: "Couple With Heart: Woman, Man",
	},
	"1F469-200D-2764-FE0F-200D-1F469": {
		Key:        "1F469-200D-2764-FE0F-200D-1F469",
		Value:      "👩‍❤️‍👩",
		Descriptor: "Couple With Heart: Woman, Woman",
	},
	"1F469-200D-2764-FE0F-200D-1F48B-200D-1F468": {
		Key:        "1F469-200D-2764-FE0F-200D-1F48B-200D-1F468",
		Value:      "👩‍❤️‍💋‍👨",
		Descriptor: "Kiss: Woman, Man",
	},
	"1F469-200D-2764-FE0F-200D-1F48B-200D-1F469": {
		Key:        "1F469-200D-2764-FE0F-200D-1F48B-200D-1F469",
		Value:      "👩‍❤️‍💋‍👩",
		Descriptor: "Kiss: Woman, Woman",
	},
	"1F46A": {
		Key:        "1F46A",
		Value:      "👪",
		Descriptor: "Family",
	},
	"1F46B": {
		Key:        "1F46B",
		Value:      "👫",
		Descriptor: "Woman and Man Holding Hands",
	},
	"1F46B-1F3FB": {
		Key:        "1F46B-1F3FB",
		Value:      "👫🏻",
		Descriptor: "Woman and Man Holding Hands: Light Skin Tone",
	},
	"1F46B-1F3FC": {
		Key:        "1F46B-1F3FC",
		Value:      "👫🏼",
		Descriptor: "Woman and Man Holding Hands: Medium-Light Skin Tone",
	},
	"1F46B-1F3FD": {
		Key:        "1F46B-1F3FD",
		Value:      "👫🏽",
		Descriptor: "Woman and Man Holding Hands: Medium Skin Tone",
	},
	"1F46B-1F3FE": {
		Key:        "1F46B-1F3FE",
		Value:      "👫🏾",
		Descriptor: "Woman and Man Holding Hands: Medium-Dark Skin Tone",
	},
	"1F46B-1F3FF": {
		Key:        "1F46B-1F3FF",
		Value:      "👫🏿",
		Descriptor: "Woman and Man Holding Hands: Dark Skin Tone",
	},
	"1F46C": {
		Key:        "1F46C",
		Value:      "👬",
		Descriptor: "Men Holding Hands",
	},
	"1F46C-1F3FB": {
		Key:        "1F46C-1F3FB",
		Value:      "👬🏻",
		Descriptor: "Men Holding Hands: Light Skin Tone",
	},
	"1F46C-1F3FC": {
		Key:        "1F46C-1F3FC",
		Value:      "👬🏼",
		Descriptor: "Men Holding Hands: Medium-Light Skin Tone",
	},
	"1F46C-1F3FD": {
		Key:        "1F46C-1F3FD",
		Value:      "👬🏽",
		Descriptor: "Men Holding Hands: Medium Skin Tone",
	},
	"1F46C-1F3FE": {
		Key:        "1F46C-1F3FE",
		Value:      "👬🏾",
		Descriptor: "Men Holding Hands: Medium-Dark Skin Tone",
	},
	"1F46C-1F3FF": {
		Key:        "1F46C-1F3FF",
		Value:      "👬🏿",
		Descriptor: "Men Holding Hands: Dark Skin Tone",
	},
	"1F46D": {
		Key:        "1F46D",
		Value:      "👭",
		Descriptor: "Women Holding Hands",
	},
	"1F46D-1F3FB": {
		Key:        "1F46D-1F3FB",
		Value:      "👭🏻",
		Descriptor: "Women Holding Hands: Light Skin Tone",
	},
	"1F46D-1F3FC": {
		Key:        "1F46D-1F3FC",
		Value:      "👭🏼",
		Descriptor: "Women Holding Hands: Medium-Light Skin Tone",
	},
	"1F46D-1F3FD": {
		Key:        "1F46D-1F3FD",
		Value:      "👭🏽",
		Descriptor: "Women Holding Hands: Medium Skin Tone",
	},
	"1F46D-1F3FE": {
		Key:        "1F46D-1F3FE",
		Value:      "👭🏾",
		Descriptor: "Women Holding Hands: Medium-Dark Skin Tone",
	},
	"1F46D-1F3FF": {
		Key:        "1F46D-1F3FF",
		Value:      "👭🏿",
		Descriptor: "Women Holding Hands: Dark Skin Tone",
	},
	"1F46E": {
		Key:        "1F46E",
		Value:      "👮",
		Descriptor: "Police Officer",
	},
	"1F46E-1F3FB": {
		Key:        "1F46E-1F3FB",
		Value:      "👮🏻",
		Descriptor: "Police Officer: Light Skin Tone",
	},
	"1F46E-1F3FB-200D-2640-FE0F": {
		Key:        "1F46E-1F3FB-200D-2640-FE0F",
		Value:      "👮🏻‍♀️",
		Descriptor: "Woman Police Officer: Light Skin Tone",
	},
	"1F46E-1F3FB-200D-2642-FE0F": {
		Key:        "1F46E-1F3FB-200D-2642-FE0F",
		Value:      "👮🏻‍♂️",
		Descriptor: "Man Police Officer: Light Skin Tone",
	},
	"1F46E-1F3FC": {
		Key:        "1F46E-1F3FC",
		Value:      "👮🏼",
		Descriptor: "Police Officer: Medium-Light Skin Tone",
	},
	"1F46E-1F3FC-200D-2640-FE0F": {
		Key:        "1F46E-1F3FC-200D-2640-FE0F",
		Value:      "👮🏼‍♀️",
		Descriptor: "Woman Police Officer: Medium-Light Skin Tone",
	},
	"1F46E-1F3FC-200D-2642-FE0F": {
		Key:        "1F46E-1F3FC-200D-2642-FE0F",
		Value:      "👮🏼‍♂️",
		Descriptor: "Man Police Officer: Medium-Light Skin Tone",
	},
	"1F46E-1F3FD": {
		Key:        "1F46E-1F3FD",
		Value:      "👮🏽",
		Descriptor: "Police Officer: Medium Skin Tone",
	},
	"1F46E-1F3FD-200D-2640-FE0F": {
		Key:        "1F46E-1F3FD-200D-2640-FE0F",
		Value:      "👮🏽‍♀️",
		Descriptor: "Woman Police Officer: Medium Skin Tone",
	},
	"1F46E-1F3FD-200D-2642-FE0F": {
		Key:        "1F46E-1F3FD-200D-2642-FE0F",
		Value:      "👮🏽‍♂️",
		Descriptor: "Man Police Officer: Medium Skin Tone",
	},
	"1F46E-1F3FE": {
		Key:        "1F46E-1F3FE",
		Value:      "👮🏾",
		Descriptor: "Police Officer: Medium-Dark Skin Tone",
	},
	"1F46E-1F3FE-200D-2640-FE0F": {
		Key:        "1F46E-1F3FE-200D-2640-FE0F",
		Value:      "👮🏾‍♀️",
		Descriptor: "Woman Police Officer: Medium-Dark Skin Tone",
	},
	"1F46E-1F3FE-200D-2642-FE0F": {
		Key:        "1F46E-1F3FE-200D-2642-FE0F",
		Value:      "👮🏾‍♂️",
		Descriptor: "Man Police Officer: Medium-Dark Skin Tone",
	},
	"1F46E-1F3FF": {
		Key:        "1F46E-1F3FF",
		Value:      "👮🏿",
		Descriptor: "Police Officer: Dark Skin Tone",
	},
	"1F46E-1F3FF-200D-2640-FE0F": {
		Key:        "1F46E-1F3FF-200D-2640-FE0F",
		Value:      "👮🏿‍♀️",
		Descriptor: "Woman Police Officer: Dark Skin Tone",
	},
	"1F46E-1F3FF-200D-2642-FE0F": {
		Key:        "1F46E-1F3FF-200D-2642-FE0F",
		Value:      "👮🏿‍♂️",
		Descriptor: "Man Police Officer: Dark Skin Tone",
	},
	"1F46E-200D-2640-FE0F": {
		Key:        "1F46E-200D-2640-FE0F",
		Value:      "👮‍♀️",
		Descriptor: "Woman Police Officer",
	},
	"1F46E-200D-2642-FE0F": {
		Key:        "1F46E-200D-2642-FE0F",
		Value:      "👮‍♂️",
		Descriptor: "Man Police Officer",
	},
	"1F46F": {
		Key:        "1F46F",
		Value:      "👯",
		Descriptor: "People With Bunny Ears",
	},
	"1F46F-200D-2640-FE0F": {
		Key:        "1F46F-200D-2640-FE0F",
		Value:      "👯‍♀️",
		Descriptor: "Women With Bunny Ears",
	},
	"1F46F-200D-2642-FE0F": {
		Key:        "1F46F-200D-2642-FE0F",
		Value:      "👯‍♂️",
		Descriptor: "Men With Bunny Ears",
	},
	"1F470": {
		Key:        "1F470",
		Value:      "👰",
		Descriptor: "Bride With Veil",
	},
	"1F470-1F3FB": {
		Key:        "1F470-1F3FB",
		Value:      "👰🏻",
		Descriptor: "Bride With Veil: Light Skin Tone",
	},
	"1F470-1F3FC": {
		Key:        "1F470-1F3FC",
		Value:      "👰🏼",
		Descriptor: "Bride With Veil: Medium-Light Skin Tone",
	},
	"1F470-1F3FD": {
		Key:        "1F470-1F3FD",
		Value:      "👰🏽",
		Descriptor: "Bride With Veil: Medium Skin Tone",
	},
	"1F470-1F3FE": {
		Key:        "1F470-1F3FE",
		Value:      "👰🏾",
		Descriptor: "Bride With Veil: Medium-Dark Skin Tone",
	},
	"1F470-1F3FF": {
		Key:        "1F470-1F3FF",
		Value:      "👰🏿",
		Descriptor: "Bride With Veil: Dark Skin Tone",
	},
	"1F471": {
		Key:        "1F471",
		Value:      "👱",
		Descriptor: "Person: Blond Hair",
	},
	"1F471-1F3FB": {
		Key:        "1F471-1F3FB",
		Value:      "👱🏻",
		Descriptor: "Person: Light Skin Tone, Blond Hair",
	},
	"1F471-1F3FB-200D-2640-FE0F": {
		Key:        "1F471-1F3FB-200D-2640-FE0F",
		Value:      "👱🏻‍♀️",
		Descriptor: "Woman: Light Skin Tone, Blond Hair",
	},
	"1F471-1F3FB-200D-2642-FE0F": {
		Key:        "1F471-1F3FB-200D-2642-FE0F",
		Value:      "👱🏻‍♂️",
		Descriptor: "Man: Light Skin Tone, Blond Hair",
	},
	"1F471-1F3FC": {
		Key:        "1F471-1F3FC",
		Value:      "👱🏼",
		Descriptor: "Person: Medium-Light Skin Tone, Blond Hair",
	},
	"1F471-1F3FC-200D-2640-FE0F": {
		Key:        "1F471-1F3FC-200D-2640-FE0F",
		Value:      "👱🏼‍♀️",
		Descriptor: "Woman: Medium-Light Skin Tone, Blond Hair",
	},
	"1F471-1F3FC-200D-2642-FE0F": {
		Key:        "1F471-1F3FC-200D-2642-FE0F",
		Value:      "👱🏼‍♂️",
		Descriptor: "Man: Medium-Light Skin Tone, Blond Hair",
	},
	"1F471-1F3FD": {
		Key:        "1F471-1F3FD",
		Value:      "👱🏽",
		Descriptor: "Person: Medium Skin Tone, Blond Hair",
	},
	"1F471-1F3FD-200D-2640-FE0F": {
		Key:        "1F471-1F3FD-200D-2640-FE0F",
		Value:      "👱🏽‍♀️",
		Descriptor: "Woman: Medium Skin Tone, Blond Hair",
	},
	"1F471-1F3FD-200D-2642-FE0F": {
		Key:        "1F471-1F3FD-200D-2642-FE0F",
		Value:      "👱🏽‍♂️",
		Descriptor: "Man: Medium Skin Tone, Blond Hair",
	},
	"1F471-1F3FE": {
		Key:        "1F471-1F3FE",
		Value:      "👱🏾",
		Descriptor: "Person: Medium-Dark Skin Tone, Blond Hair",
	},
	"1F471-1F3FE-200D-2640-FE0F": {
		Key:        "1F471-1F3FE-200D-2640-FE0F",
		Value:      "👱🏾‍♀️",
		Descriptor: "Woman: Medium-Dark Skin Tone, Blond Hair",
	},
	"1F471-1F3FE-200D-2642-FE0F": {
		Key:        "1F471-1F3FE-200D-2642-FE0F",
		Value:      "👱🏾‍♂️",
		Descriptor: "Man: Medium-Dark Skin Tone, Blond Hair",
	},
	"1F471-1F3FF": {
		Key:        "1F471-1F3FF",
		Value:      "👱🏿",
		Descriptor: "Person: Dark Skin Tone, Blond Hair",
	},
	"1F471-1F3FF-200D-2640-FE0F": {
		Key:        "1F471-1F3FF-200D-2640-FE0F",
		Value:      "👱🏿‍♀️",
		Descriptor: "Woman: Dark Skin Tone, Blond Hair",
	},
	"1F471-1F3FF-200D-2642-FE0F": {
		Key:        "1F471-1F3FF-200D-2642-FE0F",
		Value:      "👱🏿‍♂️",
		Descriptor: "Man: Dark Skin Tone, Blond Hair",
	},
	"1F471-200D-2640-FE0F": {
		Key:        "1F471-200D-2640-FE0F",
		Value:      "👱‍♀️",
		Descriptor: "Woman: Blond Hair",
	},
	"1F471-200D-2642-FE0F": {
		Key:        "1F471-200D-2642-FE0F",
		Value:      "👱‍♂️",
		Descriptor: "Man: Blond Hair",
	},
	"1F472": {
		Key:        "1F472",
		Value:      "👲",
		Descriptor: "Man With Skullcap",
	},
	"1F472-1F3FB": {
		Key:        "1F472-1F3FB",
		Value:      "👲🏻",
		Descriptor: "Man With Skullcap: Light Skin Tone",
	},
	"1F472-1F3FC": {
		Key:        "1F472-1F3FC",
		Value:      "👲🏼",
		Descriptor: "Man With Skullcap: Medium-Light Skin Tone",
	},
	"1F472-1F3FD": {
		Key:        "1F472-1F3FD",
		Value:      "👲🏽",
		Descriptor: "Man With Skullcap: Medium Skin Tone",
	},
	"1F472-1F3FE": {
		Key:        "1F472-1F3FE",
		Value:      "👲🏾",
		Descriptor: "Man With Skullcap: Medium-Dark Skin Tone",
	},
	"1F472-1F3FF": {
		Key:        "1F472-1F3FF",
		Value:      "👲🏿",
		Descriptor: "Man With Skullcap: Dark Skin Tone",
	},
	"1F473": {
		Key:        "1F473",
		Value:      "👳",
		Descriptor: "Person Wearing Turban",
	},
	"1F473-1F3FB": {
		Key:        "1F473-1F3FB",
		Value:      "👳🏻",
		Descriptor: "Person Wearing Turban: Light Skin Tone",
	},
	"1F473-1F3FB-200D-2640-FE0F": {
		Key:        "1F473-1F3FB-200D-2640-FE0F",
		Value:      "👳🏻‍♀️",
		Descriptor: "Woman Wearing Turban: Light Skin Tone",
	},
	"1F473-1F3FB-200D-2642-FE0F": {
		Key:        "1F473-1F3FB-200D-2642-FE0F",
		Value:      "👳🏻‍♂️",
		Descriptor: "Man Wearing Turban: Light Skin Tone",
	},
	"1F473-1F3FC": {
		Key:        "1F473-1F3FC",
		Value:      "👳🏼",
		Descriptor: "Person Wearing Turban: Medium-Light Skin Tone",
	},
	"1F473-1F3FC-200D-2640-FE0F": {
		Key:        "1F473-1F3FC-200D-2640-FE0F",
		Value:      "👳🏼‍♀️",
		Descriptor: "Woman Wearing Turban: Medium-Light Skin Tone",
	},
	"1F473-1F3FC-200D-2642-FE0F": {
		Key:        "1F473-1F3FC-200D-2642-FE0F",
		Value:      "👳🏼‍♂️",
		Descriptor: "Man Wearing Turban: Medium-Light Skin Tone",
	},
	"1F473-1F3FD": {
		Key:        "1F473-1F3FD",
		Value:      "👳🏽",
		Descriptor: "Person Wearing Turban: Medium Skin Tone",
	},
	"1F473-1F3FD-200D-2640-FE0F": {
		Key:        "1F473-1F3FD-200D-2640-FE0F",
		Value:      "👳🏽‍♀️",
		Descriptor: "Woman Wearing Turban: Medium Skin Tone",
	},
	"1F473-1F3FD-200D-2642-FE0F": {
		Key:        "1F473-1F3FD-200D-2642-FE0F",
		Value:      "👳🏽‍♂️",
		Descriptor: "Man Wearing Turban: Medium Skin Tone",
	},
	"1F473-1F3FE": {
		Key:        "1F473-1F3FE",
		Value:      "👳🏾",
		Descriptor: "Person Wearing Turban: Medium-Dark Skin Tone",
	},
	"1F473-1F3FE-200D-2640-FE0F": {
		Key:        "1F473-1F3FE-200D-2640-FE0F",
		Value:      "👳🏾‍♀️",
		Descriptor: "Woman Wearing Turban: Medium-Dark Skin Tone",
	},
	"1F473-1F3FE-200D-2642-FE0F": {
		Key:        "1F473-1F3FE-200D-2642-FE0F",
		Value:      "👳🏾‍♂️",
		Descriptor: "Man Wearing Turban: Medium-Dark Skin Tone",
	},
	"1F473-1F3FF": {
		Key:        "1F473-1F3FF",
		Value:      "👳🏿",
		Descriptor: "Person Wearing Turban: Dark Skin Tone",
	},
	"1F473-1F3FF-200D-2640-FE0F": {
		Key:        "1F473-1F3FF-200D-2640-FE0F",
		Value:      "👳🏿‍♀️",
		Descriptor: "Woman Wearing Turban: Dark Skin Tone",
	},
	"1F473-1F3FF-200D-2642-FE0F": {
		Key:        "1F473-1F3FF-200D-2642-FE0F",
		Value:      "👳🏿‍♂️",
		Descriptor: "Man Wearing Turban: Dark Skin Tone",
	},
	"1F473-200D-2640-FE0F": {
		Key:        "1F473-200D-2640-FE0F",
		Value:      "👳‍♀️",
		Descriptor: "Woman Wearing Turban",
	},
	"1F473-200D-2642-FE0F": {
		Key:        "1F473-200D-2642-FE0F",
		Value:      "👳‍♂️",
		Descriptor: "Man Wearing Turban",
	},
	"1F474": {
		Key:        "1F474",
		Value:      "👴",
		Descriptor: "Old Man",
	},
	"1F474-1F3FB": {
		Key:        "1F474-1F3FB",
		Value:      "👴🏻",
		Descriptor: "Old Man: Light Skin Tone",
	},
	"1F474-1F3FC": {
		Key:        "1F474-1F3FC",
		Value:      "👴🏼",
		Descriptor: "Old Man: Medium-Light Skin Tone",
	},
	"1F474-1F3FD": {
		Key:        "1F474-1F3FD",
		Value:      "👴🏽",
		Descriptor: "Old Man: Medium Skin Tone",
	},
	"1F474-1F3FE": {
		Key:        "1F474-1F3FE",
		Value:      "👴🏾",
		Descriptor: "Old Man: Medium-Dark Skin Tone",
	},
	"1F474-1F3FF": {
		Key:        "1F474-1F3FF",
		Value:      "👴🏿",
		Descriptor: "Old Man: Dark Skin Tone",
	},
	"1F475": {
		Key:        "1F475",
		Value:      "👵",
		Descriptor: "Old Woman",
	},
	"1F475-1F3FB": {
		Key:        "1F475-1F3FB",
		Value:      "👵🏻",
		Descriptor: "Old Woman: Light Skin Tone",
	},
	"1F475-1F3FC": {
		Key:        "1F475-1F3FC",
		Value:      "👵🏼",
		Descriptor: "Old Woman: Medium-Light Skin Tone",
	},
	"1F475-1F3FD": {
		Key:        "1F475-1F3FD",
		Value:      "👵🏽",
		Descriptor: "Old Woman: Medium Skin Tone",
	},
	"1F475-1F3FE": {
		Key:        "1F475-1F3FE",
		Value:      "👵🏾",
		Descriptor: "Old Woman: Medium-Dark Skin Tone",
	},
	"1F475-1F3FF": {
		Key:        "1F475-1F3FF",
		Value:      "👵🏿",
		Descriptor: "Old Woman: Dark Skin Tone",
	},
	"1F476": {
		Key:        "1F476",
		Value:      "👶",
		Descriptor: "Baby",
	},
	"1F476-1F3FB": {
		Key:        "1F476-1F3FB",
		Value:      "👶🏻",
		Descriptor: "Baby: Light Skin Tone",
	},
	"1F476-1F3FC": {
		Key:        "1F476-1F3FC",
		Value:      "👶🏼",
		Descriptor: "Baby: Medium-Light Skin Tone",
	},
	"1F476-1F3FD": {
		Key:        "1F476-1F3FD",
		Value:      "👶🏽",
		Descriptor: "Baby: Medium Skin Tone",
	},
	"1F476-1F3FE": {
		Key:        "1F476-1F3FE",
		Value:      "👶🏾",
		Descriptor: "Baby: Medium-Dark Skin Tone",
	},
	"1F476-1F3FF": {
		Key:        "1F476-1F3FF",
		Value:      "👶🏿",
		Descriptor: "Baby: Dark Skin Tone",
	},
	"1F477": {
		Key:        "1F477",
		Value:      "👷",
		Descriptor: "Construction Worker",
	},
	"1F477-1F3FB": {
		Key:        "1F477-1F3FB",
		Value:      "👷🏻",
		Descriptor: "Construction Worker: Light Skin Tone",
	},
	"1F477-1F3FB-200D-2640-FE0F": {
		Key:        "1F477-1F3FB-200D-2640-FE0F",
		Value:      "👷🏻‍♀️",
		Descriptor: "Woman Construction Worker: Light Skin Tone",
	},
	"1F477-1F3FB-200D-2642-FE0F": {
		Key:        "1F477-1F3FB-200D-2642-FE0F",
		Value:      "👷🏻‍♂️",
		Descriptor: "Man Construction Worker: Light Skin Tone",
	},
	"1F477-1F3FC": {
		Key:        "1F477-1F3FC",
		Value:      "👷🏼",
		Descriptor: "Construction Worker: Medium-Light Skin Tone",
	},
	"1F477-1F3FC-200D-2640-FE0F": {
		Key:        "1F477-1F3FC-200D-2640-FE0F",
		Value:      "👷🏼‍♀️",
		Descriptor: "Woman Construction Worker: Medium-Light Skin Tone",
	},
	"1F477-1F3FC-200D-2642-FE0F": {
		Key:        "1F477-1F3FC-200D-2642-FE0F",
		Value:      "👷🏼‍♂️",
		Descriptor: "Man Construction Worker: Medium-Light Skin Tone",
	},
	"1F477-1F3FD": {
		Key:        "1F477-1F3FD",
		Value:      "👷🏽",
		Descriptor: "Construction Worker: Medium Skin Tone",
	},
	"1F477-1F3FD-200D-2640-FE0F": {
		Key:        "1F477-1F3FD-200D-2640-FE0F",
		Value:      "👷🏽‍♀️",
		Descriptor: "Woman Construction Worker: Medium Skin Tone",
	},
	"1F477-1F3FD-200D-2642-FE0F": {
		Key:        "1F477-1F3FD-200D-2642-FE0F",
		Value:      "👷🏽‍♂️",
		Descriptor: "Man Construction Worker: Medium Skin Tone",
	},
	"1F477-1F3FE": {
		Key:        "1F477-1F3FE",
		Value:      "👷🏾",
		Descriptor: "Construction Worker: Medium-Dark Skin Tone",
	},
	"1F477-1F3FE-200D-2640-FE0F": {
		Key:        "1F477-1F3FE-200D-2640-FE0F",
		Value:      "👷🏾‍♀️",
		Descriptor: "Woman Construction Worker: Medium-Dark Skin Tone",
	},
	"1F477-1F3FE-200D-2642-FE0F": {
		Key:        "1F477-1F3FE-200D-2642-FE0F",
		Value:      "👷🏾‍♂️",
		Descriptor: "Man Construction Worker: Medium-Dark Skin Tone",
	},
	"1F477-1F3FF": {
		Key:        "1F477-1F3FF",
		Value:      "👷🏿",
		Descriptor: "Construction Worker: Dark Skin Tone",
	},
	"1F477-1F3FF-200D-2640-FE0F": {
		Key:        "1F477-1F3FF-200D-2640-FE0F",
		Value:      "👷🏿‍♀️",
		Descriptor: "Woman Construction Worker: Dark Skin Tone",
	},
	"1F477-1F3FF-200D-2642-FE0F": {
		Key:        "1F477-1F3FF-200D-2642-FE0F",
		Value:      "👷🏿‍♂️",
		Descriptor: "Man Construction Worker: Dark Skin Tone",
	},
	"1F477-200D-2640-FE0F": {
		Key:        "1F477-200D-2640-FE0F",
		Value:      "👷‍♀️",
		Descriptor: "Woman Construction Worker",
	},
	"1F477-200D-2642-FE0F": {
		Key:        "1F477-200D-2642-FE0F",
		Value:      "👷‍♂️",
		Descriptor: "Man Construction Worker",
	},
	"1F478": {
		Key:        "1F478",
		Value:      "👸",
		Descriptor: "Princess",
	},
	"1F478-1F3FB": {
		Key:        "1F478-1F3FB",
		Value:      "👸🏻",
		Descriptor: "Princess: Light Skin Tone",
	},
	"1F478-1F3FC": {
		Key:        "1F478-1F3FC",
		Value:      "👸🏼",
		Descriptor: "Princess: Medium-Light Skin Tone",
	},
	"1F478-1F3FD": {
		Key:        "1F478-1F3FD",
		Value:      "👸🏽",
		Descriptor: "Princess: Medium Skin Tone",
	},
	"1F478-1F3FE": {
		Key:        "1F478-1F3FE",
		Value:      "👸🏾",
		Descriptor: "Princess: Medium-Dark Skin Tone",
	},
	"1F478-1F3FF": {
		Key:        "1F478-1F3FF",
		Value:      "👸🏿",
		Descriptor: "Princess: Dark Skin Tone",
	},
	"1F479": {
		Key:        "1F479",
		Value:      "👹",
		Descriptor: "Ogre",
	},
	"1F47A": {
		Key:        "1F47A",
		Value:      "👺",
		Descriptor: "Goblin",
	},
	"1F47B": {
		Key:        "1F47B",
		Value:      "👻",
		Descriptor: "Ghost",
	},
	"1F47C": {
		Key:        "1F47C",
		Value:      "👼",
		Descriptor: "Baby Angel",
	},
	"1F47C-1F3FB": {
		Key:        "1F47C-1F3FB",
		Value:      "👼🏻",
		Descriptor: "Baby Angel: Light Skin Tone",
	},
	"1F47C-1F3FC": {
		Key:        "1F47C-1F3FC",
		Value:      "👼🏼",
		Descriptor: "Baby Angel: Medium-Light Skin Tone",
	},
	"1F47C-1F3FD": {
		Key:        "1F47C-1F3FD",
		Value:      "👼🏽",
		Descriptor: "Baby Angel: Medium Skin Tone",
	},
	"1F47C-1F3FE": {
		Key:        "1F47C-1F3FE",
		Value:      "👼🏾",
		Descriptor: "Baby Angel: Medium-Dark Skin Tone",
	},
	"1F47C-1F3FF": {
		Key:        "1F47C-1F3FF",
		Value:      "👼🏿",
		Descriptor: "Baby Angel: Dark Skin Tone",
	},
	"1F47D": {
		Key:        "1F47D",
		Value:      "👽",
		Descriptor: "Alien",
	},
	"1F47E": {
		Key:        "1F47E",
		Value:      "👾",
		Descriptor: "Alien Monster",
	},
	"1F47F": {
		Key:        "1F47F",
		Value:      "👿",
		Descriptor: "Angry Face With Horns",
	},
	"1F480": {
		Key:        "1F480",
		Value:      "💀",
		Descriptor: "Skull",
	},
	"1F481": {
		Key:        "1F481",
		Value:      "💁",
		Descriptor: "Person Tipping Hand",
	},
	"1F481-1F3FB": {
		Key:        "1F481-1F3FB",
		Value:      "💁🏻",
		Descriptor: "Person Tipping Hand: Light Skin Tone",
	},
	"1F481-1F3FB-200D-2640-FE0F": {
		Key:        "1F481-1F3FB-200D-2640-FE0F",
		Value:      "💁🏻‍♀️",
		Descriptor: "Woman Tipping Hand: Light Skin Tone",
	},
	"1F481-1F3FB-200D-2642-FE0F": {
		Key:        "1F481-1F3FB-200D-2642-FE0F",
		Value:      "💁🏻‍♂️",
		Descriptor: "Man Tipping Hand: Light Skin Tone",
	},
	"1F481-1F3FC": {
		Key:        "1F481-1F3FC",
		Value:      "💁🏼",
		Descriptor: "Person Tipping Hand: Medium-Light Skin Tone",
	},
	"1F481-1F3FC-200D-2640-FE0F": {
		Key:        "1F481-1F3FC-200D-2640-FE0F",
		Value:      "💁🏼‍♀️",
		Descriptor: "Woman Tipping Hand: Medium-Light Skin Tone",
	},
	"1F481-1F3FC-200D-2642-FE0F": {
		Key:        "1F481-1F3FC-200D-2642-FE0F",
		Value:      "💁🏼‍♂️",
		Descriptor: "Man Tipping Hand: Medium-Light Skin Tone",
	},
	"1F481-1F3FD": {
		Key:        "1F481-1F3FD",
		Value:      "💁🏽",
		Descriptor: "Person Tipping Hand: Medium Skin Tone",
	},
	"1F481-1F3FD-200D-2640-FE0F": {
		Key:        "1F481-1F3FD-200D-2640-FE0F",
		Value:      "💁🏽‍♀️",
		Descriptor: "Woman Tipping Hand: Medium Skin Tone",
	},
	"1F481-1F3FD-200D-2642-FE0F": {
		Key:        "1F481-1F3FD-200D-2642-FE0F",
		Value:      "💁🏽‍♂️",
		Descriptor: "Man Tipping Hand: Medium Skin Tone",
	},
	"1F481-1F3FE": {
		Key:        "1F481-1F3FE",
		Value:      "💁🏾",
		Descriptor: "Person Tipping Hand: Medium-Dark Skin Tone",
	},
	"1F481-1F3FE-200D-2640-FE0F": {
		Key:        "1F481-1F3FE-200D-2640-FE0F",
		Value:      "💁🏾‍♀️",
		Descriptor: "Woman Tipping Hand: Medium-Dark Skin Tone",
	},
	"1F481-1F3FE-200D-2642-FE0F": {
		Key:        "1F481-1F3FE-200D-2642-FE0F",
		Value:      "💁🏾‍♂️",
		Descriptor: "Man Tipping Hand: Medium-Dark Skin Tone",
	},
	"1F481-1F3FF": {
		Key:        "1F481-1F3FF",
		Value:      "💁🏿",
		Descriptor: "Person Tipping Hand: Dark Skin Tone",
	},
	"1F481-1F3FF-200D-2640-FE0F": {
		Key:        "1F481-1F3FF-200D-2640-FE0F",
		Value:      "💁🏿‍♀️",
		Descriptor: "Woman Tipping Hand: Dark Skin Tone",
	},
	"1F481-1F3FF-200D-2642-FE0F": {
		Key:        "1F481-1F3FF-200D-2642-FE0F",
		Value:      "💁🏿‍♂️",
		Descriptor: "Man Tipping Hand: Dark Skin Tone",
	},
	"1F481-200D-2640-FE0F": {
		Key:        "1F481-200D-2640-FE0F",
		Value:      "💁‍♀️",
		Descriptor: "Woman Tipping Hand",
	},
	"1F481-200D-2642-FE0F": {
		Key:        "1F481-200D-2642-FE0F",
		Value:      "💁‍♂️",
		Descriptor: "Man Tipping Hand",
	},
	"1F482": {
		Key:        "1F482",
		Value:      "💂",
		Descriptor: "Guard",
	},
	"1F482-1F3FB": {
		Key:        "1F482-1F3FB",
		Value:      "💂🏻",
		Descriptor: "Guard: Light Skin Tone",
	},
	"1F482-1F3FB-200D-2640-FE0F": {
		Key:        "1F482-1F3FB-200D-2640-FE0F",
		Value:      "💂🏻‍♀️",
		Descriptor: "Woman Guard: Light Skin Tone",
	},
	"1F482-1F3FB-200D-2642-FE0F": {
		Key:        "1F482-1F3FB-200D-2642-FE0F",
		Value:      "💂🏻‍♂️",
		Descriptor: "Man Guard: Light Skin Tone",
	},
	"1F482-1F3FC": {
		Key:        "1F482-1F3FC",
		Value:      "💂🏼",
		Descriptor: "Guard: Medium-Light Skin Tone",
	},
	"1F482-1F3FC-200D-2640-FE0F": {
		Key:        "1F482-1F3FC-200D-2640-FE0F",
		Value:      "💂🏼‍♀️",
		Descriptor: "Woman Guard: Medium-Light Skin Tone",
	},
	"1F482-1F3FC-200D-2642-FE0F": {
		Key:        "1F482-1F3FC-200D-2642-FE0F",
		Value:      "💂🏼‍♂️",
		Descriptor: "Man Guard: Medium-Light Skin Tone",
	},
	"1F482-1F3FD": {
		Key:        "1F482-1F3FD",
		Value:      "💂🏽",
		Descriptor: "Guard: Medium Skin Tone",
	},
	"1F482-1F3FD-200D-2640-FE0F": {
		Key:        "1F482-1F3FD-200D-2640-FE0F",
		Value:      "💂🏽‍♀️",
		Descriptor: "Woman Guard: Medium Skin Tone",
	},
	"1F482-1F3FD-200D-2642-FE0F": {
		Key:        "1F482-1F3FD-200D-2642-FE0F",
		Value:      "💂🏽‍♂️",
		Descriptor: "Man Guard: Medium Skin Tone",
	},
	"1F482-1F3FE": {
		Key:        "1F482-1F3FE",
		Value:      "💂🏾",
		Descriptor: "Guard: Medium-Dark Skin Tone",
	},
	"1F482-1F3FE-200D-2640-FE0F": {
		Key:        "1F482-1F3FE-200D-2640-FE0F",
		Value:      "💂🏾‍♀️",
		Descriptor: "Woman Guard: Medium-Dark Skin Tone",
	},
	"1F482-1F3FE-200D-2642-FE0F": {
		Key:        "1F482-1F3FE-200D-2642-FE0F",
		Value:      "💂🏾‍♂️",
		Descriptor: "Man Guard: Medium-Dark Skin Tone",
	},
	"1F482-1F3FF": {
		Key:        "1F482-1F3FF",
		Value:      "💂🏿",
		Descriptor: "Guard: Dark Skin Tone",
	},
	"1F482-1F3FF-200D-2640-FE0F": {
		Key:        "1F482-1F3FF-200D-2640-FE0F",
		Value:      "💂🏿‍♀️",
		Descriptor: "Woman Guard: Dark Skin Tone",
	},
	"1F482-1F3FF-200D-2642-FE0F": {
		Key:        "1F482-1F3FF-200D-2642-FE0F",
		Value:      "💂🏿‍♂️",
		Descriptor: "Man Guard: Dark Skin Tone",
	},
	"1F482-200D-2640-FE0F": {
		Key:        "1F482-200D-2640-FE0F",
		Value:      "💂‍♀️",
		Descriptor: "Woman Guard",
	},
	"1F482-200D-2642-FE0F": {
		Key:        "1F482-200D-2642-FE0F",
		Value:      "💂‍♂️",
		Descriptor: "Man Guard",
	},
	"1F483": {
		Key:        "1F483",
		Value:      "💃",
		Descriptor: "Woman Dancing",
	},
	"1F483-1F3FB": {
		Key:        "1F483-1F3FB",
		Value:      "💃🏻",
		Descriptor: "Woman Dancing: Light Skin Tone",
	},
	"1F483-1F3FC": {
		Key:        "1F483-1F3FC",
		Value:      "💃🏼",
		Descriptor: "Woman Dancing: Medium-Light Skin Tone",
	},
	"1F483-1F3FD": {
		Key:        "1F483-1F3FD",
		Value:      "💃🏽",
		Descriptor: "Woman Dancing: Medium Skin Tone",
	},
	"1F483-1F3FE": {
		Key:        "1F483-1F3FE",
		Value:      "💃🏾",
		Descriptor: "Woman Dancing: Medium-Dark Skin Tone",
	},
	"1F483-1F3FF": {
		Key:        "1F483-1F3FF",
		Value:      "💃🏿",
		Descriptor: "Woman Dancing: Dark Skin Tone",
	},
	"1F484": {
		Key:        "1F484",
		Value:      "💄",
		Descriptor: "Lipstick",
	},
	"1F485": {
		Key:        "1F485",
		Value:      "💅",
		Descriptor: "Nail Polish",
	},
	"1F485-1F3FB": {
		Key:        "1F485-1F3FB",
		Value:      "💅🏻",
		Descriptor: "Nail Polish: Light Skin Tone",
	},
	"1F485-1F3FC": {
		Key:        "1F485-1F3FC",
		Value:      "💅🏼",
		Descriptor: "Nail Polish: Medium-Light Skin Tone",
	},
	"1F485-1F3FD": {
		Key:        "1F485-1F3FD",
		Value:      "💅🏽",
		Descriptor: "Nail Polish: Medium Skin Tone",
	},
	"1F485-1F3FE": {
		Key:        "1F485-1F3FE",
		Value:      "💅🏾",
		Descriptor: "Nail Polish: Medium-Dark Skin Tone",
	},
	"1F485-1F3FF": {
		Key:        "1F485-1F3FF",
		Value:      "💅🏿",
		Descriptor: "Nail Polish: Dark Skin Tone",
	},
	"1F486": {
		Key:        "1F486",
		Value:      "💆",
		Descriptor: "Person Getting Massage",
	},
	"1F486-1F3FB": {
		Key:        "1F486-1F3FB",
		Value:      "💆🏻",
		Descriptor: "Person Getting Massage: Light Skin Tone",
	},
	"1F486-1F3FB-200D-2640-FE0F": {
		Key:        "1F486-1F3FB-200D-2640-FE0F",
		Value:      "💆🏻‍♀️",
		Descriptor: "Woman Getting Massage: Light Skin Tone",
	},
	"1F486-1F3FB-200D-2642-FE0F": {
		Key:        "1F486-1F3FB-200D-2642-FE0F",
		Value:      "💆🏻‍♂️",
		Descriptor: "Man Getting Massage: Light Skin Tone",
	},
	"1F486-1F3FC": {
		Key:        "1F486-1F3FC",
		Value:      "💆🏼",
		Descriptor: "Person Getting Massage: Medium-Light Skin Tone",
	},
	"1F486-1F3FC-200D-2640-FE0F": {
		Key:        "1F486-1F3FC-200D-2640-FE0F",
		Value:      "💆🏼‍♀️",
		Descriptor: "Woman Getting Massage: Medium-Light Skin Tone",
	},
	"1F486-1F3FC-200D-2642-FE0F": {
		Key:        "1F486-1F3FC-200D-2642-FE0F",
		Value:      "💆🏼‍♂️",
		Descriptor: "Man Getting Massage: Medium-Light Skin Tone",
	},
	"1F486-1F3FD": {
		Key:        "1F486-1F3FD",
		Value:      "💆🏽",
		Descriptor: "Person Getting Massage: Medium Skin Tone",
	},
	"1F486-1F3FD-200D-2640-FE0F": {
		Key:        "1F486-1F3FD-200D-2640-FE0F",
		Value:      "💆🏽‍♀️",
		Descriptor: "Woman Getting Massage: Medium Skin Tone",
	},
	"1F486-1F3FD-200D-2642-FE0F": {
		Key:        "1F486-1F3FD-200D-2642-FE0F",
		Value:      "💆🏽‍♂️",
		Descriptor: "Man Getting Massage: Medium Skin Tone",
	},
	"1F486-1F3FE": {
		Key:        "1F486-1F3FE",
		Value:      "💆🏾",
		Descriptor: "Person Getting Massage: Medium-Dark Skin Tone",
	},
	"1F486-1F3FE-200D-2640-FE0F": {
		Key:        "1F486-1F3FE-200D-2640-FE0F",
		Value:      "💆🏾‍♀️",
		Descriptor: "Woman Getting Massage: Medium-Dark Skin Tone",
	},
	"1F486-1F3FE-200D-2642-FE0F": {
		Key:        "1F486-1F3FE-200D-2642-FE0F",
		Value:      "💆🏾‍♂️",
		Descriptor: "Man Getting Massage: Medium-Dark Skin Tone",
	},
	"1F486-1F3FF": {
		Key:        "1F486-1F3FF",
		Value:      "💆🏿",
		Descriptor: "Person Getting Massage: Dark Skin Tone",
	},
	"1F486-1F3FF-200D-2640-FE0F": {
		Key:        "1F486-1F3FF-200D-2640-FE0F",
		Value:      "💆🏿‍♀️",
		Descriptor: "Woman Getting Massage: Dark Skin Tone",
	},
	"1F486-1F3FF-200D-2642-FE0F": {
		Key:        "1F486-1F3FF-200D-2642-FE0F",
		Value:      "💆🏿‍♂️",
		Descriptor: "Man Getting Massage: Dark Skin Tone",
	},
	"1F486-200D-2640-FE0F": {
		Key:        "1F486-200D-2640-FE0F",
		Value:      "💆‍♀️",
		Descriptor: "Woman Getting Massage",
	},
	"1F486-200D-2642-FE0F": {
		Key:        "1F486-200D-2642-FE0F",
		Value:      "💆‍♂️",
		Descriptor: "Man Getting Massage",
	},
	"1F487": {
		Key:        "1F487",
		Value:      "💇",
		Descriptor: "Person Getting Haircut",
	},
	"1F487-1F3FB": {
		Key:        "1F487-1F3FB",
		Value:      "💇🏻",
		Descriptor: "Person Getting Haircut: Light Skin Tone",
	},
	"1F487-1F3FB-200D-2640-FE0F": {
		Key:        "1F487-1F3FB-200D-2640-FE0F",
		Value:      "💇🏻‍♀️",
		Descriptor: "Woman Getting Haircut: Light Skin Tone",
	},
	"1F487-1F3FB-200D-2642-FE0F": {
		Key:        "1F487-1F3FB-200D-2642-FE0F",
		Value:      "💇🏻‍♂️",
		Descriptor: "Man Getting Haircut: Light Skin Tone",
	},
	"1F487-1F3FC": {
		Key:        "1F487-1F3FC",
		Value:      "💇🏼",
		Descriptor: "Person Getting Haircut: Medium-Light Skin Tone",
	},
	"1F487-1F3FC-200D-2640-FE0F": {
		Key:        "1F487-1F3FC-200D-2640-FE0F",
		Value:      "💇🏼‍♀️",
		Descriptor: "Woman Getting Haircut: Medium-Light Skin Tone",
	},
	"1F487-1F3FC-200D-2642-FE0F": {
		Key:        "1F487-1F3FC-200D-2642-FE0F",
		Value:      "💇🏼‍♂️",
		Descriptor: "Man Getting Haircut: Medium-Light Skin Tone",
	},
	"1F487-1F3FD": {
		Key:        "1F487-1F3FD",
		Value:      "💇🏽",
		Descriptor: "Person Getting Haircut: Medium Skin Tone",
	},
	"1F487-1F3FD-200D-2640-FE0F": {
		Key:        "1F487-1F3FD-200D-2640-FE0F",
		Value:      "💇🏽‍♀️",
		Descriptor: "Woman Getting Haircut: Medium Skin Tone",
	},
	"1F487-1F3FD-200D-2642-FE0F": {
		Key:        "1F487-1F3FD-200D-2642-FE0F",
		Value:      "💇🏽‍♂️",
		Descriptor: "Man Getting Haircut: Medium Skin Tone",
	},
	"1F487-1F3FE": {
		Key:        "1F487-1F3FE",
		Value:      "💇🏾",
		Descriptor: "Person Getting Haircut: Medium-Dark Skin Tone",
	},
	"1F487-1F3FE-200D-2640-FE0F": {
		Key:        "1F487-1F3FE-200D-2640-FE0F",
		Value:      "💇🏾‍♀️",
		Descriptor: "Woman Getting Haircut: Medium-Dark Skin Tone",
	},
	"1F487-1F3FE-200D-2642-FE0F": {
		Key:        "1F487-1F3FE-200D-2642-FE0F",
		Value:      "💇🏾‍♂️",
		Descriptor: "Man Getting Haircut: Medium-Dark Skin Tone",
	},
	"1F487-1F3FF": {
		Key:        "1F487-1F3FF",
		Value:      "💇🏿",
		Descriptor: "Person Getting Haircut: Dark Skin Tone",
	},
	"1F487-1F3FF-200D-2640-FE0F": {
		Key:        "1F487-1F3FF-200D-2640-FE0F",
		Value:      "💇🏿‍♀️",
		Descriptor: "Woman Getting Haircut: Dark Skin Tone",
	},
	"1F487-1F3FF-200D-2642-FE0F": {
		Key:        "1F487-1F3FF-200D-2642-FE0F",
		Value:      "💇🏿‍♂️",
		Descriptor: "Man Getting Haircut: Dark Skin Tone",
	},
	"1F487-200D-2640-FE0F": {
		Key:        "1F487-200D-2640-FE0F",
		Value:      "💇‍♀️",
		Descriptor: "Woman Getting Haircut",
	},
	"1F487-200D-2642-FE0F": {
		Key:        "1F487-200D-2642-FE0F",
		Value:      "💇‍♂️",
		Descriptor: "Man Getting Haircut",
	},
	"1F488": {
		Key:        "1F488",
		Value:      "💈",
		Descriptor: "Barber Pole",
	},
	"1F489": {
		Key:        "1F489",
		Value:      "💉",
		Descriptor: "Syringe",
	},
	"1F48A": {
		Key:        "1F48A",
		Value:      "💊",
		Descriptor: "Pill",
	},
	"1F48B": {
		Key:        "1F48B",
		Value:      "💋",
		Descriptor: "Kiss Mark",
	},
	"1F48C": {
		Key:        "1F48C",
		Value:      "💌",
		Descriptor: "Love Letter",
	},
	"1F48D": {
		Key:        "1F48D",
		Value:      "💍",
		Descriptor: "Ring",
	},
	"1F48E": {
		Key:        "1F48E",
		Value:      "💎",
		Descriptor: "Gem Stone",
	},
	"1F48F": {
		Key:        "1F48F",
		Value:      "💏",
		Descriptor: "Kiss",
	},
	"1F490": {
		Key:        "1F490",
		Value:      "💐",
		Descriptor: "Bouquet",
	},
	"1F491": {
		Key:        "1F491",
		Value:      "💑",
		Descriptor: "Couple With Heart",
	},
	"1F492": {
		Key:        "1F492",
		Value:      "💒",
		Descriptor: "Wedding",
	},
	"1F493": {
		Key:        "1F493",
		Value:      "💓",
		Descriptor: "Beating Heart",
	},
	"1F494": {
		Key:        "1F494",
		Value:      "💔",
		Descriptor: "Broken Heart",
	},
	"1F495": {
		Key:        "1F495",
		Value:      "💕",
		Descriptor: "Two Hearts",
	},
	"1F496": {
		Key:        "1F496",
		Value:      "💖",
		Descriptor: "Sparkling Heart",
	},
	"1F497": {
		Key:        "1F497",
		Value:      "💗",
		Descriptor: "Growing Heart",
	},
	"1F498": {
		Key:        "1F498",
		Value:      "💘",
		Descriptor: "Heart With Arrow",
	},
	"1F499": {
		Key:        "1F499",
		Value:      "💙",
		Descriptor: "Blue Heart",
	},
	"1F49A": {
		Key:        "1F49A",
		Value:      "💚",
		Descriptor: "Green Heart",
	},
	"1F49B": {
		Key:        "1F49B",
		Value:      "💛",
		Descriptor: "Yellow Heart",
	},
	"1F49C": {
		Key:        "1F49C",
		Value:      "💜",
		Descriptor: "Purple Heart",
	},
	"1F49D": {
		Key:        "1F49D",
		Value:      "💝",
		Descriptor: "Heart With Ribbon",
	},
	"1F49E": {
		Key:        "1F49E",
		Value:      "💞",
		Descriptor: "Revolving Hearts",
	},
	"1F49F": {
		Key:        "1F49F",
		Value:      "💟",
		Descriptor: "Heart Decoration",
	},
	"1F4A0": {
		Key:        "1F4A0",
		Value:      "💠",
		Descriptor: "Diamond With a Dot",
	},
	"1F4A1": {
		Key:        "1F4A1",
		Value:      "💡",
		Descriptor: "Light Bulb",
	},
	"1F4A2": {
		Key:        "1F4A2",
		Value:      "💢",
		Descriptor: "Anger Symbol",
	},
	"1F4A3": {
		Key:        "1F4A3",
		Value:      "💣",
		Descriptor: "Bomb",
	},
	"1F4A4": {
		Key:        "1F4A4",
		Value:      "💤",
		Descriptor: "Zzz",
	},
	"1F4A5": {
		Key:        "1F4A5",
		Value:      "💥",
		Descriptor: "Collision",
	},
	"1F4A6": {
		Key:        "1F4A6",
		Value:      "💦",
		Descriptor: "Sweat Droplets",
	},
	"1F4A7": {
		Key:        "1F4A7",
		Value:      "💧",
		Descriptor: "Droplet",
	},
	"1F4A8": {
		Key:        "1F4A8",
		Value:      "💨",
		Descriptor: "Dashing Away",
	},
	"1F4A9": {
		Key:        "1F4A9",
		Value:      "💩",
		Descriptor: "Pile of Poo",
	},
	"1F4AA": {
		Key:        "1F4AA",
		Value:      "💪",
		Descriptor: "Flexed Biceps",
	},
	"1F4AA-1F3FB": {
		Key:        "1F4AA-1F3FB",
		Value:      "💪🏻",
		Descriptor: "Flexed Biceps: Light Skin Tone",
	},
	"1F4AA-1F3FC": {
		Key:        "1F4AA-1F3FC",
		Value:      "💪🏼",
		Descriptor: "Flexed Biceps: Medium-Light Skin Tone",
	},
	"1F4AA-1F3FD": {
		Key:        "1F4AA-1F3FD",
		Value:      "💪🏽",
		Descriptor: "Flexed Biceps: Medium Skin Tone",
	},
	"1F4AA-1F3FE": {
		Key:        "1F4AA-1F3FE",
		Value:      "💪🏾",
		Descriptor: "Flexed Biceps: Medium-Dark Skin Tone",
	},
	"1F4AA-1F3FF": {
		Key:        "1F4AA-1F3FF",
		Value:      "💪🏿",
		Descriptor: "Flexed Biceps: Dark Skin Tone",
	},
	"1F4AB": {
		Key:        "1F4AB",
		Value:      "💫",
		Descriptor: "Dizzy",
	},
	"1F4AC": {
		Key:        "1F4AC",
		Value:      "💬",
		Descriptor: "Speech Balloon",
	},
	"1F4AD": {
		Key:        "1F4AD",
		Value:      "💭",
		Descriptor: "Thought Balloon",
	},
	"1F4AE": {
		Key:        "1F4AE",
		Value:      "💮",
		Descriptor: "White Flower",
	},
	"1F4AF": {
		Key:        "1F4AF",
		Value:      "💯",
		Descriptor: "Hundred Points",
	},
	"1F4B0": {
		Key:        "1F4B0",
		Value:      "💰",
		Descriptor: "Money Bag",
	},
	"1F4B1": {
		Key:        "1F4B1",
		Value:      "💱",
		Descriptor: "Currency Exchange",
	},
	"1F4B2": {
		Key:        "1F4B2",
		Value:      "💲",
		Descriptor: "Heavy Dollar Sign",
	},
	"1F4B3": {
		Key:        "1F4B3",
		Value:      "💳",
		Descriptor: "Credit Card",
	},
	"1F4B4": {
		Key:        "1F4B4",
		Value:      "💴",
		Descriptor: "Yen Banknote",
	},
	"1F4B5": {
		Key:        "1F4B5",
		Value:      "💵",
		Descriptor: "Dollar Banknote",
	},
	"1F4B6": {
		Key:        "1F4B6",
		Value:      "💶",
		Descriptor: "Euro Banknote",
	},
	"1F4B7": {
		Key:        "1F4B7",
		Value:      "💷",
		Descriptor: "Pound Banknote",
	},
	"1F4B8": {
		Key:        "1F4B8",
		Value:      "💸",
		Descriptor: "Money With Wings",
	},
	"1F4B9": {
		Key:        "1F4B9",
		Value:      "💹",
		Descriptor: "Chart Increasing With Yen",
	},
	"1F4BA": {
		Key:        "1F4BA",
		Value:      "💺",
		Descriptor: "Seat",
	},
	"1F4BB": {
		Key:        "1F4BB",
		Value:      "💻",
		Descriptor: "Laptop",
	},
	"1F4BC": {
		Key:        "1F4BC",
		Value:      "💼",
		Descriptor: "Briefcase",
	},
	"1F4BD": {
		Key:        "1F4BD",
		Value:      "💽",
		Descriptor: "Computer Disk",
	},
	"1F4BE": {
		Key:        "1F4BE",
		Value:      "💾",
		Descriptor: "Floppy Disk",
	},
	"1F4BF": {
		Key:        "1F4BF",
		Value:      "💿",
		Descriptor: "Optical Disk",
	},
	"1F4C0": {
		Key:        "1F4C0",
		Value:      "📀",
		Descriptor: "DVD",
	},
	"1F4C1": {
		Key:        "1F4C1",
		Value:      "📁",
		Descriptor: "File Folder",
	},
	"1F4C2": {
		Key:        "1F4C2",
		Value:      "📂",
		Descriptor: "Open File Folder",
	},
	"1F4C3": {
		Key:        "1F4C3",
		Value:      "📃",
		Descriptor: "Page With Curl",
	},
	"1F4C4": {
		Key:        "1F4C4",
		Value:      "📄",
		Descriptor: "Page Facing Up",
	},
	"1F4C5": {
		Key:        "1F4C5",
		Value:      "📅",
		Descriptor: "Calendar",
	},
	"1F4C6": {
		Key:        "1F4C6",
		Value:      "📆",
		Descriptor: "Tear-Off Calendar",
	},
	"1F4C7": {
		Key:        "1F4C7",
		Value:      "📇",
		Descriptor: "Card Index",
	},
	"1F4C8": {
		Key:        "1F4C8",
		Value:      "📈",
		Descriptor: "Chart Increasing",
	},
	"1F4C9": {
		Key:        "1F4C9",
		Value:      "📉",
		Descriptor: "Chart Decreasing",
	},
	"1F4CA": {
		Key:        "1F4CA",
		Value:      "📊",
		Descriptor: "Bar Chart",
	},
	"1F4CB": {
		Key:        "1F4CB",
		Value:      "📋",
		Descriptor: "Clipboard",
	},
	"1F4CC": {
		Key:        "1F4CC",
		Value:      "📌",
		Descriptor: "Pushpin",
	},
	"1F4CD": {
		Key:        "1F4CD",
		Value:      "📍",
		Descriptor: "Round Pushpin",
	},
	"1F4CE": {
		Key:        "1F4CE",
		Value:      "📎",
		Descriptor: "Paperclip",
	},
	"1F4CF": {
		Key:        "1F4CF",
		Value:      "📏",
		Descriptor: "Straight Ruler",
	},
	"1F4D0": {
		Key:        "1F4D0",
		Value:      "📐",
		Descriptor: "Triangular Ruler",
	},
	"1F4D1": {
		Key:        "1F4D1",
		Value:      "📑",
		Descriptor: "Bookmark Tabs",
	},
	"1F4D2": {
		Key:        "1F4D2",
		Value:      "📒",
		Descriptor: "Ledger",
	},
	"1F4D3": {
		Key:        "1F4D3",
		Value:      "📓",
		Descriptor: "Notebook",
	},
	"1F4D4": {
		Key:        "1F4D4",
		Value:      "📔",
		Descriptor: "Notebook With Decorative Cover",
	},
	"1F4D5": {
		Key:        "1F4D5",
		Value:      "📕",
		Descriptor: "Closed Book",
	},
	"1F4D6": {
		Key:        "1F4D6",
		Value:      "📖",
		Descriptor: "Open Book",
	},
	"1F4D7": {
		Key:        "1F4D7",
		Value:      "📗",
		Descriptor: "Green Book",
	},
	"1F4D8": {
		Key:        "1F4D8",
		Value:      "📘",
		Descriptor: "Blue Book",
	},
	"1F4D9": {
		Key:        "1F4D9",
		Value:      "📙",
		Descriptor: "Orange Book",
	},
	"1F4DA": {
		Key:        "1F4DA",
		Value:      "📚",
		Descriptor: "Books",
	},
	"1F4DB": {
		Key:        "1F4DB",
		Value:      "📛",
		Descriptor: "Name Badge",
	},
	"1F4DC": {
		Key:        "1F4DC",
		Value:      "📜",
		Descriptor: "Scroll",
	},
	"1F4DD": {
		Key:        "1F4DD",
		Value:      "📝",
		Descriptor: "Memo",
	},
	"1F4DE": {
		Key:        "1F4DE",
		Value:      "📞",
		Descriptor: "Telephone Receiver",
	},
	"1F4DF": {
		Key:        "1F4DF",
		Value:      "📟",
		Descriptor: "Pager",
	},
	"1F4E0": {
		Key:        "1F4E0",
		Value:      "📠",
		Descriptor: "Fax Machine",
	},
	"1F4E1": {
		Key:        "1F4E1",
		Value:      "📡",
		Descriptor: "Satellite Antenna",
	},
	"1F4E2": {
		Key:        "1F4E2",
		Value:      "📢",
		Descriptor: "Loudspeaker",
	},
	"1F4E3": {
		Key:        "1F4E3",
		Value:      "📣",
		Descriptor: "Megaphone",
	},
	"1F4E4": {
		Key:        "1F4E4",
		Value:      "📤",
		Descriptor: "Outbox Tray",
	},
	"1F4E5": {
		Key:        "1F4E5",
		Value:      "📥",
		Descriptor: "Inbox Tray",
	},
	"1F4E6": {
		Key:        "1F4E6",
		Value:      "📦",
		Descriptor: "Package",
	},
	"1F4E7": {
		Key:        "1F4E7",
		Value:      "📧",
		Descriptor: "E-Mail",
	},
	"1F4E8": {
		Key:        "1F4E8",
		Value:      "📨",
		Descriptor: "Incoming Envelope",
	},
	"1F4E9": {
		Key:        "1F4E9",
		Value:      "📩",
		Descriptor: "Envelope With Arrow",
	},
	"1F4EA": {
		Key:        "1F4EA",
		Value:      "📪",
		Descriptor: "Closed Mailbox With Lowered Flag",
	},
	"1F4EB": {
		Key:        "1F4EB",
		Value:      "📫",
		Descriptor: "Closed Mailbox With Raised Flag",
	},
	"1F4EC": {
		Key:        "1F4EC",
		Value:      "📬",
		Descriptor: "Open Mailbox With Raised Flag",
	},
	"1F4ED": {
		Key:        "1F4ED",
		Value:      "📭",
		Descriptor: "Open Mailbox With Lowered Flag",
	},
	"1F4EE": {
		Key:        "1F4EE",
		Value:      "📮",
		Descriptor: "Postbox",
	},
	"1F4EF": {
		Key:        "1F4EF",
		Value:      "📯",
		Descriptor: "Postal Horn",
	},
	"1F4F0": {
		Key:        "1F4F0",
		Value:      "📰",
		Descriptor: "Newspaper",
	},
	"1F4F1": {
		Key:        "1F4F1",
		Value:      "📱",
		Descriptor: "Mobile Phone",
	},
	"1F4F2": {
		Key:        "1F4F2",
		Value:      "📲",
		Descriptor: "Mobile Phone With Arrow",
	},
	"1F4F3": {
		Key:        "1F4F3",
		Value:      "📳",
		Descriptor: "Vibration Mode",
	},
	"1F4F4": {
		Key:        "1F4F4",
		Value:      "📴",
		Descriptor: "Mobile Phone Off",
	},
	"1F4F5": {
		Key:        "1F4F5",
		Value:      "📵",
		Descriptor: "No Mobile Phones",
	},
	"1F4F6": {
		Key:        "1F4F6",
		Value:      "📶",
		Descriptor: "Antenna Bars",
	},
	"1F4F7": {
		Key:        "1F4F7",
		Value:      "📷",
		Descriptor: "Camera",
	},
	"1F4F8": {
		Key:        "1F4F8",
		Value:      "📸",
		Descriptor: "Camera With Flash",
	},
	"1F4F9": {
		Key:        "1F4F9",
		Value:      "📹",
		Descriptor: "Video Camera",
	},
	"1F4FA": {
		Key:        "1F4FA",
		Value:      "📺",
		Descriptor: "Television",
	},
	"1F4FB": {
		Key:        "1F4FB",
		Value:      "📻",
		Descriptor: "Radio",
	},
	"1F4FC": {
		Key:        "1F4FC",
		Value:      "📼",
		Descriptor: "Videocassette",
	},
	"1F4FD-FE0F": {
		Key:        "1F4FD-FE0F",
		Value:      "📽️",
		Descriptor: "Film Projector",
	},
	"1F4FF": {
		Key:        "1F4FF",
		Value:      "📿",
		Descriptor: "Prayer Beads",
	},
	"1F500": {
		Key:        "1F500",
		Value:      "🔀",
		Descriptor: "Shuffle Tracks Button",
	},
	"1F501": {
		Key:        "1F501",
		Value:      "🔁",
		Descriptor: "Repeat Button",
	},
	"1F502": {
		Key:        "1F502",
		Value:      "🔂",
		Descriptor: "Repeat Single Button",
	},
	"1F503": {
		Key:        "1F503",
		Value:      "🔃",
		Descriptor: "Clockwise Vertical Arrows",
	},
	"1F504": {
		Key:        "1F504",
		Value:      "🔄",
		Descriptor: "Counterclockwise Arrows Button",
	},
	"1F505": {
		Key:        "1F505",
		Value:      "🔅",
		Descriptor: "Dim Button",
	},
	"1F506": {
		Key:        "1F506",
		Value:      "🔆",
		Descriptor: "Bright Button",
	},
	"1F507": {
		Key:        "1F507",
		Value:      "🔇",
		Descriptor: "Muted Speaker",
	},
	"1F508": {
		Key:        "1F508",
		Value:      "🔈",
		Descriptor: "Speaker Low Volume",
	},
	"1F509": {
		Key:        "1F509",
		Value:      "🔉",
		Descriptor: "Speaker Medium Volume",
	},
	"1F50A": {
		Key:        "1F50A",
		Value:      "🔊",
		Descriptor: "Speaker High Volume",
	},
	"1F50B": {
		Key:        "1F50B",
		Value:      "🔋",
		Descriptor: "Battery",
	},
	"1F50C": {
		Key:        "1F50C",
		Value:      "🔌",
		Descriptor: "Electric Plug",
	},
	"1F50D": {
		Key:        "1F50D",
		Value:      "🔍",
		Descriptor: "Magnifying Glass Tilted Left",
	},
	"1F50E": {
		Key:        "1F50E",
		Value:      "🔎",
		Descriptor: "Magnifying Glass Tilted Right",
	},
	"1F50F": {
		Key:        "1F50F",
		Value:      "🔏",
		Descriptor: "Locked With Pen",
	},
	"1F510": {
		Key:        "1F510",
		Value:      "🔐",
		Descriptor: "Locked With Key",
	},
	"1F511": {
		Key:        "1F511",
		Value:      "🔑",
		Descriptor: "Key",
	},
	"1F512": {
		Key:        "1F512",
		Value:      "🔒",
		Descriptor: "Locked",
	},
	"1F513": {
		Key:        "1F513",
		Value:      "🔓",
		Descriptor: "Unlocked",
	},
	"1F514": {
		Key:        "1F514",
		Value:      "🔔",
		Descriptor: "Bell",
	},
	"1F515": {
		Key:        "1F515",
		Value:      "🔕",
		Descriptor: "Bell With Slash",
	},
	"1F516": {
		Key:        "1F516",
		Value:      "🔖",
		Descriptor: "Bookmark",
	},
	"1F517": {
		Key:        "1F517",
		Value:      "🔗",
		Descriptor: "Link",
	},
	"1F518": {
		Key:        "1F518",
		Value:      "🔘",
		Descriptor: "Radio Button",
	},
	"1F519": {
		Key:        "1F519",
		Value:      "🔙",
		Descriptor: "Back Arrow",
	},
	"1F51A": {
		Key:        "1F51A",
		Value:      "🔚",
		Descriptor: "End Arrow",
	},
	"1F51B": {
		Key:        "1F51B",
		Value:      "🔛",
		Descriptor: "On! Arrow",
	},
	"1F51C": {
		Key:        "1F51C",
		Value:      "🔜",
		Descriptor: "Soon Arrow",
	},
	"1F51D": {
		Key:        "1F51D",
		Value:      "🔝",
		Descriptor: "Top Arrow",
	},
	"1F51E": {
		Key:        "1F51E",
		Value:      "🔞",
		Descriptor: "No One Under Eighteen",
	},
	"1F51F": {
		Key:        "1F51F",
		Value:      "🔟",
		Descriptor: "Keycap: 10",
	},
	"1F520": {
		Key:        "1F520",
		Value:      "🔠",
		Descriptor: "Input Latin Uppercase",
	},
	"1F521": {
		Key:        "1F521",
		Value:      "🔡",
		Descriptor: "Input Latin Lowercase",
	},
	"1F522": {
		Key:        "1F522",
		Value:      "🔢",
		Descriptor: "Input Numbers",
	},
	"1F523": {
		Key:        "1F523",
		Value:      "🔣",
		Descriptor: "Input Symbols",
	},
	"1F524": {
		Key:        "1F524",
		Value:      "🔤",
		Descriptor: "Input Latin Letters",
	},
	"1F525": {
		Key:        "1F525",
		Value:      "🔥",
		Descriptor: "Fire",
	},
	"1F526": {
		Key:        "1F526",
		Value:      "🔦",
		Descriptor: "Flashlight",
	},
	"1F527": {
		Key:        "1F527",
		Value:      "🔧",
		Descriptor: "Wrench",
	},
	"1F528": {
		Key:        "1F528",
		Value:      "🔨",
		Descriptor: "Hammer",
	},
	"1F529": {
		Key:        "1F529",
		Value:      "🔩",
		Descriptor: "Nut and Bolt",
	},
	"1F52A": {
		Key:        "1F52A",
		Value:      "🔪",
		Descriptor: "Kitchen Knife",
	},
	"1F52B": {
		Key:        "1F52B",
		Value:      "🔫",
		Descriptor: "Pistol",
	},
	"1F52C": {
		Key:        "1F52C",
		Value:      "🔬",
		Descriptor: "Microscope",
	},
	"1F52D": {
		Key:        "1F52D",
		Value:      "🔭",
		Descriptor: "Telescope",
	},
	"1F52E": {
		Key:        "1F52E",
		Value:      "🔮",
		Descriptor: "Crystal Ball",
	},
	"1F52F": {
		Key:        "1F52F",
		Value:      "🔯",
		Descriptor: "Dotted Six-Pointed Star",
	},
	"1F530": {
		Key:        "1F530",
		Value:      "🔰",
		Descriptor: "Japanese Symbol for Beginner",
	},
	"1F531": {
		Key:        "1F531",
		Value:      "🔱",
		Descriptor: "Trident Emblem",
	},
	"1F532": {
		Key:        "1F532",
		Value:      "🔲",
		Descriptor: "Black Square Button",
	},
	"1F533": {
		Key:        "1F533",
		Value:      "🔳",
		Descriptor: "White Square Button",
	},
	"1F534": {
		Key:        "1F534",
		Value:      "🔴",
		Descriptor: "Red Circle",
	},
	"1F535": {
		Key:        "1F535",
		Value:      "🔵",
		Descriptor: "Blue Circle",
	},
	"1F536": {
		Key:        "1F536",
		Value:      "🔶",
		Descriptor: "Large Orange Diamond",
	},
	"1F537": {
		Key:        "1F537",
		Value:      "🔷",
		Descriptor: "Large Blue Diamond",
	},
	"1F538": {
		Key:        "1F538",
		Value:      "🔸",
		Descriptor: "Small Orange Diamond",
	},
	"1F539": {
		Key:        "1F539",
		Value:      "🔹",
		Descriptor: "Small Blue Diamond",
	},
	"1F53A": {
		Key:        "1F53A",
		Value:      "🔺",
		Descriptor: "Red Triangle Pointed Up",
	},
	"1F53B": {
		Key:        "1F53B",
		Value:      "🔻",
		Descriptor: "Red Triangle Pointed Down",
	},
	"1F53C": {
		Key:        "1F53C",
		Value:      "🔼",
		Descriptor: "Upwards Button",
	},
	"1F53D": {
		Key:        "1F53D",
		Value:      "🔽",
		Descriptor: "Downwards Button",
	},
	"1F549-FE0F": {
		Key:        "1F549-FE0F",
		Value:      "🕉️",
		Descriptor: "Om",
	},
	"1F54A-FE0F": {
		Key:        "1F54A-FE0F",
		Value:      "🕊️",
		Descriptor: "Dove",
	},
	"1F54B": {
		Key:        "1F54B",
		Value:      "🕋",
		Descriptor: "Kaaba",
	},
	"1F54C": {
		Key:        "1F54C",
		Value:      "🕌",
		Descriptor: "Mosque",
	},
	"1F54D": {
		Key:        "1F54D",
		Value:      "🕍",
		Descriptor: "Synagogue",
	},
	"1F54E": {
		Key:        "1F54E",
		Value:      "🕎",
		Descriptor: "Menorah",
	},
	"1F550": {
		Key:        "1F550",
		Value:      "🕐",
		Descriptor: "One O’Clock",
	},
	"1F551": {
		Key:        "1F551",
		Value:      "🕑",
		Descriptor: "Two O’Clock",
	},
	"1F552": {
		Key:        "1F552",
		Value:      "🕒",
		Descriptor: "Three O’Clock",
	},
	"1F553": {
		Key:        "1F553",
		Value:      "🕓",
		Descriptor: "Four O’Clock",
	},
	"1F554": {
		Key:        "1F554",
		Value:      "🕔",
		Descriptor: "Five O’Clock",
	},
	"1F555": {
		Key:        "1F555",
		Value:      "🕕",
		Descriptor: "Six O’Clock",
	},
	"1F556": {
		Key:        "1F556",
		Value:      "🕖",
		Descriptor: "Seven O’Clock",
	},
	"1F557": {
		Key:        "1F557",
		Value:      "🕗",
		Descriptor: "Eight O’Clock",
	},
	"1F558": {
		Key:        "1F558",
		Value:      "🕘",
		Descriptor: "Nine O’Clock",
	},
	"1F559": {
		Key:        "1F559",
		Value:      "🕙",
		Descriptor: "Ten O’Clock",
	},
	"1F55A": {
		Key:        "1F55A",
		Value:      "🕚",
		Descriptor: "Eleven O’Clock",
	},
	"1F55B": {
		Key:        "1F55B",
		Value:      "🕛",
		Descriptor: "Twelve O’Clock",
	},
	"1F55C": {
		Key:        "1F55C",
		Value:      "🕜",
		Descriptor: "One-Thirty",
	},
	"1F55D": {
		Key:        "1F55D",
		Value:      "🕝",
		Descriptor: "Two-Thirty",
	},
	"1F55E": {
		Key:        "1F55E",
		Value:      "🕞",
		Descriptor: "Three-Thirty",
	},
	"1F55F": {
		Key:        "1F55F",
		Value:      "🕟",
		Descriptor: "Four-Thirty",
	},
	"1F560": {
		Key:        "1F560",
		Value:      "🕠",
		Descriptor: "Five-Thirty",
	},
	"1F561": {
		Key:        "1F561",
		Value:      "🕡",
		Descriptor: "Six-Thirty",
	},
	"1F562": {
		Key:        "1F562",
		Value:      "🕢",
		Descriptor: "Seven-Thirty",
	},
	"1F563": {
		Key:        "1F563",
		Value:      "🕣",
		Descriptor: "Eight-Thirty",
	},
	"1F564": {
		Key:        "1F564",
		Value:      "🕤",
		Descriptor: "Nine-Thirty",
	},
	"1F565": {
		Key:        "1F565",
		Value:      "🕥",
		Descriptor: "Ten-Thirty",
	},
	"1F566": {
		Key:        "1F566",
		Value:      "🕦",
		Descriptor: "Eleven-Thirty",
	},
	"1F567": {
		Key:        "1F567",
		Value:      "🕧",
		Descriptor: "Twelve-Thirty",
	},
	"1F56F-FE0F": {
		Key:        "1F56F-FE0F",
		Value:      "🕯️",
		Descriptor: "Candle",
	},
	"1F570-FE0F": {
		Key:        "1F570-FE0F",
		Value:      "🕰️",
		Descriptor: "Mantelpiece Clock",
	},
	"1F573-FE0F": {
		Key:        "1F573-FE0F",
		Value:      "🕳️",
		Descriptor: "Hole",
	},
	"1F574-1F3FB": {
		Key:        "1F574-1F3FB",
		Value:      "🕴🏻",
		Descriptor: "Man in Suit Levitating: Light Skin Tone",
	},
	"1F574-1F3FC": {
		Key:        "1F574-1F3FC",
		Value:      "🕴🏼",
		Descriptor: "Man in Suit Levitating: Medium-Light Skin Tone",
	},
	"1F574-1F3FD": {
		Key:        "1F574-1F3FD",
		Value:      "🕴🏽",
		Descriptor: "Man in Suit Levitating: Medium Skin Tone",
	},
	"1F574-1F3FE": {
		Key:        "1F574-1F3FE",
		Value:      "🕴🏾",
		Descriptor: "Man in Suit Levitating: Medium-Dark Skin Tone",
	},
	"1F574-1F3FF": {
		Key:        "1F574-1F3FF",
		Value:      "🕴🏿",
		Descriptor: "Man in Suit Levitating: Dark Skin Tone",
	},
	"1F574-FE0F": {
		Key:        "1F574-FE0F",
		Value:      "🕴️",
		Descriptor: "Man in Suit Levitating",
	},
	"1F575-1F3FB": {
		Key:        "1F575-1F3FB",
		Value:      "🕵🏻",
		Descriptor: "Detective: Light Skin Tone",
	},
	"1F575-1F3FB-200D-2640-FE0F": {
		Key:        "1F575-1F3FB-200D-2640-FE0F",
		Value:      "🕵🏻‍♀️",
		Descriptor: "Woman Detective: Light Skin Tone",
	},
	"1F575-1F3FB-200D-2642-FE0F": {
		Key:        "1F575-1F3FB-200D-2642-FE0F",
		Value:      "🕵🏻‍♂️",
		Descriptor: "Man Detective: Light Skin Tone",
	},
	"1F575-1F3FC": {
		Key:        "1F575-1F3FC",
		Value:      "🕵🏼",
		Descriptor: "Detective: Medium-Light Skin Tone",
	},
	"1F575-1F3FC-200D-2640-FE0F": {
		Key:        "1F575-1F3FC-200D-2640-FE0F",
		Value:      "🕵🏼‍♀️",
		Descriptor: "Woman Detective: Medium-Light Skin Tone",
	},
	"1F575-1F3FC-200D-2642-FE0F": {
		Key:        "1F575-1F3FC-200D-2642-FE0F",
		Value:      "🕵🏼‍♂️",
		Descriptor: "Man Detective: Medium-Light Skin Tone",
	},
	"1F575-1F3FD": {
		Key:        "1F575-1F3FD",
		Value:      "🕵🏽",
		Descriptor: "Detective: Medium Skin Tone",
	},
	"1F575-1F3FD-200D-2640-FE0F": {
		Key:        "1F575-1F3FD-200D-2640-FE0F",
		Value:      "🕵🏽‍♀️",
		Descriptor: "Woman Detective: Medium Skin Tone",
	},
	"1F575-1F3FD-200D-2642-FE0F": {
		Key:        "1F575-1F3FD-200D-2642-FE0F",
		Value:      "🕵🏽‍♂️",
		Descriptor: "Man Detective: Medium Skin Tone",
	},
	"1F575-1F3FE": {
		Key:        "1F575-1F3FE",
		Value:      "🕵🏾",
		Descriptor: "Detective: Medium-Dark Skin Tone",
	},
	"1F575-1F3FE-200D-2640-FE0F": {
		Key:        "1F575-1F3FE-200D-2640-FE0F",
		Value:      "🕵🏾‍♀️",
		Descriptor: "Woman Detective: Medium-Dark Skin Tone",
	},
	"1F575-1F3FE-200D-2642-FE0F": {
		Key:        "1F575-1F3FE-200D-2642-FE0F",
		Value:      "🕵🏾‍♂️",
		Descriptor: "Man Detective: Medium-Dark Skin Tone",
	},
	"1F575-1F3FF": {
		Key:        "1F575-1F3FF",
		Value:      "🕵🏿",
		Descriptor: "Detective: Dark Skin Tone",
	},
	"1F575-1F3FF-200D-2640-FE0F": {
		Key:        "1F575-1F3FF-200D-2640-FE0F",
		Value:      "🕵🏿‍♀️",
		Descriptor: "Woman Detective: Dark Skin Tone",
	},
	"1F575-1F3FF-200D-2642-FE0F": {
		Key:        "1F575-1F3FF-200D-2642-FE0F",
		Value:      "🕵🏿‍♂️",
		Descriptor: "Man Detective: Dark Skin Tone",
	},
	"1F575-FE0F": {
		Key:        "1F575-FE0F",
		Value:      "🕵️",
		Descriptor: "Detective",
	},
	"1F575-FE0F-200D-2640-FE0F": {
		Key:        "1F575-FE0F-200D-2640-FE0F",
		Value:      "🕵️‍♀️",
		Descriptor: "Woman Detective",
	},
	"1F575-FE0F-200D-2642-FE0F": {
		Key:        "1F575-FE0F-200D-2642-FE0F",
		Value:      "🕵️‍♂️",
		Descriptor: "Man Detective",
	},
	"1F576-FE0F": {
		Key:        "1F576-FE0F",
		Value:      "🕶️",
		Descriptor: "Sunglasses",
	},
	"1F577-FE0F": {
		Key:        "1F577-FE0F",
		Value:      "🕷️",
		Descriptor: "Spider",
	},
	"1F578-FE0F": {
		Key:        "1F578-FE0F",
		Value:      "🕸️",
		Descriptor: "Spider Web",
	},
	"1F579-FE0F": {
		Key:        "1F579-FE0F",
		Value:      "🕹️",
		Descriptor: "Joystick",
	},
	"1F57A": {
		Key:        "1F57A",
		Value:      "🕺",
		Descriptor: "Man Dancing",
	},
	"1F57A-1F3FB": {
		Key:        "1F57A-1F3FB",
		Value:      "🕺🏻",
		Descriptor: "Man Dancing: Light Skin Tone",
	},
	"1F57A-1F3FC": {
		Key:        "1F57A-1F3FC",
		Value:      "🕺🏼",
		Descriptor: "Man Dancing: Medium-Light Skin Tone",
	},
	"1F57A-1F3FD": {
		Key:        "1F57A-1F3FD",
		Value:      "🕺🏽",
		Descriptor: "Man Dancing: Medium Skin Tone",
	},
	"1F57A-1F3FE": {
		Key:        "1F57A-1F3FE",
		Value:      "🕺🏾",
		Descriptor: "Man Dancing: Medium-Dark Skin Tone",
	},
	"1F57A-1F3FF": {
		Key:        "1F57A-1F3FF",
		Value:      "🕺🏿",
		Descriptor: "Man Dancing: Dark Skin Tone",
	},
	"1F587-FE0F": {
		Key:        "1F587-FE0F",
		Value:      "🖇️",
		Descriptor: "Linked Paperclips",
	},
	"1F58A-FE0F": {
		Key:        "1F58A-FE0F",
		Value:      "🖊️",
		Descriptor: "Pen",
	},
	"1F58B-FE0F": {
		Key:        "1F58B-FE0F",
		Value:      "🖋️",
		Descriptor: "Fountain Pen",
	},
	"1F58C-FE0F": {
		Key:        "1F58C-FE0F",
		Value:      "🖌️",
		Descriptor: "Paintbrush",
	},
	"1F58D-FE0F": {
		Key:        "1F58D-FE0F",
		Value:      "🖍️",
		Descriptor: "Crayon",
	},
	"1F590-1F3FB": {
		Key:        "1F590-1F3FB",
		Value:      "🖐🏻",
		Descriptor: "Hand With Fingers Splayed: Light Skin Tone",
	},
	"1F590-1F3FC": {
		Key:        "1F590-1F3FC",
		Value:      "🖐🏼",
		Descriptor: "Hand With Fingers Splayed: Medium-Light Skin Tone",
	},
	"1F590-1F3FD": {
		Key:        "1F590-1F3FD",
		Value:      "🖐🏽",
		Descriptor: "Hand With Fingers Splayed: Medium Skin Tone",
	},
	"1F590-1F3FE": {
		Key:        "1F590-1F3FE",
		Value:      "🖐🏾",
		Descriptor: "Hand With Fingers Splayed: Medium-Dark Skin Tone",
	},
	"1F590-1F3FF": {
		Key:        "1F590-1F3FF",
		Value:      "🖐🏿",
		Descriptor: "Hand With Fingers Splayed: Dark Skin Tone",
	},
	"1F590-FE0F": {
		Key:        "1F590-FE0F",
		Value:      "🖐️",
		Descriptor: "Hand With Fingers Splayed",
	},
	"1F595": {
		Key:        "1F595",
		Value:      "🖕",
		Descriptor: "Middle Finger",
	},
	"1F595-1F3FB": {
		Key:        "1F595-1F3FB",
		Value:      "🖕🏻",
		Descriptor: "Middle Finger: Light Skin Tone",
	},
	"1F595-1F3FC": {
		Key:        "1F595-1F3FC",
		Value:      "🖕🏼",
		Descriptor: "Middle Finger: Medium-Light Skin Tone",
	},
	"1F595-1F3FD": {
		Key:        "1F595-1F3FD",
		Value:      "🖕🏽",
		Descriptor: "Middle Finger: Medium Skin Tone",
	},
	"1F595-1F3FE": {
		Key:        "1F595-1F3FE",
		Value:      "🖕🏾",
		Descriptor: "Middle Finger: Medium-Dark Skin Tone",
	},
	"1F595-1F3FF": {
		Key:        "1F595-1F3FF",
		Value:      "🖕🏿",
		Descriptor: "Middle Finger: Dark Skin Tone",
	},
	"1F596": {
		Key:        "1F596",
		Value:      "🖖",
		Descriptor: "Vulcan Salute",
	},
	"1F596-1F3FB": {
		Key:        "1F596-1F3FB",
		Value:      "🖖🏻",
		Descriptor: "Vulcan Salute: Light Skin Tone",
	},
	"1F596-1F3FC": {
		Key:        "1F596-1F3FC",
		Value:      "🖖🏼",
		Descriptor: "Vulcan Salute: Medium-Light Skin Tone",
	},
	"1F596-1F3FD": {
		Key:        "1F596-1F3FD",
		Value:      "🖖🏽",
		Descriptor: "Vulcan Salute: Medium Skin Tone",
	},
	"1F596-1F3FE": {
		Key:        "1F596-1F3FE",
		Value:      "🖖🏾",
		Descriptor: "Vulcan Salute: Medium-Dark Skin Tone",
	},
	"1F596-1F3FF": {
		Key:        "1F596-1F3FF",
		Value:      "🖖🏿",
		Descriptor: "Vulcan Salute: Dark Skin Tone",
	},
	"1F5A4": {
		Key:        "1F5A4",
		Value:      "🖤",
		Descriptor: "Black Heart",
	},
	"1F5A5-FE0F": {
		Key:        "1F5A5-FE0F",
		Value:      "🖥️",
		Descriptor: "Desktop Computer",
	},
	"1F5A8-FE0F": {
		Key:        "1F5A8-FE0F",
		Value:      "🖨️",
		Descriptor: "Printer",
	},
	"1F5B1-FE0F": {
		Key:        "1F5B1-FE0F",
		Value:      "🖱️",
		Descriptor: "Computer Mouse",
	},
	"1F5B2-FE0F": {
		Key:        "1F5B2-FE0F",
		Value:      "🖲️",
		Descriptor: "Trackball",
	},
	"1F5BC-FE0F": {
		Key:        "1F5BC-FE0F",
		Value:      "🖼️",
		Descriptor: "Framed Picture",
	},
	"1F5C2-FE0F": {
		Key:        "1F5C2-FE0F",
		Value:      "🗂️",
		Descriptor: "Card Index Dividers",
	},
	"1F5C3-FE0F": {
		Key:        "1F5C3-FE0F",
		Value:      "🗃️",
		Descriptor: "Card File Box",
	},
	"1F5C4-FE0F": {
		Key:        "1F5C4-FE0F",
		Value:      "🗄️",
		Descriptor: "File Cabinet",
	},
	"1F5D1-FE0F": {
		Key:        "1F5D1-FE0F",
		Value:      "🗑️",
		Descriptor: "Wastebasket",
	},
	"1F5D2-FE0F": {
		Key:        "1F5D2-FE0F",
		Value:      "🗒️",
		Descriptor: "Spiral Notepad",
	},
	"1F5D3-FE0F": {
		Key:        "1F5D3-FE0F",
		Value:      "🗓️",
		Descriptor: "Spiral Calendar",
	},
	"1F5DC-FE0F": {
		Key:        "1F5DC-FE0F",
		Value:      "🗜️",
		Descriptor: "Clamp",
	},
	"1F5DD-FE0F": {
		Key:        "1F5DD-FE0F",
		Value:      "🗝️",
		Descriptor: "Old Key",
	},
	"1F5DE-FE0F": {
		Key:        "1F5DE-FE0F",
		Value:      "🗞️",
		Descriptor: "Rolled-Up Newspaper",
	},
	"1F5E1-FE0F": {
		Key:        "1F5E1-FE0F",
		Value:      "🗡️",
		Descriptor: "Dagger",
	},
	"1F5E3-FE0F": {
		Key:        "1F5E3-FE0F",
		Value:      "🗣️",
		Descriptor: "Speaking Head",
	},
	"1F5E8-FE0F": {
		Key:        "1F5E8-FE0F",
		Value:      "🗨️",
		Descriptor: "Left Speech Bubble",
	},
	"1F5EF-FE0F": {
		Key:        "1F5EF-FE0F",
		Value:      "🗯️",
		Descriptor: "Right Anger Bubble",
	},
	"1F5F3-FE0F": {
		Key:        "1F5F3-FE0F",
		Value:      "🗳️",
		Descriptor: "Ballot Box With Ballot",
	},
	"1F5FA-FE0F": {
		Key:        "1F5FA-FE0F",
		Value:      "🗺️",
		Descriptor: "World Map",
	},
	"1F5FB": {
		Key:        "1F5FB",
		Value:      "🗻",
		Descriptor: "Mount Fuji",
	},
	"1F5FC": {
		Key:        "1F5FC",
		Value:      "🗼",
		Descriptor: "Tokyo Tower",
	},
	"1F5FD": {
		Key:        "1F5FD",
		Value:      "🗽",
		Descriptor: "Statue of Liberty",
	},
	"1F5FE": {
		Key:        "1F5FE",
		Value:      "🗾",
		Descriptor: "Map of Japan",
	},
	"1F5FF": {
		Key:        "1F5FF",
		Value:      "🗿",
		Descriptor: "Moai",
	},
	"1F600": {
		Key:        "1F600",
		Value:      "😀",
		Descriptor: "Grinning Face",
	},
	"1F601": {
		Key:        "1F601",
		Value:      "😁",
		Descriptor: "Beaming Face With Smiling Eyes",
	},
	"1F602": {
		Key:        "1F602",
		Value:      "😂",
		Descriptor: "Face With Tears of Joy",
	},
	"1F603": {
		Key:        "1F603",
		Value:      "😃",
		Descriptor: "Grinning Face With Big Eyes",
	},
	"1F604": {
		Key:        "1F604",
		Value:      "😄",
		Descriptor: "Grinning Face With Smiling Eyes",
	},
	"1F605": {
		Key:        "1F605",
		Value:      "😅",
		Descriptor: "Grinning Face With Sweat",
	},
	"1F606": {
		Key:        "1F606",
		Value:      "😆",
		Descriptor: "Grinning Squinting Face",
	},
	"1F607": {
		Key:        "1F607",
		Value:      "😇",
		Descriptor: "Smiling Face With Halo",
	},
	"1F608": {
		Key:        "1F608",
		Value:      "😈",
		Descriptor: "Smiling Face With Horns",
	},
	"1F609": {
		Key:        "1F609",
		Value:      "😉",
		Descriptor: "Winking Face",
	},
	"1F60A": {
		Key:        "1F60A",
		Value:      "😊",
		Descriptor: "Smiling Face With Smiling Eyes",
	},
	"1F60B": {
		Key:        "1F60B",
		Value:      "😋",
		Descriptor: "Face Savoring Food",
	},
	"1F60C": {
		Key:        "1F60C",
		Value:      "😌",
		Descriptor: "Relieved Face",
	},
	"1F60D": {
		Key:        "1F60D",
		Value:      "😍",
		Descriptor: "Smiling Face With Heart-Eyes",
	},
	"1F60E": {
		Key:        "1F60E",
		Value:      "😎",
		Descriptor: "Smiling Face With Sunglasses",
	},
	"1F60F": {
		Key:        "1F60F",
		Value:      "😏",
		Descriptor: "Smirking Face",
	},
	"1F610": {
		Key:        "1F610",
		Value:      "😐",
		Descriptor: "Neutral Face",
	},
	"1F611": {
		Key:        "1F611",
		Value:      "😑",
		Descriptor: "Expressionless Face",
	},
	"1F612": {
		Key:        "1F612",
		Value:      "😒",
		Descriptor: "Unamused Face",
	},
	"1F613": {
		Key:        "1F613",
		Value:      "😓",
		Descriptor: "Downcast Face With Sweat",
	},
	"1F614": {
		Key:        "1F614",
		Value:      "😔",
		Descriptor: "Pensive Face",
	},
	"1F615": {
		Key:        "1F615",
		Value:      "😕",
		Descriptor: "Confused Face",
	},
	"1F616": {
		Key:        "1F616",
		Value:      "😖",
		Descriptor: "Confounded Face",
	},
	"1F617": {
		Key:        "1F617",
		Value:      "😗",
		Descriptor: "Kissing Face",
	},
	"1F618": {
		Key:        "1F618",
		Value:      "😘",
		Descriptor: "Face Blowing a Kiss",
	},
	"1F619": {
		Key:        "1F619",
		Value:      "😙",
		Descriptor: "Kissing Face With Smiling Eyes",
	},
	"1F61A": {
		Key:        "1F61A",
		Value:      "😚",
		Descriptor: "Kissing Face With Closed Eyes",
	},
	"1F61B": {
		Key:        "1F61B",
		Value:      "😛",
		Descriptor: "Face With Tongue",
	},
	"1F61C": {
		Key:        "1F61C",
		Value:      "😜",
		Descriptor: "Winking Face With Tongue",
	},
	"1F61D": {
		Key:        "1F61D",
		Value:      "😝",
		Descriptor: "Squinting Face With Tongue",
	},
	"1F61E": {
		Key:        "1F61E",
		Value:      "😞",
		Descriptor: "Disappointed Face",
	},
	"1F61F": {
		Key:        "1F61F",
		Value:      "😟",
		Descriptor: "Worried Face",
	},
	"1F620": {
		Key:        "1F620",
		Value:      "😠",
		Descriptor: "Angry Face",
	},
	"1F621": {
		Key:        "1F621",
		Value:      "😡",
		Descriptor: "Pouting Face",
	},
	"1F622": {
		Key:        "1F622",
		Value:      "😢",
		Descriptor: "Crying Face",
	},
	"1F623": {
		Key:        "1F623",
		Value:      "😣",
		Descriptor: "Persevering Face",
	},
	"1F624": {
		Key:        "1F624",
		Value:      "😤",
		Descriptor: "Face With Steam From Nose",
	},
	"1F625": {
		Key:        "1F625",
		Value:      "😥",
		Descriptor: "Sad but Relieved Face",
	},
	"1F626": {
		Key:        "1F626",
		Value:      "😦",
		Descriptor: "Frowning Face With Open Mouth",
	},
	"1F627": {
		Key:        "1F627",
		Value:      "😧",
		Descriptor: "Anguished Face",
	},
	"1F628": {
		Key:        "1F628",
		Value:      "😨",
		Descriptor: "Fearful Face",
	},
	"1F629": {
		Key:        "1F629",
		Value:      "😩",
		Descriptor: "Weary Face",
	},
	"1F62A": {
		Key:        "1F62A",
		Value:      "😪",
		Descriptor: "Sleepy Face",
	},
	"1F62B": {
		Key:        "1F62B",
		Value:      "😫",
		Descriptor: "Tired Face",
	},
	"1F62C": {
		Key:        "1F62C",
		Value:      "😬",
		Descriptor: "Grimacing Face",
	},
	"1F62D": {
		Key:        "1F62D",
		Value:      "😭",
		Descriptor: "Loudly Crying Face",
	},
	"1F62E": {
		Key:        "1F62E",
		Value:      "😮",
		Descriptor: "Face With Open Mouth",
	},
	"1F62F": {
		Key:        "1F62F",
		Value:      "😯",
		Descriptor: "Hushed Face",
	},
	"1F630": {
		Key:        "1F630",
		Value:      "😰",
		Descriptor: "Anxious Face With Sweat",
	},
	"1F631": {
		Key:        "1F631",
		Value:      "😱",
		Descriptor: "Face Screaming in Fear",
	},
	"1F632": {
		Key:        "1F632",
		Value:      "😲",
		Descriptor: "Astonished Face",
	},
	"1F633": {
		Key:        "1F633",
		Value:      "😳",
		Descriptor: "Flushed Face",
	},
	"1F634": {
		Key:        "1F634",
		Value:      "😴",
		Descriptor: "Sleeping Face",
	},
	"1F635": {
		Key:        "1F635",
		Value:      "😵",
		Descriptor: "Dizzy Face",
	},
	"1F636": {
		Key:        "1F636",
		Value:      "😶",
		Descriptor: "Face Without Mouth",
	},
	"1F637": {
		Key:        "1F637",
		Value:      "😷",
		Descriptor: "Face With Medical Mask",
	},
	"1F638": {
		Key:        "1F638",
		Value:      "😸",
		Descriptor: "Grinning Cat With Smiling Eyes",
	},
	"1F639": {
		Key:        "1F639",
		Value:      "😹",
		Descriptor: "Cat With Tears of Joy",
	},
	"1F63A": {
		Key:        "1F63A",
		Value:      "😺",
		Descriptor: "Grinning Cat",
	},
	"1F63B": {
		Key:        "1F63B",
		Value:      "😻",
		Descriptor: "Smiling Cat With Heart-Eyes",
	},
	"1F63C": {
		Key:        "1F63C",
		Value:      "😼",
		Descriptor: "Cat With Wry Smile",
	},
	"1F63D": {
		Key:        "1F63D",
		Value:      "😽",
		Descriptor: "Kissing Cat",
	},
	"1F63E": {
		Key:        "1F63E",
		Value:      "😾",
		Descriptor: "Pouting Cat",
	},
	"1F63F": {
		Key:        "1F63F",
		Value:      "😿",
		Descriptor: "Crying Cat",
	},
	"1F640": {
		Key:        "1F640",
		Value:      "🙀",
		Descriptor: "Weary Cat",
	},
	"1F641": {
		Key:        "1F641",
		Value:      "🙁",
		Descriptor: "Slightly Frowning Face",
	},
	"1F642": {
		Key:        "1F642",
		Value:      "🙂",
		Descriptor: "Slightly Smiling Face",
	},
	"1F643": {
		Key:        "1F643",
		Value:      "🙃",
		Descriptor: "Upside-Down Face",
	},
	"1F644": {
		Key:        "1F644",
		Value:      "🙄",
		Descriptor: "Face With Rolling Eyes",
	},
	"1F645": {
		Key:        "1F645",
		Value:      "🙅",
		Descriptor: "Person Gesturing No",
	},
	"1F645-1F3FB": {
		Key:        "1F645-1F3FB",
		Value:      "🙅🏻",
		Descriptor: "Person Gesturing No: Light Skin Tone",
	},
	"1F645-1F3FB-200D-2640-FE0F": {
		Key:        "1F645-1F3FB-200D-2640-FE0F",
		Value:      "🙅🏻‍♀️",
		Descriptor: "Woman Gesturing No: Light Skin Tone",
	},
	"1F645-1F3FB-200D-2642-FE0F": {
		Key:        "1F645-1F3FB-200D-2642-FE0F",
		Value:      "🙅🏻‍♂️",
		Descriptor: "Man Gesturing No: Light Skin Tone",
	},
	"1F645-1F3FC": {
		Key:        "1F645-1F3FC",
		Value:      "🙅🏼",
		Descriptor: "Person Gesturing No: Medium-Light Skin Tone",
	},
	"1F645-1F3FC-200D-2640-FE0F": {
		Key:        "1F645-1F3FC-200D-2640-FE0F",
		Value:      "🙅🏼‍♀️",
		Descriptor: "Woman Gesturing No: Medium-Light Skin Tone",
	},
	"1F645-1F3FC-200D-2642-FE0F": {
		Key:        "1F645-1F3FC-200D-2642-FE0F",
		Value:      "🙅🏼‍♂️",
		Descriptor: "Man Gesturing No: Medium-Light Skin Tone",
	},
	"1F645-1F3FD": {
		Key:        "1F645-1F3FD",
		Value:      "🙅🏽",
		Descriptor: "Person Gesturing No: Medium Skin Tone",
	},
	"1F645-1F3FD-200D-2640-FE0F": {
		Key:        "1F645-1F3FD-200D-2640-FE0F",
		Value:      "🙅🏽‍♀️",
		Descriptor: "Woman Gesturing No: Medium Skin Tone",
	},
	"1F645-1F3FD-200D-2642-FE0F": {
		Key:        "1F645-1F3FD-200D-2642-FE0F",
		Value:      "🙅🏽‍♂️",
		Descriptor: "Man Gesturing No: Medium Skin Tone",
	},
	"1F645-1F3FE": {
		Key:        "1F645-1F3FE",
		Value:      "🙅🏾",
		Descriptor: "Person Gesturing No: Medium-Dark Skin Tone",
	},
	"1F645-1F3FE-200D-2640-FE0F": {
		Key:        "1F645-1F3FE-200D-2640-FE0F",
		Value:      "🙅🏾‍♀️",
		Descriptor: "Woman Gesturing No: Medium-Dark Skin Tone",
	},
	"1F645-1F3FE-200D-2642-FE0F": {
		Key:        "1F645-1F3FE-200D-2642-FE0F",
		Value:      "🙅🏾‍♂️",
		Descriptor: "Man Gesturing No: Medium-Dark Skin Tone",
	},
	"1F645-1F3FF": {
		Key:        "1F645-1F3FF",
		Value:      "🙅🏿",
		Descriptor: "Person Gesturing No: Dark Skin Tone",
	},
	"1F645-1F3FF-200D-2640-FE0F": {
		Key:        "1F645-1F3FF-200D-2640-FE0F",
		Value:      "🙅🏿‍♀️",
		Descriptor: "Woman Gesturing No: Dark Skin Tone",
	},
	"1F645-1F3FF-200D-2642-FE0F": {
		Key:        "1F645-1F3FF-200D-2642-FE0F",
		Value:      "🙅🏿‍♂️",
		Descriptor: "Man Gesturing No: Dark Skin Tone",
	},
	"1F645-200D-2640-FE0F": {
		Key:        "1F645-200D-2640-FE0F",
		Value:      "🙅‍♀️",
		Descriptor: "Woman Gesturing No",
	},
	"1F645-200D-2642-FE0F": {
		Key:        "1F645-200D-2642-FE0F",
		Value:      "🙅‍♂️",
		Descriptor: "Man Gesturing No",
	},
	"1F646": {
		Key:        "1F646",
		Value:      "🙆",
		Descriptor: "Person Gesturing OK",
	},
	"1F646-1F3FB": {
		Key:        "1F646-1F3FB",
		Value:      "🙆🏻",
		Descriptor: "Person Gesturing OK: Light Skin Tone",
	},
	"1F646-1F3FB-200D-2640-FE0F": {
		Key:        "1F646-1F3FB-200D-2640-FE0F",
		Value:      "🙆🏻‍♀️",
		Descriptor: "Woman Gesturing OK: Light Skin Tone",
	},
	"1F646-1F3FB-200D-2642-FE0F": {
		Key:        "1F646-1F3FB-200D-2642-FE0F",
		Value:      "🙆🏻‍♂️",
		Descriptor: "Man Gesturing OK: Light Skin Tone",
	},
	"1F646-1F3FC": {
		Key:        "1F646-1F3FC",
		Value:      "🙆🏼",
		Descriptor: "Person Gesturing OK: Medium-Light Skin Tone",
	},
	"1F646-1F3FC-200D-2640-FE0F": {
		Key:        "1F646-1F3FC-200D-2640-FE0F",
		Value:      "🙆🏼‍♀️",
		Descriptor: "Woman Gesturing OK: Medium-Light Skin Tone",
	},
	"1F646-1F3FC-200D-2642-FE0F": {
		Key:        "1F646-1F3FC-200D-2642-FE0F",
		Value:      "🙆🏼‍♂️",
		Descriptor: "Man Gesturing OK: Medium-Light Skin Tone",
	},
	"1F646-1F3FD": {
		Key:        "1F646-1F3FD",
		Value:      "🙆🏽",
		Descriptor: "Person Gesturing OK: Medium Skin Tone",
	},
	"1F646-1F3FD-200D-2640-FE0F": {
		Key:        "1F646-1F3FD-200D-2640-FE0F",
		Value:      "🙆🏽‍♀️",
		Descriptor: "Woman Gesturing OK: Medium Skin Tone",
	},
	"1F646-1F3FD-200D-2642-FE0F": {
		Key:        "1F646-1F3FD-200D-2642-FE0F",
		Value:      "🙆🏽‍♂️",
		Descriptor: "Man Gesturing OK: Medium Skin Tone",
	},
	"1F646-1F3FE": {
		Key:        "1F646-1F3FE",
		Value:      "🙆🏾",
		Descriptor: "Person Gesturing OK: Medium-Dark Skin Tone",
	},
	"1F646-1F3FE-200D-2640-FE0F": {
		Key:        "1F646-1F3FE-200D-2640-FE0F",
		Value:      "🙆🏾‍♀️",
		Descriptor: "Woman Gesturing OK: Medium-Dark Skin Tone",
	},
	"1F646-1F3FE-200D-2642-FE0F": {
		Key:        "1F646-1F3FE-200D-2642-FE0F",
		Value:      "🙆🏾‍♂️",
		Descriptor: "Man Gesturing OK: Medium-Dark Skin Tone",
	},
	"1F646-1F3FF": {
		Key:        "1F646-1F3FF",
		Value:      "🙆🏿",
		Descriptor: "Person Gesturing OK: Dark Skin Tone",
	},
	"1F646-1F3FF-200D-2640-FE0F": {
		Key:        "1F646-1F3FF-200D-2640-FE0F",
		Value:      "🙆🏿‍♀️",
		Descriptor: "Woman Gesturing OK: Dark Skin Tone",
	},
	"1F646-1F3FF-200D-2642-FE0F": {
		Key:        "1F646-1F3FF-200D-2642-FE0F",
		Value:      "🙆🏿‍♂️",
		Descriptor: "Man Gesturing OK: Dark Skin Tone",
	},
	"1F646-200D-2640-FE0F": {
		Key:        "1F646-200D-2640-FE0F",
		Value:      "🙆‍♀️",
		Descriptor: "Woman Gesturing OK",
	},
	"1F646-200D-2642-FE0F": {
		Key:        "1F646-200D-2642-FE0F",
		Value:      "🙆‍♂️",
		Descriptor: "Man Gesturing OK",
	},
	"1F647": {
		Key:        "1F647",
		Value:      "🙇",
		Descriptor: "Person Bowing",
	},
	"1F647-1F3FB": {
		Key:        "1F647-1F3FB",
		Value:      "🙇🏻",
		Descriptor: "Person Bowing: Light Skin Tone",
	},
	"1F647-1F3FB-200D-2640-FE0F": {
		Key:        "1F647-1F3FB-200D-2640-FE0F",
		Value:      "🙇🏻‍♀️",
		Descriptor: "Woman Bowing: Light Skin Tone",
	},
	"1F647-1F3FB-200D-2642-FE0F": {
		Key:        "1F647-1F3FB-200D-2642-FE0F",
		Value:      "🙇🏻‍♂️",
		Descriptor: "Man Bowing: Light Skin Tone",
	},
	"1F647-1F3FC": {
		Key:        "1F647-1F3FC",
		Value:      "🙇🏼",
		Descriptor: "Person Bowing: Medium-Light Skin Tone",
	},
	"1F647-1F3FC-200D-2640-FE0F": {
		Key:        "1F647-1F3FC-200D-2640-FE0F",
		Value:      "🙇🏼‍♀️",
		Descriptor: "Woman Bowing: Medium-Light Skin Tone",
	},
	"1F647-1F3FC-200D-2642-FE0F": {
		Key:        "1F647-1F3FC-200D-2642-FE0F",
		Value:      "🙇🏼‍♂️",
		Descriptor: "Man Bowing: Medium-Light Skin Tone",
	},
	"1F647-1F3FD": {
		Key:        "1F647-1F3FD",
		Value:      "🙇🏽",
		Descriptor: "Person Bowing: Medium Skin Tone",
	},
	"1F647-1F3FD-200D-2640-FE0F": {
		Key:        "1F647-1F3FD-200D-2640-FE0F",
		Value:      "🙇🏽‍♀️",
		Descriptor: "Woman Bowing: Medium Skin Tone",
	},
	"1F647-1F3FD-200D-2642-FE0F": {
		Key:        "1F647-1F3FD-200D-2642-FE0F",
		Value:      "🙇🏽‍♂️",
		Descriptor: "Man Bowing: Medium Skin Tone",
	},
	"1F647-1F3FE": {
		Key:        "1F647-1F3FE",
		Value:      "🙇🏾",
		Descriptor: "Person Bowing: Medium-Dark Skin Tone",
	},
	"1F647-1F3FE-200D-2640-FE0F": {
		Key:        "1F647-1F3FE-200D-2640-FE0F",
		Value:      "🙇🏾‍♀️",
		Descriptor: "Woman Bowing: Medium-Dark Skin Tone",
	},
	"1F647-1F3FE-200D-2642-FE0F": {
		Key:        "1F647-1F3FE-200D-2642-FE0F",
		Value:      "🙇🏾‍♂️",
		Descriptor: "Man Bowing: Medium-Dark Skin Tone",
	},
	"1F647-1F3FF": {
		Key:        "1F647-1F3FF",
		Value:      "🙇🏿",
		Descriptor: "Person Bowing: Dark Skin Tone",
	},
	"1F647-1F3FF-200D-2640-FE0F": {
		Key:        "1F647-1F3FF-200D-2640-FE0F",
		Value:      "🙇🏿‍♀️",
		Descriptor: "Woman Bowing: Dark Skin Tone",
	},
	"1F647-1F3FF-200D-2642-FE0F": {
		Key:        "1F647-1F3FF-200D-2642-FE0F",
		Value:      "🙇🏿‍♂️",
		Descriptor: "Man Bowing: Dark Skin Tone",
	},
	"1F647-200D-2640-FE0F": {
		Key:        "1F647-200D-2640-FE0F",
		Value:      "🙇‍♀️",
		Descriptor: "Woman Bowing",
	},
	"1F647-200D-2642-FE0F": {
		Key:        "1F647-200D-2642-FE0F",
		Value:      "🙇‍♂️",
		Descriptor: "Man Bowing",
	},
	"1F648": {
		Key:        "1F648",
		Value:      "🙈",
		Descriptor: "See-No-Evil Monkey",
	},
	"1F649": {
		Key:        "1F649",
		Value:      "🙉",
		Descriptor: "Hear-No-Evil Monkey",
	},
	"1F64A": {
		Key:        "1F64A",
		Value:      "🙊",
		Descriptor: "Speak-No-Evil Monkey",
	},
	"1F64B": {
		Key:        "1F64B",
		Value:      "🙋",
		Descriptor: "Person Raising Hand",
	},
	"1F64B-1F3FB": {
		Key:        "1F64B-1F3FB",
		Value:      "🙋🏻",
		Descriptor: "Person Raising Hand: Light Skin Tone",
	},
	"1F64B-1F3FB-200D-2640-FE0F": {
		Key:        "1F64B-1F3FB-200D-2640-FE0F",
		Value:      "🙋🏻‍♀️",
		Descriptor: "Woman Raising Hand: Light Skin Tone",
	},
	"1F64B-1F3FB-200D-2642-FE0F": {
		Key:        "1F64B-1F3FB-200D-2642-FE0F",
		Value:      "🙋🏻‍♂️",
		Descriptor: "Man Raising Hand: Light Skin Tone",
	},
	"1F64B-1F3FC": {
		Key:        "1F64B-1F3FC",
		Value:      "🙋🏼",
		Descriptor: "Person Raising Hand: Medium-Light Skin Tone",
	},
	"1F64B-1F3FC-200D-2640-FE0F": {
		Key:        "1F64B-1F3FC-200D-2640-FE0F",
		Value:      "🙋🏼‍♀️",
		Descriptor: "Woman Raising Hand: Medium-Light Skin Tone",
	},
	"1F64B-1F3FC-200D-2642-FE0F": {
		Key:        "1F64B-1F3FC-200D-2642-FE0F",
		Value:      "🙋🏼‍♂️",
		Descriptor: "Man Raising Hand: Medium-Light Skin Tone",
	},
	"1F64B-1F3FD": {
		Key:        "1F64B-1F3FD",
		Value:      "🙋🏽",
		Descriptor: "Person Raising Hand: Medium Skin Tone",
	},
	"1F64B-1F3FD-200D-2640-FE0F": {
		Key:        "1F64B-1F3FD-200D-2640-FE0F",
		Value:      "🙋🏽‍♀️",
		Descriptor: "Woman Raising Hand: Medium Skin Tone",
	},
	"1F64B-1F3FD-200D-2642-FE0F": {
		Key:        "1F64B-1F3FD-200D-2642-FE0F",
		Value:      "🙋🏽‍♂️",
		Descriptor: "Man Raising Hand: Medium Skin Tone",
	},
	"1F64B-1F3FE": {
		Key:        "1F64B-1F3FE",
		Value:      "🙋🏾",
		Descriptor: "Person Raising Hand: Medium-Dark Skin Tone",
	},
	"1F64B-1F3FE-200D-2640-FE0F": {
		Key:        "1F64B-1F3FE-200D-2640-FE0F",
		Value:      "🙋🏾‍♀️",
		Descriptor: "Woman Raising Hand: Medium-Dark Skin Tone",
	},
	"1F64B-1F3FE-200D-2642-FE0F": {
		Key:        "1F64B-1F3FE-200D-2642-FE0F",
		Value:      "🙋🏾‍♂️",
		Descriptor: "Man Raising Hand: Medium-Dark Skin Tone",
	},
	"1F64B-1F3FF": {
		Key:        "1F64B-1F3FF",
		Value:      "🙋🏿",
		Descriptor: "Person Raising Hand: Dark Skin Tone",
	},
	"1F64B-1F3FF-200D-2640-FE0F": {
		Key:        "1F64B-1F3FF-200D-2640-FE0F",
		Value:      "🙋🏿‍♀️",
		Descriptor: "Woman Raising Hand: Dark Skin Tone",
	},
	"1F64B-1F3FF-200D-2642-FE0F": {
		Key:        "1F64B-1F3FF-200D-2642-FE0F",
		Value:      "🙋🏿‍♂️",
		Descriptor: "Man Raising Hand: Dark Skin Tone",
	},
	"1F64B-200D-2640-FE0F": {
		Key:        "1F64B-200D-2640-FE0F",
		Value:      "🙋‍♀️",
		Descriptor: "Woman Raising Hand",
	},
	"1F64B-200D-2642-FE0F": {
		Key:        "1F64B-200D-2642-FE0F",
		Value:      "🙋‍♂️",
		Descriptor: "Man Raising Hand",
	},
	"1F64C": {
		Key:        "1F64C",
		Value:      "🙌",
		Descriptor: "Raising Hands",
	},
	"1F64C-1F3FB": {
		Key:        "1F64C-1F3FB",
		Value:      "🙌🏻",
		Descriptor: "Raising Hands: Light Skin Tone",
	},
	"1F64C-1F3FC": {
		Key:        "1F64C-1F3FC",
		Value:      "🙌🏼",
		Descriptor: "Raising Hands: Medium-Light Skin Tone",
	},
	"1F64C-1F3FD": {
		Key:        "1F64C-1F3FD",
		Value:      "🙌🏽",
		Descriptor: "Raising Hands: Medium Skin Tone",
	},
	"1F64C-1F3FE": {
		Key:        "1F64C-1F3FE",
		Value:      "🙌🏾",
		Descriptor: "Raising Hands: Medium-Dark Skin Tone",
	},
	"1F64C-1F3FF": {
		Key:        "1F64C-1F3FF",
		Value:      "🙌🏿",
		Descriptor: "Raising Hands: Dark Skin Tone",
	},
	"1F64D": {
		Key:        "1F64D",
		Value:      "🙍",
		Descriptor: "Person Frowning",
	},
	"1F64D-1F3FB": {
		Key:        "1F64D-1F3FB",
		Value:      "🙍🏻",
		Descriptor: "Person Frowning: Light Skin Tone",
	},
	"1F64D-1F3FB-200D-2640-FE0F": {
		Key:        "1F64D-1F3FB-200D-2640-FE0F",
		Value:      "🙍🏻‍♀️",
		Descriptor: "Woman Frowning: Light Skin Tone",
	},
	"1F64D-1F3FB-200D-2642-FE0F": {
		Key:        "1F64D-1F3FB-200D-2642-FE0F",
		Value:      "🙍🏻‍♂️",
		Descriptor: "Man Frowning: Light Skin Tone",
	},
	"1F64D-1F3FC": {
		Key:        "1F64D-1F3FC",
		Value:      "🙍🏼",
		Descriptor: "Person Frowning: Medium-Light Skin Tone",
	},
	"1F64D-1F3FC-200D-2640-FE0F": {
		Key:        "1F64D-1F3FC-200D-2640-FE0F",
		Value:      "🙍🏼‍♀️",
		Descriptor: "Woman Frowning: Medium-Light Skin Tone",
	},
	"1F64D-1F3FC-200D-2642-FE0F": {
		Key:        "1F64D-1F3FC-200D-2642-FE0F",
		Value:      "🙍🏼‍♂️",
		Descriptor: "Man Frowning: Medium-Light Skin Tone",
	},
	"1F64D-1F3FD": {
		Key:        "1F64D-1F3FD",
		Value:      "🙍🏽",
		Descriptor: "Person Frowning: Medium Skin Tone",
	},
	"1F64D-1F3FD-200D-2640-FE0F": {
		Key:        "1F64D-1F3FD-200D-2640-FE0F",
		Value:      "🙍🏽‍♀️",
		Descriptor: "Woman Frowning: Medium Skin Tone",
	},
	"1F64D-1F3FD-200D-2642-FE0F": {
		Key:        "1F64D-1F3FD-200D-2642-FE0F",
		Value:      "🙍🏽‍♂️",
		Descriptor: "Man Frowning: Medium Skin Tone",
	},
	"1F64D-1F3FE": {
		Key:        "1F64D-1F3FE",
		Value:      "🙍🏾",
		Descriptor: "Person Frowning: Medium-Dark Skin Tone",
	},
	"1F64D-1F3FE-200D-2640-FE0F": {
		Key:        "1F64D-1F3FE-200D-2640-FE0F",
		Value:      "🙍🏾‍♀️",
		Descriptor: "Woman Frowning: Medium-Dark Skin Tone",
	},
	"1F64D-1F3FE-200D-2642-FE0F": {
		Key:        "1F64D-1F3FE-200D-2642-FE0F",
		Value:      "🙍🏾‍♂️",
		Descriptor: "Man Frowning: Medium-Dark Skin Tone",
	},
	"1F64D-1F3FF": {
		Key:        "1F64D-1F3FF",
		Value:      "🙍🏿",
		Descriptor: "Person Frowning: Dark Skin Tone",
	},
	"1F64D-1F3FF-200D-2640-FE0F": {
		Key:        "1F64D-1F3FF-200D-2640-FE0F",
		Value:      "🙍🏿‍♀️",
		Descriptor: "Woman Frowning: Dark Skin Tone",
	},
	"1F64D-1F3FF-200D-2642-FE0F": {
		Key:        "1F64D-1F3FF-200D-2642-FE0F",
		Value:      "🙍🏿‍♂️",
		Descriptor: "Man Frowning: Dark Skin Tone",
	},
	"1F64D-200D-2640-FE0F": {
		Key:        "1F64D-200D-2640-FE0F",
		Value:      "🙍‍♀️",
		Descriptor: "Woman Frowning",
	},
	"1F64D-200D-2642-FE0F": {
		Key:        "1F64D-200D-2642-FE0F",
		Value:      "🙍‍♂️",
		Descriptor: "Man Frowning",
	},
	"1F64E": {
		Key:        "1F64E",
		Value:      "🙎",
		Descriptor: "Person Pouting",
	},
	"1F64E-1F3FB": {
		Key:        "1F64E-1F3FB",
		Value:      "🙎🏻",
		Descriptor: "Person Pouting: Light Skin Tone",
	},
	"1F64E-1F3FB-200D-2640-FE0F": {
		Key:        "1F64E-1F3FB-200D-2640-FE0F",
		Value:      "🙎🏻‍♀️",
		Descriptor: "Woman Pouting: Light Skin Tone",
	},
	"1F64E-1F3FB-200D-2642-FE0F": {
		Key:        "1F64E-1F3FB-200D-2642-FE0F",
		Value:      "🙎🏻‍♂️",
		Descriptor: "Man Pouting: Light Skin Tone",
	},
	"1F64E-1F3FC": {
		Key:        "1F64E-1F3FC",
		Value:      "🙎🏼",
		Descriptor: "Person Pouting: Medium-Light Skin Tone",
	},
	"1F64E-1F3FC-200D-2640-FE0F": {
		Key:        "1F64E-1F3FC-200D-2640-FE0F",
		Value:      "🙎🏼‍♀️",
		Descriptor: "Woman Pouting: Medium-Light Skin Tone",
	},
	"1F64E-1F3FC-200D-2642-FE0F": {
		Key:        "1F64E-1F3FC-200D-2642-FE0F",
		Value:      "🙎🏼‍♂️",
		Descriptor: "Man Pouting: Medium-Light Skin Tone",
	},
	"1F64E-1F3FD": {
		Key:        "1F64E-1F3FD",
		Value:      "🙎🏽",
		Descriptor: "Person Pouting: Medium Skin Tone",
	},
	"1F64E-1F3FD-200D-2640-FE0F": {
		Key:        "1F64E-1F3FD-200D-2640-FE0F",
		Value:      "🙎🏽‍♀️",
		Descriptor: "Woman Pouting: Medium Skin Tone",
	},
	"1F64E-1F3FD-200D-2642-FE0F": {
		Key:        "1F64E-1F3FD-200D-2642-FE0F",
		Value:      "🙎🏽‍♂️",
		Descriptor: "Man Pouting: Medium Skin Tone",
	},
	"1F64E-1F3FE": {
		Key:        "1F64E-1F3FE",
		Value:      "🙎🏾",
		Descriptor: "Person Pouting: Medium-Dark Skin Tone",
	},
	"1F64E-1F3FE-200D-2640-FE0F": {
		Key:        "1F64E-1F3FE-200D-2640-FE0F",
		Value:      "🙎🏾‍♀️",
		Descriptor: "Woman Pouting: Medium-Dark Skin Tone",
	},
	"1F64E-1F3FE-200D-2642-FE0F": {
		Key:        "1F64E-1F3FE-200D-2642-FE0F",
		Value:      "🙎🏾‍♂️",
		Descriptor: "Man Pouting: Medium-Dark Skin Tone",
	},
	"1F64E-1F3FF": {
		Key:        "1F64E-1F3FF",
		Value:      "🙎🏿",
		Descriptor: "Person Pouting: Dark Skin Tone",
	},
	"1F64E-1F3FF-200D-2640-FE0F": {
		Key:        "1F64E-1F3FF-200D-2640-FE0F",
		Value:      "🙎🏿‍♀️",
		Descriptor: "Woman Pouting: Dark Skin Tone",
	},
	"1F64E-1F3FF-200D-2642-FE0F": {
		Key:        "1F64E-1F3FF-200D-2642-FE0F",
		Value:      "🙎🏿‍♂️",
		Descriptor: "Man Pouting: Dark Skin Tone",
	},
	"1F64E-200D-2640-FE0F": {
		Key:        "1F64E-200D-2640-FE0F",
		Value:      "🙎‍♀️",
		Descriptor: "Woman Pouting",
	},
	"1F64E-200D-2642-FE0F": {
		Key:        "1F64E-200D-2642-FE0F",
		Value:      "🙎‍♂️",
		Descriptor: "Man Pouting",
	},
	"1F64F": {
		Key:        "1F64F",
		Value:      "🙏",
		Descriptor: "Folded Hands",
	},
	"1F64F-1F3FB": {
		Key:        "1F64F-1F3FB",
		Value:      "🙏🏻",
		Descriptor: "Folded Hands: Light Skin Tone",
	},
	"1F64F-1F3FC": {
		Key:        "1F64F-1F3FC",
		Value:      "🙏🏼",
		Descriptor: "Folded Hands: Medium-Light Skin Tone",
	},
	"1F64F-1F3FD": {
		Key:        "1F64F-1F3FD",
		Value:      "🙏🏽",
		Descriptor: "Folded Hands: Medium Skin Tone",
	},
	"1F64F-1F3FE": {
		Key:        "1F64F-1F3FE",
		Value:      "🙏🏾",
		Descriptor: "Folded Hands: Medium-Dark Skin Tone",
	},
	"1F64F-1F3FF": {
		Key:        "1F64F-1F3FF",
		Value:      "🙏🏿",
		Descriptor: "Folded Hands: Dark Skin Tone",
	},
	"1F680": {
		Key:        "1F680",
		Value:      "🚀",
		Descriptor: "Rocket",
	},
	"1F681": {
		Key:        "1F681",
		Value:      "🚁",
		Descriptor: "Helicopter",
	},
	"1F682": {
		Key:        "1F682",
		Value:      "🚂",
		Descriptor: "Locomotive",
	},
	"1F683": {
		Key:        "1F683",
		Value:      "🚃",
		Descriptor: "Railway Car",
	},
	"1F684": {
		Key:        "1F684",
		Value:      "🚄",
		Descriptor: "High-Speed Train",
	},
	"1F685": {
		Key:        "1F685",
		Value:      "🚅",
		Descriptor: "Bullet Train",
	},
	"1F686": {
		Key:        "1F686",
		Value:      "🚆",
		Descriptor: "Train",
	},
	"1F687": {
		Key:        "1F687",
		Value:      "🚇",
		Descriptor: "Metro",
	},
	"1F688": {
		Key:        "1F688",
		Value:      "🚈",
		Descriptor: "Light Rail",
	},
	"1F689": {
		Key:        "1F689",
		Value:      "🚉",
		Descriptor: "Station",
	},
	"1F68A": {
		Key:        "1F68A",
		Value:      "🚊",
		Descriptor: "Tram",
	},
	"1F68B": {
		Key:        "1F68B",
		Value:      "🚋",
		Descriptor: "Tram Car",
	},
	"1F68C": {
		Key:        "1F68C",
		Value:      "🚌",
		Descriptor: "Bus",
	},
	"1F68D": {
		Key:        "1F68D",
		Value:      "🚍",
		Descriptor: "Oncoming Bus",
	},
	"1F68E": {
		Key:        "1F68E",
		Value:      "🚎",
		Descriptor: "Trolleybus",
	},
	"1F68F": {
		Key:        "1F68F",
		Value:      "🚏",
		Descriptor: "Bus Stop",
	},
	"1F690": {
		Key:        "1F690",
		Value:      "🚐",
		Descriptor: "Minibus",
	},
	"1F691": {
		Key:        "1F691",
		Value:      "🚑",
		Descriptor: "Ambulance",
	},
	"1F692": {
		Key:        "1F692",
		Value:      "🚒",
		Descriptor: "Fire Engine",
	},
	"1F693": {
		Key:        "1F693",
		Value:      "🚓",
		Descriptor: "Police Car",
	},
	"1F694": {
		Key:        "1F694",
		Value:      "🚔",
		Descriptor: "Oncoming Police Car",
	},
	"1F695": {
		Key:        "1F695",
		Value:      "🚕",
		Descriptor: "Taxi",
	},
	"1F696": {
		Key:        "1F696",
		Value:      "🚖",
		Descriptor: "Oncoming Taxi",
	},
	"1F697": {
		Key:        "1F697",
		Value:      "🚗",
		Descriptor: "Automobile",
	},
	"1F698": {
		Key:        "1F698",
		Value:      "🚘",
		Descriptor: "Oncoming Automobile",
	},
	"1F699": {
		Key:        "1F699",
		Value:      "🚙",
		Descriptor: "Sport Utility Vehicle",
	},
	"1F69A": {
		Key:        "1F69A",
		Value:      "🚚",
		Descriptor: "Delivery Truck",
	},
	"1F69B": {
		Key:        "1F69B",
		Value:      "🚛",
		Descriptor: "Articulated Lorry",
	},
	"1F69C": {
		Key:        "1F69C",
		Value:      "🚜",
		Descriptor: "Tractor",
	},
	"1F69D": {
		Key:        "1F69D",
		Value:      "🚝",
		Descriptor: "Monorail",
	},
	"1F69E": {
		Key:        "1F69E",
		Value:      "🚞",
		Descriptor: "Mountain Railway",
	},
	"1F69F": {
		Key:        "1F69F",
		Value:      "🚟",
		Descriptor: "Suspension Railway",
	},
	"1F6A0": {
		Key:        "1F6A0",
		Value:      "🚠",
		Descriptor: "Mountain Cableway",
	},
	"1F6A1": {
		Key:        "1F6A1",
		Value:      "🚡",
		Descriptor: "Aerial Tramway",
	},
	"1F6A2": {
		Key:        "1F6A2",
		Value:      "🚢",
		Descriptor: "Ship",
	},
	"1F6A3": {
		Key:        "1F6A3",
		Value:      "🚣",
		Descriptor: "Person Rowing Boat",
	},
	"1F6A3-1F3FB": {
		Key:        "1F6A3-1F3FB",
		Value:      "🚣🏻",
		Descriptor: "Person Rowing Boat: Light Skin Tone",
	},
	"1F6A3-1F3FB-200D-2640-FE0F": {
		Key:        "1F6A3-1F3FB-200D-2640-FE0F",
		Value:      "🚣🏻‍♀️",
		Descriptor: "Woman Rowing Boat: Light Skin Tone",
	},
	"1F6A3-1F3FB-200D-2642-FE0F": {
		Key:        "1F6A3-1F3FB-200D-2642-FE0F",
		Value:      "🚣🏻‍♂️",
		Descriptor: "Man Rowing Boat: Light Skin Tone",
	},
	"1F6A3-1F3FC": {
		Key:        "1F6A3-1F3FC",
		Value:      "🚣🏼",
		Descriptor: "Person Rowing Boat: Medium-Light Skin Tone",
	},
	"1F6A3-1F3FC-200D-2640-FE0F": {
		Key:        "1F6A3-1F3FC-200D-2640-FE0F",
		Value:      "🚣🏼‍♀️",
		Descriptor: "Woman Rowing Boat: Medium-Light Skin Tone",
	},
	"1F6A3-1F3FC-200D-2642-FE0F": {
		Key:        "1F6A3-1F3FC-200D-2642-FE0F",
		Value:      "🚣🏼‍♂️",
		Descriptor: "Man Rowing Boat: Medium-Light Skin Tone",
	},
	"1F6A3-1F3FD": {
		Key:        "1F6A3-1F3FD",
		Value:      "🚣🏽",
		Descriptor: "Person Rowing Boat: Medium Skin Tone",
	},
	"1F6A3-1F3FD-200D-2640-FE0F": {
		Key:        "1F6A3-1F3FD-200D-2640-FE0F",
		Value:      "🚣🏽‍♀️",
		Descriptor: "Woman Rowing Boat: Medium Skin Tone",
	},
	"1F6A3-1F3FD-200D-2642-FE0F": {
		Key:        "1F6A3-1F3FD-200D-2642-FE0F",
		Value:      "🚣🏽‍♂️",
		Descriptor: "Man Rowing Boat: Medium Skin Tone",
	},
	"1F6A3-1F3FE": {
		Key:        "1F6A3-1F3FE",
		Value:      "🚣🏾",
		Descriptor: "Person Rowing Boat: Medium-Dark Skin Tone",
	},
	"1F6A3-1F3FE-200D-2640-FE0F": {
		Key:        "1F6A3-1F3FE-200D-2640-FE0F",
		Value:      "🚣🏾‍♀️",
		Descriptor: "Woman Rowing Boat: Medium-Dark Skin Tone",
	},
	"1F6A3-1F3FE-200D-2642-FE0F": {
		Key:        "1F6A3-1F3FE-200D-2642-FE0F",
		Value:      "🚣🏾‍♂️",
		Descriptor: "Man Rowing Boat: Medium-Dark Skin Tone",
	},
	"1F6A3-1F3FF": {
		Key:        "1F6A3-1F3FF",
		Value:      "🚣🏿",
		Descriptor: "Person Rowing Boat: Dark Skin Tone",
	},
	"1F6A3-1F3FF-200D-2640-FE0F": {
		Key:        "1F6A3-1F3FF-200D-2640-FE0F",
		Value:      "🚣🏿‍♀️",
		Descriptor: "Woman Rowing Boat: Dark Skin Tone",
	},
	"1F6A3-1F3FF-200D-2642-FE0F": {
		Key:        "1F6A3-1F3FF-200D-2642-FE0F",
		Value:      "🚣🏿‍♂️",
		Descriptor: "Man Rowing Boat: Dark Skin Tone",
	},
	"1F6A3-200D-2640-FE0F": {
		Key:        "1F6A3-200D-2640-FE0F",
		Value:      "🚣‍♀️",
		Descriptor: "Woman Rowing Boat",
	},
	"1F6A3-200D-2642-FE0F": {
		Key:        "1F6A3-200D-2642-FE0F",
		Value:      "🚣‍♂️",
		Descriptor: "Man Rowing Boat",
	},
	"1F6A4": {
		Key:        "1F6A4",
		Value:      "🚤",
		Descriptor: "Speedboat",
	},
	"1F6A5": {
		Key:        "1F6A5",
		Value:      "🚥",
		Descriptor: "Horizontal Traffic Light",
	},
	"1F6A6": {
		Key:        "1F6A6",
		Value:      "🚦",
		Descriptor: "Vertical Traffic Light",
	},
	"1F6A7": {
		Key:        "1F6A7",
		Value:      "🚧",
		Descriptor: "Construction",
	},
	"1F6A8": {
		Key:        "1F6A8",
		Value:      "🚨",
		Descriptor: "Police Car Light",
	},
	"1F6A9": {
		Key:        "1F6A9",
		Value:      "🚩",
		Descriptor: "Triangular Flag",
	},
	"1F6AA": {
		Key:        "1F6AA",
		Value:      "🚪",
		Descriptor: "Door",
	},
	"1F6AB": {
		Key:        "1F6AB",
		Value:      "🚫",
		Descriptor: "Prohibited",
	},
	"1F6AC": {
		Key:        "1F6AC",
		Value:      "🚬",
		Descriptor: "Cigarette",
	},
	"1F6AD": {
		Key:        "1F6AD",
		Value:      "🚭",
		Descriptor: "No Smoking",
	},
	"1F6AE": {
		Key:        "1F6AE",
		Value:      "🚮",
		Descriptor: "Litter in Bin Sign",
	},
	"1F6AF": {
		Key:        "1F6AF",
		Value:      "🚯",
		Descriptor: "No Littering",
	},
	"1F6B0": {
		Key:        "1F6B0",
		Value:      "🚰",
		Descriptor: "Potable Water",
	},
	"1F6B1": {
		Key:        "1F6B1",
		Value:      "🚱",
		Descriptor: "Non-Potable Water",
	},
	"1F6B2": {
		Key:        "1F6B2",
		Value:      "🚲",
		Descriptor: "Bicycle",
	},
	"1F6B3": {
		Key:        "1F6B3",
		Value:      "🚳",
		Descriptor: "No Bicycles",
	},
	"1F6B4": {
		Key:        "1F6B4",
		Value:      "🚴",
		Descriptor: "Person Biking",
	},
	"1F6B4-1F3FB": {
		Key:        "1F6B4-1F3FB",
		Value:      "🚴🏻",
		Descriptor: "Person Biking: Light Skin Tone",
	},
	"1F6B4-1F3FB-200D-2640-FE0F": {
		Key:        "1F6B4-1F3FB-200D-2640-FE0F",
		Value:      "🚴🏻‍♀️",
		Descriptor: "Woman Biking: Light Skin Tone",
	},
	"1F6B4-1F3FB-200D-2642-FE0F": {
		Key:        "1F6B4-1F3FB-200D-2642-FE0F",
		Value:      "🚴🏻‍♂️",
		Descriptor: "Man Biking: Light Skin Tone",
	},
	"1F6B4-1F3FC": {
		Key:        "1F6B4-1F3FC",
		Value:      "🚴🏼",
		Descriptor: "Person Biking: Medium-Light Skin Tone",
	},
	"1F6B4-1F3FC-200D-2640-FE0F": {
		Key:        "1F6B4-1F3FC-200D-2640-FE0F",
		Value:      "🚴🏼‍♀️",
		Descriptor: "Woman Biking: Medium-Light Skin Tone",
	},
	"1F6B4-1F3FC-200D-2642-FE0F": {
		Key:        "1F6B4-1F3FC-200D-2642-FE0F",
		Value:      "🚴🏼‍♂️",
		Descriptor: "Man Biking: Medium-Light Skin Tone",
	},
	"1F6B4-1F3FD": {
		Key:        "1F6B4-1F3FD",
		Value:      "🚴🏽",
		Descriptor: "Person Biking: Medium Skin Tone",
	},
	"1F6B4-1F3FD-200D-2640-FE0F": {
		Key:        "1F6B4-1F3FD-200D-2640-FE0F",
		Value:      "🚴🏽‍♀️",
		Descriptor: "Woman Biking: Medium Skin Tone",
	},
	"1F6B4-1F3FD-200D-2642-FE0F": {
		Key:        "1F6B4-1F3FD-200D-2642-FE0F",
		Value:      "🚴🏽‍♂️",
		Descriptor: "Man Biking: Medium Skin Tone",
	},
	"1F6B4-1F3FE": {
		Key:        "1F6B4-1F3FE",
		Value:      "🚴🏾",
		Descriptor: "Person Biking: Medium-Dark Skin Tone",
	},
	"1F6B4-1F3FE-200D-2640-FE0F": {
		Key:        "1F6B4-1F3FE-200D-2640-FE0F",
		Value:      "🚴🏾‍♀️",
		Descriptor: "Woman Biking: Medium-Dark Skin Tone",
	},
	"1F6B4-1F3FE-200D-2642-FE0F": {
		Key:        "1F6B4-1F3FE-200D-2642-FE0F",
		Value:      "🚴🏾‍♂️",
		Descriptor: "Man Biking: Medium-Dark Skin Tone",
	},
	"1F6B4-1F3FF": {
		Key:        "1F6B4-1F3FF",
		Value:      "🚴🏿",
		Descriptor: "Person Biking: Dark Skin Tone",
	},
	"1F6B4-1F3FF-200D-2640-FE0F": {
		Key:        "1F6B4-1F3FF-200D-2640-FE0F",
		Value:      "🚴🏿‍♀️",
		Descriptor: "Woman Biking: Dark Skin Tone",
	},
	"1F6B4-1F3FF-200D-2642-FE0F": {
		Key:        "1F6B4-1F3FF-200D-2642-FE0F",
		Value:      "🚴🏿‍♂️",
		Descriptor: "Man Biking: Dark Skin Tone",
	},
	"1F6B4-200D-2640-FE0F": {
		Key:        "1F6B4-200D-2640-FE0F",
		Value:      "🚴‍♀️",
		Descriptor: "Woman Biking",
	},
	"1F6B4-200D-2642-FE0F": {
		Key:        "1F6B4-200D-2642-FE0F",
		Value:      "🚴‍♂️",
		Descriptor: "Man Biking",
	},
	"1F6B5": {
		Key:        "1F6B5",
		Value:      "🚵",
		Descriptor: "Person Mountain Biking",
	},
	"1F6B5-1F3FB": {
		Key:        "1F6B5-1F3FB",
		Value:      "🚵🏻",
		Descriptor: "Person Mountain Biking: Light Skin Tone",
	},
	"1F6B5-1F3FB-200D-2640-FE0F": {
		Key:        "1F6B5-1F3FB-200D-2640-FE0F",
		Value:      "🚵🏻‍♀️",
		Descriptor: "Woman Mountain Biking: Light Skin Tone",
	},
	"1F6B5-1F3FB-200D-2642-FE0F": {
		Key:        "1F6B5-1F3FB-200D-2642-FE0F",
		Value:      "🚵🏻‍♂️",
		Descriptor: "Man Mountain Biking: Light Skin Tone",
	},
	"1F6B5-1F3FC": {
		Key:        "1F6B5-1F3FC",
		Value:      "🚵🏼",
		Descriptor: "Person Mountain Biking: Medium-Light Skin Tone",
	},
	"1F6B5-1F3FC-200D-2640-FE0F": {
		Key:        "1F6B5-1F3FC-200D-2640-FE0F",
		Value:      "🚵🏼‍♀️",
		Descriptor: "Woman Mountain Biking: Medium-Light Skin Tone",
	},
	"1F6B5-1F3FC-200D-2642-FE0F": {
		Key:        "1F6B5-1F3FC-200D-2642-FE0F",
		Value:      "🚵🏼‍♂️",
		Descriptor: "Man Mountain Biking: Medium-Light Skin Tone",
	},
	"1F6B5-1F3FD": {
		Key:        "1F6B5-1F3FD",
		Value:      "🚵🏽",
		Descriptor: "Person Mountain Biking: Medium Skin Tone",
	},
	"1F6B5-1F3FD-200D-2640-FE0F": {
		Key:        "1F6B5-1F3FD-200D-2640-FE0F",
		Value:      "🚵🏽‍♀️",
		Descriptor: "Woman Mountain Biking: Medium Skin Tone",
	},
	"1F6B5-1F3FD-200D-2642-FE0F": {
		Key:        "1F6B5-1F3FD-200D-2642-FE0F",
		Value:      "🚵🏽‍♂️",
		Descriptor: "Man Mountain Biking: Medium Skin Tone",
	},
	"1F6B5-1F3FE": {
		Key:        "1F6B5-1F3FE",
		Value:      "🚵🏾",
		Descriptor: "Person Mountain Biking: Medium-Dark Skin Tone",
	},
	"1F6B5-1F3FE-200D-2640-FE0F": {
		Key:        "1F6B5-1F3FE-200D-2640-FE0F",
		Value:      "🚵🏾‍♀️",
		Descriptor: "Woman Mountain Biking: Medium-Dark Skin Tone",
	},
	"1F6B5-1F3FE-200D-2642-FE0F": {
		Key:        "1F6B5-1F3FE-200D-2642-FE0F",
		Value:      "🚵🏾‍♂️",
		Descriptor: "Man Mountain Biking: Medium-Dark Skin Tone",
	},
	"1F6B5-1F3FF": {
		Key:        "1F6B5-1F3FF",
		Value:      "🚵🏿",
		Descriptor: "Person Mountain Biking: Dark Skin Tone",
	},
	"1F6B5-1F3FF-200D-2640-FE0F": {
		Key:        "1F6B5-1F3FF-200D-2640-FE0F",
		Value:      "🚵🏿‍♀️",
		Descriptor: "Woman Mountain Biking: Dark Skin Tone",
	},
	"1F6B5-1F3FF-200D-2642-FE0F": {
		Key:        "1F6B5-1F3FF-200D-2642-FE0F",
		Value:      "🚵🏿‍♂️",
		Descriptor: "Man Mountain Biking: Dark Skin Tone",
	},
	"1F6B5-200D-2640-FE0F": {
		Key:        "1F6B5-200D-2640-FE0F",
		Value:      "🚵‍♀️",
		Descriptor: "Woman Mountain Biking",
	},
	"1F6B5-200D-2642-FE0F": {
		Key:        "1F6B5-200D-2642-FE0F",
		Value:      "🚵‍♂️",
		Descriptor: "Man Mountain Biking",
	},
	"1F6B6": {
		Key:        "1F6B6",
		Value:      "🚶",
		Descriptor: "Person Walking",
	},
	"1F6B6-1F3FB": {
		Key:        "1F6B6-1F3FB",
		Value:      "🚶🏻",
		Descriptor: "Person Walking: Light Skin Tone",
	},
	"1F6B6-1F3FB-200D-2640-FE0F": {
		Key:        "1F6B6-1F3FB-200D-2640-FE0F",
		Value:      "🚶🏻‍♀️",
		Descriptor: "Woman Walking: Light Skin Tone",
	},
	"1F6B6-1F3FB-200D-2642-FE0F": {
		Key:        "1F6B6-1F3FB-200D-2642-FE0F",
		Value:      "🚶🏻‍♂️",
		Descriptor: "Man Walking: Light Skin Tone",
	},
	"1F6B6-1F3FC": {
		Key:        "1F6B6-1F3FC",
		Value:      "🚶🏼",
		Descriptor: "Person Walking: Medium-Light Skin Tone",
	},
	"1F6B6-1F3FC-200D-2640-FE0F": {
		Key:        "1F6B6-1F3FC-200D-2640-FE0F",
		Value:      "🚶🏼‍♀️",
		Descriptor: "Woman Walking: Medium-Light Skin Tone",
	},
	"1F6B6-1F3FC-200D-2642-FE0F": {
		Key:        "1F6B6-1F3FC-200D-2642-FE0F",
		Value:      "🚶🏼‍♂️",
		Descriptor: "Man Walking: Medium-Light Skin Tone",
	},
	"1F6B6-1F3FD": {
		Key:        "1F6B6-1F3FD",
		Value:      "🚶🏽",
		Descriptor: "Person Walking: Medium Skin Tone",
	},
	"1F6B6-1F3FD-200D-2640-FE0F": {
		Key:        "1F6B6-1F3FD-200D-2640-FE0F",
		Value:      "🚶🏽‍♀️",
		Descriptor: "Woman Walking: Medium Skin Tone",
	},
	"1F6B6-1F3FD-200D-2642-FE0F": {
		Key:        "1F6B6-1F3FD-200D-2642-FE0F",
		Value:      "🚶🏽‍♂️",
		Descriptor: "Man Walking: Medium Skin Tone",
	},
	"1F6B6-1F3FE": {
		Key:        "1F6B6-1F3FE",
		Value:      "🚶🏾",
		Descriptor: "Person Walking: Medium-Dark Skin Tone",
	},
	"1F6B6-1F3FE-200D-2640-FE0F": {
		Key:        "1F6B6-1F3FE-200D-2640-FE0F",
		Value:      "🚶🏾‍♀️",
		Descriptor: "Woman Walking: Medium-Dark Skin Tone",
	},
	"1F6B6-1F3FE-200D-2642-FE0F": {
		Key:        "1F6B6-1F3FE-200D-2642-FE0F",
		Value:      "🚶🏾‍♂️",
		Descriptor: "Man Walking: Medium-Dark Skin Tone",
	},
	"1F6B6-1F3FF": {
		Key:        "1F6B6-1F3FF",
		Value:      "🚶🏿",
		Descriptor: "Person Walking: Dark Skin Tone",
	},
	"1F6B6-1F3FF-200D-2640-FE0F": {
		Key:        "1F6B6-1F3FF-200D-2640-FE0F",
		Value:      "🚶🏿‍♀️",
		Descriptor: "Woman Walking: Dark Skin Tone",
	},
	"1F6B6-1F3FF-200D-2642-FE0F": {
		Key:        "1F6B6-1F3FF-200D-2642-FE0F",
		Value:      "🚶🏿‍♂️",
		Descriptor: "Man Walking: Dark Skin Tone",
	},
	"1F6B6-200D-2640-FE0F": {
		Key:        "1F6B6-200D-2640-FE0F",
		Value:      "🚶‍♀️",
		Descriptor: "Woman Walking",
	},
	"1F6B6-200D-2642-FE0F": {
		Key:        "1F6B6-200D-2642-FE0F",
		Value:      "🚶‍♂️",
		Descriptor: "Man Walking",
	},
	"1F6B7": {
		Key:        "1F6B7",
		Value:      "🚷",
		Descriptor: "No Pedestrians",
	},
	"1F6B8": {
		Key:        "1F6B8",
		Value:      "🚸",
		Descriptor: "Children Crossing",
	},
	"1F6B9": {
		Key:        "1F6B9",
		Value:      "🚹",
		Descriptor: "Men’s Room",
	},
	"1F6BA": {
		Key:        "1F6BA",
		Value:      "🚺",
		Descriptor: "Women’s Room",
	},
	"1F6BB": {
		Key:        "1F6BB",
		Value:      "🚻",
		Descriptor: "Restroom",
	},
	"1F6BC": {
		Key:        "1F6BC",
		Value:      "🚼",
		Descriptor: "Baby Symbol",
	},
	"1F6BD": {
		Key:        "1F6BD",
		Value:      "🚽",
		Descriptor: "Toilet",
	},
	"1F6BE": {
		Key:        "1F6BE",
		Value:      "🚾",
		Descriptor: "Water Closet",
	},
	"1F6BF": {
		Key:        "1F6BF",
		Value:      "🚿",
		Descriptor: "Shower",
	},
	"1F6C0": {
		Key:        "1F6C0",
		Value:      "🛀",
		Descriptor: "Person Taking Bath",
	},
	"1F6C0-1F3FB": {
		Key:        "1F6C0-1F3FB",
		Value:      "🛀🏻",
		Descriptor: "Person Taking Bath: Light Skin Tone",
	},
	"1F6C0-1F3FC": {
		Key:        "1F6C0-1F3FC",
		Value:      "🛀🏼",
		Descriptor: "Person Taking Bath: Medium-Light Skin Tone",
	},
	"1F6C0-1F3FD": {
		Key:        "1F6C0-1F3FD",
		Value:      "🛀🏽",
		Descriptor: "Person Taking Bath: Medium Skin Tone",
	},
	"1F6C0-1F3FE": {
		Key:        "1F6C0-1F3FE",
		Value:      "🛀🏾",
		Descriptor: "Person Taking Bath: Medium-Dark Skin Tone",
	},
	"1F6C0-1F3FF": {
		Key:        "1F6C0-1F3FF",
		Value:      "🛀🏿",
		Descriptor: "Person Taking Bath: Dark Skin Tone",
	},
	"1F6C1": {
		Key:        "1F6C1",
		Value:      "🛁",
		Descriptor: "Bathtub",
	},
	"1F6C2": {
		Key:        "1F6C2",
		Value:      "🛂",
		Descriptor: "Passport Control",
	},
	"1F6C3": {
		Key:        "1F6C3",
		Value:      "🛃",
		Descriptor: "Customs",
	},
	"1F6C4": {
		Key:        "1F6C4",
		Value:      "🛄",
		Descriptor: "Baggage Claim",
	},
	"1F6C5": {
		Key:        "1F6C5",
		Value:      "🛅",
		Descriptor: "Left Luggage",
	},
	"1F6CB-FE0F": {
		Key:        "1F6CB-FE0F",
		Value:      "🛋️",
		Descriptor: "Couch and Lamp",
	},
	"1F6CC": {
		Key:        "1F6CC",
		Value:      "🛌",
		Descriptor: "Person in Bed",
	},
	"1F6CD-FE0F": {
		Key:        "1F6CD-FE0F",
		Value:      "🛍️",
		Descriptor: "Shopping Bags",
	},
	"1F6CE-FE0F": {
		Key:        "1F6CE-FE0F",
		Value:      "🛎️",
		Descriptor: "Bellhop Bell",
	},
	"1F6CF-FE0F": {
		Key:        "1F6CF-FE0F",
		Value:      "🛏️",
		Descriptor: "Bed",
	},
	"1F6D0": {
		Key:        "1F6D0",
		Value:      "🛐",
		Descriptor: "Place of Worship",
	},
	"1F6D1": {
		Key:        "1F6D1",
		Value:      "🛑",
		Descriptor: "Stop Sign",
	},
	"1F6D2": {
		Key:        "1F6D2",
		Value:      "🛒",
		Descriptor: "Shopping Cart",
	},
	"1F6D5": {
		Key:        "1F6D5",
		Value:      "🛕",
		Descriptor: "Hindu Temple",
	},
	"1F6E0-FE0F": {
		Key:        "1F6E0-FE0F",
		Value:      "🛠️",
		Descriptor: "Hammer and Wrench",
	},
	"1F6E1-FE0F": {
		Key:        "1F6E1-FE0F",
		Value:      "🛡️",
		Descriptor: "Shield",
	},
	"1F6E2-FE0F": {
		Key:        "1F6E2-FE0F",
		Value:      "🛢️",
		Descriptor: "Oil Drum",
	},
	"1F6E3-FE0F": {
		Key:        "1F6E3-FE0F",
		Value:      "🛣️",
		Descriptor: "Motorway",
	},
	"1F6E4-FE0F": {
		Key:        "1F6E4-FE0F",
		Value:      "🛤️",
		Descriptor: "Railway Track",
	},
	"1F6E5-FE0F": {
		Key:        "1F6E5-FE0F",
		Value:      "🛥️",
		Descriptor: "Motor Boat",
	},
	"1F6E9-FE0F": {
		Key:        "1F6E9-FE0F",
		Value:      "🛩️",
		Descriptor: "Small Airplane",
	},
	"1F6EB": {
		Key:        "1F6EB",
		Value:      "🛫",
		Descriptor: "Airplane Departure",
	},
	"1F6EC": {
		Key:        "1F6EC",
		Value:      "🛬",
		Descriptor: "Airplane Arrival",
	},
	"1F6F0-FE0F": {
		Key:        "1F6F0-FE0F",
		Value:      "🛰️",
		Descriptor: "Satellite",
	},
	"1F6F3-FE0F": {
		Key:        "1F6F3-FE0F",
		Value:      "🛳️",
		Descriptor: "Passenger Ship",
	},
	"1F6F4": {
		Key:        "1F6F4",
		Value:      "🛴",
		Descriptor: "Kick Scooter",
	},
	"1F6F5": {
		Key:        "1F6F5",
		Value:      "🛵",
		Descriptor: "Motor Scooter",
	},
	"1F6F6": {
		Key:        "1F6F6",
		Value:      "🛶",
		Descriptor: "Canoe",
	},
	"1F6F7": {
		Key:        "1F6F7",
		Value:      "🛷",
		Descriptor: "Sled",
	},
	"1F6F8": {
		Key:        "1F6F8",
		Value:      "🛸",
		Descriptor: "Flying Saucer",
	},
	"1F6F9": {
		Key:        "1F6F9",
		Value:      "🛹",
		Descriptor: "Skateboard",
	},
	"1F6FA": {
		Key:        "1F6FA",
		Value:      "🛺",
		Descriptor: "Auto Rickshaw",
	},
	"1F7E0": {
		Key:        "1F7E0",
		Value:      "🟠",
		Descriptor: "Orange Circle",
	},
	"1F7E1": {
		Key:        "1F7E1",
		Value:      "🟡",
		Descriptor: "Yellow Circle",
	},
	"1F7E2": {
		Key:        "1F7E2",
		Value:      "🟢",
		Descriptor: "Green Circle",
	},
	"1F7E3": {
		Key:        "1F7E3",
		Value:      "🟣",
		Descriptor: "Purple Circle",
	},
	"1F7E4": {
		Key:        "1F7E4",
		Value:      "🟤",
		Descriptor: "Brown Circle",
	},
	"1F7E5": {
		Key:        "1F7E5",
		Value:      "🟥",
		Descriptor: "Red Square",
	},
	"1F7E6": {
		Key:        "1F7E6",
		Value:      "🟦",
		Descriptor: "Blue Square",
	},
	"1F7E7": {
		Key:        "1F7E7",
		Value:      "🟧",
		Descriptor: "Orange Square",
	},
	"1F7E8": {
		Key:        "1F7E8",
		Value:      "🟨",
		Descriptor: "Yellow Square",
	},
	"1F7E9": {
		Key:        "1F7E9",
		Value:      "🟩",
		Descriptor: "Green Square",
	},
	"1F7EA": {
		Key:        "1F7EA",
		Value:      "🟪",
		Descriptor: "Purple Square",
	},
	"1F7EB": {
		Key:        "1F7EB",
		Value:      "🟫",
		Descriptor: "Brown Square",
	},
	"1F90D": {
		Key:        "1F90D",
		Value:      "🤍",
		Descriptor: "White Heart",
	},
	"1F90E": {
		Key:        "1F90E",
		Value:      "🤎",
		Descriptor: "Brown Heart",
	},
	"1F90F": {
		Key:        "1F90F",
		Value:      "🤏",
		Descriptor: "Pinching Hand",
	},
	"1F90F-1F3FB": {
		Key:        "1F90F-1F3FB",
		Value:      "🤏🏻",
		Descriptor: "Pinching Hand: Light Skin Tone",
	},
	"1F90F-1F3FC": {
		Key:        "1F90F-1F3FC",
		Value:      "🤏🏼",
		Descriptor: "Pinching Hand: Medium-Light Skin Tone",
	},
	"1F90F-1F3FD": {
		Key:        "1F90F-1F3FD",
		Value:      "🤏🏽",
		Descriptor: "Pinching Hand: Medium Skin Tone",
	},
	"1F90F-1F3FE": {
		Key:        "1F90F-1F3FE",
		Value:      "🤏🏾",
		Descriptor: "Pinching Hand: Medium-Dark Skin Tone",
	},
	"1F90F-1F3FF": {
		Key:        "1F90F-1F3FF",
		Value:      "🤏🏿",
		Descriptor: "Pinching Hand: Dark Skin Tone",
	},
	"1F910": {
		Key:        "1F910",
		Value:      "🤐",
		Descriptor: "Zipper-Mouth Face",
	},
	"1F911": {
		Key:        "1F911",
		Value:      "🤑",
		Descriptor: "Money-Mouth Face",
	},
	"1F912": {
		Key:        "1F912",
		Value:      "🤒",
		Descriptor: "Face With Thermometer",
	},
	"1F913": {
		Key:        "1F913",
		Value:      "🤓",
		Descriptor: "Nerd Face",
	},
	"1F914": {
		Key:        "1F914",
		Value:      "🤔",
		Descriptor: "Thinking Face",
	},
	"1F915": {
		Key:        "1F915",
		Value:      "🤕",
		Descriptor: "Face With Head-Bandage",
	},
	"1F916": {
		Key:        "1F916",
		Value:      "🤖",
		Descriptor: "Robot",
	},
	"1F917": {
		Key:        "1F917",
		Value:      "🤗",
		Descriptor: "Hugging Face",
	},
	"1F918": {
		Key:        "1F918",
		Value:      "🤘",
		Descriptor: "Sign of the Horns",
	},
	"1F918-1F3FB": {
		Key:        "1F918-1F3FB",
		Value:      "🤘🏻",
		Descriptor: "Sign of the Horns: Light Skin Tone",
	},
	"1F918-1F3FC": {
		Key:        "1F918-1F3FC",
		Value:      "🤘🏼",
		Descriptor: "Sign of the Horns: Medium-Light Skin Tone",
	},
	"1F918-1F3FD": {
		Key:        "1F918-1F3FD",
		Value:      "🤘🏽",
		Descriptor: "Sign of the Horns: Medium Skin Tone",
	},
	"1F918-1F3FE": {
		Key:        "1F918-1F3FE",
		Value:      "🤘🏾",
		Descriptor: "Sign of the Horns: Medium-Dark Skin Tone",
	},
	"1F918-1F3FF": {
		Key:        "1F918-1F3FF",
		Value:      "🤘🏿",
		Descriptor: "Sign of the Horns: Dark Skin Tone",
	},
	"1F919": {
		Key:        "1F919",
		Value:      "🤙",
		Descriptor: "Call Me Hand",
	},
	"1F919-1F3FB": {
		Key:        "1F919-1F3FB",
		Value:      "🤙🏻",
		Descriptor: "Call Me Hand: Light Skin Tone",
	},
	"1F919-1F3FC": {
		Key:        "1F919-1F3FC",
		Value:      "🤙🏼",
		Descriptor: "Call Me Hand: Medium-Light Skin Tone",
	},
	"1F919-1F3FD": {
		Key:        "1F919-1F3FD",
		Value:      "🤙🏽",
		Descriptor: "Call Me Hand: Medium Skin Tone",
	},
	"1F919-1F3FE": {
		Key:        "1F919-1F3FE",
		Value:      "🤙🏾",
		Descriptor: "Call Me Hand: Medium-Dark Skin Tone",
	},
	"1F919-1F3FF": {
		Key:        "1F919-1F3FF",
		Value:      "🤙🏿",
		Descriptor: "Call Me Hand: Dark Skin Tone",
	},
	"1F91A": {
		Key:        "1F91A",
		Value:      "🤚",
		Descriptor: "Raised Back of Hand",
	},
	"1F91A-1F3FB": {
		Key:        "1F91A-1F3FB",
		Value:      "🤚🏻",
		Descriptor: "Raised Back of Hand: Light Skin Tone",
	},
	"1F91A-1F3FC": {
		Key:        "1F91A-1F3FC",
		Value:      "🤚🏼",
		Descriptor: "Raised Back of Hand: Medium-Light Skin Tone",
	},
	"1F91A-1F3FD": {
		Key:        "1F91A-1F3FD",
		Value:      "🤚🏽",
		Descriptor: "Raised Back of Hand: Medium Skin Tone",
	},
	"1F91A-1F3FE": {
		Key:        "1F91A-1F3FE",
		Value:      "🤚🏾",
		Descriptor: "Raised Back of Hand: Medium-Dark Skin Tone",
	},
	"1F91A-1F3FF": {
		Key:        "1F91A-1F3FF",
		Value:      "🤚🏿",
		Descriptor: "Raised Back of Hand: Dark Skin Tone",
	},
	"1F91B": {
		Key:        "1F91B",
		Value:      "🤛",
		Descriptor: "Left-Facing Fist",
	},
	"1F91B-1F3FB": {
		Key:        "1F91B-1F3FB",
		Value:      "🤛🏻",
		Descriptor: "Left-Facing Fist: Light Skin Tone",
	},
	"1F91B-1F3FC": {
		Key:        "1F91B-1F3FC",
		Value:      "🤛🏼",
		Descriptor: "Left-Facing Fist: Medium-Light Skin Tone",
	},
	"1F91B-1F3FD": {
		Key:        "1F91B-1F3FD",
		Value:      "🤛🏽",
		Descriptor: "Left-Facing Fist: Medium Skin Tone",
	},
	"1F91B-1F3FE": {
		Key:        "1F91B-1F3FE",
		Value:      "🤛🏾",
		Descriptor: "Left-Facing Fist: Medium-Dark Skin Tone",
	},
	"1F91B-1F3FF": {
		Key:        "1F91B-1F3FF",
		Value:      "🤛🏿",
		Descriptor: "Left-Facing Fist: Dark Skin Tone",
	},
	"1F91C": {
		Key:        "1F91C",
		Value:      "🤜",
		Descriptor: "Right-Facing Fist",
	},
	"1F91C-1F3FB": {
		Key:        "1F91C-1F3FB",
		Value:      "🤜🏻",
		Descriptor: "Right-Facing Fist: Light Skin Tone",
	},
	"1F91C-1F3FC": {
		Key:        "1F91C-1F3FC",
		Value:      "🤜🏼",
		Descriptor: "Right-Facing Fist: Medium-Light Skin Tone",
	},
	"1F91C-1F3FD": {
		Key:        "1F91C-1F3FD",
		Value:      "🤜🏽",
		Descriptor: "Right-Facing Fist: Medium Skin Tone",
	},
	"1F91C-1F3FE": {
		Key:        "1F91C-1F3FE",
		Value:      "🤜🏾",
		Descriptor: "Right-Facing Fist: Medium-Dark Skin Tone",
	},
	"1F91C-1F3FF": {
		Key:        "1F91C-1F3FF",
		Value:      "🤜🏿",
		Descriptor: "Right-Facing Fist: Dark Skin Tone",
	},
	"1F91D": {
		Key:        "1F91D",
		Value:      "🤝",
		Descriptor: "Handshake",
	},
	"1F91E": {
		Key:        "1F91E",
		Value:      "🤞",
		Descriptor: "Crossed Fingers",
	},
	"1F91E-1F3FB": {
		Key:        "1F91E-1F3FB",
		Value:      "🤞🏻",
		Descriptor: "Crossed Fingers: Light Skin Tone",
	},
	"1F91E-1F3FC": {
		Key:        "1F91E-1F3FC",
		Value:      "🤞🏼",
		Descriptor: "Crossed Fingers: Medium-Light Skin Tone",
	},
	"1F91E-1F3FD": {
		Key:        "1F91E-1F3FD",
		Value:      "🤞🏽",
		Descriptor: "Crossed Fingers: Medium Skin Tone",
	},
	"1F91E-1F3FE": {
		Key:        "1F91E-1F3FE",
		Value:      "🤞🏾",
		Descriptor: "Crossed Fingers: Medium-Dark Skin Tone",
	},
	"1F91E-1F3FF": {
		Key:        "1F91E-1F3FF",
		Value:      "🤞🏿",
		Descriptor: "Crossed Fingers: Dark Skin Tone",
	},
	"1F91F": {
		Key:        "1F91F",
		Value:      "🤟",
		Descriptor: "Love-You Gesture",
	},
	"1F91F-1F3FB": {
		Key:        "1F91F-1F3FB",
		Value:      "🤟🏻",
		Descriptor: "Love-You Gesture: Light Skin Tone",
	},
	"1F91F-1F3FC": {
		Key:        "1F91F-1F3FC",
		Value:      "🤟🏼",
		Descriptor: "Love-You Gesture: Medium-Light Skin Tone",
	},
	"1F91F-1F3FD": {
		Key:        "1F91F-1F3FD",
		Value:      "🤟🏽",
		Descriptor: "Love-You Gesture: Medium Skin Tone",
	},
	"1F91F-1F3FE": {
		Key:        "1F91F-1F3FE",
		Value:      "🤟🏾",
		Descriptor: "Love-You Gesture: Medium-Dark Skin Tone",
	},
	"1F91F-1F3FF": {
		Key:        "1F91F-1F3FF",
		Value:      "🤟🏿",
		Descriptor: "Love-You Gesture: Dark Skin Tone",
	},
	"1F920": {
		Key:        "1F920",
		Value:      "🤠",
		Descriptor: "Cowboy Hat Face",
	},
	"1F921": {
		Key:        "1F921",
		Value:      "🤡",
		Descriptor: "Clown Face",
	},
	"1F922": {
		Key:        "1F922",
		Value:      "🤢",
		Descriptor: "Nauseated Face",
	},
	"1F923": {
		Key:        "1F923",
		Value:      "🤣",
		Descriptor: "Rolling on the Floor Laughing",
	},
	"1F924": {
		Key:        "1F924",
		Value:      "🤤",
		Descriptor: "Drooling Face",
	},
	"1F925": {
		Key:        "1F925",
		Value:      "🤥",
		Descriptor: "Lying Face",
	},
	"1F926": {
		Key:        "1F926",
		Value:      "🤦",
		Descriptor: "Person Facepalming",
	},
	"1F926-1F3FB": {
		Key:        "1F926-1F3FB",
		Value:      "🤦🏻",
		Descriptor: "Person Facepalming: Light Skin Tone",
	},
	"1F926-1F3FB-200D-2640-FE0F": {
		Key:        "1F926-1F3FB-200D-2640-FE0F",
		Value:      "🤦🏻‍♀️",
		Descriptor: "Woman Facepalming: Light Skin Tone",
	},
	"1F926-1F3FB-200D-2642-FE0F": {
		Key:        "1F926-1F3FB-200D-2642-FE0F",
		Value:      "🤦🏻‍♂️",
		Descriptor: "Man Facepalming: Light Skin Tone",
	},
	"1F926-1F3FC": {
		Key:        "1F926-1F3FC",
		Value:      "🤦🏼",
		Descriptor: "Person Facepalming: Medium-Light Skin Tone",
	},
	"1F926-1F3FC-200D-2640-FE0F": {
		Key:        "1F926-1F3FC-200D-2640-FE0F",
		Value:      "🤦🏼‍♀️",
		Descriptor: "Woman Facepalming: Medium-Light Skin Tone",
	},
	"1F926-1F3FC-200D-2642-FE0F": {
		Key:        "1F926-1F3FC-200D-2642-FE0F",
		Value:      "🤦🏼‍♂️",
		Descriptor: "Man Facepalming: Medium-Light Skin Tone",
	},
	"1F926-1F3FD": {
		Key:        "1F926-1F3FD",
		Value:      "🤦🏽",
		Descriptor: "Person Facepalming: Medium Skin Tone",
	},
	"1F926-1F3FD-200D-2640-FE0F": {
		Key:        "1F926-1F3FD-200D-2640-FE0F",
		Value:      "🤦🏽‍♀️",
		Descriptor: "Woman Facepalming: Medium Skin Tone",
	},
	"1F926-1F3FD-200D-2642-FE0F": {
		Key:        "1F926-1F3FD-200D-2642-FE0F",
		Value:      "🤦🏽‍♂️",
		Descriptor: "Man Facepalming: Medium Skin Tone",
	},
	"1F926-1F3FE": {
		Key:        "1F926-1F3FE",
		Value:      "🤦🏾",
		Descriptor: "Person Facepalming: Medium-Dark Skin Tone",
	},
	"1F926-1F3FE-200D-2640-FE0F": {
		Key:        "1F926-1F3FE-200D-2640-FE0F",
		Value:      "🤦🏾‍♀️",
		Descriptor: "Woman Facepalming: Medium-Dark Skin Tone",
	},
	"1F926-1F3FE-200D-2642-FE0F": {
		Key:        "1F926-1F3FE-200D-2642-FE0F",
		Value:      "🤦🏾‍♂️",
		Descriptor: "Man Facepalming: Medium-Dark Skin Tone",
	},
	"1F926-1F3FF": {
		Key:        "1F926-1F3FF",
		Value:      "🤦🏿",
		Descriptor: "Person Facepalming: Dark Skin Tone",
	},
	"1F926-1F3FF-200D-2640-FE0F": {
		Key:        "1F926-1F3FF-200D-2640-FE0F",
		Value:      "🤦🏿‍♀️",
		Descriptor: "Woman Facepalming: Dark Skin Tone",
	},
	"1F926-1F3FF-200D-2642-FE0F": {
		Key:        "1F926-1F3FF-200D-2642-FE0F",
		Value:      "🤦🏿‍♂️",
		Descriptor: "Man Facepalming: Dark Skin Tone",
	},
	"1F926-200D-2640-FE0F": {
		Key:        "1F926-200D-2640-FE0F",
		Value:      "🤦‍♀️",
		Descriptor: "Woman Facepalming",
	},
	"1F926-200D-2642-FE0F": {
		Key:        "1F926-200D-2642-FE0F",
		Value:      "🤦‍♂️",
		Descriptor: "Man Facepalming",
	},
	"1F927": {
		Key:        "1F927",
		Value:      "🤧",
		Descriptor: "Sneezing Face",
	},
	"1F928": {
		Key:        "1F928",
		Value:      "🤨",
		Descriptor: "Face With Raised Eyebrow",
	},
	"1F929": {
		Key:        "1F929",
		Value:      "🤩",
		Descriptor: "Star-Struck",
	},
	"1F92A": {
		Key:        "1F92A",
		Value:      "🤪",
		Descriptor: "Zany Face",
	},
	"1F92B": {
		Key:        "1F92B",
		Value:      "🤫",
		Descriptor: "Shushing Face",
	},
	"1F92C": {
		Key:        "1F92C",
		Value:      "🤬",
		Descriptor: "Face With Symbols on Mouth",
	},
	"1F92D": {
		Key:        "1F92D",
		Value:      "🤭",
		Descriptor: "Face With Hand Over Mouth",
	},
	"1F92E": {
		Key:        "1F92E",
		Value:      "🤮",
		Descriptor: "Face Vomiting",
	},
	"1F92F": {
		Key:        "1F92F",
		Value:      "🤯",
		Descriptor: "Exploding Head",
	},
	"1F930": {
		Key:        "1F930",
		Value:      "🤰",
		Descriptor: "Pregnant Woman",
	},
	"1F930-1F3FB": {
		Key:        "1F930-1F3FB",
		Value:      "🤰🏻",
		Descriptor: "Pregnant Woman: Light Skin Tone",
	},
	"1F930-1F3FC": {
		Key:        "1F930-1F3FC",
		Value:      "🤰🏼",
		Descriptor: "Pregnant Woman: Medium-Light Skin Tone",
	},
	"1F930-1F3FD": {
		Key:        "1F930-1F3FD",
		Value:      "🤰🏽",
		Descriptor: "Pregnant Woman: Medium Skin Tone",
	},
	"1F930-1F3FE": {
		Key:        "1F930-1F3FE",
		Value:      "🤰🏾",
		Descriptor: "Pregnant Woman: Medium-Dark Skin Tone",
	},
	"1F930-1F3FF": {
		Key:        "1F930-1F3FF",
		Value:      "🤰🏿",
		Descriptor: "Pregnant Woman: Dark Skin Tone",
	},
	"1F931": {
		Key:        "1F931",
		Value:      "🤱",
		Descriptor: "Breast-Feeding",
	},
	"1F931-1F3FB": {
		Key:        "1F931-1F3FB",
		Value:      "🤱🏻",
		Descriptor: "Breast-Feeding: Light Skin Tone",
	},
	"1F931-1F3FC": {
		Key:        "1F931-1F3FC",
		Value:      "🤱🏼",
		Descriptor: "Breast-Feeding: Medium-Light Skin Tone",
	},
	"1F931-1F3FD": {
		Key:        "1F931-1F3FD",
		Value:      "🤱🏽",
		Descriptor: "Breast-Feeding: Medium Skin Tone",
	},
	"1F931-1F3FE": {
		Key:        "1F931-1F3FE",
		Value:      "🤱🏾",
		Descriptor: "Breast-Feeding: Medium-Dark Skin Tone",
	},
	"1F931-1F3FF": {
		Key:        "1F931-1F3FF",
		Value:      "🤱🏿",
		Descriptor: "Breast-Feeding: Dark Skin Tone",
	},
	"1F932": {
		Key:        "1F932",
		Value:      "🤲",
		Descriptor: "Palms Up Together",
	},
	"1F932-1F3FB": {
		Key:        "1F932-1F3FB",
		Value:      "🤲🏻",
		Descriptor: "Palms Up Together: Light Skin Tone",
	},
	"1F932-1F3FC": {
		Key:        "1F932-1F3FC",
		Value:      "🤲🏼",
		Descriptor: "Palms Up Together: Medium-Light Skin Tone",
	},
	"1F932-1F3FD": {
		Key:        "1F932-1F3FD",
		Value:      "🤲🏽",
		Descriptor: "Palms Up Together: Medium Skin Tone",
	},
	"1F932-1F3FE": {
		Key:        "1F932-1F3FE",
		Value:      "🤲🏾",
		Descriptor: "Palms Up Together: Medium-Dark Skin Tone",
	},
	"1F932-1F3FF": {
		Key:        "1F932-1F3FF",
		Value:      "🤲🏿",
		Descriptor: "Palms Up Together: Dark Skin Tone",
	},
	"1F933": {
		Key:        "1F933",
		Value:      "🤳",
		Descriptor: "Selfie",
	},
	"1F933-1F3FB": {
		Key:        "1F933-1F3FB",
		Value:      "🤳🏻",
		Descriptor: "Selfie: Light Skin Tone",
	},
	"1F933-1F3FC": {
		Key:        "1F933-1F3FC",
		Value:      "🤳🏼",
		Descriptor: "Selfie: Medium-Light Skin Tone",
	},
	"1F933-1F3FD": {
		Key:        "1F933-1F3FD",
		Value:      "🤳🏽",
		Descriptor: "Selfie: Medium Skin Tone",
	},
	"1F933-1F3FE": {
		Key:        "1F933-1F3FE",
		Value:      "🤳🏾",
		Descriptor: "Selfie: Medium-Dark Skin Tone",
	},
	"1F933-1F3FF": {
		Key:        "1F933-1F3FF",
		Value:      "🤳🏿",
		Descriptor: "Selfie: Dark Skin Tone",
	},
	"1F934": {
		Key:        "1F934",
		Value:      "🤴",
		Descriptor: "Prince",
	},
	"1F934-1F3FB": {
		Key:        "1F934-1F3FB",
		Value:      "🤴🏻",
		Descriptor: "Prince: Light Skin Tone",
	},
	"1F934-1F3FC": {
		Key:        "1F934-1F3FC",
		Value:      "🤴🏼",
		Descriptor: "Prince: Medium-Light Skin Tone",
	},
	"1F934-1F3FD": {
		Key:        "1F934-1F3FD",
		Value:      "🤴🏽",
		Descriptor: "Prince: Medium Skin Tone",
	},
	"1F934-1F3FE": {
		Key:        "1F934-1F3FE",
		Value:      "🤴🏾",
		Descriptor: "Prince: Medium-Dark Skin Tone",
	},
	"1F934-1F3FF": {
		Key:        "1F934-1F3FF",
		Value:      "🤴🏿",
		Descriptor: "Prince: Dark Skin Tone",
	},
	"1F935": {
		Key:        "1F935",
		Value:      "🤵",
		Descriptor: "Man in Tuxedo",
	},
	"1F935-1F3FB": {
		Key:        "1F935-1F3FB",
		Value:      "🤵🏻",
		Descriptor: "Man in Tuxedo: Light Skin Tone",
	},
	"1F935-1F3FC": {
		Key:        "1F935-1F3FC",
		Value:      "🤵🏼",
		Descriptor: "Man in Tuxedo: Medium-Light Skin Tone",
	},
	"1F935-1F3FD": {
		Key:        "1F935-1F3FD",
		Value:      "🤵🏽",
		Descriptor: "Man in Tuxedo: Medium Skin Tone",
	},
	"1F935-1F3FE": {
		Key:        "1F935-1F3FE",
		Value:      "🤵🏾",
		Descriptor: "Man in Tuxedo: Medium-Dark Skin Tone",
	},
	"1F935-1F3FF": {
		Key:        "1F935-1F3FF",
		Value:      "🤵🏿",
		Descriptor: "Man in Tuxedo: Dark Skin Tone",
	},
	"1F936": {
		Key:        "1F936",
		Value:      "🤶",
		Descriptor: "Mrs. Claus",
	},
	"1F936-1F3FB": {
		Key:        "1F936-1F3FB",
		Value:      "🤶🏻",
		Descriptor: "Mrs. Claus: Light Skin Tone",
	},
	"1F936-1F3FC": {
		Key:        "1F936-1F3FC",
		Value:      "🤶🏼",
		Descriptor: "Mrs. Claus: Medium-Light Skin Tone",
	},
	"1F936-1F3FD": {
		Key:        "1F936-1F3FD",
		Value:      "🤶🏽",
		Descriptor: "Mrs. Claus: Medium Skin Tone",
	},
	"1F936-1F3FE": {
		Key:        "1F936-1F3FE",
		Value:      "🤶🏾",
		Descriptor: "Mrs. Claus: Medium-Dark Skin Tone",
	},
	"1F936-1F3FF": {
		Key:        "1F936-1F3FF",
		Value:      "🤶🏿",
		Descriptor: "Mrs. Claus: Dark Skin Tone",
	},
	"1F937": {
		Key:        "1F937",
		Value:      "🤷",
		Descriptor: "Person Shrugging",
	},
	"1F937-1F3FB": {
		Key:        "1F937-1F3FB",
		Value:      "🤷🏻",
		Descriptor: "Person Shrugging: Light Skin Tone",
	},
	"1F937-1F3FB-200D-2640-FE0F": {
		Key:        "1F937-1F3FB-200D-2640-FE0F",
		Value:      "🤷🏻‍♀️",
		Descriptor: "Woman Shrugging: Light Skin Tone",
	},
	"1F937-1F3FB-200D-2642-FE0F": {
		Key:        "1F937-1F3FB-200D-2642-FE0F",
		Value:      "🤷🏻‍♂️",
		Descriptor: "Man Shrugging: Light Skin Tone",
	},
	"1F937-1F3FC": {
		Key:        "1F937-1F3FC",
		Value:      "🤷🏼",
		Descriptor: "Person Shrugging: Medium-Light Skin Tone",
	},
	"1F937-1F3FC-200D-2640-FE0F": {
		Key:        "1F937-1F3FC-200D-2640-FE0F",
		Value:      "🤷🏼‍♀️",
		Descriptor: "Woman Shrugging: Medium-Light Skin Tone",
	},
	"1F937-1F3FC-200D-2642-FE0F": {
		Key:        "1F937-1F3FC-200D-2642-FE0F",
		Value:      "🤷🏼‍♂️",
		Descriptor: "Man Shrugging: Medium-Light Skin Tone",
	},
	"1F937-1F3FD": {
		Key:        "1F937-1F3FD",
		Value:      "🤷🏽",
		Descriptor: "Person Shrugging: Medium Skin Tone",
	},
	"1F937-1F3FD-200D-2640-FE0F": {
		Key:        "1F937-1F3FD-200D-2640-FE0F",
		Value:      "🤷🏽‍♀️",
		Descriptor: "Woman Shrugging: Medium Skin Tone",
	},
	"1F937-1F3FD-200D-2642-FE0F": {
		Key:        "1F937-1F3FD-200D-2642-FE0F",
		Value:      "🤷🏽‍♂️",
		Descriptor: "Man Shrugging: Medium Skin Tone",
	},
	"1F937-1F3FE": {
		Key:        "1F937-1F3FE",
		Value:      "🤷🏾",
		Descriptor: "Person Shrugging: Medium-Dark Skin Tone",
	},
	"1F937-1F3FE-200D-2640-FE0F": {
		Key:        "1F937-1F3FE-200D-2640-FE0F",
		Value:      "🤷🏾‍♀️",
		Descriptor: "Woman Shrugging: Medium-Dark Skin Tone",
	},
	"1F937-1F3FE-200D-2642-FE0F": {
		Key:        "1F937-1F3FE-200D-2642-FE0F",
		Value:      "🤷🏾‍♂️",
		Descriptor: "Man Shrugging: Medium-Dark Skin Tone",
	},
	"1F937-1F3FF": {
		Key:        "1F937-1F3FF",
		Value:      "🤷🏿",
		Descriptor: "Person Shrugging: Dark Skin Tone",
	},
	"1F937-1F3FF-200D-2640-FE0F": {
		Key:        "1F937-1F3FF-200D-2640-FE0F",
		Value:      "🤷🏿‍♀️",
		Descriptor: "Woman Shrugging: Dark Skin Tone",
	},
	"1F937-1F3FF-200D-2642-FE0F": {
		Key:        "1F937-1F3FF-200D-2642-FE0F",
		Value:      "🤷🏿‍♂️",
		Descriptor: "Man Shrugging: Dark Skin Tone",
	},
	"1F937-200D-2640-FE0F": {
		Key:        "1F937-200D-2640-FE0F",
		Value:      "🤷‍♀️",
		Descriptor: "Woman Shrugging",
	},
	"1F937-200D-2642-FE0F": {
		Key:        "1F937-200D-2642-FE0F",
		Value:      "🤷‍♂️",
		Descriptor: "Man Shrugging",
	},
	"1F938": {
		Key:        "1F938",
		Value:      "🤸",
		Descriptor: "Person Cartwheeling",
	},
	"1F938-1F3FB": {
		Key:        "1F938-1F3FB",
		Value:      "🤸🏻",
		Descriptor: "Person Cartwheeling: Light Skin Tone",
	},
	"1F938-1F3FB-200D-2640-FE0F": {
		Key:        "1F938-1F3FB-200D-2640-FE0F",
		Value:      "🤸🏻‍♀️",
		Descriptor: "Woman Cartwheeling: Light Skin Tone",
	},
	"1F938-1F3FB-200D-2642-FE0F": {
		Key:        "1F938-1F3FB-200D-2642-FE0F",
		Value:      "🤸🏻‍♂️",
		Descriptor: "Man Cartwheeling: Light Skin Tone",
	},
	"1F938-1F3FC": {
		Key:        "1F938-1F3FC",
		Value:      "🤸🏼",
		Descriptor: "Person Cartwheeling: Medium-Light Skin Tone",
	},
	"1F938-1F3FC-200D-2640-FE0F": {
		Key:        "1F938-1F3FC-200D-2640-FE0F",
		Value:      "🤸🏼‍♀️",
		Descriptor: "Woman Cartwheeling: Medium-Light Skin Tone",
	},
	"1F938-1F3FC-200D-2642-FE0F": {
		Key:        "1F938-1F3FC-200D-2642-FE0F",
		Value:      "🤸🏼‍♂️",
		Descriptor: "Man Cartwheeling: Medium-Light Skin Tone",
	},
	"1F938-1F3FD": {
		Key:        "1F938-1F3FD",
		Value:      "🤸🏽",
		Descriptor: "Person Cartwheeling: Medium Skin Tone",
	},
	"1F938-1F3FD-200D-2640-FE0F": {
		Key:        "1F938-1F3FD-200D-2640-FE0F",
		Value:      "🤸🏽‍♀️",
		Descriptor: "Woman Cartwheeling: Medium Skin Tone",
	},
	"1F938-1F3FD-200D-2642-FE0F": {
		Key:        "1F938-1F3FD-200D-2642-FE0F",
		Value:      "🤸🏽‍♂️",
		Descriptor: "Man Cartwheeling: Medium Skin Tone",
	},
	"1F938-1F3FE": {
		Key:        "1F938-1F3FE",
		Value:      "🤸🏾",
		Descriptor: "Person Cartwheeling: Medium-Dark Skin Tone",
	},
	"1F938-1F3FE-200D-2640-FE0F": {
		Key:        "1F938-1F3FE-200D-2640-FE0F",
		Value:      "🤸🏾‍♀️",
		Descriptor: "Woman Cartwheeling: Medium-Dark Skin Tone",
	},
	"1F938-1F3FE-200D-2642-FE0F": {
		Key:        "1F938-1F3FE-200D-2642-FE0F",
		Value:      "🤸🏾‍♂️",
		Descriptor: "Man Cartwheeling: Medium-Dark Skin Tone",
	},
	"1F938-1F3FF": {
		Key:        "1F938-1F3FF",
		Value:      "🤸🏿",
		Descriptor: "Person Cartwheeling: Dark Skin Tone",
	},
	"1F938-1F3FF-200D-2640-FE0F": {
		Key:        "1F938-1F3FF-200D-2640-FE0F",
		Value:      "🤸🏿‍♀️",
		Descriptor: "Woman Cartwheeling: Dark Skin Tone",
	},
	"1F938-1F3FF-200D-2642-FE0F": {
		Key:        "1F938-1F3FF-200D-2642-FE0F",
		Value:      "🤸🏿‍♂️",
		Descriptor: "Man Cartwheeling: Dark Skin Tone",
	},
	"1F938-200D-2640-FE0F": {
		Key:        "1F938-200D-2640-FE0F",
		Value:      "🤸‍♀️",
		Descriptor: "Woman Cartwheeling",
	},
	"1F938-200D-2642-FE0F": {
		Key:        "1F938-200D-2642-FE0F",
		Value:      "🤸‍♂️",
		Descriptor: "Man Cartwheeling",
	},
	"1F939": {
		Key:        "1F939",
		Value:      "🤹",
		Descriptor: "Person Juggling",
	},
	"1F939-1F3FB": {
		Key:        "1F939-1F3FB",
		Value:      "🤹🏻",
		Descriptor: "Person Juggling: Light Skin Tone",
	},
	"1F939-1F3FB-200D-2640-FE0F": {
		Key:        "1F939-1F3FB-200D-2640-FE0F",
		Value:      "🤹🏻‍♀️",
		Descriptor: "Woman Juggling: Light Skin Tone",
	},
	"1F939-1F3FB-200D-2642-FE0F": {
		Key:        "1F939-1F3FB-200D-2642-FE0F",
		Value:      "🤹🏻‍♂️",
		Descriptor: "Man Juggling: Light Skin Tone",
	},
	"1F939-1F3FC": {
		Key:        "1F939-1F3FC",
		Value:      "🤹🏼",
		Descriptor: "Person Juggling: Medium-Light Skin Tone",
	},
	"1F939-1F3FC-200D-2640-FE0F": {
		Key:        "1F939-1F3FC-200D-2640-FE0F",
		Value:      "🤹🏼‍♀️",
		Descriptor: "Woman Juggling: Medium-Light Skin Tone",
	},
	"1F939-1F3FC-200D-2642-FE0F": {
		Key:        "1F939-1F3FC-200D-2642-FE0F",
		Value:      "🤹🏼‍♂️",
		Descriptor: "Man Juggling: Medium-Light Skin Tone",
	},
	"1F939-1F3FD": {
		Key:        "1F939-1F3FD",
		Value:      "🤹🏽",
		Descriptor: "Person Juggling: Medium Skin Tone",
	},
	"1F939-1F3FD-200D-2640-FE0F": {
		Key:        "1F939-1F3FD-200D-2640-FE0F",
		Value:      "🤹🏽‍♀️",
		Descriptor: "Woman Juggling: Medium Skin Tone",
	},
	"1F939-1F3FD-200D-2642-FE0F": {
		Key:        "1F939-1F3FD-200D-2642-FE0F",
		Value:      "🤹🏽‍♂️",
		Descriptor: "Man Juggling: Medium Skin Tone",
	},
	"1F939-1F3FE": {
		Key:        "1F939-1F3FE",
		Value:      "🤹🏾",
		Descriptor: "Person Juggling: Medium-Dark Skin Tone",
	},
	"1F939-1F3FE-200D-2640-FE0F": {
		Key:        "1F939-1F3FE-200D-2640-FE0F",
		Value:      "🤹🏾‍♀️",
		Descriptor: "Woman Juggling: Medium-Dark Skin Tone",
	},
	"1F939-1F3FE-200D-2642-FE0F": {
		Key:        "1F939-1F3FE-200D-2642-FE0F",
		Value:      "🤹🏾‍♂️",
		Descriptor: "Man Juggling: Medium-Dark Skin Tone",
	},
	"1F939-1F3FF": {
		Key:        "1F939-1F3FF",
		Value:      "🤹🏿",
		Descriptor: "Person Juggling: Dark Skin Tone",
	},
	"1F939-1F3FF-200D-2640-FE0F": {
		Key:        "1F939-1F3FF-200D-2640-FE0F",
		Value:      "🤹🏿‍♀️",
		Descriptor: "Woman Juggling: Dark Skin Tone",
	},
	"1F939-1F3FF-200D-2642-FE0F": {
		Key:        "1F939-1F3FF-200D-2642-FE0F",
		Value:      "🤹🏿‍♂️",
		Descriptor: "Man Juggling: Dark Skin Tone",
	},
	"1F939-200D-2640-FE0F": {
		Key:        "1F939-200D-2640-FE0F",
		Value:      "🤹‍♀️",
		Descriptor: "Woman Juggling",
	},
	"1F939-200D-2642-FE0F": {
		Key:        "1F939-200D-2642-FE0F",
		Value:      "🤹‍♂️",
		Descriptor: "Man Juggling",
	},
	"1F93A": {
		Key:        "1F93A",
		Value:      "🤺",
		Descriptor: "Person Fencing",
	},
	"1F93C": {
		Key:        "1F93C",
		Value:      "🤼",
		Descriptor: "People Wrestling",
	},
	"1F93C-200D-2640-FE0F": {
		Key:        "1F93C-200D-2640-FE0F",
		Value:      "🤼‍♀️",
		Descriptor: "Women Wrestling",
	},
	"1F93C-200D-2642-FE0F": {
		Key:        "1F93C-200D-2642-FE0F",
		Value:      "🤼‍♂️",
		Descriptor: "Men Wrestling",
	},
	"1F93D": {
		Key:        "1F93D",
		Value:      "🤽",
		Descriptor: "Person Playing Water Polo",
	},
	"1F93D-1F3FB": {
		Key:        "1F93D-1F3FB",
		Value:      "🤽🏻",
		Descriptor: "Person Playing Water Polo: Light Skin Tone",
	},
	"1F93D-1F3FB-200D-2640-FE0F": {
		Key:        "1F93D-1F3FB-200D-2640-FE0F",
		Value:      "🤽🏻‍♀️",
		Descriptor: "Woman Playing Water Polo: Light Skin Tone",
	},
	"1F93D-1F3FB-200D-2642-FE0F": {
		Key:        "1F93D-1F3FB-200D-2642-FE0F",
		Value:      "🤽🏻‍♂️",
		Descriptor: "Man Playing Water Polo: Light Skin Tone",
	},
	"1F93D-1F3FC": {
		Key:        "1F93D-1F3FC",
		Value:      "🤽🏼",
		Descriptor: "Person Playing Water Polo: Medium-Light Skin Tone",
	},
	"1F93D-1F3FC-200D-2640-FE0F": {
		Key:        "1F93D-1F3FC-200D-2640-FE0F",
		Value:      "🤽🏼‍♀️",
		Descriptor: "Woman Playing Water Polo: Medium-Light Skin Tone",
	},
	"1F93D-1F3FC-200D-2642-FE0F": {
		Key:        "1F93D-1F3FC-200D-2642-FE0F",
		Value:      "🤽🏼‍♂️",
		Descriptor: "Man Playing Water Polo: Medium-Light Skin Tone",
	},
	"1F93D-1F3FD": {
		Key:        "1F93D-1F3FD",
		Value:      "🤽🏽",
		Descriptor: "Person Playing Water Polo: Medium Skin Tone",
	},
	"1F93D-1F3FD-200D-2640-FE0F": {
		Key:        "1F93D-1F3FD-200D-2640-FE0F",
		Value:      "🤽🏽‍♀️",
		Descriptor: "Woman Playing Water Polo: Medium Skin Tone",
	},
	"1F93D-1F3FD-200D-2642-FE0F": {
		Key:        "1F93D-1F3FD-200D-2642-FE0F",
		Value:      "🤽🏽‍♂️",
		Descriptor: "Man Playing Water Polo: Medium Skin Tone",
	},
	"1F93D-1F3FE": {
		Key:        "1F93D-1F3FE",
		Value:      "🤽🏾",
		Descriptor: "Person Playing Water Polo: Medium-Dark Skin Tone",
	},
	"1F93D-1F3FE-200D-2640-FE0F": {
		Key:        "1F93D-1F3FE-200D-2640-FE0F",
		Value:      "🤽🏾‍♀️",
		Descriptor: "Woman Playing Water Polo: Medium-Dark Skin Tone",
	},
	"1F93D-1F3FE-200D-2642-FE0F": {
		Key:        "1F93D-1F3FE-200D-2642-FE0F",
		Value:      "🤽🏾‍♂️",
		Descriptor: "Man Playing Water Polo: Medium-Dark Skin Tone",
	},
	"1F93D-1F3FF": {
		Key:        "1F93D-1F3FF",
		Value:      "🤽🏿",
		Descriptor: "Person Playing Water Polo: Dark Skin Tone",
	},
	"1F93D-1F3FF-200D-2640-FE0F": {
		Key:        "1F93D-1F3FF-200D-2640-FE0F",
		Value:      "🤽🏿‍♀️",
		Descriptor: "Woman Playing Water Polo: Dark Skin Tone",
	},
	"1F93D-1F3FF-200D-2642-FE0F": {
		Key:        "1F93D-1F3FF-200D-2642-FE0F",
		Value:      "🤽🏿‍♂️",
		Descriptor: "Man Playing Water Polo: Dark Skin Tone",
	},
	"1F93D-200D-2640-FE0F": {
		Key:        "1F93D-200D-2640-FE0F",
		Value:      "🤽‍♀️",
		Descriptor: "Woman Playing Water Polo",
	},
	"1F93D-200D-2642-FE0F": {
		Key:        "1F93D-200D-2642-FE0F",
		Value:      "🤽‍♂️",
		Descriptor: "Man Playing Water Polo",
	},
	"1F93E": {
		Key:        "1F93E",
		Value:      "🤾",
		Descriptor: "Person Playing Handball",
	},
	"1F93E-1F3FB": {
		Key:        "1F93E-1F3FB",
		Value:      "🤾🏻",
		Descriptor: "Person Playing Handball: Light Skin Tone",
	},
	"1F93E-1F3FB-200D-2640-FE0F": {
		Key:        "1F93E-1F3FB-200D-2640-FE0F",
		Value:      "🤾🏻‍♀️",
		Descriptor: "Woman Playing Handball: Light Skin Tone",
	},
	"1F93E-1F3FB-200D-2642-FE0F": {
		Key:        "1F93E-1F3FB-200D-2642-FE0F",
		Value:      "🤾🏻‍♂️",
		Descriptor: "Man Playing Handball: Light Skin Tone",
	},
	"1F93E-1F3FC": {
		Key:        "1F93E-1F3FC",
		Value:      "🤾🏼",
		Descriptor: "Person Playing Handball: Medium-Light Skin Tone",
	},
	"1F93E-1F3FC-200D-2640-FE0F": {
		Key:        "1F93E-1F3FC-200D-2640-FE0F",
		Value:      "🤾🏼‍♀️",
		Descriptor: "Woman Playing Handball: Medium-Light Skin Tone",
	},
	"1F93E-1F3FC-200D-2642-FE0F": {
		Key:        "1F93E-1F3FC-200D-2642-FE0F",
		Value:      "🤾🏼‍♂️",
		Descriptor: "Man Playing Handball: Medium-Light Skin Tone",
	},
	"1F93E-1F3FD": {
		Key:        "1F93E-1F3FD",
		Value:      "🤾🏽",
		Descriptor: "Person Playing Handball: Medium Skin Tone",
	},
	"1F93E-1F3FD-200D-2640-FE0F": {
		Key:        "1F93E-1F3FD-200D-2640-FE0F",
		Value:      "🤾🏽‍♀️",
		Descriptor: "Woman Playing Handball: Medium Skin Tone",
	},
	"1F93E-1F3FD-200D-2642-FE0F": {
		Key:        "1F93E-1F3FD-200D-2642-FE0F",
		Value:      "🤾🏽‍♂️",
		Descriptor: "Man Playing Handball: Medium Skin Tone",
	},
	"1F93E-1F3FE": {
		Key:        "1F93E-1F3FE",
		Value:      "🤾🏾",
		Descriptor: "Person Playing Handball: Medium-Dark Skin Tone",
	},
	"1F93E-1F3FE-200D-2640-FE0F": {
		Key:        "1F93E-1F3FE-200D-2640-FE0F",
		Value:      "🤾🏾‍♀️",
		Descriptor: "Woman Playing Handball: Medium-Dark Skin Tone",
	},
	"1F93E-1F3FE-200D-2642-FE0F": {
		Key:        "1F93E-1F3FE-200D-2642-FE0F",
		Value:      "🤾🏾‍♂️",
		Descriptor: "Man Playing Handball: Medium-Dark Skin Tone",
	},
	"1F93E-1F3FF": {
		Key:        "1F93E-1F3FF",
		Value:      "🤾🏿",
		Descriptor: "Person Playing Handball: Dark Skin Tone",
	},
	"1F93E-1F3FF-200D-2640-FE0F": {
		Key:        "1F93E-1F3FF-200D-2640-FE0F",
		Value:      "🤾🏿‍♀️",
		Descriptor: "Woman Playing Handball: Dark Skin Tone",
	},
	"1F93E-1F3FF-200D-2642-FE0F": {
		Key:        "1F93E-1F3FF-200D-2642-FE0F",
		Value:      "🤾🏿‍♂️",
		Descriptor: "Man Playing Handball: Dark Skin Tone",
	},
	"1F93E-200D-2640-FE0F": {
		Key:        "1F93E-200D-2640-FE0F",
		Value:      "🤾‍♀️",
		Descriptor: "Woman Playing Handball",
	},
	"1F93E-200D-2642-FE0F": {
		Key:        "1F93E-200D-2642-FE0F",
		Value:      "🤾‍♂️",
		Descriptor: "Man Playing Handball",
	},
	"1F93F": {
		Key:        "1F93F",
		Value:      "🤿",
		Descriptor: "Diving Mask",
	},
	"1F940": {
		Key:        "1F940",
		Value:      "🥀",
		Descriptor: "Wilted Flower",
	},
	"1F941": {
		Key:        "1F941",
		Value:      "🥁",
		Descriptor: "Drum",
	},
	"1F942": {
		Key:        "1F942",
		Value:      "🥂",
		Descriptor: "Clinking Glasses",
	},
	"1F943": {
		Key:        "1F943",
		Value:      "🥃",
		Descriptor: "Tumbler Glass",
	},
	"1F944": {
		Key:        "1F944",
		Value:      "🥄",
		Descriptor: "Spoon",
	},
	"1F945": {
		Key:        "1F945",
		Value:      "🥅",
		Descriptor: "Goal Net",
	},
	"1F947": {
		Key:        "1F947",
		Value:      "🥇",
		Descriptor: "1st Place Medal",
	},
	"1F948": {
		Key:        "1F948",
		Value:      "🥈",
		Descriptor: "2nd Place Medal",
	},
	"1F949": {
		Key:        "1F949",
		Value:      "🥉",
		Descriptor: "3rd Place Medal",
	},
	"1F94A": {
		Key:        "1F94A",
		Value:      "🥊",
		Descriptor: "Boxing Glove",
	},
	"1F94B": {
		Key:        "1F94B",
		Value:      "🥋",
		Descriptor: "Martial Arts Uniform",
	},
	"1F94C": {
		Key:        "1F94C",
		Value:      "🥌",
		Descriptor: "Curling Stone",
	},
	"1F94D": {
		Key:        "1F94D",
		Value:      "🥍",
		Descriptor: "Lacrosse",
	},
	"1F94E": {
		Key:        "1F94E",
		Value:      "🥎",
		Descriptor: "Softball",
	},
	"1F94F": {
		Key:        "1F94F",
		Value:      "🥏",
		Descriptor: "Flying Disc",
	},
	"1F950": {
		Key:        "1F950",
		Value:      "🥐",
		Descriptor: "Croissant",
	},
	"1F951": {
		Key:        "1F951",
		Value:      "🥑",
		Descriptor: "Avocado",
	},
	"1F952": {
		Key:        "1F952",
		Value:      "🥒",
		Descriptor: "Cucumber",
	},
	"1F953": {
		Key:        "1F953",
		Value:      "🥓",
		Descriptor: "Bacon",
	},
	"1F954": {
		Key:        "1F954",
		Value:      "🥔",
		Descriptor: "Potato",
	},
	"1F955": {
		Key:        "1F955",
		Value:      "🥕",
		Descriptor: "Carrot",
	},
	"1F956": {
		Key:        "1F956",
		Value:      "🥖",
		Descriptor: "Baguette Bread",
	},
	"1F957": {
		Key:        "1F957",
		Value:      "🥗",
		Descriptor: "Green Salad",
	},
	"1F958": {
		Key:        "1F958",
		Value:      "🥘",
		Descriptor: "Shallow Pan of Food",
	},
	"1F959": {
		Key:        "1F959",
		Value:      "🥙",
		Descriptor: "Stuffed Flatbread",
	},
	"1F95A": {
		Key:        "1F95A",
		Value:      "🥚",
		Descriptor: "Egg",
	},
	"1F95B": {
		Key:        "1F95B",
		Value:      "🥛",
		Descriptor: "Glass of Milk",
	},
	"1F95C": {
		Key:        "1F95C",
		Value:      "🥜",
		Descriptor: "Peanuts",
	},
	"1F95D": {
		Key:        "1F95D",
		Value:      "🥝",
		Descriptor: "Kiwi Fruit",
	},
	"1F95E": {
		Key:        "1F95E",
		Value:      "🥞",
		Descriptor: "Pancakes",
	},
	"1F95F": {
		Key:        "1F95F",
		Value:      "🥟",
		Descriptor: "Dumpling",
	},
	"1F960": {
		Key:        "1F960",
		Value:      "🥠",
		Descriptor: "Fortune Cookie",
	},
	"1F961": {
		Key:        "1F961",
		Value:      "🥡",
		Descriptor: "Takeout Box",
	},
	"1F962": {
		Key:        "1F962",
		Value:      "🥢",
		Descriptor: "Chopsticks",
	},
	"1F963": {
		Key:        "1F963",
		Value:      "🥣",
		Descriptor: "Bowl With Spoon",
	},
	"1F964": {
		Key:        "1F964",
		Value:      "🥤",
		Descriptor: "Cup With Straw",
	},
	"1F965": {
		Key:        "1F965",
		Value:      "🥥",
		Descriptor: "Coconut",
	},
	"1F966": {
		Key:        "1F966",
		Value:      "🥦",
		Descriptor: "Broccoli",
	},
	"1F967": {
		Key:        "1F967",
		Value:      "🥧",
		Descriptor: "Pie",
	},
	"1F968": {
		Key:        "1F968",
		Value:      "🥨",
		Descriptor: "Pretzel",
	},
	"1F969": {
		Key:        "1F969",
		Value:      "🥩",
		Descriptor: "Cut of Meat",
	},
	"1F96A": {
		Key:        "1F96A",
		Value:      "🥪",
		Descriptor: "Sandwich",
	},
	"1F96B": {
		Key:        "1F96B",
		Value:      "🥫",
		Descriptor: "Canned Food",
	},
	"1F96C": {
		Key:        "1F96C",
		Value:      "🥬",
		Descriptor: "Leafy Green",
	},
	"1F96D": {
		Key:        "1F96D",
		Value:      "🥭",
		Descriptor: "Mango",
	},
	"1F96E": {
		Key:        "1F96E",
		Value:      "🥮",
		Descriptor: "Moon Cake",
	},
	"1F96F": {
		Key:        "1F96F",
		Value:      "🥯",
		Descriptor: "Bagel",
	},
	"1F970": {
		Key:        "1F970",
		Value:      "🥰",
		Descriptor: "Smiling Face With Hearts",
	},
	"1F971": {
		Key:        "1F971",
		Value:      "🥱",
		Descriptor: "Yawning Face",
	},
	"1F973": {
		Key:        "1F973",
		Value:      "🥳",
		Descriptor: "Partying Face",
	},
	"1F974": {
		Key:        "1F974",
		Value:      "🥴",
		Descriptor: "Woozy Face",
	},
	"1F975": {
		Key:        "1F975",
		Value:      "🥵",
		Descriptor: "Hot Face",
	},
	"1F976": {
		Key:        "1F976",
		Value:      "🥶",
		Descriptor: "Cold Face",
	},
	"1F97A": {
		Key:        "1F97A",
		Value:      "🥺",
		Descriptor: "Pleading Face",
	},
	"1F97B": {
		Key:        "1F97B",
		Value:      "🥻",
		Descriptor: "Sari",
	},
	"1F97C": {
		Key:        "1F97C",
		Value:      "🥼",
		Descriptor: "Lab Coat",
	},
	"1F97D": {
		Key:        "1F97D",
		Value:      "🥽",
		Descriptor: "Goggles",
	},
	"1F97E": {
		Key:        "1F97E",
		Value:      "🥾",
		Descriptor: "Hiking Boot",
	},
	"1F97F": {
		Key:        "1F97F",
		Value:      "🥿",
		Descriptor: "Flat Shoe",
	},
	"1F980": {
		Key:        "1F980",
		Value:      "🦀",
		Descriptor: "Crab",
	},
	"1F981": {
		Key:        "1F981",
		Value:      "🦁",
		Descriptor: "Lion",
	},
	"1F982": {
		Key:        "1F982",
		Value:      "🦂",
		Descriptor: "Scorpion",
	},
	"1F983": {
		Key:        "1F983",
		Value:      "🦃",
		Descriptor: "Turkey",
	},
	"1F984": {
		Key:        "1F984",
		Value:      "🦄",
		Descriptor: "Unicorn",
	},
	"1F985": {
		Key:        "1F985",
		Value:      "🦅",
		Descriptor: "Eagle",
	},
	"1F986": {
		Key:        "1F986",
		Value:      "🦆",
		Descriptor: "Duck",
	},
	"1F987": {
		Key:        "1F987",
		Value:      "🦇",
		Descriptor: "Bat",
	},
	"1F988": {
		Key:        "1F988",
		Value:      "🦈",
		Descriptor: "Shark",
	},
	"1F989": {
		Key:        "1F989",
		Value:      "🦉",
		Descriptor: "Owl",
	},
	"1F98A": {
		Key:        "1F98A",
		Value:      "🦊",
		Descriptor: "Fox",
	},
	"1F98B": {
		Key:        "1F98B",
		Value:      "🦋",
		Descriptor: "Butterfly",
	},
	"1F98C": {
		Key:        "1F98C",
		Value:      "🦌",
		Descriptor: "Deer",
	},
	"1F98D": {
		Key:        "1F98D",
		Value:      "🦍",
		Descriptor: "Gorilla",
	},
	"1F98E": {
		Key:        "1F98E",
		Value:      "🦎",
		Descriptor: "Lizard",
	},
	"1F98F": {
		Key:        "1F98F",
		Value:      "🦏",
		Descriptor: "Rhinoceros",
	},
	"1F990": {
		Key:        "1F990",
		Value:      "🦐",
		Descriptor: "Shrimp",
	},
	"1F991": {
		Key:        "1F991",
		Value:      "🦑",
		Descriptor: "Squid",
	},
	"1F992": {
		Key:        "1F992",
		Value:      "🦒",
		Descriptor: "Giraffe",
	},
	"1F993": {
		Key:        "1F993",
		Value:      "🦓",
		Descriptor: "Zebra",
	},
	"1F994": {
		Key:        "1F994",
		Value:      "🦔",
		Descriptor: "Hedgehog",
	},
	"1F995": {
		Key:        "1F995",
		Value:      "🦕",
		Descriptor: "Sauropod",
	},
	"1F996": {
		Key:        "1F996",
		Value:      "🦖",
		Descriptor: "T-Rex",
	},
	"1F997": {
		Key:        "1F997",
		Value:      "🦗",
		Descriptor: "Cricket",
	},
	"1F998": {
		Key:        "1F998",
		Value:      "🦘",
		Descriptor: "Kangaroo",
	},
	"1F999": {
		Key:        "1F999",
		Value:      "🦙",
		Descriptor: "Llama",
	},
	"1F99A": {
		Key:        "1F99A",
		Value:      "🦚",
		Descriptor: "Peacock",
	},
	"1F99B": {
		Key:        "1F99B",
		Value:      "🦛",
		Descriptor: "Hippopotamus",
	},
	"1F99C": {
		Key:        "1F99C",
		Value:      "🦜",
		Descriptor: "Parrot",
	},
	"1F99D": {
		Key:        "1F99D",
		Value:      "🦝",
		Descriptor: "Raccoon",
	},
	"1F99E": {
		Key:        "1F99E",
		Value:      "🦞",
		Descriptor: "Lobster",
	},
	"1F99F": {
		Key:        "1F99F",
		Value:      "🦟",
		Descriptor: "Mosquito",
	},
	"1F9A0": {
		Key:        "1F9A0",
		Value:      "🦠",
		Descriptor: "Microbe",
	},
	"1F9A1": {
		Key:        "1F9A1",
		Value:      "🦡",
		Descriptor: "Badger",
	},
	"1F9A2": {
		Key:        "1F9A2",
		Value:      "🦢",
		Descriptor: "Swan",
	},
	"1F9A5": {
		Key:        "1F9A5",
		Value:      "🦥",
		Descriptor: "Sloth",
	},
	"1F9A6": {
		Key:        "1F9A6",
		Value:      "🦦",
		Descriptor: "Otter",
	},
	"1F9A7": {
		Key:        "1F9A7",
		Value:      "🦧",
		Descriptor: "Orangutan",
	},
	"1F9A8": {
		Key:        "1F9A8",
		Value:      "🦨",
		Descriptor: "Skunk",
	},
	"1F9A9": {
		Key:        "1F9A9",
		Value:      "🦩",
		Descriptor: "Flamingo",
	},
	"1F9AA": {
		Key:        "1F9AA",
		Value:      "🦪",
		Descriptor: "Oyster",
	},
	"1F9AE": {
		Key:        "1F9AE",
		Value:      "🦮",
		Descriptor: "Guide Dog",
	},
	"1F9AF": {
		Key:        "1F9AF",
		Value:      "🦯",
		Descriptor: "Probing Cane",
	},
	"1F9B0": {
		Key:        "1F9B0",
		Value:      "🦰",
		Descriptor: "Red Hair",
	},
	"1F9B1": {
		Key:        "1F9B1",
		Value:      "🦱",
		Descriptor: "Curly Hair",
	},
	"1F9B2": {
		Key:        "1F9B2",
		Value:      "🦲",
		Descriptor: "Bald",
	},
	"1F9B3": {
		Key:        "1F9B3",
		Value:      "🦳",
		Descriptor: "White Hair",
	},
	"1F9B4": {
		Key:        "1F9B4",
		Value:      "🦴",
		Descriptor: "Bone",
	},
	"1F9B5": {
		Key:        "1F9B5",
		Value:      "🦵",
		Descriptor: "Leg",
	},
	"1F9B5-1F3FB": {
		Key:        "1F9B5-1F3FB",
		Value:      "🦵🏻",
		Descriptor: "Leg: Light Skin Tone",
	},
	"1F9B5-1F3FC": {
		Key:        "1F9B5-1F3FC",
		Value:      "🦵🏼",
		Descriptor: "Leg: Medium-Light Skin Tone",
	},
	"1F9B5-1F3FD": {
		Key:        "1F9B5-1F3FD",
		Value:      "🦵🏽",
		Descriptor: "Leg: Medium Skin Tone",
	},
	"1F9B5-1F3FE": {
		Key:        "1F9B5-1F3FE",
		Value:      "🦵🏾",
		Descriptor: "Leg: Medium-Dark Skin Tone",
	},
	"1F9B5-1F3FF": {
		Key:        "1F9B5-1F3FF",
		Value:      "🦵🏿",
		Descriptor: "Leg: Dark Skin Tone",
	},
	"1F9B6": {
		Key:        "1F9B6",
		Value:      "🦶",
		Descriptor: "Foot",
	},
	"1F9B6-1F3FB": {
		Key:        "1F9B6-1F3FB",
		Value:      "🦶🏻",
		Descriptor: "Foot: Light Skin Tone",
	},
	"1F9B6-1F3FC": {
		Key:        "1F9B6-1F3FC",
		Value:      "🦶🏼",
		Descriptor: "Foot: Medium-Light Skin Tone",
	},
	"1F9B6-1F3FD": {
		Key:        "1F9B6-1F3FD",
		Value:      "🦶🏽",
		Descriptor: "Foot: Medium Skin Tone",
	},
	"1F9B6-1F3FE": {
		Key:        "1F9B6-1F3FE",
		Value:      "🦶🏾",
		Descriptor: "Foot: Medium-Dark Skin Tone",
	},
	"1F9B6-1F3FF": {
		Key:        "1F9B6-1F3FF",
		Value:      "🦶🏿",
		Descriptor: "Foot: Dark Skin Tone",
	},
	"1F9B7": {
		Key:        "1F9B7",
		Value:      "🦷",
		Descriptor: "Tooth",
	},
	"1F9B8": {
		Key:        "1F9B8",
		Value:      "🦸",
		Descriptor: "Superhero",
	},
	"1F9B8-1F3FB": {
		Key:        "1F9B8-1F3FB",
		Value:      "🦸🏻",
		Descriptor: "Superhero: Light Skin Tone",
	},
	"1F9B8-1F3FB-200D-2640-FE0F": {
		Key:        "1F9B8-1F3FB-200D-2640-FE0F",
		Value:      "🦸🏻‍♀️",
		Descriptor: "Woman Superhero: Light Skin Tone",
	},
	"1F9B8-1F3FB-200D-2642-FE0F": {
		Key:        "1F9B8-1F3FB-200D-2642-FE0F",
		Value:      "🦸🏻‍♂️",
		Descriptor: "Man Superhero: Light Skin Tone",
	},
	"1F9B8-1F3FC": {
		Key:        "1F9B8-1F3FC",
		Value:      "🦸🏼",
		Descriptor: "Superhero: Medium-Light Skin Tone",
	},
	"1F9B8-1F3FC-200D-2640-FE0F": {
		Key:        "1F9B8-1F3FC-200D-2640-FE0F",
		Value:      "🦸🏼‍♀️",
		Descriptor: "Woman Superhero: Medium-Light Skin Tone",
	},
	"1F9B8-1F3FC-200D-2642-FE0F": {
		Key:        "1F9B8-1F3FC-200D-2642-FE0F",
		Value:      "🦸🏼‍♂️",
		Descriptor: "Man Superhero: Medium-Light Skin Tone",
	},
	"1F9B8-1F3FD": {
		Key:        "1F9B8-1F3FD",
		Value:      "🦸🏽",
		Descriptor: "Superhero: Medium Skin Tone",
	},
	"1F9B8-1F3FD-200D-2640-FE0F": {
		Key:        "1F9B8-1F3FD-200D-2640-FE0F",
		Value:      "🦸🏽‍♀️",
		Descriptor: "Woman Superhero: Medium Skin Tone",
	},
	"1F9B8-1F3FD-200D-2642-FE0F": {
		Key:        "1F9B8-1F3FD-200D-2642-FE0F",
		Value:      "🦸🏽‍♂️",
		Descriptor: "Man Superhero: Medium Skin Tone",
	},
	"1F9B8-1F3FE": {
		Key:        "1F9B8-1F3FE",
		Value:      "🦸🏾",
		Descriptor: "Superhero: Medium-Dark Skin Tone",
	},
	"1F9B8-1F3FE-200D-2640-FE0F": {
		Key:        "1F9B8-1F3FE-200D-2640-FE0F",
		Value:      "🦸🏾‍♀️",
		Descriptor: "Woman Superhero: Medium-Dark Skin Tone",
	},
	"1F9B8-1F3FE-200D-2642-FE0F": {
		Key:        "1F9B8-1F3FE-200D-2642-FE0F",
		Value:      "🦸🏾‍♂️",
		Descriptor: "Man Superhero: Medium-Dark Skin Tone",
	},
	"1F9B8-1F3FF": {
		Key:        "1F9B8-1F3FF",
		Value:      "🦸🏿",
		Descriptor: "Superhero: Dark Skin Tone",
	},
	"1F9B8-1F3FF-200D-2640-FE0F": {
		Key:        "1F9B8-1F3FF-200D-2640-FE0F",
		Value:      "🦸🏿‍♀️",
		Descriptor: "Woman Superhero: Dark Skin Tone",
	},
	"1F9B8-1F3FF-200D-2642-FE0F": {
		Key:        "1F9B8-1F3FF-200D-2642-FE0F",
		Value:      "🦸🏿‍♂️",
		Descriptor: "Man Superhero: Dark Skin Tone",
	},
	"1F9B8-200D-2640-FE0F": {
		Key:        "1F9B8-200D-2640-FE0F",
		Value:      "🦸‍♀️",
		Descriptor: "Woman Superhero",
	},
	"1F9B8-200D-2642-FE0F": {
		Key:        "1F9B8-200D-2642-FE0F",
		Value:      "🦸‍♂️",
		Descriptor: "Man Superhero",
	},
	"1F9B9": {
		Key:        "1F9B9",
		Value:      "🦹",
		Descriptor: "Supervillain",
	},
	"1F9B9-1F3FB": {
		Key:        "1F9B9-1F3FB",
		Value:      "🦹🏻",
		Descriptor: "Supervillain: Light Skin Tone",
	},
	"1F9B9-1F3FB-200D-2640-FE0F": {
		Key:        "1F9B9-1F3FB-200D-2640-FE0F",
		Value:      "🦹🏻‍♀️",
		Descriptor: "Woman Supervillain: Light Skin Tone",
	},
	"1F9B9-1F3FB-200D-2642-FE0F": {
		Key:        "1F9B9-1F3FB-200D-2642-FE0F",
		Value:      "🦹🏻‍♂️",
		Descriptor: "Man Supervillain: Light Skin Tone",
	},
	"1F9B9-1F3FC": {
		Key:        "1F9B9-1F3FC",
		Value:      "🦹🏼",
		Descriptor: "Supervillain: Medium-Light Skin Tone",
	},
	"1F9B9-1F3FC-200D-2640-FE0F": {
		Key:        "1F9B9-1F3FC-200D-2640-FE0F",
		Value:      "🦹🏼‍♀️",
		Descriptor: "Woman Supervillain: Medium-Light Skin Tone",
	},
	"1F9B9-1F3FC-200D-2642-FE0F": {
		Key:        "1F9B9-1F3FC-200D-2642-FE0F",
		Value:      "🦹🏼‍♂️",
		Descriptor: "Man Supervillain: Medium-Light Skin Tone",
	},
	"1F9B9-1F3FD": {
		Key:        "1F9B9-1F3FD",
		Value:      "🦹🏽",
		Descriptor: "Supervillain: Medium Skin Tone",
	},
	"1F9B9-1F3FD-200D-2640-FE0F": {
		Key:        "1F9B9-1F3FD-200D-2640-FE0F",
		Value:      "🦹🏽‍♀️",
		Descriptor: "Woman Supervillain: Medium Skin Tone",
	},
	"1F9B9-1F3FD-200D-2642-FE0F": {
		Key:        "1F9B9-1F3FD-200D-2642-FE0F",
		Value:      "🦹🏽‍♂️",
		Descriptor: "Man Supervillain: Medium Skin Tone",
	},
	"1F9B9-1F3FE": {
		Key:        "1F9B9-1F3FE",
		Value:      "🦹🏾",
		Descriptor: "Supervillain: Medium-Dark Skin Tone",
	},
	"1F9B9-1F3FE-200D-2640-FE0F": {
		Key:        "1F9B9-1F3FE-200D-2640-FE0F",
		Value:      "🦹🏾‍♀️",
		Descriptor: "Woman Supervillain: Medium-Dark Skin Tone",
	},
	"1F9B9-1F3FE-200D-2642-FE0F": {
		Key:        "1F9B9-1F3FE-200D-2642-FE0F",
		Value:      "🦹🏾‍♂️",
		Descriptor: "Man Supervillain: Medium-Dark Skin Tone",
	},
	"1F9B9-1F3FF": {
		Key:        "1F9B9-1F3FF",
		Value:      "🦹🏿",
		Descriptor: "Supervillain: Dark Skin Tone",
	},
	"1F9B9-1F3FF-200D-2640-FE0F": {
		Key:        "1F9B9-1F3FF-200D-2640-FE0F",
		Value:      "🦹🏿‍♀️",
		Descriptor: "Woman Supervillain: Dark Skin Tone",
	},
	"1F9B9-1F3FF-200D-2642-FE0F": {
		Key:        "1F9B9-1F3FF-200D-2642-FE0F",
		Value:      "🦹🏿‍♂️",
		Descriptor: "Man Supervillain: Dark Skin Tone",
	},
	"1F9B9-200D-2640-FE0F": {
		Key:        "1F9B9-200D-2640-FE0F",
		Value:      "🦹‍♀️",
		Descriptor: "Woman Supervillain",
	},
	"1F9B9-200D-2642-FE0F": {
		Key:        "1F9B9-200D-2642-FE0F",
		Value:      "🦹‍♂️",
		Descriptor: "Man Supervillain",
	},
	"1F9BA": {
		Key:        "1F9BA",
		Value:      "🦺",
		Descriptor: "Safety Vest",
	},
	"1F9BB": {
		Key:        "1F9BB",
		Value:      "🦻",
		Descriptor: "Ear With Hearing Aid",
	},
	"1F9BB-1F3FB": {
		Key:        "1F9BB-1F3FB",
		Value:      "🦻🏻",
		Descriptor: "Ear With Hearing Aid: Light Skin Tone",
	},
	"1F9BB-1F3FC": {
		Key:        "1F9BB-1F3FC",
		Value:      "🦻🏼",
		Descriptor: "Ear With Hearing Aid: Medium-Light Skin Tone",
	},
	"1F9BB-1F3FD": {
		Key:        "1F9BB-1F3FD",
		Value:      "🦻🏽",
		Descriptor: "Ear With Hearing Aid: Medium Skin Tone",
	},
	"1F9BB-1F3FE": {
		Key:        "1F9BB-1F3FE",
		Value:      "🦻🏾",
		Descriptor: "Ear With Hearing Aid: Medium-Dark Skin Tone",
	},
	"1F9BB-1F3FF": {
		Key:        "1F9BB-1F3FF",
		Value:      "🦻🏿",
		Descriptor: "Ear With Hearing Aid: Dark Skin Tone",
	},
	"1F9BC": {
		Key:        "1F9BC",
		Value:      "🦼",
		Descriptor: "Motorized Wheelchair",
	},
	"1F9BD": {
		Key:        "1F9BD",
		Value:      "🦽",
		Descriptor: "Manual Wheelchair",
	},
	"1F9BE": {
		Key:        "1F9BE",
		Value:      "🦾",
		Descriptor: "Mechanical Arm",
	},
	"1F9BF": {
		Key:        "1F9BF",
		Value:      "🦿",
		Descriptor: "Mechanical Leg",
	},
	"1F9C0": {
		Key:        "1F9C0",
		Value:      "🧀",
		Descriptor: "Cheese Wedge",
	},
	"1F9C1": {
		Key:        "1F9C1",
		Value:      "🧁",
		Descriptor: "Cupcake",
	},
	"1F9C2": {
		Key:        "1F9C2",
		Value:      "🧂",
		Descriptor: "Salt",
	},
	"1F9C3": {
		Key:        "1F9C3",
		Value:      "🧃",
		Descriptor: "Beverage Box",
	},
	"1F9C4": {
		Key:        "1F9C4",
		Value:      "🧄",
		Descriptor: "Garlic",
	},
	"1F9C5": {
		Key:        "1F9C5",
		Value:      "🧅",
		Descriptor: "Onion",
	},
	"1F9C6": {
		Key:        "1F9C6",
		Value:      "🧆",
		Descriptor: "Falafel",
	},
	"1F9C7": {
		Key:        "1F9C7",
		Value:      "🧇",
		Descriptor: "Waffle",
	},
	"1F9C8": {
		Key:        "1F9C8",
		Value:      "🧈",
		Descriptor: "Butter",
	},
	"1F9C9": {
		Key:        "1F9C9",
		Value:      "🧉",
		Descriptor: "Mate",
	},
	"1F9CA": {
		Key:        "1F9CA",
		Value:      "🧊",
		Descriptor: "Ice",
	},
	"1F9CD": {
		Key:        "1F9CD",
		Value:      "🧍",
		Descriptor: "Person Standing",
	},
	"1F9CD-1F3FB": {
		Key:        "1F9CD-1F3FB",
		Value:      "🧍🏻",
		Descriptor: "Person Standing: Light Skin Tone",
	},
	"1F9CD-1F3FB-200D-2640-FE0F": {
		Key:        "1F9CD-1F3FB-200D-2640-FE0F",
		Value:      "🧍🏻‍♀️",
		Descriptor: "Woman Standing: Light Skin Tone",
	},
	"1F9CD-1F3FB-200D-2642-FE0F": {
		Key:        "1F9CD-1F3FB-200D-2642-FE0F",
		Value:      "🧍🏻‍♂️",
		Descriptor: "Man Standing: Light Skin Tone",
	},
	"1F9CD-1F3FC": {
		Key:        "1F9CD-1F3FC",
		Value:      "🧍🏼",
		Descriptor: "Person Standing: Medium-Light Skin Tone",
	},
	"1F9CD-1F3FC-200D-2640-FE0F": {
		Key:        "1F9CD-1F3FC-200D-2640-FE0F",
		Value:      "🧍🏼‍♀️",
		Descriptor: "Woman Standing: Medium-Light Skin Tone",
	},
	"1F9CD-1F3FC-200D-2642-FE0F": {
		Key:        "1F9CD-1F3FC-200D-2642-FE0F",
		Value:      "🧍🏼‍♂️",
		Descriptor: "Man Standing: Medium-Light Skin Tone",
	},
	"1F9CD-1F3FD": {
		Key:        "1F9CD-1F3FD",
		Value:      "🧍🏽",
		Descriptor: "Person Standing: Medium Skin Tone",
	},
	"1F9CD-1F3FD-200D-2640-FE0F": {
		Key:        "1F9CD-1F3FD-200D-2640-FE0F",
		Value:      "🧍🏽‍♀️",
		Descriptor: "Woman Standing: Medium Skin Tone",
	},
	"1F9CD-1F3FD-200D-2642-FE0F": {
		Key:        "1F9CD-1F3FD-200D-2642-FE0F",
		Value:      "🧍🏽‍♂️",
		Descriptor: "Man Standing: Medium Skin Tone",
	},
	"1F9CD-1F3FE": {
		Key:        "1F9CD-1F3FE",
		Value:      "🧍🏾",
		Descriptor: "Person Standing: Medium-Dark Skin Tone",
	},
	"1F9CD-1F3FE-200D-2640-FE0F": {
		Key:        "1F9CD-1F3FE-200D-2640-FE0F",
		Value:      "🧍🏾‍♀️",
		Descriptor: "Woman Standing: Medium-Dark Skin Tone",
	},
	"1F9CD-1F3FE-200D-2642-FE0F": {
		Key:        "1F9CD-1F3FE-200D-2642-FE0F",
		Value:      "🧍🏾‍♂️",
		Descriptor: "Man Standing: Medium-Dark Skin Tone",
	},
	"1F9CD-1F3FF": {
		Key:        "1F9CD-1F3FF",
		Value:      "🧍🏿",
		Descriptor: "Person Standing: Dark Skin Tone",
	},
	"1F9CD-1F3FF-200D-2640-FE0F": {
		Key:        "1F9CD-1F3FF-200D-2640-FE0F",
		Value:      "🧍🏿‍♀️",
		Descriptor: "Woman Standing: Dark Skin Tone",
	},
	"1F9CD-1F3FF-200D-2642-FE0F": {
		Key:        "1F9CD-1F3FF-200D-2642-FE0F",
		Value:      "🧍🏿‍♂️",
		Descriptor: "Man Standing: Dark Skin Tone",
	},
	"1F9CD-200D-2640-FE0F": {
		Key:        "1F9CD-200D-2640-FE0F",
		Value:      "🧍‍♀️",
		Descriptor: "Woman Standing",
	},
	"1F9CD-200D-2642-FE0F": {
		Key:        "1F9CD-200D-2642-FE0F",
		Value:      "🧍‍♂️",
		Descriptor: "Man Standing",
	},
	"1F9CE": {
		Key:        "1F9CE",
		Value:      "🧎",
		Descriptor: "Person Kneeling",
	},
	"1F9CE-1F3FB": {
		Key:        "1F9CE-1F3FB",
		Value:      "🧎🏻",
		Descriptor: "Person Kneeling: Light Skin Tone",
	},
	"1F9CE-1F3FB-200D-2640-FE0F": {
		Key:        "1F9CE-1F3FB-200D-2640-FE0F",
		Value:      "🧎🏻‍♀️",
		Descriptor: "Woman Kneeling: Light Skin Tone",
	},
	"1F9CE-1F3FB-200D-2642-FE0F": {
		Key:        "1F9CE-1F3FB-200D-2642-FE0F",
		Value:      "🧎🏻‍♂️",
		Descriptor: "Man Kneeling: Light Skin Tone",
	},
	"1F9CE-1F3FC": {
		Key:        "1F9CE-1F3FC",
		Value:      "🧎🏼",
		Descriptor: "Person Kneeling: Medium-Light Skin Tone",
	},
	"1F9CE-1F3FC-200D-2640-FE0F": {
		Key:        "1F9CE-1F3FC-200D-2640-FE0F",
		Value:      "🧎🏼‍♀️",
		Descriptor: "Woman Kneeling: Medium-Light Skin Tone",
	},
	"1F9CE-1F3FC-200D-2642-FE0F": {
		Key:        "1F9CE-1F3FC-200D-2642-FE0F",
		Value:      "🧎🏼‍♂️",
		Descriptor: "Man Kneeling: Medium-Light Skin Tone",
	},
	"1F9CE-1F3FD": {
		Key:        "1F9CE-1F3FD",
		Value:      "🧎🏽",
		Descriptor: "Person Kneeling: Medium Skin Tone",
	},
	"1F9CE-1F3FD-200D-2640-FE0F": {
		Key:        "1F9CE-1F3FD-200D-2640-FE0F",
		Value:      "🧎🏽‍♀️",
		Descriptor: "Woman Kneeling: Medium Skin Tone",
	},
	"1F9CE-1F3FD-200D-2642-FE0F": {
		Key:        "1F9CE-1F3FD-200D-2642-FE0F",
		Value:      "🧎🏽‍♂️",
		Descriptor: "Man Kneeling: Medium Skin Tone",
	},
	"1F9CE-1F3FE": {
		Key:        "1F9CE-1F3FE",
		Value:      "🧎🏾",
		Descriptor: "Person Kneeling: Medium-Dark Skin Tone",
	},
	"1F9CE-1F3FE-200D-2640-FE0F": {
		Key:        "1F9CE-1F3FE-200D-2640-FE0F",
		Value:      "🧎🏾‍♀️",
		Descriptor: "Woman Kneeling: Medium-Dark Skin Tone",
	},
	"1F9CE-1F3FE-200D-2642-FE0F": {
		Key:        "1F9CE-1F3FE-200D-2642-FE0F",
		Value:      "🧎🏾‍♂️",
		Descriptor: "Man Kneeling: Medium-Dark Skin Tone",
	},
	"1F9CE-1F3FF": {
		Key:        "1F9CE-1F3FF",
		Value:      "🧎🏿",
		Descriptor: "Person Kneeling: Dark Skin Tone",
	},
	"1F9CE-1F3FF-200D-2640-FE0F": {
		Key:        "1F9CE-1F3FF-200D-2640-FE0F",
		Value:      "🧎🏿‍♀️",
		Descriptor: "Woman Kneeling: Dark Skin Tone",
	},
	"1F9CE-1F3FF-200D-2642-FE0F": {
		Key:        "1F9CE-1F3FF-200D-2642-FE0F",
		Value:      "🧎🏿‍♂️",
		Descriptor: "Man Kneeling: Dark Skin Tone",
	},
	"1F9CE-200D-2640-FE0F": {
		Key:        "1F9CE-200D-2640-FE0F",
		Value:      "🧎‍♀️",
		Descriptor: "Woman Kneeling",
	},
	"1F9CE-200D-2642-FE0F": {
		Key:        "1F9CE-200D-2642-FE0F",
		Value:      "🧎‍♂️",
		Descriptor: "Man Kneeling",
	},
	"1F9CF": {
		Key:        "1F9CF",
		Value:      "🧏",
		Descriptor: "Deaf Person",
	},
	"1F9CF-1F3FB": {
		Key:        "1F9CF-1F3FB",
		Value:      "🧏🏻",
		Descriptor: "Deaf Person: Light Skin Tone",
	},
	"1F9CF-1F3FB-200D-2640-FE0F": {
		Key:        "1F9CF-1F3FB-200D-2640-FE0F",
		Value:      "🧏🏻‍♀️",
		Descriptor: "Deaf Woman: Light Skin Tone",
	},
	"1F9CF-1F3FB-200D-2642-FE0F": {
		Key:        "1F9CF-1F3FB-200D-2642-FE0F",
		Value:      "🧏🏻‍♂️",
		Descriptor: "Deaf Man: Light Skin Tone",
	},
	"1F9CF-1F3FC": {
		Key:        "1F9CF-1F3FC",
		Value:      "🧏🏼",
		Descriptor: "Deaf Person: Medium-Light Skin Tone",
	},
	"1F9CF-1F3FC-200D-2640-FE0F": {
		Key:        "1F9CF-1F3FC-200D-2640-FE0F",
		Value:      "🧏🏼‍♀️",
		Descriptor: "Deaf Woman: Medium-Light Skin Tone",
	},
	"1F9CF-1F3FC-200D-2642-FE0F": {
		Key:        "1F9CF-1F3FC-200D-2642-FE0F",
		Value:      "🧏🏼‍♂️",
		Descriptor: "Deaf Man: Medium-Light Skin Tone",
	},
	"1F9CF-1F3FD": {
		Key:        "1F9CF-1F3FD",
		Value:      "🧏🏽",
		Descriptor: "Deaf Person: Medium Skin Tone",
	},
	"1F9CF-1F3FD-200D-2640-FE0F": {
		Key:        "1F9CF-1F3FD-200D-2640-FE0F",
		Value:      "🧏🏽‍♀️",
		Descriptor: "Deaf Woman: Medium Skin Tone",
	},
	"1F9CF-1F3FD-200D-2642-FE0F": {
		Key:        "1F9CF-1F3FD-200D-2642-FE0F",
		Value:      "🧏🏽‍♂️",
		Descriptor: "Deaf Man: Medium Skin Tone",
	},
	"1F9CF-1F3FE": {
		Key:        "1F9CF-1F3FE",
		Value:      "🧏🏾",
		Descriptor: "Deaf Person: Medium-Dark Skin Tone",
	},
	"1F9CF-1F3FE-200D-2640-FE0F": {
		Key:        "1F9CF-1F3FE-200D-2640-FE0F",
		Value:      "🧏🏾‍♀️",
		Descriptor: "Deaf Woman: Medium-Dark Skin Tone",
	},
	"1F9CF-1F3FE-200D-2642-FE0F": {
		Key:        "1F9CF-1F3FE-200D-2642-FE0F",
		Value:      "🧏🏾‍♂️",
		Descriptor: "Deaf Man: Medium-Dark Skin Tone",
	},
	"1F9CF-1F3FF": {
		Key:        "1F9CF-1F3FF",
		Value:      "🧏🏿",
		Descriptor: "Deaf Person: Dark Skin Tone",
	},
	"1F9CF-1F3FF-200D-2640-FE0F": {
		Key:        "1F9CF-1F3FF-200D-2640-FE0F",
		Value:      "🧏🏿‍♀️",
		Descriptor: "Deaf Woman: Dark Skin Tone",
	},
	"1F9CF-1F3FF-200D-2642-FE0F": {
		Key:        "1F9CF-1F3FF-200D-2642-FE0F",
		Value:      "🧏🏿‍♂️",
		Descriptor: "Deaf Man: Dark Skin Tone",
	},
	"1F9CF-200D-2640-FE0F": {
		Key:        "1F9CF-200D-2640-FE0F",
		Value:      "🧏‍♀️",
		Descriptor: "Deaf Woman",
	},
	"1F9CF-200D-2642-FE0F": {
		Key:        "1F9CF-200D-2642-FE0F",
		Value:      "🧏‍♂️",
		Descriptor: "Deaf Man",
	},
	"1F9D0": {
		Key:        "1F9D0",
		Value:      "🧐",
		Descriptor: "Face With Monocle",
	},
	"1F9D1": {
		Key:        "1F9D1",
		Value:      "🧑",
		Descriptor: "Person",
	},
	"1F9D1-1F3FB": {
		Key:        "1F9D1-1F3FB",
		Value:      "🧑🏻",
		Descriptor: "Person: Light Skin Tone",
	},
	"1F9D1-1F3FB-200D-1F33E": {
		Key:        "1F9D1-1F3FB-200D-1F33E",
		Value:      "🧑🏻‍🌾",
		Descriptor: "Farmer: Light Skin Tone",
	},
	"1F9D1-1F3FB-200D-1F373": {
		Key:        "1F9D1-1F3FB-200D-1F373",
		Value:      "🧑🏻‍🍳",
		Descriptor: "Cook: Light Skin Tone",
	},
	"1F9D1-1F3FB-200D-1F393": {
		Key:        "1F9D1-1F3FB-200D-1F393",
		Value:      "🧑🏻‍🎓",
		Descriptor: "Student: Light Skin Tone",
	},
	"1F9D1-1F3FB-200D-1F3A4": {
		Key:        "1F9D1-1F3FB-200D-1F3A4",
		Value:      "🧑🏻‍🎤",
		Descriptor: "Singer: Light Skin Tone",
	},
	"1F9D1-1F3FB-200D-1F3A8": {
		Key:        "1F9D1-1F3FB-200D-1F3A8",
		Value:      "🧑🏻‍🎨",
		Descriptor: "Artist: Light Skin Tone",
	},
	"1F9D1-1F3FB-200D-1F3EB": {
		Key:        "1F9D1-1F3FB-200D-1F3EB",
		Value:      "🧑🏻‍🏫",
		Descriptor: "Teacher: Light Skin Tone",
	},
	"1F9D1-1F3FB-200D-1F3ED": {
		Key:        "1F9D1-1F3FB-200D-1F3ED",
		Value:      "🧑🏻‍🏭",
		Descriptor: "Factory Worker: Light Skin Tone",
	},
	"1F9D1-1F3FB-200D-1F4BB": {
		Key:        "1F9D1-1F3FB-200D-1F4BB",
		Value:      "🧑🏻‍💻",
		Descriptor: "Technologist: Light Skin Tone",
	},
	"1F9D1-1F3FB-200D-1F4BC": {
		Key:        "1F9D1-1F3FB-200D-1F4BC",
		Value:      "🧑🏻‍💼",
		Descriptor: "Office Worker: Light Skin Tone",
	},
	"1F9D1-1F3FB-200D-1F527": {
		Key:        "1F9D1-1F3FB-200D-1F527",
		Value:      "🧑🏻‍🔧",
		Descriptor: "Mechanic: Light Skin Tone",
	},
	"1F9D1-1F3FB-200D-1F52C": {
		Key:        "1F9D1-1F3FB-200D-1F52C",
		Value:      "🧑🏻‍🔬",
		Descriptor: "Scientist: Light Skin Tone",
	},
	"1F9D1-1F3FB-200D-1F680": {
		Key:        "1F9D1-1F3FB-200D-1F680",
		Value:      "🧑🏻‍🚀",
		Descriptor: "Astronaut: Light Skin Tone",
	},
	"1F9D1-1F3FB-200D-1F692": {
		Key:        "1F9D1-1F3FB-200D-1F692",
		Value:      "🧑🏻‍🚒",
		Descriptor: "Firefighter: Light Skin Tone",
	},
	"1F9D1-1F3FB-200D-1F91D-200D-1F9D1-1F3FB": {
		Key:        "1F9D1-1F3FB-200D-1F91D-200D-1F9D1-1F3FB",
		Value:      "🧑🏻‍🤝‍🧑🏻",
		Descriptor: "People Holding Hands: Light Skin Tone",
	},
	"1F9D1-1F3FB-200D-1F91D-200D-1F9D1-1F3FC": {
		Key:        "1F9D1-1F3FB-200D-1F91D-200D-1F9D1-1F3FC",
		Value:      "🧑🏻‍🤝‍🧑🏼",
		Descriptor: "People Holding Hands: Light Skin Tone, Medium-Light Skin Tone",
	},
	"1F9D1-1F3FB-200D-1F91D-200D-1F9D1-1F3FD": {
		Key:        "1F9D1-1F3FB-200D-1F91D-200D-1F9D1-1F3FD",
		Value:      "🧑🏻‍🤝‍🧑🏽",
		Descriptor: "People Holding Hands: Light Skin Tone, Medium Skin Tone",
	},
	"1F9D1-1F3FB-200D-1F91D-200D-1F9D1-1F3FE": {
		Key:        "1F9D1-1F3FB-200D-1F91D-200D-1F9D1-1F3FE",
		Value:      "🧑🏻‍🤝‍🧑🏾",
		Descriptor: "People Holding Hands: Light Skin Tone, Medium-Dark Skin Tone",
	},
	"1F9D1-1F3FB-200D-1F91D-200D-1F9D1-1F3FF": {
		Key:        "1F9D1-1F3FB-200D-1F91D-200D-1F9D1-1F3FF",
		Value:      "🧑🏻‍🤝‍🧑🏿",
		Descriptor: "People Holding Hands: Light Skin Tone, Dark Skin Tone",
	},
	"1F9D1-1F3FB-200D-1F9AF": {
		Key:        "1F9D1-1F3FB-200D-1F9AF",
		Value:      "🧑🏻‍🦯",
		Descriptor: "Person With Probing Cane: Light Skin Tone",
	},
	"1F9D1-1F3FB-200D-1F9B0": {
		Key:        "1F9D1-1F3FB-200D-1F9B0",
		Value:      "🧑🏻‍🦰",
		Descriptor: "Person: Light Skin Tone, Red Hair",
	},
	"1F9D1-1F3FB-200D-1F9B1": {
		Key:        "1F9D1-1F3FB-200D-1F9B1",
		Value:      "🧑🏻‍🦱",
		Descriptor: "Person: Light Skin Tone, Curly Hair",
	},
	"1F9D1-1F3FB-200D-1F9B2": {
		Key:        "1F9D1-1F3FB-200D-1F9B2",
		Value:      "🧑🏻‍🦲",
		Descriptor: "Person: Light Skin Tone, Bald",
	},
	"1F9D1-1F3FB-200D-1F9B3": {
		Key:        "1F9D1-1F3FB-200D-1F9B3",
		Value:      "🧑🏻‍🦳",
		Descriptor: "Person: Light Skin Tone, White Hair",
	},
	"1F9D1-1F3FB-200D-1F9BC": {
		Key:        "1F9D1-1F3FB-200D-1F9BC",
		Value:      "🧑🏻‍🦼",
		Descriptor: "Person in Motorized Wheelchair: Light Skin Tone",
	},
	"1F9D1-1F3FB-200D-1F9BD": {
		Key:        "1F9D1-1F3FB-200D-1F9BD",
		Value:      "🧑🏻‍🦽",
		Descriptor: "Person in Manual Wheelchair: Light Skin Ton",
	},
	"1F9D1-1F3FB-200D-2695-FE0F": {
		Key:        "1F9D1-1F3FB-200D-2695-FE0F",
		Value:      "🧑🏻‍⚕️",
		Descriptor: "Health Worker: Light Skin Tone",
	},
	"1F9D1-1F3FB-200D-2696-FE0F": {
		Key:        "1F9D1-1F3FB-200D-2696-FE0F",
		Value:      "🧑🏻‍⚖️",
		Descriptor: "Judge: Light Skin Tone",
	},
	"1F9D1-1F3FB-200D-2708-FE0F": {
		Key:        "1F9D1-1F3FB-200D-2708-FE0F",
		Value:      "🧑🏻‍✈️",
		Descriptor: "Pilot: Light Skin Tone",
	},
	"1F9D1-1F3FC": {
		Key:        "1F9D1-1F3FC",
		Value:      "🧑🏼",
		Descriptor: "Person: Medium-Light Skin Tone",
	},
	"1F9D1-1F3FC-200D-1F33E": {
		Key:        "1F9D1-1F3FC-200D-1F33E",
		Value:      "🧑🏼‍🌾",
		Descriptor: "Farmer: Medium-Light Skin Tone",
	},
	"1F9D1-1F3FC-200D-1F373": {
		Key:        "1F9D1-1F3FC-200D-1F373",
		Value:      "🧑🏼‍🍳",
		Descriptor: "Cook: Medium-Light Skin Tone",
	},
	"1F9D1-1F3FC-200D-1F393": {
		Key:        "1F9D1-1F3FC-200D-1F393",
		Value:      "🧑🏼‍🎓",
		Descriptor: "Student: Medium-Light Skin Tone",
	},
	"1F9D1-1F3FC-200D-1F3A4": {
		Key:        "1F9D1-1F3FC-200D-1F3A4",
		Value:      "🧑🏼‍🎤",
		Descriptor: "Singer: Medium-Light Skin Tone",
	},
	"1F9D1-1F3FC-200D-1F3A8": {
		Key:        "1F9D1-1F3FC-200D-1F3A8",
		Value:      "🧑🏼‍🎨",
		Descriptor: "Artist: Medium-Light Skin Tone",
	},
	"1F9D1-1F3FC-200D-1F3EB": {
		Key:        "1F9D1-1F3FC-200D-1F3EB",
		Value:      "🧑🏼‍🏫",
		Descriptor: "Teacher: Medium-Light Skin Tone",
	},
	"1F9D1-1F3FC-200D-1F3ED": {
		Key:        "1F9D1-1F3FC-200D-1F3ED",
		Value:      "🧑🏼‍🏭",
		Descriptor: "Factory Worker: Medium-Light Skin Tone",
	},
	"1F9D1-1F3FC-200D-1F4BB": {
		Key:        "1F9D1-1F3FC-200D-1F4BB",
		Value:      "🧑🏼‍💻",
		Descriptor: "Technologist: Medium-Light Skin Tone",
	},
	"1F9D1-1F3FC-200D-1F4BC": {
		Key:        "1F9D1-1F3FC-200D-1F4BC",
		Value:      "🧑🏼‍💼",
		Descriptor: "Office Worker: Medium-Light Skin Tone",
	},
	"1F9D1-1F3FC-200D-1F527": {
		Key:        "1F9D1-1F3FC-200D-1F527",
		Value:      "🧑🏼‍🔧",
		Descriptor: "Mechanic: Medium-Light Skin Tone",
	},
	"1F9D1-1F3FC-200D-1F52C": {
		Key:        "1F9D1-1F3FC-200D-1F52C",
		Value:      "🧑🏼‍🔬",
		Descriptor: "Scientist: Medium-Light Skin Tone",
	},
	"1F9D1-1F3FC-200D-1F680": {
		Key:        "1F9D1-1F3FC-200D-1F680",
		Value:      "🧑🏼‍🚀",
		Descriptor: "Astronaut: Medium-Light Skin Tone",
	},
	"1F9D1-1F3FC-200D-1F692": {
		Key:        "1F9D1-1F3FC-200D-1F692",
		Value:      "🧑🏼‍🚒",
		Descriptor: "Firefighter: Medium-Light Skin Tone",
	},
	"1F9D1-1F3FC-200D-1F91D-200D-1F9D1-1F3FB": {
		Key:        "1F9D1-1F3FC-200D-1F91D-200D-1F9D1-1F3FB",
		Value:      "🧑🏼‍🤝‍🧑🏻",
		Descriptor: "People Holding Hands: Medium-Light Skin Tone, Light Skin Tone",
	},
	"1F9D1-1F3FC-200D-1F91D-200D-1F9D1-1F3FC": {
		Key:        "1F9D1-1F3FC-200D-1F91D-200D-1F9D1-1F3FC",
		Value:      "🧑🏼‍🤝‍🧑🏼",
		Descriptor: "People Holding Hands: Medium-Light Skin Tone",
	},
	"1F9D1-1F3FC-200D-1F91D-200D-1F9D1-1F3FD": {
		Key:        "1F9D1-1F3FC-200D-1F91D-200D-1F9D1-1F3FD",
		Value:      "🧑🏼‍🤝‍🧑🏽",
		Descriptor: "People Holding Hands: Medium-Light Skin Tone, Medium Skin Tone",
	},
	"1F9D1-1F3FC-200D-1F91D-200D-1F9D1-1F3FE": {
		Key:        "1F9D1-1F3FC-200D-1F91D-200D-1F9D1-1F3FE",
		Value:      "🧑🏼‍🤝‍🧑🏾",
		Descriptor: "People Holding Hands: Medium-Light Skin Tone, Medium-Dark Skin Tone",
	},
	"1F9D1-1F3FC-200D-1F91D-200D-1F9D1-1F3FF": {
		Key:        "1F9D1-1F3FC-200D-1F91D-200D-1F9D1-1F3FF",
		Value:      "🧑🏼‍🤝‍🧑🏿",
		Descriptor: "People Holding Hands: Medium-Light Skin Tone, Dark Skin Tone",
	},
	"1F9D1-1F3FC-200D-1F9AF": {
		Key:        "1F9D1-1F3FC-200D-1F9AF",
		Value:      "🧑🏼‍🦯",
		Descriptor: "Person With Probing Cane: Medium-Light Skin Tone",
	},
	"1F9D1-1F3FC-200D-1F9B0": {
		Key:        "1F9D1-1F3FC-200D-1F9B0",
		Value:      "🧑🏼‍🦰",
		Descriptor: "Person: Medium-Light Skin Tone, Red Hair",
	},
	"1F9D1-1F3FC-200D-1F9B1": {
		Key:        "1F9D1-1F3FC-200D-1F9B1",
		Value:      "🧑🏼‍🦱",
		Descriptor: "Person: Medium-Light Skin Tone, Curly Hair",
	},
	"1F9D1-1F3FC-200D-1F9B2": {
		Key:        "1F9D1-1F3FC-200D-1F9B2",
		Value:      "🧑🏼‍🦲",
		Descriptor: "Person: Medium-Light Skin Tone, Bald",
	},
	"1F9D1-1F3FC-200D-1F9B3": {
		Key:        "1F9D1-1F3FC-200D-1F9B3",
		Value:      "🧑🏼‍🦳",
		Descriptor: "Person: Medium-Light Skin Tone, White Hair",
	},
	"1F9D1-1F3FC-200D-1F9BC": {
		Key:        "1F9D1-1F3FC-200D-1F9BC",
		Value:      "🧑🏼‍🦼",
		Descriptor: "Person in Motorized Wheelchair: Medium-Light Skin Tone",
	},
	"1F9D1-1F3FC-200D-1F9BD": {
		Key:        "1F9D1-1F3FC-200D-1F9BD",
		Value:      "🧑🏼‍🦽",
		Descriptor: "Person in Manual Wheelchair: Medium-Light Skin Tone",
	},
	"1F9D1-1F3FC-200D-2695-FE0F": {
		Key:        "1F9D1-1F3FC-200D-2695-FE0F",
		Value:      "🧑🏼‍⚕️",
		Descriptor: "Health Worker: Medium-Light Skin Tone",
	},
	"1F9D1-1F3FC-200D-2696-FE0F": {
		Key:        "1F9D1-1F3FC-200D-2696-FE0F",
		Value:      "🧑🏼‍⚖️",
		Descriptor: "Judge: Medium-Light Skin Tone",
	},
	"1F9D1-1F3FC-200D-2708-FE0F": {
		Key:        "1F9D1-1F3FC-200D-2708-FE0F",
		Value:      "🧑🏼‍✈️",
		Descriptor: "Pilot: Medium-Light Skin Tone",
	},
	"1F9D1-1F3FD": {
		Key:        "1F9D1-1F3FD",
		Value:      "🧑🏽",
		Descriptor: "Person: Medium Skin Tone",
	},
	"1F9D1-1F3FD-200D-1F33E": {
		Key:        "1F9D1-1F3FD-200D-1F33E",
		Value:      "🧑🏽‍🌾",
		Descriptor: "Farmer: Medium Skin Tone",
	},
	"1F9D1-1F3FD-200D-1F373": {
		Key:        "1F9D1-1F3FD-200D-1F373",
		Value:      "🧑🏽‍🍳",
		Descriptor: "Cook: Medium Skin Tone",
	},
	"1F9D1-1F3FD-200D-1F393": {
		Key:        "1F9D1-1F3FD-200D-1F393",
		Value:      "🧑🏽‍🎓",
		Descriptor: "Student: Medium Skin Tone",
	},
	"1F9D1-1F3FD-200D-1F3A4": {
		Key:        "1F9D1-1F3FD-200D-1F3A4",
		Value:      "🧑🏽‍🎤",
		Descriptor: "Singer: Medium Skin Tone",
	},
	"1F9D1-1F3FD-200D-1F3A8": {
		Key:        "1F9D1-1F3FD-200D-1F3A8",
		Value:      "🧑🏽‍🎨",
		Descriptor: "Artist: Medium Skin Tone",
	},
	"1F9D1-1F3FD-200D-1F3EB": {
		Key:        "1F9D1-1F3FD-200D-1F3EB",
		Value:      "🧑🏽‍🏫",
		Descriptor: "Teacher: Medium Skin Tone",
	},
	"1F9D1-1F3FD-200D-1F3ED": {
		Key:        "1F9D1-1F3FD-200D-1F3ED",
		Value:      "🧑🏽‍🏭",
		Descriptor: "Factory Worker: Medium Skin Tone",
	},
	"1F9D1-1F3FD-200D-1F4BB": {
		Key:        "1F9D1-1F3FD-200D-1F4BB",
		Value:      "🧑🏽‍💻",
		Descriptor: "Technologist: Medium Skin Tone",
	},
	"1F9D1-1F3FD-200D-1F4BC": {
		Key:        "1F9D1-1F3FD-200D-1F4BC",
		Value:      "🧑🏽‍💼",
		Descriptor: "Office Worker: Medium Skin Tone",
	},
	"1F9D1-1F3FD-200D-1F527": {
		Key:        "1F9D1-1F3FD-200D-1F527",
		Value:      "🧑🏽‍🔧",
		Descriptor: "Mechanic: Medium Skin Tone",
	},
	"1F9D1-1F3FD-200D-1F52C": {
		Key:        "1F9D1-1F3FD-200D-1F52C",
		Value:      "🧑🏽‍🔬",
		Descriptor: "Scientist: Medium Skin Tone",
	},
	"1F9D1-1F3FD-200D-1F680": {
		Key:        "1F9D1-1F3FD-200D-1F680",
		Value:      "🧑🏽‍🚀",
		Descriptor: "Astronaut: Medium Skin Tone",
	},
	"1F9D1-1F3FD-200D-1F692": {
		Key:        "1F9D1-1F3FD-200D-1F692",
		Value:      "🧑🏽‍🚒",
		Descriptor: "Firefighter: Medium Skin Tone",
	},
	"1F9D1-1F3FD-200D-1F91D-200D-1F9D1-1F3FB": {
		Key:        "1F9D1-1F3FD-200D-1F91D-200D-1F9D1-1F3FB",
		Value:      "🧑🏽‍🤝‍🧑🏻",
		Descriptor: "People Holding Hands: Medium Skin Tone, Light Skin Tone",
	},
	"1F9D1-1F3FD-200D-1F91D-200D-1F9D1-1F3FC": {
		Key:        "1F9D1-1F3FD-200D-1F91D-200D-1F9D1-1F3FC",
		Value:      "🧑🏽‍🤝‍🧑🏼",
		Descriptor: "People Holding Hands: Medium Skin Tone, Medium-Light Skin Tone",
	},
	"1F9D1-1F3FD-200D-1F91D-200D-1F9D1-1F3FD": {
		Key:        "1F9D1-1F3FD-200D-1F91D-200D-1F9D1-1F3FD",
		Value:      "🧑🏽‍🤝‍🧑🏽",
		Descriptor: "People Holding Hands: Medium Skin Tone",
	},
	"1F9D1-1F3FD-200D-1F91D-200D-1F9D1-1F3FE": {
		Key:        "1F9D1-1F3FD-200D-1F91D-200D-1F9D1-1F3FE",
		Value:      "🧑🏽‍🤝‍🧑🏾",
		Descriptor: "People Holding Hands: Medium Skin Tone, Medium-Dark Skin Tone",
	},
	"1F9D1-1F3FD-200D-1F91D-200D-1F9D1-1F3FF": {
		Key:        "1F9D1-1F3FD-200D-1F91D-200D-1F9D1-1F3FF",
		Value:      "🧑🏽‍🤝‍🧑🏿",
		Descriptor: "People Holding Hands: Medium Skin Tone, Dark Skin Tone",
	},
	"1F9D1-1F3FD-200D-1F9AF": {
		Key:        "1F9D1-1F3FD-200D-1F9AF",
		Value:      "🧑🏽‍🦯",
		Descriptor: "Person With Probing Cane: Medium Skin Tone",
	},
	"1F9D1-1F3FD-200D-1F9B0": {
		Key:        "1F9D1-1F3FD-200D-1F9B0",
		Value:      "🧑🏽‍🦰",
		Descriptor: "Person: Medium Skin Tone, Red Hair",
	},
	"1F9D1-1F3FD-200D-1F9B1": {
		Key:        "1F9D1-1F3FD-200D-1F9B1",
		Value:      "🧑🏽‍🦱",
		Descriptor: "Person: Medium Skin Tone, Curly Hair",
	},
	"1F9D1-1F3FD-200D-1F9B2": {
		Key:        "1F9D1-1F3FD-200D-1F9B2",
		Value:      "🧑🏽‍🦲",
		Descriptor: "Person: Medium Skin Tone, Bald",
	},
	"1F9D1-1F3FD-200D-1F9B3": {
		Key:        "1F9D1-1F3FD-200D-1F9B3",
		Value:      "🧑🏽‍🦳",
		Descriptor: "Person: Medium Skin Tone, White Hair",
	},
	"1F9D1-1F3FD-200D-1F9BC": {
		Key:        "1F9D1-1F3FD-200D-1F9BC",
		Value:      "🧑🏽‍🦼",
		Descriptor: "Person in Motorized Wheelchair: Medium Skin Tone",
	},
	"1F9D1-1F3FD-200D-1F9BD": {
		Key:        "1F9D1-1F3FD-200D-1F9BD",
		Value:      "🧑🏽‍🦽",
		Descriptor: "Person in Manual Wheelchair: Medium Skin Tone",
	},
	"1F9D1-1F3FD-200D-2695-FE0F": {
		Key:        "1F9D1-1F3FD-200D-2695-FE0F",
		Value:      "🧑🏽‍⚕️",
		Descriptor: "Health Worker: Medium Skin Tone",
	},
	"1F9D1-1F3FD-200D-2696-FE0F": {
		Key:        "1F9D1-1F3FD-200D-2696-FE0F",
		Value:      "🧑🏽‍⚖️",
		Descriptor: "Judge: Medium Skin Tone",
	},
	"1F9D1-1F3FD-200D-2708-FE0F": {
		Key:        "1F9D1-1F3FD-200D-2708-FE0F",
		Value:      "🧑🏽‍✈️",
		Descriptor: "Pilot: Medium Skin Tone",
	},
	"1F9D1-1F3FE": {
		Key:        "1F9D1-1F3FE",
		Value:      "🧑🏾",
		Descriptor: "Person: Medium-Dark Skin Tone",
	},
	"1F9D1-1F3FE-200D-1F33E": {
		Key:        "1F9D1-1F3FE-200D-1F33E",
		Value:      "🧑🏾‍🌾",
		Descriptor: "Farmer: Medium-Dark Skin Tone",
	},
	"1F9D1-1F3FE-200D-1F373": {
		Key:        "1F9D1-1F3FE-200D-1F373",
		Value:      "🧑🏾‍🍳",
		Descriptor: "Cook: Medium-Dark Skin Tone",
	},
	"1F9D1-1F3FE-200D-1F393": {
		Key:        "1F9D1-1F3FE-200D-1F393",
		Value:      "🧑🏾‍🎓",
		Descriptor: "Student: Medium-Dark Skin Tone",
	},
	"1F9D1-1F3FE-200D-1F3A4": {
		Key:        "1F9D1-1F3FE-200D-1F3A4",
		Value:      "🧑🏾‍🎤",
		Descriptor: "Singer: Medium-Dark Skin Tone",
	},
	"1F9D1-1F3FE-200D-1F3A8": {
		Key:        "1F9D1-1F3FE-200D-1F3A8",
		Value:      "🧑🏾‍🎨",
		Descriptor: "Artist: Medium-Dark Skin Tone",
	},
	"1F9D1-1F3FE-200D-1F3EB": {
		Key:        "1F9D1-1F3FE-200D-1F3EB",
		Value:      "🧑🏾‍🏫",
		Descriptor: "Teacher: Medium-Dark Skin Tone",
	},
	"1F9D1-1F3FE-200D-1F3ED": {
		Key:        "1F9D1-1F3FE-200D-1F3ED",
		Value:      "🧑🏾‍🏭",
		Descriptor: "Factory Worker: Medium-Dark Skin Tone",
	},
	"1F9D1-1F3FE-200D-1F4BB": {
		Key:        "1F9D1-1F3FE-200D-1F4BB",
		Value:      "🧑🏾‍💻",
		Descriptor: "Technologist: Medium-Dark Skin Tone",
	},
	"1F9D1-1F3FE-200D-1F4BC": {
		Key:        "1F9D1-1F3FE-200D-1F4BC",
		Value:      "🧑🏾‍💼",
		Descriptor: "Office Worker: Medium-Dark Skin Tone",
	},
	"1F9D1-1F3FE-200D-1F527": {
		Key:        "1F9D1-1F3FE-200D-1F527",
		Value:      "🧑🏾‍🔧",
		Descriptor: "Mechanic: Medium-Dark Skin Tone",
	},
	"1F9D1-1F3FE-200D-1F52C": {
		Key:        "1F9D1-1F3FE-200D-1F52C",
		Value:      "🧑🏾‍🔬",
		Descriptor: "Scientist: Medium-Dark Skin Tone",
	},
	"1F9D1-1F3FE-200D-1F680": {
		Key:        "1F9D1-1F3FE-200D-1F680",
		Value:      "🧑🏾‍🚀",
		Descriptor: "Astronaut: Medium-Dark Skin Tone",
	},
	"1F9D1-1F3FE-200D-1F692": {
		Key:        "1F9D1-1F3FE-200D-1F692",
		Value:      "🧑🏾‍🚒",
		Descriptor: "Firefighter: Medium-Dark Skin Tone",
	},
	"1F9D1-1F3FE-200D-1F91D-200D-1F9D1-1F3FB": {
		Key:        "1F9D1-1F3FE-200D-1F91D-200D-1F9D1-1F3FB",
		Value:      "🧑🏾‍🤝‍🧑🏻",
		Descriptor: "People Holding Hands: Medium-Dark Skin Tone, Light Skin Tone",
	},
	"1F9D1-1F3FE-200D-1F91D-200D-1F9D1-1F3FC": {
		Key:        "1F9D1-1F3FE-200D-1F91D-200D-1F9D1-1F3FC",
		Value:      "🧑🏾‍🤝‍🧑🏼",
		Descriptor: "People Holding Hands: Medium-Dark Skin Tone, Medium-Light Skin Tone",
	},
	"1F9D1-1F3FE-200D-1F91D-200D-1F9D1-1F3FD": {
		Key:        "1F9D1-1F3FE-200D-1F91D-200D-1F9D1-1F3FD",
		Value:      "🧑🏾‍🤝‍🧑🏽",
		Descriptor: "People Holding Hands: Medium-Dark Skin Tone, Medium Skin Tone",
	},
	"1F9D1-1F3FE-200D-1F91D-200D-1F9D1-1F3FE": {
		Key:        "1F9D1-1F3FE-200D-1F91D-200D-1F9D1-1F3FE",
		Value:      "🧑🏾‍🤝‍🧑🏾",
		Descriptor: "People Holding Hands: Medium-Dark Skin Tone",
	},
	"1F9D1-1F3FE-200D-1F91D-200D-1F9D1-1F3FF": {
		Key:        "1F9D1-1F3FE-200D-1F91D-200D-1F9D1-1F3FF",
		Value:      "🧑🏾‍🤝‍🧑🏿",
		Descriptor: "People Holding Hands: Medium-Dark Skin Tone, Dark Skin Tone",
	},
	"1F9D1-1F3FE-200D-1F9AF": {
		Key:        "1F9D1-1F3FE-200D-1F9AF",
		Value:      "🧑🏾‍🦯",
		Descriptor: "Person With Probing Cane: Medium-Dark Skin Tone",
	},
	"1F9D1-1F3FE-200D-1F9B0": {
		Key:        "1F9D1-1F3FE-200D-1F9B0",
		Value:      "🧑🏾‍🦰",
		Descriptor: "Person: Medium-Dark Skin Tone, Red Hair",
	},
	"1F9D1-1F3FE-200D-1F9B1": {
		Key:        "1F9D1-1F3FE-200D-1F9B1",
		Value:      "🧑🏾‍🦱",
		Descriptor: "Person: Medium-Dark Skin Tone, Curly Hair",
	},
	"1F9D1-1F3FE-200D-1F9B2": {
		Key:        "1F9D1-1F3FE-200D-1F9B2",
		Value:      "🧑🏾‍🦲",
		Descriptor: "Person: Medium-Dark Skin Tone, Bald",
	},
	"1F9D1-1F3FE-200D-1F9B3": {
		Key:        "1F9D1-1F3FE-200D-1F9B3",
		Value:      "🧑🏾‍🦳",
		Descriptor: "Person: Medium-Dark Skin Tone, White Hair",
	},
	"1F9D1-1F3FE-200D-1F9BC": {
		Key:        "1F9D1-1F3FE-200D-1F9BC",
		Value:      "🧑🏾‍🦼",
		Descriptor: "Person in Motorized Wheelchair: Medium-Dark Skin Tone",
	},
	"1F9D1-1F3FE-200D-1F9BD": {
		Key:        "1F9D1-1F3FE-200D-1F9BD",
		Value:      "🧑🏾‍🦽",
		Descriptor: "Person in Manual Wheelchair: Medium-Dark Skin Tone",
	},
	"1F9D1-1F3FE-200D-2695-FE0F": {
		Key:        "1F9D1-1F3FE-200D-2695-FE0F",
		Value:      "🧑🏾‍⚕️",
		Descriptor: "Health Worker: Medium-Dark Skin Tone",
	},
	"1F9D1-1F3FE-200D-2696-FE0F": {
		Key:        "1F9D1-1F3FE-200D-2696-FE0F",
		Value:      "🧑🏾‍⚖️",
		Descriptor: "Judge: Medium-Dark Skin Tone",
	},
	"1F9D1-1F3FE-200D-2708-FE0F": {
		Key:        "1F9D1-1F3FE-200D-2708-FE0F",
		Value:      "🧑🏾‍✈️",
		Descriptor: "Pilot: Medium-Dark Skin Tone",
	},
	"1F9D1-1F3FF": {
		Key:        "1F9D1-1F3FF",
		Value:      "🧑🏿",
		Descriptor: "Person: Dark Skin Tone",
	},
	"1F9D1-1F3FF-200D-1F33E": {
		Key:        "1F9D1-1F3FF-200D-1F33E",
		Value:      "🧑🏿‍🌾",
		Descriptor: "Farmer: Dark Skin Tone",
	},
	"1F9D1-1F3FF-200D-1F373": {
		Key:        "1F9D1-1F3FF-200D-1F373",
		Value:      "🧑🏿‍🍳",
		Descriptor: "Cook: Dark Skin Tone",
	},
	"1F9D1-1F3FF-200D-1F393": {
		Key:        "1F9D1-1F3FF-200D-1F393",
		Value:      "🧑🏿‍🎓",
		Descriptor: "Student: Dark Skin Tone",
	},
	"1F9D1-1F3FF-200D-1F3A4": {
		Key:        "1F9D1-1F3FF-200D-1F3A4",
		Value:      "🧑🏿‍🎤",
		Descriptor: "Singer: Dark Skin Tone",
	},
	"1F9D1-1F3FF-200D-1F3A8": {
		Key:        "1F9D1-1F3FF-200D-1F3A8",
		Value:      "🧑🏿‍🎨",
		Descriptor: "Artist: Dark Skin Tone",
	},
	"1F9D1-1F3FF-200D-1F3EB": {
		Key:        "1F9D1-1F3FF-200D-1F3EB",
		Value:      "🧑🏿‍🏫",
		Descriptor: "Teacher: Dark Skin Tone",
	},
	"1F9D1-1F3FF-200D-1F3ED": {
		Key:        "1F9D1-1F3FF-200D-1F3ED",
		Value:      "🧑🏿‍🏭",
		Descriptor: "Factory Worker: Dark Skin Tone",
	},
	"1F9D1-1F3FF-200D-1F4BB": {
		Key:        "1F9D1-1F3FF-200D-1F4BB",
		Value:      "🧑🏿‍💻",
		Descriptor: "Technologist: Dark Skin Tone",
	},
	"1F9D1-1F3FF-200D-1F4BC": {
		Key:        "1F9D1-1F3FF-200D-1F4BC",
		Value:      "🧑🏿‍💼",
		Descriptor: "Office Worker: Dark Skin Tone",
	},
	"1F9D1-1F3FF-200D-1F527": {
		Key:        "1F9D1-1F3FF-200D-1F527",
		Value:      "🧑🏿‍🔧",
		Descriptor: "Mechanic: Dark Skin Tone",
	},
	"1F9D1-1F3FF-200D-1F52C": {
		Key:        "1F9D1-1F3FF-200D-1F52C",
		Value:      "🧑🏿‍🔬",
		Descriptor: "Scientist: Dark Skin Tone",
	},
	"1F9D1-1F3FF-200D-1F680": {
		Key:        "1F9D1-1F3FF-200D-1F680",
		Value:      "🧑🏿‍🚀",
		Descriptor: "Astronaut: Dark Skin Tone",
	},
	"1F9D1-1F3FF-200D-1F692": {
		Key:        "1F9D1-1F3FF-200D-1F692",
		Value:      "🧑🏿‍🚒",
		Descriptor: "Firefighter: Dark Skin Tone",
	},
	"1F9D1-1F3FF-200D-1F91D-200D-1F9D1-1F3FB": {
		Key:        "1F9D1-1F3FF-200D-1F91D-200D-1F9D1-1F3FB",
		Value:      "🧑🏿‍🤝‍🧑🏻",
		Descriptor: "People Holding Hands: Dark Skin Tone, Light Skin Tone",
	},
	"1F9D1-1F3FF-200D-1F91D-200D-1F9D1-1F3FC": {
		Key:        "1F9D1-1F3FF-200D-1F91D-200D-1F9D1-1F3FC",
		Value:      "🧑🏿‍🤝‍🧑🏼",
		Descriptor: "People Holding Hands: Dark Skin Tone, Medium-Light Skin Tone",
	},
	"1F9D1-1F3FF-200D-1F91D-200D-1F9D1-1F3FD": {
		Key:        "1F9D1-1F3FF-200D-1F91D-200D-1F9D1-1F3FD",
		Value:      "🧑🏿‍🤝‍🧑🏽",
		Descriptor: "People Holding Hands: Dark Skin Tone, Medium Skin Tone",
	},
	"1F9D1-1F3FF-200D-1F91D-200D-1F9D1-1F3FE": {
		Key:        "1F9D1-1F3FF-200D-1F91D-200D-1F9D1-1F3FE",
		Value:      "🧑🏿‍🤝‍🧑🏾",
		Descriptor: "People Holding Hands: Dark Skin Tone, Medium-Dark Skin Tone",
	},
	"1F9D1-1F3FF-200D-1F91D-200D-1F9D1-1F3FF": {
		Key:        "1F9D1-1F3FF-200D-1F91D-200D-1F9D1-1F3FF",
		Value:      "🧑🏿‍🤝‍🧑🏿",
		Descriptor: "People Holding Hands: Dark Skin Tone",
	},
	"1F9D1-1F3FF-200D-1F9AF": {
		Key:        "1F9D1-1F3FF-200D-1F9AF",
		Value:      "🧑🏿‍🦯",
		Descriptor: "Person With Probing Cane: Dark Skin Tone",
	},
	"1F9D1-1F3FF-200D-1F9B0": {
		Key:        "1F9D1-1F3FF-200D-1F9B0",
		Value:      "🧑🏿‍🦰",
		Descriptor: "Person: Dark Skin Tone, Red Hair",
	},
	"1F9D1-1F3FF-200D-1F9B1": {
		Key:        "1F9D1-1F3FF-200D-1F9B1",
		Value:      "🧑🏿‍🦱",
		Descriptor: "Person: Dark Skin Tone, Curly Hair",
	},
	"1F9D1-1F3FF-200D-1F9B2": {
		Key:        "1F9D1-1F3FF-200D-1F9B2",
		Value:      "🧑🏿‍🦲",
		Descriptor: "Person: Dark Skin Tone, Bald",
	},
	"1F9D1-1F3FF-200D-1F9B3": {
		Key:        "1F9D1-1F3FF-200D-1F9B3",
		Value:      "🧑🏿‍🦳",
		Descriptor: "Person: Dark Skin Tone, White Hair",
	},
	"1F9D1-1F3FF-200D-1F9BC": {
		Key:        "1F9D1-1F3FF-200D-1F9BC",
		Value:      "🧑🏿‍🦼",
		Descriptor: "Person in Motorized Wheelchair: Dark Skin Tone",
	},
	"1F9D1-1F3FF-200D-1F9BD": {
		Key:        "1F9D1-1F3FF-200D-1F9BD",
		Value:      "🧑🏿‍🦽",
		Descriptor: "Person in Manual Wheelchair: Dark Skin Tone",
	},
	"1F9D1-1F3FF-200D-2695-FE0F": {
		Key:        "1F9D1-1F3FF-200D-2695-FE0F",
		Value:      "🧑🏿‍⚕️",
		Descriptor: "Health Worker: Dark Skin Tone",
	},
	"1F9D1-1F3FF-200D-2696-FE0F": {
		Key:        "1F9D1-1F3FF-200D-2696-FE0F",
		Value:      "🧑🏿‍⚖️",
		Descriptor: "Judge: Dark Skin Tone",
	},
	"1F9D1-1F3FF-200D-2708-FE0F": {
		Key:        "1F9D1-1F3FF-200D-2708-FE0F",
		Value:      "🧑🏿‍✈️",
		Descriptor: "Pilot: Dark Skin Tone",
	},
	"1F9D1-200D-1F33E": {
		Key:        "1F9D1-200D-1F33E",
		Value:      "🧑‍🌾",
		Descriptor: "Farmer",
	},
	"1F9D1-200D-1F373": {
		Key:        "1F9D1-200D-1F373",
		Value:      "🧑‍🍳",
		Descriptor: "Cook",
	},
	"1F9D1-200D-1F393": {
		Key:        "1F9D1-200D-1F393",
		Value:      "🧑‍🎓",
		Descriptor: "Student",
	},
	"1F9D1-200D-1F3A4": {
		Key:        "1F9D1-200D-1F3A4",
		Value:      "🧑‍🎤",
		Descriptor: "Singer",
	},
	"1F9D1-200D-1F3A8": {
		Key:        "1F9D1-200D-1F3A8",
		Value:      "🧑‍🎨",
		Descriptor: "Artist",
	},
	"1F9D1-200D-1F3EB": {
		Key:        "1F9D1-200D-1F3EB",
		Value:      "🧑‍🏫",
		Descriptor: "Teacher",
	},
	"1F9D1-200D-1F3ED": {
		Key:        "1F9D1-200D-1F3ED",
		Value:      "🧑‍🏭",
		Descriptor: "Factory Worker",
	},
	"1F9D1-200D-1F4BB": {
		Key:        "1F9D1-200D-1F4BB",
		Value:      "🧑‍💻",
		Descriptor: "Technologist",
	},
	"1F9D1-200D-1F4BC": {
		Key:        "1F9D1-200D-1F4BC",
		Value:      "🧑‍💼",
		Descriptor: "Office Worker",
	},
	"1F9D1-200D-1F527": {
		Key:        "1F9D1-200D-1F527",
		Value:      "🧑‍🔧",
		Descriptor: "Mechanic",
	},
	"1F9D1-200D-1F52C": {
		Key:        "1F9D1-200D-1F52C",
		Value:      "🧑‍🔬",
		Descriptor: "Scientist",
	},
	"1F9D1-200D-1F680": {
		Key:        "1F9D1-200D-1F680",
		Value:      "🧑‍🚀",
		Descriptor: "Astronaut",
	},
	"1F9D1-200D-1F692": {
		Key:        "1F9D1-200D-1F692",
		Value:      "🧑‍🚒",
		Descriptor: "Firefighter",
	},
	"1F9D1-200D-1F91D-200D-1F9D1": {
		Key:        "1F9D1-200D-1F91D-200D-1F9D1",
		Value:      "🧑‍🤝‍🧑",
		Descriptor: "People Holding Hands",
	},
	"1F9D1-200D-1F9AF": {
		Key:        "1F9D1-200D-1F9AF",
		Value:      "🧑‍🦯",
		Descriptor: "Person With Probing Cane",
	},
	"1F9D1-200D-1F9B0": {
		Key:        "1F9D1-200D-1F9B0",
		Value:      "🧑‍🦰",
		Descriptor: "Person: Red Hair",
	},
	"1F9D1-200D-1F9B1": {
		Key:        "1F9D1-200D-1F9B1",
		Value:      "🧑‍🦱",
		Descriptor: "Person: Curly Hair",
	},
	"1F9D1-200D-1F9B2": {
		Key:        "1F9D1-200D-1F9B2",
		Value:      "🧑‍🦲",
		Descriptor: "Person: Bald",
	},
	"1F9D1-200D-1F9B3": {
		Key:        "1F9D1-200D-1F9B3",
		Value:      "🧑‍🦳",
		Descriptor: "Person: White Hair",
	},
	"1F9D1-200D-1F9BC": {
		Key:        "1F9D1-200D-1F9BC",
		Value:      "🧑‍🦼",
		Descriptor: "Person in Motorized Wheelchair",
	},
	"1F9D1-200D-1F9BD": {
		Key:        "1F9D1-200D-1F9BD",
		Value:      "🧑‍🦽",
		Descriptor: "Person in Manual Wheelchair",
	},
	"1F9D1-200D-2695-FE0F": {
		Key:        "1F9D1-200D-2695-FE0F",
		Value:      "🧑‍⚕️",
		Descriptor: "Health Worker",
	},
	"1F9D1-200D-2696-FE0F": {
		Key:        "1F9D1-200D-2696-FE0F",
		Value:      "🧑‍⚖️",
		Descriptor: "Judge",
	},
	"1F9D1-200D-2708-FE0F": {
		Key:        "1F9D1-200D-2708-FE0F",
		Value:      "🧑‍✈️",
		Descriptor: "Pilot",
	},
	"1F9D2": {
		Key:        "1F9D2",
		Value:      "🧒",
		Descriptor: "Child",
	},
	"1F9D2-1F3FB": {
		Key:        "1F9D2-1F3FB",
		Value:      "🧒🏻",
		Descriptor: "Child: Light Skin Tone",
	},
	"1F9D2-1F3FC": {
		Key:        "1F9D2-1F3FC",
		Value:      "🧒🏼",
		Descriptor: "Child: Medium-Light Skin Tone",
	},
	"1F9D2-1F3FD": {
		Key:        "1F9D2-1F3FD",
		Value:      "🧒🏽",
		Descriptor: "Child: Medium Skin Tone",
	},
	"1F9D2-1F3FE": {
		Key:        "1F9D2-1F3FE",
		Value:      "🧒🏾",
		Descriptor: "Child: Medium-Dark Skin Tone",
	},
	"1F9D2-1F3FF": {
		Key:        "1F9D2-1F3FF",
		Value:      "🧒🏿",
		Descriptor: "Child: Dark Skin Tone",
	},
	"1F9D3": {
		Key:        "1F9D3",
		Value:      "🧓",
		Descriptor: "Older Person",
	},
	"1F9D3-1F3FB": {
		Key:        "1F9D3-1F3FB",
		Value:      "🧓🏻",
		Descriptor: "Older Person: Light Skin Tone",
	},
	"1F9D3-1F3FC": {
		Key:        "1F9D3-1F3FC",
		Value:      "🧓🏼",
		Descriptor: "Older Person: Medium-Light Skin Tone",
	},
	"1F9D3-1F3FD": {
		Key:        "1F9D3-1F3FD",
		Value:      "🧓🏽",
		Descriptor: "Older Person: Medium Skin Tone",
	},
	"1F9D3-1F3FE": {
		Key:        "1F9D3-1F3FE",
		Value:      "🧓🏾",
		Descriptor: "Older Person: Medium-Dark Skin Tone",
	},
	"1F9D3-1F3FF": {
		Key:        "1F9D3-1F3FF",
		Value:      "🧓🏿",
		Descriptor: "Older Person: Dark Skin Tone",
	},
	"1F9D4": {
		Key:        "1F9D4",
		Value:      "🧔",
		Descriptor: "Man: Beard",
	},
	"1F9D4-1F3FB": {
		Key:        "1F9D4-1F3FB",
		Value:      "🧔🏻",
		Descriptor: "Man: Light Skin Tone, Beard",
	},
	"1F9D4-1F3FC": {
		Key:        "1F9D4-1F3FC",
		Value:      "🧔🏼",
		Descriptor: "Man: Medium-Light Skin Tone, Beard",
	},
	"1F9D4-1F3FD": {
		Key:        "1F9D4-1F3FD",
		Value:      "🧔🏽",
		Descriptor: "Man: Medium Skin Tone, Beard",
	},
	"1F9D4-1F3FE": {
		Key:        "1F9D4-1F3FE",
		Value:      "🧔🏾",
		Descriptor: "Man: Medium-Dark Skin Tone, Beard",
	},
	"1F9D4-1F3FF": {
		Key:        "1F9D4-1F3FF",
		Value:      "🧔🏿",
		Descriptor: "Man: Dark Skin Tone, Beard",
	},
	"1F9D5": {
		Key:        "1F9D5",
		Value:      "🧕",
		Descriptor: "Woman With Headscarf",
	},
	"1F9D5-1F3FB": {
		Key:        "1F9D5-1F3FB",
		Value:      "🧕🏻",
		Descriptor: "Woman With Headscarf: Light Skin Tone",
	},
	"1F9D5-1F3FC": {
		Key:        "1F9D5-1F3FC",
		Value:      "🧕🏼",
		Descriptor: "Woman With Headscarf: Medium-Light Skin Tone",
	},
	"1F9D5-1F3FD": {
		Key:        "1F9D5-1F3FD",
		Value:      "🧕🏽",
		Descriptor: "Woman With Headscarf: Medium Skin Tone",
	},
	"1F9D5-1F3FE": {
		Key:        "1F9D5-1F3FE",
		Value:      "🧕🏾",
		Descriptor: "Woman With Headscarf: Medium-Dark Skin Tone",
	},
	"1F9D5-1F3FF": {
		Key:        "1F9D5-1F3FF",
		Value:      "🧕🏿",
		Descriptor: "Woman With Headscarf: Dark Skin Tone",
	},
	"1F9D6": {
		Key:        "1F9D6",
		Value:      "🧖",
		Descriptor: "Person in Steamy Room",
	},
	"1F9D6-1F3FB": {
		Key:        "1F9D6-1F3FB",
		Value:      "🧖🏻",
		Descriptor: "Person in Steamy Room: Light Skin Tone",
	},
	"1F9D6-1F3FB-200D-2640-FE0F": {
		Key:        "1F9D6-1F3FB-200D-2640-FE0F",
		Value:      "🧖🏻‍♀️",
		Descriptor: "Woman in Steamy Room: Light Skin Tone",
	},
	"1F9D6-1F3FB-200D-2642-FE0F": {
		Key:        "1F9D6-1F3FB-200D-2642-FE0F",
		Value:      "🧖🏻‍♂️",
		Descriptor: "Man in Steamy Room: Light Skin Tone",
	},
	"1F9D6-1F3FC": {
		Key:        "1F9D6-1F3FC",
		Value:      "🧖🏼",
		Descriptor: "Person in Steamy Room: Medium-Light Skin Tone",
	},
	"1F9D6-1F3FC-200D-2640-FE0F": {
		Key:        "1F9D6-1F3FC-200D-2640-FE0F",
		Value:      "🧖🏼‍♀️",
		Descriptor: "Woman in Steamy Room: Medium-Light Skin Tone",
	},
	"1F9D6-1F3FC-200D-2642-FE0F": {
		Key:        "1F9D6-1F3FC-200D-2642-FE0F",
		Value:      "🧖🏼‍♂️",
		Descriptor: "Man in Steamy Room: Medium-Light Skin Tone",
	},
	"1F9D6-1F3FD": {
		Key:        "1F9D6-1F3FD",
		Value:      "🧖🏽",
		Descriptor: "Person in Steamy Room: Medium Skin Tone",
	},
	"1F9D6-1F3FD-200D-2640-FE0F": {
		Key:        "1F9D6-1F3FD-200D-2640-FE0F",
		Value:      "🧖🏽‍♀️",
		Descriptor: "Woman in Steamy Room: Medium Skin Tone",
	},
	"1F9D6-1F3FD-200D-2642-FE0F": {
		Key:        "1F9D6-1F3FD-200D-2642-FE0F",
		Value:      "🧖🏽‍♂️",
		Descriptor: "Man in Steamy Room: Medium Skin Tone",
	},
	"1F9D6-1F3FE": {
		Key:        "1F9D6-1F3FE",
		Value:      "🧖🏾",
		Descriptor: "Person in Steamy Room: Medium-Dark Skin Tone",
	},
	"1F9D6-1F3FE-200D-2640-FE0F": {
		Key:        "1F9D6-1F3FE-200D-2640-FE0F",
		Value:      "🧖🏾‍♀️",
		Descriptor: "Woman in Steamy Room: Medium-Dark Skin Tone",
	},
	"1F9D6-1F3FE-200D-2642-FE0F": {
		Key:        "1F9D6-1F3FE-200D-2642-FE0F",
		Value:      "🧖🏾‍♂️",
		Descriptor: "Man in Steamy Room: Medium-Dark Skin Tone",
	},
	"1F9D6-1F3FF": {
		Key:        "1F9D6-1F3FF",
		Value:      "🧖🏿",
		Descriptor: "Person in Steamy Room: Dark Skin Tone",
	},
	"1F9D6-1F3FF-200D-2640-FE0F": {
		Key:        "1F9D6-1F3FF-200D-2640-FE0F",
		Value:      "🧖🏿‍♀️",
		Descriptor: "Woman in Steamy Room: Dark Skin Tone",
	},
	"1F9D6-1F3FF-200D-2642-FE0F": {
		Key:        "1F9D6-1F3FF-200D-2642-FE0F",
		Value:      "🧖🏿‍♂️",
		Descriptor: "Man in Steamy Room: Dark Skin Tone",
	},
	"1F9D6-200D-2640-FE0F": {
		Key:        "1F9D6-200D-2640-FE0F",
		Value:      "🧖‍♀️",
		Descriptor: "Woman in Steamy Room",
	},
	"1F9D6-200D-2642-FE0F": {
		Key:        "1F9D6-200D-2642-FE0F",
		Value:      "🧖‍♂️",
		Descriptor: "Man in Steamy Room",
	},
	"1F9D7": {
		Key:        "1F9D7",
		Value:      "🧗",
		Descriptor: "Person Climbing",
	},
	"1F9D7-1F3FB": {
		Key:        "1F9D7-1F3FB",
		Value:      "🧗🏻",
		Descriptor: "Person Climbing: Light Skin Tone",
	},
	"1F9D7-1F3FB-200D-2640-FE0F": {
		Key:        "1F9D7-1F3FB-200D-2640-FE0F",
		Value:      "🧗🏻‍♀️",
		Descriptor: "Woman Climbing: Light Skin Tone",
	},
	"1F9D7-1F3FB-200D-2642-FE0F": {
		Key:        "1F9D7-1F3FB-200D-2642-FE0F",
		Value:      "🧗🏻‍♂️",
		Descriptor: "Man Climbing: Light Skin Tone",
	},
	"1F9D7-1F3FC": {
		Key:        "1F9D7-1F3FC",
		Value:      "🧗🏼",
		Descriptor: "Person Climbing: Medium-Light Skin Tone",
	},
	"1F9D7-1F3FC-200D-2640-FE0F": {
		Key:        "1F9D7-1F3FC-200D-2640-FE0F",
		Value:      "🧗🏼‍♀️",
		Descriptor: "Woman Climbing: Medium-Light Skin Tone",
	},
	"1F9D7-1F3FC-200D-2642-FE0F": {
		Key:        "1F9D7-1F3FC-200D-2642-FE0F",
		Value:      "🧗🏼‍♂️",
		Descriptor: "Man Climbing: Medium-Light Skin Tone",
	},
	"1F9D7-1F3FD": {
		Key:        "1F9D7-1F3FD",
		Value:      "🧗🏽",
		Descriptor: "Person Climbing: Medium Skin Tone",
	},
	"1F9D7-1F3FD-200D-2640-FE0F": {
		Key:        "1F9D7-1F3FD-200D-2640-FE0F",
		Value:      "🧗🏽‍♀️",
		Descriptor: "Woman Climbing: Medium Skin Tone",
	},
	"1F9D7-1F3FD-200D-2642-FE0F": {
		Key:        "1F9D7-1F3FD-200D-2642-FE0F",
		Value:      "🧗🏽‍♂️",
		Descriptor: "Man Climbing: Medium Skin Tone",
	},
	"1F9D7-1F3FE": {
		Key:        "1F9D7-1F3FE",
		Value:      "🧗🏾",
		Descriptor: "Person Climbing: Medium-Dark Skin Tone",
	},
	"1F9D7-1F3FE-200D-2640-FE0F": {
		Key:        "1F9D7-1F3FE-200D-2640-FE0F",
		Value:      "🧗🏾‍♀️",
		Descriptor: "Woman Climbing: Medium-Dark Skin Tone",
	},
	"1F9D7-1F3FE-200D-2642-FE0F": {
		Key:        "1F9D7-1F3FE-200D-2642-FE0F",
		Value:      "🧗🏾‍♂️",
		Descriptor: "Man Climbing: Medium-Dark Skin Tone",
	},
	"1F9D7-1F3FF": {
		Key:        "1F9D7-1F3FF",
		Value:      "🧗🏿",
		Descriptor: "Person Climbing: Dark Skin Tone",
	},
	"1F9D7-1F3FF-200D-2640-FE0F": {
		Key:        "1F9D7-1F3FF-200D-2640-FE0F",
		Value:      "🧗🏿‍♀️",
		Descriptor: "Woman Climbing: Dark Skin Tone",
	},
	"1F9D7-1F3FF-200D-2642-FE0F": {
		Key:        "1F9D7-1F3FF-200D-2642-FE0F",
		Value:      "🧗🏿‍♂️",
		Descriptor: "Man Climbing: Dark Skin Tone",
	},
	"1F9D7-200D-2640-FE0F": {
		Key:        "1F9D7-200D-2640-FE0F",
		Value:      "🧗‍♀️",
		Descriptor: "Woman Climbing",
	},
	"1F9D7-200D-2642-FE0F": {
		Key:        "1F9D7-200D-2642-FE0F",
		Value:      "🧗‍♂️",
		Descriptor: "Man Climbing",
	},
	"1F9D8": {
		Key:        "1F9D8",
		Value:      "🧘",
		Descriptor: "Person in Lotus Position",
	},
	"1F9D8-1F3FB": {
		Key:        "1F9D8-1F3FB",
		Value:      "🧘🏻",
		Descriptor: "Person in Lotus Position: Light Skin Tone",
	},
	"1F9D8-1F3FB-200D-2640-FE0F": {
		Key:        "1F9D8-1F3FB-200D-2640-FE0F",
		Value:      "🧘🏻‍♀️",
		Descriptor: "Woman in Lotus Position: Light Skin Tone",
	},
	"1F9D8-1F3FB-200D-2642-FE0F": {
		Key:        "1F9D8-1F3FB-200D-2642-FE0F",
		Value:      "🧘🏻‍♂️",
		Descriptor: "Man in Lotus Position: Light Skin Tone",
	},
	"1F9D8-1F3FC": {
		Key:        "1F9D8-1F3FC",
		Value:      "🧘🏼",
		Descriptor: "Person in Lotus Position: Medium-Light Skin Tone",
	},
	"1F9D8-1F3FC-200D-2640-FE0F": {
		Key:        "1F9D8-1F3FC-200D-2640-FE0F",
		Value:      "🧘🏼‍♀️",
		Descriptor: "Woman in Lotus Position: Medium-Light Skin Tone",
	},
	"1F9D8-1F3FC-200D-2642-FE0F": {
		Key:        "1F9D8-1F3FC-200D-2642-FE0F",
		Value:      "🧘🏼‍♂️",
		Descriptor: "Man in Lotus Position: Medium-Light Skin Tone",
	},
	"1F9D8-1F3FD": {
		Key:        "1F9D8-1F3FD",
		Value:      "🧘🏽",
		Descriptor: "Person in Lotus Position: Medium Skin Tone",
	},
	"1F9D8-1F3FD-200D-2640-FE0F": {
		Key:        "1F9D8-1F3FD-200D-2640-FE0F",
		Value:      "🧘🏽‍♀️",
		Descriptor: "Woman in Lotus Position: Medium Skin Tone",
	},
	"1F9D8-1F3FD-200D-2642-FE0F": {
		Key:        "1F9D8-1F3FD-200D-2642-FE0F",
		Value:      "🧘🏽‍♂️",
		Descriptor: "Man in Lotus Position: Medium Skin Tone",
	},
	"1F9D8-1F3FE": {
		Key:        "1F9D8-1F3FE",
		Value:      "🧘🏾",
		Descriptor: "Person in Lotus Position: Medium-Dark Skin Tone",
	},
	"1F9D8-1F3FE-200D-2640-FE0F": {
		Key:        "1F9D8-1F3FE-200D-2640-FE0F",
		Value:      "🧘🏾‍♀️",
		Descriptor: "Woman in Lotus Position: Medium-Dark Skin Tone",
	},
	"1F9D8-1F3FE-200D-2642-FE0F": {
		Key:        "1F9D8-1F3FE-200D-2642-FE0F",
		Value:      "🧘🏾‍♂️",
		Descriptor: "Man in Lotus Position: Medium-Dark Skin Tone",
	},
	"1F9D8-1F3FF": {
		Key:        "1F9D8-1F3FF",
		Value:      "🧘🏿",
		Descriptor: "Person in Lotus Position: Dark Skin Tone",
	},
	"1F9D8-1F3FF-200D-2640-FE0F": {
		Key:        "1F9D8-1F3FF-200D-2640-FE0F",
		Value:      "🧘🏿‍♀️",
		Descriptor: "Woman in Lotus Position: Dark Skin Tone",
	},
	"1F9D8-1F3FF-200D-2642-FE0F": {
		Key:        "1F9D8-1F3FF-200D-2642-FE0F",
		Value:      "🧘🏿‍♂️",
		Descriptor: "Man in Lotus Position: Dark Skin Tone",
	},
	"1F9D8-200D-2640-FE0F": {
		Key:        "1F9D8-200D-2640-FE0F",
		Value:      "🧘‍♀️",
		Descriptor: "Woman in Lotus Position",
	},
	"1F9D8-200D-2642-FE0F": {
		Key:        "1F9D8-200D-2642-FE0F",
		Value:      "🧘‍♂️",
		Descriptor: "Man in Lotus Position",
	},
	"1F9D9": {
		Key:        "1F9D9",
		Value:      "🧙",
		Descriptor: "Mage",
	},
	"1F9D9-1F3FB": {
		Key:        "1F9D9-1F3FB",
		Value:      "🧙🏻",
		Descriptor: "Mage: Light Skin Tone",
	},
	"1F9D9-1F3FB-200D-2640-FE0F": {
		Key:        "1F9D9-1F3FB-200D-2640-FE0F",
		Value:      "🧙🏻‍♀️",
		Descriptor: "Woman Mage: Light Skin Tone",
	},
	"1F9D9-1F3FB-200D-2642-FE0F": {
		Key:        "1F9D9-1F3FB-200D-2642-FE0F",
		Value:      "🧙🏻‍♂️",
		Descriptor: "Man Mage: Light Skin Tone",
	},
	"1F9D9-1F3FC": {
		Key:        "1F9D9-1F3FC",
		Value:      "🧙🏼",
		Descriptor: "Mage: Medium-Light Skin Tone",
	},
	"1F9D9-1F3FC-200D-2640-FE0F": {
		Key:        "1F9D9-1F3FC-200D-2640-FE0F",
		Value:      "🧙🏼‍♀️",
		Descriptor: "Woman Mage: Medium-Light Skin Tone",
	},
	"1F9D9-1F3FC-200D-2642-FE0F": {
		Key:        "1F9D9-1F3FC-200D-2642-FE0F",
		Value:      "🧙🏼‍♂️",
		Descriptor: "Man Mage: Medium-Light Skin Tone",
	},
	"1F9D9-1F3FD": {
		Key:        "1F9D9-1F3FD",
		Value:      "🧙🏽",
		Descriptor: "Mage: Medium Skin Tone",
	},
	"1F9D9-1F3FD-200D-2640-FE0F": {
		Key:        "1F9D9-1F3FD-200D-2640-FE0F",
		Value:      "🧙🏽‍♀️",
		Descriptor: "Woman Mage: Medium Skin Tone",
	},
	"1F9D9-1F3FD-200D-2642-FE0F": {
		Key:        "1F9D9-1F3FD-200D-2642-FE0F",
		Value:      "🧙🏽‍♂️",
		Descriptor: "Man Mage: Medium Skin Tone",
	},
	"1F9D9-1F3FE": {
		Key:        "1F9D9-1F3FE",
		Value:      "🧙🏾",
		Descriptor: "Mage: Medium-Dark Skin Tone",
	},
	"1F9D9-1F3FE-200D-2640-FE0F": {
		Key:        "1F9D9-1F3FE-200D-2640-FE0F",
		Value:      "🧙🏾‍♀️",
		Descriptor: "Woman Mage: Medium-Dark Skin Tone",
	},
	"1F9D9-1F3FE-200D-2642-FE0F": {
		Key:        "1F9D9-1F3FE-200D-2642-FE0F",
		Value:      "🧙🏾‍♂️",
		Descriptor: "Man Mage: Medium-Dark Skin Tone",
	},
	"1F9D9-1F3FF": {
		Key:        "1F9D9-1F3FF",
		Value:      "🧙🏿",
		Descriptor: "Mage: Dark Skin Tone",
	},
	"1F9D9-1F3FF-200D-2640-FE0F": {
		Key:        "1F9D9-1F3FF-200D-2640-FE0F",
		Value:      "🧙🏿‍♀️",
		Descriptor: "Woman Mage: Dark Skin Tone",
	},
	"1F9D9-1F3FF-200D-2642-FE0F": {
		Key:        "1F9D9-1F3FF-200D-2642-FE0F",
		Value:      "🧙🏿‍♂️",
		Descriptor: "Man Mage: Dark Skin Tone",
	},
	"1F9D9-200D-2640-FE0F": {
		Key:        "1F9D9-200D-2640-FE0F",
		Value:      "🧙‍♀️",
		Descriptor: "Woman Mage",
	},
	"1F9D9-200D-2642-FE0F": {
		Key:        "1F9D9-200D-2642-FE0F",
		Value:      "🧙‍♂️",
		Descriptor: "Man Mage",
	},
	"1F9DA": {
		Key:        "1F9DA",
		Value:      "🧚",
		Descriptor: "Fairy",
	},
	"1F9DA-1F3FB": {
		Key:        "1F9DA-1F3FB",
		Value:      "🧚🏻",
		Descriptor: "Fairy: Light Skin Tone",
	},
	"1F9DA-1F3FB-200D-2640-FE0F": {
		Key:        "1F9DA-1F3FB-200D-2640-FE0F",
		Value:      "🧚🏻‍♀️",
		Descriptor: "Woman Fairy: Light Skin Tone",
	},
	"1F9DA-1F3FB-200D-2642-FE0F": {
		Key:        "1F9DA-1F3FB-200D-2642-FE0F",
		Value:      "🧚🏻‍♂️",
		Descriptor: "Man Fairy: Light Skin Tone",
	},
	"1F9DA-1F3FC": {
		Key:        "1F9DA-1F3FC",
		Value:      "🧚🏼",
		Descriptor: "Fairy: Medium-Light Skin Tone",
	},
	"1F9DA-1F3FC-200D-2640-FE0F": {
		Key:        "1F9DA-1F3FC-200D-2640-FE0F",
		Value:      "🧚🏼‍♀️",
		Descriptor: "Woman Fairy: Medium-Light Skin Tone",
	},
	"1F9DA-1F3FC-200D-2642-FE0F": {
		Key:        "1F9DA-1F3FC-200D-2642-FE0F",
		Value:      "🧚🏼‍♂️",
		Descriptor: "Man Fairy: Medium-Light Skin Tone",
	},
	"1F9DA-1F3FD": {
		Key:        "1F9DA-1F3FD",
		Value:      "🧚🏽",
		Descriptor: "Fairy: Medium Skin Tone",
	},
	"1F9DA-1F3FD-200D-2640-FE0F": {
		Key:        "1F9DA-1F3FD-200D-2640-FE0F",
		Value:      "🧚🏽‍♀️",
		Descriptor: "Woman Fairy: Medium Skin Tone",
	},
	"1F9DA-1F3FD-200D-2642-FE0F": {
		Key:        "1F9DA-1F3FD-200D-2642-FE0F",
		Value:      "🧚🏽‍♂️",
		Descriptor: "Man Fairy: Medium Skin Tone",
	},
	"1F9DA-1F3FE": {
		Key:        "1F9DA-1F3FE",
		Value:      "🧚🏾",
		Descriptor: "Fairy: Medium-Dark Skin Tone",
	},
	"1F9DA-1F3FE-200D-2640-FE0F": {
		Key:        "1F9DA-1F3FE-200D-2640-FE0F",
		Value:      "🧚🏾‍♀️",
		Descriptor: "Woman Fairy: Medium-Dark Skin Tone",
	},
	"1F9DA-1F3FE-200D-2642-FE0F": {
		Key:        "1F9DA-1F3FE-200D-2642-FE0F",
		Value:      "🧚🏾‍♂️",
		Descriptor: "Man Fairy: Medium-Dark Skin Tone",
	},
	"1F9DA-1F3FF": {
		Key:        "1F9DA-1F3FF",
		Value:      "🧚🏿",
		Descriptor: "Fairy: Dark Skin Tone",
	},
	"1F9DA-1F3FF-200D-2640-FE0F": {
		Key:        "1F9DA-1F3FF-200D-2640-FE0F",
		Value:      "🧚🏿‍♀️",
		Descriptor: "Woman Fairy: Dark Skin Tone",
	},
	"1F9DA-1F3FF-200D-2642-FE0F": {
		Key:        "1F9DA-1F3FF-200D-2642-FE0F",
		Value:      "🧚🏿‍♂️",
		Descriptor: "Man Fairy: Dark Skin Tone",
	},
	"1F9DA-200D-2640-FE0F": {
		Key:        "1F9DA-200D-2640-FE0F",
		Value:      "🧚‍♀️",
		Descriptor: "Woman Fairy",
	},
	"1F9DA-200D-2642-FE0F": {
		Key:        "1F9DA-200D-2642-FE0F",
		Value:      "🧚‍♂️",
		Descriptor: "Man Fairy",
	},
	"1F9DB": {
		Key:        "1F9DB",
		Value:      "🧛",
		Descriptor: "Vampire",
	},
	"1F9DB-1F3FB": {
		Key:        "1F9DB-1F3FB",
		Value:      "🧛🏻",
		Descriptor: "Vampire: Light Skin Tone",
	},
	"1F9DB-1F3FB-200D-2640-FE0F": {
		Key:        "1F9DB-1F3FB-200D-2640-FE0F",
		Value:      "🧛🏻‍♀️",
		Descriptor: "Woman Vampire: Light Skin Tone",
	},
	"1F9DB-1F3FB-200D-2642-FE0F": {
		Key:        "1F9DB-1F3FB-200D-2642-FE0F",
		Value:      "🧛🏻‍♂️",
		Descriptor: "Man Vampire: Light Skin Tone",
	},
	"1F9DB-1F3FC": {
		Key:        "1F9DB-1F3FC",
		Value:      "🧛🏼",
		Descriptor: "Vampire: Medium-Light Skin Tone",
	},
	"1F9DB-1F3FC-200D-2640-FE0F": {
		Key:        "1F9DB-1F3FC-200D-2640-FE0F",
		Value:      "🧛🏼‍♀️",
		Descriptor: "Woman Vampire: Medium-Light Skin Tone",
	},
	"1F9DB-1F3FC-200D-2642-FE0F": {
		Key:        "1F9DB-1F3FC-200D-2642-FE0F",
		Value:      "🧛🏼‍♂️",
		Descriptor: "Man Vampire: Medium-Light Skin Tone",
	},
	"1F9DB-1F3FD": {
		Key:        "1F9DB-1F3FD",
		Value:      "🧛🏽",
		Descriptor: "Vampire: Medium Skin Tone",
	},
	"1F9DB-1F3FD-200D-2640-FE0F": {
		Key:        "1F9DB-1F3FD-200D-2640-FE0F",
		Value:      "🧛🏽‍♀️",
		Descriptor: "Woman Vampire: Medium Skin Tone",
	},
	"1F9DB-1F3FD-200D-2642-FE0F": {
		Key:        "1F9DB-1F3FD-200D-2642-FE0F",
		Value:      "🧛🏽‍♂️",
		Descriptor: "Man Vampire: Medium Skin Tone",
	},
	"1F9DB-1F3FE": {
		Key:        "1F9DB-1F3FE",
		Value:      "🧛🏾",
		Descriptor: "Vampire: Medium-Dark Skin Tone",
	},
	"1F9DB-1F3FE-200D-2640-FE0F": {
		Key:        "1F9DB-1F3FE-200D-2640-FE0F",
		Value:      "🧛🏾‍♀️",
		Descriptor: "Woman Vampire: Medium-Dark Skin Tone",
	},
	"1F9DB-1F3FE-200D-2642-FE0F": {
		Key:        "1F9DB-1F3FE-200D-2642-FE0F",
		Value:      "🧛🏾‍♂️",
		Descriptor: "Man Vampire: Medium-Dark Skin Tone",
	},
	"1F9DB-1F3FF": {
		Key:        "1F9DB-1F3FF",
		Value:      "🧛🏿",
		Descriptor: "Vampire: Dark Skin Tone",
	},
	"1F9DB-1F3FF-200D-2640-FE0F": {
		Key:        "1F9DB-1F3FF-200D-2640-FE0F",
		Value:      "🧛🏿‍♀️",
		Descriptor: "Woman Vampire: Dark Skin Tone",
	},
	"1F9DB-1F3FF-200D-2642-FE0F": {
		Key:        "1F9DB-1F3FF-200D-2642-FE0F",
		Value:      "🧛🏿‍♂️",
		Descriptor: "Man Vampire: Dark Skin Tone",
	},
	"1F9DB-200D-2640-FE0F": {
		Key:        "1F9DB-200D-2640-FE0F",
		Value:      "🧛‍♀️",
		Descriptor: "Woman Vampire",
	},
	"1F9DB-200D-2642-FE0F": {
		Key:        "1F9DB-200D-2642-FE0F",
		Value:      "🧛‍♂️",
		Descriptor: "Man Vampire",
	},
	"1F9DC": {
		Key:        "1F9DC",
		Value:      "🧜",
		Descriptor: "Merperson",
	},
	"1F9DC-1F3FB": {
		Key:        "1F9DC-1F3FB",
		Value:      "🧜🏻",
		Descriptor: "Merperson: Light Skin Tone",
	},
	"1F9DC-1F3FB-200D-2640-FE0F": {
		Key:        "1F9DC-1F3FB-200D-2640-FE0F",
		Value:      "🧜🏻‍♀️",
		Descriptor: "Mermaid: Light Skin Tone",
	},
	"1F9DC-1F3FB-200D-2642-FE0F": {
		Key:        "1F9DC-1F3FB-200D-2642-FE0F",
		Value:      "🧜🏻‍♂️",
		Descriptor: "Merman: Light Skin Tone",
	},
	"1F9DC-1F3FC": {
		Key:        "1F9DC-1F3FC",
		Value:      "🧜🏼",
		Descriptor: "Merperson: Medium-Light Skin Tone",
	},
	"1F9DC-1F3FC-200D-2640-FE0F": {
		Key:        "1F9DC-1F3FC-200D-2640-FE0F",
		Value:      "🧜🏼‍♀️",
		Descriptor: "Mermaid: Medium-Light Skin Tone",
	},
	"1F9DC-1F3FC-200D-2642-FE0F": {
		Key:        "1F9DC-1F3FC-200D-2642-FE0F",
		Value:      "🧜🏼‍♂️",
		Descriptor: "Merman: Medium-Light Skin Tone",
	},
	"1F9DC-1F3FD": {
		Key:        "1F9DC-1F3FD",
		Value:      "🧜🏽",
		Descriptor: "Merperson: Medium Skin Tone",
	},
	"1F9DC-1F3FD-200D-2640-FE0F": {
		Key:        "1F9DC-1F3FD-200D-2640-FE0F",
		Value:      "🧜🏽‍♀️",
		Descriptor: "Mermaid: Medium Skin Tone",
	},
	"1F9DC-1F3FD-200D-2642-FE0F": {
		Key:        "1F9DC-1F3FD-200D-2642-FE0F",
		Value:      "🧜🏽‍♂️",
		Descriptor: "Merman: Medium Skin Tone",
	},
	"1F9DC-1F3FE": {
		Key:        "1F9DC-1F3FE",
		Value:      "🧜🏾",
		Descriptor: "Merperson: Medium-Dark Skin Tone",
	},
	"1F9DC-1F3FE-200D-2640-FE0F": {
		Key:        "1F9DC-1F3FE-200D-2640-FE0F",
		Value:      "🧜🏾‍♀️",
		Descriptor: "Mermaid: Medium-Dark Skin Tone",
	},
	"1F9DC-1F3FE-200D-2642-FE0F": {
		Key:        "1F9DC-1F3FE-200D-2642-FE0F",
		Value:      "🧜🏾‍♂️",
		Descriptor: "Merman: Medium-Dark Skin Tone",
	},
	"1F9DC-1F3FF": {
		Key:        "1F9DC-1F3FF",
		Value:      "🧜🏿",
		Descriptor: "Merperson: Dark Skin Tone",
	},
	"1F9DC-1F3FF-200D-2640-FE0F": {
		Key:        "1F9DC-1F3FF-200D-2640-FE0F",
		Value:      "🧜🏿‍♀️",
		Descriptor: "Mermaid: Dark Skin Tone",
	},
	"1F9DC-1F3FF-200D-2642-FE0F": {
		Key:        "1F9DC-1F3FF-200D-2642-FE0F",
		Value:      "🧜🏿‍♂️",
		Descriptor: "Merman: Dark Skin Tone",
	},
	"1F9DC-200D-2640-FE0F": {
		Key:        "1F9DC-200D-2640-FE0F",
		Value:      "🧜‍♀️",
		Descriptor: "Mermaid",
	},
	"1F9DC-200D-2642-FE0F": {
		Key:        "1F9DC-200D-2642-FE0F",
		Value:      "🧜‍♂️",
		Descriptor: "Merman",
	},
	"1F9DD": {
		Key:        "1F9DD",
		Value:      "🧝",
		Descriptor: "Elf",
	},
	"1F9DD-1F3FB": {
		Key:        "1F9DD-1F3FB",
		Value:      "🧝🏻",
		Descriptor: "Elf: Light Skin Tone",
	},
	"1F9DD-1F3FB-200D-2640-FE0F": {
		Key:        "1F9DD-1F3FB-200D-2640-FE0F",
		Value:      "🧝🏻‍♀️",
		Descriptor: "Woman Elf: Light Skin Tone",
	},
	"1F9DD-1F3FB-200D-2642-FE0F": {
		Key:        "1F9DD-1F3FB-200D-2642-FE0F",
		Value:      "🧝🏻‍♂️",
		Descriptor: "Man Elf: Light Skin Tone",
	},
	"1F9DD-1F3FC": {
		Key:        "1F9DD-1F3FC",
		Value:      "🧝🏼",
		Descriptor: "Elf: Medium-Light Skin Tone",
	},
	"1F9DD-1F3FC-200D-2640-FE0F": {
		Key:        "1F9DD-1F3FC-200D-2640-FE0F",
		Value:      "🧝🏼‍♀️",
		Descriptor: "Woman Elf: Medium-Light Skin Tone",
	},
	"1F9DD-1F3FC-200D-2642-FE0F": {
		Key:        "1F9DD-1F3FC-200D-2642-FE0F",
		Value:      "🧝🏼‍♂️",
		Descriptor: "Man Elf: Medium-Light Skin Tone",
	},
	"1F9DD-1F3FD": {
		Key:        "1F9DD-1F3FD",
		Value:      "🧝🏽",
		Descriptor: "Elf: Medium Skin Tone",
	},
	"1F9DD-1F3FD-200D-2640-FE0F": {
		Key:        "1F9DD-1F3FD-200D-2640-FE0F",
		Value:      "🧝🏽‍♀️",
		Descriptor: "Woman Elf: Medium Skin Tone",
	},
	"1F9DD-1F3FD-200D-2642-FE0F": {
		Key:        "1F9DD-1F3FD-200D-2642-FE0F",
		Value:      "🧝🏽‍♂️",
		Descriptor: "Man Elf: Medium Skin Tone",
	},
	"1F9DD-1F3FE": {
		Key:        "1F9DD-1F3FE",
		Value:      "🧝🏾",
		Descriptor: "Elf: Medium-Dark Skin Tone",
	},
	"1F9DD-1F3FE-200D-2640-FE0F": {
		Key:        "1F9DD-1F3FE-200D-2640-FE0F",
		Value:      "🧝🏾‍♀️",
		Descriptor: "Woman Elf: Medium-Dark Skin Tone",
	},
	"1F9DD-1F3FE-200D-2642-FE0F": {
		Key:        "1F9DD-1F3FE-200D-2642-FE0F",
		Value:      "🧝🏾‍♂️",
		Descriptor: "Man Elf: Medium-Dark Skin Tone",
	},
	"1F9DD-1F3FF": {
		Key:        "1F9DD-1F3FF",
		Value:      "🧝🏿",
		Descriptor: "Elf: Dark Skin Tone",
	},
	"1F9DD-1F3FF-200D-2640-FE0F": {
		Key:        "1F9DD-1F3FF-200D-2640-FE0F",
		Value:      "🧝🏿‍♀️",
		Descriptor: "Woman Elf: Dark Skin Tone",
	},
	"1F9DD-1F3FF-200D-2642-FE0F": {
		Key:        "1F9DD-1F3FF-200D-2642-FE0F",
		Value:      "🧝🏿‍♂️",
		Descriptor: "Man Elf: Dark Skin Tone",
	},
	"1F9DD-200D-2640-FE0F": {
		Key:        "1F9DD-200D-2640-FE0F",
		Value:      "🧝‍♀️",
		Descriptor: "Woman Elf",
	},
	"1F9DD-200D-2642-FE0F": {
		Key:        "1F9DD-200D-2642-FE0F",
		Value:      "🧝‍♂️",
		Descriptor: "Man Elf",
	},
	"1F9DE": {
		Key:        "1F9DE",
		Value:      "🧞",
		Descriptor: "Genie",
	},
	"1F9DE-200D-2640-FE0F": {
		Key:        "1F9DE-200D-2640-FE0F",
		Value:      "🧞‍♀️",
		Descriptor: "Woman Genie",
	},
	"1F9DE-200D-2642-FE0F": {
		Key:        "1F9DE-200D-2642-FE0F",
		Value:      "🧞‍♂️",
		Descriptor: "Man Genie",
	},
	"1F9DF": {
		Key:        "1F9DF",
		Value:      "🧟",
		Descriptor: "Zombie",
	},
	"1F9DF-200D-2640-FE0F": {
		Key:        "1F9DF-200D-2640-FE0F",
		Value:      "🧟‍♀️",
		Descriptor: "Woman Zombie",
	},
	"1F9DF-200D-2642-FE0F": {
		Key:        "1F9DF-200D-2642-FE0F",
		Value:      "🧟‍♂️",
		Descriptor: "Man Zombie",
	},
	"1F9E0": {
		Key:        "1F9E0",
		Value:      "🧠",
		Descriptor: "Brain",
	},
	"1F9E1": {
		Key:        "1F9E1",
		Value:      "🧡",
		Descriptor: "Orange Heart",
	},
	"1F9E2": {
		Key:        "1F9E2",
		Value:      "🧢",
		Descriptor: "Billed Cap",
	},
	"1F9E3": {
		Key:        "1F9E3",
		Value:      "🧣",
		Descriptor: "Scarf",
	},
	"1F9E4": {
		Key:        "1F9E4",
		Value:      "🧤",
		Descriptor: "Gloves",
	},
	"1F9E5": {
		Key:        "1F9E5",
		Value:      "🧥",
		Descriptor: "Coat",
	},
	"1F9E6": {
		Key:        "1F9E6",
		Value:      "🧦",
		Descriptor: "Socks",
	},
	"1F9E7": {
		Key:        "1F9E7",
		Value:      "🧧",
		Descriptor: "Red Envelope",
	},
	"1F9E8": {
		Key:        "1F9E8",
		Value:      "🧨",
		Descriptor: "Firecracker",
	},
	"1F9E9": {
		Key:        "1F9E9",
		Value:      "🧩",
		Descriptor: "Puzzle Piece",
	},
	"1F9EA": {
		Key:        "1F9EA",
		Value:      "🧪",
		Descriptor: "Test Tube",
	},
	"1F9EB": {
		Key:        "1F9EB",
		Value:      "🧫",
		Descriptor: "Petri Dish",
	},
	"1F9EC": {
		Key:        "1F9EC",
		Value:      "🧬",
		Descriptor: "DNA",
	},
	"1F9ED": {
		Key:        "1F9ED",
		Value:      "🧭",
		Descriptor: "Compass",
	},
	"1F9EE": {
		Key:        "1F9EE",
		Value:      "🧮",
		Descriptor: "Abacus",
	},
	"1F9EF": {
		Key:        "1F9EF",
		Value:      "🧯",
		Descriptor: "Fire Extinguisher",
	},
	"1F9F0": {
		Key:        "1F9F0",
		Value:      "🧰",
		Descriptor: "Toolbox",
	},
	"1F9F1": {
		Key:        "1F9F1",
		Value:      "🧱",
		Descriptor: "Brick",
	},
	"1F9F2": {
		Key:        "1F9F2",
		Value:      "🧲",
		Descriptor: "Magnet",
	},
	"1F9F3": {
		Key:        "1F9F3",
		Value:      "🧳",
		Descriptor: "Luggage",
	},
	"1F9F4": {
		Key:        "1F9F4",
		Value:      "🧴",
		Descriptor: "Lotion Bottle",
	},
	"1F9F5": {
		Key:        "1F9F5",
		Value:      "🧵",
		Descriptor: "Thread",
	},
	"1F9F6": {
		Key:        "1F9F6",
		Value:      "🧶",
		Descriptor: "Yarn",
	},
	"1F9F7": {
		Key:        "1F9F7",
		Value:      "🧷",
		Descriptor: "Safety Pin",
	},
	"1F9F8": {
		Key:        "1F9F8",
		Value:      "🧸",
		Descriptor: "Teddy Bear",
	},
	"1F9F9": {
		Key:        "1F9F9",
		Value:      "🧹",
		Descriptor: "Broom",
	},
	"1F9FA": {
		Key:        "1F9FA",
		Value:      "🧺",
		Descriptor: "Basket",
	},
	"1F9FB": {
		Key:        "1F9FB",
		Value:      "🧻",
		Descriptor: "Roll of Paper",
	},
	"1F9FC": {
		Key:        "1F9FC",
		Value:      "🧼",
		Descriptor: "Soap",
	},
	"1F9FD": {
		Key:        "1F9FD",
		Value:      "🧽",
		Descriptor: "Sponge",
	},
	"1F9FE": {
		Key:        "1F9FE",
		Value:      "🧾",
		Descriptor: "Receipt",
	},
	"1F9FF": {
		Key:        "1F9FF",
		Value:      "🧿",
		Descriptor: "Nazar Amulet",
	},
	"1FA70": {
		Key:        "1FA70",
		Value:      "🩰",
		Descriptor: "Ballet Shoes",
	},
	"1FA71": {
		Key:        "1FA71",
		Value:      "🩱",
		Descriptor: "One-Piece Swimsuit",
	},
	"1FA72": {
		Key:        "1FA72",
		Value:      "🩲",
		Descriptor: "Briefs",
	},
	"1FA73": {
		Key:        "1FA73",
		Value:      "🩳",
		Descriptor: "Shorts",
	},
	"1FA78": {
		Key:        "1FA78",
		Value:      "🩸",
		Descriptor: "Drop of Blood",
	},
	"1FA79": {
		Key:        "1FA79",
		Value:      "🩹",
		Descriptor: "Adhesive Bandage",
	},
	"1FA7A": {
		Key:        "1FA7A",
		Value:      "🩺",
		Descriptor: "Stethoscope",
	},
	"1FA80": {
		Key:        "1FA80",
		Value:      "🪀",
		Descriptor: "Yo-Yo",
	},
	"1FA81": {
		Key:        "1FA81",
		Value:      "🪁",
		Descriptor: "Kite",
	},
	"1FA82": {
		Key:        "1FA82",
		Value:      "🪂",
		Descriptor: "Parachute",
	},
	"1FA90": {
		Key:        "1FA90",
		Value:      "🪐",
		Descriptor: "Ringed Planet",
	},
	"1FA91": {
		Key:        "1FA91",
		Value:      "🪑",
		Descriptor: "Chair",
	},
	"1FA92": {
		Key:        "1FA92",
		Value:      "🪒",
		Descriptor: "Razor",
	},
	"1FA93": {
		Key:        "1FA93",
		Value:      "🪓",
		Descriptor: "Axe",
	},
	"1FA94": {
		Key:        "1FA94",
		Value:      "🪔",
		Descriptor: "Diya Lamp",
	},
	"1FA95": {
		Key:        "1FA95",
		Value:      "🪕",
		Descriptor: "Banjo",
	},
	"203C-FE0F": {
		Key:        "203C-FE0F",
		Value:      "‼️",
		Descriptor: "Double Exclamation Mark",
	},
	"2049-FE0F": {
		Key:        "2049-FE0F",
		Value:      "⁉️",
		Descriptor: "Exclamation Question Mark",
	},
	"2122-FE0F": {
		Key:        "2122-FE0F",
		Value:      "™️",
		Descriptor: "Trade Mark",
	},
	"2139-FE0F": {
		Key:        "2139-FE0F",
		Value:      "ℹ️",
		Descriptor: "Information",
	},
	"2194-FE0F": {
		Key:        "2194-FE0F",
		Value:      "↔️",
		Descriptor: "Left-Right Arrow",
	},
	"2195-FE0F": {
		Key:        "2195-FE0F",
		Value:      "↕️",
		Descriptor: "Up-Down Arrow",
	},
	"2196-FE0F": {
		Key:        "2196-FE0F",
		Value:      "↖️",
		Descriptor: "Up-Left Arrow",
	},
	"2197-FE0F": {
		Key:        "2197-FE0F",
		Value:      "↗️",
		Descriptor: "Up-Right Arrow",
	},
	"2198-FE0F": {
		Key:        "2198-FE0F",
		Value:      "↘️",
		Descriptor: "Down-Right Arrow",
	},
	"2199-FE0F": {
		Key:        "2199-FE0F",
		Value:      "↙️",
		Descriptor: "Down-Left Arrow",
	},
	"21A9-FE0F": {
		Key:        "21A9-FE0F",
		Value:      "↩️",
		Descriptor: "Right Arrow Curving Left",
	},
	"21AA-FE0F": {
		Key:        "21AA-FE0F",
		Value:      "↪️",
		Descriptor: "Left Arrow Curving Right",
	},
	"23-FE0F-20E3": {
		Key:        "23-FE0F-20E3",
		Value:      "#️⃣",
		Descriptor: "Keycap Number Sign",
	},
	"231A": {
		Key:        "231A",
		Value:      "⌚",
		Descriptor: "Watch",
	},
	"231B": {
		Key:        "231B",
		Value:      "⌛",
		Descriptor: "Hourglass Done",
	},
	"2328-FE0F": {
		Key:        "2328-FE0F",
		Value:      "⌨️",
		Descriptor: "Keyboard",
	},
	"23CF-FE0F": {
		Key:        "23CF-FE0F",
		Value:      "⏏️",
		Descriptor: "Eject Button",
	},
	"23E9": {
		Key:        "23E9",
		Value:      "⏩",
		Descriptor: "Fast-Forward Button",
	},
	"23EA": {
		Key:        "23EA",
		Value:      "⏪",
		Descriptor: "Fast Reverse Button",
	},
	"23EB": {
		Key:        "23EB",
		Value:      "⏫",
		Descriptor: "Fast Up Button",
	},
	"23EC": {
		Key:        "23EC",
		Value:      "⏬",
		Descriptor: "Fast Down Button",
	},
	"23ED-FE0F": {
		Key:        "23ED-FE0F",
		Value:      "⏭️",
		Descriptor: "Next Track Button",
	},
	"23EE-FE0F": {
		Key:        "23EE-FE0F",
		Value:      "⏮️",
		Descriptor: "Last Track Button",
	},
	"23EF-FE0F": {
		Key:        "23EF-FE0F",
		Value:      "⏯️",
		Descriptor: "Play or Pause Button",
	},
	"23F0": {
		Key:        "23F0",
		Value:      "⏰",
		Descriptor: "Alarm Clock",
	},
	"23F1-FE0F": {
		Key:        "23F1-FE0F",
		Value:      "⏱️",
		Descriptor: "Stopwatch",
	},
	"23F2-FE0F": {
		Key:        "23F2-FE0F",
		Value:      "⏲️",
		Descriptor: "Timer Clock",
	},
	"23F3": {
		Key:        "23F3",
		Value:      "⏳",
		Descriptor: "Hourglass Not Done",
	},
	"23F8-FE0F": {
		Key:        "23F8-FE0F",
		Value:      "⏸️",
		Descriptor: "Pause Button",
	},
	"23F9-FE0F": {
		Key:        "23F9-FE0F",
		Value:      "⏹️",
		Descriptor: "Stop Button",
	},
	"23FA-FE0F": {
		Key:        "23FA-FE0F",
		Value:      "⏺️",
		Descriptor: "Record Button",
	},
	"24C2-FE0F": {
		Key:        "24C2-FE0F",
		Value:      "Ⓜ️",
		Descriptor: "Circled M",
	},
	"25AA-FE0F": {
		Key:        "25AA-FE0F",
		Value:      "▪️",
		Descriptor: "Black Small Square",
	},
	"25AB-FE0F": {
		Key:        "25AB-FE0F",
		Value:      "▫️",
		Descriptor: "White Small Square",
	},
	"25B6-FE0F": {
		Key:        "25B6-FE0F",
		Value:      "▶️",
		Descriptor: "Play Button",
	},
	"25C0-FE0F": {
		Key:        "25C0-FE0F",
		Value:      "◀️",
		Descriptor: "Reverse Button",
	},
	"25FB-FE0F": {
		Key:        "25FB-FE0F",
		Value:      "◻️",
		Descriptor: "White Medium Square",
	},
	"25FC-FE0F": {
		Key:        "25FC-FE0F",
		Value:      "◼️",
		Descriptor: "Black Medium Square",
	},
	"25FD": {
		Key:        "25FD",
		Value:      "◽",
		Descriptor: "White Medium-Small Square",
	},
	"25FE": {
		Key:        "25FE",
		Value:      "◾",
		Descriptor: "Black Medium-Small Square",
	},
	"2600-FE0F": {
		Key:        "2600-FE0F",
		Value:      "☀️",
		Descriptor: "Sun",
	},
	"2601-FE0F": {
		Key:        "2601-FE0F",
		Value:      "☁️",
		Descriptor: "Cloud",
	},
	"2602-FE0F": {
		Key:        "2602-FE0F",
		Value:      "☂️",
		Descriptor: "Umbrella",
	},
	"2603-FE0F": {
		Key:        "2603-FE0F",
		Value:      "☃️",
		Descriptor: "Snowman",
	},
	"2604-FE0F": {
		Key:        "2604-FE0F",
		Value:      "☄️",
		Descriptor: "Comet",
	},
	"260E-FE0F": {
		Key:        "260E-FE0F",
		Value:      "☎️",
		Descriptor: "Telephone",
	},
	"2611-FE0F": {
		Key:        "2611-FE0F",
		Value:      "☑️",
		Descriptor: "Check Box With Check",
	},
	"2614": {
		Key:        "2614",
		Value:      "☔",
		Descriptor: "Umbrella With Rain Drops",
	},
	"2615": {
		Key:        "2615",
		Value:      "☕",
		Descriptor: "Hot Beverage",
	},
	"2618-FE0F": {
		Key:        "2618-FE0F",
		Value:      "☘️",
		Descriptor: "Shamrock",
	},
	"261D-1F3FB": {
		Key:        "261D-1F3FB",
		Value:      "☝🏻",
		Descriptor: "Index Pointing Up: Light Skin Tone",
	},
	"261D-1F3FC": {
		Key:        "261D-1F3FC",
		Value:      "☝🏼",
		Descriptor: "Index Pointing Up: Medium-Light Skin Tone",
	},
	"261D-1F3FD": {
		Key:        "261D-1F3FD",
		Value:      "☝🏽",
		Descriptor: "Index Pointing Up: Medium Skin Tone",
	},
	"261D-1F3FE": {
		Key:        "261D-1F3FE",
		Value:      "☝🏾",
		Descriptor: "Index Pointing Up: Medium-Dark Skin Tone",
	},
	"261D-1F3FF": {
		Key:        "261D-1F3FF",
		Value:      "☝🏿",
		Descriptor: "Index Pointing Up: Dark Skin Tone",
	},
	"261D-FE0F": {
		Key:        "261D-FE0F",
		Value:      "☝️",
		Descriptor: "Index Pointing Up",
	},
	"2620-FE0F": {
		Key:        "2620-FE0F",
		Value:      "☠️",
		Descriptor: "Skull and Crossbones",
	},
	"2622-FE0F": {
		Key:        "2622-FE0F",
		Value:      "☢️",
		Descriptor: "Radioactive",
	},
	"2623-FE0F": {
		Key:        "2623-FE0F",
		Value:      "☣️",
		Descriptor: "Biohazard",
	},
	"2626-FE0F": {
		Key:        "2626-FE0F",
		Value:      "☦️",
		Descriptor: "Orthodox Cross",
	},
	"262A-FE0F": {
		Key:        "262A-FE0F",
		Value:      "☪️",
		Descriptor: "Star and Crescent",
	},
	"262E-FE0F": {
		Key:        "262E-FE0F",
		Value:      "☮️",
		Descriptor: "Peace Symbol",
	},
	"262F-FE0F": {
		Key:        "262F-FE0F",
		Value:      "☯️",
		Descriptor: "Yin Yang",
	},
	"2638-FE0F": {
		Key:        "2638-FE0F",
		Value:      "☸️",
		Descriptor: "Wheel of Dharma",
	},
	"2639-FE0F": {
		Key:        "2639-FE0F",
		Value:      "☹️",
		Descriptor: "Frowning Face",
	},
	"263A-FE0F": {
		Key:        "263A-FE0F",
		Value:      "☺️",
		Descriptor: "Smiling Face",
	},
	"2648": {
		Key:        "2648",
		Value:      "♈",
		Descriptor: "Aries",
	},
	"2649": {
		Key:        "2649",
		Value:      "♉",
		Descriptor: "Taurus",
	},
	"264A": {
		Key:        "264A",
		Value:      "♊",
		Descriptor: "Gemini",
	},
	"264B": {
		Key:        "264B",
		Value:      "♋",
		Descriptor: "Cancer",
	},
	"264C": {
		Key:        "264C",
		Value:      "♌",
		Descriptor: "Leo",
	},
	"264D": {
		Key:        "264D",
		Value:      "♍",
		Descriptor: "Virgo",
	},
	"264E": {
		Key:        "264E",
		Value:      "♎",
		Descriptor: "Libra",
	},
	"264F": {
		Key:        "264F",
		Value:      "♏",
		Descriptor: "Scorpio",
	},
	"2650": {
		Key:        "2650",
		Value:      "♐",
		Descriptor: "Sagittarius",
	},
	"2651": {
		Key:        "2651",
		Value:      "♑",
		Descriptor: "Capricorn",
	},
	"2652": {
		Key:        "2652",
		Value:      "♒",
		Descriptor: "Aquarius",
	},
	"2653": {
		Key:        "2653",
		Value:      "♓",
		Descriptor: "Pisces",
	},
	"265F-FE0F": {
		Key:        "265F-FE0F",
		Value:      "♟️",
		Descriptor: "Chess Pawn",
	},
	"2660-FE0F": {
		Key:        "2660-FE0F",
		Value:      "♠️",
		Descriptor: "Spade Suit",
	},
	"2663-FE0F": {
		Key:        "2663-FE0F",
		Value:      "♣️",
		Descriptor: "Club Suit",
	},
	"2665-FE0F": {
		Key:        "2665-FE0F",
		Value:      "♥️",
		Descriptor: "Heart Suit",
	},
	"2666-FE0F": {
		Key:        "2666-FE0F",
		Value:      "♦️",
		Descriptor: "Diamond Suit",
	},
	"2668-FE0F": {
		Key:        "2668-FE0F",
		Value:      "♨️",
		Descriptor: "Hot Springs",
	},
	"267B-FE0F": {
		Key:        "267B-FE0F",
		Value:      "♻️",
		Descriptor: "Recycling Symbol",
	},
	"267E-FE0F": {
		Key:        "267E-FE0F",
		Value:      "♾️",
		Descriptor: "Infinity",
	},
	"267F": {
		Key:        "267F",
		Value:      "♿",
		Descriptor: "Wheelchair Symbol",
	},
	"2692-FE0F": {
		Key:        "2692-FE0F",
		Value:      "⚒️",
		Descriptor: "Hammer and Pick",
	},
	"2693": {
		Key:        "2693",
		Value:      "⚓",
		Descriptor: "Anchor",
	},
	"2694-FE0F": {
		Key:        "2694-FE0F",
		Value:      "⚔️",
		Descriptor: "Crossed Swords",
	},
	"2696-FE0F": {
		Key:        "2696-FE0F",
		Value:      "⚖️",
		Descriptor: "Balance Scale",
	},
	"2697-FE0F": {
		Key:        "2697-FE0F",
		Value:      "⚗️",
		Descriptor: "Alembic",
	},
	"2699-FE0F": {
		Key:        "2699-FE0F",
		Value:      "⚙️",
		Descriptor: "Gear",
	},
	"269B-FE0F": {
		Key:        "269B-FE0F",
		Value:      "⚛️",
		Descriptor: "Atom Symbol",
	},
	"269C-FE0F": {
		Key:        "269C-FE0F",
		Value:      "⚜️",
		Descriptor: "Fleur-de-lis",
	},
	"26A0-FE0F": {
		Key:        "26A0-FE0F",
		Value:      "⚠️",
		Descriptor: "Warning",
	},
	"26A1": {
		Key:        "26A1",
		Value:      "⚡",
		Descriptor: "High Voltage",
	},
	"26AA": {
		Key:        "26AA",
		Value:      "⚪",
		Descriptor: "White Circle",
	},
	"26AB": {
		Key:        "26AB",
		Value:      "⚫",
		Descriptor: "Black Circle",
	},
	"26B0-FE0F": {
		Key:        "26B0-FE0F",
		Value:      "⚰️",
		Descriptor: "Coffin",
	},
	"26B1-FE0F": {
		Key:        "26B1-FE0F",
		Value:      "⚱️",
		Descriptor: "Funeral Urn",
	},
	"26BD": {
		Key:        "26BD",
		Value:      "⚽",
		Descriptor: "Soccer Ball",
	},
	"26BE": {
		Key:        "26BE",
		Value:      "⚾",
		Descriptor: "Baseball",
	},
	"26C4": {
		Key:        "26C4",
		Value:      "⛄",
		Descriptor: "Snowman Without Snow",
	},
	"26C5": {
		Key:        "26C5",
		Value:      "⛅",
		Descriptor: "Sun Behind Cloud",
	},
	"26C8-FE0F": {
		Key:        "26C8-FE0F",
		Value:      "⛈️",
		Descriptor: "Cloud With Lightning and Rain",
	},
	"26CE": {
		Key:        "26CE",
		Value:      "⛎",
		Descriptor: "Ophiuchus",
	},
	"26CF-FE0F": {
		Key:        "26CF-FE0F",
		Value:      "⛏️",
		Descriptor: "Pick",
	},
	"26D1-FE0F": {
		Key:        "26D1-FE0F",
		Value:      "⛑️",
		Descriptor: "Rescue Worker’s Helmet",
	},
	"26D3-FE0F": {
		Key:        "26D3-FE0F",
		Value:      "⛓️",
		Descriptor: "Chains",
	},
	"26D4": {
		Key:        "26D4",
		Value:      "⛔",
		Descriptor: "No Entry",
	},
	"26E9-FE0F": {
		Key:        "26E9-FE0F",
		Value:      "⛩️",
		Descriptor: "Shinto Shrine",
	},
	"26EA": {
		Key:        "26EA",
		Value:      "⛪",
		Descriptor: "Church",
	},
	"26F0-FE0F": {
		Key:        "26F0-FE0F",
		Value:      "⛰️",
		Descriptor: "Mountain",
	},
	"26F1-FE0F": {
		Key:        "26F1-FE0F",
		Value:      "⛱️",
		Descriptor: "Umbrella on Ground",
	},
	"26F2": {
		Key:        "26F2",
		Value:      "⛲",
		Descriptor: "Fountain",
	},
	"26F3": {
		Key:        "26F3",
		Value:      "⛳",
		Descriptor: "Flag in Hole",
	},
	"26F4-FE0F": {
		Key:        "26F4-FE0F",
		Value:      "⛴️",
		Descriptor: "Ferry",
	},
	"26F5": {
		Key:        "26F5",
		Value:      "⛵",
		Descriptor: "Sailboat",
	},
	"26F7-FE0F": {
		Key:        "26F7-FE0F",
		Value:      "⛷️",
		Descriptor: "Skier",
	},
	"26F8-FE0F": {
		Key:        "26F8-FE0F",
		Value:      "⛸️",
		Descriptor: "Ice Skate",
	},
	"26F9-1F3FB": {
		Key:        "26F9-1F3FB",
		Value:      "⛹🏻",
		Descriptor: "Person Bouncing Ball: Light Skin Tone",
	},
	"26F9-1F3FB-200D-2640-FE0F": {
		Key:        "26F9-1F3FB-200D-2640-FE0F",
		Value:      "⛹🏻‍♀️",
		Descriptor: "Woman Bouncing Ball: Light Skin Tone",
	},
	"26F9-1F3FB-200D-2642-FE0F": {
		Key:        "26F9-1F3FB-200D-2642-FE0F",
		Value:      "⛹🏻‍♂️",
		Descriptor: "Man Bouncing Ball: Light Skin Tone",
	},
	"26F9-1F3FC": {
		Key:        "26F9-1F3FC",
		Value:      "⛹🏼",
		Descriptor: "Person Bouncing Ball: Medium-Light Skin Tone",
	},
	"26F9-1F3FC-200D-2640-FE0F": {
		Key:        "26F9-1F3FC-200D-2640-FE0F",
		Value:      "⛹🏼‍♀️",
		Descriptor: "Woman Bouncing Ball: Medium-Light Skin Tone",
	},
	"26F9-1F3FC-200D-2642-FE0F": {
		Key:        "26F9-1F3FC-200D-2642-FE0F",
		Value:      "⛹🏼‍♂️",
		Descriptor: "Man Bouncing Ball: Medium-Light Skin Tone",
	},
	"26F9-1F3FD": {
		Key:        "26F9-1F3FD",
		Value:      "⛹🏽",
		Descriptor: "Person Bouncing Ball: Medium Skin Tone",
	},
	"26F9-1F3FD-200D-2640-FE0F": {
		Key:        "26F9-1F3FD-200D-2640-FE0F",
		Value:      "⛹🏽‍♀️",
		Descriptor: "Woman Bouncing Ball: Medium Skin Tone",
	},
	"26F9-1F3FD-200D-2642-FE0F": {
		Key:        "26F9-1F3FD-200D-2642-FE0F",
		Value:      "⛹🏽‍♂️",
		Descriptor: "Man Bouncing Ball: Medium Skin Tone",
	},
	"26F9-1F3FE": {
		Key:        "26F9-1F3FE",
		Value:      "⛹🏾",
		Descriptor: "Person Bouncing Ball: Medium-Dark Skin Tone",
	},
	"26F9-1F3FE-200D-2640-FE0F": {
		Key:        "26F9-1F3FE-200D-2640-FE0F",
		Value:      "⛹🏾‍♀️",
		Descriptor: "Woman Bouncing Ball: Medium-Dark Skin Tone",
	},
	"26F9-1F3FE-200D-2642-FE0F": {
		Key:        "26F9-1F3FE-200D-2642-FE0F",
		Value:      "⛹🏾‍♂️",
		Descriptor: "Man Bouncing Ball: Medium-Dark Skin Tone",
	},
	"26F9-1F3FF": {
		Key:        "26F9-1F3FF",
		Value:      "⛹🏿",
		Descriptor: "Person Bouncing Ball: Dark Skin Tone",
	},
	"26F9-1F3FF-200D-2640-FE0F": {
		Key:        "26F9-1F3FF-200D-2640-FE0F",
		Value:      "⛹🏿‍♀️",
		Descriptor: "Woman Bouncing Ball: Dark Skin Tone",
	},
	"26F9-1F3FF-200D-2642-FE0F": {
		Key:        "26F9-1F3FF-200D-2642-FE0F",
		Value:      "⛹🏿‍♂️",
		Descriptor: "Man Bouncing Ball: Dark Skin Tone",
	},
	"26F9-FE0F": {
		Key:        "26F9-FE0F",
		Value:      "⛹️",
		Descriptor: "Person Bouncing Ball",
	},
	"26F9-FE0F-200D-2640-FE0F": {
		Key:        "26F9-FE0F-200D-2640-FE0F",
		Value:      "⛹️‍♀️",
		Descriptor: "Woman Bouncing Ball",
	},
	"26F9-FE0F-200D-2642-FE0F": {
		Key:        "26F9-FE0F-200D-2642-FE0F",
		Value:      "⛹️‍♂️",
		Descriptor: "Man Bouncing Ball",
	},
	"26FA": {
		Key:        "26FA",
		Value:      "⛺",
		Descriptor: "Tent",
	},
	"26FD": {
		Key:        "26FD",
		Value:      "⛽",
		Descriptor: "Fuel Pump",
	},
	"2702-FE0F": {
		Key:        "2702-FE0F",
		Value:      "✂️",
		Descriptor: "Scissors",
	},
	"2705": {
		Key:        "2705",
		Value:      "✅",
		Descriptor: "Check Mark Button",
	},
	"2708-FE0F": {
		Key:        "2708-FE0F",
		Value:      "✈️",
		Descriptor: "Airplane",
	},
	"2709-FE0F": {
		Key:        "2709-FE0F",
		Value:      "✉️",
		Descriptor: "Envelope",
	},
	"270A": {
		Key:        "270A",
		Value:      "✊",
		Descriptor: "Raised Fist",
	},
	"270A-1F3FB": {
		Key:        "270A-1F3FB",
		Value:      "✊🏻",
		Descriptor: "Raised Fist: Light Skin Tone",
	},
	"270A-1F3FC": {
		Key:        "270A-1F3FC",
		Value:      "✊🏼",
		Descriptor: "Raised Fist: Medium-Light Skin Tone",
	},
	"270A-1F3FD": {
		Key:        "270A-1F3FD",
		Value:      "✊🏽",
		Descriptor: "Raised Fist: Medium Skin Tone",
	},
	"270A-1F3FE": {
		Key:        "270A-1F3FE",
		Value:      "✊🏾",
		Descriptor: "Raised Fist: Medium-Dark Skin Tone",
	},
	"270A-1F3FF": {
		Key:        "270A-1F3FF",
		Value:      "✊🏿",
		Descriptor: "Raised Fist: Dark Skin Tone",
	},
	"270B": {
		Key:        "270B",
		Value:      "✋",
		Descriptor: "Raised Hand",
	},
	"270B-1F3FB": {
		Key:        "270B-1F3FB",
		Value:      "✋🏻",
		Descriptor: "Raised Hand: Light Skin Tone",
	},
	"270B-1F3FC": {
		Key:        "270B-1F3FC",
		Value:      "✋🏼",
		Descriptor: "Raised Hand: Medium-Light Skin Tone",
	},
	"270B-1F3FD": {
		Key:        "270B-1F3FD",
		Value:      "✋🏽",
		Descriptor: "Raised Hand: Medium Skin Tone",
	},
	"270B-1F3FE": {
		Key:        "270B-1F3FE",
		Value:      "✋🏾",
		Descriptor: "Raised Hand: Medium-Dark Skin Tone",
	},
	"270B-1F3FF": {
		Key:        "270B-1F3FF",
		Value:      "✋🏿",
		Descriptor: "Raised Hand: Dark Skin Tone",
	},
	"270C-1F3FB": {
		Key:        "270C-1F3FB",
		Value:      "✌🏻",
		Descriptor: "Victory Hand: Light Skin Tone",
	},
	"270C-1F3FC": {
		Key:        "270C-1F3FC",
		Value:      "✌🏼",
		Descriptor: "Victory Hand: Medium-Light Skin Tone",
	},
	"270C-1F3FD": {
		Key:        "270C-1F3FD",
		Value:      "✌🏽",
		Descriptor: "Victory Hand: Medium Skin Tone",
	},
	"270C-1F3FE": {
		Key:        "270C-1F3FE",
		Value:      "✌🏾",
		Descriptor: "Victory Hand: Medium-Dark Skin Tone",
	},
	"270C-1F3FF": {
		Key:        "270C-1F3FF",
		Value:      "✌🏿",
		Descriptor: "Victory Hand: Dark Skin Tone",
	},
	"270C-FE0F": {
		Key:        "270C-FE0F",
		Value:      "✌️",
		Descriptor: "Victory Hand",
	},
	"270D-1F3FB": {
		Key:        "270D-1F3FB",
		Value:      "✍🏻",
		Descriptor: "Writing Hand: Light Skin Tone",
	},
	"270D-1F3FC": {
		Key:        "270D-1F3FC",
		Value:      "✍🏼",
		Descriptor: "Writing Hand: Medium-Light Skin Tone",
	},
	"270D-1F3FD": {
		Key:        "270D-1F3FD",
		Value:      "✍🏽",
		Descriptor: "Writing Hand: Medium Skin Tone",
	},
	"270D-1F3FE": {
		Key:        "270D-1F3FE",
		Value:      "✍🏾",
		Descriptor: "Writing Hand: Medium-Dark Skin Tone",
	},
	"270D-1F3FF": {
		Key:        "270D-1F3FF",
		Value:      "✍🏿",
		Descriptor: "Writing Hand: Dark Skin Tone",
	},
	"270D-FE0F": {
		Key:        "270D-FE0F",
		Value:      "✍️",
		Descriptor: "Writing Hand",
	},
	"270F-FE0F": {
		Key:        "270F-FE0F",
		Value:      "✏️",
		Descriptor: "Pencil",
	},
	"2712-FE0F": {
		Key:        "2712-FE0F",
		Value:      "✒️",
		Descriptor: "Black Nib",
	},
	"2714-FE0F": {
		Key:        "2714-FE0F",
		Value:      "✔️",
		Descriptor: "Check Mark",
	},
	"2716-FE0F": {
		Key:        "2716-FE0F",
		Value:      "✖️",
		Descriptor: "Multiplication Sign",
	},
	"271D-FE0F": {
		Key:        "271D-FE0F",
		Value:      "✝️",
		Descriptor: "Latin Cross",
	},
	"2721-FE0F": {
		Key:        "2721-FE0F",
		Value:      "✡️",
		Descriptor: "Star of David",
	},
	"2728": {
		Key:        "2728",
		Value:      "✨",
		Descriptor: "Sparkles",
	},
	"2733-FE0F": {
		Key:        "2733-FE0F",
		Value:      "✳️",
		Descriptor: "Eight-Spoked Asterisk",
	},
	"2734-FE0F": {
		Key:        "2734-FE0F",
		Value:      "✴️",
		Descriptor: "Eight-Pointed Star",
	},
	"2744-FE0F": {
		Key:        "2744-FE0F",
		Value:      "❄️",
		Descriptor: "Snowflake",
	},
	"2747-FE0F": {
		Key:        "2747-FE0F",
		Value:      "❇️",
		Descriptor: "Sparkle",
	},
	"274C": {
		Key:        "274C",
		Value:      "❌",
		Descriptor: "Cross Mark",
	},
	"274E": {
		Key:        "274E",
		Value:      "❎",
		Descriptor: "Cross Mark Button",
	},
	"2753": {
		Key:        "2753",
		Value:      "❓",
		Descriptor: "Question Mark",
	},
	"2754": {
		Key:        "2754",
		Value:      "❔",
		Descriptor: "White Question Mark",
	},
	"2755": {
		Key:        "2755",
		Value:      "❕",
		Descriptor: "White Exclamation Mark",
	},
	"2757": {
		Key:        "2757",
		Value:      "❗",
		Descriptor: "Exclamation Mark",
	},
	"2763-FE0F": {
		Key:        "2763-FE0F",
		Value:      "❣️",
		Descriptor: "Heart Exclamation",
	},
	"2764-FE0F": {
		Key:        "2764-FE0F",
		Value:      "❤️",
		Descriptor: "Red Heart",
	},
	"2795": {
		Key:        "2795",
		Value:      "➕",
		Descriptor: "Plus Sign",
	},
	"2796": {
		Key:        "2796",
		Value:      "➖",
		Descriptor: "Minus Sign",
	},
	"2797": {
		Key:        "2797",
		Value:      "➗",
		Descriptor: "Division Sign",
	},
	"27A1-FE0F": {
		Key:        "27A1-FE0F",
		Value:      "➡️",
		Descriptor: "Right Arrow",
	},
	"27B0": {
		Key:        "27B0",
		Value:      "➰",
		Descriptor: "Curly Loop",
	},
	"27BF": {
		Key:        "27BF",
		Value:      "➿",
		Descriptor: "Double Curly Loop",
	},
	"2934-FE0F": {
		Key:        "2934-FE0F",
		Value:      "⤴️",
		Descriptor: "Right Arrow Curving Up",
	},
	"2935-FE0F": {
		Key:        "2935-FE0F",
		Value:      "⤵️",
		Descriptor: "Right Arrow Curving Down",
	},
	"2A-FE0F-20E3": {
		Key:        "2A-FE0F-20E3",
		Value:      "*️⃣",
		Descriptor: "Keycap Asterisk",
	},
	"2B05-FE0F": {
		Key:        "2B05-FE0F",
		Value:      "⬅️",
		Descriptor: "Left Arrow",
	},
	"2B06-FE0F": {
		Key:        "2B06-FE0F",
		Value:      "⬆️",
		Descriptor: "Up Arrow",
	},
	"2B07-FE0F": {
		Key:        "2B07-FE0F",
		Value:      "⬇️",
		Descriptor: "Down Arrow",
	},
	"2B1B": {
		Key:        "2B1B",
		Value:      "⬛",
		Descriptor: "Black Large Square",
	},
	"2B1C": {
		Key:        "2B1C",
		Value:      "⬜",
		Descriptor: "White Large Square",
	},
	"2B50": {
		Key:        "2B50",
		Value:      "⭐",
		Descriptor: "Star",
	},
	"2B55": {
		Key:        "2B55",
		Value:      "⭕",
		Descriptor: "Hollow Red Circle",
	},
	"30-FE0F-20E3": {
		Key:        "30-FE0F-20E3",
		Value:      "0️⃣",
		Descriptor: "Keycap Digit Zero",
	},
	"3030-FE0F": {
		Key:        "3030-FE0F",
		Value:      "〰️",
		Descriptor: "Wavy Dash",
	},
	"303D-FE0F": {
		Key:        "303D-FE0F",
		Value:      "〽️",
		Descriptor: "Part Alternation Mark",
	},
	"31-FE0F-20E3": {
		Key:        "31-FE0F-20E3",
		Value:      "1️⃣",
		Descriptor: "Keycap Digit One",
	},
	"32-FE0F-20E3": {
		Key:        "32-FE0F-20E3",
		Value:      "2️⃣",
		Descriptor: "Keycap Digit Two",
	},
	"3297-FE0F": {
		Key:        "3297-FE0F",
		Value:      "㊗️",
		Descriptor: "Japanese “Congratulations” Button",
	},
	"3299-FE0F": {
		Key:        "3299-FE0F",
		Value:      "㊙️",
		Descriptor: "Japanese “Secret” Button",
	},
	"33-FE0F-20E3": {
		Key:        "33-FE0F-20E3",
		Value:      "3️⃣",
		Descriptor: "Keycap Digit Three",
	},
	"34-FE0F-20E3": {
		Key:        "34-FE0F-20E3",
		Value:      "4️⃣",
		Descriptor: "Keycap Digit Four",
	},
	"35-FE0F-20E3": {
		Key:        "35-FE0F-20E3",
		Value:      "5️⃣",
		Descriptor: "Keycap Digit Five",
	},
	"36-FE0F-20E3": {
		Key:        "36-FE0F-20E3",
		Value:      "6️⃣",
		Descriptor: "Keycap Digit Six",
	},
	"37-FE0F-20E3": {
		Key:        "37-FE0F-20E3",
		Value:      "7️⃣",
		Descriptor: "Keycap Digit Seven",
	},
	"38-FE0F-20E3": {
		Key:        "38-FE0F-20E3",
		Value:      "8️⃣",
		Descriptor: "Keycap Digit Eight",
	},
	"39-FE0F-20E3": {
		Key:        "39-FE0F-20E3",
		Value:      "9️⃣",
		Descriptor: "Keycap Digit Nine",
	},
	"A9-FE0F": {
		Key:        "A9-FE0F",
		Value:      "©️",
		Descriptor: "Copyright",
	},
	"AE-FE0F": {
		Key:        "AE-FE0F",
		Value:      "®️",
		Descriptor: "Registered",
	},
}
