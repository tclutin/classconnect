<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
    <!-- Bootstrap CSS -->
    <link href="https://stackpath.bootstrapcdn.com/bootstrap/4.5.2/css/bootstrap.min.css" rel="stylesheet">
    <style>
        .nav-link {
            transition: background-color 0.3s ease-in-out;
        }
        .nav-link:hover {
            background-color: #6c757d !important;
        }
        .active .nav-link {
            background-color: #6c757d !important;
        }
        .logout-btn {
            color: #fff;
            background-color: #dc3545;
            border: none;
            border-radius: 5px;
            padding: 5px 10px;
            transition: background-color 0.3s ease-in-out;
        }
        .logout-btn:hover {
            background-color: #c82333;
        }
    </style>
</head>
<body>
@php use Illuminate\Support\Facades\Route; @endphp
<nav class="navbar navbar-expand-lg navbar-light bg-light">
    <a class="navbar-brand" href="#">ClassConnect</a>
    <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarNav"
            aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
    </button>
    <div class="collapse navbar-collapse" id="navbarNav">
        <ul class="navbar-nav mr-auto">
            <li class="nav-item">
                <a class="nav-link {{ Route::is('me') ? 'active' : '' }}" href="{{ route('me') }}">Me</a>
            </li>
            <li class="nav-item">
                <a class="nav-link {{ Route::is('groups.index') ? 'active' : '' }}" href="{{ route('groups.index') }}">Groups</a>
            </li>
            <li class="nav-item">
                <a class="nav-link {{ Route::is('groups.create_form') ? 'active' : '' }}" href="{{ route('groups.create_form') }}">Create Group</a>
            </li>
            <li class="nav-item">
                <a class="nav-link {{ Route::is('schedules.upload_form') ? 'active' : '' }}" href="{{ route('schedules.upload_form') }}">Create Schedule</a>
            </li>
        </ul>
        <ul class="navbar-nav">
            @if(session()->has('access_token'))
            <li class="nav-item">
                <form action="{{ route('logout') }}" method="POST">
                    @csrf
                    <button type="submit" class="btn btn-link logout-btn">Logout</button>
                </form>
            </li>
            @else
            <li class="nav-item">
                <a class="nav-link" href="{{ route('login') }}">Login</a>
            </li>
            <li class="nav-item">
                <a class="nav-link" href="{{ route('signup') }}">Sign Up</a>
            </li>
            @endif
        </ul>
    </div>
</nav>

<div class="container mt-4">
    @yield('content')
</div>

<!-- Bootstrap JS -->
<script src="https://code.jquery.com/jquery-3.5.1.slim.min.js"></script>
<script src="https://cdn.jsdelivr.net/npm/@popperjs/core@2.5.4/dist/umd/popper.min.js"></script>
<script src="https://stackpath.bootstrapcdn.com/bootstrap/4.5.2/js/bootstrap.min.js"></script>
</body>
</html>
