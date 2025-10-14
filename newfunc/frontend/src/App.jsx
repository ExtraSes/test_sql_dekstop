import {useState, useEffect} from 'react';
import logo from './assets/images/logo-universal.png';
import './App.css';
import {GetAllUsers, AddUser} from "../wailsjs/go/main/App";

function App() {
    const [users, setUsers] = useState([]);
    const [loading, setLoading] = useState(false);
    const [showAddForm, setShowAddForm] = useState(false);
    
    // Форма для добавления пользователя
    const [formData, setFormData] = useState({
        name: '',
        sex: 'male',
        sumgavna: ''
    });

    // Загружаем пользователей при запуске
    useEffect(() => {
        loadUsers();
    }, []);

    const loadUsers = async () => {
        setLoading(true);
        try {
            const result = await GetAllUsers();
            setUsers(result);
        } catch (error) {
            console.error('Ошибка при загрузке пользователей:', error);
        } finally {
            setLoading(false);
        }
    };

    const handleInputChange = (e) => {
        const { name, value } = e.target;
        setFormData(prev => ({
            ...prev,
            [name]: value
        }));
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        if (!formData.name.trim() || !formData.sumgavna) {
            alert('Пожалуйста, заполните все поля');
            return;
        }

        try {
            await AddUser(formData.name, formData.sex, parseInt(formData.sumgavna));
            setFormData({ name: '', sex: 'male', sumgavna: '' });
            setShowAddForm(false);
            loadUsers(); // Перезагружаем список
        } catch (error) {
            console.error('Ошибка при добавлении пользователя:', error);
            alert('Ошибка при добавлении пользователя');
        }
    };

    return (
        <div id="App">
            <img src={logo} id="logo" alt="logo"/>
            
            <div className="header">
                <h1>Управление пользователями</h1>
                <button 
                    className="btn btn-primary" 
                    onClick={() => setShowAddForm(!showAddForm)}
                >
                    {showAddForm ? 'Отменить' : 'Добавить пользователя'}
                </button>
            </div>

            {showAddForm && (
                <div className="add-form">
                    <h2>Добавить нового пользователя</h2>
                    <form onSubmit={handleSubmit}>
                        <div className="form-group">
                            <label>Имя:</label>
                            <input
                                type="text"
                                name="name"
                                value={formData.name}
                                onChange={handleInputChange}
                                required
                            />
                        </div>
                        <div className="form-group">
                            <label>Пол:</label>
                            <select
                                name="sex"
                                value={formData.sex}
                                onChange={handleInputChange}
                            >
                                <option value="male">Мужской</option>
                                <option value="female">Женский</option>
                            </select>
                        </div>
                        <div className="form-group">
                            <label>Sumgavna:</label>
                            <input
                                type="number"
                                name="sumgavna"
                                value={formData.sumgavna}
                                onChange={handleInputChange}
                                required
                            />
                        </div>
                        <button type="submit" className="btn btn-success">Добавить</button>
                    </form>
                </div>
            )}

            <div className="users-section">
                <h2>Список пользователей</h2>
                {loading ? (
                    <div className="loading">Загрузка...</div>
                ) : users.length === 0 ? (
                    <div className="empty">Пользователи не найдены</div>
                ) : (
                    <div className="users-grid">
                        {users.map(user => (
                            <div key={user.id} className="user-card">
                                <div className="user-info">
                                    <h3>{user.name}</h3>
                                    <p><strong>Пол:</strong> {user.sex === 'male' ? 'Мужской' : 'Женский'}</p>
                                    <p><strong>Sumgavna:</strong> {user.sumgavna}</p>
                                    <p><strong>ID:</strong> {user.id}</p>
                                </div>
                            </div>
                        ))}
                    </div>
                )}
            </div>
        </div>
    )
}

export default App
