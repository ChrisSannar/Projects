@extends('layouts.app')

@section('content')
<div class="container">
    <div class="row justify-content-center">
        <div class="col-md-8">
            <div class="card">
                <div class="card-header">{{ __('Dashboard') }}</div>

                <div class="card-body">
                    @if (session('status'))
                        <div class="alert alert-success" role="alert">
                            {{ session('status') }}
                        </div>
                    @endif
                    <h1>{{ $user->username }}</h1>
                    <h2>{{ $user->profile->title }}</h2>
                    <div>{{ $user->profile->description }}</div>
                </div>
            </div>
        </div>
    </div>
</div>
@endsection
