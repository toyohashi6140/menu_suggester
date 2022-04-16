import { useState } from 'react';
import sampleImg from '../img/sample.jpg';
const Dish = ({alt}) => {
    const [recipe, setRecipe] = useState(0);
    return (
        <div className='dish'>
            <img src={sampleImg} alt={alt}/>
            <button 
                className="btn" 
                onClick={
                    () => {
                        setRecipe(recipe+1);  
                    }
                }
            >
                レシピを取得する
            </button>
        </div>
    )
}
export default Dish