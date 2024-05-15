<!-- index.blade.php -->

@extends('layouts.app')

@section('content')
    <div class="container">
        <div class="row justify-content-center">
            <div class="col-md-8">
                <div class="card">
                    <div class="card-header">Group List</div>

                    <div class="card-body">
                        @if ($groups->isEmpty())
                            <p>No groups found.</p>
                        @else
                            <ul>
                                @foreach ($groups as $group)
                                    <li>{{ $group->name }}</li>
                                @endforeach
                            </ul>
                        @endif
                    </div>
                </div>
            </div>
        </div>
    </div>
@endsection
