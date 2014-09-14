Feature: Setup a new season
  In order to have volunteers show up at the toy library
  As the rostering coordinator
  I want to have a roster arranged

  Scenario: Start a new season
    When I log in to the admin section
     And I upload the member & session CSV files
    Then I should see the review page
    When I confirm the season looks good
    Then user emails should be sent out
     And I should see an empty roster page

    When 'John Smith' follows his login link
     And 'John Smith' selects his duty sessions

    Then I should see 'John Smith's sessions filled in on the roster page
    When I auto-fill the roster
    Then I should see 'Joshua Footle's sessions filled in on the roster page
