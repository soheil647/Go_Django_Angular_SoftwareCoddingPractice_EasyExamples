# Generated by Django 2.2.5 on 2019-09-09 18:25

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('Scotech', '0001_initial'),
    ]

    operations = [
        migrations.AddField(
            model_name='news',
            name='disc',
            field=models.TextField(default=''),
        ),
        migrations.AddField(
            model_name='news',
            name='title',
            field=models.CharField(default=0, max_length=20),
            preserve_default=False,
        ),
    ]
