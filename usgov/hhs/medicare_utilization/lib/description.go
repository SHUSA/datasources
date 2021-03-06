package lib

import (
  "io"
  "regexp"
  "encoding/csv"
  "github.com/bloomapi/dataloading"
  "github.com/bloomapi/dataloading/helpers"
)

type Description struct {}

func (d *Description) Available() ([]dataloading.Source, error) {
  return []dataloading.Source{
    dataloading.Source{
      Name: "usgov.hhs.medicare_utilization",
      Version: "2013-00",
    },
  }, nil
}

func (d *Description) FieldNames(sourceName string) ([]string, error) {
  fileMatch := regexp.MustCompile(`Medicare_Provider_Util_Payment_PUF_CY2013.txt$`)

  reader, err := getFileReader("http://download.cms.gov/Research-Statistics-Data-and-Systems/Statistics-Trends-and-Reports/Medicare-Provider-Charge-Data/Downloads/Medicare_Provider_Util_Payment_PUF_CY2013.zip?agree=yes&next=Accept", fileMatch)
  if err != nil {
    return nil, err
  }

  csvReader := csv.NewReader(reader)
  if err != nil {
    return nil, err
  }

  csvReader.Comma = '\t'

  columns, err := csvReader.Read()
  if err != nil {
    return nil, err
  }

  return columns, nil
}

func getFileReader(uri string, zipPattern *regexp.Regexp) (io.Reader, error) {
  downloader := dataloading.NewDownloader("data/", nil)
  path, err := downloader.Fetch(uri)
  if err != nil {
    return nil, err
  }

  reader, err := helpers.OpenExtractZipReader(path, zipPattern)
  if err != nil {
    return nil, err
  }

  return reader, nil
}

func (d *Description) Reader(source dataloading.Source) (dataloading.ValueReader, error) {
  fileMatch := regexp.MustCompile(`Medicare_Provider_Util_Payment_PUF_CY2013.txt$`)
  reader, err := getFileReader("http://download.cms.gov/Research-Statistics-Data-and-Systems/Statistics-Trends-and-Reports/Medicare-Provider-Charge-Data/Downloads/Medicare_Provider_Util_Payment_PUF_CY2013.zip?agree=yes&next=Accept", fileMatch)
  if err != nil {
    return nil, err
  }

  zipReader := helpers.NewCsvTabReader(reader)

  return zipReader, nil
}