from django.shortcuts import render, redirect
from django.contrib import messages
from django.contrib.auth.models import User, auth
# Create your views here.


def logout(request):
    auth.logout(request)
    return redirect('/index.html')


def login(request):
    if request.method == 'POST':
        name = request.POST['name']
        password = request.POST['password']
        email = request.POST['email']

        user = auth.authenticate(username=name, password=password, email=email)

        if user is not None:
            auth.login(request, user)
            return redirect('/index.html')
        else:
            messages.info(request, 'invalid Login')
            return redirect("login")

    else:
        return render(request, "login.html")


def register(request):

    if request.method == 'POST':
        name = request.POST['name']
        password = request.POST['password']
        email = request.POST['email']
        if User.objects.filter(username=name).exists():
            messages.info(request, 'UserName Taken')
            return redirect('/account/register')
        elif User.objects.filter(email=email).exists():
            messages.info(request, 'Email Taken')
            return redirect('/account/register')
        elif name == "":
            messages.info(request, 'UserName Cant Be Null')
            return redirect('/account/register')
        elif email == "":
            messages.info(request, 'Email Cant be Null')
            return redirect('/account/register')
        elif password == "":
            messages.info(request, 'Password cant be Null')
            return redirect('/account/register')
        else:
            user = User.objects.create_user(username=name, password=password, email=email)
            user.save
            auth.login(request, user)
            return redirect('/index.html')
    else:
        return render(request, 'register.html')




