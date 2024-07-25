from models import db, User
from werkzeug.security import generate_password_hash, check_password_hash

class UserService:
    def create_user(self, username, password):
        hashed_password = generate_password_hash(password, method='pbkdf2:sha256')  # Poprawiona metoda
        new_user = User(username=username, password=hashed_password)
        db.session.add(new_user)
        db.session.commit()
    
    def authenticate(self, username, password):
        user = User.query.filter_by(username=username).first()
        if user and check_password_hash(user.password, password):
            return user
        return None