from django.shortcuts import render
from django.http import HttpResponse
from .models import Games, News
# Create your views here.


def hello(request):
    return render(request, 'home.html', {'name': 'Iran'})


def home(request):
    news = News.objects.all()
    RecetGame = Games.objects.all()
    print(news)

    return render(request, 'index.html', {'news': news, 'Recengame': RecetGame})


def review(request):
    return render(request, 'review.html')


def categories(request):
    return render(request, 'categories.html')


def community(request):
    return render(request, 'community.html')


def contact(request):
    return render(request, 'contact.html')




