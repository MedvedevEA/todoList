-- +goose Up
-- +goose StatementBegin
INSERT INTO public.tasks
VALUES 
(default,'task 1',null,'done',default,default),
(default,'task 2','50%','in_progress',default,default),
(default,'task 3','25%','in_progress',default,default),
(default,'task 4','waiting for task 2 to complete','new',default,default),
(default,'task 5','low priority','new',default,default),
(default,'task 6','low priority','new',default,default);
-- +goose StatementEnd
