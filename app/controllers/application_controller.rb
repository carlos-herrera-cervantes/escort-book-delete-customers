# frozen_string_literal: true

# Main application controller
class ApplicationController < ActionController::API
  def health
    render json: { message: 'OK', date: DateTime.now }
  end
end
