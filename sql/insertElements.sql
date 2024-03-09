INSERT INTO BAND(bandName, genre, youtube, instagram, tunecore, twitter, imagePath) VALUES ('新亜並行空間', 'ロック', 'https://www.youtube.com/channel/UCUGsHePFZpv4aR2jJzon5BA', 'https://www.instagram.com/neo_a_parallel/', 'https://www.tunecore.co.jp/artists/N.A.Parallel?lang=ja', 'https://twitter.com/Neo_A_Parallel', './images/新亜並行空間.png');
INSERT INTO BAND(bandName, genre, youtube, instagram, twitter, imagePath) VALUES ('東野ルーカス', 'ジャズ', 'https://www.youtube.com/@lucashigashino5535', 'https://www.instagram.com/lucashigashino/', 'https://twitter.com/Lucas22982914', './images/東野ルーカス.png');
INSERT INTO BAND(bandName, genre, youtube, instagram, tunecore, twitter, homePage, imagePath) VALUES ('Supernova Awakening', 'メタル', 'https://www.youtube.com/c/SupernovaAwakeningOfficial', 'https://www.instagram.com/supernovaawakening', 'https://www.tunecore.com/music/supernovaawakening', 'https://twitter.com/SuprNovaAwakens', 'https://twitter.com/SuprNovaAwakens	https://www.supernovaawakening.net', './images/SupernovaAwakening.png');

INSERT INTO livehouse (livehouseName, longitude, latitude, homePage) VALUES('The Red Curtain', 35.0237056, 135.7545413, 'https://theredcurtain.live');
INSERT INTO livehouse (livehouseName, longitude, latitude, homePage) VALUES('Divine Cathedral', 35.0237009, 135.7579916, 'https://divinecathedral.xyz');
INSERT INTO livehouse (livehouseName, longitude, latitude, homePage) VALUES ('Groove Foundry', 35.0172265, 135.7590776, 'https://groovefoundry.xyz');

INSERT INTO event (eventName, livehouseId, eventDate ,openTime, startTime, fee) VALUES ('夜空の耀き', 2, '2023-03-15' ,'18:00', '21:00', 2000);
INSERT INTO event (eventName, livehouseId, eventDate ,openTime, startTime, fee) VALUES ('絶唱の宴', 3,  '2023-03-11' ,'15:00', '18:00', 2200);
INSERT INTO event (eventName, livehouseId, eventDate ,openTime, startTime, fee) VALUES ('混沌の咆哮', 1, '2023-03-20' ,'16:00', '19:00', 2500);

INSERT INTO paticipant (eventId, bandId) VALUES (1, 1);
INSERT INTO paticipant (eventId, bandId) VALUES (1, 2);
INSERT INTO paticipant (eventId, bandId) VALUES (2, 2);
INSERT INTO paticipant (eventId, bandId) VALUES (2, 3);
INSERT INTO paticipant (eventId, bandId) VALUES (3, 1);
INSERT INTO paticipant (eventId, bandId) VALUES (3, 3);
