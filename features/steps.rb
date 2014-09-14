require 'bundler'
Bundler.require(:all)
require 'capybara'
require 'selenium-webdriver'
require File.dirname(__FILE__) + '/spec_helper'

include Capybara::DSL
Capybara.default_driver = :selenium

def base_url
  ENV['BASE_URL'] || "http://localhost:8080"
end

def as_admin(url)
  url.gsub(/:\/\//, "://#{username}:#{password}@")
end

def username
  ENV['ADMIN_USER'] || "admin"
end

def password
  ENV['ADMIN_PASS'] || "admin"
end

When(/^I log in to the admin section$/) do
  visit as_admin base_url
end

When(/^I upload the member & session CSV files$/) do
  attach_file("members.csv", File.expand_path("assets/sample_members.csv"))
  attach_file("sessions.csv", File.expand_path("assets/sample_roster.csv"))
end

Then(/^I should see the review page$/) do
  # I should see the members
  within(".members") do
    members = page.all(".member")
    expect(members).to have(2).members

    smiths = page.all(".member").first
    expect(smiths.find('memberNo').text).to eq "1"
    expect(smiths.find('name').text).to eq "John & Mary Smith"
    expect(smiths.find('mobiles').text).to eq "0412131920 1049853948"
    expect(smiths.find('emails').text).to eq "john@smith.ca mary@smith.ca"
    expect(smiths.find('hours_owed').text).to eq "6"

    footle = page.all(".member").last
    expect(smiths.find('memberNo').text).to eq "2"
    expect(smiths.find('name').text).to eq "Joshua Footle"
    expect(smiths.find('mobiles').text).to eq "004939402"
    expect(smiths.find('emails').text).to eq "josh@footle.com"
    expect(smiths.find('hours_owed').text).to eq "8"
  end

  # I should see a calendar
  within(".calendar") do
    duties = page.all(" .duty")
    expect(duties).to have(7).duty_sessions

    # Dates & times are correct
    times = duty_sessions.map {|s| Time.at s.find(".datetime")['data-timestamp'].to_i }
    expect(times[0]).to eq Time.new(2015, 01, 02, 19, 30, 0, "+11:00")
    expect(times[1]).to eq Time.new(2015, 01, 02, 19, 30, 0, "+11:00")
    expect(times[2]).to eq Time.new(2015, 01, 05, 10, 0, 0, "+11:00")
    expect(times[3]).to eq Time.new(2015, 01, 05, 10, 0, 0, "+11:00")
    expect(times[4]).to eq Time.new(2015, 01, 05, 10, 0, 0, "+11:00")
    expect(times[5]).to eq Time.new(2015, 01, 09, 19, 30, 0, "+11:00")
    expect(times[6]).to eq Time.new(2015, 01, 09, 19, 30, 0, "+11:00")

    # durations
    durations = duty_sessions.map {|s| s.find(".duration").text }
    expect(durations[0]).to eq "2 hours"
    expect(durations[1]).to eq "2 hours"
    expect(durations[2]).to eq "3 hours"
    expect(durations[3]).to eq "3 hours"
    expect(durations[4]).to eq "3 hours"
    expect(durations[5]).to eq "2 hours"
    expect(durations[6]).to eq "2 hours"

    # durations
    locations = duty_sessions.map {|s| s.find(".location").text }
    expect(locations[0]).to eq "Oakleigh community hall"
    expect(locations[1]).to eq "Oakleigh community hall"
    expect(locations[2]).to eq "Oakleigh community hall"
    expect(locations[3]).to eq "Oakleigh community hall"
    expect(locations[4]).to eq "Oakleigh community hall"
    expect(locations[5]).to eq "Oakleigh community hall"
    expect(locations[6]).to eq "Oakleigh community hall"

    # role
    locations = duty_sessions.map {|s| s.find(".role").text }
    expect(locations[0]).to eq "Key"
    expect(locations[1]).to eq "Second"
    expect(locations[2]).to eq "Key"
    expect(locations[3]).to eq "Second"
    expect(locations[4]).to eq "Second"
    expect(locations[5]).to eq "Key"
    expect(locations[6]).to eq "Second"

    # filled yet?
    locations = duty_sessions.map {|s| s.find(".member").text }
    expect(locations[0]).to eq ""
    expect(locations[1]).to eq ""
    expect(locations[2]).to eq "John & Mary"
    expect(locations[3]).to eq ""
    expect(locations[4]).to eq ""
    expect(locations[5]).to eq ""
    expect(locations[6]).to eq ""

  end
end

When(/^I confirm the season looks good$/) do
  click_on ".confirm_roster_preview"
end

Then(/^user emails should be sent out$/) do
  pending # express the regexp above with the code you wish you had
end

Then(/^I should see an empty roster page$/) do
  pending # express the regexp above with the code you wish you had
end

When(/^'[^']+' follows his login link$/) do
  pending # express the regexp above with the code you wish you had
end

When(/^'[^']+' selects his duty sessions$/) do
  pending # express the regexp above with the code you wish you had
end

Then(/^I should see '[^']+'s sessions filled in on the roster page$/) do
  pending # express the regexp above with the code you wish you had
end

When(/^I auto\-fill the roster$/) do
  pending # express the regexp above with the code you wish you had
end
