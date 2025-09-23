-- === USERS ===
INSERT INTO users (login, password, is_moderator)
VALUES
    ('demo','demo',true),
    ('moder','moder',true),
    ('alice','alice',false),
    ('bob','bob',false)
    ON CONFLICT (login) DO NOTHING;


-- === CRITERIA с правильными image_url ===
INSERT INTO criteria (code, name, indicator, duration, home_visit, image_url, description)
VALUES
    ('№1','Оценка возраста пациента','> 55 лет','1 календарный день', true,
     'http://127.0.0.1:9000/services-images/n1_age.png',
     'Возраст > 55 лет — критерий при поступлении.'),

    ('№2','Анализ лейкоцитов крови','> 16 000/мм³','1 календарный день', true,
     'http://127.0.0.1:9000/services-images/n2_wbc.png',
     'Повышение лейкоцитов может указывать на выраженный воспалительный процесс.'),

    ('№3','Измерение уровня глюкозы','> 200 мг/дл (11,1 ммоль/л)','1 календарный день', true,
     'http://127.0.0.1:9000/services-images/n3_glucose.png',
     'Гипергликемия — один из ранних критериев.'),

    ('№4','Определение уровня ЛДГ','> 350 Ед/л','1 календарный день', true,
     'http://127.0.0.1:9000/services-images/n4_ldh.png',
     'ЛДГ > 350 МЕ/л — критерий тяжести.'),

    ('№5','Анализ активности АСТ','> 250 Ед/л','1 календарный день', true,
     'http://127.0.0.1:9000/services-images/n5_ast.png',
     'АСТ > 250 МЕ/л — критерий тяжести.'),

    ('№6','Контроль изменения гематокрита','Падение > 10% (за 48 ч)','1 календарный день через 48 часов', true,
     'http://127.0.0.1:9000/services-images/n6_hct.png',
     'Снижение гематокрита в динамике — неблагоприятный признак.'),

    ('№7','Измерение уровня мочевины (BUN)','Повышение > 5 мг/дл','1 календарный день через 48 часов', true,
     'http://127.0.0.1:9000/services-images/n7_bun.png',
     'Рост мочевины указывает на ухудшение.'),

    ('№8','Измерение уровня кальция сыворотки','< 8,0 мг/дл (2,0 ммоль/л)','1 календарный день', true,
     'http://127.0.0.1:9000/services-images/n8_ca.png',
     'Гипокальциемия — прогностический критерий.'),

    ('№9','Измерение PaO₂','< 60 мм рт.ст.','1 календарный день', true,
     'http://127.0.0.1:9000/services-images/n9_pao2.png',
     'PaO₂ < 60 мм рт.ст. — критерий в шкале Рэнсона.'),

    ('№10','Оценка кислотно-щелочного состояния','Дефицит оснований > 4 мЭкв/л','1 календарный день через 48 часов', true,
     'http://127.0.0.1:9000/services-images/n10_acidbase.png',
     'Декомпенсированный ацидоз — неблагоприятен.'),

    ('№11','Оценка объёма секвестрированной жидкости','> 6 л','1 календарный день через 48 часов', true,
     'http://127.0.0.1:9000/services-images/n11_sequestration.png',
     'Большой объём — высокий риск осложнений.');

-- При необходимости можно деактивировать критерий явно:
-- UPDATE criteria SET is_active = false WHERE code = '№11';

-- === Примерные ЗАКАЗЫ (orders) ===
-- статусы: draft|deleted|formed|finished|rejected
INSERT INTO orders (status, creator_id, formed_at, moderator_id, computed_result)
VALUES
    ('draft',
     (SELECT id FROM users WHERE login='alice'),
     NULL, NULL, NULL),

    ('formed',
     (SELECT id FROM users WHERE login='alice'),
     NOW(),
     (SELECT id FROM users WHERE login='moder'),
     'Промежуточная сумма баллов: 3'),

    ('finished',
     (SELECT id FROM users WHERE login='bob'),
     NOW(),
     (SELECT id FROM users WHERE login='moder'),
     'Итоговая оценка тяжести: 4 балла');

-- === Позиции заказов (order_items) для примера ===
-- Привязка по code критериев через подзапрос (чтобы не знать id заранее)
-- Заказ 1 (draft) — добавим пару позиций
INSERT INTO order_items (order_id, criterion_id, position, value_num, value_indicator)
VALUES
    (
        (SELECT id FROM orders WHERE status='draft' AND creator_id=(SELECT id FROM users WHERE login='alice') ORDER BY id ASC LIMIT 1),
    (SELECT id FROM criteria WHERE code='№1'), 1, NULL, TRUE
    ),
  (
    (SELECT id FROM orders WHERE status='draft' AND creator_id=(SELECT id FROM users WHERE login='alice') ORDER BY id ASC LIMIT 1),
    (SELECT id FROM criteria WHERE code='№3'), 2, 12.3, TRUE
  );

-- Заказ 2 (formed) — три позиции
INSERT INTO order_items (order_id, criterion_id, position, value_num, value_indicator)
VALUES
    (
        (SELECT id FROM orders WHERE status='formed' AND creator_id=(SELECT id FROM users WHERE login='alice') ORDER BY id DESC LIMIT 1),
    (SELECT id FROM criteria WHERE code='№2'), 1, 17.2, TRUE
    ),
  (
    (SELECT id FROM orders WHERE status='formed' AND creator_id=(SELECT id FROM users WHERE login='alice') ORDER BY id DESC LIMIT 1),
    (SELECT id FROM criteria WHERE code='№4'), 2, 420, TRUE
  ),
  (
    (SELECT id FROM orders WHERE status='formed' AND creator_id=(SELECT id FROM users WHERE login='alice') ORDER BY id DESC LIMIT 1),
    (SELECT id FROM criteria WHERE code='№8'), 3, 1.95, TRUE
  );

-- Заказ 3 (finished) — четыре позиции
INSERT INTO order_items (order_id, criterion_id, position, value_num, value_indicator)
VALUES
    (
        (SELECT id FROM orders WHERE status='finished' AND creator_id=(SELECT id FROM users WHERE login='bob') ORDER BY id DESC LIMIT 1),
    (SELECT id FROM criteria WHERE code='№1'), 1, NULL, TRUE
    ),
  (
    (SELECT id FROM orders WHERE status='finished' AND creator_id=(SELECT id FROM users WHERE login='bob') ORDER BY id DESC LIMIT 1),
    (SELECT id FROM criteria WHERE code='№5'), 2, 310, TRUE
  ),
  (
    (SELECT id FROM orders WHERE status='finished' AND creator_id=(SELECT id FROM users WHERE login='bob') ORDER BY id DESC LIMIT 1),
    (SELECT id FROM criteria WHERE code='№9'), 3, 58, TRUE
  ),
  (
    (SELECT id FROM orders WHERE status='finished' AND creator_id=(SELECT id FROM users WHERE login='bob') ORDER BY id DESC LIMIT 1),
    (SELECT id FROM criteria WHERE code='№10'), 4, 5.1, TRUE
  );