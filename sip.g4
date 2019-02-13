alphanum  :  alpha | digit;
reserved    :  SEMICOLON | SLASH | QUESTION | COLON | AT | AMPERSAND | EQUALS | PLUS
                     | DOLLAR | COMMA;
unreserved  :  alphanum | mark;
mark        :  DASH | UNDERSCORE | PERIOD | EXCLAMATION | TILDE | ASTERISK | APOSTROPHE
                     | LEFT_PAREN | RIGHT_PAREN;
escaped     :  PERCENT hexdig hexdig;

lws  :  (wsp* crlf)? wsp+; // linear whitespace
sws  :  (lws)?; // sep whitespace

hcolon  :  ( sp | htab )* COLON sws;

text_utf8_trim  :  text_utf8char+ (lws* text_utf8char)*;
text_utf8char   :  (EXCLAMATION | QUOTE | POUND | DOLLAR | PERCENT | AMPERSAND | APOSTROPHE | LEFT_PAREN | RIGHT_PAREN | ASTERISK | PLUS | COMMA | DASH | PERIOD | SLASH | ZERO | ONE | TWO | THREE | FOUR | FIVE | SIX | SEVEN | EIGHT | NINE | COLON | SEMICOLON | LESS_THAN | EQUALS | GREATER_THAN | QUESTION | AT | CAP_A | CAP_B | CAP_C | CAP_D | CAP_E | CAP_F | CAP_G | CAP_H | CAP_I | CAP_J | CAP_K | CAP_L | CAP_M | CAP_N | CAP_O | CAP_P | CAP_Q | CAP_R | CAP_S | CAP_T | CAP_U | CAP_V | CAP_W | CAP_X | CAP_Y | CAP_Z | LEFT_BRACE | BACKSLASH | RIGHT_BRACE | CARAT | UNDERSCORE | ACCENT | A | B | C | D | E | F | G | H | I | J | K | L | M | N | O | P | Q | R | S | T | U | V | W | X | Y | Z | LEFT_CURLY_BRACE | PIPE | RIGHT_CURLY_BRACE | TILDE) | utf8_nonascii;
utf8_nonascii   :  ((U_00C0 | U_00C1 | U_00C2 | U_00C3 | U_00C4 | U_00C5 | U_00C6 | U_00C7 | U_00C8 | U_00C9 | U_00CA | U_00CB | U_00CC | U_00CD | U_00CE | U_00CF | U_00D0 | U_00D1 | U_00D2 | U_00D3 | U_00D4 | U_00D5 | U_00D6 | U_00D7 | U_00D8 | U_00D9 | U_00DA | U_00DB | U_00DC | U_00DD | U_00DE | U_00DF) utf8_cont)
                      |  ((U_00E0 | U_00E1 | U_00E2 | U_00E3 | U_00E4 | U_00E5 | U_00E6 | U_00E7 | U_00E8 | U_00E9 | U_00EA | U_00EB | U_00EC | U_00ED | U_00EE | U_00EF) utf8_cont utf8_cont)
                      |  ((U_00F0 | U_00F1 | U_00F2 | U_00F3 | U_00F4 | U_00F5 | U_00F6 | U_00F7) utf8_cont utf8_cont utf8_cont)
                      |  ((U_00F8 | U_00F9 | U_00FA | U_00FB) utf8_cont utf8_cont utf8_cont utf8_cont)
                      |  ((U_00FC | U_00FD) utf8_cont utf8_cont utf8_cont utf8_cont utf8_cont);
utf8_cont       :  (U_0080 | U_0081 | U_0082 | U_0083 | U_0084 | U_0085 | U_0086 | U_0087 | U_0088 | U_0089 | U_008A | U_008B | U_008C | U_008D | U_008E | U_008F | U_0090 | U_0091 | U_0092 | U_0093 | U_0094 | U_0095 | U_0096 | U_0097 | U_0098 | U_0099 | U_009A | U_009B | U_009C | U_009D | U_009E | U_009F | U_00A0 | U_00A1 | U_00A2 | U_00A3 | U_00A4 | U_00A5 | U_00A6 | U_00A7 | U_00A8 | U_00A9 | U_00AA | U_00AB | U_00AC | U_00AD | U_00AE | U_00AF | U_00B0 | U_00B1 | U_00B2 | U_00B3 | U_00B4 | U_00B5 | U_00B6 | U_00B7 | U_00B8 | U_00B9 | U_00BA | U_00BB | U_00BC | U_00BD | U_00BE | U_00BF);
lhex  :  digit | (A | B | C | D | E | F); //lowercase a-f

token       :  (alphanum | DASH | PERIOD | EXCLAMATION | PERCENT | ASTERISK
                     | UNDERSCORE | PLUS | ACCENT | APOSTROPHE | TILDE )+;

word        :  (alphanum | DASH | PERIOD | EXCLAMATION | PERCENT | ASTERISK |
                     UNDERSCORE | PLUS | ACCENT | APOSTROPHE | TILDE |
                     LEFT_PAREN | RIGHT_PAREN | LESS_THAN | GREATER_THAN |
                     COLON | BACKSLASH | dquote |
                     SLASH | LEFT_BRACE | RIGHT_BRACE | QUESTION |
                     LEFT_CURLY_BRACE | RIGHT_CURLY_BRACE )+;

star    :  sws ASTERISK sws; // asterisk
slash   :  sws SLASH sws; // slash
equal   :  sws EQUALS sws; // equal
lparen  :  sws LEFT_PAREN sws; // left parenthesis
rparen  :  sws RIGHT_PAREN sws; // right parenthesis
raquot  :  GREATER_THAN sws; // right angle quote
laquot  :  sws LESS_THAN;// left angle quote
comma   :  sws COMMA sws; // comma
semi    :  sws SEMICOLON sws; // semicolon
colon   :  sws COLON sws; // colon
ldquot  :  sws dquote;// open double quotation mark
rdquot  :  dquote sws; // close double quotation mark
comment  :  lparen (ctext | quoted_pair | comment)* rparen;
ctext    :  (EXCLAMATION | QUOTE | POUND | DOLLAR | PERCENT | AMPERSAND | APOSTROPHE) | (ASTERISK | PLUS | COMMA | DASH | PERIOD | SLASH | ZERO | ONE | TWO | THREE | FOUR | FIVE | SIX | SEVEN | EIGHT | NINE | COLON | SEMICOLON | LESS_THAN | EQUALS | GREATER_THAN | QUESTION | AT | CAP_A | CAP_B | CAP_C | CAP_D | CAP_E | CAP_F | CAP_G | CAP_H | CAP_I | CAP_J | CAP_K | CAP_L | CAP_M | CAP_N | CAP_O | CAP_P | CAP_Q | CAP_R | CAP_S | CAP_T | CAP_U | CAP_V | CAP_W | CAP_X | CAP_Y | CAP_Z | LEFT_BRACE) | (RIGHT_BRACE | CARAT | UNDERSCORE | ACCENT | A | B | C | D | E | F | G | H | I | J | K | L | M | N | O | P | Q | R | S | T | U | V | W | X | Y | Z | LEFT_CURLY_BRACE | PIPE | RIGHT_CURLY_BRACE | TILDE) | utf8_nonascii
                  | lws;

quoted_string  :  sws dquote (qdtext | quoted_pair )* dquote;
qdtext         :  lws | EXCLAMATION | (POUND | DOLLAR | PERCENT | AMPERSAND | APOSTROPHE | LEFT_PAREN | RIGHT_PAREN | ASTERISK | PLUS | COMMA | DASH | PERIOD | SLASH | ZERO | ONE | TWO | THREE | FOUR | FIVE | SIX | SEVEN | EIGHT | NINE | COLON | SEMICOLON | LESS_THAN | EQUALS | GREATER_THAN | QUESTION | AT | CAP_A | CAP_B | CAP_C | CAP_D | CAP_E | CAP_F | CAP_G | CAP_H | CAP_I | CAP_J | CAP_K | CAP_L | CAP_M | CAP_N | CAP_O | CAP_P | CAP_Q | CAP_R | CAP_S | CAP_T | CAP_U | CAP_V | CAP_W | CAP_X | CAP_Y | CAP_Z | LEFT_BRACE) | (RIGHT_BRACE | CARAT | UNDERSCORE | ACCENT | A | B | C | D | E | F | G | H | I | J | K | L | M | N | O | P | Q | R | S | T | U | V | W | X | Y | Z | LEFT_CURLY_BRACE | PIPE | RIGHT_CURLY_BRACE | TILDE)
                        | utf8_nonascii;

quoted_pair  :  BACKSLASH ((U_0000 | U_0001 | U_0002 | U_0003 | U_0004 | U_0005 | U_0006 | U_0007 | U_0008 | TAB) | (U_000B | U_000C)
                | (U_000E | U_000F | U_0010 | U_0011 | U_0012 | U_0013 | U_0014 | U_0015 | U_0016 | U_0017 | U_0018 | U_0019 | U_001A | U_001B | U_001C | U_001D | U_001E | U_001F | SPACE | EXCLAMATION | QUOTE | POUND | DOLLAR | PERCENT | AMPERSAND | APOSTROPHE | LEFT_PAREN | RIGHT_PAREN | ASTERISK | PLUS | COMMA | DASH | PERIOD | SLASH | ZERO | ONE | TWO | THREE | FOUR | FIVE | SIX | SEVEN | EIGHT | NINE | COLON | SEMICOLON | LESS_THAN | EQUALS | GREATER_THAN | QUESTION | AT | CAP_A | CAP_B | CAP_C | CAP_D | CAP_E | CAP_F | CAP_G | CAP_H | CAP_I | CAP_J | CAP_K | CAP_L | CAP_M | CAP_N | CAP_O | CAP_P | CAP_Q | CAP_R | CAP_S | CAP_T | CAP_U | CAP_V | CAP_W | CAP_X | CAP_Y | CAP_Z | LEFT_BRACE | BACKSLASH | RIGHT_BRACE | CARAT | UNDERSCORE | ACCENT | A | B | C | D | E | F | G | H | I | J | K | L | M | N | O | P | Q | R | S | T | U | V | W | X | Y | Z | LEFT_CURLY_BRACE | PIPE | RIGHT_CURLY_BRACE | TILDE | U_007F));

sip_uri          :  ((CAP_S | S) (CAP_I | I) (CAP_P | P) COLON) ( userinfo )? hostport
                    uri_parameters ( headers )?;
sips_uri         :  ((CAP_S | S) (CAP_I | I) (CAP_P | P) (CAP_S | S) COLON) ( userinfo )? hostport
                    uri_parameters ( headers )?;
userinfo         :  ( user | telephone_subscriber ) ( COLON password )? AT;
user             :  ( unreserved | escaped | user_unreserved )+;
user_unreserved  :  AMPERSAND | EQUALS | PLUS | DOLLAR | COMMA | SEMICOLON | QUESTION | SLASH;
password         :  ( unreserved | escaped |
                    AMPERSAND | EQUALS | PLUS | DOLLAR | COMMA )*;
hostport         :  host ( COLON port )?;
host             :  hostname | ipv4address | ipv6reference;
hostname         :  ( domainlabel PERIOD )* toplabel ( PERIOD )?;
domainlabel      :  alphanum
                    | (alphanum ( alphanum | DASH )* alphanum);
toplabel         :  alpha | (alpha ( alphanum | DASH )* alphanum);
ipv4address    :  digit (digit? | (digit digit)) PERIOD digit (digit? | (digit digit)) PERIOD digit (digit? | (digit digit)) PERIOD digit (digit? | (digit digit));
ipv6reference  :  LEFT_BRACE ipv6address RIGHT_BRACE;
ipv6address    :  hexpart ( COLON ipv4address )?;
hexpart        :  hexseq | (hexseq (COLON COLON) ( hexseq )?) | ((COLON COLON) ( hexseq )?);
hexseq         :  hex4 ( COLON hex4)*;
hex4           :  hexdig (hexdig? | (hexdig hexdig) | (hexdig hexdig hexdig));
port           :  digit+;

//   The BNF for telephone-subscriber can be found in RFC 2806 [9].  Note,
//   however, that any characters allowed there that are not allowed in
//   the user part of the SIP URI MUST be escaped.

uri_parameters    :  ( SEMICOLON uri_parameter)*;
uri_parameter     :  transport_param | user_param | method_param
                     | ttl_param | maddr_param | lr_param | other_param;
transport_param   :  ((CAP_T | T) (CAP_R | R) (CAP_A | A) (CAP_N | N) (CAP_S | S) (CAP_P | P) (CAP_O | O) (CAP_R | R) (CAP_T | T) EQUALS)
                     ( ((CAP_U | U) (CAP_D | D) (CAP_P | P)) | ((CAP_T | T) (CAP_C | C) (CAP_P | P)) | ((CAP_S | S) (CAP_C | C) (CAP_T | T) (CAP_P | P)) | ((CAP_T | T) (CAP_L | L) (CAP_S | S))
                     | other_transport);
other_transport   :  token;
user_param        :  ((CAP_U | U) (CAP_S | S) (CAP_E | E) (CAP_R | R) EQUALS) ( ((CAP_P | P) (CAP_H | H) (CAP_O | O) (CAP_N | N) (CAP_E | E)) | ((CAP_I | I) (CAP_P | P)) | other_user);
other_user        :  token;
method_param      :  ((CAP_M | M) (CAP_E | E) (CAP_T | T) (CAP_H | H) (CAP_O | O) (CAP_D | D) EQUALS) method;
ttl_param         :  ((CAP_T | T) (CAP_T | T) (CAP_L | L) EQUALS) ttl;
maddr_param       :  ((CAP_M | M) (CAP_A | A) (CAP_D | D) (CAP_D | D) (CAP_R | R) EQUALS) host;
lr_param          :  ((CAP_L | L) (CAP_R | R));
other_param       :  pname ( EQUALS pvalue )?;
pname             :  paramchar+;
pvalue            :  paramchar+;
paramchar         :  param_unreserved | unreserved | escaped;
param_unreserved  :  LEFT_BRACE | RIGHT_BRACE | SLASH | COLON | AMPERSAND | PLUS | DOLLAR;

headers         :  QUESTION header ( AMPERSAND header )*;
header          :  hname EQUALS hvalue;
hname           :  ( hnv_unreserved | unreserved | escaped )+;
hvalue          :  ( hnv_unreserved | unreserved | escaped )*;
hnv_unreserved  :  LEFT_BRACE | RIGHT_BRACE | SLASH | QUESTION | COLON | PLUS | DOLLAR;

sip_message    :  request | response;
request        :  request_line
                  ( message_header )*
                  crlf
                  ( message_body )?;
request_line   :  method sp request_uri sp sip_version crlf;
request_uri    :  sip_uri | sips_uri | absoluteuri;
absoluteuri    :  scheme COLON ( hier_part | opaque_part );
hier_part      :  ( net_path | abs_path ) ( QUESTION query )?;
net_path       :  (SLASH SLASH) authority ( abs_path )?;
abs_path       :  SLASH path_segments;
opaque_part    :  uric_no_slash uric*;
uric           :  reserved | unreserved | escaped;
uric_no_slash  :  unreserved | escaped | SEMICOLON | QUESTION | COLON | AT
                  | AMPERSAND | EQUALS | PLUS | DOLLAR | COMMA;
path_segments  :  segment ( SLASH segment )*;
segment        :  pchar* ( SEMICOLON param )*;
param          :  pchar*;
pchar          :  unreserved | escaped |
                  COLON | AT | AMPERSAND | EQUALS | PLUS | DOLLAR | COMMA;
scheme         :  alpha ( alpha | digit | PLUS | DASH | PERIOD )*;
authority      :  srvr | reg_name;
srvr           :  ( ( userinfo )? hostport )?;       // remove extra @ per LDT work
reg_name       :  ( unreserved | escaped | DOLLAR | COMMA
                  | SEMICOLON | COLON | AT | AMPERSAND | EQUALS | PLUS )+;
query          :  uric*;
sip_version    :  ((CAP_S | S) (CAP_I | I) (CAP_P | P)) SLASH digit+ PERIOD digit+;

message_header  :  (accept
                |  accept_encoding
                |  accept_language
                |  alert_info
                |  allow
                |  authentication_info
                |  authorization
                |  call_id
                |  call_info
                |  contact
                |  content_disposition
                |  content_encoding
                |  content_language
                |  content_length
                |  content_type
                |  cseq
                |  date
                |  error_info
                |  expires
                |  from
                |  in_reply_to
                |  max_forwards
                |  mime_version
                |  min_expires
                |  organization
                |  priority
                |  proxy_authenticate
                |  proxy_authorization
                |  proxy_require
                |  record_route
                |  reply_to
                |  require
                |  retry_after
                |  route
                |  server
                |  subject
                |  supported
                |  timestamp
                |  to
                |  unsupported
                |  user_agent
                |  via
                |  warning
                |  www_authenticate
                |  extension_header) crlf;

invitem           :  (CAP_I|CAP_N|CAP_V|CAP_I|CAP_T|CAP_E); // INVITE in caps
ackm              :  (CAP_A|CAP_C|CAP_K); // ACK in caps
optionsm          :  (CAP_O|CAP_P|CAP_T|CAP_I|CAP_O|CAP_N|CAP_S); // OPTIONS in caps
byem              :  (CAP_B|CAP_Y|CAP_E); // BYE in caps
cancelm           :  (CAP_C|CAP_A|CAP_N|CAP_C|CAP_E|CAP_L); // CANCEL in caps
registerm         :  (CAP_R|CAP_E|CAP_G|CAP_I|CAP_S|CAP_T|CAP_E|CAP_R); // REGISTER in caps
method            :  invitem | ackm | optionsm | byem
                     | cancelm | registerm
                     | extension_method;
extension_method  :  token;
response          :  status_line
                     ( message_header )*
                     crlf
                     ( message_body )?;

status_line     :  sip_version sp status_code sp reason_phrase crlf;
status_code     :  informational
               |   redirection
               |   success
               |   client_error
               |   server_error
               |   global_failure
               |   extension_code;
extension_code  :  digit digit digit;
reason_phrase   :  (reserved | unreserved | escaped
                   | utf8_nonascii | utf8_cont | sp | htab)*;

informational  :  (ONE ZERO ZERO)  //  Trying
              |   (ONE EIGHT ZERO)  //  Ringing
              |   (ONE EIGHT ONE)  //  Call Is Being Forwarded
              |   (ONE EIGHT TWO)  //  Queued
              |   (ONE EIGHT THREE);  //  Session Progress
success  :  (TWO ZERO ZERO);  //  OK

redirection  :  (THREE ZERO ZERO)  //  Multiple Choices
            |   (THREE ZERO ONE)  //  Moved Permanently
            |   (THREE ZERO TWO)  //  Moved Temporarily
            |   (THREE ZERO FIVE)  //  Use Proxy
            |   (THREE EIGHT ZERO);  //  Alternative Service

client_error  :  (FOUR ZERO ZERO)  //  Bad Request
             |   (FOUR ZERO ONE)  //  Unauthorized
             |   (FOUR ZERO TWO)  //  Payment Required
             |   (FOUR ZERO THREE)  //  Forbidden
             |   (FOUR ZERO FOUR)  //  Not Found
             |   (FOUR ZERO FIVE)  //  Method Not Allowed
             |   (FOUR ZERO SIX)  //  Not Acceptable
             |   (FOUR ZERO SEVEN)  //  Proxy Authentication Required
             |   (FOUR ZERO EIGHT)  //  Request Timeout
             |   (FOUR ONE ZERO)  //  Gone
             |   (FOUR ONE THREE)  //  Request Entity Too Large
             |   (FOUR ONE FOUR)  //  Request-URI Too Large
             |   (FOUR ONE FIVE)  //  Unsupported Media Type
             |   (FOUR ONE SIX)  //  Unsupported URI Scheme
             |   (FOUR TWO ZERO)  //  Bad Extension
             |   (FOUR TWO ONE)  //  Extension Required
             |   (FOUR TWO THREE)  //  Interval Too Brief
             |   (FOUR EIGHT ZERO)  //  Temporarily not available
             |   (FOUR EIGHT ONE)  //  Call Leg/Transaction Does Not Exist
             |   (FOUR EIGHT TWO)  //  Loop Detected
             |   (FOUR EIGHT THREE)  //  Too Many Hops
             |   (FOUR EIGHT FOUR)  //  Address Incomplete
             |   (FOUR EIGHT FIVE)  //  Ambiguous
             |   (FOUR EIGHT SIX)  //  Busy Here
             |   (FOUR EIGHT SEVEN)  //  Request Terminated
             |   (FOUR EIGHT EIGHT)  //  Not Acceptable Here
             |   (FOUR NINE ONE)  //  Request Pending
             |   (FOUR NINE THREE);  //  Undecipherable

server_error  :  (FIVE ZERO ZERO)  //  Internal Server Error
             |   (FIVE ZERO ONE)  //  Not Implemented
             |   (FIVE ZERO TWO)  //  Bad Gateway
             |   (FIVE ZERO THREE)  //  Service Unavailable
             |   (FIVE ZERO FOUR)  //  Server Time-out
             |   (FIVE ZERO FIVE)  //  SIP Version not supported
             |   (FIVE ONE THREE);  //  Message Too Large
global_failure  :  (SIX ZERO ZERO)  //  Busy Everywhere
               |   (SIX ZERO THREE)  //  Decline
               |   (SIX ZERO FOUR)  //  Does not exist anywhere
               |   (SIX ZERO SIX);  //  Not Acceptable

accept         :  ((CAP_A | A) (CAP_C | C) (CAP_C | C) (CAP_E | E) (CAP_P | P) (CAP_T | T)) hcolon
                   ( accept_range (comma accept_range)* )?;
accept_range   :  media_range (semi accept_param)*;
media_range    :  ( (ASTERISK SLASH ASTERISK)
                  | ( m_type slash ASTERISK )
                  | ( m_type slash m_subtype )
                  ) ( semi m_parameter )*;
accept_param   :  ((CAP_Q | Q) equal qvalue) | generic_param;
qvalue         :  ( ZERO ( PERIOD  (digit? | (digit digit) | (digit digit digit)) )? )
                  | ( ONE ( PERIOD  ((ZERO)? | ((ZERO) (ZERO)) | ((ZERO) (ZERO) (ZERO))) )? );
generic_param  :  token ( equal gen_value )?;
gen_value      :  token | host | quoted_string;

accept_encoding  :  ((CAP_A | A) (CAP_C | C) (CAP_C | C) (CAP_E | E) (CAP_P | P) (CAP_T | T) DASH (CAP_E | E) (CAP_N | N) (CAP_C | C) (CAP_O | O) (CAP_D | D) (CAP_I | I) (CAP_N | N) (CAP_G | G)) hcolon
                     ( encoding (comma encoding)* )?;
encoding         :  codings (semi accept_param)*;
codings          :  content_coding | ASTERISK;
content_coding   :  token;

accept_language  :  ((CAP_A | A) (CAP_C | C) (CAP_C | C) (CAP_E | E) (CAP_P | P) (CAP_T | T) DASH (CAP_L | L) (CAP_A | A) (CAP_N | N) (CAP_G | G) (CAP_U | U) (CAP_A | A) (CAP_G | G) (CAP_E | E)) hcolon
                     ( language (comma language)* )?;
language         :  language_range (semi accept_param)*;
language_range   :  ( ( alpha (alpha? | (alpha alpha) | (alpha alpha alpha) | (alpha alpha alpha alpha) | (alpha alpha alpha alpha alpha) | (alpha alpha alpha alpha alpha alpha) | (alpha alpha alpha alpha alpha alpha alpha)) ( DASH alpha (alpha? | (alpha alpha) | (alpha alpha alpha) | (alpha alpha alpha alpha) | (alpha alpha alpha alpha alpha) | (alpha alpha alpha alpha alpha alpha) | (alpha alpha alpha alpha alpha alpha alpha)) )* ) | ASTERISK );

alert_info   :  ((CAP_A | A) (CAP_L | L) (CAP_E | E) (CAP_R | R) (CAP_T | T) DASH (CAP_I | I) (CAP_N | N) (CAP_F | F) (CAP_O | O)) hcolon alert_param (comma alert_param)*;
alert_param  :  laquot absoluteuri raquot ( semi generic_param )*;

allow  :  ((CAP_A | A) (CAP_L | L) (CAP_L | L) (CAP_O | O) (CAP_W | W)) hcolon (method (comma method)*)?;

authorization     :  ((CAP_A | A) (CAP_U | U) (CAP_T | T) (CAP_H | H) (CAP_O | O) (CAP_R | R) (CAP_I | I) (CAP_Z | Z) (CAP_A | A) (CAP_T | T) (CAP_I | I) (CAP_O | O) (CAP_N | N)) hcolon credentials;
credentials       :  (((CAP_D | D) (CAP_I | I) (CAP_G | G) (CAP_E | E) (CAP_S | S) (CAP_T | T)) lws digest_response)
                     | other_response;
digest_response   :  dig_resp (comma dig_resp)*;
dig_resp          :  username | realm | nonce | digest_uri
                      | dresponse | algorithm | cnonce
                      | opaque | message_qop
                      | nonce_count | auth_param;
username          :  ((CAP_U | U) (CAP_S | S) (CAP_E | E) (CAP_R | R) (CAP_N | N) (CAP_A | A) (CAP_M | M) (CAP_E | E)) equal username_value;
username_value    :  quoted_string;
digest_uri        :  ((CAP_U | U) (CAP_R | R) (CAP_I | I)) equal ldquot digest_uri_value rdquot;
digest_uri_value  :  ((SPACE | EXCLAMATION)|(POUND | DOLLAR | PERCENT | AMPERSAND | APOSTROPHE | LEFT_PAREN | RIGHT_PAREN | ASTERISK | PLUS | COMMA | DASH | PERIOD | SLASH | ZERO | ONE | TWO | THREE | FOUR | FIVE | SIX | SEVEN | EIGHT | NINE | COLON | SEMICOLON | LESS_THAN | EQUALS | GREATER_THAN | QUESTION | AT | CAP_A | CAP_B | CAP_C | CAP_D | CAP_E | CAP_F | CAP_G | CAP_H | CAP_I | CAP_J | CAP_K | CAP_L | CAP_M | CAP_N | CAP_O | CAP_P | CAP_Q | CAP_R | CAP_S | CAP_T | CAP_U | CAP_V | CAP_W | CAP_X | CAP_Y | CAP_Z | LEFT_BRACE | BACKSLASH | RIGHT_BRACE | CARAT | UNDERSCORE | ACCENT | A | B | C | D | E | F | G | H | I | J | K | L | M | N | O | P | Q | R | S | T | U | V | W | X | Y | Z | LEFT_CURLY_BRACE | PIPE | RIGHT_CURLY_BRACE | TILDE | U_007F | U_0080 | U_0081 | U_0082 | U_0083 | U_0084 | U_0085 | U_0086 | U_0087 | U_0088 | U_0089 | U_008A | U_008B | U_008C | U_008D | U_008E | U_008F | U_0090 | U_0091 | U_0092 | U_0093 | U_0094 | U_0095 | U_0096 | U_0097 | U_0098 | U_0099 | U_009A | U_009B | U_009C | U_009D | U_009E | U_009F | U_00A0 | U_00A1 | U_00A2 | U_00A3 | U_00A4 | U_00A5 | U_00A6 | U_00A7 | U_00A8 | U_00A9 | U_00AA | U_00AB | U_00AC | U_00AD | U_00AE | U_00AF | U_00B0 | U_00B1 | U_00B2 | U_00B3 | U_00B4 | U_00B5 | U_00B6 | U_00B7 | U_00B8 | U_00B9 | U_00BA | U_00BB | U_00BC | U_00BD | U_00BE | U_00BF | U_00C0 | U_00C1 | U_00C2 | U_00C3 | U_00C4 | U_00C5 | U_00C6 | U_00C7 | U_00C8 | U_00C9 | U_00CA | U_00CB | U_00CC | U_00CD | U_00CE | U_00CF | U_00D0 | U_00D1 | U_00D2 | U_00D3 | U_00D4 | U_00D5 | U_00D6 | U_00D7 | U_00D8 | U_00D9 | U_00DA | U_00DB | U_00DC | U_00DD | U_00DE | U_00DF | U_00E0 | U_00E1 | U_00E2 | U_00E3 | U_00E4 | U_00E5 | U_00E6 | U_00E7 | U_00E8 | U_00E9 | U_00EA | U_00EB | U_00EC | U_00ED | U_00EE | U_00EF | U_00F0 | U_00F1 | U_00F2 | U_00F3 | U_00F4 | U_00F5 | U_00F6 | U_00F7 | U_00F8 | U_00F9 | U_00FA | U_00FB | U_00FC | U_00FD | U_00FE | U_00FF))*;
//digest-uri-value  =  rquest-uri ; Equal to request-uri as specified by HTTP/1.1
message_qop       :  ((CAP_Q | Q) (CAP_O | O) (CAP_P | P)) equal qop_value;
cnonce            :  ((CAP_C | C) (CAP_N | N) (CAP_O | O) (CAP_N | N) (CAP_C | C) (CAP_E | E)) equal cnonce_value;
cnonce_value      :  nonce_value;
nonce_count       :  ((CAP_N | N) (CAP_C | C)) equal nc_value;
nc_value          :  lhex lhex lhex lhex lhex lhex lhex lhex;
dresponse         :  ((CAP_R | R) (CAP_E | E) (CAP_S | S) (CAP_P | P) (CAP_O | O) (CAP_N | N) (CAP_S | S) (CAP_E | E)) equal request_digest;
request_digest    :  ldquot lhex lhex lhex lhex lhex lhex lhex lhex lhex lhex lhex lhex lhex lhex lhex lhex lhex lhex lhex lhex lhex lhex lhex lhex lhex lhex lhex lhex lhex lhex lhex lhex rdquot;
auth_param        :  auth_param_name equal
                     ( token | quoted_string );
auth_param_name   :  token;
other_response    :  auth_scheme lws auth_param
                     (comma auth_param)*;
auth_scheme       :  token;

authentication_info  :  ((CAP_A | A) (CAP_U | U) (CAP_T | T) (CAP_H | H) (CAP_E | E) (CAP_N | N) (CAP_T | T) (CAP_I | I) (CAP_C | C) (CAP_A | A) (CAP_T | T) (CAP_I | I) (CAP_O | O) (CAP_N | N) DASH (CAP_I | I) (CAP_N | N) (CAP_F | F) (CAP_O | O)) hcolon ainfo
                        (comma ainfo)*;
ainfo                :  nextnonce | message_qop
                         | response_auth | cnonce
                         | nonce_count;
nextnonce            :  ((CAP_N | N) (CAP_E | E) (CAP_X | X) (CAP_T | T) (CAP_N | N) (CAP_O | O) (CAP_N | N) (CAP_C | C) (CAP_E | E)) equal nonce_value;
response_auth        :  ((CAP_R | R) (CAP_S | S) (CAP_P | P) (CAP_A | A) (CAP_U | U) (CAP_T | T) (CAP_H | H)) equal response_digest;
response_digest      :  ldquot lhex* rdquot;

call_id  :  ( ((CAP_C | C) (CAP_A | A) (CAP_L | L) (CAP_L | L) DASH (CAP_I | I) (CAP_D | D)) | (CAP_I | I) ) hcolon callid;
callid   :  word ( AT word )?;

call_info   :  ((CAP_C | C) (CAP_A | A) (CAP_L | L) (CAP_L | L) DASH (CAP_I | I) (CAP_N | N) (CAP_F | F) (CAP_O | O)) hcolon info (comma info)*;
info        :  laquot absoluteuri raquot ( semi info_param)*;
info_param  :  ( ((CAP_P | P) (CAP_U | U) (CAP_R | R) (CAP_P | P) (CAP_O | O) (CAP_S | S) (CAP_E | E)) equal ( ((CAP_I | I) (CAP_C | C) (CAP_O | O) (CAP_N | N)) | ((CAP_I | I) (CAP_N | N) (CAP_F | F) (CAP_O | O))
               | ((CAP_C | C) (CAP_A | A) (CAP_R | R) (CAP_D | D)) | token ) ) | generic_param;

contact        :  (((CAP_C | C) (CAP_O | O) (CAP_N | N) (CAP_T | T) (CAP_A | A) (CAP_C | C) (CAP_T | T)) | (CAP_M | M) ) hcolon
                  ( star | (contact_param (comma contact_param)*));
contact_param  :  (name_addr | addr_spec) (semi contact_params)*;
name_addr      :  ( display_name )? laquot addr_spec raquot;
addr_spec      :  sip_uri | sips_uri | absoluteuri;
display_name   :  (token lws)*| quoted_string;

contact_params     :  c_p_q | c_p_expires
                      | contact_extension;
c_p_q              :  (CAP_Q | Q) equal qvalue;
c_p_expires        :  ((CAP_E | E) (CAP_X | X) (CAP_P | P) (CAP_I | I) (CAP_R | R) (CAP_E | E) (CAP_S | S)) equal delta_seconds;
contact_extension  :  generic_param;
delta_seconds      :  digit+;

content_disposition   :  ((CAP_C | C) (CAP_O | O) (CAP_N | N) (CAP_T | T) (CAP_E | E) (CAP_N | N) (CAP_T | T) DASH (CAP_D | D) (CAP_I | I) (CAP_S | S) (CAP_P | P) (CAP_O | O) (CAP_S | S) (CAP_I | I) (CAP_T | T) (CAP_I | I) (CAP_O | O) (CAP_N | N)) hcolon
                         disp_type ( semi disp_param )*;
disp_type             :  ((CAP_R | R) (CAP_E | E) (CAP_N | N) (CAP_D | D) (CAP_E | E) (CAP_R | R)) | ((CAP_S | S) (CAP_E | E) (CAP_S | S) (CAP_S | S) (CAP_I | I) (CAP_O | O) (CAP_N | N)) | ((CAP_I | I) (CAP_C | C) (CAP_O | O) (CAP_N | N)) | ((CAP_A | A) (CAP_L | L) (CAP_E | E) (CAP_R | R) (CAP_T | T))
                         | disp_extension_token;
disp_param            :  handling_param | generic_param;
handling_param        :  ((CAP_H | H) (CAP_A | A) (CAP_N | N) (CAP_D | D) (CAP_L | L) (CAP_I | I) (CAP_N | N) (CAP_G | G)) equal
                         ( ((CAP_O | O) (CAP_P | P) (CAP_T | T) (CAP_I | I) (CAP_O | O) (CAP_N | N) (CAP_A | A) (CAP_L | L)) | ((CAP_R | R) (CAP_E | E) (CAP_Q | Q) (CAP_U | U) (CAP_I | I) (CAP_R | R) (CAP_E | E) (CAP_D | D))
                         | other_handling );
other_handling        :  token;
disp_extension_token  :  token;

content_encoding  :  ( ((CAP_C | C) (CAP_O | O) (CAP_N | N) (CAP_T | T) (CAP_E | E) (CAP_N | N) (CAP_T | T) DASH (CAP_E | E) (CAP_N | N) (CAP_C | C) (CAP_O | O) (CAP_D | D) (CAP_I | I) (CAP_N | N) (CAP_G | G)) | (CAP_E | E) ) hcolon
                     content_coding (comma content_coding)*;

content_language  :  ((CAP_C | C) (CAP_O | O) (CAP_N | N) (CAP_T | T) (CAP_E | E) (CAP_N | N) (CAP_T | T) DASH (CAP_L | L) (CAP_A | A) (CAP_N | N) (CAP_G | G) (CAP_U | U) (CAP_A | A) (CAP_G | G) (CAP_E | E)) hcolon
                     language_tag (comma language_tag)*;
language_tag      :  primary_tag ( DASH subtag )*;
primary_tag       :  alpha (alpha? | (alpha alpha) | (alpha alpha alpha) | (alpha alpha alpha alpha) | (alpha alpha alpha alpha alpha) | (alpha alpha alpha alpha alpha alpha) | (alpha alpha alpha alpha alpha alpha alpha));
subtag            :  alpha (alpha? | (alpha alpha) | (alpha alpha alpha) | (alpha alpha alpha alpha) | (alpha alpha alpha alpha alpha) | (alpha alpha alpha alpha alpha alpha) | (alpha alpha alpha alpha alpha alpha alpha));

content_length  :  ( ((CAP_C | C) (CAP_O | O) (CAP_N | N) (CAP_T | T) (CAP_E | E) (CAP_N | N) (CAP_T | T) DASH (CAP_L | L) (CAP_E | E) (CAP_N | N) (CAP_G | G) (CAP_T | T) (CAP_H | H)) | (CAP_L | L) ) hcolon digit+;
content_type     :  ( ((CAP_C | C) (CAP_O | O) (CAP_N | N) (CAP_T | T) (CAP_E | E) (CAP_N | N) (CAP_T | T) DASH (CAP_T | T) (CAP_Y | Y) (CAP_P | P) (CAP_E | E)) | (CAP_C | C) ) hcolon media_type;
media_type       :  m_type slash m_subtype (semi m_parameter)*;
m_type           :  discrete_type | composite_type;
discrete_type    :  ((CAP_T | T) (CAP_E | E) (CAP_X | X) (CAP_T | T)) | ((CAP_I | I) (CAP_M | M) (CAP_A | A) (CAP_G | G) (CAP_E | E)) | ((CAP_A | A) (CAP_U | U) (CAP_D | D) (CAP_I | I) (CAP_O | O)) | ((CAP_V | V) (CAP_I | I) (CAP_D | D) (CAP_E | E) (CAP_O | O))
                    | ((CAP_A | A) (CAP_P | P) (CAP_P | P) (CAP_L | L) (CAP_I | I) (CAP_C | C) (CAP_A | A) (CAP_T | T) (CAP_I | I) (CAP_O | O) (CAP_N | N)) | extension_token;
composite_type   :  ((CAP_M | M) (CAP_E | E) (CAP_S | S) (CAP_S | S) (CAP_A | A) (CAP_G | G) (CAP_E | E)) | ((CAP_M | M) (CAP_U | U) (CAP_L | L) (CAP_T | T) (CAP_I | I) (CAP_P | P) (CAP_A | A) (CAP_R | R) (CAP_T | T)) | extension_token;
extension_token  :  ietf_token | x_token;
ietf_token       :  token;
x_token          :  ((CAP_X | X) DASH) token;
m_subtype        :  extension_token | iana_token;
iana_token       :  token;
m_parameter      :  m_attribute equal m_value;
m_attribute      :  token;
m_value          :  token | quoted_string;

cseq  :  ((CAP_C | C) (CAP_S | S) (CAP_E | E) (CAP_Q | Q)) hcolon digit+ lws method;

date          :  ((CAP_D | D) (CAP_A | A) (CAP_T | T) (CAP_E | E)) hcolon sip_date;
sip_date      :  rfc1123_date;
rfc1123_date  :  wkday COMMA sp date1 sp time sp ((CAP_G | G) (CAP_M | M) (CAP_T | T));
date1         :  digit digit sp month sp digit digit digit digit;
                 // day month year (e.g., 02 Jun 1982)
time          :  digit digit COLON digit digit COLON digit digit;
                 // 00:00:00 - 23:59:59
wkday         :  ((CAP_M | M) (CAP_O | O) (CAP_N | N)) | ((CAP_T | T) (CAP_U | U) (CAP_E | E)) | ((CAP_W | W) (CAP_E | E) (CAP_D | D))
                 | ((CAP_T | T) (CAP_H | H) (CAP_U | U)) | ((CAP_F | F) (CAP_R | R) (CAP_I | I)) | ((CAP_S | S) (CAP_A | A) (CAP_T | T)) | ((CAP_S | S) (CAP_U | U) (CAP_N | N));
month         :  ((CAP_J | J) (CAP_A | A) (CAP_N | N)) | ((CAP_F | F) (CAP_E | E) (CAP_B | B)) | ((CAP_M | M) (CAP_A | A) (CAP_R | R)) | ((CAP_A | A) (CAP_P | P) (CAP_R | R))
                 | ((CAP_M | M) (CAP_A | A) (CAP_Y | Y)) | ((CAP_J | J) (CAP_U | U) (CAP_N | N)) | ((CAP_J | J) (CAP_U | U) (CAP_L | L)) | ((CAP_A | A) (CAP_U | U) (CAP_G | G))
                 | ((CAP_S | S) (CAP_E | E) (CAP_P | P)) | ((CAP_O | O) (CAP_C | C) (CAP_T | T)) | ((CAP_N | N) (CAP_O | O) (CAP_V | V)) | ((CAP_D | D) (CAP_E | E) (CAP_C | C));

error_info  :  ((CAP_E | E) (CAP_R | R) (CAP_R | R) (CAP_O | O) (CAP_R | R) DASH (CAP_I | I) (CAP_N | N) (CAP_F | F) (CAP_O | O)) hcolon error_uri (comma error_uri)*;
error_uri   :  laquot absoluteuri raquot ( semi generic_param )*;

expires     :  ((CAP_E | E) (CAP_X | X) (CAP_P | P) (CAP_I | I) (CAP_R | R) (CAP_E | E) (CAP_S | S)) hcolon delta_seconds;
from        :  ( ((CAP_F | F) (CAP_R | R) (CAP_O | O) (CAP_M | M)) | (CAP_F | F) ) hcolon from_spec;
from_spec   :  ( name_addr | addr_spec )
               ( semi from_param )*;
from_param  :  tag_param | generic_param;
tag_param   :  ((CAP_T | T) (CAP_A | A) (CAP_G | G)) equal token;

in_reply_to  :  ((CAP_I | I) (CAP_N | N) DASH (CAP_R | R) (CAP_E | E) (CAP_P | P) (CAP_L | L) (CAP_Y | Y) DASH (CAP_T | T) (CAP_O | O)) hcolon callid (comma callid)*;

max_forwards  :  ((CAP_M | M) (CAP_A | A) (CAP_X | X) DASH (CAP_F | F) (CAP_O | O) (CAP_R | R) (CAP_W | W) (CAP_A | A) (CAP_R | R) (CAP_D | D) (CAP_S | S)) hcolon digit+;

mime_version  :  ((CAP_M | M) (CAP_I | I) (CAP_M | M) (CAP_E | E) DASH (CAP_V | V) (CAP_E | E) (CAP_R | R) (CAP_S | S) (CAP_I | I) (CAP_O | O) (CAP_N | N)) hcolon digit+ PERIOD digit+;

min_expires  :  ((CAP_M | M) (CAP_I | I) (CAP_N | N) DASH (CAP_E | E) (CAP_X | X) (CAP_P | P) (CAP_I | I) (CAP_R | R) (CAP_E | E) (CAP_S | S)) hcolon delta_seconds;

organization  :  ((CAP_O | O) (CAP_R | R) (CAP_G | G) (CAP_A | A) (CAP_N | N) (CAP_I | I) (CAP_Z | Z) (CAP_A | A) (CAP_T | T) (CAP_I | I) (CAP_O | O) (CAP_N | N)) hcolon (text_utf8_trim)?;

priority        :  ((CAP_P | P) (CAP_R | R) (CAP_I | I) (CAP_O | O) (CAP_R | R) (CAP_I | I) (CAP_T | T) (CAP_Y | Y)) hcolon priority_value;
priority_value  :  ((CAP_E | E) (CAP_M | M) (CAP_E | E) (CAP_R | R) (CAP_G | G) (CAP_E | E) (CAP_N | N) (CAP_C | C) (CAP_Y | Y)) | ((CAP_U | U) (CAP_R | R) (CAP_G | G) (CAP_E | E) (CAP_N | N) (CAP_T | T)) | ((CAP_N | N) (CAP_O | O) (CAP_R | R) (CAP_M | M) (CAP_A | A) (CAP_L | L))
                   | ((CAP_N | N) (CAP_O | O) (CAP_N | N) DASH (CAP_U | U) (CAP_R | R) (CAP_G | G) (CAP_E | E) (CAP_N | N) (CAP_T | T)) | other_priority;
other_priority  :  token;

proxy_authenticate  :  ((CAP_P | P) (CAP_R | R) (CAP_O | O) (CAP_X | X) (CAP_Y | Y) DASH (CAP_A | A) (CAP_U | U) (CAP_T | T) (CAP_H | H) (CAP_E | E) (CAP_N | N) (CAP_T | T) (CAP_I | I) (CAP_C | C) (CAP_A | A) (CAP_T | T) (CAP_E | E)) hcolon challenge;
challenge           :  (((CAP_D | D) (CAP_I | I) (CAP_G | G) (CAP_E | E) (CAP_S | S) (CAP_T | T)) lws digest_cln (comma digest_cln)*)
                       | other_challenge;
other_challenge     :  auth_scheme lws auth_param
                       (comma auth_param)*;
digest_cln          :  realm | domain | nonce
                        | opaque | stale | algorithm
                        | qop_options | auth_param;
realm               :  ((CAP_R | R) (CAP_E | E) (CAP_A | A) (CAP_L | L) (CAP_M | M)) equal realm_value;
realm_value         :  quoted_string;
domain              :  ((CAP_D | D) (CAP_O | O) (CAP_M | M) (CAP_A | A) (CAP_I | I) (CAP_N | N)) equal ldquot uri
                       ( sp+ uri )* rdquot;
uri                 :  absoluteuri | abs_path;
nonce               :  ((CAP_N | N) (CAP_O | O) (CAP_N | N) (CAP_C | C) (CAP_E | E)) equal nonce_value;
nonce_value         :  quoted_string;
opaque              :  ((CAP_O | O) (CAP_P | P) (CAP_A | A) (CAP_Q | Q) (CAP_U | U) (CAP_E | E)) equal quoted_string;
stale               :  ((CAP_S | S) (CAP_T | T) (CAP_A | A) (CAP_L | L) (CAP_E | E)) equal ( ((CAP_T | T) (CAP_R | R) (CAP_U | U) (CAP_E | E)) | ((CAP_F | F) (CAP_A | A) (CAP_L | L) (CAP_S | S) (CAP_E | E)) );
algorithm           :  ((CAP_A | A) (CAP_L | L) (CAP_G | G) (CAP_O | O) (CAP_R | R) (CAP_I | I) (CAP_T | T) (CAP_H | H) (CAP_M | M)) equal ( ((CAP_M | M) (CAP_D | D) FIVE) | ((CAP_M | M) (CAP_D | D) FIVE DASH (CAP_S | S) (CAP_E | E) (CAP_S | S) (CAP_S | S))
                       | token );
qop_options         :  ((CAP_Q | Q) (CAP_O | O) (CAP_P | P)) equal ldquot qop_value
                       (COMMA qop_value)* rdquot;
qop_value           :  ((CAP_A | A) (CAP_U | U) (CAP_T | T) (CAP_H | H)) | ((CAP_A | A) (CAP_U | U) (CAP_T | T) (CAP_H | H) DASH (CAP_I | I) (CAP_N | N) (CAP_T | T)) | token;

proxy_authorization  :  ((CAP_P | P) (CAP_R | R) (CAP_O | O) (CAP_X | X) (CAP_Y | Y) DASH (CAP_A | A) (CAP_U | U) (CAP_T | T) (CAP_H | H) (CAP_O | O) (CAP_R | R) (CAP_I | I) (CAP_Z | Z) (CAP_A | A) (CAP_T | T) (CAP_I | I) (CAP_O | O) (CAP_N | N)) hcolon credentials;
proxy_require  :  ((CAP_P | P) (CAP_R | R) (CAP_O | O) (CAP_X | X) (CAP_Y | Y) DASH (CAP_R | R) (CAP_E | E) (CAP_Q | Q) (CAP_U | U) (CAP_I | I) (CAP_R | R) (CAP_E | E)) hcolon option_tag
                  (comma option_tag)*;
option_tag     :  token;

record_route  :  ((CAP_R | R) (CAP_E | E) (CAP_C | C) (CAP_O | O) (CAP_R | R) (CAP_D | D) DASH (CAP_R | R) (CAP_O | O) (CAP_U | U) (CAP_T | T) (CAP_E | E)) hcolon rec_route (comma rec_route)*;
rec_route     :  name_addr ( semi rr_param )*;
rr_param      :  generic_param;

reply_to      :  ((CAP_R | R) (CAP_E | E) (CAP_P | P) (CAP_L | L) (CAP_Y | Y) DASH (CAP_T | T) (CAP_O | O)) hcolon rplyto_spec;
rplyto_spec   :  ( name_addr | addr_spec )
                 ( semi rplyto_param )*;
rplyto_param  :  generic_param;
require       :  ((CAP_R | R) (CAP_E | E) (CAP_Q | Q) (CAP_U | U) (CAP_I | I) (CAP_R | R) (CAP_E | E)) hcolon option_tag (comma option_tag)*;

retry_after  :  ((CAP_R | R) (CAP_E | E) (CAP_T | T) (CAP_R | R) (CAP_Y | Y) DASH (CAP_A | A) (CAP_F | F) (CAP_T | T) (CAP_E | E) (CAP_R | R)) hcolon delta_seconds
                ( comment )? ( semi retry_param )*;

retry_param  :  (((CAP_D | D) (CAP_U | U) (CAP_R | R) (CAP_A | A) (CAP_T | T) (CAP_I | I) (CAP_O | O) (CAP_N | N)) equal delta_seconds)
                | generic_param;

route        :  ((CAP_R | R) (CAP_O | O) (CAP_U | U) (CAP_T | T) (CAP_E | E)) hcolon route_param (comma route_param)*;
route_param  :  name_addr ( semi rr_param )*;

server           :  ((CAP_S | S) (CAP_E | E) (CAP_R | R) (CAP_V | V) (CAP_E | E) (CAP_R | R)) hcolon server_val (lws server_val)*;
server_val       :  product | comment;
product          :  token (slash product_version)?;
product_version  :  token;

subject  :  ( ((CAP_S | S) (CAP_U | U) (CAP_B | B) (CAP_J | J) (CAP_E | E) (CAP_C | C) (CAP_T | T)) | (CAP_S | S) ) hcolon (text_utf8_trim)?;

supported  :  ( ((CAP_S | S) (CAP_U | U) (CAP_P | P) (CAP_P | P) (CAP_O | O) (CAP_R | R) (CAP_T | T) (CAP_E | E) (CAP_D | D)) | (CAP_K | K) ) hcolon
              (option_tag (comma option_tag)*)?;

timestamp  :  ((CAP_T | T) (CAP_I | I) (CAP_M | M) (CAP_E | E) (CAP_S | S) (CAP_T | T) (CAP_A | A) (CAP_M | M) (CAP_P | P)) hcolon (digit)+
               ( PERIOD (digit)* )? ( lws delay )?;
delay      :  (digit)* ( PERIOD (digit)* )?;

to        :  ( ((CAP_T | T) (CAP_O | O)) | (CAP_T | T) ) hcolon ( name_addr
             | addr_spec ) ( semi to_param )*;
to_param  :  tag_param | generic_param;

unsupported  :  ((CAP_U | U) (CAP_N | N) (CAP_S | S) (CAP_U | U) (CAP_P | P) (CAP_P | P) (CAP_O | O) (CAP_R | R) (CAP_T | T) (CAP_E | E) (CAP_D | D)) hcolon option_tag (comma option_tag)*;
user_agent  :  ((CAP_U | U) (CAP_S | S) (CAP_E | E) (CAP_R | R) DASH (CAP_A | A) (CAP_G | G) (CAP_E | E) (CAP_N | N) (CAP_T | T)) hcolon server_val (lws server_val)*;
via               :  ( ((CAP_V | V) (CAP_I | I) (CAP_A | A)) | (CAP_V | V) ) hcolon via_parm (comma via_parm)*;
via_parm          :  sent_protocol lws sent_by ( semi via_params )*;
via_params        :  via_ttl | via_maddr
                     | via_received | via_branch
                     | via_extension;
via_ttl           :  ((CAP_T | T) (CAP_T | T) (CAP_L | L)) equal ttl;
via_maddr         :  ((CAP_M | M) (CAP_A | A) (CAP_D | D) (CAP_D | D) (CAP_R | R)) equal host;
via_received      :  ((CAP_R | R) (CAP_E | E) (CAP_C | C) (CAP_E | E) (CAP_I | I) (CAP_V | V) (CAP_E | E) (CAP_D | D)) equal (ipv4address | ipv6address);
via_branch        :  ((CAP_B | B) (CAP_R | R) (CAP_A | A) (CAP_N | N) (CAP_C | C) (CAP_H | H)) equal token;
via_extension     :  generic_param;
sent_protocol     :  protocol_name slash protocol_version
                     slash transport;
protocol_name     :  ((CAP_S | S) (CAP_I | I) (CAP_P | P)) | token;
protocol_version  :  token;
transport         :  ((CAP_U | U) (CAP_D | D) (CAP_P | P)) | ((CAP_T | T) (CAP_C | C) (CAP_P | P)) | ((CAP_T | T) (CAP_L | L) (CAP_S | S)) | ((CAP_S | S) (CAP_C | C) (CAP_T | T) (CAP_P | P))
                     | other_transport;
sent_by           :  host ( colon port )?;
ttl               :  digit (digit? | (digit digit)); // 0 to 255

warning        :  ((CAP_W | W) (CAP_A | A) (CAP_R | R) (CAP_N | N) (CAP_I | I) (CAP_N | N) (CAP_G | G)) hcolon warning_value (comma warning_value)*;
warning_value  :  warn_code sp warn_agent sp warn_text;
warn_code      :  digit digit digit;
warn_agent     :  hostport | pseudonym;
                  //  the name or pseudonym of the server adding
                  //  the Warning header, for use in debugging
warn_text      :  quoted_string;
pseudonym      :  token;

www_authenticate  :  ((CAP_W | W) (CAP_W | W) (CAP_W | W) DASH (CAP_A | A) (CAP_U | U) (CAP_T | T) (CAP_H | H) (CAP_E | E) (CAP_N | N) (CAP_T | T) (CAP_I | I) (CAP_C | C) (CAP_A | A) (CAP_T | T) (CAP_E | E)) hcolon challenge;

extension_header  :  header_name hcolon header_value;
header_name       :  token;
header_value      :  (text_utf8char | utf8_cont | lws)*;
message_body  :  octet*;

// rfc2806-2
telephone_subscriber  : global_phone_number | local_phone_number;
global_phone_number   : PLUS base_phone_number (isdn_subaddress)?
                        (post_dial)? (area_specifier |
                        service_provider | future_extension)*;
base_phone_number     : phonedigit+;
local_phone_number    : (phonedigit | dtmf_digit |
                        pause_character)+ (isdn_subaddress)?
                        (post_dial)? area_specifier
                        (area_specifier | service_provider |
                        future_extension)*;
isdn_subaddress       : (SEMICOLON (CAP_I | I) (CAP_S | S) (CAP_U | U) (CAP_B | B) EQUALS) phonedigit+;
post_dial             : (SEMICOLON (CAP_P | P) (CAP_O | O) (CAP_S | S) (CAP_T | T) (CAP_D | D) EQUALS) (phonedigit |
                        dtmf_digit | pause_character)+;
area_specifier        : SEMICOLON phone_context_tag EQUALS phone_context_ident;
phone_context_tag     : ((CAP_P | P) (CAP_H | H) (CAP_O | O) (CAP_N | N) (CAP_E | E) DASH (CAP_C | C) (CAP_O | O) (CAP_N | N) (CAP_T | T) (CAP_E | E) (CAP_X | X) (CAP_T | T));
phone_context_ident   : network_prefix | private_prefix;
network_prefix        : global_network_prefix | local_network_prefix;
global_network_prefix : PLUS phonedigit+;
local_network_prefix  : (phonedigit | dtmf_digit | pause_character)+;
private_prefix        : ((EXCLAMATION | QUOTE) | (DOLLAR | PERCENT | AMPERSAND | APOSTROPHE) | COMMA | SLASH | COLON |
                        (LESS_THAN | EQUALS | GREATER_THAN | QUESTION | AT) | (CAP_E | CAP_F | CAP_G | CAP_H | CAP_I | CAP_J | CAP_K | CAP_L | CAP_M | CAP_N | CAP_O) | (CAP_Q | CAP_R | CAP_S | CAP_T | CAP_U | CAP_V) | (CAP_X | CAP_Y | CAP_Z | LEFT_BRACE | BACKSLASH | RIGHT_BRACE | CARAT | UNDERSCORE | ACCENT) |
                        (E | F | G | H | I | J | K | L | M | N | O) | (Q | R | S | T | U | V) | (X | Y | Z | LEFT_CURLY_BRACE | PIPE | RIGHT_CURLY_BRACE | TILDE))
                        ((EXCLAMATION | QUOTE | POUND | DOLLAR | PERCENT | AMPERSAND | APOSTROPHE | LEFT_PAREN | RIGHT_PAREN | ASTERISK | PLUS | COMMA | DASH | PERIOD | SLASH | ZERO | ONE | TWO | THREE | FOUR | FIVE | SIX | SEVEN | EIGHT | NINE | COLON) | (LESS_THAN | EQUALS | GREATER_THAN | QUESTION | AT | CAP_A | CAP_B | CAP_C | CAP_D | CAP_E | CAP_F | CAP_G | CAP_H | CAP_I | CAP_J | CAP_K | CAP_L | CAP_M | CAP_N | CAP_O | CAP_P | CAP_Q | CAP_R | CAP_S | CAP_T | CAP_U | CAP_V | CAP_W | CAP_X | CAP_Y | CAP_Z | LEFT_BRACE | BACKSLASH | RIGHT_BRACE | CARAT | UNDERSCORE | ACCENT | A | B | C | D | E | F | G | H | I | J | K | L | M | N | O | P | Q | R | S | T | U | V | W | X | Y | Z | LEFT_CURLY_BRACE | PIPE | RIGHT_CURLY_BRACE | TILDE))*;
                        // Characters in URLs must follow escaping rules
                        // as explained in [RFC2396]
                        // See sections 1.2 and 2.5.2
service_provider      : SEMICOLON provider_tag EQUALS provider_hostname;
provider_tag          : ((CAP_T | T) (CAP_S | S) (CAP_P | P));
provider_hostname     : domain; // <domain> is defined in [RFC1035]
                        // See section 2.5.10
future_extension      : SEMICOLON (token_char)+ (EQUALS (((token_char)+
                        (QUESTION (token_char)+)?) | quoted_string ))?;
                        // See section 2.5.11 and [RFC2543]
token_char            : (EXCLAMATION | (POUND | DOLLAR | PERCENT | AMPERSAND | APOSTROPHE) | (ASTERISK | PLUS) | (DASH | PERIOD) | (ZERO | ONE | TWO | THREE | FOUR | FIVE | SIX | SEVEN | EIGHT | NINE)
                        | (CAP_A | CAP_B | CAP_C | CAP_D | CAP_E | CAP_F | CAP_G | CAP_H | CAP_I | CAP_J | CAP_K | CAP_L | CAP_M | CAP_N | CAP_O | CAP_P | CAP_Q | CAP_R | CAP_S | CAP_T | CAP_U | CAP_V | CAP_W | CAP_X | CAP_Y | CAP_Z) | (CARAT | UNDERSCORE | ACCENT | A | B | C | D | E | F | G | H | I | J | K | L | M | N | O | P | Q | R | S | T | U | V | W | X | Y | Z) | PIPE | TILDE);
                        // Characters in URLs must follow escaping rules
                        // as explained in [RFC2396]
                        // See sections 1.2 and 2.5.11
quoted_string         : QUOTE ( (BACKSLASH char_1) | ((SPACE | EXCLAMATION) | (POUND | DOLLAR | PERCENT | AMPERSAND | APOSTROPHE | LEFT_PAREN | RIGHT_PAREN | ASTERISK | PLUS | COMMA | DASH | PERIOD | SLASH | ZERO | ONE | TWO | THREE | FOUR | FIVE | SIX | SEVEN | EIGHT | NINE | COLON | SEMICOLON | LESS_THAN | EQUALS | GREATER_THAN | QUESTION | AT | CAP_A | CAP_B | CAP_C | CAP_D | CAP_E | CAP_F | CAP_G | CAP_H | CAP_I | CAP_J | CAP_K | CAP_L | CAP_M | CAP_N | CAP_O | CAP_P | CAP_Q | CAP_R | CAP_S | CAP_T | CAP_U | CAP_V | CAP_W | CAP_X | CAP_Y | CAP_Z | LEFT_BRACE | BACKSLASH | RIGHT_BRACE | CARAT | UNDERSCORE | ACCENT | A | B | C | D | E | F | G | H | I | J | K | L | M | N | O | P | Q | R | S | T | U | V | W | X | Y | Z | LEFT_CURLY_BRACE | PIPE | RIGHT_CURLY_BRACE | TILDE)
                        | (U_0080 | U_0081 | U_0082 | U_0083 | U_0084 | U_0085 | U_0086 | U_0087 | U_0088 | U_0089 | U_008A | U_008B | U_008C | U_008D | U_008E | U_008F | U_0090 | U_0091 | U_0092 | U_0093 | U_0094 | U_0095 | U_0096 | U_0097 | U_0098 | U_0099 | U_009A | U_009B | U_009C | U_009D | U_009E | U_009F | U_00A0 | U_00A1 | U_00A2 | U_00A3 | U_00A4 | U_00A5 | U_00A6 | U_00A7 | U_00A8 | U_00A9 | U_00AA | U_00AB | U_00AC | U_00AD | U_00AE | U_00AF | U_00B0 | U_00B1 | U_00B2 | U_00B3 | U_00B4 | U_00B5 | U_00B6 | U_00B7 | U_00B8 | U_00B9 | U_00BA | U_00BB | U_00BC | U_00BD | U_00BE | U_00BF | U_00C0 | U_00C1 | U_00C2 | U_00C3 | U_00C4 | U_00C5 | U_00C6 | U_00C7 | U_00C8 | U_00C9 | U_00CA | U_00CB | U_00CC | U_00CD | U_00CE | U_00CF | U_00D0 | U_00D1 | U_00D2 | U_00D3 | U_00D4 | U_00D5 | U_00D6 | U_00D7 | U_00D8 | U_00D9 | U_00DA | U_00DB | U_00DC | U_00DD | U_00DE | U_00DF | U_00E0 | U_00E1 | U_00E2 | U_00E3 | U_00E4 | U_00E5 | U_00E6 | U_00E7 | U_00E8 | U_00E9 | U_00EA | U_00EB | U_00EC | U_00ED | U_00EE | U_00EF | U_00F0 | U_00F1 | U_00F2 | U_00F3 | U_00F4 | U_00F5 | U_00F6 | U_00F7 | U_00F8 | U_00F9 | U_00FA | U_00FB | U_00FC | U_00FD | U_00FE | U_00FF) ))* QUOTE;
                        // Characters in URLs must follow escaping rules
                        // as explained in [RFC2396]
                        // See sections 1.2 and 2.5.11
phonedigit            : digit | visual_separator;
visual_separator      : DASH | PERIOD | LEFT_PAREN | RIGHT_PAREN;
pause_character       : one_second_pause | wait_for_dial_tone;
one_second_pause      : (CAP_P | P);
wait_for_dial_tone    : (CAP_W | W);
dtmf_digit            : ASTERISK | POUND | (CAP_A | A) | (CAP_B | B) | (CAP_C | C) | (CAP_D | D);





// Full Copyright Statement
//
//  Copyright (C) The Internet Society (2002).  All Rights Reserved.
//
//  This document and translations of it may be copied and furnished to
//  others, and derivative works that comment on or otherwise explain it
//  or assist in its implementation may be prepared, copied, published
//  and distributed, in whole or in part, without restriction of any
//  kind, provided that the above copyright notice and this paragraph are
//  included on all such copies and derivative works.  However, this
//  document itself may not be modified in any way, such as by removing
//  the copyright notice or references to the Internet Society or other
//  Internet organizations, except as needed for the purpose of
//  developing Internet standards in which case the procedures for
//  copyrights defined in the Internet Standards process must be
//  followed, or as required to translate it into languages other than
//  English.
//
//  The limited permissions granted above are perpetual and will not be
//  revoked by the Internet Society or its successors or assigns.
//
//  This document and the information contained herein is provided on an
//  "AS IS" basis and THE INTERNET SOCIETY AND THE INTERNET ENGINEERING
//  TASK FORCE DISCLAIMS ALL WARRANTIES, EXPRESS OR IMPLIED, INCLUDING
//  BUT NOT LIMITED TO ANY WARRANTY THAT THE USE OF THE INFORMATION
//  HEREIN WILL NOT INFRINGE ANY RIGHTS OR ANY IMPLIED WARRANTIES OF
//  MERCHANTABILITY OR FITNESS FOR A PARTICULAR PURPOSE.

//////////////////////////////////////////////////////////////////////////
// Lexer rules generated for each distinct character in original grammar
// per http://www.unicode.org/charts/PDF/U0000.pdf
//////////////////////////////////////////////////////////////////////////

TAB : '\u0009';
SPACE : ' ';
EXCLAMATION : '!';
QUOTE : '"';
POUND : '#';
DOLLAR : '$';
PERCENT : '%';
AMPERSAND : '&';
APOSTROPHE : '\'';
LEFT_PAREN : '(';
RIGHT_PAREN : ')';
ASTERISK : '*';
PLUS : '+';
COMMA : ',';
DASH : '-';
PERIOD : '.';
SLASH : '/';
ZERO : '0';
ONE : '1';
TWO : '2';
THREE : '3';
FOUR : '4';
FIVE : '5';
SIX : '6';
SEVEN : '7';
EIGHT : '8';
NINE : '9';
COLON : ':';
SEMICOLON : ';';
LESS_THAN : '<';
EQUALS : '=';
GREATER_THAN : '>';
QUESTION : '?';
AT : '@';
CAP_A : 'A';
CAP_B : 'B';
CAP_C : 'C';
CAP_D : 'D';
CAP_E : 'E';
CAP_F : 'F';
CAP_G : 'G';
CAP_H : 'H';
CAP_I : 'I';
CAP_J : 'J';
CAP_K : 'K';
CAP_L : 'L';
CAP_M : 'M';
CAP_N : 'N';
CAP_O : 'O';
CAP_P : 'P';
CAP_Q : 'Q';
CAP_R : 'R';
CAP_S : 'S';
CAP_T : 'T';
CAP_U : 'U';
CAP_V : 'V';
CAP_W : 'W';
CAP_X : 'X';
CAP_Y : 'Y';
CAP_Z : 'Z';
LEFT_BRACE : '[';
BACKSLASH : '\\';
RIGHT_BRACE : ']';
CARAT : '^';
UNDERSCORE : '_';
ACCENT : '`';
A : 'a';
B : 'b';
C : 'c';
D : 'd';
E : 'e';
F : 'f';
G : 'g';
H : 'h';
I : 'i';
J : 'j';
K : 'k';
L : 'l';
M : 'm';
N : 'n';
O : 'o';
P : 'p';
Q : 'q';
R : 'r';
S : 's';
T : 't';
U : 'u';
V : 'v';
W : 'w';
X : 'x';
Y : 'y';
Z : 'z';
LEFT_CURLY_BRACE : '{';
PIPE : '|';
RIGHT_CURLY_BRACE : '}';
TILDE : '~';
U_0000 : '\u0000';
U_0001 : '\u0001';
U_0002 : '\u0002';
U_0003 : '\u0003';
U_0004 : '\u0004';
U_0005 : '\u0005';
U_0006 : '\u0006';
U_0007 : '\u0007';
U_0008 : '\u0008';
U_000B : '\u000B';
U_000C : '\u000C';
U_000E : '\u000E';
U_000F : '\u000F';
U_0010 : '\u0010';
U_0011 : '\u0011';
U_0012 : '\u0012';
U_0013 : '\u0013';
U_0014 : '\u0014';
U_0015 : '\u0015';
U_0016 : '\u0016';
U_0017 : '\u0017';
U_0018 : '\u0018';
U_0019 : '\u0019';
U_001A : '\u001A';
U_001B : '\u001B';
U_001C : '\u001C';
U_001D : '\u001D';
U_001E : '\u001E';
U_001F : '\u001F';
U_007F : '\u007F';
U_0080 : '\u0080';
U_0081 : '\u0081';
U_0082 : '\u0082';
U_0083 : '\u0083';
U_0084 : '\u0084';
U_0085 : '\u0085';
U_0086 : '\u0086';
U_0087 : '\u0087';
U_0088 : '\u0088';
U_0089 : '\u0089';
U_008A : '\u008A';
U_008B : '\u008B';
U_008C : '\u008C';
U_008D : '\u008D';
U_008E : '\u008E';
U_008F : '\u008F';
U_0090 : '\u0090';
U_0091 : '\u0091';
U_0092 : '\u0092';
U_0093 : '\u0093';
U_0094 : '\u0094';
U_0095 : '\u0095';
U_0096 : '\u0096';
U_0097 : '\u0097';
U_0098 : '\u0098';
U_0099 : '\u0099';
U_009A : '\u009A';
U_009B : '\u009B';
U_009C : '\u009C';
U_009D : '\u009D';
U_009E : '\u009E';
U_009F : '\u009F';
U_00A0 : '\u00A0';
U_00A1 : '\u00A1';
U_00A2 : '\u00A2';
U_00A3 : '\u00A3';
U_00A4 : '\u00A4';
U_00A5 : '\u00A5';
U_00A6 : '\u00A6';
U_00A7 : '\u00A7';
U_00A8 : '\u00A8';
U_00A9 : '\u00A9';
U_00AA : '\u00AA';
U_00AB : '\u00AB';
U_00AC : '\u00AC';
U_00AD : '\u00AD';
U_00AE : '\u00AE';
U_00AF : '\u00AF';
U_00B0 : '\u00B0';
U_00B1 : '\u00B1';
U_00B2 : '\u00B2';
U_00B3 : '\u00B3';
U_00B4 : '\u00B4';
U_00B5 : '\u00B5';
U_00B6 : '\u00B6';
U_00B7 : '\u00B7';
U_00B8 : '\u00B8';
U_00B9 : '\u00B9';
U_00BA : '\u00BA';
U_00BB : '\u00BB';
U_00BC : '\u00BC';
U_00BD : '\u00BD';
U_00BE : '\u00BE';
U_00BF : '\u00BF';
U_00C0 : '\u00C0';
U_00C1 : '\u00C1';
U_00C2 : '\u00C2';
U_00C3 : '\u00C3';
U_00C4 : '\u00C4';
U_00C5 : '\u00C5';
U_00C6 : '\u00C6';
U_00C7 : '\u00C7';
U_00C8 : '\u00C8';
U_00C9 : '\u00C9';
U_00CA : '\u00CA';
U_00CB : '\u00CB';
U_00CC : '\u00CC';
U_00CD : '\u00CD';
U_00CE : '\u00CE';
U_00CF : '\u00CF';
U_00D0 : '\u00D0';
U_00D1 : '\u00D1';
U_00D2 : '\u00D2';
U_00D3 : '\u00D3';
U_00D4 : '\u00D4';
U_00D5 : '\u00D5';
U_00D6 : '\u00D6';
U_00D7 : '\u00D7';
U_00D8 : '\u00D8';
U_00D9 : '\u00D9';
U_00DA : '\u00DA';
U_00DB : '\u00DB';
U_00DC : '\u00DC';
U_00DD : '\u00DD';
U_00DE : '\u00DE';
U_00DF : '\u00DF';
U_00E0 : '\u00E0';
U_00E1 : '\u00E1';
U_00E2 : '\u00E2';
U_00E3 : '\u00E3';
U_00E4 : '\u00E4';
U_00E5 : '\u00E5';
U_00E6 : '\u00E6';
U_00E7 : '\u00E7';
U_00E8 : '\u00E8';
U_00E9 : '\u00E9';
U_00EA : '\u00EA';
U_00EB : '\u00EB';
U_00EC : '\u00EC';
U_00ED : '\u00ED';
U_00EE : '\u00EE';
U_00EF : '\u00EF';
U_00F0 : '\u00F0';
U_00F1 : '\u00F1';
U_00F2 : '\u00F2';
U_00F3 : '\u00F3';
U_00F4 : '\u00F4';
U_00F5 : '\u00F5';
U_00F6 : '\u00F6';
U_00F7 : '\u00F7';
U_00F8 : '\u00F8';
U_00F9 : '\u00F9';
U_00FA : '\u00FA';
U_00FB : '\u00FB';
U_00FC : '\u00FC';
U_00FD : '\u00FD';
U_00FE : '\u00FE';
U_00FF : '\u00FF';
