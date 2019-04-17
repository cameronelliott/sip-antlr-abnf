// Code generated from Sip.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Sip

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSipListener is a complete listener for a parse tree produced by SipParser.
type BaseSipListener struct{}

var _ SipListener = &BaseSipListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSipListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSipListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSipListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSipListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterAlphanum is called when production alphanum is entered.
func (s *BaseSipListener) EnterAlphanum(ctx *AlphanumContext) {}

// ExitAlphanum is called when production alphanum is exited.
func (s *BaseSipListener) ExitAlphanum(ctx *AlphanumContext) {}

// EnterReserved is called when production reserved is entered.
func (s *BaseSipListener) EnterReserved(ctx *ReservedContext) {}

// ExitReserved is called when production reserved is exited.
func (s *BaseSipListener) ExitReserved(ctx *ReservedContext) {}

// EnterUnreserved is called when production unreserved is entered.
func (s *BaseSipListener) EnterUnreserved(ctx *UnreservedContext) {}

// ExitUnreserved is called when production unreserved is exited.
func (s *BaseSipListener) ExitUnreserved(ctx *UnreservedContext) {}

// EnterMark is called when production mark is entered.
func (s *BaseSipListener) EnterMark(ctx *MarkContext) {}

// ExitMark is called when production mark is exited.
func (s *BaseSipListener) ExitMark(ctx *MarkContext) {}

// EnterEscaped is called when production escaped is entered.
func (s *BaseSipListener) EnterEscaped(ctx *EscapedContext) {}

// ExitEscaped is called when production escaped is exited.
func (s *BaseSipListener) ExitEscaped(ctx *EscapedContext) {}

// EnterLws is called when production lws is entered.
func (s *BaseSipListener) EnterLws(ctx *LwsContext) {}

// ExitLws is called when production lws is exited.
func (s *BaseSipListener) ExitLws(ctx *LwsContext) {}

// EnterSws is called when production sws is entered.
func (s *BaseSipListener) EnterSws(ctx *SwsContext) {}

// ExitSws is called when production sws is exited.
func (s *BaseSipListener) ExitSws(ctx *SwsContext) {}

// EnterHcolon is called when production hcolon is entered.
func (s *BaseSipListener) EnterHcolon(ctx *HcolonContext) {}

// ExitHcolon is called when production hcolon is exited.
func (s *BaseSipListener) ExitHcolon(ctx *HcolonContext) {}

// EnterText_utf8_trim is called when production text_utf8_trim is entered.
func (s *BaseSipListener) EnterText_utf8_trim(ctx *Text_utf8_trimContext) {}

// ExitText_utf8_trim is called when production text_utf8_trim is exited.
func (s *BaseSipListener) ExitText_utf8_trim(ctx *Text_utf8_trimContext) {}

// EnterText_utf8char is called when production text_utf8char is entered.
func (s *BaseSipListener) EnterText_utf8char(ctx *Text_utf8charContext) {}

// ExitText_utf8char is called when production text_utf8char is exited.
func (s *BaseSipListener) ExitText_utf8char(ctx *Text_utf8charContext) {}

// EnterUtf8_nonascii is called when production utf8_nonascii is entered.
func (s *BaseSipListener) EnterUtf8_nonascii(ctx *Utf8_nonasciiContext) {}

// ExitUtf8_nonascii is called when production utf8_nonascii is exited.
func (s *BaseSipListener) ExitUtf8_nonascii(ctx *Utf8_nonasciiContext) {}

// EnterUtf8_cont is called when production utf8_cont is entered.
func (s *BaseSipListener) EnterUtf8_cont(ctx *Utf8_contContext) {}

// ExitUtf8_cont is called when production utf8_cont is exited.
func (s *BaseSipListener) ExitUtf8_cont(ctx *Utf8_contContext) {}

// EnterLhex is called when production lhex is entered.
func (s *BaseSipListener) EnterLhex(ctx *LhexContext) {}

// ExitLhex is called when production lhex is exited.
func (s *BaseSipListener) ExitLhex(ctx *LhexContext) {}

// EnterToken is called when production token is entered.
func (s *BaseSipListener) EnterToken(ctx *TokenContext) {}

// ExitToken is called when production token is exited.
func (s *BaseSipListener) ExitToken(ctx *TokenContext) {}

// EnterWord is called when production word is entered.
func (s *BaseSipListener) EnterWord(ctx *WordContext) {}

// ExitWord is called when production word is exited.
func (s *BaseSipListener) ExitWord(ctx *WordContext) {}

// EnterStar is called when production star is entered.
func (s *BaseSipListener) EnterStar(ctx *StarContext) {}

// ExitStar is called when production star is exited.
func (s *BaseSipListener) ExitStar(ctx *StarContext) {}

// EnterSlash is called when production slash is entered.
func (s *BaseSipListener) EnterSlash(ctx *SlashContext) {}

// ExitSlash is called when production slash is exited.
func (s *BaseSipListener) ExitSlash(ctx *SlashContext) {}

// EnterEqual is called when production equal is entered.
func (s *BaseSipListener) EnterEqual(ctx *EqualContext) {}

// ExitEqual is called when production equal is exited.
func (s *BaseSipListener) ExitEqual(ctx *EqualContext) {}

// EnterLparen is called when production lparen is entered.
func (s *BaseSipListener) EnterLparen(ctx *LparenContext) {}

// ExitLparen is called when production lparen is exited.
func (s *BaseSipListener) ExitLparen(ctx *LparenContext) {}

// EnterRparen is called when production rparen is entered.
func (s *BaseSipListener) EnterRparen(ctx *RparenContext) {}

// ExitRparen is called when production rparen is exited.
func (s *BaseSipListener) ExitRparen(ctx *RparenContext) {}

// EnterRaquot is called when production raquot is entered.
func (s *BaseSipListener) EnterRaquot(ctx *RaquotContext) {}

// ExitRaquot is called when production raquot is exited.
func (s *BaseSipListener) ExitRaquot(ctx *RaquotContext) {}

// EnterLaquot is called when production laquot is entered.
func (s *BaseSipListener) EnterLaquot(ctx *LaquotContext) {}

// ExitLaquot is called when production laquot is exited.
func (s *BaseSipListener) ExitLaquot(ctx *LaquotContext) {}

// EnterComma is called when production comma is entered.
func (s *BaseSipListener) EnterComma(ctx *CommaContext) {}

// ExitComma is called when production comma is exited.
func (s *BaseSipListener) ExitComma(ctx *CommaContext) {}

// EnterSemi is called when production semi is entered.
func (s *BaseSipListener) EnterSemi(ctx *SemiContext) {}

// ExitSemi is called when production semi is exited.
func (s *BaseSipListener) ExitSemi(ctx *SemiContext) {}

// EnterColon is called when production colon is entered.
func (s *BaseSipListener) EnterColon(ctx *ColonContext) {}

// ExitColon is called when production colon is exited.
func (s *BaseSipListener) ExitColon(ctx *ColonContext) {}

// EnterLdquot is called when production ldquot is entered.
func (s *BaseSipListener) EnterLdquot(ctx *LdquotContext) {}

// ExitLdquot is called when production ldquot is exited.
func (s *BaseSipListener) ExitLdquot(ctx *LdquotContext) {}

// EnterRdquot is called when production rdquot is entered.
func (s *BaseSipListener) EnterRdquot(ctx *RdquotContext) {}

// ExitRdquot is called when production rdquot is exited.
func (s *BaseSipListener) ExitRdquot(ctx *RdquotContext) {}

// EnterComment is called when production comment is entered.
func (s *BaseSipListener) EnterComment(ctx *CommentContext) {}

// ExitComment is called when production comment is exited.
func (s *BaseSipListener) ExitComment(ctx *CommentContext) {}

// EnterCtext is called when production ctext is entered.
func (s *BaseSipListener) EnterCtext(ctx *CtextContext) {}

// ExitCtext is called when production ctext is exited.
func (s *BaseSipListener) ExitCtext(ctx *CtextContext) {}

// EnterQuoted_string is called when production quoted_string is entered.
func (s *BaseSipListener) EnterQuoted_string(ctx *Quoted_stringContext) {}

// ExitQuoted_string is called when production quoted_string is exited.
func (s *BaseSipListener) ExitQuoted_string(ctx *Quoted_stringContext) {}

// EnterQdtext is called when production qdtext is entered.
func (s *BaseSipListener) EnterQdtext(ctx *QdtextContext) {}

// ExitQdtext is called when production qdtext is exited.
func (s *BaseSipListener) ExitQdtext(ctx *QdtextContext) {}

// EnterQuoted_pair is called when production quoted_pair is entered.
func (s *BaseSipListener) EnterQuoted_pair(ctx *Quoted_pairContext) {}

// ExitQuoted_pair is called when production quoted_pair is exited.
func (s *BaseSipListener) ExitQuoted_pair(ctx *Quoted_pairContext) {}

// EnterSip_uri is called when production sip_uri is entered.
func (s *BaseSipListener) EnterSip_uri(ctx *Sip_uriContext) {}

// ExitSip_uri is called when production sip_uri is exited.
func (s *BaseSipListener) ExitSip_uri(ctx *Sip_uriContext) {}

// EnterSips_uri is called when production sips_uri is entered.
func (s *BaseSipListener) EnterSips_uri(ctx *Sips_uriContext) {}

// ExitSips_uri is called when production sips_uri is exited.
func (s *BaseSipListener) ExitSips_uri(ctx *Sips_uriContext) {}

// EnterUserinfo is called when production userinfo is entered.
func (s *BaseSipListener) EnterUserinfo(ctx *UserinfoContext) {}

// ExitUserinfo is called when production userinfo is exited.
func (s *BaseSipListener) ExitUserinfo(ctx *UserinfoContext) {}

// EnterUser is called when production user is entered.
func (s *BaseSipListener) EnterUser(ctx *UserContext) {}

// ExitUser is called when production user is exited.
func (s *BaseSipListener) ExitUser(ctx *UserContext) {}

// EnterUser_unreserved is called when production user_unreserved is entered.
func (s *BaseSipListener) EnterUser_unreserved(ctx *User_unreservedContext) {}

// ExitUser_unreserved is called when production user_unreserved is exited.
func (s *BaseSipListener) ExitUser_unreserved(ctx *User_unreservedContext) {}

// EnterPassword is called when production password is entered.
func (s *BaseSipListener) EnterPassword(ctx *PasswordContext) {}

// ExitPassword is called when production password is exited.
func (s *BaseSipListener) ExitPassword(ctx *PasswordContext) {}

// EnterHostport is called when production hostport is entered.
func (s *BaseSipListener) EnterHostport(ctx *HostportContext) {}

// ExitHostport is called when production hostport is exited.
func (s *BaseSipListener) ExitHostport(ctx *HostportContext) {}

// EnterHost is called when production host is entered.
func (s *BaseSipListener) EnterHost(ctx *HostContext) {}

// ExitHost is called when production host is exited.
func (s *BaseSipListener) ExitHost(ctx *HostContext) {}

// EnterHostname is called when production hostname is entered.
func (s *BaseSipListener) EnterHostname(ctx *HostnameContext) {}

// ExitHostname is called when production hostname is exited.
func (s *BaseSipListener) ExitHostname(ctx *HostnameContext) {}

// EnterDomainlabel is called when production domainlabel is entered.
func (s *BaseSipListener) EnterDomainlabel(ctx *DomainlabelContext) {}

// ExitDomainlabel is called when production domainlabel is exited.
func (s *BaseSipListener) ExitDomainlabel(ctx *DomainlabelContext) {}

// EnterToplabel is called when production toplabel is entered.
func (s *BaseSipListener) EnterToplabel(ctx *ToplabelContext) {}

// ExitToplabel is called when production toplabel is exited.
func (s *BaseSipListener) ExitToplabel(ctx *ToplabelContext) {}

// EnterIpv4address is called when production ipv4address is entered.
func (s *BaseSipListener) EnterIpv4address(ctx *Ipv4addressContext) {}

// ExitIpv4address is called when production ipv4address is exited.
func (s *BaseSipListener) ExitIpv4address(ctx *Ipv4addressContext) {}

// EnterIpv6reference is called when production ipv6reference is entered.
func (s *BaseSipListener) EnterIpv6reference(ctx *Ipv6referenceContext) {}

// ExitIpv6reference is called when production ipv6reference is exited.
func (s *BaseSipListener) ExitIpv6reference(ctx *Ipv6referenceContext) {}

// EnterIpv6address is called when production ipv6address is entered.
func (s *BaseSipListener) EnterIpv6address(ctx *Ipv6addressContext) {}

// ExitIpv6address is called when production ipv6address is exited.
func (s *BaseSipListener) ExitIpv6address(ctx *Ipv6addressContext) {}

// EnterHexpart is called when production hexpart is entered.
func (s *BaseSipListener) EnterHexpart(ctx *HexpartContext) {}

// ExitHexpart is called when production hexpart is exited.
func (s *BaseSipListener) ExitHexpart(ctx *HexpartContext) {}

// EnterHexseq is called when production hexseq is entered.
func (s *BaseSipListener) EnterHexseq(ctx *HexseqContext) {}

// ExitHexseq is called when production hexseq is exited.
func (s *BaseSipListener) ExitHexseq(ctx *HexseqContext) {}

// EnterHex4 is called when production hex4 is entered.
func (s *BaseSipListener) EnterHex4(ctx *Hex4Context) {}

// ExitHex4 is called when production hex4 is exited.
func (s *BaseSipListener) ExitHex4(ctx *Hex4Context) {}

// EnterPort is called when production port is entered.
func (s *BaseSipListener) EnterPort(ctx *PortContext) {}

// ExitPort is called when production port is exited.
func (s *BaseSipListener) ExitPort(ctx *PortContext) {}

// EnterUri_parameters is called when production uri_parameters is entered.
func (s *BaseSipListener) EnterUri_parameters(ctx *Uri_parametersContext) {}

// ExitUri_parameters is called when production uri_parameters is exited.
func (s *BaseSipListener) ExitUri_parameters(ctx *Uri_parametersContext) {}

// EnterUri_parameter is called when production uri_parameter is entered.
func (s *BaseSipListener) EnterUri_parameter(ctx *Uri_parameterContext) {}

// ExitUri_parameter is called when production uri_parameter is exited.
func (s *BaseSipListener) ExitUri_parameter(ctx *Uri_parameterContext) {}

// EnterTransport_param is called when production transport_param is entered.
func (s *BaseSipListener) EnterTransport_param(ctx *Transport_paramContext) {}

// ExitTransport_param is called when production transport_param is exited.
func (s *BaseSipListener) ExitTransport_param(ctx *Transport_paramContext) {}

// EnterOther_transport is called when production other_transport is entered.
func (s *BaseSipListener) EnterOther_transport(ctx *Other_transportContext) {}

// ExitOther_transport is called when production other_transport is exited.
func (s *BaseSipListener) ExitOther_transport(ctx *Other_transportContext) {}

// EnterUser_param is called when production user_param is entered.
func (s *BaseSipListener) EnterUser_param(ctx *User_paramContext) {}

// ExitUser_param is called when production user_param is exited.
func (s *BaseSipListener) ExitUser_param(ctx *User_paramContext) {}

// EnterOther_user is called when production other_user is entered.
func (s *BaseSipListener) EnterOther_user(ctx *Other_userContext) {}

// ExitOther_user is called when production other_user is exited.
func (s *BaseSipListener) ExitOther_user(ctx *Other_userContext) {}

// EnterMethod_param is called when production method_param is entered.
func (s *BaseSipListener) EnterMethod_param(ctx *Method_paramContext) {}

// ExitMethod_param is called when production method_param is exited.
func (s *BaseSipListener) ExitMethod_param(ctx *Method_paramContext) {}

// EnterTtl_param is called when production ttl_param is entered.
func (s *BaseSipListener) EnterTtl_param(ctx *Ttl_paramContext) {}

// ExitTtl_param is called when production ttl_param is exited.
func (s *BaseSipListener) ExitTtl_param(ctx *Ttl_paramContext) {}

// EnterMaddr_param is called when production maddr_param is entered.
func (s *BaseSipListener) EnterMaddr_param(ctx *Maddr_paramContext) {}

// ExitMaddr_param is called when production maddr_param is exited.
func (s *BaseSipListener) ExitMaddr_param(ctx *Maddr_paramContext) {}

// EnterLr_param is called when production lr_param is entered.
func (s *BaseSipListener) EnterLr_param(ctx *Lr_paramContext) {}

// ExitLr_param is called when production lr_param is exited.
func (s *BaseSipListener) ExitLr_param(ctx *Lr_paramContext) {}

// EnterOther_param is called when production other_param is entered.
func (s *BaseSipListener) EnterOther_param(ctx *Other_paramContext) {}

// ExitOther_param is called when production other_param is exited.
func (s *BaseSipListener) ExitOther_param(ctx *Other_paramContext) {}

// EnterPname is called when production pname is entered.
func (s *BaseSipListener) EnterPname(ctx *PnameContext) {}

// ExitPname is called when production pname is exited.
func (s *BaseSipListener) ExitPname(ctx *PnameContext) {}

// EnterPvalue is called when production pvalue is entered.
func (s *BaseSipListener) EnterPvalue(ctx *PvalueContext) {}

// ExitPvalue is called when production pvalue is exited.
func (s *BaseSipListener) ExitPvalue(ctx *PvalueContext) {}

// EnterParamchar is called when production paramchar is entered.
func (s *BaseSipListener) EnterParamchar(ctx *ParamcharContext) {}

// ExitParamchar is called when production paramchar is exited.
func (s *BaseSipListener) ExitParamchar(ctx *ParamcharContext) {}

// EnterParam_unreserved is called when production param_unreserved is entered.
func (s *BaseSipListener) EnterParam_unreserved(ctx *Param_unreservedContext) {}

// ExitParam_unreserved is called when production param_unreserved is exited.
func (s *BaseSipListener) ExitParam_unreserved(ctx *Param_unreservedContext) {}

// EnterHeaders is called when production headers is entered.
func (s *BaseSipListener) EnterHeaders(ctx *HeadersContext) {}

// ExitHeaders is called when production headers is exited.
func (s *BaseSipListener) ExitHeaders(ctx *HeadersContext) {}

// EnterHeader is called when production header is entered.
func (s *BaseSipListener) EnterHeader(ctx *HeaderContext) {}

// ExitHeader is called when production header is exited.
func (s *BaseSipListener) ExitHeader(ctx *HeaderContext) {}

// EnterHname is called when production hname is entered.
func (s *BaseSipListener) EnterHname(ctx *HnameContext) {}

// ExitHname is called when production hname is exited.
func (s *BaseSipListener) ExitHname(ctx *HnameContext) {}

// EnterHvalue is called when production hvalue is entered.
func (s *BaseSipListener) EnterHvalue(ctx *HvalueContext) {}

// ExitHvalue is called when production hvalue is exited.
func (s *BaseSipListener) ExitHvalue(ctx *HvalueContext) {}

// EnterHnv_unreserved is called when production hnv_unreserved is entered.
func (s *BaseSipListener) EnterHnv_unreserved(ctx *Hnv_unreservedContext) {}

// ExitHnv_unreserved is called when production hnv_unreserved is exited.
func (s *BaseSipListener) ExitHnv_unreserved(ctx *Hnv_unreservedContext) {}

// EnterSip_message is called when production sip_message is entered.
func (s *BaseSipListener) EnterSip_message(ctx *Sip_messageContext) {}

// ExitSip_message is called when production sip_message is exited.
func (s *BaseSipListener) ExitSip_message(ctx *Sip_messageContext) {}

// EnterRequest is called when production request is entered.
func (s *BaseSipListener) EnterRequest(ctx *RequestContext) {}

// ExitRequest is called when production request is exited.
func (s *BaseSipListener) ExitRequest(ctx *RequestContext) {}

// EnterRequest_line is called when production request_line is entered.
func (s *BaseSipListener) EnterRequest_line(ctx *Request_lineContext) {}

// ExitRequest_line is called when production request_line is exited.
func (s *BaseSipListener) ExitRequest_line(ctx *Request_lineContext) {}

// EnterRequest_uri is called when production request_uri is entered.
func (s *BaseSipListener) EnterRequest_uri(ctx *Request_uriContext) {}

// ExitRequest_uri is called when production request_uri is exited.
func (s *BaseSipListener) ExitRequest_uri(ctx *Request_uriContext) {}

// EnterAbsoluteuri is called when production absoluteuri is entered.
func (s *BaseSipListener) EnterAbsoluteuri(ctx *AbsoluteuriContext) {}

// ExitAbsoluteuri is called when production absoluteuri is exited.
func (s *BaseSipListener) ExitAbsoluteuri(ctx *AbsoluteuriContext) {}

// EnterHier_part is called when production hier_part is entered.
func (s *BaseSipListener) EnterHier_part(ctx *Hier_partContext) {}

// ExitHier_part is called when production hier_part is exited.
func (s *BaseSipListener) ExitHier_part(ctx *Hier_partContext) {}

// EnterNet_path is called when production net_path is entered.
func (s *BaseSipListener) EnterNet_path(ctx *Net_pathContext) {}

// ExitNet_path is called when production net_path is exited.
func (s *BaseSipListener) ExitNet_path(ctx *Net_pathContext) {}

// EnterAbs_path is called when production abs_path is entered.
func (s *BaseSipListener) EnterAbs_path(ctx *Abs_pathContext) {}

// ExitAbs_path is called when production abs_path is exited.
func (s *BaseSipListener) ExitAbs_path(ctx *Abs_pathContext) {}

// EnterOpaque_part is called when production opaque_part is entered.
func (s *BaseSipListener) EnterOpaque_part(ctx *Opaque_partContext) {}

// ExitOpaque_part is called when production opaque_part is exited.
func (s *BaseSipListener) ExitOpaque_part(ctx *Opaque_partContext) {}

// EnterUric is called when production uric is entered.
func (s *BaseSipListener) EnterUric(ctx *UricContext) {}

// ExitUric is called when production uric is exited.
func (s *BaseSipListener) ExitUric(ctx *UricContext) {}

// EnterUric_no_slash is called when production uric_no_slash is entered.
func (s *BaseSipListener) EnterUric_no_slash(ctx *Uric_no_slashContext) {}

// ExitUric_no_slash is called when production uric_no_slash is exited.
func (s *BaseSipListener) ExitUric_no_slash(ctx *Uric_no_slashContext) {}

// EnterPath_segments is called when production path_segments is entered.
func (s *BaseSipListener) EnterPath_segments(ctx *Path_segmentsContext) {}

// ExitPath_segments is called when production path_segments is exited.
func (s *BaseSipListener) ExitPath_segments(ctx *Path_segmentsContext) {}

// EnterSegment is called when production segment is entered.
func (s *BaseSipListener) EnterSegment(ctx *SegmentContext) {}

// ExitSegment is called when production segment is exited.
func (s *BaseSipListener) ExitSegment(ctx *SegmentContext) {}

// EnterParam is called when production param is entered.
func (s *BaseSipListener) EnterParam(ctx *ParamContext) {}

// ExitParam is called when production param is exited.
func (s *BaseSipListener) ExitParam(ctx *ParamContext) {}

// EnterPchar is called when production pchar is entered.
func (s *BaseSipListener) EnterPchar(ctx *PcharContext) {}

// ExitPchar is called when production pchar is exited.
func (s *BaseSipListener) ExitPchar(ctx *PcharContext) {}

// EnterScheme is called when production scheme is entered.
func (s *BaseSipListener) EnterScheme(ctx *SchemeContext) {}

// ExitScheme is called when production scheme is exited.
func (s *BaseSipListener) ExitScheme(ctx *SchemeContext) {}

// EnterAuthority is called when production authority is entered.
func (s *BaseSipListener) EnterAuthority(ctx *AuthorityContext) {}

// ExitAuthority is called when production authority is exited.
func (s *BaseSipListener) ExitAuthority(ctx *AuthorityContext) {}

// EnterSrvr is called when production srvr is entered.
func (s *BaseSipListener) EnterSrvr(ctx *SrvrContext) {}

// ExitSrvr is called when production srvr is exited.
func (s *BaseSipListener) ExitSrvr(ctx *SrvrContext) {}

// EnterReg_name is called when production reg_name is entered.
func (s *BaseSipListener) EnterReg_name(ctx *Reg_nameContext) {}

// ExitReg_name is called when production reg_name is exited.
func (s *BaseSipListener) ExitReg_name(ctx *Reg_nameContext) {}

// EnterQuery is called when production query is entered.
func (s *BaseSipListener) EnterQuery(ctx *QueryContext) {}

// ExitQuery is called when production query is exited.
func (s *BaseSipListener) ExitQuery(ctx *QueryContext) {}

// EnterSip_version is called when production sip_version is entered.
func (s *BaseSipListener) EnterSip_version(ctx *Sip_versionContext) {}

// ExitSip_version is called when production sip_version is exited.
func (s *BaseSipListener) ExitSip_version(ctx *Sip_versionContext) {}

// EnterMessage_header is called when production message_header is entered.
func (s *BaseSipListener) EnterMessage_header(ctx *Message_headerContext) {}

// ExitMessage_header is called when production message_header is exited.
func (s *BaseSipListener) ExitMessage_header(ctx *Message_headerContext) {}

// EnterInvitem is called when production invitem is entered.
func (s *BaseSipListener) EnterInvitem(ctx *InvitemContext) {}

// ExitInvitem is called when production invitem is exited.
func (s *BaseSipListener) ExitInvitem(ctx *InvitemContext) {}

// EnterAckm is called when production ackm is entered.
func (s *BaseSipListener) EnterAckm(ctx *AckmContext) {}

// ExitAckm is called when production ackm is exited.
func (s *BaseSipListener) ExitAckm(ctx *AckmContext) {}

// EnterOptionsm is called when production optionsm is entered.
func (s *BaseSipListener) EnterOptionsm(ctx *OptionsmContext) {}

// ExitOptionsm is called when production optionsm is exited.
func (s *BaseSipListener) ExitOptionsm(ctx *OptionsmContext) {}

// EnterByem is called when production byem is entered.
func (s *BaseSipListener) EnterByem(ctx *ByemContext) {}

// ExitByem is called when production byem is exited.
func (s *BaseSipListener) ExitByem(ctx *ByemContext) {}

// EnterCancelm is called when production cancelm is entered.
func (s *BaseSipListener) EnterCancelm(ctx *CancelmContext) {}

// ExitCancelm is called when production cancelm is exited.
func (s *BaseSipListener) ExitCancelm(ctx *CancelmContext) {}

// EnterRegisterm is called when production registerm is entered.
func (s *BaseSipListener) EnterRegisterm(ctx *RegistermContext) {}

// ExitRegisterm is called when production registerm is exited.
func (s *BaseSipListener) ExitRegisterm(ctx *RegistermContext) {}

// EnterMethod is called when production method is entered.
func (s *BaseSipListener) EnterMethod(ctx *MethodContext) {}

// ExitMethod is called when production method is exited.
func (s *BaseSipListener) ExitMethod(ctx *MethodContext) {}

// EnterExtension_method is called when production extension_method is entered.
func (s *BaseSipListener) EnterExtension_method(ctx *Extension_methodContext) {}

// ExitExtension_method is called when production extension_method is exited.
func (s *BaseSipListener) ExitExtension_method(ctx *Extension_methodContext) {}

// EnterResponse is called when production response is entered.
func (s *BaseSipListener) EnterResponse(ctx *ResponseContext) {}

// ExitResponse is called when production response is exited.
func (s *BaseSipListener) ExitResponse(ctx *ResponseContext) {}

// EnterStatus_line is called when production status_line is entered.
func (s *BaseSipListener) EnterStatus_line(ctx *Status_lineContext) {}

// ExitStatus_line is called when production status_line is exited.
func (s *BaseSipListener) ExitStatus_line(ctx *Status_lineContext) {}

// EnterStatus_code is called when production status_code is entered.
func (s *BaseSipListener) EnterStatus_code(ctx *Status_codeContext) {}

// ExitStatus_code is called when production status_code is exited.
func (s *BaseSipListener) ExitStatus_code(ctx *Status_codeContext) {}

// EnterExtension_code is called when production extension_code is entered.
func (s *BaseSipListener) EnterExtension_code(ctx *Extension_codeContext) {}

// ExitExtension_code is called when production extension_code is exited.
func (s *BaseSipListener) ExitExtension_code(ctx *Extension_codeContext) {}

// EnterReason_phrase is called when production reason_phrase is entered.
func (s *BaseSipListener) EnterReason_phrase(ctx *Reason_phraseContext) {}

// ExitReason_phrase is called when production reason_phrase is exited.
func (s *BaseSipListener) ExitReason_phrase(ctx *Reason_phraseContext) {}

// EnterInformational is called when production informational is entered.
func (s *BaseSipListener) EnterInformational(ctx *InformationalContext) {}

// ExitInformational is called when production informational is exited.
func (s *BaseSipListener) ExitInformational(ctx *InformationalContext) {}

// EnterSuccess is called when production success is entered.
func (s *BaseSipListener) EnterSuccess(ctx *SuccessContext) {}

// ExitSuccess is called when production success is exited.
func (s *BaseSipListener) ExitSuccess(ctx *SuccessContext) {}

// EnterRedirection is called when production redirection is entered.
func (s *BaseSipListener) EnterRedirection(ctx *RedirectionContext) {}

// ExitRedirection is called when production redirection is exited.
func (s *BaseSipListener) ExitRedirection(ctx *RedirectionContext) {}

// EnterClient_error is called when production client_error is entered.
func (s *BaseSipListener) EnterClient_error(ctx *Client_errorContext) {}

// ExitClient_error is called when production client_error is exited.
func (s *BaseSipListener) ExitClient_error(ctx *Client_errorContext) {}

// EnterServer_error is called when production server_error is entered.
func (s *BaseSipListener) EnterServer_error(ctx *Server_errorContext) {}

// ExitServer_error is called when production server_error is exited.
func (s *BaseSipListener) ExitServer_error(ctx *Server_errorContext) {}

// EnterGlobal_failure is called when production global_failure is entered.
func (s *BaseSipListener) EnterGlobal_failure(ctx *Global_failureContext) {}

// ExitGlobal_failure is called when production global_failure is exited.
func (s *BaseSipListener) ExitGlobal_failure(ctx *Global_failureContext) {}

// EnterAccept is called when production accept is entered.
func (s *BaseSipListener) EnterAccept(ctx *AcceptContext) {}

// ExitAccept is called when production accept is exited.
func (s *BaseSipListener) ExitAccept(ctx *AcceptContext) {}

// EnterAccept_range is called when production accept_range is entered.
func (s *BaseSipListener) EnterAccept_range(ctx *Accept_rangeContext) {}

// ExitAccept_range is called when production accept_range is exited.
func (s *BaseSipListener) ExitAccept_range(ctx *Accept_rangeContext) {}

// EnterMedia_range is called when production media_range is entered.
func (s *BaseSipListener) EnterMedia_range(ctx *Media_rangeContext) {}

// ExitMedia_range is called when production media_range is exited.
func (s *BaseSipListener) ExitMedia_range(ctx *Media_rangeContext) {}

// EnterAccept_param is called when production accept_param is entered.
func (s *BaseSipListener) EnterAccept_param(ctx *Accept_paramContext) {}

// ExitAccept_param is called when production accept_param is exited.
func (s *BaseSipListener) ExitAccept_param(ctx *Accept_paramContext) {}

// EnterQvalue is called when production qvalue is entered.
func (s *BaseSipListener) EnterQvalue(ctx *QvalueContext) {}

// ExitQvalue is called when production qvalue is exited.
func (s *BaseSipListener) ExitQvalue(ctx *QvalueContext) {}

// EnterGeneric_param is called when production generic_param is entered.
func (s *BaseSipListener) EnterGeneric_param(ctx *Generic_paramContext) {}

// ExitGeneric_param is called when production generic_param is exited.
func (s *BaseSipListener) ExitGeneric_param(ctx *Generic_paramContext) {}

// EnterGen_value is called when production gen_value is entered.
func (s *BaseSipListener) EnterGen_value(ctx *Gen_valueContext) {}

// ExitGen_value is called when production gen_value is exited.
func (s *BaseSipListener) ExitGen_value(ctx *Gen_valueContext) {}

// EnterAccept_encoding is called when production accept_encoding is entered.
func (s *BaseSipListener) EnterAccept_encoding(ctx *Accept_encodingContext) {}

// ExitAccept_encoding is called when production accept_encoding is exited.
func (s *BaseSipListener) ExitAccept_encoding(ctx *Accept_encodingContext) {}

// EnterEncoding is called when production encoding is entered.
func (s *BaseSipListener) EnterEncoding(ctx *EncodingContext) {}

// ExitEncoding is called when production encoding is exited.
func (s *BaseSipListener) ExitEncoding(ctx *EncodingContext) {}

// EnterCodings is called when production codings is entered.
func (s *BaseSipListener) EnterCodings(ctx *CodingsContext) {}

// ExitCodings is called when production codings is exited.
func (s *BaseSipListener) ExitCodings(ctx *CodingsContext) {}

// EnterContent_coding is called when production content_coding is entered.
func (s *BaseSipListener) EnterContent_coding(ctx *Content_codingContext) {}

// ExitContent_coding is called when production content_coding is exited.
func (s *BaseSipListener) ExitContent_coding(ctx *Content_codingContext) {}

// EnterAccept_language is called when production accept_language is entered.
func (s *BaseSipListener) EnterAccept_language(ctx *Accept_languageContext) {}

// ExitAccept_language is called when production accept_language is exited.
func (s *BaseSipListener) ExitAccept_language(ctx *Accept_languageContext) {}

// EnterLanguage is called when production language is entered.
func (s *BaseSipListener) EnterLanguage(ctx *LanguageContext) {}

// ExitLanguage is called when production language is exited.
func (s *BaseSipListener) ExitLanguage(ctx *LanguageContext) {}

// EnterLanguage_range is called when production language_range is entered.
func (s *BaseSipListener) EnterLanguage_range(ctx *Language_rangeContext) {}

// ExitLanguage_range is called when production language_range is exited.
func (s *BaseSipListener) ExitLanguage_range(ctx *Language_rangeContext) {}

// EnterAlert_info is called when production alert_info is entered.
func (s *BaseSipListener) EnterAlert_info(ctx *Alert_infoContext) {}

// ExitAlert_info is called when production alert_info is exited.
func (s *BaseSipListener) ExitAlert_info(ctx *Alert_infoContext) {}

// EnterAlert_param is called when production alert_param is entered.
func (s *BaseSipListener) EnterAlert_param(ctx *Alert_paramContext) {}

// ExitAlert_param is called when production alert_param is exited.
func (s *BaseSipListener) ExitAlert_param(ctx *Alert_paramContext) {}

// EnterAllow is called when production allow is entered.
func (s *BaseSipListener) EnterAllow(ctx *AllowContext) {}

// ExitAllow is called when production allow is exited.
func (s *BaseSipListener) ExitAllow(ctx *AllowContext) {}

// EnterAuthorization is called when production authorization is entered.
func (s *BaseSipListener) EnterAuthorization(ctx *AuthorizationContext) {}

// ExitAuthorization is called when production authorization is exited.
func (s *BaseSipListener) ExitAuthorization(ctx *AuthorizationContext) {}

// EnterCredentials is called when production credentials is entered.
func (s *BaseSipListener) EnterCredentials(ctx *CredentialsContext) {}

// ExitCredentials is called when production credentials is exited.
func (s *BaseSipListener) ExitCredentials(ctx *CredentialsContext) {}

// EnterDigest_response is called when production digest_response is entered.
func (s *BaseSipListener) EnterDigest_response(ctx *Digest_responseContext) {}

// ExitDigest_response is called when production digest_response is exited.
func (s *BaseSipListener) ExitDigest_response(ctx *Digest_responseContext) {}

// EnterDig_resp is called when production dig_resp is entered.
func (s *BaseSipListener) EnterDig_resp(ctx *Dig_respContext) {}

// ExitDig_resp is called when production dig_resp is exited.
func (s *BaseSipListener) ExitDig_resp(ctx *Dig_respContext) {}

// EnterUsername is called when production username is entered.
func (s *BaseSipListener) EnterUsername(ctx *UsernameContext) {}

// ExitUsername is called when production username is exited.
func (s *BaseSipListener) ExitUsername(ctx *UsernameContext) {}

// EnterUsername_value is called when production username_value is entered.
func (s *BaseSipListener) EnterUsername_value(ctx *Username_valueContext) {}

// ExitUsername_value is called when production username_value is exited.
func (s *BaseSipListener) ExitUsername_value(ctx *Username_valueContext) {}

// EnterDigest_uri is called when production digest_uri is entered.
func (s *BaseSipListener) EnterDigest_uri(ctx *Digest_uriContext) {}

// ExitDigest_uri is called when production digest_uri is exited.
func (s *BaseSipListener) ExitDigest_uri(ctx *Digest_uriContext) {}

// EnterDigest_uri_value is called when production digest_uri_value is entered.
func (s *BaseSipListener) EnterDigest_uri_value(ctx *Digest_uri_valueContext) {}

// ExitDigest_uri_value is called when production digest_uri_value is exited.
func (s *BaseSipListener) ExitDigest_uri_value(ctx *Digest_uri_valueContext) {}

// EnterMessage_qop is called when production message_qop is entered.
func (s *BaseSipListener) EnterMessage_qop(ctx *Message_qopContext) {}

// ExitMessage_qop is called when production message_qop is exited.
func (s *BaseSipListener) ExitMessage_qop(ctx *Message_qopContext) {}

// EnterCnonce is called when production cnonce is entered.
func (s *BaseSipListener) EnterCnonce(ctx *CnonceContext) {}

// ExitCnonce is called when production cnonce is exited.
func (s *BaseSipListener) ExitCnonce(ctx *CnonceContext) {}

// EnterCnonce_value is called when production cnonce_value is entered.
func (s *BaseSipListener) EnterCnonce_value(ctx *Cnonce_valueContext) {}

// ExitCnonce_value is called when production cnonce_value is exited.
func (s *BaseSipListener) ExitCnonce_value(ctx *Cnonce_valueContext) {}

// EnterNonce_count is called when production nonce_count is entered.
func (s *BaseSipListener) EnterNonce_count(ctx *Nonce_countContext) {}

// ExitNonce_count is called when production nonce_count is exited.
func (s *BaseSipListener) ExitNonce_count(ctx *Nonce_countContext) {}

// EnterNc_value is called when production nc_value is entered.
func (s *BaseSipListener) EnterNc_value(ctx *Nc_valueContext) {}

// ExitNc_value is called when production nc_value is exited.
func (s *BaseSipListener) ExitNc_value(ctx *Nc_valueContext) {}

// EnterDresponse is called when production dresponse is entered.
func (s *BaseSipListener) EnterDresponse(ctx *DresponseContext) {}

// ExitDresponse is called when production dresponse is exited.
func (s *BaseSipListener) ExitDresponse(ctx *DresponseContext) {}

// EnterRequest_digest is called when production request_digest is entered.
func (s *BaseSipListener) EnterRequest_digest(ctx *Request_digestContext) {}

// ExitRequest_digest is called when production request_digest is exited.
func (s *BaseSipListener) ExitRequest_digest(ctx *Request_digestContext) {}

// EnterAuth_param is called when production auth_param is entered.
func (s *BaseSipListener) EnterAuth_param(ctx *Auth_paramContext) {}

// ExitAuth_param is called when production auth_param is exited.
func (s *BaseSipListener) ExitAuth_param(ctx *Auth_paramContext) {}

// EnterAuth_param_name is called when production auth_param_name is entered.
func (s *BaseSipListener) EnterAuth_param_name(ctx *Auth_param_nameContext) {}

// ExitAuth_param_name is called when production auth_param_name is exited.
func (s *BaseSipListener) ExitAuth_param_name(ctx *Auth_param_nameContext) {}

// EnterOther_response is called when production other_response is entered.
func (s *BaseSipListener) EnterOther_response(ctx *Other_responseContext) {}

// ExitOther_response is called when production other_response is exited.
func (s *BaseSipListener) ExitOther_response(ctx *Other_responseContext) {}

// EnterAuth_scheme is called when production auth_scheme is entered.
func (s *BaseSipListener) EnterAuth_scheme(ctx *Auth_schemeContext) {}

// ExitAuth_scheme is called when production auth_scheme is exited.
func (s *BaseSipListener) ExitAuth_scheme(ctx *Auth_schemeContext) {}

// EnterAuthentication_info is called when production authentication_info is entered.
func (s *BaseSipListener) EnterAuthentication_info(ctx *Authentication_infoContext) {}

// ExitAuthentication_info is called when production authentication_info is exited.
func (s *BaseSipListener) ExitAuthentication_info(ctx *Authentication_infoContext) {}

// EnterAinfo is called when production ainfo is entered.
func (s *BaseSipListener) EnterAinfo(ctx *AinfoContext) {}

// ExitAinfo is called when production ainfo is exited.
func (s *BaseSipListener) ExitAinfo(ctx *AinfoContext) {}

// EnterNextnonce is called when production nextnonce is entered.
func (s *BaseSipListener) EnterNextnonce(ctx *NextnonceContext) {}

// ExitNextnonce is called when production nextnonce is exited.
func (s *BaseSipListener) ExitNextnonce(ctx *NextnonceContext) {}

// EnterResponse_auth is called when production response_auth is entered.
func (s *BaseSipListener) EnterResponse_auth(ctx *Response_authContext) {}

// ExitResponse_auth is called when production response_auth is exited.
func (s *BaseSipListener) ExitResponse_auth(ctx *Response_authContext) {}

// EnterResponse_digest is called when production response_digest is entered.
func (s *BaseSipListener) EnterResponse_digest(ctx *Response_digestContext) {}

// ExitResponse_digest is called when production response_digest is exited.
func (s *BaseSipListener) ExitResponse_digest(ctx *Response_digestContext) {}

// EnterCall_id is called when production call_id is entered.
func (s *BaseSipListener) EnterCall_id(ctx *Call_idContext) {}

// ExitCall_id is called when production call_id is exited.
func (s *BaseSipListener) ExitCall_id(ctx *Call_idContext) {}

// EnterCallid is called when production callid is entered.
func (s *BaseSipListener) EnterCallid(ctx *CallidContext) {}

// ExitCallid is called when production callid is exited.
func (s *BaseSipListener) ExitCallid(ctx *CallidContext) {}

// EnterCall_info is called when production call_info is entered.
func (s *BaseSipListener) EnterCall_info(ctx *Call_infoContext) {}

// ExitCall_info is called when production call_info is exited.
func (s *BaseSipListener) ExitCall_info(ctx *Call_infoContext) {}

// EnterInfo is called when production info is entered.
func (s *BaseSipListener) EnterInfo(ctx *InfoContext) {}

// ExitInfo is called when production info is exited.
func (s *BaseSipListener) ExitInfo(ctx *InfoContext) {}

// EnterInfo_param is called when production info_param is entered.
func (s *BaseSipListener) EnterInfo_param(ctx *Info_paramContext) {}

// ExitInfo_param is called when production info_param is exited.
func (s *BaseSipListener) ExitInfo_param(ctx *Info_paramContext) {}

// EnterContact is called when production contact is entered.
func (s *BaseSipListener) EnterContact(ctx *ContactContext) {}

// ExitContact is called when production contact is exited.
func (s *BaseSipListener) ExitContact(ctx *ContactContext) {}

// EnterContact_param is called when production contact_param is entered.
func (s *BaseSipListener) EnterContact_param(ctx *Contact_paramContext) {}

// ExitContact_param is called when production contact_param is exited.
func (s *BaseSipListener) ExitContact_param(ctx *Contact_paramContext) {}

// EnterName_addr is called when production name_addr is entered.
func (s *BaseSipListener) EnterName_addr(ctx *Name_addrContext) {}

// ExitName_addr is called when production name_addr is exited.
func (s *BaseSipListener) ExitName_addr(ctx *Name_addrContext) {}

// EnterAddr_spec is called when production addr_spec is entered.
func (s *BaseSipListener) EnterAddr_spec(ctx *Addr_specContext) {}

// ExitAddr_spec is called when production addr_spec is exited.
func (s *BaseSipListener) ExitAddr_spec(ctx *Addr_specContext) {}

// EnterDisplay_name is called when production display_name is entered.
func (s *BaseSipListener) EnterDisplay_name(ctx *Display_nameContext) {}

// ExitDisplay_name is called when production display_name is exited.
func (s *BaseSipListener) ExitDisplay_name(ctx *Display_nameContext) {}

// EnterContact_params is called when production contact_params is entered.
func (s *BaseSipListener) EnterContact_params(ctx *Contact_paramsContext) {}

// ExitContact_params is called when production contact_params is exited.
func (s *BaseSipListener) ExitContact_params(ctx *Contact_paramsContext) {}

// EnterC_p_q is called when production c_p_q is entered.
func (s *BaseSipListener) EnterC_p_q(ctx *C_p_qContext) {}

// ExitC_p_q is called when production c_p_q is exited.
func (s *BaseSipListener) ExitC_p_q(ctx *C_p_qContext) {}

// EnterC_p_expires is called when production c_p_expires is entered.
func (s *BaseSipListener) EnterC_p_expires(ctx *C_p_expiresContext) {}

// ExitC_p_expires is called when production c_p_expires is exited.
func (s *BaseSipListener) ExitC_p_expires(ctx *C_p_expiresContext) {}

// EnterContact_extension is called when production contact_extension is entered.
func (s *BaseSipListener) EnterContact_extension(ctx *Contact_extensionContext) {}

// ExitContact_extension is called when production contact_extension is exited.
func (s *BaseSipListener) ExitContact_extension(ctx *Contact_extensionContext) {}

// EnterDelta_seconds is called when production delta_seconds is entered.
func (s *BaseSipListener) EnterDelta_seconds(ctx *Delta_secondsContext) {}

// ExitDelta_seconds is called when production delta_seconds is exited.
func (s *BaseSipListener) ExitDelta_seconds(ctx *Delta_secondsContext) {}

// EnterContent_disposition is called when production content_disposition is entered.
func (s *BaseSipListener) EnterContent_disposition(ctx *Content_dispositionContext) {}

// ExitContent_disposition is called when production content_disposition is exited.
func (s *BaseSipListener) ExitContent_disposition(ctx *Content_dispositionContext) {}

// EnterDisp_type is called when production disp_type is entered.
func (s *BaseSipListener) EnterDisp_type(ctx *Disp_typeContext) {}

// ExitDisp_type is called when production disp_type is exited.
func (s *BaseSipListener) ExitDisp_type(ctx *Disp_typeContext) {}

// EnterDisp_param is called when production disp_param is entered.
func (s *BaseSipListener) EnterDisp_param(ctx *Disp_paramContext) {}

// ExitDisp_param is called when production disp_param is exited.
func (s *BaseSipListener) ExitDisp_param(ctx *Disp_paramContext) {}

// EnterHandling_param is called when production handling_param is entered.
func (s *BaseSipListener) EnterHandling_param(ctx *Handling_paramContext) {}

// ExitHandling_param is called when production handling_param is exited.
func (s *BaseSipListener) ExitHandling_param(ctx *Handling_paramContext) {}

// EnterOther_handling is called when production other_handling is entered.
func (s *BaseSipListener) EnterOther_handling(ctx *Other_handlingContext) {}

// ExitOther_handling is called when production other_handling is exited.
func (s *BaseSipListener) ExitOther_handling(ctx *Other_handlingContext) {}

// EnterDisp_extension_token is called when production disp_extension_token is entered.
func (s *BaseSipListener) EnterDisp_extension_token(ctx *Disp_extension_tokenContext) {}

// ExitDisp_extension_token is called when production disp_extension_token is exited.
func (s *BaseSipListener) ExitDisp_extension_token(ctx *Disp_extension_tokenContext) {}

// EnterContent_encoding is called when production content_encoding is entered.
func (s *BaseSipListener) EnterContent_encoding(ctx *Content_encodingContext) {}

// ExitContent_encoding is called when production content_encoding is exited.
func (s *BaseSipListener) ExitContent_encoding(ctx *Content_encodingContext) {}

// EnterContent_language is called when production content_language is entered.
func (s *BaseSipListener) EnterContent_language(ctx *Content_languageContext) {}

// ExitContent_language is called when production content_language is exited.
func (s *BaseSipListener) ExitContent_language(ctx *Content_languageContext) {}

// EnterLanguage_tag is called when production language_tag is entered.
func (s *BaseSipListener) EnterLanguage_tag(ctx *Language_tagContext) {}

// ExitLanguage_tag is called when production language_tag is exited.
func (s *BaseSipListener) ExitLanguage_tag(ctx *Language_tagContext) {}

// EnterPrimary_tag is called when production primary_tag is entered.
func (s *BaseSipListener) EnterPrimary_tag(ctx *Primary_tagContext) {}

// ExitPrimary_tag is called when production primary_tag is exited.
func (s *BaseSipListener) ExitPrimary_tag(ctx *Primary_tagContext) {}

// EnterSubtag is called when production subtag is entered.
func (s *BaseSipListener) EnterSubtag(ctx *SubtagContext) {}

// ExitSubtag is called when production subtag is exited.
func (s *BaseSipListener) ExitSubtag(ctx *SubtagContext) {}

// EnterContent_length is called when production content_length is entered.
func (s *BaseSipListener) EnterContent_length(ctx *Content_lengthContext) {}

// ExitContent_length is called when production content_length is exited.
func (s *BaseSipListener) ExitContent_length(ctx *Content_lengthContext) {}

// EnterContent_type is called when production content_type is entered.
func (s *BaseSipListener) EnterContent_type(ctx *Content_typeContext) {}

// ExitContent_type is called when production content_type is exited.
func (s *BaseSipListener) ExitContent_type(ctx *Content_typeContext) {}

// EnterMedia_type is called when production media_type is entered.
func (s *BaseSipListener) EnterMedia_type(ctx *Media_typeContext) {}

// ExitMedia_type is called when production media_type is exited.
func (s *BaseSipListener) ExitMedia_type(ctx *Media_typeContext) {}

// EnterM_type is called when production m_type is entered.
func (s *BaseSipListener) EnterM_type(ctx *M_typeContext) {}

// ExitM_type is called when production m_type is exited.
func (s *BaseSipListener) ExitM_type(ctx *M_typeContext) {}

// EnterDiscrete_type is called when production discrete_type is entered.
func (s *BaseSipListener) EnterDiscrete_type(ctx *Discrete_typeContext) {}

// ExitDiscrete_type is called when production discrete_type is exited.
func (s *BaseSipListener) ExitDiscrete_type(ctx *Discrete_typeContext) {}

// EnterComposite_type is called when production composite_type is entered.
func (s *BaseSipListener) EnterComposite_type(ctx *Composite_typeContext) {}

// ExitComposite_type is called when production composite_type is exited.
func (s *BaseSipListener) ExitComposite_type(ctx *Composite_typeContext) {}

// EnterExtension_token is called when production extension_token is entered.
func (s *BaseSipListener) EnterExtension_token(ctx *Extension_tokenContext) {}

// ExitExtension_token is called when production extension_token is exited.
func (s *BaseSipListener) ExitExtension_token(ctx *Extension_tokenContext) {}

// EnterIetf_token is called when production ietf_token is entered.
func (s *BaseSipListener) EnterIetf_token(ctx *Ietf_tokenContext) {}

// ExitIetf_token is called when production ietf_token is exited.
func (s *BaseSipListener) ExitIetf_token(ctx *Ietf_tokenContext) {}

// EnterX_token is called when production x_token is entered.
func (s *BaseSipListener) EnterX_token(ctx *X_tokenContext) {}

// ExitX_token is called when production x_token is exited.
func (s *BaseSipListener) ExitX_token(ctx *X_tokenContext) {}

// EnterM_subtype is called when production m_subtype is entered.
func (s *BaseSipListener) EnterM_subtype(ctx *M_subtypeContext) {}

// ExitM_subtype is called when production m_subtype is exited.
func (s *BaseSipListener) ExitM_subtype(ctx *M_subtypeContext) {}

// EnterIana_token is called when production iana_token is entered.
func (s *BaseSipListener) EnterIana_token(ctx *Iana_tokenContext) {}

// ExitIana_token is called when production iana_token is exited.
func (s *BaseSipListener) ExitIana_token(ctx *Iana_tokenContext) {}

// EnterM_parameter is called when production m_parameter is entered.
func (s *BaseSipListener) EnterM_parameter(ctx *M_parameterContext) {}

// ExitM_parameter is called when production m_parameter is exited.
func (s *BaseSipListener) ExitM_parameter(ctx *M_parameterContext) {}

// EnterM_attribute is called when production m_attribute is entered.
func (s *BaseSipListener) EnterM_attribute(ctx *M_attributeContext) {}

// ExitM_attribute is called when production m_attribute is exited.
func (s *BaseSipListener) ExitM_attribute(ctx *M_attributeContext) {}

// EnterM_value is called when production m_value is entered.
func (s *BaseSipListener) EnterM_value(ctx *M_valueContext) {}

// ExitM_value is called when production m_value is exited.
func (s *BaseSipListener) ExitM_value(ctx *M_valueContext) {}

// EnterCseq is called when production cseq is entered.
func (s *BaseSipListener) EnterCseq(ctx *CseqContext) {}

// ExitCseq is called when production cseq is exited.
func (s *BaseSipListener) ExitCseq(ctx *CseqContext) {}

// EnterDate is called when production date is entered.
func (s *BaseSipListener) EnterDate(ctx *DateContext) {}

// ExitDate is called when production date is exited.
func (s *BaseSipListener) ExitDate(ctx *DateContext) {}

// EnterSip_date is called when production sip_date is entered.
func (s *BaseSipListener) EnterSip_date(ctx *Sip_dateContext) {}

// ExitSip_date is called when production sip_date is exited.
func (s *BaseSipListener) ExitSip_date(ctx *Sip_dateContext) {}

// EnterRfc1123_date is called when production rfc1123_date is entered.
func (s *BaseSipListener) EnterRfc1123_date(ctx *Rfc1123_dateContext) {}

// ExitRfc1123_date is called when production rfc1123_date is exited.
func (s *BaseSipListener) ExitRfc1123_date(ctx *Rfc1123_dateContext) {}

// EnterDate1 is called when production date1 is entered.
func (s *BaseSipListener) EnterDate1(ctx *Date1Context) {}

// ExitDate1 is called when production date1 is exited.
func (s *BaseSipListener) ExitDate1(ctx *Date1Context) {}

// EnterTime is called when production time is entered.
func (s *BaseSipListener) EnterTime(ctx *TimeContext) {}

// ExitTime is called when production time is exited.
func (s *BaseSipListener) ExitTime(ctx *TimeContext) {}

// EnterWkday is called when production wkday is entered.
func (s *BaseSipListener) EnterWkday(ctx *WkdayContext) {}

// ExitWkday is called when production wkday is exited.
func (s *BaseSipListener) ExitWkday(ctx *WkdayContext) {}

// EnterMonth is called when production month is entered.
func (s *BaseSipListener) EnterMonth(ctx *MonthContext) {}

// ExitMonth is called when production month is exited.
func (s *BaseSipListener) ExitMonth(ctx *MonthContext) {}

// EnterError_info is called when production error_info is entered.
func (s *BaseSipListener) EnterError_info(ctx *Error_infoContext) {}

// ExitError_info is called when production error_info is exited.
func (s *BaseSipListener) ExitError_info(ctx *Error_infoContext) {}

// EnterError_uri is called when production error_uri is entered.
func (s *BaseSipListener) EnterError_uri(ctx *Error_uriContext) {}

// ExitError_uri is called when production error_uri is exited.
func (s *BaseSipListener) ExitError_uri(ctx *Error_uriContext) {}

// EnterExpires is called when production expires is entered.
func (s *BaseSipListener) EnterExpires(ctx *ExpiresContext) {}

// ExitExpires is called when production expires is exited.
func (s *BaseSipListener) ExitExpires(ctx *ExpiresContext) {}

// EnterFrom is called when production from is entered.
func (s *BaseSipListener) EnterFrom(ctx *FromContext) {}

// ExitFrom is called when production from is exited.
func (s *BaseSipListener) ExitFrom(ctx *FromContext) {}

// EnterFrom_spec is called when production from_spec is entered.
func (s *BaseSipListener) EnterFrom_spec(ctx *From_specContext) {}

// ExitFrom_spec is called when production from_spec is exited.
func (s *BaseSipListener) ExitFrom_spec(ctx *From_specContext) {}

// EnterFrom_param is called when production from_param is entered.
func (s *BaseSipListener) EnterFrom_param(ctx *From_paramContext) {}

// ExitFrom_param is called when production from_param is exited.
func (s *BaseSipListener) ExitFrom_param(ctx *From_paramContext) {}

// EnterTag_param is called when production tag_param is entered.
func (s *BaseSipListener) EnterTag_param(ctx *Tag_paramContext) {}

// ExitTag_param is called when production tag_param is exited.
func (s *BaseSipListener) ExitTag_param(ctx *Tag_paramContext) {}

// EnterIn_reply_to is called when production in_reply_to is entered.
func (s *BaseSipListener) EnterIn_reply_to(ctx *In_reply_toContext) {}

// ExitIn_reply_to is called when production in_reply_to is exited.
func (s *BaseSipListener) ExitIn_reply_to(ctx *In_reply_toContext) {}

// EnterMax_forwards is called when production max_forwards is entered.
func (s *BaseSipListener) EnterMax_forwards(ctx *Max_forwardsContext) {}

// ExitMax_forwards is called when production max_forwards is exited.
func (s *BaseSipListener) ExitMax_forwards(ctx *Max_forwardsContext) {}

// EnterMime_version is called when production mime_version is entered.
func (s *BaseSipListener) EnterMime_version(ctx *Mime_versionContext) {}

// ExitMime_version is called when production mime_version is exited.
func (s *BaseSipListener) ExitMime_version(ctx *Mime_versionContext) {}

// EnterMin_expires is called when production min_expires is entered.
func (s *BaseSipListener) EnterMin_expires(ctx *Min_expiresContext) {}

// ExitMin_expires is called when production min_expires is exited.
func (s *BaseSipListener) ExitMin_expires(ctx *Min_expiresContext) {}

// EnterOrganization is called when production organization is entered.
func (s *BaseSipListener) EnterOrganization(ctx *OrganizationContext) {}

// ExitOrganization is called when production organization is exited.
func (s *BaseSipListener) ExitOrganization(ctx *OrganizationContext) {}

// EnterPriority is called when production priority is entered.
func (s *BaseSipListener) EnterPriority(ctx *PriorityContext) {}

// ExitPriority is called when production priority is exited.
func (s *BaseSipListener) ExitPriority(ctx *PriorityContext) {}

// EnterPriority_value is called when production priority_value is entered.
func (s *BaseSipListener) EnterPriority_value(ctx *Priority_valueContext) {}

// ExitPriority_value is called when production priority_value is exited.
func (s *BaseSipListener) ExitPriority_value(ctx *Priority_valueContext) {}

// EnterOther_priority is called when production other_priority is entered.
func (s *BaseSipListener) EnterOther_priority(ctx *Other_priorityContext) {}

// ExitOther_priority is called when production other_priority is exited.
func (s *BaseSipListener) ExitOther_priority(ctx *Other_priorityContext) {}

// EnterProxy_authenticate is called when production proxy_authenticate is entered.
func (s *BaseSipListener) EnterProxy_authenticate(ctx *Proxy_authenticateContext) {}

// ExitProxy_authenticate is called when production proxy_authenticate is exited.
func (s *BaseSipListener) ExitProxy_authenticate(ctx *Proxy_authenticateContext) {}

// EnterChallenge is called when production challenge is entered.
func (s *BaseSipListener) EnterChallenge(ctx *ChallengeContext) {}

// ExitChallenge is called when production challenge is exited.
func (s *BaseSipListener) ExitChallenge(ctx *ChallengeContext) {}

// EnterOther_challenge is called when production other_challenge is entered.
func (s *BaseSipListener) EnterOther_challenge(ctx *Other_challengeContext) {}

// ExitOther_challenge is called when production other_challenge is exited.
func (s *BaseSipListener) ExitOther_challenge(ctx *Other_challengeContext) {}

// EnterDigest_cln is called when production digest_cln is entered.
func (s *BaseSipListener) EnterDigest_cln(ctx *Digest_clnContext) {}

// ExitDigest_cln is called when production digest_cln is exited.
func (s *BaseSipListener) ExitDigest_cln(ctx *Digest_clnContext) {}

// EnterRealm is called when production realm is entered.
func (s *BaseSipListener) EnterRealm(ctx *RealmContext) {}

// ExitRealm is called when production realm is exited.
func (s *BaseSipListener) ExitRealm(ctx *RealmContext) {}

// EnterRealm_value is called when production realm_value is entered.
func (s *BaseSipListener) EnterRealm_value(ctx *Realm_valueContext) {}

// ExitRealm_value is called when production realm_value is exited.
func (s *BaseSipListener) ExitRealm_value(ctx *Realm_valueContext) {}

// EnterDomain is called when production domain is entered.
func (s *BaseSipListener) EnterDomain(ctx *DomainContext) {}

// ExitDomain is called when production domain is exited.
func (s *BaseSipListener) ExitDomain(ctx *DomainContext) {}

// EnterUri is called when production uri is entered.
func (s *BaseSipListener) EnterUri(ctx *UriContext) {}

// ExitUri is called when production uri is exited.
func (s *BaseSipListener) ExitUri(ctx *UriContext) {}

// EnterNonce is called when production nonce is entered.
func (s *BaseSipListener) EnterNonce(ctx *NonceContext) {}

// ExitNonce is called when production nonce is exited.
func (s *BaseSipListener) ExitNonce(ctx *NonceContext) {}

// EnterNonce_value is called when production nonce_value is entered.
func (s *BaseSipListener) EnterNonce_value(ctx *Nonce_valueContext) {}

// ExitNonce_value is called when production nonce_value is exited.
func (s *BaseSipListener) ExitNonce_value(ctx *Nonce_valueContext) {}

// EnterOpaque is called when production opaque is entered.
func (s *BaseSipListener) EnterOpaque(ctx *OpaqueContext) {}

// ExitOpaque is called when production opaque is exited.
func (s *BaseSipListener) ExitOpaque(ctx *OpaqueContext) {}

// EnterStale is called when production stale is entered.
func (s *BaseSipListener) EnterStale(ctx *StaleContext) {}

// ExitStale is called when production stale is exited.
func (s *BaseSipListener) ExitStale(ctx *StaleContext) {}

// EnterAlgorithm is called when production algorithm is entered.
func (s *BaseSipListener) EnterAlgorithm(ctx *AlgorithmContext) {}

// ExitAlgorithm is called when production algorithm is exited.
func (s *BaseSipListener) ExitAlgorithm(ctx *AlgorithmContext) {}

// EnterQop_options is called when production qop_options is entered.
func (s *BaseSipListener) EnterQop_options(ctx *Qop_optionsContext) {}

// ExitQop_options is called when production qop_options is exited.
func (s *BaseSipListener) ExitQop_options(ctx *Qop_optionsContext) {}

// EnterQop_value is called when production qop_value is entered.
func (s *BaseSipListener) EnterQop_value(ctx *Qop_valueContext) {}

// ExitQop_value is called when production qop_value is exited.
func (s *BaseSipListener) ExitQop_value(ctx *Qop_valueContext) {}

// EnterProxy_authorization is called when production proxy_authorization is entered.
func (s *BaseSipListener) EnterProxy_authorization(ctx *Proxy_authorizationContext) {}

// ExitProxy_authorization is called when production proxy_authorization is exited.
func (s *BaseSipListener) ExitProxy_authorization(ctx *Proxy_authorizationContext) {}

// EnterProxy_require is called when production proxy_require is entered.
func (s *BaseSipListener) EnterProxy_require(ctx *Proxy_requireContext) {}

// ExitProxy_require is called when production proxy_require is exited.
func (s *BaseSipListener) ExitProxy_require(ctx *Proxy_requireContext) {}

// EnterOption_tag is called when production option_tag is entered.
func (s *BaseSipListener) EnterOption_tag(ctx *Option_tagContext) {}

// ExitOption_tag is called when production option_tag is exited.
func (s *BaseSipListener) ExitOption_tag(ctx *Option_tagContext) {}

// EnterRecord_route is called when production record_route is entered.
func (s *BaseSipListener) EnterRecord_route(ctx *Record_routeContext) {}

// ExitRecord_route is called when production record_route is exited.
func (s *BaseSipListener) ExitRecord_route(ctx *Record_routeContext) {}

// EnterRec_route is called when production rec_route is entered.
func (s *BaseSipListener) EnterRec_route(ctx *Rec_routeContext) {}

// ExitRec_route is called when production rec_route is exited.
func (s *BaseSipListener) ExitRec_route(ctx *Rec_routeContext) {}

// EnterRr_param is called when production rr_param is entered.
func (s *BaseSipListener) EnterRr_param(ctx *Rr_paramContext) {}

// ExitRr_param is called when production rr_param is exited.
func (s *BaseSipListener) ExitRr_param(ctx *Rr_paramContext) {}

// EnterReply_to is called when production reply_to is entered.
func (s *BaseSipListener) EnterReply_to(ctx *Reply_toContext) {}

// ExitReply_to is called when production reply_to is exited.
func (s *BaseSipListener) ExitReply_to(ctx *Reply_toContext) {}

// EnterRplyto_spec is called when production rplyto_spec is entered.
func (s *BaseSipListener) EnterRplyto_spec(ctx *Rplyto_specContext) {}

// ExitRplyto_spec is called when production rplyto_spec is exited.
func (s *BaseSipListener) ExitRplyto_spec(ctx *Rplyto_specContext) {}

// EnterRplyto_param is called when production rplyto_param is entered.
func (s *BaseSipListener) EnterRplyto_param(ctx *Rplyto_paramContext) {}

// ExitRplyto_param is called when production rplyto_param is exited.
func (s *BaseSipListener) ExitRplyto_param(ctx *Rplyto_paramContext) {}

// EnterRequire is called when production require is entered.
func (s *BaseSipListener) EnterRequire(ctx *RequireContext) {}

// ExitRequire is called when production require is exited.
func (s *BaseSipListener) ExitRequire(ctx *RequireContext) {}

// EnterRetry_after is called when production retry_after is entered.
func (s *BaseSipListener) EnterRetry_after(ctx *Retry_afterContext) {}

// ExitRetry_after is called when production retry_after is exited.
func (s *BaseSipListener) ExitRetry_after(ctx *Retry_afterContext) {}

// EnterRetry_param is called when production retry_param is entered.
func (s *BaseSipListener) EnterRetry_param(ctx *Retry_paramContext) {}

// ExitRetry_param is called when production retry_param is exited.
func (s *BaseSipListener) ExitRetry_param(ctx *Retry_paramContext) {}

// EnterRoute is called when production route is entered.
func (s *BaseSipListener) EnterRoute(ctx *RouteContext) {}

// ExitRoute is called when production route is exited.
func (s *BaseSipListener) ExitRoute(ctx *RouteContext) {}

// EnterRoute_param is called when production route_param is entered.
func (s *BaseSipListener) EnterRoute_param(ctx *Route_paramContext) {}

// ExitRoute_param is called when production route_param is exited.
func (s *BaseSipListener) ExitRoute_param(ctx *Route_paramContext) {}

// EnterServer is called when production server is entered.
func (s *BaseSipListener) EnterServer(ctx *ServerContext) {}

// ExitServer is called when production server is exited.
func (s *BaseSipListener) ExitServer(ctx *ServerContext) {}

// EnterServer_val is called when production server_val is entered.
func (s *BaseSipListener) EnterServer_val(ctx *Server_valContext) {}

// ExitServer_val is called when production server_val is exited.
func (s *BaseSipListener) ExitServer_val(ctx *Server_valContext) {}

// EnterProduct is called when production product is entered.
func (s *BaseSipListener) EnterProduct(ctx *ProductContext) {}

// ExitProduct is called when production product is exited.
func (s *BaseSipListener) ExitProduct(ctx *ProductContext) {}

// EnterProduct_version is called when production product_version is entered.
func (s *BaseSipListener) EnterProduct_version(ctx *Product_versionContext) {}

// ExitProduct_version is called when production product_version is exited.
func (s *BaseSipListener) ExitProduct_version(ctx *Product_versionContext) {}

// EnterSubject is called when production subject is entered.
func (s *BaseSipListener) EnterSubject(ctx *SubjectContext) {}

// ExitSubject is called when production subject is exited.
func (s *BaseSipListener) ExitSubject(ctx *SubjectContext) {}

// EnterSupported is called when production supported is entered.
func (s *BaseSipListener) EnterSupported(ctx *SupportedContext) {}

// ExitSupported is called when production supported is exited.
func (s *BaseSipListener) ExitSupported(ctx *SupportedContext) {}

// EnterTimestamp is called when production timestamp is entered.
func (s *BaseSipListener) EnterTimestamp(ctx *TimestampContext) {}

// ExitTimestamp is called when production timestamp is exited.
func (s *BaseSipListener) ExitTimestamp(ctx *TimestampContext) {}

// EnterDelay is called when production delay is entered.
func (s *BaseSipListener) EnterDelay(ctx *DelayContext) {}

// ExitDelay is called when production delay is exited.
func (s *BaseSipListener) ExitDelay(ctx *DelayContext) {}

// EnterTo is called when production to is entered.
func (s *BaseSipListener) EnterTo(ctx *ToContext) {}

// ExitTo is called when production to is exited.
func (s *BaseSipListener) ExitTo(ctx *ToContext) {}

// EnterTo_param is called when production to_param is entered.
func (s *BaseSipListener) EnterTo_param(ctx *To_paramContext) {}

// ExitTo_param is called when production to_param is exited.
func (s *BaseSipListener) ExitTo_param(ctx *To_paramContext) {}

// EnterUnsupported is called when production unsupported is entered.
func (s *BaseSipListener) EnterUnsupported(ctx *UnsupportedContext) {}

// ExitUnsupported is called when production unsupported is exited.
func (s *BaseSipListener) ExitUnsupported(ctx *UnsupportedContext) {}

// EnterUser_agent is called when production user_agent is entered.
func (s *BaseSipListener) EnterUser_agent(ctx *User_agentContext) {}

// ExitUser_agent is called when production user_agent is exited.
func (s *BaseSipListener) ExitUser_agent(ctx *User_agentContext) {}

// EnterVia is called when production via is entered.
func (s *BaseSipListener) EnterVia(ctx *ViaContext) {}

// ExitVia is called when production via is exited.
func (s *BaseSipListener) ExitVia(ctx *ViaContext) {}

// EnterVia_parm is called when production via_parm is entered.
func (s *BaseSipListener) EnterVia_parm(ctx *Via_parmContext) {}

// ExitVia_parm is called when production via_parm is exited.
func (s *BaseSipListener) ExitVia_parm(ctx *Via_parmContext) {}

// EnterVia_params is called when production via_params is entered.
func (s *BaseSipListener) EnterVia_params(ctx *Via_paramsContext) {}

// ExitVia_params is called when production via_params is exited.
func (s *BaseSipListener) ExitVia_params(ctx *Via_paramsContext) {}

// EnterVia_ttl is called when production via_ttl is entered.
func (s *BaseSipListener) EnterVia_ttl(ctx *Via_ttlContext) {}

// ExitVia_ttl is called when production via_ttl is exited.
func (s *BaseSipListener) ExitVia_ttl(ctx *Via_ttlContext) {}

// EnterVia_maddr is called when production via_maddr is entered.
func (s *BaseSipListener) EnterVia_maddr(ctx *Via_maddrContext) {}

// ExitVia_maddr is called when production via_maddr is exited.
func (s *BaseSipListener) ExitVia_maddr(ctx *Via_maddrContext) {}

// EnterVia_received is called when production via_received is entered.
func (s *BaseSipListener) EnterVia_received(ctx *Via_receivedContext) {}

// ExitVia_received is called when production via_received is exited.
func (s *BaseSipListener) ExitVia_received(ctx *Via_receivedContext) {}

// EnterVia_branch is called when production via_branch is entered.
func (s *BaseSipListener) EnterVia_branch(ctx *Via_branchContext) {}

// ExitVia_branch is called when production via_branch is exited.
func (s *BaseSipListener) ExitVia_branch(ctx *Via_branchContext) {}

// EnterVia_extension is called when production via_extension is entered.
func (s *BaseSipListener) EnterVia_extension(ctx *Via_extensionContext) {}

// ExitVia_extension is called when production via_extension is exited.
func (s *BaseSipListener) ExitVia_extension(ctx *Via_extensionContext) {}

// EnterSent_protocol is called when production sent_protocol is entered.
func (s *BaseSipListener) EnterSent_protocol(ctx *Sent_protocolContext) {}

// ExitSent_protocol is called when production sent_protocol is exited.
func (s *BaseSipListener) ExitSent_protocol(ctx *Sent_protocolContext) {}

// EnterProtocol_name is called when production protocol_name is entered.
func (s *BaseSipListener) EnterProtocol_name(ctx *Protocol_nameContext) {}

// ExitProtocol_name is called when production protocol_name is exited.
func (s *BaseSipListener) ExitProtocol_name(ctx *Protocol_nameContext) {}

// EnterProtocol_version is called when production protocol_version is entered.
func (s *BaseSipListener) EnterProtocol_version(ctx *Protocol_versionContext) {}

// ExitProtocol_version is called when production protocol_version is exited.
func (s *BaseSipListener) ExitProtocol_version(ctx *Protocol_versionContext) {}

// EnterTransport is called when production transport is entered.
func (s *BaseSipListener) EnterTransport(ctx *TransportContext) {}

// ExitTransport is called when production transport is exited.
func (s *BaseSipListener) ExitTransport(ctx *TransportContext) {}

// EnterSent_by is called when production sent_by is entered.
func (s *BaseSipListener) EnterSent_by(ctx *Sent_byContext) {}

// ExitSent_by is called when production sent_by is exited.
func (s *BaseSipListener) ExitSent_by(ctx *Sent_byContext) {}

// EnterTtl is called when production ttl is entered.
func (s *BaseSipListener) EnterTtl(ctx *TtlContext) {}

// ExitTtl is called when production ttl is exited.
func (s *BaseSipListener) ExitTtl(ctx *TtlContext) {}

// EnterWarning is called when production warning is entered.
func (s *BaseSipListener) EnterWarning(ctx *WarningContext) {}

// ExitWarning is called when production warning is exited.
func (s *BaseSipListener) ExitWarning(ctx *WarningContext) {}

// EnterWarning_value is called when production warning_value is entered.
func (s *BaseSipListener) EnterWarning_value(ctx *Warning_valueContext) {}

// ExitWarning_value is called when production warning_value is exited.
func (s *BaseSipListener) ExitWarning_value(ctx *Warning_valueContext) {}

// EnterWarn_code is called when production warn_code is entered.
func (s *BaseSipListener) EnterWarn_code(ctx *Warn_codeContext) {}

// ExitWarn_code is called when production warn_code is exited.
func (s *BaseSipListener) ExitWarn_code(ctx *Warn_codeContext) {}

// EnterWarn_agent is called when production warn_agent is entered.
func (s *BaseSipListener) EnterWarn_agent(ctx *Warn_agentContext) {}

// ExitWarn_agent is called when production warn_agent is exited.
func (s *BaseSipListener) ExitWarn_agent(ctx *Warn_agentContext) {}

// EnterWarn_text is called when production warn_text is entered.
func (s *BaseSipListener) EnterWarn_text(ctx *Warn_textContext) {}

// ExitWarn_text is called when production warn_text is exited.
func (s *BaseSipListener) ExitWarn_text(ctx *Warn_textContext) {}

// EnterPseudonym is called when production pseudonym is entered.
func (s *BaseSipListener) EnterPseudonym(ctx *PseudonymContext) {}

// ExitPseudonym is called when production pseudonym is exited.
func (s *BaseSipListener) ExitPseudonym(ctx *PseudonymContext) {}

// EnterWww_authenticate is called when production www_authenticate is entered.
func (s *BaseSipListener) EnterWww_authenticate(ctx *Www_authenticateContext) {}

// ExitWww_authenticate is called when production www_authenticate is exited.
func (s *BaseSipListener) ExitWww_authenticate(ctx *Www_authenticateContext) {}

// EnterExtension_header is called when production extension_header is entered.
func (s *BaseSipListener) EnterExtension_header(ctx *Extension_headerContext) {}

// ExitExtension_header is called when production extension_header is exited.
func (s *BaseSipListener) ExitExtension_header(ctx *Extension_headerContext) {}

// EnterHeader_name is called when production header_name is entered.
func (s *BaseSipListener) EnterHeader_name(ctx *Header_nameContext) {}

// ExitHeader_name is called when production header_name is exited.
func (s *BaseSipListener) ExitHeader_name(ctx *Header_nameContext) {}

// EnterHeader_value is called when production header_value is entered.
func (s *BaseSipListener) EnterHeader_value(ctx *Header_valueContext) {}

// ExitHeader_value is called when production header_value is exited.
func (s *BaseSipListener) ExitHeader_value(ctx *Header_valueContext) {}

// EnterMessage_body is called when production message_body is entered.
func (s *BaseSipListener) EnterMessage_body(ctx *Message_bodyContext) {}

// ExitMessage_body is called when production message_body is exited.
func (s *BaseSipListener) ExitMessage_body(ctx *Message_bodyContext) {}

// EnterTelephone_subscriber is called when production telephone_subscriber is entered.
func (s *BaseSipListener) EnterTelephone_subscriber(ctx *Telephone_subscriberContext) {}

// ExitTelephone_subscriber is called when production telephone_subscriber is exited.
func (s *BaseSipListener) ExitTelephone_subscriber(ctx *Telephone_subscriberContext) {}

// EnterGlobal_phone_number is called when production global_phone_number is entered.
func (s *BaseSipListener) EnterGlobal_phone_number(ctx *Global_phone_numberContext) {}

// ExitGlobal_phone_number is called when production global_phone_number is exited.
func (s *BaseSipListener) ExitGlobal_phone_number(ctx *Global_phone_numberContext) {}

// EnterBase_phone_number is called when production base_phone_number is entered.
func (s *BaseSipListener) EnterBase_phone_number(ctx *Base_phone_numberContext) {}

// ExitBase_phone_number is called when production base_phone_number is exited.
func (s *BaseSipListener) ExitBase_phone_number(ctx *Base_phone_numberContext) {}

// EnterLocal_phone_number is called when production local_phone_number is entered.
func (s *BaseSipListener) EnterLocal_phone_number(ctx *Local_phone_numberContext) {}

// ExitLocal_phone_number is called when production local_phone_number is exited.
func (s *BaseSipListener) ExitLocal_phone_number(ctx *Local_phone_numberContext) {}

// EnterIsdn_subaddress is called when production isdn_subaddress is entered.
func (s *BaseSipListener) EnterIsdn_subaddress(ctx *Isdn_subaddressContext) {}

// ExitIsdn_subaddress is called when production isdn_subaddress is exited.
func (s *BaseSipListener) ExitIsdn_subaddress(ctx *Isdn_subaddressContext) {}

// EnterPost_dial is called when production post_dial is entered.
func (s *BaseSipListener) EnterPost_dial(ctx *Post_dialContext) {}

// ExitPost_dial is called when production post_dial is exited.
func (s *BaseSipListener) ExitPost_dial(ctx *Post_dialContext) {}

// EnterArea_specifier is called when production area_specifier is entered.
func (s *BaseSipListener) EnterArea_specifier(ctx *Area_specifierContext) {}

// ExitArea_specifier is called when production area_specifier is exited.
func (s *BaseSipListener) ExitArea_specifier(ctx *Area_specifierContext) {}

// EnterPhone_context_tag is called when production phone_context_tag is entered.
func (s *BaseSipListener) EnterPhone_context_tag(ctx *Phone_context_tagContext) {}

// ExitPhone_context_tag is called when production phone_context_tag is exited.
func (s *BaseSipListener) ExitPhone_context_tag(ctx *Phone_context_tagContext) {}

// EnterPhone_context_ident is called when production phone_context_ident is entered.
func (s *BaseSipListener) EnterPhone_context_ident(ctx *Phone_context_identContext) {}

// ExitPhone_context_ident is called when production phone_context_ident is exited.
func (s *BaseSipListener) ExitPhone_context_ident(ctx *Phone_context_identContext) {}

// EnterNetwork_prefix is called when production network_prefix is entered.
func (s *BaseSipListener) EnterNetwork_prefix(ctx *Network_prefixContext) {}

// ExitNetwork_prefix is called when production network_prefix is exited.
func (s *BaseSipListener) ExitNetwork_prefix(ctx *Network_prefixContext) {}

// EnterGlobal_network_prefix is called when production global_network_prefix is entered.
func (s *BaseSipListener) EnterGlobal_network_prefix(ctx *Global_network_prefixContext) {}

// ExitGlobal_network_prefix is called when production global_network_prefix is exited.
func (s *BaseSipListener) ExitGlobal_network_prefix(ctx *Global_network_prefixContext) {}

// EnterLocal_network_prefix is called when production local_network_prefix is entered.
func (s *BaseSipListener) EnterLocal_network_prefix(ctx *Local_network_prefixContext) {}

// ExitLocal_network_prefix is called when production local_network_prefix is exited.
func (s *BaseSipListener) ExitLocal_network_prefix(ctx *Local_network_prefixContext) {}

// EnterPrivate_prefix is called when production private_prefix is entered.
func (s *BaseSipListener) EnterPrivate_prefix(ctx *Private_prefixContext) {}

// ExitPrivate_prefix is called when production private_prefix is exited.
func (s *BaseSipListener) ExitPrivate_prefix(ctx *Private_prefixContext) {}

// EnterService_provider is called when production service_provider is entered.
func (s *BaseSipListener) EnterService_provider(ctx *Service_providerContext) {}

// ExitService_provider is called when production service_provider is exited.
func (s *BaseSipListener) ExitService_provider(ctx *Service_providerContext) {}

// EnterProvider_tag is called when production provider_tag is entered.
func (s *BaseSipListener) EnterProvider_tag(ctx *Provider_tagContext) {}

// ExitProvider_tag is called when production provider_tag is exited.
func (s *BaseSipListener) ExitProvider_tag(ctx *Provider_tagContext) {}

// EnterProvider_hostname is called when production provider_hostname is entered.
func (s *BaseSipListener) EnterProvider_hostname(ctx *Provider_hostnameContext) {}

// ExitProvider_hostname is called when production provider_hostname is exited.
func (s *BaseSipListener) ExitProvider_hostname(ctx *Provider_hostnameContext) {}

// EnterFuture_extension is called when production future_extension is entered.
func (s *BaseSipListener) EnterFuture_extension(ctx *Future_extensionContext) {}

// ExitFuture_extension is called when production future_extension is exited.
func (s *BaseSipListener) ExitFuture_extension(ctx *Future_extensionContext) {}

// EnterToken_char is called when production token_char is entered.
func (s *BaseSipListener) EnterToken_char(ctx *Token_charContext) {}

// ExitToken_char is called when production token_char is exited.
func (s *BaseSipListener) ExitToken_char(ctx *Token_charContext) {}

// EnterQuoted_string_telnum is called when production quoted_string_telnum is entered.
func (s *BaseSipListener) EnterQuoted_string_telnum(ctx *Quoted_string_telnumContext) {}

// ExitQuoted_string_telnum is called when production quoted_string_telnum is exited.
func (s *BaseSipListener) ExitQuoted_string_telnum(ctx *Quoted_string_telnumContext) {}

// EnterPhonedigit is called when production phonedigit is entered.
func (s *BaseSipListener) EnterPhonedigit(ctx *PhonedigitContext) {}

// ExitPhonedigit is called when production phonedigit is exited.
func (s *BaseSipListener) ExitPhonedigit(ctx *PhonedigitContext) {}

// EnterVisual_separator is called when production visual_separator is entered.
func (s *BaseSipListener) EnterVisual_separator(ctx *Visual_separatorContext) {}

// ExitVisual_separator is called when production visual_separator is exited.
func (s *BaseSipListener) ExitVisual_separator(ctx *Visual_separatorContext) {}

// EnterPause_character is called when production pause_character is entered.
func (s *BaseSipListener) EnterPause_character(ctx *Pause_characterContext) {}

// ExitPause_character is called when production pause_character is exited.
func (s *BaseSipListener) ExitPause_character(ctx *Pause_characterContext) {}

// EnterOne_second_pause is called when production one_second_pause is entered.
func (s *BaseSipListener) EnterOne_second_pause(ctx *One_second_pauseContext) {}

// ExitOne_second_pause is called when production one_second_pause is exited.
func (s *BaseSipListener) ExitOne_second_pause(ctx *One_second_pauseContext) {}

// EnterWait_for_dial_tone is called when production wait_for_dial_tone is entered.
func (s *BaseSipListener) EnterWait_for_dial_tone(ctx *Wait_for_dial_toneContext) {}

// ExitWait_for_dial_tone is called when production wait_for_dial_tone is exited.
func (s *BaseSipListener) ExitWait_for_dial_tone(ctx *Wait_for_dial_toneContext) {}

// EnterDtmf_digit is called when production dtmf_digit is entered.
func (s *BaseSipListener) EnterDtmf_digit(ctx *Dtmf_digitContext) {}

// ExitDtmf_digit is called when production dtmf_digit is exited.
func (s *BaseSipListener) ExitDtmf_digit(ctx *Dtmf_digitContext) {}

// EnterAlpha is called when production alpha is entered.
func (s *BaseSipListener) EnterAlpha(ctx *AlphaContext) {}

// ExitAlpha is called when production alpha is exited.
func (s *BaseSipListener) ExitAlpha(ctx *AlphaContext) {}

// EnterBit is called when production bit is entered.
func (s *BaseSipListener) EnterBit(ctx *BitContext) {}

// ExitBit is called when production bit is exited.
func (s *BaseSipListener) ExitBit(ctx *BitContext) {}

// EnterChar_1 is called when production char_1 is entered.
func (s *BaseSipListener) EnterChar_1(ctx *Char_1Context) {}

// ExitChar_1 is called when production char_1 is exited.
func (s *BaseSipListener) ExitChar_1(ctx *Char_1Context) {}

// EnterCr is called when production cr is entered.
func (s *BaseSipListener) EnterCr(ctx *CrContext) {}

// ExitCr is called when production cr is exited.
func (s *BaseSipListener) ExitCr(ctx *CrContext) {}

// EnterCrlf is called when production crlf is entered.
func (s *BaseSipListener) EnterCrlf(ctx *CrlfContext) {}

// ExitCrlf is called when production crlf is exited.
func (s *BaseSipListener) ExitCrlf(ctx *CrlfContext) {}

// EnterCtl is called when production ctl is entered.
func (s *BaseSipListener) EnterCtl(ctx *CtlContext) {}

// ExitCtl is called when production ctl is exited.
func (s *BaseSipListener) ExitCtl(ctx *CtlContext) {}

// EnterDigit is called when production digit is entered.
func (s *BaseSipListener) EnterDigit(ctx *DigitContext) {}

// ExitDigit is called when production digit is exited.
func (s *BaseSipListener) ExitDigit(ctx *DigitContext) {}

// EnterDquote is called when production dquote is entered.
func (s *BaseSipListener) EnterDquote(ctx *DquoteContext) {}

// ExitDquote is called when production dquote is exited.
func (s *BaseSipListener) ExitDquote(ctx *DquoteContext) {}

// EnterHexdig is called when production hexdig is entered.
func (s *BaseSipListener) EnterHexdig(ctx *HexdigContext) {}

// ExitHexdig is called when production hexdig is exited.
func (s *BaseSipListener) ExitHexdig(ctx *HexdigContext) {}

// EnterHtab is called when production htab is entered.
func (s *BaseSipListener) EnterHtab(ctx *HtabContext) {}

// ExitHtab is called when production htab is exited.
func (s *BaseSipListener) ExitHtab(ctx *HtabContext) {}

// EnterLf is called when production lf is entered.
func (s *BaseSipListener) EnterLf(ctx *LfContext) {}

// ExitLf is called when production lf is exited.
func (s *BaseSipListener) ExitLf(ctx *LfContext) {}

// EnterLwsp is called when production lwsp is entered.
func (s *BaseSipListener) EnterLwsp(ctx *LwspContext) {}

// ExitLwsp is called when production lwsp is exited.
func (s *BaseSipListener) ExitLwsp(ctx *LwspContext) {}

// EnterOctet is called when production octet is entered.
func (s *BaseSipListener) EnterOctet(ctx *OctetContext) {}

// ExitOctet is called when production octet is exited.
func (s *BaseSipListener) ExitOctet(ctx *OctetContext) {}

// EnterSp is called when production sp is entered.
func (s *BaseSipListener) EnterSp(ctx *SpContext) {}

// ExitSp is called when production sp is exited.
func (s *BaseSipListener) ExitSp(ctx *SpContext) {}

// EnterVchar is called when production vchar is entered.
func (s *BaseSipListener) EnterVchar(ctx *VcharContext) {}

// ExitVchar is called when production vchar is exited.
func (s *BaseSipListener) ExitVchar(ctx *VcharContext) {}

// EnterWsp is called when production wsp is entered.
func (s *BaseSipListener) EnterWsp(ctx *WspContext) {}

// ExitWsp is called when production wsp is exited.
func (s *BaseSipListener) ExitWsp(ctx *WspContext) {}
