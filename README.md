Can I write the tests in cucumber & ensure it all works OK?

# Roster

## TODO
Figure out architecture.
 * Can I use the filesystem instead of a RDBMS?
   * much simpler; 1 box, no concurrent writes.
   * sqlite. Backup to dropbox as .CSV is a nice-to-have.

Roster manager user flow:
1) Visit site, enter http auth to get in to admin area.
2) Fill in members file, sessions file, hit OK
  * server saves members.csv & sessions.csv to disk under uploads/timestamp folder
  * redirects to review page for that uploaded timestamp
3) Check over review page (sample email, sample session)
  * Server shows a calendar to show that sessions CSV is OK.
  * Server shows a table of people to show that people CSV is OK.
4) Approve schedule
  * Server backs up old schedule
  * Server writes new schedule
  * Generates email for each member address, sends them

User gets email, follows link
  * Sees a calendar & a number of hours owed
  * pick dates to fill in the hours
  * hit save

Server updates schedule with who is doing what.
Show user an 'Add to calendar' link to an .ics feed.
Server emails user to tell them.

5) Admin checks back in ~2 weeks
 * See a report of which members haven't filled in their duty days
 * See a report of which days are still needed
 * Auto-assign members to duty days (algorithm? far apart? close together?)

Cronjob runs <sometimes> and invokes an RPC on the servers debug port.
 * Server checks whether anything is coming in 5..7 days
  * Updates CSV with 'email_reminder_sent_date' column
  * Sends email
  * Updates CSV with 'sms_reminder_sent_date' column
  * Sends SMS







### DB (?)
 * Member
  * email/phone/name/address/??? - json?
 * Admin
  * email/??? (member subtype? annex table?)

### Admin section
 * Login (one-time code?)
 * Secure Sessions (not in goweb, gorilla has one)
 * Import member CSV
 * show member list
 * make member an admin

#### Rostering
 * Setup rostered times & roles in advance
  * in this case, key/support/support
 * Send members an email with login code.
 * Set members rostered times.

### Member section
 * Choose my rostered times

hmac+sha email address + expiry date to create login code?
