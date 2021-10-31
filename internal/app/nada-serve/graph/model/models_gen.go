// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type Node interface {
	IsNode()
}

// Represents an airport
type Airport struct {
	// IATA airport code
	ID string `json:"id"`
	// Airport literal name
	Name *string `json:"name"`
	// Airport elevation in feet
	ElevationFt *int `json:"elevationFt"`
	// Airport continent
	Continent *string `json:"continent"`
	// Airport continent
	IsoCountry *string `json:"isoCountry"`
	// Airport iso region
	IsoRegion *string `json:"isoRegion"`
	// Airport municipality
	Municipality *string `json:"municipality"`
	// Airport gps code
	GpsCode *string `json:"gpsCode"`
	// Airport local code
	LocalCode *string `json:"localCode"`
	// Airport coordinate lat (Decimal Degree) in WGS84 coordinate system
	Lat *float64 `json:"lat"`
	// Airport coordinate lat (Decimal Degree) in WGS84 coordinate system
	Lon *float64 `json:"lon"`
	// List of available sensors for this airport
	Sensors []*Sensor `json:"sensors"`
	// Get a subset of sensors
	GetSubsetOfSensors []*Sensor `json:"getSubsetOfSensors"`
}

func (Airport) IsNode() {}

type MeasureMeanData struct {
	// Measure id
	ID string `json:"id"`
	// Sensor used to capture this measure
	Sensor *Sensor `json:"sensor"`
	// Measure airport location
	Airport *Airport `json:"airport"`
	// Measure value
	Value float64 `json:"value"`
	// Start date of mean calculation
	StartDate time.Time `json:"startDate"`
	// End date of mean calculation
	EndDate time.Time `json:"endDate"`
}

func (MeasureMeanData) IsNode() {}

type Measurement struct {
	// Measurement id
	ID string `json:"id"`
	// Measurement name
	Name string `json:"name"`
	// Measurement unit
	Unit string `json:"unit"`
}

func (Measurement) IsNode() {}

type Sensor struct {
	// Sensor id
	ID string `json:"id"`
	// Sensor's airport
	Airport *Airport `json:"airport"`
	// Type of measurement that provides this sensor
	Measurement *Measurement `json:"measurement"`
	// Get a serie of mean measures for this sensor
	GetMeanMeasureInterval []*MeasureMeanData `json:"getMeanMeasureInterval"`
	// Get mean measures for a specific day, for this sensor.
	GetMeanMeasures []*MeasureMeanData `json:"getMeanMeasures"`
}

func (Sensor) IsNode() {}

type FakeColor struct {
	Red255   *int `json:"red255"`
	Green255 *int `json:"green255"`
	Blue255  *int `json:"blue255"`
}

type FakeImageSize struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type FakeOptions struct {
	UseFullAddress    *bool          `json:"useFullAddress"`
	MinMoney          *float64       `json:"minMoney"`
	MaxMoney          *float64       `json:"maxMoney"`
	DecimalPlaces     *int           `json:"decimalPlaces"`
	ImageSize         *FakeImageSize `json:"imageSize"`
	ImageKeywords     []string       `json:"imageKeywords"`
	RandomizeImageURL *bool          `json:"randomizeImageUrl"`
	EmailProvider     *string        `json:"emailProvider"`
	PasswordLength    *int           `json:"passwordLength"`
	LoremSize         *FakeLoremSize `json:"loremSize"`
	DateFormat        *string        `json:"dateFormat"`
	DateFrom          *string        `json:"dateFrom"`
	DateTo            *string        `json:"dateTo"`
	BaseColor         *FakeColor     `json:"baseColor"`
	MinNumber         *float64       `json:"minNumber"`
	MaxNumber         *float64       `json:"maxNumber"`
	PrecisionNumber   *float64       `json:"precisionNumber"`
}

type MeanMeasureMode string

const (
	// The value specified by discretize will divide in X MeasureMeanData specified by the duration
	MeanMeasureModeFluxDuration MeanMeasureMode = "FLUX_DURATION"
	// The value specified by discretize will divide in X MeasureMeanData for the whole interval
	MeanMeasureModeForInterval MeanMeasureMode = "FOR_INTERVAL"
)

var AllMeanMeasureMode = []MeanMeasureMode{
	MeanMeasureModeFluxDuration,
	MeanMeasureModeForInterval,
}

func (e MeanMeasureMode) IsValid() bool {
	switch e {
	case MeanMeasureModeFluxDuration, MeanMeasureModeForInterval:
		return true
	}
	return false
}

func (e MeanMeasureMode) String() string {
	return string(e)
}

func (e *MeanMeasureMode) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = MeanMeasureMode(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid MeanMeasureMode", str)
	}
	return nil
}

func (e MeanMeasureMode) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type FakeLocale string

const (
	FakeLocaleAz        FakeLocale = "az"
	FakeLocaleCz        FakeLocale = "cz"
	FakeLocaleDe        FakeLocale = "de"
	FakeLocaleDeAt      FakeLocale = "de_AT"
	FakeLocaleDeCh      FakeLocale = "de_CH"
	FakeLocaleEn        FakeLocale = "en"
	FakeLocaleEnAu      FakeLocale = "en_AU"
	FakeLocaleEnBork    FakeLocale = "en_BORK"
	FakeLocaleEnCa      FakeLocale = "en_CA"
	FakeLocaleEnGb      FakeLocale = "en_GB"
	FakeLocaleEnIe      FakeLocale = "en_IE"
	FakeLocaleEnInd     FakeLocale = "en_IND"
	FakeLocaleEnUs      FakeLocale = "en_US"
	FakeLocaleEnAuOcker FakeLocale = "en_au_ocker"
	FakeLocaleEs        FakeLocale = "es"
	FakeLocaleEsMx      FakeLocale = "es_MX"
	FakeLocaleFa        FakeLocale = "fa"
	FakeLocaleFr        FakeLocale = "fr"
	FakeLocaleFrCa      FakeLocale = "fr_CA"
	FakeLocaleGe        FakeLocale = "ge"
	FakeLocaleIDID      FakeLocale = "id_ID"
	FakeLocaleIt        FakeLocale = "it"
	FakeLocaleJa        FakeLocale = "ja"
	FakeLocaleKo        FakeLocale = "ko"
	FakeLocaleNbNo      FakeLocale = "nb_NO"
	FakeLocaleNep       FakeLocale = "nep"
	FakeLocaleNl        FakeLocale = "nl"
	FakeLocalePl        FakeLocale = "pl"
	FakeLocalePtBr      FakeLocale = "pt_BR"
	FakeLocaleRu        FakeLocale = "ru"
	FakeLocaleSk        FakeLocale = "sk"
	FakeLocaleSv        FakeLocale = "sv"
	FakeLocaleTr        FakeLocale = "tr"
	FakeLocaleUk        FakeLocale = "uk"
	FakeLocaleVi        FakeLocale = "vi"
	FakeLocaleZhCn      FakeLocale = "zh_CN"
	FakeLocaleZhTw      FakeLocale = "zh_TW"
)

var AllFakeLocale = []FakeLocale{
	FakeLocaleAz,
	FakeLocaleCz,
	FakeLocaleDe,
	FakeLocaleDeAt,
	FakeLocaleDeCh,
	FakeLocaleEn,
	FakeLocaleEnAu,
	FakeLocaleEnBork,
	FakeLocaleEnCa,
	FakeLocaleEnGb,
	FakeLocaleEnIe,
	FakeLocaleEnInd,
	FakeLocaleEnUs,
	FakeLocaleEnAuOcker,
	FakeLocaleEs,
	FakeLocaleEsMx,
	FakeLocaleFa,
	FakeLocaleFr,
	FakeLocaleFrCa,
	FakeLocaleGe,
	FakeLocaleIDID,
	FakeLocaleIt,
	FakeLocaleJa,
	FakeLocaleKo,
	FakeLocaleNbNo,
	FakeLocaleNep,
	FakeLocaleNl,
	FakeLocalePl,
	FakeLocalePtBr,
	FakeLocaleRu,
	FakeLocaleSk,
	FakeLocaleSv,
	FakeLocaleTr,
	FakeLocaleUk,
	FakeLocaleVi,
	FakeLocaleZhCn,
	FakeLocaleZhTw,
}

func (e FakeLocale) IsValid() bool {
	switch e {
	case FakeLocaleAz, FakeLocaleCz, FakeLocaleDe, FakeLocaleDeAt, FakeLocaleDeCh, FakeLocaleEn, FakeLocaleEnAu, FakeLocaleEnBork, FakeLocaleEnCa, FakeLocaleEnGb, FakeLocaleEnIe, FakeLocaleEnInd, FakeLocaleEnUs, FakeLocaleEnAuOcker, FakeLocaleEs, FakeLocaleEsMx, FakeLocaleFa, FakeLocaleFr, FakeLocaleFrCa, FakeLocaleGe, FakeLocaleIDID, FakeLocaleIt, FakeLocaleJa, FakeLocaleKo, FakeLocaleNbNo, FakeLocaleNep, FakeLocaleNl, FakeLocalePl, FakeLocalePtBr, FakeLocaleRu, FakeLocaleSk, FakeLocaleSv, FakeLocaleTr, FakeLocaleUk, FakeLocaleVi, FakeLocaleZhCn, FakeLocaleZhTw:
		return true
	}
	return false
}

func (e FakeLocale) String() string {
	return string(e)
}

func (e *FakeLocale) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = FakeLocale(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid fake__Locale", str)
	}
	return nil
}

func (e FakeLocale) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type FakeTypes string

const (
	FakeTypesZipCode            FakeTypes = "zipCode"
	FakeTypesCity               FakeTypes = "city"
	FakeTypesStreetName         FakeTypes = "streetName"
	FakeTypesStreetAddress      FakeTypes = "streetAddress"
	FakeTypesSecondaryAddress   FakeTypes = "secondaryAddress"
	FakeTypesCounty             FakeTypes = "county"
	FakeTypesCountry            FakeTypes = "country"
	FakeTypesCountryCode        FakeTypes = "countryCode"
	FakeTypesState              FakeTypes = "state"
	FakeTypesStateAbbr          FakeTypes = "stateAbbr"
	FakeTypesLatitude           FakeTypes = "latitude"
	FakeTypesLongitude          FakeTypes = "longitude"
	FakeTypesColorName          FakeTypes = "colorName"
	FakeTypesProductCategory    FakeTypes = "productCategory"
	FakeTypesProductName        FakeTypes = "productName"
	FakeTypesMoney              FakeTypes = "money"
	FakeTypesProductMaterial    FakeTypes = "productMaterial"
	FakeTypesProduct            FakeTypes = "product"
	FakeTypesCompanyName        FakeTypes = "companyName"
	FakeTypesCompanyCatchPhrase FakeTypes = "companyCatchPhrase"
	FakeTypesCompanyBs          FakeTypes = "companyBS"
	FakeTypesDbColumn           FakeTypes = "dbColumn"
	FakeTypesDbType             FakeTypes = "dbType"
	FakeTypesDbCollation        FakeTypes = "dbCollation"
	FakeTypesDbEngine           FakeTypes = "dbEngine"
	// By default returns dates beetween 2000-01-01 and 2030-01-01.
	// Configure date format with options \`dateFormat\` \`dateFrom\` \`dateTo\`.
	FakeTypesDate                           FakeTypes = "date"
	FakeTypesPastDate                       FakeTypes = "pastDate"
	FakeTypesFutureDate                     FakeTypes = "futureDate"
	FakeTypesRecentDate                     FakeTypes = "recentDate"
	FakeTypesFinanceAccountName             FakeTypes = "financeAccountName"
	FakeTypesFinanceTransactionType         FakeTypes = "financeTransactionType"
	FakeTypesCurrencyCode                   FakeTypes = "currencyCode"
	FakeTypesCurrencyName                   FakeTypes = "currencyName"
	FakeTypesCurrencySymbol                 FakeTypes = "currencySymbol"
	FakeTypesBitcoinAddress                 FakeTypes = "bitcoinAddress"
	FakeTypesInternationalBankAccountNumber FakeTypes = "internationalBankAccountNumber"
	FakeTypesBankIdentifierCode             FakeTypes = "bankIdentifierCode"
	FakeTypesHackerAbbreviation             FakeTypes = "hackerAbbreviation"
	FakeTypesHackerPhrase                   FakeTypes = "hackerPhrase"
	FakeTypesImageURL                       FakeTypes = "imageUrl"
	FakeTypesAvatarURL                      FakeTypes = "avatarUrl"
	FakeTypesEmail                          FakeTypes = "email"
	FakeTypesURL                            FakeTypes = "url"
	FakeTypesDomainName                     FakeTypes = "domainName"
	FakeTypesIpv4Address                    FakeTypes = "ipv4Address"
	FakeTypesIpv6Address                    FakeTypes = "ipv6Address"
	FakeTypesUserAgent                      FakeTypes = "userAgent"
	FakeTypesColorHex                       FakeTypes = "colorHex"
	FakeTypesMacAddress                     FakeTypes = "macAddress"
	FakeTypesPassword                       FakeTypes = "password"
	FakeTypesLorem                          FakeTypes = "lorem"
	FakeTypesFirstName                      FakeTypes = "firstName"
	FakeTypesLastName                       FakeTypes = "lastName"
	FakeTypesFullName                       FakeTypes = "fullName"
	FakeTypesJobTitle                       FakeTypes = "jobTitle"
	FakeTypesPhoneNumber                    FakeTypes = "phoneNumber"
	FakeTypesNumber                         FakeTypes = "number"
	FakeTypesUUID                           FakeTypes = "uuid"
	FakeTypesWord                           FakeTypes = "word"
	FakeTypesWords                          FakeTypes = "words"
	FakeTypesLocale                         FakeTypes = "locale"
	FakeTypesFilename                       FakeTypes = "filename"
	FakeTypesMimeType                       FakeTypes = "mimeType"
	FakeTypesFileExtension                  FakeTypes = "fileExtension"
	FakeTypesSemver                         FakeTypes = "semver"
)

var AllFakeTypes = []FakeTypes{
	FakeTypesZipCode,
	FakeTypesCity,
	FakeTypesStreetName,
	FakeTypesStreetAddress,
	FakeTypesSecondaryAddress,
	FakeTypesCounty,
	FakeTypesCountry,
	FakeTypesCountryCode,
	FakeTypesState,
	FakeTypesStateAbbr,
	FakeTypesLatitude,
	FakeTypesLongitude,
	FakeTypesColorName,
	FakeTypesProductCategory,
	FakeTypesProductName,
	FakeTypesMoney,
	FakeTypesProductMaterial,
	FakeTypesProduct,
	FakeTypesCompanyName,
	FakeTypesCompanyCatchPhrase,
	FakeTypesCompanyBs,
	FakeTypesDbColumn,
	FakeTypesDbType,
	FakeTypesDbCollation,
	FakeTypesDbEngine,
	FakeTypesDate,
	FakeTypesPastDate,
	FakeTypesFutureDate,
	FakeTypesRecentDate,
	FakeTypesFinanceAccountName,
	FakeTypesFinanceTransactionType,
	FakeTypesCurrencyCode,
	FakeTypesCurrencyName,
	FakeTypesCurrencySymbol,
	FakeTypesBitcoinAddress,
	FakeTypesInternationalBankAccountNumber,
	FakeTypesBankIdentifierCode,
	FakeTypesHackerAbbreviation,
	FakeTypesHackerPhrase,
	FakeTypesImageURL,
	FakeTypesAvatarURL,
	FakeTypesEmail,
	FakeTypesURL,
	FakeTypesDomainName,
	FakeTypesIpv4Address,
	FakeTypesIpv6Address,
	FakeTypesUserAgent,
	FakeTypesColorHex,
	FakeTypesMacAddress,
	FakeTypesPassword,
	FakeTypesLorem,
	FakeTypesFirstName,
	FakeTypesLastName,
	FakeTypesFullName,
	FakeTypesJobTitle,
	FakeTypesPhoneNumber,
	FakeTypesNumber,
	FakeTypesUUID,
	FakeTypesWord,
	FakeTypesWords,
	FakeTypesLocale,
	FakeTypesFilename,
	FakeTypesMimeType,
	FakeTypesFileExtension,
	FakeTypesSemver,
}

func (e FakeTypes) IsValid() bool {
	switch e {
	case FakeTypesZipCode, FakeTypesCity, FakeTypesStreetName, FakeTypesStreetAddress, FakeTypesSecondaryAddress, FakeTypesCounty, FakeTypesCountry, FakeTypesCountryCode, FakeTypesState, FakeTypesStateAbbr, FakeTypesLatitude, FakeTypesLongitude, FakeTypesColorName, FakeTypesProductCategory, FakeTypesProductName, FakeTypesMoney, FakeTypesProductMaterial, FakeTypesProduct, FakeTypesCompanyName, FakeTypesCompanyCatchPhrase, FakeTypesCompanyBs, FakeTypesDbColumn, FakeTypesDbType, FakeTypesDbCollation, FakeTypesDbEngine, FakeTypesDate, FakeTypesPastDate, FakeTypesFutureDate, FakeTypesRecentDate, FakeTypesFinanceAccountName, FakeTypesFinanceTransactionType, FakeTypesCurrencyCode, FakeTypesCurrencyName, FakeTypesCurrencySymbol, FakeTypesBitcoinAddress, FakeTypesInternationalBankAccountNumber, FakeTypesBankIdentifierCode, FakeTypesHackerAbbreviation, FakeTypesHackerPhrase, FakeTypesImageURL, FakeTypesAvatarURL, FakeTypesEmail, FakeTypesURL, FakeTypesDomainName, FakeTypesIpv4Address, FakeTypesIpv6Address, FakeTypesUserAgent, FakeTypesColorHex, FakeTypesMacAddress, FakeTypesPassword, FakeTypesLorem, FakeTypesFirstName, FakeTypesLastName, FakeTypesFullName, FakeTypesJobTitle, FakeTypesPhoneNumber, FakeTypesNumber, FakeTypesUUID, FakeTypesWord, FakeTypesWords, FakeTypesLocale, FakeTypesFilename, FakeTypesMimeType, FakeTypesFileExtension, FakeTypesSemver:
		return true
	}
	return false
}

func (e FakeTypes) String() string {
	return string(e)
}

func (e *FakeTypes) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = FakeTypes(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid fake__Types", str)
	}
	return nil
}

func (e FakeTypes) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type FakeLoremSize string

const (
	FakeLoremSizeWord       FakeLoremSize = "word"
	FakeLoremSizeWords      FakeLoremSize = "words"
	FakeLoremSizeSentence   FakeLoremSize = "sentence"
	FakeLoremSizeSentences  FakeLoremSize = "sentences"
	FakeLoremSizeParagraph  FakeLoremSize = "paragraph"
	FakeLoremSizeParagraphs FakeLoremSize = "paragraphs"
)

var AllFakeLoremSize = []FakeLoremSize{
	FakeLoremSizeWord,
	FakeLoremSizeWords,
	FakeLoremSizeSentence,
	FakeLoremSizeSentences,
	FakeLoremSizeParagraph,
	FakeLoremSizeParagraphs,
}

func (e FakeLoremSize) IsValid() bool {
	switch e {
	case FakeLoremSizeWord, FakeLoremSizeWords, FakeLoremSizeSentence, FakeLoremSizeSentences, FakeLoremSizeParagraph, FakeLoremSizeParagraphs:
		return true
	}
	return false
}

func (e FakeLoremSize) String() string {
	return string(e)
}

func (e *FakeLoremSize) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = FakeLoremSize(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid fake__loremSize", str)
	}
	return nil
}

func (e FakeLoremSize) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
