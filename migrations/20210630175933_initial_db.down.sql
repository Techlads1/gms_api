DROP TABLE IF EXISTS public.regions; -- drop region table
DROP TABLE IF EXISTS public.users;
-- drop types, eg enums if exisits
DROP TYPE IF EXISTS public.appeal_decision_enum;

-- drop functions
DROP FUNCTION IF EXISTS public.update_updated_at_column();

-- drop sequencies
DROP SEQUENCE IF EXISTS public.regions_id_seq;
