-- USERS
GRANT SELECT, INSERT, UPDATE, DELETE ON TABLE public.users TO c2n_admin;
GRANT USAGE, SELECT ON SEQUENCE public.users_id_seq TO c2n_admin;

GRANT SELECT, INSERT, UPDATE, DELETE ON TABLE public.users TO c2n_user;
GRANT USAGE, SELECT ON SEQUENCE public.users_id_seq TO c2n_user;


-- CATEGORIES
GRANT SELECT, INSERT, UPDATE, DELETE ON TABLE public.categories TO c2n_admin;
GRANT USAGE, SELECT ON SEQUENCE public.categories_id_seq TO c2n_admin;

GRANT SELECT ON TABLE public.categories TO c2n_user;


-- PRODUCTS
GRANT SELECT, INSERT, UPDATE, DELETE ON TABLE public.products TO c2n_admin;
GRANT USAGE, SELECT ON SEQUENCE public.products_id_seq TO c2n_admin;

GRANT SELECT ON TABLE public.products TO c2n_user;


-- REVIEWS
GRANT SELECT, INSERT, UPDATE, DELETE ON TABLE public.reviews TO c2n_admin;
GRANT USAGE, SELECT ON SEQUENCE public.reviews_id_seq TO c2n_admin;

GRANT SELECT, INSERT, UPDATE, DELETE ON TABLE public.reviews TO c2n_user;
GRANT USAGE, SELECT ON SEQUENCE public.reviews_id_seq TO c2n_user;