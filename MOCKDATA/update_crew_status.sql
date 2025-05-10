-- Set crew status randomly
UPDATE crew 
SET status = (
    SELECT status 
    FROM (
        VALUES 
            ('active'),
            ('on_leave'),
            ('inactive'),
            ('suspended'),
            ('retired')
    ) AS s(status)
    ORDER BY RANDOM()
    LIMIT 1
);