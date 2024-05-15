<!-- resources/views/me.blade.php -->

@extends('layouts.app')
@php
    use Illuminate\Support\Facades\Route;
@endphp
@section('content')
    <div class="container">
        <div class="row justify-content-center">
            <div class="col-md-8">
                <div class="card">
                    <div class="card-header">{{ __('User Information') }}</div>

                    <div class="card-body">
                        <p>{{ __('Username') }}: {{ $user['username'] }}</p>
                        <p>{{ __('Email') }}: {{ $user['email'] }}</p>
                        @if ($user['group'])
                            <p>{{ __('Group Name') }}: {{ $user['group']['name'] }}</p>
                            <p>{{ __('Group Code') }}: {{ $user['group']['code'] }}</p>
                            <p>{{ __('Group Members Count') }}: {{ $user['group']['members_count'] }}</p>
                            <p>{{ __('Group Created At') }}: {{ $user['group']['created_at'] }}</p>
                        @endif
                    </div>
                </div>
            </div>
        </div>
    </div>
@endsection
