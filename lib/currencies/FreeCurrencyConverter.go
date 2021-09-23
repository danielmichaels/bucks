package currencies

type Currency struct {
	Value float64 `json:"value,string"`
	From  string  `json:"fromCurrency"`
	To    string  `json:"toCurrency"`
	Total float64 `json:"total,string"`
}

type FreeCurrencyURL struct {
	Url string
}

type CurrencyList struct {
	Id             string
	CurrencyName   string
	CurrencySymbol string
}

type CurrencyConverterAPI struct {
	Key     string `mapstructure:"key"`
	Default bool   `mapstructure:"default"`
}

var FreeCurrencyConverterList = []CurrencyList{
	{
		CurrencyName:   "Albanian Lek",
		CurrencySymbol: "Lek",
		Id:             "ALL",
	},
	{
		CurrencyName:   "East Caribbean Dollar",
		CurrencySymbol: "$",
		Id:             "XCD",
	},
	{
		CurrencyName:   "Euro",
		CurrencySymbol: "€",
		Id:             "EUR",
	},
	{
		CurrencyName:   "Barbadian Dollar",
		CurrencySymbol: "$",
		Id:             "BBD",
	},
	{
		CurrencyName: "Bhutanese Ngultrum",
		Id:           "BTN",
	},
	{
		CurrencyName:   "Brunei Dollar",
		CurrencySymbol: "$",
		Id:             "BND",
	},
	{
		CurrencyName: "Central African CFA Franc",
		Id:           "XAF",
	},
	{
		CurrencyName:   "Cuban Peso",
		CurrencySymbol: "$",
		Id:             "CUP",
	},
	{
		CurrencyName:   "United States Dollar",
		CurrencySymbol: "$",
		Id:             "USD",
	},
	{
		CurrencyName:   "Falkland Islands Pound",
		CurrencySymbol: "£",
		Id:             "FKP",
	},
	{
		CurrencyName:   "Gibraltar Pound",
		CurrencySymbol: "£",
		Id:             "GIP",
	},
	{
		CurrencyName:   "Hungarian Forint",
		CurrencySymbol: "Ft",
		Id:             "HUF",
	},
	{
		CurrencyName:   "Iranian Rial",
		CurrencySymbol: "﷼",
		Id:             "IRR",
	},
	{
		CurrencyName:   "Jamaican Dollar",
		CurrencySymbol: "J$",
		Id:             "JMD",
	},
	{
		CurrencyName:   "Australian Dollar",
		CurrencySymbol: "$",
		Id:             "AUD",
	},
	{
		CurrencyName:   "Lao Kip",
		CurrencySymbol: "₭",
		Id:             "LAK",
	},
	{
		CurrencyName: "Libyan Dinar",
		Id:           "LYD",
	},
	{
		CurrencyName:   "Macedonian Denar",
		CurrencySymbol: "ден",
		Id:             "MKD",
	},
	{
		CurrencyName: "West African CFA Franc",
		Id:           "XOF",
	},
	{
		CurrencyName:   "New Zealand Dollar",
		CurrencySymbol: "$",
		Id:             "NZD",
	},
	{
		CurrencyName:   "Omani Rial",
		CurrencySymbol: "﷼",
		Id:             "OMR",
	},
	{
		CurrencyName: "Papua New Guinean Kina",
		Id:           "PGK",
	},
	{
		CurrencyName: "Rwandan Franc",
		Id:           "RWF",
	},
	{
		CurrencyName: "Samoan Tala",
		Id:           "WST",
	},
	{
		CurrencyName:   "Serbian Dinar",
		CurrencySymbol: "Дин.",
		Id:             "RSD",
	},
	{
		CurrencyName:   "Swedish Krona",
		CurrencySymbol: "kr",
		Id:             "SEK",
	},
	{
		CurrencyName:   "Tanzanian Shilling",
		CurrencySymbol: "TSh",
		Id:             "TZS",
	},
	{
		CurrencyName: "Armenian Dram",
		Id:           "AMD",
	},
	{
		CurrencyName:   "Bahamian Dollar",
		CurrencySymbol: "$",
		Id:             "BSD",
	},
	{
		CurrencyName:   "Bosnia And Herzegovina Konvertibilna Marka",
		CurrencySymbol: "KM",
		Id:             "BAM",
	},
	{
		CurrencyName: "Cape Verdean Escudo",
		Id:           "CVE",
	},
	{
		CurrencyName:   "Chinese Yuan",
		CurrencySymbol: "¥",
		Id:             "CNY",
	},
	{
		CurrencyName:   "Costa Rican Colon",
		CurrencySymbol: "₡",
		Id:             "CRC",
	},
	{
		CurrencyName:   "Czech Koruna",
		CurrencySymbol: "Kč",
		Id:             "CZK",
	},
	{
		CurrencyName: "Eritrean Nakfa",
		Id:           "ERN",
	},
	{
		CurrencyName: "Georgian Lari",
		Id:           "GEL",
	},
	{
		CurrencyName: "Haitian Gourde",
		Id:           "HTG",
	},
	{
		CurrencyName:   "Indian Rupee",
		CurrencySymbol: "₹",
		Id:             "INR",
	},
	{
		CurrencyName: "Jordanian Dinar",
		Id:           "JOD",
	},
	{
		CurrencyName:   "South Korean Won",
		CurrencySymbol: "₩",
		Id:             "KRW",
	},
	{
		CurrencyName:   "Lebanese Lira",
		CurrencySymbol: "£",
		Id:             "LBP",
	},
	{
		CurrencyName: "Malawian Kwacha",
		Id:           "MWK",
	},
	{
		CurrencyName: "Mauritanian Ouguiya",
		Id:           "MRO",
	},
	{
		CurrencyName: "Mozambican Metical",
		Id:           "MZN",
	},
	{
		CurrencyName:   "Netherlands Antillean Gulden",
		CurrencySymbol: "ƒ",
		Id:             "ANG",
	},
	{
		CurrencyName:   "Peruvian Nuevo Sol",
		CurrencySymbol: "S/.",
		Id:             "PEN",
	},
	{
		CurrencyName:   "Qatari Riyal",
		CurrencySymbol: "﷼",
		Id:             "QAR",
	},
	{
		CurrencyName: "Sao Tome And Principe Dobra",
		Id:           "STD",
	},
	{
		CurrencyName: "Sierra Leonean Leone",
		Id:           "SLL",
	},
	{
		CurrencyName:   "Somali Shilling",
		CurrencySymbol: "S",
		Id:             "SOS",
	},
	{
		CurrencyName: "Sudanese Pound",
		Id:           "SDG",
	},
	{
		CurrencyName:   "Syrian Pound",
		CurrencySymbol: "£",
		Id:             "SYP",
	},
	{
		CurrencyName: "Angolan Kwanza",
		Id:           "AOA",
	},
	{
		CurrencyName:   "Aruban Florin",
		CurrencySymbol: "ƒ",
		Id:             "AWG",
	},
	{
		CurrencyName: "Bahraini Dinar",
		Id:           "BHD",
	},
	{
		CurrencyName:   "Belize Dollar",
		CurrencySymbol: "BZ$",
		Id:             "BZD",
	},
	{
		CurrencyName:   "Botswana Pula",
		CurrencySymbol: "P",
		Id:             "BWP",
	},
	{
		CurrencyName: "Burundi Franc",
		Id:           "BIF",
	},
	{
		CurrencyName:   "Cayman Islands Dollar",
		CurrencySymbol: "$",
		Id:             "KYD",
	},
	{
		CurrencyName:   "Colombian Peso",
		CurrencySymbol: "$",
		Id:             "COP",
	},
	{
		CurrencyName:   "Danish Krone",
		CurrencySymbol: "kr",
		Id:             "DKK",
	},
	{
		CurrencyName:   "Guatemalan Quetzal",
		CurrencySymbol: "Q",
		Id:             "GTQ",
	},
	{
		CurrencyName:   "Honduran Lempira",
		CurrencySymbol: "L",
		Id:             "HNL",
	},
	{
		CurrencyName:   "Indonesian Rupiah",
		CurrencySymbol: "Rp",
		Id:             "IDR",
	},
	{
		CurrencyName:   "Israeli New Sheqel",
		CurrencySymbol: "₪",
		Id:             "ILS",
	},
	{
		CurrencyName:   "Kazakhstani Tenge",
		CurrencySymbol: "лв",
		Id:             "KZT",
	},
	{
		CurrencyName: "Kuwaiti Dinar",
		Id:           "KWD",
	},
	{
		CurrencyName: "Lesotho Loti",
		Id:           "LSL",
	},
	{
		CurrencyName:   "Malaysian Ringgit",
		CurrencySymbol: "RM",
		Id:             "MYR",
	},
	{
		CurrencyName:   "Mauritian Rupee",
		CurrencySymbol: "₨",
		Id:             "MUR",
	},
	{
		CurrencyName:   "Mongolian Tugrik",
		CurrencySymbol: "₮",
		Id:             "MNT",
	},
	{
		CurrencyName: "Myanma Kyat",
		Id:           "MMK",
	},
	{
		CurrencyName:   "Nigerian Naira",
		CurrencySymbol: "₦",
		Id:             "NGN",
	},
	{
		CurrencyName:   "Panamanian Balboa",
		CurrencySymbol: "B/.",
		Id:             "PAB",
	},
	{
		CurrencyName:   "Philippine Peso",
		CurrencySymbol: "₱",
		Id:             "PHP",
	},
	{
		CurrencyName:   "Romanian Leu",
		CurrencySymbol: "lei",
		Id:             "RON",
	},
	{
		CurrencyName:   "Saudi Riyal",
		CurrencySymbol: "﷼",
		Id:             "SAR",
	},
	{
		CurrencyName:   "Singapore Dollar",
		CurrencySymbol: "$",
		Id:             "SGD",
	},
	{
		CurrencyName:   "South African Rand",
		CurrencySymbol: "R",
		Id:             "ZAR",
	},
	{
		CurrencyName:   "Surinamese Dollar",
		CurrencySymbol: "$",
		Id:             "SRD",
	},
	{
		CurrencyName:   "New Taiwan Dollar",
		CurrencySymbol: "NT$",
		Id:             "TWD",
	},
	{
		CurrencyName: "Paanga",
		Id:           "TOP",
	},
	{
		CurrencyName: "Venezuelan Bolivar",
		Id:           "VEF",
	},
	{
		CurrencyName: "Algerian Dinar",
		Id:           "DZD",
	},
	{
		CurrencyName:   "Argentine Peso",
		CurrencySymbol: "$",
		Id:             "ARS",
	},
	{
		CurrencyName:   "Azerbaijani Manat",
		CurrencySymbol: "ман",
		Id:             "AZN",
	},
	{
		CurrencyName:   "Belarusian Ruble",
		CurrencySymbol: "p.",
		Id:             "BYR",
	},
	{
		CurrencyName:   "Bolivian Boliviano",
		CurrencySymbol: "$b",
		Id:             "BOB",
	},
	{
		CurrencyName:   "Bulgarian Lev",
		CurrencySymbol: "лв",
		Id:             "BGN",
	},
	{
		CurrencyName:   "Canadian Dollar",
		CurrencySymbol: "$",
		Id:             "CAD",
	},
	{
		CurrencyName:   "Chilean Peso",
		CurrencySymbol: "$",
		Id:             "CLP",
	},
	{
		CurrencyName: "Congolese Franc",
		Id:           "CDF",
	},
	{
		CurrencyName:   "Dominican Peso",
		CurrencySymbol: "RD$",
		Id:             "DOP",
	},
	{
		CurrencyName:   "Fijian Dollar",
		CurrencySymbol: "$",
		Id:             "FJD",
	},
	{
		CurrencyName: "Gambian Dalasi",
		Id:           "GMD",
	},
	{
		CurrencyName:   "Guyanese Dollar",
		CurrencySymbol: "$",
		Id:             "GYD",
	},
	{
		CurrencyName:   "Icelandic Króna",
		CurrencySymbol: "kr",
		Id:             "ISK",
	},
	{
		CurrencyName: "Iraqi Dinar",
		Id:           "IQD",
	},
	{
		CurrencyName:   "Japanese Yen",
		CurrencySymbol: "¥",
		Id:             "JPY",
	},
	{
		CurrencyName:   "North Korean Won",
		CurrencySymbol: "₩",
		Id:             "KPW",
	},
	{
		CurrencyName:   "Latvian Lats",
		CurrencySymbol: "Ls",
		Id:             "LVL",
	},
	{
		CurrencyName:   "Swiss Franc",
		CurrencySymbol: "Fr.",
		Id:             "CHF",
	},
	{
		CurrencyName: "Malagasy Ariary",
		Id:           "MGA",
	},
	{
		CurrencyName: "Moldovan Leu",
		Id:           "MDL",
	},
	{
		CurrencyName: "Moroccan Dirham",
		Id:           "MAD",
	},
	{
		CurrencyName:   "Nepalese Rupee",
		CurrencySymbol: "₨",
		Id:             "NPR",
	},
	{
		CurrencyName:   "Nicaraguan Cordoba",
		CurrencySymbol: "C$",
		Id:             "NIO",
	},
	{
		CurrencyName:   "Pakistani Rupee",
		CurrencySymbol: "₨",
		Id:             "PKR",
	},
	{
		CurrencyName:   "Paraguayan Guarani",
		CurrencySymbol: "Gs",
		Id:             "PYG",
	},
	{
		CurrencyName:   "Saint Helena Pound",
		CurrencySymbol: "£",
		Id:             "SHP",
	},
	{
		CurrencyName:   "Seychellois Rupee",
		CurrencySymbol: "₨",
		Id:             "SCR",
	},
	{
		CurrencyName:   "Solomon Islands Dollar",
		CurrencySymbol: "$",
		Id:             "SBD",
	},
	{
		CurrencyName:   "Sri Lankan Rupee",
		CurrencySymbol: "₨",
		Id:             "LKR",
	},
	{
		CurrencyName:   "Thai Baht",
		CurrencySymbol: "฿",
		Id:             "THB",
	},
	{
		CurrencyName: "Turkish New Lira",
		Id:           "TRY",
	},
	{
		CurrencyName: "UAE Dirham",
		Id:           "AED",
	},
	{
		CurrencyName: "Vanuatu Vatu",
		Id:           "VUV",
	},
	{
		CurrencyName:   "Yemeni Rial",
		CurrencySymbol: "﷼",
		Id:             "YER",
	},
	{
		CurrencyName:   "Afghan Afghani",
		CurrencySymbol: "؋",
		Id:             "AFN",
	},
	{
		CurrencyName: "Bangladeshi Taka",
		Id:           "BDT",
	},
	{
		CurrencyName:   "Brazilian Real",
		CurrencySymbol: "R$",
		Id:             "BRL",
	},
	{
		CurrencyName:   "Cambodian Riel",
		CurrencySymbol: "៛",
		Id:             "KHR",
	},
	{
		CurrencyName: "Comorian Franc",
		Id:           "KMF",
	},
	{
		CurrencyName:   "Croatian Kuna",
		CurrencySymbol: "kn",
		Id:             "HRK",
	},
	{
		CurrencyName: "Djiboutian Franc",
		Id:           "DJF",
	},
	{
		CurrencyName:   "Egyptian Pound",
		CurrencySymbol: "£",
		Id:             "EGP",
	},
	{
		CurrencyName: "Ethiopian Birr",
		Id:           "ETB",
	},
	{
		CurrencyName: "CFP Franc",
		Id:           "XPF",
	},
	{
		CurrencyName: "Ghanaian Cedi",
		Id:           "GHS",
	},
	{
		CurrencyName: "Guinean Franc",
		Id:           "GNF",
	},
	{
		CurrencyName:   "Hong Kong Dollar",
		CurrencySymbol: "$",
		Id:             "HKD",
	},
	{
		CurrencyName: "Special Drawing Rights",
		Id:           "XDR",
	},
	{
		CurrencyName:   "Kenyan Shilling",
		CurrencySymbol: "KSh",
		Id:             "KES",
	},
	{
		CurrencyName:   "Kyrgyzstani Som",
		CurrencySymbol: "лв",
		Id:             "KGS",
	},
	{
		CurrencyName:   "Liberian Dollar",
		CurrencySymbol: "$",
		Id:             "LRD",
	},
	{
		CurrencyName: "Macanese Pataca",
		Id:           "MOP",
	},
	{
		CurrencyName: "Maldivian Rufiyaa",
		Id:           "MVR",
	},
	{
		CurrencyName:   "Mexican Peso",
		CurrencySymbol: "$",
		Id:             "MXN",
	},
	{
		CurrencyName:   "Namibian Dollar",
		CurrencySymbol: "$",
		Id:             "NAD",
	},
	{
		CurrencyName:   "Norwegian Krone",
		CurrencySymbol: "kr",
		Id:             "NOK",
	},
	{
		CurrencyName:   "Polish Zloty",
		CurrencySymbol: "zł",
		Id:             "PLN",
	},
	{
		CurrencyName:   "Russian Ruble",
		CurrencySymbol: "руб",
		Id:             "RUB",
	},
	{
		CurrencyName: "Swazi Lilangeni",
		Id:           "SZL",
	},
	{
		CurrencyName: "Tajikistani Somoni",
		Id:           "TJS",
	},
	{
		CurrencyName:   "Trinidad and Tobago Dollar",
		CurrencySymbol: "TT$",
		Id:             "TTD",
	},
	{
		CurrencyName:   "Ugandan Shilling",
		CurrencySymbol: "USh",
		Id:             "UGX",
	},
	{
		CurrencyName:   "Uruguayan Peso",
		CurrencySymbol: "$U",
		Id:             "UYU",
	},
	{
		CurrencyName:   "Vietnamese Dong",
		CurrencySymbol: "₫",
		Id:             "VND",
	},
	{
		CurrencyName: "Tunisian Dinar",
		Id:           "TND",
	},
	{
		CurrencyName:   "Ukrainian Hryvnia",
		CurrencySymbol: "₴",
		Id:             "UAH",
	},
	{
		CurrencyName:   "Uzbekistani Som",
		CurrencySymbol: "лв",
		Id:             "UZS",
	},
	{
		CurrencyName: "Turkmenistan Manat",
		Id:           "TMT",
	},
	{
		CurrencyName:   "British Pound",
		CurrencySymbol: "£",
		Id:             "GBP",
	},
	{
		CurrencyName: "Zambian Kwacha",
		Id:           "ZMW",
	},
	{
		CurrencyName:   "Bitcoin",
		CurrencySymbol: "BTC",
		Id:             "BTC",
	},
	{
		CurrencyName:   "New Belarusian Ruble",
		CurrencySymbol: "p.",
		Id:             "BYN",
	},
	{
		CurrencyName: "Bermudan Dollar",
		Id:           "BMD",
	},
	{
		CurrencyName: "Guernsey Pound",
		Id:           "GGP",
	},
	{
		CurrencyName: "Chilean Unit Of Account",
		Id:           "CLF",
	},
	{
		CurrencyName: "Cuban Convertible Peso",
		Id:           "CUC",
	},
	{
		CurrencyName: "Manx pound",
		Id:           "IMP",
	},
	{
		CurrencyName: "Jersey Pound",
		Id:           "JEP",
	},
	{
		CurrencyName: "Salvadoran Colón",
		Id:           "SVC",
	},
	{
		CurrencyName: "Old Zambian Kwacha",
		Id:           "ZMK",
	},
	{
		CurrencyName: "Silver (troy ounce)",
		Id:           "XAG",
	},
	{
		CurrencyName: "Zimbabwean Dollar",
		Id:           "ZWL",
	},
}
