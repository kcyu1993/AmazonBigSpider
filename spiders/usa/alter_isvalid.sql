# Setup the corresponding category for the SQL command.

# BigPID for the corresponding results

USE smart_base;
# List the results here
SELECT distinct url,id,bigpid ,name,bigpname,page
  FROM smart_category where isvalid=1 order by bigpid limit 1000000;

# Set the corresponding
UPDATE smart_category SET isvalid=1 WHERE bigpid=34;
UPDATE smart_category SET isvalid=0 WHERE NOT bigpid=34;
SELECT @original_isvalid := isvalid FROM smart_category where bigpid=34 ORDER BY bigpid LIMIT 100000;

SELECT distinct url,id,bigpid,name,bigpname,page FROM smart_category where isvalid=1 ORDER BY name limit 1000000;



# Setup the distinct
# SELECT @original_isvalid := isvalid FROM smart_category where bigpid=34 ORDER BY bigpid LIMIT 100000;

# SET @original_isvalid := 1;
#
# SELECT @original_isvalid;



