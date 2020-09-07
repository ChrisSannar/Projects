<?php

namespace App\Http\Resources;

use Illuminate\Http\Resources\Json\JsonResource;

/*
    A layer with a set of rules as to what the API returns
    Typically used to format the data and limit what is returned from the api
*/

class TestResource extends JsonResource
{
    /**
     * Transform the resource into an array.
     *
     * @param  \Illuminate\Http\Request  $request
     * @return array
     */
    public function toArray($request)
    {
        return [
            'id' => $this->id,
            'text' => $this->text,
            'completed' => $this->completed
        ];
        // return parent::toArray($request);
    }
}
