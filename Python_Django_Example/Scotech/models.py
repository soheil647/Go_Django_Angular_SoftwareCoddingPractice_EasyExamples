from django.db import models

# Create your models here.


class Games(models.Model):
    title = models.CharField(max_length=20)
    name = models.CharField(max_length=20)
    img = models.ImageField(upload_to='game_pics')
    description = models.TextField()
    comment = models.IntegerField()
    gameLink = models.CharField(max_length=50)
    commentLink = models.CharField(max_length=50)


class News(models.Model):
    title = models.CharField(max_length=20)
    disc = models.TextField()



