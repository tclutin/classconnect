<!-- resources/views/includes/header.blade.php -->

<nav class="navbar navbar-expand-lg navbar-dark bg-dark">
    <div class="container-fluid">
        <a class="navbar-brand" href="{{ route('me') }}">Your App</a>
        <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNav"
            aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
            <span class="navbar-toggler-icon"></span>
        </button>
        <div class="collapse navbar-collapse" id="navbarNav">
            <ul class="navbar-nav">
                <li class="nav-item">
                    <a class="nav-link" href="{{ route('me') }}">Me</a>
                </li>
                <li class="nav-item">
                    <a class="nav-link" href="{{ route('groups.index') }}">Groups</a>
                </li>
                <li class="nav-item">
                    <a class="nav-link" href="{{ route('groups.create_form') }}">Create Group</a>
                </li>
            </ul>
            <!-- Перемещаем кнопку Log Out вправо -->
            <ul class="navbar-nav ml-auto">
                <li class="nav-item">
                    <form method="POST" action="{{ route('logout') }}">
                        @csrf
                        <button type="submit" class="btn btn-outline-danger">Log Out</button>
                    </form>
                </li>
            </ul>
        </div>
    </div>
</nav>
