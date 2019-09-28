# TabeoTest
Screening test for Tabeo


Here is my solution, it is nearly working, i started bottom-up with the tests, I ran into some import difficulty
and I am unsure how to fix it ..


possible improvements:
figure out way of writing top level test consistently with a realistic example
perhaps by providing a getschedule endpoint?
don't assume database exists
wrap errors and return more precise HTTP error codes  by make custom error type corresponding to failure modes
no time to do getall endpoint, although it seems trivial

use prepared statements in sql queries
move to Gorm for better time handling
refactor launchpadid to be consistent