# GolangWebApplicationFirewall
Test repository for Web Application Firewall code experiments in Golang

int retVal = msc_intervention( transaction, &intervention );

what sets this retVal?

deny.cc			evaluate()
drop.cc			evaluate()
redirect.cc		evaluate()

~/ModSecurity/src/actions/disruptive
transaction->m_it.disruptive = true;

transaction object has ModSecurityIntervention property
ModSecurityIntervention object has disruptive  property int

so what code calls e.g.
deny.cc
evaluate()

msc_append_request_body
msc_append_response_body

e.g.
msc_append_response_body
this->m_rules->evaluate(modsecurity::RequestBodyPhase, this);

which calls deny.cc			evaluate()
which sets  transaction->m_it.disruptive = true;


then by the time msc_intervention() is called this evalates to true
bool Transaction::intervention(ModSecurityIntervention *it) {

thus the code con'ts and sets
it->disruptive = m_it.disruptive;

which is the disruptive flag of the ModSecurityIntervention object passed in to the function 
and the msc_intervention() ends as

return it->disruptive;




How does a malicious request get detected?
believe everything channelled here:
int inter = msc_intervention( transaction, &intervention );

and checks intervention struct
ModSecurityIntervention
disruptive

what functions set "disruptive" flag?
transaction.cc
appendRequestBody
appendResponseBody

wrappers?
msc_append_request_body
msc_append_response_body

deny.cc			evaluate()
drop.cc			evaluate()
redirect.cc		evaluate()


IMPORTANT
log sql injection attack details from here

rule_with_actions.cc
void RuleWithActions::performLogging

modsecurity.cc
void ModSecurity::serverLog(void *data, std::shared_ptr<RuleMessage> rm) {

rule_message.cc
std::string RuleMessage::log(const RuleMessage *rm, int props, int code) {
std::string RuleMessage::_details(const RuleMessage *rm) {


LOGGING
transaction.cc
extern "C" int msc_process_request_body( Transaction *transaction )
int Transaction::processRequestBody()

this->m_rules->evaluate(modsecurity::RequestBodyPhase, this);


IMPOSSIBLE
having to hack ModSecurity code base directly now and insert log statements!

StevePro RulesSet::evaluate() 2
StevePro RulesSet::evaluate() 3

phase 3


Violate
942100
detected SQLi using libinjection


serverlog
REQUEST-942-APPLICATION-ATTACK-SQLI.conf

line "45"
line "1102"
line "1188"
line "1570"


    "id:942100,\
    phase:2,\

    "id:942480,\
    phase:2,\



RuleMessage::log()
msg.append("[client " + std::string(*rm->m_clientIpAddress.get()) + "] ");
msg.append("ModSecurity: Warning. ");
msg.append("ModSecurity: Warning. ");
msg.append(rm->m_match);
msg.append(_details(rm));



ENCOURAGING
I found this meesage
i.e. m_match
m_match_message.assign("detected SQLi using libinjection.");

in operators
detect_sqi.h


Do I need to modify the SecRule to something like this?
SecRule REQUEST_BODY  "blocktest" "id:2,phase:2,deny,status:400,msg:'Test rule'"


Web Application Firewall in Go
OWASP ModSecurity Core RuleSet
https://medium.com/lightbaseio/web-application-firewall-in-go-feat-owasp-modsecurity-core-rule-set-3f97a26e3311