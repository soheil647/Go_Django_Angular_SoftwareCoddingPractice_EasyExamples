class game_bombs(models.Model):
    username = models.CharField(max_length=30)
    scoin = models.IntegerField()
    bomb = models.ForeignKey('game_bombs_data', on_delete=models.CASCADE)
    player_rank = models.IntegerField()
    rate = models.IntegerField()


class game_bombs_data(models.Model):
    time_start = models.DateTimeField(default=one_day_hence())
    time_end = models.DateTimeField(default=one_day_hence())
    title = models.CharField(max_length=30)
    link = models.CharField(max_length=30, default='1')
    award = models.CharField(max_length=50, default='1')
    capacity = models.IntegerField(default=1)
    cost = models.CharField(max_length=30, default='a')
    description = models.CharField(max_length=300)


class event_game(models.Model):
    event = models.ForeignKey('game_bombs_data', on_delete=models.CASCADE)
    game = models.ForeignKey('Games', on_delete=models.CASCADE)
    username = models.CharField(max_length=30, default='aicam')
    score = models.IntegerField(max_length=30, default=0)


def getEventId(userID, game_id, score):
    events_ids = []
    final_events = []
    x = []
    for row_data in event_game.objects.filter(game=game_id):
        x.append(row_data.event.id)

    x = list(set(x))
    for i in x:
        if event_game.objects.filter(game=game_id, username=userID, event=game_bombs_data.objects.get(id=i)).count() < 1:
            current_time = timezone.now()
            t = game_bombs_data.objects.get(id=i)
            if t.time_start < current_time < t.time_end:
                event_game(event=game_bombs_data.objects.get(id=i), game=Games.objects.get(id=game_id), username=userID,score=score).save()

    for row_data in event_game.objects.filter(game=game_id, username=userID):
        events_ids.append(row_data.event.id)

    for eid in events_ids:
        if game_bombs.objects.filter(username=userID, bomb=eid).count() > 0:
                current_time = timezone.now()
                t = game_bombs_data.objects.get(id=eid)
                if t.time_start < current_time < t.time_end:
                    final_events.append(eid)

    for eid in event_game.objects.all():
        for fid in final_events:
            if eid.event.id == fid:
                row_data = event_game.objects.get(event=eid.event, game=Games.objects.get(id=int(game_id)),
                                                  username=userID)
                # row_dataa = game_bombs.objects.get(username=userID, bomb=eid.event)
                if row_data.score < int(score):
                    # row_dataa.rate += (int(score) - row_data.score)
                    row_data.score = int(score)
                    row_data.save()
                    # row_dataa.save()

    for fid in final_events:
        row_dataa = game_bombs.objects.get(username=userID, bomb=game_bombs_data.objects.get(id=fid))
        row_dataa.rate = 0
        for row_data in event_game.objects.filter(username=userID, event=game_bombs_data.objects.get(id=fid)):
            row_dataa.rate += row_data.score
        row_dataa.save()

    return final_events


def rank(events_id):
    for events in events_id:
        user_score_in_event = []
        for row_data in game_bombs.objects.filter(bomb=game_bombs_data.objects.get(id=events)):
            user_score_in_event.append(row_data)

        for i in range(0, len(user_score_in_event)):
            for j in range(0, len(user_score_in_event)):
                if user_score_in_event[i].rate < user_score_in_event[j].rate:
                    if user_score_in_event[j].player_rank > user_score_in_event[i].player_rank:
                        t = user_score_in_event[i].player_rank
                        user_score_in_event[i].player_rank = user_score_in_event[j].player_rank
                        user_score_in_event[j].player_rank = t

        for i in range(0, len(user_score_in_event)):
            user_score_in_event[i].player_rank = i + 1
            user_score_in_event[i].save()
