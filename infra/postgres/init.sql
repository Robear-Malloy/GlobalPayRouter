-- GlobalRoute Pay — per-service schema bootstrap
-- Runs automatically on first postgres container start.
-- Each service owns exactly one schema; no cross-schema queries allowed.

CREATE SCHEMA IF NOT EXISTS routing;
CREATE SCHEMA IF NOT EXISTS fraud;
CREATE SCHEMA IF NOT EXISTS compliance;
CREATE SCHEMA IF NOT EXISTS ledger;

-- Grant the app user access to every schema
GRANT ALL PRIVILEGES ON SCHEMA routing    TO grp;
GRANT ALL PRIVILEGES ON SCHEMA fraud      TO grp;
GRANT ALL PRIVILEGES ON SCHEMA compliance TO grp;
GRANT ALL PRIVILEGES ON SCHEMA ledger     TO grp;
