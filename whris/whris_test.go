package whris

import (
	"testing"
)

const TestRipeWhoisResponse = `% IANA WHOIS server
% for more information on IANA, visit http://www.iana.org
% This query returned 1 object

refer:        whois.ripe.net

inetnum:      93.0.0.0 - 93.255.255.255
organisation: RIPE NCC
status:       ALLOCATED

whois:        whois.ripe.net

changed:      2007-03
source:       IANA

% This is the RIPE Database query service.
% The objects are in RPSL format.
%
% The RIPE Database is subject to Terms and Conditions.
% See http://www.ripe.net/db/support/db-terms-conditions.pdf

% Note: this output has been filtered.
%       To receive output for a database update, use the "-B" flag.

% Information related to '93.184.216.0 - 93.184.216.255'

% Abuse contact for '93.184.216.0 - 93.184.216.255' is 'abuse@verizondigitalmedia.com'

inetnum:        93.184.216.0 - 93.184.216.255
netname:        EDGECAST-NETBLK-03
descr:          NETBLK-03-EU-93-184-216-0-24
country:        EU
admin-c:        DS7892-RIPE
tech-c:         DS7892-RIPE
status:         ASSIGNED PA
mnt-by:         MNT-EDGECAST
created:        2012-06-22T21:48:41Z
last-modified:  2012-06-22T21:48:41Z
source:         RIPE # Filtered

person:         Derrick Sawyer
address:        13031 W Jefferson Blvd #900, Los Angeles, CA 90094
phone:          +18773343236
nic-hdl:        DS7892-RIPE
created:        2010-08-25T18:44:19Z
last-modified:  2017-03-03T09:06:18Z
source:         RIPE
mnt-by:         MNT-EDGECAST

% This query was served by the RIPE Database Query Service version 1.102.2 (WAGYU)

;; Query time: 1194 msec
;; WHEN: Tue Feb 01 21:12:31 JST 2022`

func TestResolve(t *testing.T) {
	type args struct {
		domain  string
		verbose bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Resolve(tt.args.domain, tt.args.verbose); (err != nil) != tt.wantErr {
				t.Errorf("Resolve() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSummary_ParseWhoisResponse(t *testing.T) {
	type fields struct {
		TargetDomain        string
		TargetIp            string
		WhoisResponseServer string
		WhoisResponse       string
		ParseResult         []NetworkAdmin
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Summary{
				TargetDomain:        tt.fields.TargetDomain,
				TargetIp:            tt.fields.TargetIp,
				WhoisResponseServer: tt.fields.WhoisResponseServer,
				WhoisResponse:       tt.fields.WhoisResponse,
				ParseResult:         tt.fields.ParseResult,
			}
			if err := s.ParseWhoisResponse(); (err != nil) != tt.wantErr {
				t.Errorf("Summary.ParseWhoisResponse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSummary_SetWhoisResponseServerFromWhoisResponse(t *testing.T) {
	type fields struct {
		WhoisResponseServer string
		WhoisResponse       string
	}
	tests := []struct {
		name   string
		fields fields
		want   fields
	}{
		{
			name: "Be able to correctly split the response from ripe",
			fields: fields{
				WhoisResponseServer: "whois.apnic.net",
				WhoisResponse:       TestRipeWhoisResponse,
			},
			want: fields{
				WhoisResponseServer: "whois.ripe.net",
				WhoisResponse:       TestRipeWhoisResponse,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Summary{
				WhoisResponseServer: tt.fields.WhoisResponseServer,
				WhoisResponse:       tt.fields.WhoisResponse,
			}
			s.SetWhoisResponseServerFromWhoisResponse()
			if s.WhoisResponseServer != tt.want.WhoisResponseServer {
				t.Errorf("Summary.SetWhoisResponseServerFromWhoisResponse() = %v, want %v", s.WhoisResponseServer, tt.want.WhoisResponseServer)
			}
		})
	}
}

func TestSummary_ParseCheck(t *testing.T) {
	type fields struct {
		TargetDomain        string
		TargetIp            string
		WhoisResponseServer string
		WhoisResponse       string
		ParseResult         []NetworkAdmin
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Summary{
				TargetDomain:        tt.fields.TargetDomain,
				TargetIp:            tt.fields.TargetIp,
				WhoisResponseServer: tt.fields.WhoisResponseServer,
				WhoisResponse:       tt.fields.WhoisResponse,
				ParseResult:         tt.fields.ParseResult,
			}
			if got := s.ParseCheck(); got != tt.want {
				t.Errorf("Summary.ParseCheck() = %v, want %v", got, tt.want)
			}
		})
	}
}
