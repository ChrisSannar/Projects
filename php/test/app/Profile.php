<?php

namespace App;

use Illuminate\Database\Eloquent\Model;

/*
    This part of the application *defines* how the data is to be used on the page
    (Kind of like relationships between schemas... ya know. Like a model)
*/

class Profile extends Model
{
    public function user() {
        return $this->belongsTo(User::class);
    }
}
