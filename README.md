# go-EmailChecker
Verify if the sender email domain is authentic 

**Summary of Differences**

Record	Purpose	                                                      Stored As	    Example
SPF	    Prevents email spoofing by listing authorized mail servers.	  TXT	          v=spf1 include:_spf.google.com ~all
DMARC	  Defines how to handle emails that fail SPF/DKIM checks.	      TXT           (under _dmarc.domain)	v=DMARC1; p=reject; rua=mailto:dmarc-reports@example.com
MX	    Directs incoming emails to the correct mail servers.	        MX	          10 mail.example.com.

**MX (Mail Exchange) Record**

Purpose:
MX records specify the mail servers responsible for receiving emails for a domain.

Format:
MX records point to mail server hostnames and include priority values.

Example MX Records for example.com:

example.com.    3600   IN   MX   10 mail.example.com.
example.com.    3600   IN   MX   20 backup.example.com.
10 mail.example.com. → Primary mail server (lower number = higher priority).
20 backup.example.com. → Backup mail server (used if primary is unavailable).
How it works:

When someone sends an email to user@example.com, the sender's mail server queries the MX records for example.com.
The email is delivered to the highest-priority (lowest number) mail server.
If the primary server is down, the backup server is used.


**1. SPF (Sender Policy Framework) Record**

Purpose:
SPF records define which mail servers are authorized to send emails on behalf of a domain. This helps prevent email spoofing (when attackers send emails pretending to be from your domain).

Format:
SPF records are stored as TXT records in DNS and usually start with v=spf1, followed by a list of allowed mail servers.

Example SPF Record:

v=spf1 include:_spf.google.com ~all
v=spf1 → SPF version 1.
include:_spf.google.com → Allows Google servers to send emails for this domain.
~all → Soft fail (unauthorized emails may be accepted but marked as spam).
How it works:

When an email is received, the recipient's mail server checks the SPF record of the sender’s domain.
If the sending mail server matches an allowed IP, the email is accepted.
If not, it may be rejected or marked as spam.

**2. DMARC (Domain-based Message Authentication, Reporting, and Conformance) Record**

Purpose:
DMARC builds on SPF and DKIM (DomainKeys Identified Mail) to tell email servers how to handle unauthenticated emails from a domain. It also allows reporting of email authentication results.

Format:
DMARC records are stored as TXT records under _dmarc.domain.com and usually start with v=DMARC1.

Example DMARC Record:

v=DMARC1; p=reject; rua=mailto:dmarc-reports@example.com
v=DMARC1 → DMARC version 1.
p=reject → Reject emails that fail authentication.
rua=mailto:dmarc-reports@example.com → Send reports to this email.
How it works:

The recipient mail server checks if the email passes SPF or DKIM.
If SPF/DKIM fails, DMARC tells the recipient what to do:
none → Take no action.
quarantine → Mark as spam.
reject → Block the email.
A report is sent to the domain owner about email authentication failures.
