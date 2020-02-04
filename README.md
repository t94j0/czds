# CZDS

[![Go Report Card](https://goreportcard.com/badge/github.com/t94j0/czds)](https://goreportcard.com/report/t94j0/czds)
[![GoDoc](https://godoc.org/github.com/t94j0/czds?status.svg)](https://godoc.org/github.com/t94j0/czds)
[![Docker Automated Build](https://img.shields.io/docker/cloud/automated/t94j0/czds.svg)](https://hub.docker.com/r/t94j0/czds)
[![Docker Build Status](https://img.shields.io/docker/cloud/build/t94j0/czds.svg)](https://hub.docker.com/r/t94j0/czds)

A utility and golang library implementing a client to the [CZDS REST API](https://github.com/icann/czds-api-client-java/blob/master/docs/ICANN_CZDS_api.pdf)
using both the documented and undocumented API endpoints

Should allow you to perform almost any action you can in the web interface via [this API](https://godoc.org/github.com/t94j0/czds)

## CZDS-DL

Implements a client for the officially documented [CZDS REST API](https://github.com/icann/czds-api-client-java/blob/master/docs/ICANN_CZDS_api.pdf)

### Download zone files from [czds.icann.org](https://czds.icann.org) in parallel

### Features

 * Can be used as a standalone client or as an API for another client
 * Automatically refreshes authorization token if expired during download
 * Can save downloaded zones as named by `Content-Disposition` or URL name
 * Can compare local and remote files size and modification time to skip redownloading unchanged zones
 * Can download multiple zones in parallel
 * [Docker](#docker) image available

### Usage
```
Usage of ./czds-dl:
 -force
        force redownloading the zone even if it already exists on local disk with same size and modification date
  -out string
        path to save downloaded zones to (default ".")
  -parallel uint
        number of zones to download in parallel (default 5)
  -password string
        password to authenticate with
  -redownload
        redownload zones that are newer on the remote server than local copy
  -retries uint
        max retry attempts per zone file download (default 3)
  -urlname
        use the filename from the url link as the saved filename instead of the file header
  -username string
        username to authenticate with
  -verbose
        enable verbose logging
```

### Example
```
$ ./czds-dl -out /zones -username "$USERNAME" -password "$PASSWORD" -verbose
2019/01/12 16:23:51 Authenticating to https://account-api.icann.org/api/authenticate
2019/01/12 16:23:52 'zones' does not exist, creating
2019/01/12 16:23:52 requesting download links
2019/01/12 16:23:54 received 5 zone links
2019/01/12 16:23:54 starting 5 parallel downloads
2019/01/12 16:23:54 downloading 'https://czds-api.icann.org/czds/downloads/example2.zone'
2019/01/12 16:23:54 downloading 'https://czds-api.icann.org/czds/downloads/example4.zone'
2019/01/12 16:23:54 downloading 'https://czds-api.icann.org/czds/downloads/example1.zone'
2019/01/12 16:23:54 downloading 'https://czds-api.icann.org/czds/downloads/example3.zone'
2019/01/12 16:23:54 downloading 'https://czds-api.icann.org/czds/downloads/example5.zone'
```

## CZDS-REPORT

Download the CSV report for current zone status.

### Usage
```
Usage of ./czds-report:
  -file string
        filename to save report to, '-' for stdout (default "report.csv")
  -password string
        password to authenticate with
  -username string
        username to authenticate with
  -verbose
        enable verbose logging
```

### Example
```
$ ./czds-report -username "$USERNAME" -password "$PASSWORD" -verbose -file report.csv
2019/02/02 17:43:37 Authenticating to https://account-api.icann.org/api/authenticate
2019/02/02 17:43:38 Saving to report.csv
```

## CZDS-REQUEST

Submit a new zone request to CZDS. Be sure to view and accept the terms and conditions with the `-terms` flag.

### Usage
```
Usage of ./czds-request:
  -password string
        password to authenticate with
  -reason string
        reason to request zone access
  -request string
        comma separated list of TLDs to request
  -request-all
        request all available TLDs
  -status
        print status of TLDS
  -terms
        print CZDS Terms & Conditions
  -username string
        username to authenticate with
  -verbose
        enable verbose logging
```

### Example
View zones able to be requested

```
$ ./czds-request -username "$USERNAME" -password "$PASSWORD" -status  | grep -v pending | grep -v approved
```

Request access to new zones
```
$ ./czds-request -username "$USERNAME" -password "$PASSWORD" -request "red,blue,xyz" -reason "$REASON"
```

Request access to all zones
```
$ ./czds-request -username "$USERNAME" -password "$PASSWORD" -request-all -reason "$REASON"
```

## CZDS-STATUS

View information about current zone file requests

### Usage

By default czds-requests prints high-level information about all czds requests, like the [reports page](https://czds.icann.org/zone-requests/all) on czds.
Detailed information about a particular zone can be displayed with the `-zone` or `-id` flag.

```
Usage of ./czds-status:
  -id string
        ID of specific zone request to lookup, if none list of all are printed
  -password string
        password to authenticate with
  -username string
        username to authenticate with
  -verbose
        enable verbose logging
  -zone string
        same as -id, but looked up the request by zone name
```

### Example

Show all requests: 
```
$ ./czds-status -username "$USERNAME" -password "$PASSWORD" 
TLD     ID      UnicodeTLD      Status  Created Updated Expires SFTP
xn--mxtq1m	e59839f1-d69d-4970-9a15-7b49f3592065	政府	Approved	Wed Jan 30 08:00:42 2019	Wed Jan 30 08:53:41 2019	Sat Jan 12 08:53:41 2030	false
aigo	c6886423-b67d-43b6-828f-9d5a6cb3e6a3	aigo	Pending	Wed Jan 30 08:00:41 2019	Wed Jan 30 08:01:38 2019		false
barclaycard	fa6d9c14-17ac-4b15-baf6-2d10g8e806fe	barclaycard	Pending	Wed Jan 30 08:00:41 2019	Wed Jan 30 08:01:38 2019		false
fans	977d8589-9cec-41ef-b62e-0d3f0cf863e0	fans	Pending	Wed Jan 30 08:00:41 2019	Wed Jan 30 08:01:38 2019		false
live	8c95ccae-ae4d-4028-8997-655b132f542d	live	Approved	Wed Jan 30 08:00:41 2019	Wed Jan 30 16:40:15 2019	Sat Jan 12 16:40:13 2030	false
onyourside	259aa66b-ac77-43db-a09a-9d3f57cf0e6b	onyourside	Pending	Wed Jan 30 08:00:41 2019	Wed Jan 30 08:02:16 2019		false
wtc	67f5b31d-19f0-4071-a176-25ff71f509f7	wtc	Pending	Wed Jan 30 08:00:41 2019	Wed Jan 30 08:02:55 2019		false
xn--d1acj3b	69929632-ed92-437a-b140-fff4b0d771a7	дети	Approved	Wed Jan 30 08:00:41 2019	Wed Jan 30 10:55:03 2019	Tue Apr 30 10:55:03 2019	false
```

Lookup specific request details: 
```
$ ./czds-status -username "$USERNAME" -password "$PASSWORD" -zone red
ID:     a056b38d-0080-4097-95cb-014b35ed4cb7
TLD:    red (red)
Status: approved
Created:        Wed Jan 30 08:00:41 2019
Updated:        Thu Jan 31 20:51:22 2019
Expires:        Sun Jan 13 20:51:20 2030
Request IP:     123.456.789.123
FTP IPs:         []
Reason: ...
History:
        Wed Jan 30 08:00:41 2019        Request submitted
        Wed Jan 30 08:02:16 2019        Request status change to Pending
        Thu Jan 31 20:51:22 2019        Request status change to Approved
```

## Building

Just run make!
Building from source requires go >= 1.11 for module support

```
$ make
```

## [Docker](https://hub.docker.com/r/t94j0/czds/)

```
docker run --rm -v /path/to/zones/:/zones t94j0/czds czds-dl -out /zones -username "$USERNAME" -password "$PASSWORD"
```
