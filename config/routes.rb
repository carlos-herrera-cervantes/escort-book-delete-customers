# frozen_string_literal: true

Rails.application.routes.draw do
  get 'api/v1/health' => 'application#health'
end
