
from django.urls import path
from . import views

urlpatterns = [
    path('', views.hello, name='home'),
    path('index.html', views.home, name='index'),
    path('review.html', views.review, name='Games'),
    path('categories.html', views.categories, name='Blog'),
    path('community.html', views.community, name='Forums'),
    path('contact.html', views.contact, name='Contact'),
]
